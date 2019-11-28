/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcegroups

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"
)

// var AzureResourceGroupManager ResourceGroupManager = &azureResourceGroupManager{}

func NewAzureResourceGroupManager() *AzureResourceGroupManager {
	return &AzureResourceGroupManager{}
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
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
}
