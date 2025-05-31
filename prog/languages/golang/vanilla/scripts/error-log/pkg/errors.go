package pkg

import (
	"errors"
	"fmt"
	"strings"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type layer string

var (
	ControllerLayer layer = "controller"
	ServiceLayer    layer = "service"
	RepositoryLayer layer = "repository"
)

var (
	ErrInternal   = errors.New("internal error")
	ErrNoData     = errors.New("no data caught")
	ErrValidation = errors.New("validation failed")
	ErrCheck      = errors.New("checking failed")
)

func HasAnyTheseErrors(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func hasLayer(layer layer, errors ...error) bool {
	var errStrings []string
	for _, err := range errors {
		errStrings = append(errStrings, err.Error())
	}
	return strings.Contains(strings.Join(errStrings, "\n"), string("layer: "+layer))
}

func BuildErrorLog(layer layer, funcName, message string, errors ...error) error {
	const (
		corner   = "└─"
		vertical = "├─"
	)

	var tabs string
	switch layer {
	case ControllerLayer:
		tabs = "\n\t"
	case ServiceLayer:
		tabs = "\n\t\t"
	case RepositoryLayer:
		tabs = "\n\t\t\t"
	}

	serviceWithRepository := layer == ServiceLayer && hasLayer(RepositoryLayer, errors...)
	breakable := len(errors) <= 0 || (len(errors) == 1 && (hasLayer(ServiceLayer, errors...) || hasLayer(RepositoryLayer, errors...)))

	separator := vertical
	if message == "" && breakable {
		separator = corner
	}
	text := fmt.Errorf(
		"%s%s layer: %s%s%s func: %s",
		tabs, vertical, layer, tabs, separator, funcName,
	)

	separator = vertical
	if breakable {
		separator = corner
	}
	if message != "" {
		text = fmt.Errorf("%w%s%s message: %s", text, tabs, separator, message)
	}

	separator = vertical
	for index, err := range errors {
		prefix := "sentinel"
		if index >= len(errors)-1 {
			if (layer == ControllerLayer && hasLayer(ServiceLayer, errors...)) || serviceWithRepository {
				return fmt.Errorf("%w%w", text, err)
			}
			return fmt.Errorf("%w%s%s error: %w", text, tabs, corner, err)
		}
		if index >= len(errors)-2 && serviceWithRepository {
			separator = corner
			prefix = "error"
		}
		text = fmt.Errorf("%w%s%s %s: %w", text, tabs, separator, prefix, err)
	}

	return text
}
