// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
)

func TestIncludesWebhooks(t *testing.T) {
	assert.Equal(t, true, config.OperatorModeWebhooks.IncludesWebhooks())
	assert.Equal(t, true, config.OperatorModeBoth.IncludesWebhooks())
	assert.Equal(t, false, config.OperatorModeWatchers.IncludesWebhooks())
}

func TestIncludesWatchers(t *testing.T) {
	assert.Equal(t, false, config.OperatorModeWebhooks.IncludesWatchers())
	assert.Equal(t, true, config.OperatorModeBoth.IncludesWatchers())
	assert.Equal(t, true, config.OperatorModeWatchers.IncludesWatchers())
}

func TestRoundtripString(t *testing.T) {
	values := []string{"watchers-and-webhooks", "webhooks", "watchers"}
	for _, value := range values {
		mode, err := config.ParseOperatorMode(value)
		require.Nil(t, err, value)
		assert.Equal(t, value, mode.String())
	}
}

func TestRoundtripValue(t *testing.T) {
	modes := []config.OperatorMode{
		config.OperatorModeWatchers,
		config.OperatorModeWebhooks,
		config.OperatorModeBoth,
	}
	for _, mode := range modes {
		parsed, err := config.ParseOperatorMode(mode.String())
		require.Nil(t, err, mode)
		assert.Equal(t, mode, parsed)
	}
}

func TestParseInvalid(t *testing.T) {
	_, err := config.ParseOperatorMode("dimmer")
	require.Contains(t, err.Error(), `operator mode value must be one of "watchers-and-webhooks", "webhooks" or "watchers" but was "dimmer"`)
}

func TestStringInvalid(t *testing.T) {
	require.Panics(t, func() {
		_ = config.OperatorMode(0).String()
	})
}
