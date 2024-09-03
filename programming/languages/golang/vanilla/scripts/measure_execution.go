package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof"
)

var memory runtime.MemStats

func init() {
	runtime.ReadMemStats(&memory)
	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
}

// to make graphs in Debian run `apt install graphviz` and to run program:
// go run ./ & sleep 1; go tool pprof 'http://localhost:6060/debug/pprof/{profile|heap|goroutine}[?seconds=<seconds>]'
func main() {
	defer func(start time.Time) {
		fmt.Printf("alloced memory: %dMiB\n", memory.Alloc/1024^2)
		fmt.Printf("total alloced memory: %dMiB\n", memory.TotalAlloc/1024^2)
		fmt.Printf("os memory obtained: %dMiB\n", memory.Sys/1024^2)
		fmt.Printf("garbage collector cycles: %dc\n", memory.NumGC)
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
