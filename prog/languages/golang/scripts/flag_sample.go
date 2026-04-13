package main

import (
	"flag"
	"fmt"
)

// build: go build [-ldflags '-s -w'] ./
// help: ./golang -h
// usage: ./golang [-b 1] [-f xpto]
func main() {
	var foo string
	var bar int
	flag.StringVar(&foo, "f", "foo", "foo flag description")
	flag.IntVar(&bar, "b", 0, "bar flag description")
	flag.Parse()
	fmt.Printf("foo: %s\nbar: %d\n", foo, bar)
}
