package test

import (
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestRate(t *testing.T) {
	limiter := rate.NewLimiter(10, 10)

	for i := 0; i < 100; i++ {
		if limiter.Allow() {
			//time.Sleep(100)
			println("ok", i)
		} else {
			time.Sleep(time.Second)
		}
	}
}
