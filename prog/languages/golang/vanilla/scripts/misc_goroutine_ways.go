package main

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

const concurrency = 3

// semaphore itera sobre os dados e cria uma goroutine para cada dado a ser processado
func semaphore() {
	trafficlight := make(chan struct{}, concurrency)
	payload := payload()

	var wg sync.WaitGroup
	for index, data := range payload {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer func() { <-trafficlight }()
			trafficlight <- struct{}{}

			fmt.Printf("process %d: %v\n", index, data)
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
}

// workers cria a quantidade máxima de goroutines e elas leem os dados a serem processados
func workers() {
	payload := payload()
	dataChan := make(chan string, len(payload))

	go func() {
		for _, data := range payload {
			dataChan <- data
		}
		close(dataChan)
	}()

	var wg sync.WaitGroup
	for index := range concurrency {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for data := range dataChan {
				fmt.Printf("process %d: %v\n", index, data)
				time.Sleep(time.Second)
			}
		}()
	}
	wg.Wait()
}

// go playground: https://go.dev/play/p/jvJCKPj7zIR
func main() {
	println("trafficlight")
	semaphore()
	println("workers")
	workers()
}

func payload() []string {
	var payload []string
	for range 10 {
		payload = append(payload, rand.Text())
	}
	return payload
}
