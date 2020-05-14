// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmext

import (
	"context"
	"encoding/json"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type AzureVirtualMachineExtensionClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVirtualMachineExtensionClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVirtualMachineExtensionClient {
	return &AzureVirtualMachineExtensionClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getVirtualMachineExtensionClient() compute.VirtualMachineExtensionsClient {
	computeClient := compute.NewVirtualMachineExtensionsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (m *AzureVirtualMachineExtensionClient) CreateVirtualMachineExtension(ctx context.Context, location string, resourceGroupName string, vmName string, extName string, autoUpgradeMinorVersion bool, forceUpdateTag string, publisher string, typeName string, typeHandlerVersion string, settings string, protectedSettings string) (future compute.VirtualMachineExtensionsCreateOrUpdateFuture, err error) {

	client := getVirtualMachineExtensionClient()

	var extensionSettings map[string]*string
	json.Unmarshal([]byte(settings), &extensionSettings)

	var extensionProtectedSettings map[string]*string
	json.Unmarshal([]byte(protectedSettings), &extensionProtectedSettings)

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		vmName,
		extName,
		compute.VirtualMachineExtension{
			Location: &location,
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				ForceUpdateTag:          &forceUpdateTag,
				Publisher:               &publisher,
				Type:                    &typeName,
				TypeHandlerVersion:      &typeHandlerVersion,
				AutoUpgradeMinorVersion: &autoUpgradeMinorVersion,
				Settings:                &extensionSettings,
				ProtectedSettings:       &extensionProtectedSettings,
			},
		},
	)

	return future, err
}

func (m *AzureVirtualMachineExtensionClient) DeleteVirtualMachineExtension(ctx context.Context, extName string, vmName string, resourcegroup string) (status string, err error) {

	client := getVirtualMachineExtensionClient()

	_, err = client.Get(ctx, resourcegroup, vmName, extName, "")
	if err == nil { // vm present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmName, extName)
		return future.Status(), err
	}
	// VM extension not present so return success anyway
	return "VM extension not present", nil

}

func (m *AzureVirtualMachineExtensionClient) GetVirtualMachineExtension(ctx context.Context, resourcegroup string, vmName string, extName string) (vm compute.VirtualMachineExtension, err error) {

	client := getVirtualMachineExtensionClient()

	return client.Get(ctx, resourcegroup, vmName, extName, "")
}

func (p *AzureVirtualMachineExtensionClient) AddVirtualMachineExtensionCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.AzureVirtualMachineExtension) error {
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

func (p *AzureVirtualMachineExtensionClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureVirtualMachineExtension) (map[string][]byte, error) {
	name := instance.Name

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := p.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	emptyProtectedSettings := "{}"
	secret["protectedSettings"] = []byte(emptyProtectedSettings)

	return secret, nil
}
