package go_trycatch_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/public_bucket/go_trycatch"
)

func TestReturnsErrorsFromFuncsWhichThrowErrors(t *testing.T) {
	err := go_trycatch.TryCatch(func() {
		ThrowingFunc()
	})
	require.ErrorContains(t, err, "ThrowingFunc")
}

func TestPanicsWhenAnErrorIsPassedToPanic(t *testing.T) {
	require.Panics(t, func() {
		go_trycatch.TryCatch(func() {
			PanicWithErrorFunc()
		})
	})
}

func ThrowingFunc() {
	go_trycatch.Throw(errors.New("error from ThrowingFunc"))
}

func PanicWithErrorFunc() {
	panic(errors.New("a panicing error from PanicWithErrorFunc"))
}
