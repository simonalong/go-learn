package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {

	fmt.Println(time.ParseDuration("12s"))

	//randTime, err := time.ParseDuration(isc.ToString(rand.Int63n(60 * 1000)) + "ms")
	// time.Sleep(randTime)
}
