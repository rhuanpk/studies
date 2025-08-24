package errer

import "errors"

// HasAnyThese check if has at least one of passed errors.
func HasAnyThese(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
