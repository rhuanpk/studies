package runtime

import (
	"regexp"
	"runtime"
)

// FuncName returns the function name of your caller.
func FuncName(fullPath bool) string {
	pc, _, _, _ := runtime.Caller(1)
	return map[bool]string{
		true:  runtime.FuncForPC(pc).Name(),
		false: regexp.MustCompile(`\.(.*)$`).FindStringSubmatch(runtime.FuncForPC(pc).Name())[1],
	}[fullPath]
}
