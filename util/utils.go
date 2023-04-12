package util

import "fmt"

func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}
	return fmt.Errorf("%v,%w", existErr, newErr) //收集错误
}
