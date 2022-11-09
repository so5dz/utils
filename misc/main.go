package misc

import (
	"log"
	"os"
)

const _FailureExitCode = 1
const _SuccessExitCode = 0

func WrapMain(mainWithError func() error) func() {
	return func() {
		err := mainWithError()
		if err != nil {
			log.Println(err)
			os.Exit(_FailureExitCode)
		} else {
			os.Exit(_SuccessExitCode)
		}
	}
}
