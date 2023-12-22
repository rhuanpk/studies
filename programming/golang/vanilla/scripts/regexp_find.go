package main

import (
	"fmt"
	"regexp"
)

// go playground: https://go.dev/play/p/9c5KCLHwVCw
func main() {
	fmt.Println("FindAllString ................:", regexp.MustCompile(`(/[^/]+)`).FindAllString("/foo/bar", -1))
	fmt.Println("FindAllStringIndex ...........:", regexp.MustCompile(`(/[^/]+)`).FindAllStringIndex("/foo/bar", -1))
	fmt.Println("FindAllStringSubmatch ........:", regexp.MustCompile(`(/[^/]+)`).FindAllStringSubmatch("/foo/bar", -1))
	fmt.Println("FindAllStringSubmatchIndex ...:", regexp.MustCompile(`(/[^/]+)`).FindAllStringSubmatchIndex("/foo/bar", -1))
	fmt.Println("FindString ...................:", regexp.MustCompile(`(/[^/]+)`).FindString("/foo/bar"))
	fmt.Println("FindStringIndex ..............:", regexp.MustCompile(`(/[^/]+)`).FindStringIndex("/foo/bar"))
	fmt.Println("FindStringSubmatch ...........:", regexp.MustCompile(`(/[^/]+)`).FindStringSubmatch("/foo/bar"))
	fmt.Println("FindStringSubmatchIndex ......:", regexp.MustCompile(`(/[^/]+)`).FindStringSubmatchIndex("/foo/bar"))
}
