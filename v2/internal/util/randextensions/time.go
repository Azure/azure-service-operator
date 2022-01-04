// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package randextensions

import (
	"math/rand"
	"time"
)

// Jitter returns the provided duration, scaled by a jitter factor.
// jitter must be between 0 and 1.
// After jitter is applied the result will be in the range: [(1-jitter)t, (1+jitter)t]
// UNLESS (1+jitter)*t overflows int64. Overflows are not checked by this function so
// use with caution.
func Jitter(r *rand.Rand, t time.Duration, jitter float64) time.Duration {
	if jitter <= 0 || jitter > 1 {
		panic("jitter must be in the range (0, 1]")
	}
	low := 1 - jitter
	val := low + r.Float64()*(2*jitter)
	return time.Duration(float64(t) * val)
}
