package main

import (
	"fmt"
	"regexp"
)

// go playground: https://go.dev/play/p/faq8g5aECLp
func main() {
	const (
		regex = `/[^/]+`
		str   = "/foo/bar"
	)
	println("REGEX:", regex, "\nSTRING:", str)
	fmt.Printf("FindAllString ................: %#v\n", regexp.MustCompile(regex).FindAllString(str, -1))
	fmt.Printf("FindAllStringIndex ...........: %#v\n", regexp.MustCompile(regex).FindAllStringIndex(str, -1))
	fmt.Printf("FindAllStringSubmatch ........: %#v\n", regexp.MustCompile(regex).FindAllStringSubmatch(str, -1))
	fmt.Printf("FindAllStringSubmatchIndex ...: %#v\n", regexp.MustCompile(regex).FindAllStringSubmatchIndex(str, -1))
	fmt.Printf("FindString ...................: %#v\n", regexp.MustCompile(regex).FindString(str))
	fmt.Printf("FindStringIndex ..............: %#v\n", regexp.MustCompile(regex).FindStringIndex(str))
	fmt.Printf("FindStringSubmatch ...........: %#v\n", regexp.MustCompile(regex).FindStringSubmatch(str))
	fmt.Printf("FindStringSubmatchIndex ......: %#v\n", regexp.MustCompile(regex).FindStringSubmatchIndex(str))
}
