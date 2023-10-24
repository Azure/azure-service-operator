// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

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
	creds config.Credentials
}

func getNamespacesClient(creds config.Credentials) (eventhub.NamespacesClient, error) {
	nsClient := eventhub.NewNamespacesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	auth, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return eventhub.NamespacesClient{}, err
	}
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(config.UserAgent())
	return nsClient, nil
}

func NewEventHubNamespaceClient(creds config.Credentials) *azureEventHubNamespaceManager {
	return &azureEventHubNamespaceManager{creds: creds}
}

// DeleteNamespace deletes an existing namespace. This operation also removes all associated resources under the namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (m *azureEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (autorest.Response, error) {

	nsClient, err := getNamespacesClient(m.creds)
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	future, err := nsClient.Delete(ctx,
		resourceGroupName,
		namespaceName)

	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	return autorest.Response{Response: future.Response()}, err
}

// Get gets the description of the specified namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (m *azureEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	nsClient, err := getNamespacesClient(m.creds)
	if err != nil {
		return nil, err
	}

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
func (m *azureEventHubNamespaceManager) CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	nsClient, err := getNamespacesClient(m.creds)
	if err != nil {
		return nil, err
	}

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

func (m *azureEventHubNamespaceManager) CreateNamespace(ctx context.Context, resourceGroupName string, namespaceName string, location string, sku v1alpha1.EventhubNamespaceSku, properties v1alpha1.EventhubNamespaceProperties) (eventhub.EHNamespace, error) {
	nsClient, err := getNamespacesClient(m.creds)
	if err != nil {
		return eventhub.EHNamespace{}, err
	}

	// Construct the Sku struct for the namespace
	namespaceSku := eventhub.Sku{
		Capacity: to.Int32Ptr(1),
	}

	if sku.Capacity != 0 {
		namespaceSku.Capacity = &sku.Capacity
	}

	namespaceSku.Name = eventhub.Standard
	if strings.ToLower(sku.Name) == "basic" {
		namespaceSku.Name = eventhub.Basic
	}

	namespaceSku.Tier = eventhub.SkuTierStandard
	if strings.ToLower(sku.Tier) == "basic" {
		namespaceSku.Tier = eventhub.SkuTierBasic
	}

	namespaceProperties := eventhub.EHNamespaceProperties{}

	// Construct the properties struct if tier is not "Basic"
	// These properties are only valid if tier is "Standard"
	// We need to do this as the SDK doesnt handle this gracefully if we set them,
	// and causes the eventhubs to not open up in the portal

	if namespaceSku.Tier == eventhub.SkuTierStandard {
		namespaceProperties = eventhub.EHNamespaceProperties{
			IsAutoInflateEnabled:   &properties.IsAutoInflateEnabled,
			MaximumThroughputUnits: &properties.MaximumThroughputUnits,
			KafkaEnabled:           &properties.KafkaEnabled,
		}
	}

	future, err := nsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventhub.EHNamespace{
			Location:              to.StringPtr(location),
			EHNamespaceProperties: &namespaceProperties,
			Sku:                   &namespaceSku,
		},
	)

	if err != nil {
		return eventhub.EHNamespace{}, err
	}

	return future.Result(nsClient)
}

func (m *azureEventHubNamespaceManager) CreateNetworkRuleSet(ctx context.Context, groupname string, namespace string, rules eventhub.NetworkRuleSet) (result eventhub.NetworkRuleSet, err error) {
	namespaceclient, err := getNamespacesClient(m.creds)
	if err != nil {
		return eventhub.NetworkRuleSet{}, err
	}

	return namespaceclient.CreateOrUpdateNetworkRuleSet(ctx, groupname, namespace, rules)
}

func (m *azureEventHubNamespaceManager) GetNetworkRuleSet(ctx context.Context, groupName string, namespace string) (ruleset eventhub.NetworkRuleSet, err error) {
	namespaceclient, err := getNamespacesClient(m.creds)
	if err != nil {
		return eventhub.NetworkRuleSet{}, err
	}

	return namespaceclient.GetNetworkRuleSet(ctx, groupName, namespace)
}

func (m *azureEventHubNamespaceManager) DeleteNetworkRuleSet(ctx context.Context, groupName string, namespace string) (result eventhub.NetworkRuleSet, err error) {
	namespaceclient, err := getNamespacesClient(m.creds)
	if err != nil {
		return eventhub.NetworkRuleSet{}, err
	}

	// SDK does not have a DeleteNetworkRuleSet function, so setting rules to empty
	// and calling Create to delete the rules

	emptyrules := eventhub.NetworkRuleSet{
		NetworkRuleSetProperties: &eventhub.NetworkRuleSetProperties{
			DefaultAction: eventhub.Allow,
		},
	}
	return namespaceclient.CreateOrUpdateNetworkRuleSet(ctx, groupName, namespace, emptyrules)

}

