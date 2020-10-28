/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"flag"
	"log"
	"testing"
	"time"
)

// TODO: it's unfortunate we can't just use g.Eventually all the time
func WaitFor(ctx context.Context, timeout time.Duration, check func(context.Context) (bool, error)) error {
	waitCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := false

	for !done {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		default:
			var err error
			done, err = check(waitCtx)
			if err != nil {
				return err
			}
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

// testMainWrapper is a wrapper that can be called by TestMain so that we can use defer
func SetupTeardownTestMain(
	m *testing.M,
	skipShortTests bool,
	setup func() error,
	teardown func() error) int {

	if skipShortTests {
		flag.Parse()
		if testing.Short() {
			log.Println("Skipping slow tests in short mode")
			return 0
		}
	}

	setupErr := setup()
	if setupErr != nil {
		panic(setupErr)
	}

	defer func() {
		if teardownErr := teardown(); teardownErr != nil {
			panic(teardownErr)
		}
	}()

	return m.Run()
}
