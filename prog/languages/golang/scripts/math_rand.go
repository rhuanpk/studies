package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go playground: https://go.dev/play/p/n9d_RD6GuDt
func main() {
	println("math/rand")

	// generate a new data sorce every time that request random numbers
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("random.Int():", random.Int())
	fmt.Println("random.Intn(42):", random.Intn(42))
	fmt.Println("random.Float64():", random.Float64())

	// println("rand.Int():", strconv.Itoa(rand.Int()))
	// println("rand.Intn(42):", strconv.Itoa(rand.Intn(42)))
	// println("rand.Float64():", strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
}
