// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package manageddisks

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest"
)

// AzureManagedDiskManager is the struct that the manager functions hang off
type AzureManagedDiskManager struct {
	Telemetry telemetry.Telemetry
}

// getManagedDiskClient returns a new instance of an ManagedDisks client
func getManagedDiskClient() (compute.DisksClient, error) {
	client := compute.NewDisksClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = compute.DisksClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// CreateManageDisk creates ManageDisks
func (_ *AzureManagedDiskManager) CreateManageDisk(ctx context.Context, location string, resourceGroupName string, resourceName string, createOption string, diskSizeGB int) (compute.Disk, error) {
	client, err := getManagedDiskClient()
	if err != nil {
		return compute.Disk{}, err
	}

	diskSizeGBInt32 := (int32)(diskSizeGB)
	createOptionEnum := (compute.DiskCreateOption)(compute.DiskCreateOptionTypesEmpty)
	if createOption == string(compute.DiskCreateOptionTypesEmpty) {
		createOptionEnum = (compute.DiskCreateOption)(compute.DiskCreateOptionTypesEmpty)
	} else if createOption == string(compute.DiskCreateOptionTypesAttach) {
		createOptionEnum = (compute.DiskCreateOption)(compute.DiskCreateOptionTypesAttach)
	} else if createOption == string(compute.DiskCreateOptionTypesFromImage) {
		createOptionEnum = (compute.DiskCreateOption)(compute.DiskCreateOptionTypesFromImage)
	}

	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		compute.Disk{
			Location: &location,
			DiskProperties: &compute.DiskProperties{
				CreationData: &compute.CreationData{
					CreateOption: createOptionEnum,
				},
				DiskSizeGB: &diskSizeGBInt32,
			},
		},
	)
	if err != nil {
		return compute.Disk{}, err
	}

	return future.Result(client)
}

// DeleteManagedDisk deletes a ManagedDisk
func (_ *AzureManagedDiskManager) DeleteManagedDisk(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error) {
	client, err := getManagedDiskClient()
	if err != nil {
		return autorest.Response{}, err
	}

	future, err := client.Delete(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return autorest.Response{}, err
	}

	return future.Result(client)
}

// ManagedDiskExists checks to see if a ManagedDisk exists
func (_ *AzureManagedDiskManager) ManagedDiskExists(ctx context.Context, resourceGroupName string, resourceName string) (bool, error) {
	client, err := getManagedDiskClient()
	if err != nil {
		return false, err
	}

	result, err := client.Get(
		ctx,
		resourceGroupName,
		resourceName)
	if err != nil {
		return false, err
	} else if result.Name == nil {
		return false, nil
	}

	return true, nil
}
