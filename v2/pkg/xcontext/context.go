/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package xcontext

import (
	"context"
	"os"
	"os/signal"
)

// MakeInterruptibleContext returns a context that will be cancelled when an interrupt signal is received.
// This allows a command line tool to be cancelled by Ctrl+C.
func MakeInterruptibleContext(ctx context.Context) context.Context {
	result, cancel := context.WithCancel(ctx)

	// Wait for a signal to quit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	go func() {
		<-signalChan
		cancel()
	}()

	return result
}
