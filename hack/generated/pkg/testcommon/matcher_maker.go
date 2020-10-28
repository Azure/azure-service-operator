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
type MatcherMaker struct {
	ensure *Ensure
}

func NewMatcherMaker(ensure *Ensure) *MatcherMaker {
	return &MatcherMaker{
		ensure: ensure,
	}
}

func (m *MatcherMaker) BeProvisioned(ctx context.Context) types.GomegaMatcher {
	return &BeProvisionedMatcher{
		ensure: m.ensure,
		ctx:    ctx,
	}
}

func (m *MatcherMaker) BeDeleted(ctx context.Context) types.GomegaMatcher {
	return &BeDeletedMatcher{
		ensure: m.ensure,
		ctx:    ctx,
	}
}
