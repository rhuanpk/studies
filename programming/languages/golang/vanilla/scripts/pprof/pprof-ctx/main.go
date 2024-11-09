package main

import (
	"context"
	"dev/pkg"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
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

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	go func() {
		cmd := exec.CommandContext(ctx,
			"go", "tool", "pprof", "http://localhost:6060/debug/pprof/profile",
		)
		if err := cmd.Start(); err != nil {
			fmt.Println("pprof error:", err)
		}
		go cmd.Wait()
		// cmd.Cancel()
		// cmd.Process.Kill()
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
	cancel()
}
