package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTSTP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-signalChan
		fmt.Fprintln(os.Stdout, "\rsignal received:", sig.String())
		fmt.Fprintln(os.Stdout, "execution time:", time.Since(startTime).String())
		close(signalChan)
		os.Exit(1)
	}()

	time.Sleep(time.Second * 60)
}
