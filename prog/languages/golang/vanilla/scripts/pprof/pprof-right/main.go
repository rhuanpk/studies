package main

import (
	"dev/pkg"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
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
		channel chan bool
		wg      sync.WaitGroup
		ints    []int
		sum     int
	)

	// print("press <enter> to continue ")
	// bufio.NewReader(os.Stdin).ReadString('\n')
	time.Sleep(time.Second * 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for round := 0; round < 3; round++ {
			pkg.CreazySum(start, end)
		}
	}()

	for round := 0; round < 5; round++ {
		for index := start; index < end; index++ {
			ints = append(ints, index)
		}
	}
	for range ints {
		sum++
	}
	println("total:", sum)
	pkg.PprofHeap()

	wg.Wait()
	println("finish")
	<-channel
}
