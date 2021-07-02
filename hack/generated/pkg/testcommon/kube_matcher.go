/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
	"github.com/onsi/gomega/types"
)

// TODO: Would we rather these just be on testcontext? Might read better
type KubeMatcher struct {
	ctx    context.Context
	ensure *Ensure
}

func NewKubeMatcher(ensure *Ensure, ctx context.Context) *KubeMatcher {
	return &KubeMatcher{
		ensure: ensure,
		ctx:    ctx,
	}
}

func (m *KubeMatcher) BeProvisioned() types.GomegaMatcher {
	return &DesiredStateMatcher{
		ensure:    m.ensure,
		ctx:       m.ctx,
		goalState: armclient.SucceededProvisioningState,
	}
}

func (m *KubeMatcher) BeProvisionedAfter(previousState armclient.ProvisioningState) types.GomegaMatcher {
	return &DesiredStateMatcher{
		ensure:        m.ensure,
		ctx:           m.ctx,
		goalState:     armclient.SucceededProvisioningState,
		previousState: &previousState,
	}
}

func (m *KubeMatcher) BeFailed() types.GomegaMatcher {
	return &DesiredStateMatcher{
		ensure:    m.ensure,
		ctx:       m.ctx,
		goalState: armclient.FailedProvisioningState,
	}
}

func (m *KubeMatcher) BeDeleted() types.GomegaMatcher {
	return &BeDeletedMatcher{
		ensure: m.ensure,
		ctx:    m.ctx,
	}
}
