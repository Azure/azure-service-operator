// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package customizations

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20220701"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.PostReconciliationChecker = &PrivateEndpointExtension{}

func (extension *PrivateEndpointExtension) PostReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger) (extensions.PostReconcileCheckResult, error) {

	if endpoint, ok := obj.(*network.PrivateEndpoint); ok && endpoint.Status.PrivateLinkServiceConnections != nil {

		for _, connection := range endpoint.Status.PrivateLinkServiceConnections {
			if *connection.PrivateLinkServiceConnectionState.Status != "Approved" {
				// Returns 'conditions.NewReadyConditionImpactingError' error
				return extensions.PostReconcileCheckResultFailure(
					fmt.Sprintf(
						"Private connection '%s' to the endpoint requires approval %q",
						*connection.Id,
						*connection.PrivateLinkServiceConnectionState)), nil
			}
		}
	}

	return extensions.PostReconcileCheckResultSuccess(), nil
}
