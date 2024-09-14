package main

import (
	"regexp"
	"runtime"
)

// func funcName(fullPath bool) string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	function := runtime.FuncForPC(pc)
// 	name := function.Name()
// 	if !fullPath {
// 		regex := regexp.MustCompile(`\.(.*)$`)
// 		submatchs := regex.FindStringSubmatch(name)
// 		name = submatchs[1]
// 	}
// 	return name
// }

func funcName(fullPath bool) string {
	pc, _, _, _ := runtime.Caller(1)
	return map[bool]string{
		true:  runtime.FuncForPC(pc).Name(),
		false: regexp.MustCompile(`\.(.*)$`).FindStringSubmatch(runtime.FuncForPC(pc).Name())[1],
	}[fullPath]
}

// go playground: https://go.dev/play/p/QCwDmoJT-07
func main() {
	println(funcName(false))
}
