// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmss

import (
	"context"
	"fmt"
	"strings"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type AzureVMScaleSetClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVMScaleSetClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVMScaleSetClient {
	return &AzureVMScaleSetClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getVMScaleSetClient() compute.VirtualMachineScaleSetsClient {
	computeClient := compute.NewVirtualMachineScaleSetsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (m *AzureVMScaleSetClient) CreateVMScaleSet(ctx context.Context, location string, resourceGroupName string, resourceName string, vmSize string, capacity int64, osType string, adminUserName string, adminPassword string, sshPublicKeyData string, platformImageURN string, vnetName string, subnetName string, loadBalancerName string, backendAddressPoolName string, inboundNatPoolName string) (future compute.VirtualMachineScaleSetsCreateOrUpdateFuture, err error) {

	client := getVMScaleSetClient()

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
		compute.SSHPublicKey{
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
		compute.VirtualMachineScaleSetIPConfiguration{
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
					compute.SubResource{
						ID: &bePoolIDInput,
					},
				},
				LoadBalancerInboundNatPools: &[]compute.SubResource{
					compute.SubResource{
						ID: &natPoolIDInput,
					},
				},
			},
		},
	}

	isPrimaryNic := true
	nicConfigsToAdd := []compute.VirtualMachineScaleSetNetworkConfiguration{
		compute.VirtualMachineScaleSetNetworkConfiguration{
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

func (m *AzureVMScaleSetClient) DeleteVMScaleSet(ctx context.Context, vmssName string, resourcegroup string) (status string, err error) {

	client := getVMScaleSetClient()

	_, err = client.Get(ctx, resourcegroup, vmssName)
	if err == nil { // vmss present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmssName)
		return future.Status(), err
	}
	// VM  not present so return success anyway
	return "VMSS not present", nil

}

func (m *AzureVMScaleSetClient) GetVMScaleSet(ctx context.Context, resourcegroup string, vmssName string) (vmss compute.VirtualMachineScaleSet, err error) {

	client := getVMScaleSetClient()
	return client.Get(ctx, resourcegroup, vmssName)
}

func (p *AzureVMScaleSetClient) AddVMScaleSetCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.AzureVMScaleSet) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	err := p.SecretClient.Upsert(ctx,
		key,
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
	name := instance.Name

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := p.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	randomPassword := helpers.NewPassword()
	secret["password"] = []byte(randomPassword)

	return secret, nil
}
