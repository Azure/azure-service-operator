// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vm

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

type AzureVirtualMachineClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVirtualMachineClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVirtualMachineClient {
	return &AzureVirtualMachineClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getVirtualMachineClient() compute.VirtualMachinesClient {
	computeClient := compute.NewVirtualMachinesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (m *AzureVirtualMachineClient) CreateVirtualMachine(ctx context.Context, location string, resourceGroupName string, resourceName string, vmSize string, osType string, adminUserName string, adminPassword string, sshPublicKeyData string, networkInterfaceName string, platformImageURN string) (future compute.VirtualMachinesCreateOrUpdateFuture, err error) {

	client := getVirtualMachineClient()

	vmSizeInput := compute.VirtualMachineSizeTypes(vmSize)
	provisionVMAgent := true
	platformImageUrnTokens := strings.Split(platformImageURN, ":")

	adminPasswordInput := ""
	adminPasswordBase64Decoded := helpers.FromBase64EncodedString(adminPassword)
	if adminPasswordBase64Decoded != "" {
		adminPasswordInput = adminPasswordBase64Decoded
	}

	addAsPrimaryNic := true
	nicIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"networkInterfaces",
		networkInterfaceName,
		"",
		"",
	)

	nicsToAdd := []compute.NetworkInterfaceReference{
		compute.NetworkInterfaceReference{
			ID: &nicIDInput,
			NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
				Primary: &addAsPrimaryNic,
			},
		},
	}

	sshKeyPath := fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUserName)
	sshKeysToAdd := []compute.SSHPublicKey{
		compute.SSHPublicKey{
			Path:    &sshKeyPath,
			KeyData: &sshPublicKeyData,
		},
	}
	linuxProfile := compute.OSProfile{
		ComputerName:  &resourceName,
		AdminUsername: &adminUserName,
		AdminPassword: &adminPasswordInput,
		LinuxConfiguration: &compute.LinuxConfiguration{
			SSH: &compute.SSHConfiguration{
				PublicKeys: &sshKeysToAdd,
			},
			ProvisionVMAgent: &provisionVMAgent,
		},
	}

	windowsProfile := compute.OSProfile{
		ComputerName:  &resourceName,
		AdminUsername: &adminUserName,
		AdminPassword: &adminPasswordInput,
		WindowsConfiguration: &compute.WindowsConfiguration{
			ProvisionVMAgent: &provisionVMAgent,
		},
	}

	osProfile := linuxProfile
	if osType == "Windows" {
		osProfile = windowsProfile
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		compute.VirtualMachine{
			Location: &location,
			VirtualMachineProperties: &compute.VirtualMachineProperties{
				HardwareProfile: &compute.HardwareProfile{
					VMSize: vmSizeInput,
				},
				StorageProfile: &compute.StorageProfile{
					ImageReference: &compute.ImageReference{
						Publisher: &platformImageUrnTokens[0],
						Offer:     &platformImageUrnTokens[1],
						Sku:       &platformImageUrnTokens[2],
						Version:   &platformImageUrnTokens[3],
					},
				},
				OsProfile: &osProfile,
				NetworkProfile: &compute.NetworkProfile{
					NetworkInterfaces: &nicsToAdd,
				},
			},
		},
	)

	return future, err
}

func (m *AzureVirtualMachineClient) DeleteVirtualMachine(ctx context.Context, vmName string, resourcegroup string) (status string, err error) {

	client := getVirtualMachineClient()

	_, err = client.Get(ctx, resourcegroup, vmName, "")
	if err == nil { // vm present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmName)
		return future.Status(), err
	}
	// VM  not present so return success anyway
	return "VM not present", nil

}

func (m *AzureVirtualMachineClient) GetVirtualMachine(ctx context.Context, resourcegroup string, vmName string) (vm compute.VirtualMachine, err error) {

	client := getVirtualMachineClient()

	return client.Get(ctx, resourcegroup, vmName, "")
}

func (p *AzureVirtualMachineClient) AddVirtualMachineCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.AzureVirtualMachine) error {
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

func (p *AzureVirtualMachineClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureVirtualMachine) (map[string][]byte, error) {
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
