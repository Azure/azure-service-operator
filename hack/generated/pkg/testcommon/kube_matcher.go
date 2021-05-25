/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/onsi/gomega/types"
)

// TODO: Would we rather these just be on testcontext? Might read better
type KubeMatcher struct {
	ensure *Ensure
}

func NewKubeMatcher(ensure *Ensure) *KubeMatcher {
	return &KubeMatcher{
		ensure: ensure,
	}
}

func (m *KubeMatcher) BeProvisioned(ctx context.Context) types.GomegaMatcher {
	return &BeProvisionedMatcher{
		ensure: m.ensure,
		ctx:    ctx,
	}
}

func (m *KubeMatcher) BeDeleted(ctx context.Context) types.GomegaMatcher {
	return &BeDeletedMatcher{
		ensure: m.ensure,
		ctx:    ctx,
	}
}
