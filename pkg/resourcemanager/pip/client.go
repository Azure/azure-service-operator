// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package pip

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzurePublicIPAddressClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzurePublicIPAddressClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzurePublicIPAddressClient {
	return &AzurePublicIPAddressClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewAzurePublicIPAddressClient(creds, secretClient, scheme)
}

func getPublicIPAddressClient(creds config.Credentials) vnetwork.PublicIPAddressesClient {
	pipClient := vnetwork.NewPublicIPAddressesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	pipClient.Authorizer = a
	pipClient.AddToUserAgent(config.UserAgent())
	return pipClient
}

func (m *AzurePublicIPAddressClient) CreatePublicIPAddress(ctx context.Context,
	location string,
	resourceGroupName string,
	resourceName string,
	publicIPAllocationMethod string,
	idleTimeoutInMinutes int,
	publicIPAddressVersion string,
	skuName string,
	ipTags map[string]string) (future vnetwork.PublicIPAddressesCreateOrUpdateFuture, err error) {

	client := getPublicIPAddressClient(m.Creds)

	publicIPAllocationMethodField := vnetwork.Static
	if publicIPAllocationMethod == string(vnetwork.Dynamic) {
		publicIPAllocationMethodField = vnetwork.Dynamic
	}
	idleTimeoutInMinutesInt32 := (int32)(idleTimeoutInMinutes)
	publicIPAddressVersionField := vnetwork.IPv4
	if publicIPAddressVersion == string(vnetwork.IPv6) {
		publicIPAddressVersionField = vnetwork.IPv6
	}
	skuNameInput := vnetwork.PublicIPAddressSkuNameBasic
	if skuName == string(vnetwork.PublicIPAddressSkuNameStandard) {
		skuNameInput = vnetwork.PublicIPAddressSkuNameStandard
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		vnetwork.PublicIPAddress{
			Location: &location,
			PublicIPAddressPropertiesFormat: &vnetwork.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: publicIPAllocationMethodField,
				IdleTimeoutInMinutes:     &idleTimeoutInMinutesInt32,
				PublicIPAddressVersion:   publicIPAddressVersionField,
				IPTags:                   getIPTagsForPublicIP(ipTags),
			},
			Sku: &vnetwork.PublicIPAddressSku{
				Name: skuNameInput,
			},
		},
	)

	return future, err
}

func getIPTagsForPublicIP(tags map[string]string) *[]vnetwork.IPTag {
	if tags == nil || len(tags) == 0 {
		return nil
	}

	outputTags := []vnetwork.IPTag{}
	for k, v := range tags {
		outputTags = append(outputTags, vnetwork.IPTag{
			IPTagType: &k,
			Tag:       &v,
		})
	}
	return &outputTags
}

func (m *AzurePublicIPAddressClient) DeletePublicIPAddress(ctx context.Context, publicIPAddressName string, resourcegroup string) (status string, err error) {

	client := getPublicIPAddressClient(m.Creds)

	_, err = client.Get(ctx, resourcegroup, publicIPAddressName, "")
	if err == nil { // pip present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, publicIPAddressName)
		return future.Status(), err
	}
	// pip not present so return success anyway
	return "pip not present", nil

}

func (m *AzurePublicIPAddressClient) GetPublicIPAddress(ctx context.Context, resourcegroup string, publicIPAddressName string) (pip network.PublicIPAddress, err error) {

	client := getPublicIPAddressClient(m.Creds)

	return client.Get(ctx, resourcegroup, publicIPAddressName, "")
}
