package resourcegroups

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest"
)

var AzureResourceGroupManager ResourceGroupManager = &azureResourceGroupManager{}

type ResourceGroupManager interface {
	// CreateGroup creates a new resource group named by env var
	CreateGroup(ctx context.Context, groupName string, location string) (resources.Group, error)

	// DeleteGroup removes the resource group named by env var
	DeleteGroup(ctx context.Context, groupName string) (result autorest.Response, err error)

	// CheckExistence checks whether a resource exists
	CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)
}
