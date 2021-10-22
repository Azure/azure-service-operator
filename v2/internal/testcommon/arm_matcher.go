/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/onsi/gomega/types"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

type ARMMatcher struct {
	client *genericarmclient.GenericClient
}

func NewARMMatcher(client *genericarmclient.GenericClient) *ARMMatcher {
	return &ARMMatcher{
		client: client,
	}
}

func (m *ARMMatcher) BeProvisioned(ctx context.Context) types.GomegaMatcher {
	return &AzureBeProvisionedMatcher{
		azureClient: m.client,
		ctx:         ctx,
	}
}

func (m *ARMMatcher) BeDeleted(ctx context.Context) types.GomegaMatcher {
	return &AzureBeDeletedMatcher{
		azureClient: m.client,
		ctx:         ctx,
	}
}
