// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcegroups

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func NewAzureResourceGroupManager(creds config.Credentials) *AzureResourceGroupManager {
	return &AzureResourceGroupManager{creds: creds}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewAzureResourceGroupManager(creds)
}

type ResourceGroupManager interface {
	// CreateGroup creates a new resource group named by env var
	CreateGroup(ctx context.Context, groupName string, location string) (resources.Group, error)

	// DeleteGroup removes the resource group named by env var
	DeleteGroup(ctx context.Context, groupName string) (result autorest.Response, err error)

	// DeleteGroup removes the resource group named by env var
	DeleteGroupAsync(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error)

	// CheckExistence checks whether a resource exists
	CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)

	// also embed methods from AsyncClient
	resourcemanager.ARMClient
}
