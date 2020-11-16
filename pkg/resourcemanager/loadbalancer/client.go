// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package loadbalancer

import (
	"context"
	"strings"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzureLoadBalancerClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureLoadBalancerClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureLoadBalancerClient {
	return &AzureLoadBalancerClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewAzureLoadBalancerClient(creds, secretClient, scheme)
}

func getLoadBalancerClient(creds config.Credentials) vnetwork.LoadBalancersClient {
	lbClient := vnetwork.NewLoadBalancersClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	lbClient.Authorizer = a
	lbClient.AddToUserAgent(config.UserAgent())
	return lbClient
}

func (c *AzureLoadBalancerClient) CreateLoadBalancer(ctx context.Context, location string, resourceGroupName string, resourceName string, publicIPAddressName string, backendAddressPoolName string, inboundNatPoolName string, frontendPortRangeStart int, frontendPortRangeEnd int, backendPort int) (future vnetwork.LoadBalancersCreateOrUpdateFuture, err error) {

	client := getLoadBalancerClient(c.Creds)

	publicIPAddressIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"publicIPAddresses",
		publicIPAddressName,
		"",
		"",
	)

	publicIPAddress := vnetwork.PublicIPAddress{
		ID: &publicIPAddressIDInput,
	}

	frontEndIPConfigName := strings.Join([]string{resourceName, "IpCfg"}, "-")
	frontendIPConfiguration := vnetwork.FrontendIPConfiguration{
		Name: &frontEndIPConfigName,
		FrontendIPConfigurationPropertiesFormat: &vnetwork.FrontendIPConfigurationPropertiesFormat{
			PublicIPAddress: &publicIPAddress,
		},
	}

	var ipConfigsToAdd []vnetwork.FrontendIPConfiguration
	ipConfigsToAdd = append(
		ipConfigsToAdd,
		frontendIPConfiguration,
	)

	frontendIPConfigId := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"loadBalancers",
		resourceName,
		"frontendIPConfigurations",
		frontEndIPConfigName,
	)

	frontendIPConfigurationSubResource := vnetwork.SubResource{
		ID: &frontendIPConfigId,
	}

	var bcPoolsToAdd []vnetwork.BackendAddressPool
	bcPoolsToAdd = append(
		bcPoolsToAdd,
		vnetwork.BackendAddressPool{
			Name: &backendAddressPoolName,
		},
	)

	frontendPortRangeStartInt32 := int32(frontendPortRangeStart)
	frontendPortRangeEndInt32 := int32(frontendPortRangeEnd)
	backendPortInt32 := int32(backendPort)

	var natPoolsToAdd []vnetwork.InboundNatPool
	natPoolsToAdd = append(
		natPoolsToAdd,
		vnetwork.InboundNatPool{
			Name: &inboundNatPoolName,
			InboundNatPoolPropertiesFormat: &vnetwork.InboundNatPoolPropertiesFormat{
				FrontendIPConfiguration: &frontendIPConfigurationSubResource,
				Protocol:                vnetwork.TransportProtocolTCP,
				FrontendPortRangeStart:  &frontendPortRangeStartInt32,
				FrontendPortRangeEnd:    &frontendPortRangeEndInt32,
				BackendPort:             &backendPortInt32,
			},
		},
	)

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		vnetwork.LoadBalancer{
			Location: &location,
			LoadBalancerPropertiesFormat: &vnetwork.LoadBalancerPropertiesFormat{
				FrontendIPConfigurations: &ipConfigsToAdd,
				BackendAddressPools:      &bcPoolsToAdd,
				InboundNatPools:          &natPoolsToAdd,
			},
		},
	)

	return future, err
}

func (c *AzureLoadBalancerClient) DeleteLoadBalancer(ctx context.Context, loadBalancerName string, resourcegroup string) (status string, err error) {

	client := getLoadBalancerClient(c.Creds)

	_, err = client.Get(ctx, resourcegroup, loadBalancerName, "")
	if err == nil { // load balancer present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, loadBalancerName)
		return future.Status(), err
	}
	// load balancer not present so return success anyway
	return "load balancer not present", nil

}

func (c *AzureLoadBalancerClient) GetLoadBalancer(ctx context.Context, resourcegroup string, loadBalancerName string) (lb vnetwork.LoadBalancer, err error) {

	client := getLoadBalancerClient(c.Creds)

	return client.Get(ctx, resourcegroup, loadBalancerName, "")
}
