package main

import (
	"dev/pkg"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

const (
	start = 0
	end   = 999999
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	// runtime.SetCPUProfileRate(1)
}

func main() {
	var (
		ints []int
		sum  int
	)

	go pkg.CreazySum(start, end)

	for index := start; index < end; index++ {
		ints = append(ints, index)
	}
	for range ints {
		sum++
	}
	println("total:", sum)

	log.Println(http.ListenAndServe("localhost:6060", nil))
}
