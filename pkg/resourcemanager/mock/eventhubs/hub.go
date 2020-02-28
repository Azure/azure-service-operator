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

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	pkghelpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

type eventHubAccess struct {
	rule eventhub.AuthorizationRule
	keys eventhub.AccessKeys
}

type eventHubResource struct {
	resourceGroupName string
	namespaceName     string
	eventHubName      string
	eventHub          eventhub.Model
	eventHubAccesses  []eventHubAccess
}

func findEventHub(res []eventHubResource, predicate func(eventHubResource) bool) (int, eventHubResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, eventHubResource{}
}

func findAccess(res []eventHubAccess, name string) (int, eventHubAccess) {
	for index, r := range res {
		if *r.rule.Name == name {
			return index, r
		}
	}
	return -1, eventHubAccess{}
}

type mockEventHubManager struct {
	eventHubResources []eventHubResource
	SecretClient      secrets.SecretClient
	Scheme            *runtime.Scheme
}

func NewMockEventHubClient(secretClient secrets.SecretClient, scheme *runtime.Scheme) *mockEventHubManager {
	return &mockEventHubManager{
		SecretClient:      secretClient,
		Scheme:            scheme,
		eventHubResources: []eventHubResource{},
	}
}

func (manager *mockEventHubManager) DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error) {
	hubs := manager.eventHubResources

	index, _ := findEventHub(hubs, func(g eventHubResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("eventhub not found")
	}

	manager.eventHubResources = append(hubs[:index], hubs[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

func (manager *mockEventHubManager) CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, messageRetentionInDays int32, partitionCount int32, captureDescription *eventhub.CaptureDescription) (eventhub.Model, error) {

	var eventHub = eventhub.Model{
		Response: helpers.GetRestResponse(201),
		Properties: &eventhub.Properties{
			MessageRetentionInDays: to.Int64Ptr(int64(messageRetentionInDays)),
			PartitionCount:         to.Int64Ptr(int64(partitionCount)),
			Status:                 "",
			CaptureDescription:     captureDescription,
		},
		Name: &eventHubName,
	}
	manager.eventHubResources = append(manager.eventHubResources, eventHubResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eventHubName:      eventHubName,
		eventHub:          eventHub,
		eventHubAccesses:  []eventHubAccess{},
	})

	return eventHub, nil
}

func (manager *mockEventHubManager) GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubs := manager.eventHubResources

	index, hub := findEventHub(hubs, func(g eventHubResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName
	})

	if index == -1 {
		return eventhub.Model{}, errors.New("eventhub not found")
	}

	return hub.eventHub, nil
}

func (manager *mockEventHubManager) getHubAccess(resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (eventHubResource, int, eventHubAccess, error) {
	hubs := manager.eventHubResources
	hubIndex, hub := findEventHub(hubs, func(g eventHubResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName
	})
	if hubIndex == -1 {
		return eventHubResource{}, 0, eventHubAccess{}, errors.New("eventhub not found")
	}
	authRules := hub.eventHubAccesses
	ruleIndex, rule := findAccess(authRules, authorizationRuleName)

	return hub, ruleIndex, rule, nil
}

func (manager *mockEventHubManager) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (eventhub.AuthorizationRule, error) {
	hub, accessIndex, _, err := manager.getHubAccess(resourceGroupName, namespaceName, eventHubName, authorizationRuleName)
	if err != nil {
		return eventhub.AuthorizationRule{}, err
	}

	if accessIndex == -1 {
		hub.eventHubAccesses = append(hub.eventHubAccesses, eventHubAccess{
			rule: parameters,
			keys: eventhub.AccessKeys{
				Response:                       helpers.GetRestResponse(http.StatusOK),
				PrimaryConnectionString:        to.StringPtr(pkghelpers.RandomString(40)),
				SecondaryConnectionString:      to.StringPtr(pkghelpers.RandomString(40)),
				AliasPrimaryConnectionString:   to.StringPtr(pkghelpers.RandomString(40)),
				AliasSecondaryConnectionString: to.StringPtr(pkghelpers.RandomString(40)),
				PrimaryKey:                     to.StringPtr(pkghelpers.RandomString(15)),
				SecondaryKey:                   to.StringPtr(pkghelpers.RandomString(15)),
				KeyName:                        to.StringPtr(pkghelpers.RandomString(10)),
			},
		})
	} else {
		hub.eventHubAccesses[accessIndex].rule = parameters
	}

	return parameters, nil
}

func (manager *mockEventHubManager) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (eventhub.AccessKeys, error) {
	_, accessIndex, access, err := manager.getHubAccess(resourceGroupName, namespaceName, eventHubName, authorizationRuleName)

	if err != nil {
		return eventhub.AccessKeys{}, err
	}

	if accessIndex == -1 {
		return eventhub.AccessKeys{}, errors.New("eventhub access rule not found")
	}

	return access.keys, nil
}

func (e *mockEventHubManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := e.convert(obj)
	if err != nil {
		return false, err
	}

	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	partitionCount := instance.Spec.Properties.PartitionCount
	messageRetentionInDays := instance.Spec.Properties.MessageRetentionInDays
	secretName := instance.Spec.SecretName
	captureDescription := instance.Spec.Properties.CaptureDescription

	if len(secretName) == 0 {
		secretName = eventhubName
		instance.Spec.SecretName = eventhubName
	}

	// write information back to instance
	instance.Status.Provisioning = true

	capturePtr := getCaptureDescriptionPtr(captureDescription)

	_, err = e.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName, messageRetentionInDays, partitionCount, capturePtr)
	if err != nil {
		instance.Status.Provisioning = false
		return false, err
	}

	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

func (e *mockEventHubManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := e.convert(obj)
	if err != nil {
		return false, err
	}

	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	_, err = e.DeleteHub(ctx, resourcegroup, eventhubNamespace, eventhubName)

	return true, err
}

func (e *mockEventHubManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

func (g *mockEventHubManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (e *mockEventHubManager) convert(obj runtime.Object) (*azurev1alpha1.Eventhub, error) {
	local, ok := obj.(*azurev1alpha1.Eventhub)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

const storageAccountResourceFmt = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s"

func getCaptureDescriptionPtr(captureDescription azurev1alpha1.CaptureDescription) *eventhub.CaptureDescription {
	// add capture details
	var capturePtr *eventhub.CaptureDescription

	storage := captureDescription.Destination.StorageAccount
	storageAccountResourceID := fmt.Sprintf(storageAccountResourceFmt, config.SubscriptionID(), storage.ResourceGroup, storage.AccountName)

	if captureDescription.Enabled {
		capturePtr = &eventhub.CaptureDescription{
			Enabled:           to.BoolPtr(true),
			Encoding:          eventhub.Avro,
			IntervalInSeconds: &captureDescription.IntervalInSeconds,
			SizeLimitInBytes:  &captureDescription.SizeLimitInBytes,
			Destination: &eventhub.Destination{
				Name: &captureDescription.Destination.Name,
				DestinationProperties: &eventhub.DestinationProperties{
					StorageAccountResourceID: &storageAccountResourceID,
					BlobContainer:            &captureDescription.Destination.BlobContainer,
					ArchiveNameFormat:        &captureDescription.Destination.ArchiveNameFormat,
				},
			},
			SkipEmptyArchives: to.BoolPtr(true),
		}
	}
	return capturePtr
}
