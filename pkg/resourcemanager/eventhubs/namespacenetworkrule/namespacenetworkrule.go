package namespacenetworkrule

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

type AzureNamespaceNetworkRuleManager struct {
}

func NewAzureNamespaceNetworkRuleManager() *AzureNamespaceNetworkRuleManager {
	return &AzureNamespaceNetworkRuleManager{}
}

func getNamespacesClient() eventhub.NamespacesClient {
	nsClient := eventhub.NewNamespacesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(config.UserAgent())
	return nsClient
}

func (nr *AzureNamespaceNetworkRuleManager) CreateNetworkRuleSet(ctx context.Context, groupname string, namespace string, rules eventhub.NetworkRuleSet) (result eventhub.NetworkRuleSet, err error) {
	namespaceclient := getNamespacesClient()
	return namespaceclient.CreateOrUpdateNetworkRuleSet(ctx, groupname, namespace, rules)
}

func (nr *AzureNamespaceNetworkRuleManager) GetNetworkRuleSet(ctx context.Context, groupName string, namespace string) (ruleset eventhub.NetworkRuleSet, err error) {
	namespaceclient := getNamespacesClient()
	return namespaceclient.GetNetworkRuleSet(ctx, groupName, namespace)
}

func (nr *AzureNamespaceNetworkRuleManager) DeleteNetworkRuleSet(ctx context.Context, groupName string, namespace string) (result eventhub.NetworkRuleSet, err error) {
	namespaceclient := getNamespacesClient()

	// SDK does not have a DeleteNetworkRuleSet function, so setting rules to empty
	// and calling Create to delete the rules

	emptyrules := eventhub.NetworkRuleSet{
		NetworkRuleSetProperties: &eventhub.NetworkRuleSetProperties{
			DefaultAction: eventhub.Allow,
		},
	}
	return namespaceclient.CreateOrUpdateNetworkRuleSet(ctx, groupName, namespace, emptyrules)

}
