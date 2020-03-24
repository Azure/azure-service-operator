// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package namespacenetworkrule

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type NamespaceNetworkRuleManager interface {
	CreateNetworkRuleSet(ctx context.Context, groupname string, namespace string, rules eventhub.NetworkRuleSet) (result eventhub.NetworkRuleSet, err error)
	GetNetworkRuleSet(ctx context.Context, groupName string, namespace string) (ruleset eventhub.NetworkRuleSet, err error)
	DeleteNetworkRuleSet(ctx context.Context, groupName string, namespace string) (err error)
	resourcemanager.ARMClient
}
