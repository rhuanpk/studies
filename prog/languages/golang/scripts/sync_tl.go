package main

import (
	"sync"
	"time"
)

const delay = time.Millisecond * 500

// go playground: https://go.dev/play/p/pMIuiUtnpFg
func main() {
	var wg sync.WaitGroup
	tl := make(chan byte, 2) // up max to 2
	for index := range 10 {
		wg.Add(1)
		tl <- 0
		go func(index int) {
			defer func() {
				wg.Done()
				<-tl
			}()
			println("index:", index)
			time.Sleep(delay)
		}(index)
	}
	wg.Wait()
	close(tl)
	println("finish")
}
