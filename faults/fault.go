package faults

import (
	"runtime"

	"github.com/pkg/errors"
)

func Trace(err error) error {
	if err == nil {
		return nil
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return errors.Wrap(err, f.Name())
}

func New(message string) error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	return errors.Wrap(errors.New(message), f.Name())
}
