package utils

import (
	"os"
	"os/signal"
	"time"
)

func BlockUntilInterrupted() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}

func BlockForSeconds(seconds float64) {
	time.Sleep(time.Duration(seconds * 1e9))
}
