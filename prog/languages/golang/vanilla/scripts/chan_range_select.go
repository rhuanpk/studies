package main

import (
	"sync"
	"time"
)

func init() {
	// only for do not get deadlock while running
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

// synchronous
func first() {
	var count int

	for index := 0; index <= 99; index++ {
		count += index
	}

	println("count:", count)
}

// asynchronous with DRC (data race condition)
func second() {
	var count int

	for index := 0; index <= 99; index++ {
		go func() {
			count += index
		}()
	}
	time.Sleep(time.Millisecond)

	println("count:", count)
}

// asynchronous with DRC (data race condition)
func third() {
	var count int

	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count += index
		}()
	}
	wg.Wait()

	println("count:", count)
}

// range chans
func fourth() {
	ch := make(chan int, 5)

	for index := range 5 {
		ch <- index
	}
	close(ch)

	// range automatically sees the ok value of chan to know when breaks
	for value := range ch {
		println("value:", value)
	}
}

// synchronous again
func fifth() {
	var count int
	ch := make(chan int, 100)

	// safe async
	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- index
		}()
	}
	wg.Wait()
	close(ch) // needed for the range

	// but sync again
	for value := range ch {
		count += value
	}

	println("count:", count)
}

// synchronous again
func sixth() {
	var count int
	ch := make(chan int, 100)

	// safe async
	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- index
		}()
	}
	wg.Wait() // needed for the default

loop:
	// but sync again
	for {
		select {
		case value := <-ch:
			count += value
		default:
			break loop
		}
	}

	println("count:", count)
}

// safe asynchronous
func seventh() {
	var count int
	ch := make(chan int, 100)

	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- index
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

loop:
	// better for multi chans control
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				break loop
			}
			count += value
		}
	}

	println("count:", count)
}

// safe asynchronous
func eighth() {
	var count int
	ch := make(chan int, 100)

	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- index
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// better for multi chans control
	var skip bool
	for !skip {
		select {
		case value, ok := <-ch:
			if !ok {
				skip = true
			}
			count += value
		}
	}

	println("count:", count)
}

// safe asynchronous
func ninth() {
	var count int
	ch := make(chan int, 100)

	var wg sync.WaitGroup
	for index := 0; index <= 99; index++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- index
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// better for single chan control
	for value := range ch {
		count += value
	}

	println("count:", count)
}

// go playground: https://go.dev/play/p/N1Xb3SJHrlB
func main() {
	first()
	println()
	second()
	println()
	third()
	println()
	fourth()
	println()
	fifth()
	println()
	sixth()
	println()
	seventh()
	println()
	eighth()
	println()
	ninth()
}
