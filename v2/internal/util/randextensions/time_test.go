// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package randextensions_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"

	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
)

func Test_Jitter(t *testing.T) {
	t.Parallel()

	seed := time.Now().UnixNano()
	//nolint:gosec // do not want cryptographic randomness here
	r := rand.New(lockedrand.NewSource(seed))
	properties := gopter.NewProperties(gopter.DefaultTestParametersWithSeed(seed))

	properties.Property(
		"Is in expected jitter range",
		prop.ForAll(
			func(t int64, jitter float64) bool {
				result := randextensions.Jitter(r, time.Duration(t), jitter)

				expectedLow := time.Duration((1 - jitter) * float64(t))

				halfRange := int64(jitter * float64(t))
				expectedHigh := time.Duration((1 + jitter) * float64(t)) // This could overflow -- check is below
				if math.MaxInt64-halfRange < t {
					expectedHigh = math.MaxInt64
				}
				return result >= expectedLow && result <= expectedHigh
			},
			gen.Int64Range(0, math.MaxInt64),
			gen.Float64Range(0.001, 1)))

	properties.TestingRun(t)
}
