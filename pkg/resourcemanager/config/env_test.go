// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
)

const parseTestVar = "---TestParseStringListFromEnvironment---"

func checkValue(value string) []string {
	// Try to avoid polluting global state, although there's no way to
	// unset a variable.
	oldValue := envy.Get(parseTestVar, "")
	envy.Set(parseTestVar, value)
	result := config.ParseStringListFromEnvironment(parseTestVar)
	envy.Set(parseTestVar, oldValue)
	return result
}

func TestParseStringListFromEnvironment(t *testing.T) {
	require := require.New(t)
	require.Empty(checkValue(""))
	require.Empty(checkValue("    "))
	require.Equal(checkValue("a"), []string{"a"})
	require.Equal(checkValue("a,b,c,d"), []string{"a", "b", "c", "d"})
	require.Equal(checkValue("a , b, c "), []string{"a", "b", "c"})
}
