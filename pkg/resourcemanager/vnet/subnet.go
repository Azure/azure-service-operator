// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"
	"strings"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// AzureSubnetManager is the struct that the manager functions hang off
type AzureSubnetManager struct {
	creds config.Credentials
}

// NewAzureSubnetManager returns a new client for subnets
func NewAzureSubnetManager(creds config.Credentials) *AzureSubnetManager {
	return &AzureSubnetManager{creds: creds}
}

// getSubnetClient returns a new instance of an subnet client
func getSubnetClient(creds config.Credentials) (vnetwork.SubnetsClient, error) {
	client := vnetwork.NewSubnetsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		client = vnetwork.SubnetsClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// Get gets a Subnet from Azure
func (m *AzureSubnetManager) Get(ctx context.Context, resourceGroup, vnet, subnet string) (vnetwork.Subnet, error) {
	client, err := getSubnetClient(m.creds)
	if err != nil {
		return vnetwork.Subnet{}, err
	}

	return client.Get(ctx, resourceGroup, vnet, subnet, "")
}

// SubnetID models the parts of a subnet resource id
type SubnetID struct {
	Name          string
	VNet          string
	Subnet        string
	ResourceGroup string
	Subscription  string
}

// ParseSubnetID takes a resource id for a subnet and parses it into its parts
func ParseSubnetID(sid string) SubnetID {
	parts := strings.Split(sid, "/")
	subid := SubnetID{}

	for i, v := range parts {
		if i == 0 {
			continue
		}
		switch parts[i-1] {
		case "subscriptions":
			subid.Subscription = v
		case "resourceGroups":
			subid.ResourceGroup = v
		case "virtualNetworks":
			subid.VNet = v
		case "subnets":
			subid.Subnet = v
		}
	}
	return subid
}
