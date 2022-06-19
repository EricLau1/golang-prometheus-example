package job

import (
	"flag"
	"golang-prometheus-example/app/cache"
	"golang-prometheus-example/app/random"
	"log"
	"time"
)

const defaultTimeout = time.Minute * 2

var (
	timeout time.Duration
)

func init() {
	flag.DurationVar(&timeout, "timeout", defaultTimeout, "max time to complete job")
	flag.Parse()
}

func Run() {
	var (
		c     = cache.New()
		start = time.Now()
	)

	log.Println("Starting job...")
	log.Println("Stop at:", start.Add(timeout).String())

	for {
		delay := random.Intn(3000)

		for i := 0; i < 1000; i++ {
			data := create()
			c.Set(data.ID, data.Bytes())
		}

		cache.Push()

		if time.Since(start) >= timeout {
			break
		}

		time.Sleep(time.Millisecond * time.Duration(delay))
	}

	if c.Mb() > 0 {
		log.Printf("%2.fMb cached\n", c.Mb())
		c.Reset()
		cache.Push()
	}
	log.Println("Job finished.")
}
