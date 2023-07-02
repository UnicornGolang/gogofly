package utils

import "fmt"

func AppendError(existsErr, newErr error) error {
	if existsErr == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", existsErr, newErr)
}