// ParseNetworkRules parses the network rules specified in the yaml and converts to the SDK struct
func ParseNetworkRules(networkRule *v1alpha1.EventhubNamespaceNetworkRule) eventhub.NetworkRuleSet {

	defaultAction := eventhub.Allow

	if len(networkRule.DefaultAction) != 0 {
		defActionStr := string(networkRule.DefaultAction)

		switch defActionStr {
		case "allow":
			defaultAction = eventhub.Allow
		case "deny":
			defaultAction = eventhub.Deny
		default:
			defaultAction = eventhub.Deny
		}
	}

	var vNetRulesSet []eventhub.NWRuleSetVirtualNetworkRules
	if networkRule.VirtualNetworkRules != nil {
		for _, i := range *networkRule.VirtualNetworkRules {
			subnetID := i.SubnetID
			ignoreEndpoint := i.IgnoreMissingServiceEndpoint
			vNetRulesSet = append(vNetRulesSet, eventhub.NWRuleSetVirtualNetworkRules{
				Subnet: &eventhub.Subnet{
					ID: &subnetID,
				},
				IgnoreMissingVnetServiceEndpoint: &ignoreEndpoint,
			})
		}
	}

	var ipRulesSet []eventhub.NWRuleSetIPRules
	if networkRule.IPRules != nil {
		for _, i := range *networkRule.IPRules {
			ipmask := i.IPMask
			ipRulesSet = append(ipRulesSet, eventhub.NWRuleSetIPRules{
				IPMask: ipmask,
				Action: eventhub.NetworkRuleIPActionAllow, // only applicable value
			})
		}
	}

	return eventhub.NetworkRuleSet{
		NetworkRuleSetProperties: &eventhub.NetworkRuleSetProperties{
			DefaultAction:       defaultAction,
			VirtualNetworkRules: &vNetRulesSet,
			IPRules:             &ipRulesSet,
		},
	}

}

func (ns *azureEventHubNamespaceManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return false, err
	}

	// set a spec hash if one hasn't been set
	hash := helpers.Hash256(instance.Spec)
	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	if instance.Status.SpecHash == "" {
		instance.Status.SpecHash = hash
	}

	namespaceLocation := instance.Spec.Location
	namespaceName := instance.Name
	resourcegroup := instance.Spec.ResourceGroup
	eventHubNSSku := instance.Spec.Sku
	eventHubNSProperties := instance.Spec.Properties

	// write information back to instance
	instance.Status.Provisioning = true
	instance.Status.Provisioned = false
	instance.Status.FailedProvisioning = false

	evhns, err := ns.GetNamespace(ctx, resourcegroup, namespaceName)
	if err == nil {
		if instance.Status.SpecHash == hash { // SpecHash matches what is asked for
			instance.Status.State = *evhns.ProvisioningState
			instance.Status.Message = "namespace exists but may not be ready"

			if *evhns.ProvisioningState == "Succeeded" {
				// Provision network rules

				instance.Status.Message = "namespace created, creating network rule"

				networkRuleSet := eventhub.NetworkRuleSet{
					NetworkRuleSetProperties: &eventhub.NetworkRuleSetProperties{
						DefaultAction: eventhub.Allow,
					},
				}
				if instance.Spec.NetworkRule != nil {
					networkRuleSet = ParseNetworkRules(instance.Spec.NetworkRule)
				}

				_, err := ns.CreateNetworkRuleSet(ctx, resourcegroup, namespaceName, networkRuleSet)
				if err != nil {
					instance.Status.Message = errhelp.StripErrorIDs(err)
					azerr := errhelp.NewAzureError(err)

					ignorableErrors := []string{
						errhelp.ResourceGroupNotFoundErrorCode,
						errhelp.ParentNotFoundErrorCode,
						errhelp.ResourceNotFound,
					}

					unrecoverableErrors := []string{
						errhelp.BadRequest, // error we get when namespace is "Basic" SKU, invalid rule etc.
						"InvalidSkuForNetworkRuleSet",
					}

					if helpers.ContainsString(ignorableErrors, azerr.Type) {
						instance.Status.Provisioning = false
						return false, nil //requeue
					}

					if helpers.ContainsString(unrecoverableErrors, azerr.Type) {
						instance.Status.Provisioning = false
						return true, nil // Stop reconciliation
					}

					return false, err
				}

				instance.Status.SpecHash = hash
				instance.Status.Message = resourcemanager.SuccessMsg
				instance.Status.Provisioned = true
				instance.Status.Provisioning = false
				instance.Status.ResourceId = *evhns.ID
				return true, nil
			}

			return false, nil
		}
	}

	// create Event Hubs namespace
	newNs, err := ns.CreateNamespace(ctx, resourcegroup, namespaceName, namespaceLocation, eventHubNSSku, eventHubNSProperties)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		if strings.Contains(azerr.Type, errhelp.AsyncOpIncompleteError) {
			// resource creation request sent to Azure
			instance.Status.SpecHash = hash
			return false, nil
		}

		instance.Status.Provisioning = false
		if strings.Contains(azerr.Type, errhelp.ResourceGroupNotFoundErrorCode) || strings.Contains(err.Error(), "validation failed") {
			return false, nil
		}

		return false, errors.Wrap(err, "EventhubNamespace create error")
	}

	// write information back to instance
	instance.Status.State = *newNs.ProvisioningState
	return true, nil
}

func (ns *azureEventHubNamespaceManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return false, err
	}

	namespaceName := instance.Name
	resourcegroup := instance.Spec.ResourceGroup

	resp, err := ns.DeleteNamespace(ctx, resourcegroup, namespaceName)
	if err != nil {
		azerr := errhelp.NewAzureError(err)

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
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		// some error we don't know about or can't handle happened
		instance.Status.Provisioning = false

		return true, errors.Wrap(err, "EventhubNamespace delete error")

	}

	// if NoContent response is returned, namespace is already gone
	if resp.StatusCode == http.StatusNoContent {
		return false, nil
	}

	return true, nil
}

func (ns *azureEventHubNamespaceManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := ns.convert(obj)
	if err != nil {
		return nil, err
	}

	key := types.NamespacedName{Namespace: instance.Namespace, Name: instance.Spec.ResourceGroup}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &v1alpha1.ResourceGroup{}},
	}, nil
}

func (g *azureEventHubNamespaceManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (ns *azureEventHubNamespaceManager) convert(obj runtime.Object) (*azurev1alpha1.EventhubNamespace, error) {
	local, ok := obj.(*azurev1alpha1.EventhubNamespace)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
