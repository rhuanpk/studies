package main

import (
	"fmt"
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

func first() {
	ch := make(chan int) // chan buffer = 0
	fmt.Println("chan memory address:", ch)

	// // infinitily waiting because chan does not have scape in buffer to
	// // allocate nothing
	// ch <- 1

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// fmt.Println("chan value:", <-ch)

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// _, ok := <-ch
	// fmt.Println("chan ok:", ok)

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// value, ok := <-ch
	// fmt.Println("chan returns:", value, ok)
}

func second() {
	ch := make(chan int, 1) // chan buffer = 1
	fmt.Println("chan memory address:", ch)

	ch <- 1

	// // infinitily waiting because chan already reached maximum buffer capacity
	// ch <- 1

	// fmt.Println("chan value:", <-ch)

	// _, ok := <-ch
	// fmt.Println("chan ok:", ok)

	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// fmt.Println("chan value:", <-ch)

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// _, ok = <-ch
	// fmt.Println("chan ok:", ok)

	// // infinitily waiting because chan does not have nothing in buffer to
	// // consume
	// value, ok = <-ch
	// fmt.Println("chan returns:", value, ok)
}

func third() {
	ch := make(chan int, 2) // chan buffer = 2
	fmt.Println("chan memory address:", ch)

	ch <- 1

	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)

	ch <- 1

	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)
}

func fourth() {
	ch := make(chan int, 2) // chan buffer = 2
	fmt.Println("chan memory address:", ch)

	ch <- 1
	ch <- 1

	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)

	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)
}

func fifth() {
	ch := make(chan int, 2) // chan buffer = 2
	fmt.Println("chan memory address:", ch)

	ch <- 1

	close(ch)

	// // panics because chan is closed to allocate new values
	// ch <- 1

	// despite chan already closed not yet empty so ok value is true
	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)

	// because chan is closed AND empty returns the zero value of chan and false
	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)

	// because chan is closed AND empty returns the zero value of chan and false
	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)
}

func sixth() {
	ch := make(chan int, 2) // chan buffer = 2
	fmt.Println("chan memory address:", ch)

	ch <- 1
	ch <- 1

	close(ch)

	// despite chan already closed not yet empty so ok value is true
	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)

	// despite chan already closed not yet empty so ok value is true
	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)

	// because chan is closed AND empty returns the zero value of chan and false
	value, ok = <-ch
	fmt.Println("chan returns:", value, ok)
}

func seventh() {
	ch := make(chan int) // chan buffer = 0
	fmt.Println("chan memory address:", ch)

	close(ch)

	// closed chans are not blockers
	value, ok := <-ch
	fmt.Println("chan returns:", value, ok)
}

// go playground: https://go.dev/play/p/lIBsgtxzsh2
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
}
