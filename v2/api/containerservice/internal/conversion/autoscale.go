// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package conversion

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// PullAutoscaleProfiles reads both the singleton preview representation and the later list representation.
// The caller must pass a cloned property bag because pulling removes the stored value.
func PullAutoscaleProfiles[T any](bag genruntime.PropertyBag) ([]T, error) {
	value, ok := bag["Autoscale"]
	if !ok {
		return nil, nil
	}

	if strings.HasPrefix(strings.TrimSpace(value), "{") {
		var profile T
		if err := bag.Pull("Autoscale", &profile); err != nil {
			return nil, err
		}
		return []T{profile}, nil
	}

	var profiles []T
	if err := bag.Pull("Autoscale", &profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}
