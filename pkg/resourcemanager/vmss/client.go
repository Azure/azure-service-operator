// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmss

import (
	"context"
	"fmt"
	"strings"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"k8s.io/apimachinery/pkg/runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzureVMScaleSetClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVMScaleSetClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVMScaleSetClient {
	return &AzureVMScaleSetClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getVMScaleSetClient(creds config.Credentials) compute.VirtualMachineScaleSetsClient {
	computeClient := compute.NewVirtualMachineScaleSetsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (c *AzureVMScaleSetClient) CreateVMScaleSet(ctx context.Context, location string, resourceGroupName string, resourceName string, vmSize string, capacity int64, osType string, adminUserName string, adminPassword string, sshPublicKeyData string, platformImageURN string, vnetName string, subnetName string, loadBalancerName string, backendAddressPoolName string, inboundNatPoolName string) (future compute.VirtualMachineScaleSetsCreateOrUpdateFuture, err error) {

	client := getVMScaleSetClient(c.Creds)

	// Construct OS Profile
	provisionVMAgent := true
	platformImageUrnTokens := strings.Split(platformImageURN, ":")

	adminPasswordInput := ""
	adminPasswordBase64Decoded := helpers.FromBase64EncodedString(adminPassword)
	if adminPasswordBase64Decoded != "" {
		adminPasswordInput = adminPasswordBase64Decoded
	}

	sshKeyPath := fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUserName)
	sshKeysToAdd := []compute.SSHPublicKey{
		{
			Path:    &sshKeyPath,
			KeyData: &sshPublicKeyData,
		},
	}
	linuxProfile := compute.VirtualMachineScaleSetOSProfile{
		ComputerNamePrefix: &resourceName,
		AdminUsername:      &adminUserName,
		AdminPassword:      &adminPasswordInput,
		LinuxConfiguration: &compute.LinuxConfiguration{
			SSH: &compute.SSHConfiguration{
				PublicKeys: &sshKeysToAdd,
			},
			ProvisionVMAgent: &provisionVMAgent,
		},
	}

	windowsProfile := compute.VirtualMachineScaleSetOSProfile{
		ComputerNamePrefix: &resourceName,
		AdminUsername:      &adminUserName,
		AdminPassword:      &adminPasswordInput,
		WindowsConfiguration: &compute.WindowsConfiguration{
			ProvisionVMAgent: &provisionVMAgent,
		},
	}

	osProfile := linuxProfile
	if osType == "Windows" {
		osProfile = windowsProfile
	}

	// Construct Network Profile
	vmssIPConfigName := resourceName + "-ipconfig"
	publicIPAddressConfigurationName := resourceName + "-pip"
	idleTimeoutInMinutes := int32(15)
	subnetIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"virtualNetworks",
		vnetName,
		"subnets",
		subnetName,
	)
	bePoolIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"loadBalancers",
		loadBalancerName,
		"backendAddressPools",
		backendAddressPoolName,
	)
	natPoolIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"loadBalancers",
		loadBalancerName,
		"inboundNatPools",
		inboundNatPoolName,
	)
	ipConfigs := []compute.VirtualMachineScaleSetIPConfiguration{
		{
			Name: &vmssIPConfigName,
			VirtualMachineScaleSetIPConfigurationProperties: &compute.VirtualMachineScaleSetIPConfigurationProperties{
				Subnet: &compute.APIEntityReference{
					ID: &subnetIDInput,
				},
				PublicIPAddressConfiguration: &compute.VirtualMachineScaleSetPublicIPAddressConfiguration{
					Name: &publicIPAddressConfigurationName,
					VirtualMachineScaleSetPublicIPAddressConfigurationProperties: &compute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
						IdleTimeoutInMinutes: &idleTimeoutInMinutes,
					},
				},
				LoadBalancerBackendAddressPools: &[]compute.SubResource{
					{
						ID: &bePoolIDInput,
					},
				},
				LoadBalancerInboundNatPools: &[]compute.SubResource{
					{
						ID: &natPoolIDInput,
					},
				},
			},
		},
	}

	isPrimaryNic := true
	nicConfigsToAdd := []compute.VirtualMachineScaleSetNetworkConfiguration{
		{
			Name: &resourceName,
			VirtualMachineScaleSetNetworkConfigurationProperties: &compute.VirtualMachineScaleSetNetworkConfigurationProperties{
				Primary:          &isPrimaryNic,
				IPConfigurations: &ipConfigs,
			},
		},
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		compute.VirtualMachineScaleSet{
			Location: &location,
			Sku: &compute.Sku{
				Name:     &vmSize,
				Capacity: &capacity,
			},
			VirtualMachineScaleSetProperties: &compute.VirtualMachineScaleSetProperties{
				UpgradePolicy: &compute.UpgradePolicy{
					Mode: compute.Automatic,
				},
				VirtualMachineProfile: &compute.VirtualMachineScaleSetVMProfile{
					StorageProfile: &compute.VirtualMachineScaleSetStorageProfile{
						OsDisk: &compute.VirtualMachineScaleSetOSDisk{
							Caching:      compute.CachingTypesReadOnly,
							CreateOption: compute.DiskCreateOptionTypesFromImage,
						},
						ImageReference: &compute.ImageReference{
							Publisher: &platformImageUrnTokens[0],
							Offer:     &platformImageUrnTokens[1],
							Sku:       &platformImageUrnTokens[2],
							Version:   &platformImageUrnTokens[3],
						},
					},
					OsProfile: &osProfile,
					NetworkProfile: &compute.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: &nicConfigsToAdd,
					},
				},
			},
		},
	)

	return future, err
}

func (c *AzureVMScaleSetClient) DeleteVMScaleSet(ctx context.Context, vmssName string, resourcegroup string) (status string, err error) {

	client := getVMScaleSetClient(c.Creds)

	_, err = client.Get(ctx, resourcegroup, vmssName)
	if err == nil { // vmss present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmssName)
		return future.Status(), err
	}
	// VM  not present so return success anyway
	return "VMSS not present", nil

}

func (c *AzureVMScaleSetClient) GetVMScaleSet(ctx context.Context, resourcegroup string, vmssName string) (vmss compute.VirtualMachineScaleSet, err error) {

	client := getVMScaleSetClient(c.Creds)
	return client.Get(ctx, resourcegroup, vmssName)
}

func (p *AzureVMScaleSetClient) AddVMScaleSetCredsToSecrets(ctx context.Context, data map[string][]byte, instance *azurev1alpha1.AzureVMScaleSet) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	err := p.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(p.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *AzureVMScaleSetClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureVMScaleSet) (map[string][]byte, error) {
	secret := map[string][]byte{}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	if stored, err := p.SecretClient.Get(ctx, secretKey); err == nil {
		return stored, nil
	}

	randomPassword := helpers.NewPassword()
	secret["password"] = []byte(randomPassword)

	return secret, nil
}
