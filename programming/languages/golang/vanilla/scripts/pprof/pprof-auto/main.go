package main

import (
	"dev/pkg"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
	"runtime"
	"sync"
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
	go http.ListenAndServe("localhost:6060", nil)
}

func main() {
	var (
		wg   sync.WaitGroup
		ints []int
		sum  int
	)

	go func() {
		wg.Add(1)
		if err := exec.Command(
			"go", "tool", "pprof", "http://localhost:6060/debug/pprof/profile",
		).Run(); err != nil {
			fmt.Println("pprof error:", err)
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		pkg.CreazySum(start, end)
		wg.Done()
	}()

	for index := start; index < end; index++ {
		ints = append(ints, index)
	}
	for range ints {
		sum++
	}
	println("total:", sum)

	wg.Wait()
}
