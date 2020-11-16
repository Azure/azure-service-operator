// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmext

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"k8s.io/apimachinery/pkg/runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzureVirtualMachineExtensionClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureVirtualMachineExtensionClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureVirtualMachineExtensionClient {
	return &AzureVirtualMachineExtensionClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewAzureVirtualMachineExtensionClient(creds, secretClient, scheme)
}

func getVirtualMachineExtensionClient(creds config.Credentials) compute.VirtualMachineExtensionsClient {
	computeClient := compute.NewVirtualMachineExtensionsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	computeClient.Authorizer = a
	computeClient.AddToUserAgent(config.UserAgent())
	return computeClient
}

func (c *AzureVirtualMachineExtensionClient) CreateVirtualMachineExtension(ctx context.Context, location string, resourceGroupName string, vmName string, extName string, autoUpgradeMinorVersion bool, forceUpdateTag string, publisher string, typeName string, typeHandlerVersion string, settings string, protectedSettings string) (future compute.VirtualMachineExtensionsCreateOrUpdateFuture, err error) {

	client := getVirtualMachineExtensionClient(c.Creds)

	var extensionSettings map[string]*string

	err = json.Unmarshal([]byte(settings), &extensionSettings)
	if err != nil {
		return future, err
	}

	var extensionProtectedSettings map[string]*string
	err = json.Unmarshal([]byte(protectedSettings), &extensionProtectedSettings)
	if err != nil {
		return future, err
	}

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

func (c *AzureVirtualMachineExtensionClient) DeleteVirtualMachineExtension(ctx context.Context, extName string, vmName string, resourcegroup string) (status string, err error) {

	client := getVirtualMachineExtensionClient(c.Creds)

	_, err = client.Get(ctx, resourcegroup, vmName, extName, "")
	if err == nil { // vm present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, vmName, extName)
		return future.Status(), err
	}
	// VM extension not present so return success anyway
	return "VM extension not present", nil

}

func (c *AzureVirtualMachineExtensionClient) GetVirtualMachineExtension(ctx context.Context, resourcegroup string, vmName string, extName string) (vm compute.VirtualMachineExtension, err error) {

	client := getVirtualMachineExtensionClient(c.Creds)

	return client.Get(ctx, resourcegroup, vmName, extName, "")
}

func (p *AzureVirtualMachineExtensionClient) AddVirtualMachineExtensionCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.AzureVirtualMachineExtension) error {
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

func (p *AzureVirtualMachineExtensionClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureVirtualMachineExtension) (map[string][]byte, error) {
	secret := map[string][]byte{}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	if stored, err := p.SecretClient.Get(ctx, secretKey); err == nil {
		return stored, nil
	}

	emptyProtectedSettings := "{}"
	secret["protectedSettings"] = []byte(emptyProtectedSettings)

	return secret, nil
}
