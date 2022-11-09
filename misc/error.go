package misc

import (
	"errors"
	"fmt"
	"strings"
)

func NewError(a ...any) error {
	return errors.New(fmt.Sprintln(a...))
}

func WrapError(newError string, cause error) error {
	return fmt.Errorf("%s: \n%v", newError, cause)
}

func WrapMultiple(newError string, errors []error) error {
	if len(errors) == 0 {
		return nil
	}
	if len(errors) == 1 {
		return WrapError(newError, errors[0])
	}
	formatString := "%s: \n" + strings.Repeat("%v, ", len(errors)-1) + "%v"
	return fmt.Errorf(formatString, newError, errors)
}
