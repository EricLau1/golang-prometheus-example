package exit

import (
	"os"
	"os/signal"
)

const (
	SUCCESS = 1
	FAILURE = 2
)

func Graceful(onStop ...func()) {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	go func() {
		for {
			select {
			case _ = <-stop:
				for _, fn := range onStop {
					fn()
				}
				os.Exit(SUCCESS)
			}
		}
	}()
}
