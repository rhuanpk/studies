//go:build channels

/*
	Create some program that:
		- what is concurrency?
			R: concurrency, do lots of things at once. That’s different from doing lots of things at the same time Channels.
		- when do you need channels?
			R: to communicate between goroutines.
		- how can you send data into a channel?
			R: `channel <- "message"`
		- how can you read data from a channel?
			R: `message := <- channel`

	source: https://golangr.com/channels/
*/

package main

import "fmt"

func function_1(channel chan string) {
    channel <- "function_1() was here"
}

func function_2(channel chan string) {
    message := <- channel
    fmt.Println("f2",message)
}


func main() {
   var channel chan string = make(chan string)
   go function_1(channel)
   go function_2(channel)
   fmt.Scanln()
}
