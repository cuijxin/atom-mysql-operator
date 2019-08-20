package signals

import (
	"context"
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler sets up a signal handler that calls the given CancelFunc
// on SIGTERM/SIGINT. If a second signal is caught, the program is terminated
// immediately with exit code 1.
func SetupSignalHandler(cancelFunc context.CancelFunc) {
	close(onlyOneSignalHandler) // panics when called twice

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancelFunc()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()
}
