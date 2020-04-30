/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package xcobra

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/devigned/tab"
	"github.com/spf13/cobra"
)

type (
	// ErrorWithCode is an error that contains an os.Exit code
	ErrorWithCode struct {
		Code int
	}
)

func (ewc ErrorWithCode) Error() string {
	return fmt.Sprintf("failed with error code: %d", ewc.Code)
}

// NewErrorWithCode will return a new error with an os.Exit code
func NewErrorWithCode(code int) *ErrorWithCode {
	return &ErrorWithCode{
		Code: code,
	}
}

// RunWithCtx will run a command which will respect os signals and propagate the context to children
func RunWithCtx(run func(ctx context.Context, cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	go func() {
		<-signalChan
		cancel()
	}()

	return func(cmd *cobra.Command, args []string) {
		ctx, span := tab.StartSpan(ctx, cmd.Name()+".Run")
		defer span.End()
		defer cancel()

		var err error
		cmd.PostRunE = func(c *cobra.Command, args []string) error {
			defer exitWithCode(err)
			return err
		}

		err = run(ctx, cmd, args)
	}
}
