package logging

import (
	"fmt"
	"slices"
	"strings"
)

type layer string

const (
	controller layer = "controller"
	service    layer = "service"
	repository layer = "repository"
)

const (
	vertical = "├─"
	corner   = "└─"
)

func separator(ok bool) string {
	return map[bool]string{true: corner, false: vertical}[ok]
}

func hasLayer(err error) bool {
	return strings.Contains(err.Error(), "layer: ")
}

func nextIndexIsLast[Type any](index int, slice []Type) (int, bool) {
	lastIndex := len(slice) - 1
	nextIndex := index + 1
	nenextIndex := index + 2

	nextIndexIsLast := nenextIndex > lastIndex

	return nextIndex, nextIndexIsLast
}

func isAtualLast[Type any](index int, errs []Type) bool {
	return index >= len(errs)-1
}

func isNextLastAndNil(index int, errs []error) bool {
	nextIndex, nextIndexIsLast := nextIndexIsLast(index, errs)
	return nextIndexIsLast && errs[nextIndex] == nil
}

func isNextLastAndEmpty(index int, messages []string) bool {
	nextIndex, nextIndexIsLast := nextIndexIsLast(index, messages)
	return nextIndexIsLast && messages[nextIndex] == ""
}

func buildError(tabs string, layer layer, funcName string, messages []string, sentinels []error, errs ...error) error {
	var errBuild, errLayer error

	for index, err := range errs {
		if hasLayer(err) {
			errLayer = err
			// errs = append(errs[:index], errs[index+1:]...)
			errs = slices.Delete(errs, index, index+1)
		}
	}

	noFuncName := funcName == ""
	noMessages := len(messages) <= 0
	noSentinels := len(sentinels) <= 0
	noErrors := len(errs) <= 0

	// build layer
	separatorLayer := separator(
		noFuncName && noMessages && noSentinels && noErrors,
	)
	errBuild = fmt.Errorf(
		"%s%s layer: %s",
		tabs, separatorLayer, strings.ToUpper(string(layer)),
	)

	// build func name
	if funcName != "" {
		separatorFuncName := separator(
			noMessages && noSentinels && noErrors,
		)

		errBuild = fmt.Errorf(
			"%w%s%s function: %s",
			errBuild, tabs, separatorFuncName, funcName,
		)
	}

	// build messages
	for index, message := range messages {
		if message == "" {
			continue
		}

		separatorMessage := separator(
			(noSentinels && noErrors) && (isAtualLast(index, errs) || isNextLastAndEmpty(index, messages)),
		)

		errBuild = fmt.Errorf(
			"%w%s%s message: %s",
			errBuild, tabs, separatorMessage, message,
		)
	}

	// build sentinels
	for index, sentinel := range sentinels {
		if sentinel == nil {
			continue
		}

		separatorSentinel := separator(
			noErrors && (isAtualLast(index, sentinels) || isNextLastAndNil(index, sentinels)),
		)

		errBuild = fmt.Errorf(
			"%w%s%s sentinel: %w",
			errBuild, tabs, separatorSentinel, sentinel,
		)
	}

	// build errors
	for index, err := range errs {
		if err == nil {
			continue
		}

		separatorError := separator(
			isAtualLast(index, errs) || isNextLastAndNil(index, errs),
		)

		errBuild = fmt.Errorf(
			"%w%s%s error: %w",
			errBuild, tabs, separatorError, err,
		)
	}

	// build error layer
	if errLayer != nil {
		errBuild = fmt.Errorf("%w%w", errBuild, errLayer)
	}

	return errBuild
}

// Controller build the controller layer error log.
func Controller(funcName string, messages []string, sentinels []error, errs ...error) error {
	const tabs = "\n\t"
	return buildError(tabs, controller, funcName, messages, sentinels, errs...)
}

// Service build the service layer error log.
func Service(funcName string, messages []string, sentinels []error, errs ...error) error {
	const tabs = "\n\t\t"
	return buildError(tabs, service, funcName, messages, sentinels, errs...)
}

// Repository build the repository layer error log.
func Repository(funcName string, messages []string, sentinels []error, errs ...error) error {
	const tabs = "\n\t\t\t"
	return buildError(tabs, repository, funcName, messages, sentinels, errs...)
}
