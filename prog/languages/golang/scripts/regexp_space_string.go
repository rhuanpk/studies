package main

import (
	"fmt"
	"regexp"
	"strings"
)

// go playground: https://go.dev/play/p/gyM2POjZh1-
func main() {
	str := "hello\nworld"
	println("string:", str)

	regex := regexp.MustCompile(`\\n.*$`).ReplaceAllString(str, "")
	println("\nregex:", regex)

	raw := regexp.MustCompile(`\\n.*$`).ReplaceAllString(strings.Trim(fmt.Sprintf("%#v", str), `"`), "")
	println("\nraw:", raw)
}
