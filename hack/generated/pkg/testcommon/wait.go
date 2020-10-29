/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"
)

// TODO: it's unfortunate we can't just use g.Eventually all the time
func WaitFor(ctx context.Context, timeout time.Duration, check func(context.Context) (bool, error)) error {
	waitCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		default:
			done, err := check(waitCtx)
			if err != nil {
				return err
			}

			if done {
				return nil
			}

			time.Sleep(5 * time.Second)
		}
	}
}

func SetupTeardownTestMain(
	m *testing.M,
	skipSlowTests bool,
	setup func() error,
	teardown func() error) int {

	// safety check before calling testing.Short()
	if !flag.CommandLine.Parsed() {
		panic("flag.Parse must have been invoked")
	}

	if skipSlowTests && testing.Short() {
		log.Println("Skipping slow tests in short mode")
		return 0
	}

	setupErr := setup()
	if setupErr != nil {
		panic(fmt.Sprintf("setup error: %s", setupErr.Error()))
	}

	defer func() {
		if teardownErr := teardown(); teardownErr != nil {
			panic(fmt.Sprintf("teardown error: %s", teardownErr.Error()))
		}
	}()

	return m.Run()
}
