package impl

import (
	"context"
	"os"
	"os/signal"
)

func WithSignals(parent context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	done := make(chan struct{})
	chS := make(chan os.Signal, 1)
	signal.Notify(chS, signals...)
	go func() {
		select {
		case <-chS:
		case <-done:
		case <-parent.Done():
		}
		signal.Stop(chS)
		close(chS)
		cancel()
	}()
	return ctx, func() {
		select {
		case <-done:
		default:
			close(done)
		}
	}
}
