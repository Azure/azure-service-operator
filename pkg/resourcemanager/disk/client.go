// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package disk

import (
	"context"
	"strings"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
)

type AzureDiskClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureDiskClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureDiskClient {
	return &AzureDiskClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getDiskClient() compute.DisksClient {
	computeClient := compute.NewDisksClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (m *AzureDiskClient) CreateDisk(ctx context.Context, location string, resourceGroupName string, resourceName string, createOption string, diskSizeGB int32) (future compute.DisksCreateOrUpdateFuture, err error) {

	client := getDiskClient()

	diskCreateOption := compute.DiskCreateOption(strings.TrimSpace(createOption))

	// Initialize Disk API Spec properties.
	diskAPISpec := compute.Disk{
		Location: &location,
		DiskProperties: &compute.DiskProperties{
			CreationData: &compute.CreationData{
				CreateOption: diskCreateOption,
			},
			DiskSizeGB: &diskSizeGB,
		},
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		diskAPISpec,
	)

	return future, err
}

func (m *AzureDiskClient) DeleteDisk(ctx context.Context, diskName string, resourcegroup string) (status string, err error) {

	client := getDiskClient()

	_, err = client.Get(ctx, resourcegroup, diskName)
	if err == nil { // vm present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, diskName)
		return future.Status(), err
	}
	// VM  not present so return success anyway
	return "Disk not present", nil

}

func (m *AzureDiskClient) GetDisk(ctx context.Context, resourcegroup string, diskName string) (vm compute.Disk, err error) {

	client := getDiskClient()

	return client.Get(ctx, resourcegroup, diskName)
}
