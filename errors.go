package utils

import (
	"fmt"
)

// WrapError wraps error with their callers name
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	context := CallerFunctionName()
	err = fmt.Errorf("%s: %w", context, err)
	return err
}

// WrapErrorP wraps a pointer to an error, only to be used in functions with named return values
func WrapErrorP(err *error) {
	if *err == nil {
		return
	}
	context := CallerFunctionName()
	*err = fmt.Errorf("%s: %w", context, *err)
}
