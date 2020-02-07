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
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type consumerGroupResource struct {
	resourceGroupName string
	namespaceName     string
	eventHubName      string
	consumerGroupName string
	ConsumerGroup     eventhub.ConsumerGroup
}

type mockConsumerGroupManager struct {
	consumerGroupResources []consumerGroupResource
}

func NewMockConsumerGroupClient() *mockConsumerGroupManager {
	return &mockConsumerGroupManager{}
}

func findConsumerGroup(res []consumerGroupResource, predicate func(consumerGroupResource) bool) (int, consumerGroupResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, consumerGroupResource{}
}

func find(res []interface{}, predicate func(interface{}) bool) (int, interface{}) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, consumerGroupResource{}
}

func (manager *mockConsumerGroupManager) CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	var consumerGroup = eventhub.ConsumerGroup{
		Response: helpers.GetRestResponse(201),
		Name:     to.StringPtr(consumerGroupName),
	}
	manager.consumerGroupResources = append(manager.consumerGroupResources, consumerGroupResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eventHubName:      eventHubName,
		consumerGroupName: consumerGroupName,
		ConsumerGroup:     consumerGroup,
	})
	return consumerGroup, nil
}

func (manager *mockConsumerGroupManager) DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (autorest.Response, error) {
	groups := manager.consumerGroupResources

	index, _ := findConsumerGroup(groups, func(g consumerGroupResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName &&
			g.consumerGroupName == consumerGroupName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("consumer group not found")
	}

	manager.consumerGroupResources = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

func (manager *mockConsumerGroupManager) GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	groups := manager.consumerGroupResources

	index, group := findConsumerGroup(groups, func(g consumerGroupResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName &&
			g.consumerGroupName == consumerGroupName
	})

	if index == -1 {
		return eventhub.ConsumerGroup{
			Response: helpers.GetRestResponse(http.StatusNotFound),
		}, errors.New("consumer group not found")
	}

	group.ConsumerGroup.Response = helpers.GetRestResponse(http.StatusOK)
	return group.ConsumerGroup, nil
}

func (cg *mockConsumerGroupManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.EnsureOption) (bool, error) {

	instance, err := cg.convert(obj)
	if err != nil {
		return false, err
	}

	// write information back to instance
	instance.Status.Provisioning = true

	kubeObjectName := instance.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	eventhubName := instance.Spec.Eventhub
	azureConsumerGroupName := instance.Spec.ConsumerGroupName

	// if no need for shared consumer group name, use the kube name
	if len(azureConsumerGroupName) == 0 {
		azureConsumerGroupName = kubeObjectName
	}

	resp, err := cg.CreateConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, azureConsumerGroupName)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		return false, err
	}
	instance.Status.State = resp.Status
	instance.Status.Message = "success"
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

func (cg *mockConsumerGroupManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := cg.convert(obj)
	if err != nil {
		return false, err
	}

	kubeObjectName := instance.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	eventhubName := instance.Spec.Eventhub
	azureConsumerGroupName := instance.Spec.ConsumerGroupName

	// if no need for shared consumer group name, use the kube name
	if len(azureConsumerGroupName) == 0 {
		azureConsumerGroupName = kubeObjectName
	}

	_, err = cg.DeleteConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, azureConsumerGroupName)
	if err != nil {
		return false, err
	}

	instance.Status.Message = "deleted"

	return true, nil
}

func (cg *mockConsumerGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := cg.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Namespace,
			},
			Target: &v1alpha1.EventhubNamespace{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil

}

func (cg *mockConsumerGroupManager) convert(obj runtime.Object) (*v1alpha1.ConsumerGroup, error) {
	local, ok := obj.(*v1alpha1.ConsumerGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
