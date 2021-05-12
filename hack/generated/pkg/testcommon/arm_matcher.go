/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/onsi/gomega/types"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

type ArmMatcher struct {
	client armclient.Applier
}

func NewArmMatcher(armClient armclient.Applier) *ArmMatcher {
	return &ArmMatcher{
		client: armClient,
	}
}

func (m *ArmMatcher) BeProvisioned(ctx context.Context) types.GomegaMatcher {
	return &AzureBeProvisionedMatcher{
		azureClient: m.client,
		ctx:         ctx,
	}
}

func (m *ArmMatcher) BeDeleted(ctx context.Context) types.GomegaMatcher {
	return &AzureBeDeletedMatcher{
		azureClient: m.client,
		ctx:         ctx,
	}
}
