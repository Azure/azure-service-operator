// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"fmt"
	"testing"
	"time"
)

func TestRetryTimeout(t *testing.T) {
	start := time.Now()
	_ = Retry(5*time.Second, 1*time.Second, func() error {
		return fmt.Errorf("test")
	})
	stop := time.Since(start)
	if stop < 5*time.Second {
		t.Errorf("retry ended too soon: %s", stop)
	}
}
func TestRetryStopErr(t *testing.T) {
	start := time.Now()
	_ = Retry(5*time.Second, 1*time.Second, func() error {
		return NewStop(fmt.Errorf("test"))
	})
	stop := time.Since(start)
	if stop > 1*time.Second {
		t.Errorf("retry with stop should not take so long: %s", stop)
	}
}
