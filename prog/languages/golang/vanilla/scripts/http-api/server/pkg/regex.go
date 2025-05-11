package pkg

import "regexp"

// RegEx simulates the grep GNU tool.
func RegEx(pattern, str string) bool {
	return regexp.MustCompile(pattern).MatchString(str)
}
