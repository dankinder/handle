package handle

import (
	"errors"
	"testing"
)

func TestErrPanicsOnError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected to recover a panic, but no panic happened")
		}
	}()
	Err(errors.New("Whups"))
}

func TestErrNoPanicOnNilError(t *testing.T) {
	Err(nil)
}
