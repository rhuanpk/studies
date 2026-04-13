package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("error in open file:", err)
	}
	defer file.Close()

	log.SetOutput(file)
	for index := 0; index < 5; index++ {
		fmt.Fprintf(file, "index %d) ", index)
		log.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	}
	fmt.Fprintln(file)
}
