package go_trycatch

import (
)

type tryCatchError struct {
	Err error
}

func Throw(err error) {
	panic(tryCatchError{Err: err})
}

func TryCatch(
	try func(),
) (err error) {

	defer func() {
		if r := recover(); r != nil {
			rErr, ok := r.(tryCatchError)
			if !ok {
				panic(r)
			}
			err = rErr.Err
		}
	}()

	try()

	return err
}
