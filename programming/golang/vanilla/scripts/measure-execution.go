package main

import "time"

func main() {
	defer func(start time.Time) {
		println("duration:", time.Since(start).String())
	}(time.Now())
	time.Sleep(time.Second * 3)
}
