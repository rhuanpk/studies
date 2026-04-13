package main

import "time"

func main() {
	defer func(start time.Time) {
		println("duration:", time.Since(start).String())
	}(time.Now())
	// start := time.Now()

	var xpto []string
	for index := 0; index < 999999; index++ {
		xpto = append(xpto, "hello world")
		time.Sleep(time.Millisecond * 250)
	}

	// println("duration:", time.Since(start).String())
}
