package main

import (
	"fmt"
	"github.com/goburrow/cache"
	"time"
)

func main() {
	load := func(k cache.Key) (cache.Value, error) {
		time.Sleep(100 * time.Millisecond) // Slow task
		return fmt.Sprintf("%d", k), nil
	}
	// Create a loading cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(100),                 // Limit number of entries in the cache.
		cache.WithExpireAfterAccess(1*time.Minute), // Expire entries after 1 minute since last accessed.
		cache.WithRefreshAfterWrite(2*time.Minute), // Expire entries after 2 minutes since last created.
	)

	//getTicker := time.Tick(100 * time.Millisecond)
	//reportTicker := time.Tick(5 * time.Second)
	//for {
	//	select {
	//	case <-getTicker:
	//		_, _ = c.Get(rand.Intn(200))
	//	case <-reportTicker:
	//		st := cache.Stats{}
	//		c.Stats(&st)
	//		fmt.Printf("%+v\n", st)
	//	}
	//}

	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
	time.Sleep(2 * time.Second)
	fmt.Println(c.Get(12))
	fmt.Println(c.Get(12))
}
