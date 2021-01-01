// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package containerregistry

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"
)

// NewAzureContainerRegistryManager creates a new container registry manager
func NewAzureContainerRegistryManager(creds config.Credentials, scheme *runtime.Scheme, secretClient secrets.SecretClient) *azureContainerRegistryManager {
	return &azureContainerRegistryManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

type ContainerRegistryManager interface {
	// CreateRegistry creates the contain registry in Azure
	CreateRegistry(ctx context.Context, instance *azurev1alpha1.AzureContainerRegistry) (containerregistry.Registry, error)

	// UpdateRegistry updates the contain registry in Azure
	UpdateRegistry(ctx context.Context, instance *azurev1alpha1.AzureContainerRegistry) (containerregistry.Registry, error)

	// DeleteVault removes the container registry in Azure
	DeleteRegistry(ctx context.Context, groupName string, registryName string) (result autorest.Response, err error)

	// CheckExistence checks for the presence of a container registry instance on Azure
	GetRegistry(ctx context.Context, groupName string, registryName string) (result containerregistry.Registry, err error)

	// ListCredentials returns the list of login credentials for the speciried container registry
	ListCredentials(ctx context.Context, groupName string, registryName string) (result containerregistry.RegistryListCredentialsResult, err error)

	// also embed async client methods
	resourcemanager.ARMClient
}
