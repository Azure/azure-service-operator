package eventhubs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	pkghelpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type eventHubAccess struct {
	rule eventhub.AuthorizationRule
	keys eventhub.AccessKeys
}

type EventHubResource struct {
	ResourceGroupName  string
	NamespaceName      string
	EventHubName       string
	EventHub           eventhub.Model
	AuthorizationRules []eventHubAccess
}

func findEventHub(res []EventHubResource, predicate func(EventHubResource) bool) (int, EventHubResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, EventHubResource{}
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
	eventHubResources []EventHubResource
}

func (manager *mockEventHubManager) DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error) {
	return autorest.Response{}, nil
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
	manager.eventHubResources = append(manager.eventHubResources, EventHubResource{
		ResourceGroupName:  resourceGroupName,
		NamespaceName:      namespaceName,
		EventHubName:       eventHubName,
		EventHub:           eventHub,
		AuthorizationRules: []eventHubAccess{},
	})
	return eventHub, nil
}

func (manager *mockEventHubManager) GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubs := manager.eventHubResources

	index, hub := findEventHub(hubs, func(g EventHubResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName &&
			g.EventHubName == eventHubName
	})

	if index == -1 {
		return eventhub.Model{}, errors.New("eventhub not found")
	}

	return hub.EventHub, nil
}

func (manager *mockEventHubManager) getHubAccess(resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (EventHubResource, int, eventHubAccess, error) {
	hubs := manager.eventHubResources
	hubIndex, hub := findEventHub(hubs, func(g EventHubResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName &&
			g.EventHubName == eventHubName
	})
	if hubIndex == -1 {
		return EventHubResource{}, 0, eventHubAccess{}, errors.New("eventhub not found")
	}
	authRules := hub.AuthorizationRules
	ruleIndex, rule := findAccess(authRules, authorizationRuleName)

	return hub, ruleIndex, rule, nil
}

func (manager *mockEventHubManager) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (eventhub.AuthorizationRule, error) {
	hub, accessIndex, _, err := manager.getHubAccess(resourceGroupName, namespaceName, eventHubName, authorizationRuleName)
	if err != nil {
		return eventhub.AuthorizationRule{}, err
	}

	if accessIndex == -1 {
		hub.AuthorizationRules = append(hub.AuthorizationRules, eventHubAccess{
			rule: parameters,
			keys: eventhub.AccessKeys{
				Response:                       helpers.GetRestResponse(200),
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
		hub.AuthorizationRules[accessIndex].rule = parameters
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
