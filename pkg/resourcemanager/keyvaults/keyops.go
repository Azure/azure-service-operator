package keyvaults

import (
	"context"
	"fmt"

	kvops "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type KeyvaultKeyClient struct {
	KeyvaultClient *azureKeyVaultManager
}

func (k *KeyvaultKeyClient) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	instance.Status.Provisioning = true

	// Check if this KeyVault already exists and its state if it does.

	kvopsclient := NewOpsClient(instance.Name)

	keyvault, err := k.KeyvaultClient.GetVault(ctx, instance.Spec.ResourceGroup, instance.Spec.KeyVault)
	if err == nil {
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false

		vaultBaseURL := *keyvault.Properties.VaultURI
		var ksize int32 = instance.Spec.KeySize
		kops := kvops.PossibleJSONWebKeyOperationValues()
		katts := kvops.KeyAttributes{
			Enabled: to.BoolPtr(true),
		}
		params := kvops.KeyCreateParameters{
			Kty:           kvops.RSA,
			KeySize:       &ksize,
			KeyOps:        &kops,
			KeyAttributes: &katts,
		}
		_, err := kvopsclient.CreateKey(ctx, vaultBaseURL, instance.Name, params)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return true, nil
}

func (k *KeyvaultKeyClient) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	keyv, err := k.KeyvaultClient.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		vaultBaseURL := *keyv.Properties.VaultURI
		kvopsclient := NewOpsClient(instance.Name)

		_, err := kvopsclient.DeleteKey(ctx, vaultBaseURL, instance.Name)
		if err != nil {
			return true, err
		}
	}

	return false, nil
}

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

func (k *KeyvaultKeyClient) convert(obj runtime.Object) (*v1alpha1.KeyVaultKey, error) {
	local, ok := obj.(*v1alpha1.KeyVaultKey)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
