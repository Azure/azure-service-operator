/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package matchers

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// TODO: Put this into a subpackage
// TODO: Would we rather these just be on testcontext? Might read better
type Azure struct {
	ctx    context.Context
	client *genericarmclient.GenericClient
}

func NewAzure(ctx context.Context, client *genericarmclient.GenericClient) *Azure {
	return &Azure{
		ctx:    ctx,
		client: client,
	}
}

func (m *Azure) BeDeleted(apiVersion string) *BeDeletedInAzureMatcher {
	return &BeDeletedInAzureMatcher{
		ctx:        m.ctx,
		client:     m.client,
		apiVersion: apiVersion,
	}
}
