// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
)

type azureConsumerGroupManager struct {
	Log logr.Logger
}

func NewConsumerGroupClient(log logr.Logger) *azureConsumerGroupManager {
	return &azureConsumerGroupManager{
		Log: log,
	}
}

func getConsumerGroupsClient() eventhub.ConsumerGroupsClient {
	consumerGroupClient := eventhub.NewConsumerGroupsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	consumerGroupClient.Authorizer = auth
	consumerGroupClient.AddToUserAgent(config.UserAgent())
	return consumerGroupClient
}

// CreateConsumerGroup creates an Event Hub Consumer Group
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
// parameters - parameters supplied to create or update a consumer group resource.
func (_ *azureConsumerGroupManager) CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	consumerGroupClient := getConsumerGroupsClient()

	parameters := eventhub.ConsumerGroup{}
	return consumerGroupClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		consumerGroupName,
		parameters,
	)

}

// DeleteConsumerGroup deletes an Event Hub Consumer Group
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
func (_ *azureConsumerGroupManager) DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error) {
	consumerGroupClient := getConsumerGroupsClient()
	return consumerGroupClient.Delete(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		consumerGroupName,
	)

}

//GetConsumerGroup gets consumer group description for the specified Consumer Group.
func (_ *azureConsumerGroupManager) GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	consumerGroupClient := getConsumerGroupsClient()
	return consumerGroupClient.Get(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName)
}

func (cg *azureConsumerGroupManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

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

	newCg, err := cg.CreateConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, azureConsumerGroupName)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// reconciliation is not done but error is acceptable
			return false, nil
		}

		return false, err
	}

	instance.Status.State = "Active"
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.ResourceId = *newCg.ID

	return true, nil
}

func (cg *azureConsumerGroupManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

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

	// deletions to non existing groups lead to 'conflicting operation' errors so we GET it here
	_, err = cg.GetConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, azureConsumerGroupName)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			// these things mean the entity is already gone
			return false, nil
		}
		return true, err
	}

	_, err = cg.DeleteConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, azureConsumerGroupName)
	if err != nil {
		return false, err
	}

	instance.Status.Message = "deleted"

	return true, nil
}

func (cg *azureConsumerGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := cg.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Eventhub,
			},
			Target: &v1alpha1.Eventhub{},
		},
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

func (g *azureConsumerGroupManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (cg *azureConsumerGroupManager) convert(obj runtime.Object) (*v1alpha1.ConsumerGroup, error) {
	local, ok := obj.(*v1alpha1.ConsumerGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
