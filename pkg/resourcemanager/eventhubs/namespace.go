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

package eventhubs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type azureEventHubNamespaceManager struct {
	Log logr.Logger
}

func getNamespacesClient() eventhub.NamespacesClient {
	nsClient := eventhub.NewNamespacesClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(config.UserAgent())
	return nsClient
}

func NewEventHubNamespaceClient(log logr.Logger) *azureEventHubNamespaceManager {
	return &azureEventHubNamespaceManager{
		Log: log,
	}
}

// DeleteNamespace deletes an existing namespace. This operation also removes all associated resources under the namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (_ *azureEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (autorest.Response, error) {

	nsClient := getNamespacesClient()
	future, err := nsClient.Delete(ctx,
		resourceGroupName,
		namespaceName)

	return autorest.Response{Response: future.Response()}, err
}

// Get gets the description of the specified namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (_ *azureEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	nsClient := getNamespacesClient()
	x, err := nsClient.Get(ctx, resourceGroupName, namespaceName)

	if err != nil {
		return &eventhub.EHNamespace{
			Response: x.Response,
		}, err
	}

	return &x, err
}

// CreateNamespaceAndWait creates an Event Hubs namespace
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// location - azure region
func (_ *azureEventHubNamespaceManager) CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	nsClient := getNamespacesClient()
	future, err := nsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventhub.EHNamespace{
			Location: to.StringPtr(location),
		},
	)
	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, nsClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(nsClient)
	return &result, err
}

func (_ *azureEventHubNamespaceManager) CreateNamespace(ctx context.Context, resourceGroupName string, namespaceName string, location string) (eventhub.EHNamespace, error) {
	nsClient := getNamespacesClient()

	future, err := nsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventhub.EHNamespace{
			Location: to.StringPtr(location),
		},
	)
	if err != nil {
		return eventhub.EHNamespace{}, err
	}

	return future.Result(nsClient)
}

func (ns *azureEventHubNamespaceManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return false, err
	}

	namespaceLocation := instance.Spec.Location
	namespaceName := instance.Name
	resourcegroup := instance.Spec.ResourceGroup

	// write information back to instance
	instance.Status.Provisioning = true

	// @todo handle updates
	evhns, err := ns.GetNamespace(ctx, resourcegroup, namespaceName)
	if err == nil {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.State = *evhns.ProvisioningState
		instance.Status.Message = "Namespace already existed"

		if *evhns.ProvisioningState == "Succeeded" {
			return true, nil
		}

		return false, nil
	}

	// create Event Hubs namespace
	_, err = ns.CreateNamespace(ctx, resourcegroup, namespaceName, namespaceLocation)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = err.Error()
			return false, nil
		}

		instance.Status.Provisioning = false

		return true, fmt.Errorf("EventhubNamespace create error %v", err)

	}
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

func (ns *azureEventHubNamespaceManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return false, err
	}

	namespaceName := instance.Name
	resourcegroup := instance.Spec.ResourceGroup

	resp, err := ns.DeleteNamespace(ctx, resourcegroup, namespaceName)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)

		// check if async op from previous delete was still happening
		catch := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.RequestConflictError,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		}

		// check if namespace was already gone
		catch = []string{
			errhelp.NotFoundErrorCode,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		// some error we don't know about or can't handle happened
		instance.Status.Provisioning = false

		return true, fmt.Errorf("EventhubNamespace delete error %v", err)

	}

	// if NoContent response is returned, namespace is already gone
	if resp.StatusCode == http.StatusNoContent {
		return false, nil
	}

	return true, nil
}

func (ns *azureEventHubNamespaceManager) Parents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return nil, err
	}

	key := types.NamespacedName{Namespace: instance.Namespace, Name: instance.Spec.ResourceGroup}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &v1alpha1.ResourceGroup{}},
	}, nil

}

func (ns *azureEventHubNamespaceManager) convert(obj runtime.Object) (*azurev1alpha1.EventhubNamespace, error) {
	local, ok := obj.(*azurev1alpha1.EventhubNamespace)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
