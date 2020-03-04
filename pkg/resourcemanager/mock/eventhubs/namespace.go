// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	apphelpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type eventHubNamespaceResource struct {
	resourceGroupName string
	namespaceName     string
	eHNamespace       eventhub.EHNamespace
}

type mockEventHubNamespaceManager struct {
	eventHubNamespaceResources []eventHubNamespaceResource
}

func NewMockEventHubNamespaceClient() *mockEventHubNamespaceManager {
	return &mockEventHubNamespaceManager{}
}

func findEventHubNamespace(res []eventHubNamespaceResource, predicate func(eventHubNamespaceResource) bool) (int, eventHubNamespaceResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, eventHubNamespaceResource{}
}

func (manager *mockEventHubNamespaceManager) CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	var eventHubNamespace = eventhub.EHNamespace{
		Response: helpers.GetRestResponse(201),
		EHNamespaceProperties: &eventhub.EHNamespaceProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
		Location: &location,
		Name:     to.StringPtr(namespaceName),
	}
	manager.eventHubNamespaceResources = append(manager.eventHubNamespaceResources, eventHubNamespaceResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eHNamespace:       eventHubNamespace,
	})
	return &eventHubNamespace, nil
}

func (manager *mockEventHubNamespaceManager) CreateNamespace(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	var eventHubNamespace = eventhub.EHNamespace{
		Response: helpers.GetRestResponse(201),
		EHNamespaceProperties: &eventhub.EHNamespaceProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
		Location: &location,
		Name:     to.StringPtr(namespaceName),
	}
	manager.eventHubNamespaceResources = append(manager.eventHubNamespaceResources, eventHubNamespaceResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eHNamespace:       eventHubNamespace,
	})
	return &eventHubNamespace, nil
}

func (manager *mockEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (autorest.Response, error) {
	namespaces := manager.eventHubNamespaceResources

	index, _ := findEventHubNamespace(namespaces, func(g eventHubNamespaceResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("eventhub namespace not found")
	}

	manager.eventHubNamespaceResources = append(namespaces[:index], namespaces[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

func (manager *mockEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	groups := manager.eventHubNamespaceResources

	index, group := findEventHubNamespace(groups, func(g eventHubNamespaceResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName
	})

	if index == -1 {
		return &eventhub.EHNamespace{}, errors.New("eventhub namespace not found")
	}

	return &group.eHNamespace, nil
}

func (ns *mockEventHubNamespaceManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

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
	_, err = ns.GetNamespace(ctx, resourcegroup, namespaceName)
	if err == nil {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true

		return true, nil
	}

	// set this message so the tests are happy
	if instance.Spec.ResourceGroup == "gone" {
		instance.Status.Message = "ResourceGroupNotFound"
	}

	// create Event Hubs namespace
	_, err = ns.CreateNamespace(ctx, resourcegroup, namespaceName, namespaceLocation)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if apphelpers.ContainsString(catch, azerr.Type) {
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

func (ns *mockEventHubNamespaceManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return false, err
	}

	namespaceName := instance.Name
	resourcegroup := instance.Spec.ResourceGroup

	_, err = ns.DeleteNamespace(ctx, resourcegroup, namespaceName)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)

		// check if async op from previous delete was still happening
		catch := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.RequestConflictError,
		}
		if apphelpers.ContainsString(catch, azerr.Type) {
			return true, nil
		}

		// check if namespace was already gone
		catch = []string{
			errhelp.NotFoundErrorCode,
		}
		if apphelpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		// some error we don't know about or can't handle happened
		instance.Status.Provisioning = false

		return true, fmt.Errorf("EventhubNamespace delete error %v", err)

	}

	return false, nil
}

func (ns *mockEventHubNamespaceManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return nil, err
	}

	key := types.NamespacedName{Namespace: instance.Namespace, Name: instance.Spec.ResourceGroup}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &v1alpha1.ResourceGroup{}},
	}, nil
}

func (g *mockEventHubNamespaceManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (ns *mockEventHubNamespaceManager) convert(obj runtime.Object) (*v1alpha1.EventhubNamespace, error) {
	local, ok := obj.(*v1alpha1.EventhubNamespace)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
