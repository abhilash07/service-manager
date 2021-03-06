package util

import (
	"context"
	"os"
	"os/signal"

	"github.com/Peripli/service-manager/pkg/log"
)

// HandleInterrupts handles process signal interrupts
func HandleInterrupts() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt)
	go func() {
		select {
		case <-term:
			log.C(ctx).Error("Received OS interrupt, exiting gracefully...")
			cancel()
		case <-ctx.Done():
			return
		}
	}()
	return ctx, cancel
}
