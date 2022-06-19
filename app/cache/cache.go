package cache

import (
	"golang-prometheus-example/app/random"
	"log"
	"sync"
	"time"
)

const (
	defaultExpirations  = time.Second * 10
	defaultCleanerDelay = time.Second * 5
)

type Cache interface {
	Set(key string, value []byte)
	Get(key string) (value []byte)
	Del(key string)
	Mb() float64
	ForEach(each func(k string, v []byte) bool)
	Reset()
}

type cache struct {
	data        map[string]Item
	lk          *sync.Mutex
	memory      float64
	expirations time.Duration
}

func New() Cache {
	c := &cache{
		data:        make(map[string]Item),
		lk:          &sync.Mutex{},
		memory:      0,
		expirations: defaultExpirations,
	}
	cleaner(c, defaultCleanerDelay)
	return c
}

func (c *cache) onLock(rw func()) {
	c.lk.Lock()
	rw()
	c.lk.Unlock()
}

func (c *cache) Set(key string, value []byte) {
	c.onLock(func() {
		expirations := random.Intn(int(c.expirations))
		expiration := time.Now().Add(time.Second * time.Duration(expirations)).Unix()
		c.data[key] = Item{Data: value, Expiration: expiration}
		c.memory += float64(len(value))
		memoryGauge.Set(c.Mb())
	})
}

func (c *cache) Get(key string) (value []byte) {
	c.onLock(func() {
		if item, ok := c.data[key]; ok {
			value = item.Data
		}
	})
	return
}

func (c *cache) Del(key string) {
	c.onLock(func() {
		c.del(key)
	})
}

func (c *cache) del(key string) {
	if item, ok := c.data[key]; ok {
		delete(c.data, key)
		c.memory -= float64(len(item.Data))
		memoryGauge.Set(c.Mb())
	}
}

func (c *cache) Mb() float64 {
	// convert bytes to mega bytes
	return c.memory / 1024 / 1024
}

func (c *cache) ForEach(each func(k string, v []byte) bool) {
	c.onLock(func() {
		for key, item := range c.data {
			if !each(key, item.Data) {
				break
			}
		}
	})
}

func (c *cache) Reset() {
	c.onLock(func() {
		for key := range c.data {
			c.del(key)
		}
	})
}

func cleaner(c *cache, delay time.Duration) {
	go func() {
		for {
			c.onLock(func() {
				log.Println("Starting cleanup cache...")
				log.Printf("Usage: %.2fMb\n", c.Mb())
				deleted := 0
				for key, item := range c.data {
					if item.IsExpired(time.Now()) {
						c.del(key)
						deleted++
					}
				}
				log.Println("Cleanup finished.")
				log.Printf("%d items removed.\n", deleted)
				log.Printf("Usage: %.2fMb\n", c.Mb())
			})
			time.Sleep(delay)
		}
	}()
}
