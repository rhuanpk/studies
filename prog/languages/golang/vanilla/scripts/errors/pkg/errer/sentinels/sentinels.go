package sentinels

import "errors"

var (
	// ErrInternal indicates any internal application error.
	ErrInternal = add("internal error")
	// ErrNoData indicates any missing data.
	ErrNoData = add("no data caught")
	// ErrValidation indicates any validation failed.
	ErrValidation = add("validation failed")
	// ErrCheck indiactes any checking failed.
	ErrCheck = add("checking failed")
)

var sentinels []error

func add(msg string) error {
	err := errors.New(msg)
	sentinels = append(sentinels, err)

	return err
}

// Has check if some of errs are on registered sentinels.
func Has(errs ...error) bool {
	for _, err := range errs {
		for _, e := range sentinels {
			if errors.Is(err, e) {
				return true
			}
		}
	}

	return false
}

// HasGettingFirst check if some of errs are on registered sentinels and return the first found.
func HasGettingFirst(errs ...error) (bool, error) {
	for _, err := range errs {
		for _, e := range sentinels {
			if errors.Is(err, e) {
				return true, e
			}
		}
	}

	return false, nil
}

// HasGettingAll check if some of errs are on registered sentinels and return all found.
func HasGettingAll(errs ...error) (bool, []error) {
	var (
		has bool
		all []error
	)

	for _, err := range errs {
		for _, e := range sentinels {
			if errors.Is(err, e) {
				has = true
				all = append(all, e)
			}
		}
	}

	return has, all
}
