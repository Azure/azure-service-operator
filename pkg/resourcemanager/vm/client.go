// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vm

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

type AzureVirtualMachineClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVirtualMachineClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVirtualMachineClient {
	return &AzureVirtualMachineClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getVirtualMachineClient(creds config.Credentials) compute.VirtualMachinesClient {
	computeClient := compute.NewVirtualMachinesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (c *AzureVirtualMachineClient) CreateVirtualMachine(ctx context.Context, location string, resourceGroupName string, resourceName string, vmSize string, osType string, adminUserName string, adminPassword string, sshPublicKeyData string, networkInterfaceName string, platformImageURN string) (future compute.VirtualMachinesCreateOrUpdateFuture, err error) {

	client := getVirtualMachineClient(c.Creds)

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
		{
			ID: &nicIDInput,
			NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
				Primary: &addAsPrimaryNic,
			},
		},
	}

	sshKeyPath := fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUserName)
	sshKeysToAdd := []compute.SSHPublicKey{
		{
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

func (c *AzureVirtualMachineClient) DeleteVirtualMachine(ctx context.Context, vmName string, resourcegroup string) (status string, err error) {

	client := getVirtualMachineClient(c.Creds)

	_, err = client.Get(ctx, resourcegroup, vmName, "")
	if err == nil { // vm present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmName)
		return future.Status(), err
	}
	// VM  not present so return success anyway
	return "VM not present", nil

}

func (c *AzureVirtualMachineClient) GetVirtualMachine(ctx context.Context, resourcegroup string, vmName string) (vm compute.VirtualMachine, err error) {

	client := getVirtualMachineClient(c.Creds)

	return client.Get(ctx, resourcegroup, vmName, "")
}

func (c *AzureVirtualMachineClient) AddVirtualMachineCredsToSecrets(ctx context.Context, data map[string][]byte, instance *azurev1alpha1.AzureVirtualMachine) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	err := c.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(c.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *AzureVirtualMachineClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureVirtualMachine) (map[string][]byte, error) {
	secret := map[string][]byte{}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	if stored, err := c.SecretClient.Get(ctx, secretKey); err == nil {
		return stored, nil
	}

	randomPassword := helpers.NewPassword()
	secret["password"] = []byte(randomPassword)

	return secret, nil
}
