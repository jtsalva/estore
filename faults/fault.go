package faults

import (
	"errors"
	"fmt"
	"runtime"
)

func Trace(err error) error {
	if err == nil {
		return nil
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return errors.New(fmt.Sprintf("%s: %s", f.Name(), err.Error()))
}
