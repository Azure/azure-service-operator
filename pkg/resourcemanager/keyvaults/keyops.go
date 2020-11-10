// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"context"
	"fmt"
	"net/http"

	kvops "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// KeyvaultKeyClient emcompasses the methods needed for the keyops client to fulfill the ARMClient interface
type KeyvaultKeyClient struct {
	Creds          config.Credentials
	KeyvaultClient *azureKeyVaultManager
}

func NewKeyvaultKeyClient(creds config.Credentials, client *azureKeyVaultManager) *KeyvaultKeyClient {
	return &KeyvaultKeyClient{
		Creds:          creds,
		KeyvaultClient: client,
	}
}

// Ensure idempotently implements the user's requested state
func (k *KeyvaultKeyClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	instance.Status.Provisioning = true

	// Check if this KeyVault already exists and its state if it does.

	kvopsclient := NewOpsClient(k.Creds, instance.Name)

	keyvault, err := k.KeyvaultClient.GetVault(ctx, instance.Spec.ResourceGroup, instance.Spec.KeyVault)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		return false, err
	}

	vaultBaseURL := *keyvault.Properties.VaultURI

	// exit successfully if key already exists, @todo, determine how to roll these in a way users would expect
	if _, err := kvopsclient.GetKey(ctx, vaultBaseURL, instance.GetName(), ""); err == nil {
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return true, nil
	}

	katts := kvops.KeyAttributes{
		Enabled: to.BoolPtr(true),
	}

	// default to all operations if none specified
	ops := instance.Spec.Operations
	if len(ops) == 0 {
		ops = kvops.PossibleJSONWebKeyOperationValues()
	}

	params := kvops.KeyCreateParameters{
		Kty:           instance.Spec.Type,
		KeySize:       &instance.Spec.KeySize,
		KeyOps:        &instance.Spec.Operations,
		KeyAttributes: &katts,
	}

	switch instance.Spec.Type {
	case kvops.RSA, kvops.RSAHSM:
		if instance.Spec.KeySize == 0 {
			instance.Status.Message = "no keysize provided for rsa key"
			instance.Status.FailedProvisioning = true
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			return true, nil
		}
		params.KeySize = &instance.Spec.KeySize
	case kvops.EC, kvops.ECHSM:
		if instance.Spec.Curve == "" {
			instance.Status.Message = "no curve provided for ec key"
			instance.Status.FailedProvisioning = true
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			return true, nil
		}
		params.Curve = instance.Spec.Curve
	}

	req, err := kvopsclient.CreateKey(ctx, vaultBaseURL, instance.Name, params)
	if err != nil {
		instance.Status.Message = err.Error()

		// this generally means the operator doesn't have access to the keyvault
		// this can be resolved elsewhere so we should keep trying
		// See https://github.com/Azure/azure-sdk-for-go/issues/10975 for more details
		if req.Response.Response != nil && req.Response.StatusCode == http.StatusForbidden {
			return false, nil
		}

		return false, err
	}

	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false

	return true, nil
}

// Delete ensures the requested resource is gone from Azure
func (k *KeyvaultKeyClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	keyv, err := k.KeyvaultClient.GetVault(ctx, instance.Spec.ResourceGroup, instance.Spec.KeyVault)
	if err != nil {

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		return true, err
	}

	vaultBaseURL := *keyv.Properties.VaultURI
	kvopsclient := NewOpsClient(k.Creds, instance.Spec.KeyVault)

	req, err := kvopsclient.DeleteKey(ctx, vaultBaseURL, instance.Name)
	if err != nil {

		if req.Response.StatusCode == http.StatusNotFound {
			return false, nil
		}

		return true, err
	}

	// @Todo figure out when to purge or if we should ever
	// _, err = kvopsclient.PurgeDeletedKey(ctx, vaultBaseURL, instance.Name)
	// if err != nil {
	// 	return true, err
	// }

	return false, nil
}

// GetParents returns the kube resources most likely to be parents to this resource
func (k *KeyvaultKeyClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := k.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.KeyVault,
			},
			Target: &azurev1alpha1.KeyVault{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *KeyvaultKeyClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (k *KeyvaultKeyClient) convert(obj runtime.Object) (*v1alpha1.KeyVaultKey, error) {
	local, ok := obj.(*v1alpha1.KeyVaultKey)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
