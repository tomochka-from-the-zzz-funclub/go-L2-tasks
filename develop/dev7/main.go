package main

import (
	"fmt"
	"time"
)

// or merge one or more done-channels into a single-channel if one of its representatives closes.
func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			for {
				select {
				case <-channel:
					close(single)
				case <-single:
					return
				}
			}
		}(channel)
	}
	return single
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))

}
