// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimgmt

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// APIManager is an interface for API Management
type APIManager interface {
	CreateAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string, properties azurev1alpha1.APIProperties, eTag string) (apimanagement.APIContract, error)
	DeleteAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string, eTag string, deleteRevisions bool) (autorest.Response, error)
	GetAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string) (apimanagement.APIContract, error)
	resourcemanager.ARMClient
}
