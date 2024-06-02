package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// statements
			if i > 1 {
				println("returning")
				return
			}
			println("index", i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	println("finished")
}
