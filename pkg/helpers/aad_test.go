/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractIdentityName_ExtractsCorrectName(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	resourceID := "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/mymi"
	identityName := extractIdentityName(resourceID)
	assert.Equal("mymi", identityName)
}
