package cache

import "time"

type Item struct {
	Data       []byte
	Expiration int64
}

func (i *Item) IsExpired(t time.Time) bool {
	return t.After(time.Unix(i.Expiration, 0))
}
