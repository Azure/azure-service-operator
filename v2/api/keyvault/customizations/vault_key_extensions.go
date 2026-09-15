/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701/storage"
	keys "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const (
	DeleteMode_Detach  = "detach"
	DeleteMode_Delete  = "delete"
	DeleteMode_Disable = "disable"
)

var _ extensions.Deleter = &VaultKeyExtension{}

// Delete implements extensions.Deleter. VaultKey deliberately exposes no ARM DELETE operation
// (guarded by the supported-operations allowlist test), so the reconciler routes deletion through
// DeleteNotPossibleInAzure, which invokes this extension instead of issuing an ARM delete.
// The operatorSpec.deleteMode value selects what happens to the key in Azure:
//
//   - detach (default): leave the key untouched; only the Kubernetes resource goes away.
//   - disable: set the key's enabled attribute to false via the data plane, then let the
//     Kubernetes resource go. The key remains in the vault.
//   - delete: soft-delete the key via the data plane; it remains recoverable for the vault's
//     soft-delete retention period.
//
// next is never called on a successful path: the default behaviour for a resource without ARM
// delete is an error instructing the user to detach, and this extension replaces that outcome.
func (ex *VaultKeyExtension) Delete(
	ctx context.Context,
	log logr.Logger,
	reslv *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	obj genruntime.ARMMetaObject,
	_ extensions.DeleteFunc,
) (extensions.DeleteResult, error) {
	key, ok := obj.(*keys.VaultKey)
	if !ok {
		return extensions.DeleteResult{}, eris.Errorf(
			"cannot run VaultKeyExtension.Delete() with unexpected resource type %T",
			obj,
		)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = key

	mode := vaultKeyDeleteMode(key)
	switch mode {
	case DeleteMode_Detach:
		log.Info(
			"deleteMode is detach; leaving key in place in Azure",
			"key", key.AzureName(),
		)
		return extensions.DeleteCompleted(), nil
	case DeleteMode_Disable, DeleteMode_Delete:
		// Handled below, after we have a data-plane client
	default:
		// The webhook enum validation should make this unreachable
		return extensions.DeleteResult{}, eris.Errorf(
			"unexpected operatorSpec.deleteMode %q for VaultKey %s",
			mode,
			key.Name,
		)
	}

	keyClient, err := newKeyClient(ctx, key, reslv, armClient)
	if err != nil {
		return extensions.DeleteResult{}, err
	}

	name := key.AzureName()
	switch mode {
	case DeleteMode_Disable:
		_, err = keyClient.UpdateKey(
			ctx,
			name,
			"", // latest version
			azkeys.UpdateKeyParameters{
				KeyAttributes: &azkeys.KeyAttributes{
					Enabled: to.Ptr(false),
				},
			},
			nil,
		)
		if err != nil && !genericarmclient.IsNotFoundError(err) {
			return extensions.DeleteResult{}, eris.Wrapf(err, "failed to disable key %s", name)
		}

		log.Info(
			"deleteMode is disable; disabled key and left it in Azure",
			"key", name,
		)
	case DeleteMode_Delete:
		_, err = keyClient.DeleteKey(ctx, name, nil)
		if err != nil && !genericarmclient.IsNotFoundError(err) {
			return extensions.DeleteResult{}, eris.Wrapf(err, "failed to delete key %s", name)
		}

		log.Info(
			"deleteMode is delete; soft-deleted key in Azure",
			"key", name,
		)
	}

	return extensions.DeleteCompleted(), nil
}

// vaultKeyDeleteMode returns the effective operatorSpec.deleteMode, defaulting to detach - the
// safe option that never touches the key in Azure.
func vaultKeyDeleteMode(key *keys.VaultKey) string {
	if key.Spec.OperatorSpec != nil && key.Spec.OperatorSpec.DeleteMode != nil {
		return *key.Spec.OperatorSpec.DeleteMode
	}

	return DeleteMode_Detach
}

// newKeyClient creates a Key Vault data-plane keys client for the vault holding this key,
// sharing the ARM client's credentials and HTTP pipeline options (so recorded tests capture
// data-plane traffic the same way they capture ARM traffic).
func newKeyClient(
	ctx context.Context,
	key *keys.VaultKey,
	reslv *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
) (*azkeys.Client, error) {
	vaultURL, err := getVaultURL(ctx, key, reslv, armClient)
	if err != nil {
		return nil, err
	}

	options := &azkeys.ClientOptions{}
	if armOptions := armClient.ClientOptions(); armOptions != nil {
		options.ClientOptions = armOptions.ClientOptions
	}

	keyClient, err := azkeys.NewClient(vaultURL, armClient.Creds(), options)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create Key Vault data-plane client for %s", vaultURL)
	}

	return keyClient, nil
}

// getVaultURL determines the data-plane URL of the vault holding this key. The key's own status
// carries the full key URI once the key has been created; before that (or if the status was never
// populated) we fall back to reading vaultUri from the owning vault via ARM, which works whether
// the owner is an ASO-managed Vault or a plain ARM ID reference, and in every cloud (no hardcoded
// DNS suffix).
func getVaultURL(
	ctx context.Context,
	key *keys.VaultKey,
	reslv *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
) (string, error) {
	if key.Status.KeyUri != nil && *key.Status.KeyUri != "" {
		return vaultURLFromKeyURI(*key.Status.KeyUri)
	}

	vaultID, err := getVaultID(ctx, key, reslv)
	if err != nil {
		return "", err
	}

	vc, err := armkeyvault.NewVaultsClient(vaultID.SubscriptionID, armClient.Creds(), armClient.ClientOptions())
	if err != nil {
		return "", eris.Wrap(err, "failed to create new VaultsClient")
	}

	resp, err := vc.Get(ctx, vaultID.ResourceGroupName, vaultID.Name, nil)
	if err != nil {
		return "", eris.Wrapf(err, "failed to read vault %s to determine its data-plane URI", vaultID.Name)
	}

	if resp.Properties == nil || resp.Properties.VaultURI == nil || *resp.Properties.VaultURI == "" {
		return "", eris.Errorf("vault %s did not report a vaultUri", vaultID.Name)
	}

	return *resp.Properties.VaultURI, nil
}

// vaultURLFromKeyURI reduces a key URI such as https://myvault.vault.azure.net/keys/mykey to the
// vault's base URL, https://myvault.vault.azure.net.
func vaultURLFromKeyURI(keyURI string) (string, error) {
	u, err := url.Parse(keyURI)
	if err != nil {
		return "", eris.Wrapf(err, "failed to parse key URI %q", keyURI)
	}

	if u.Scheme == "" || u.Host == "" {
		return "", eris.Errorf("key URI %q does not contain a scheme and host", keyURI)
	}

	return u.Scheme + "://" + u.Host, nil
}

// getVaultID resolves the ARM ID of the vault owning this key.
func getVaultID(
	ctx context.Context,
	key *keys.VaultKey,
	reslv *resolver.Resolver,
) (*arm.ResourceID, error) {
	owner, err := reslv.ResolveOwner(ctx, key)
	if err != nil {
		return nil, eris.Wrapf(err, "unable to resolve owner of VaultKey %s", key.Name)
	}

	switch owner.Result {
	case resolver.OwnerFoundKubernetes:
		vault, ok := owner.Owner.(*keyvault.Vault)
		if !ok {
			return nil, eris.Errorf("expected owner of VaultKey %s to be a Vault", key.Name)
		}

		// Type assert that the Vault is the hub type. This will fail to compile if
		// the hub type has been changed but this extension has not been updated to match
		var _ conversion.Hub = vault

		id, err := genruntime.GetAndParseResourceID(vault)
		if err != nil {
			return nil, eris.Wrap(err, "failed to get and parse resource ID from VaultKey owner")
		}

		return id, nil
	case resolver.OwnerFoundARM:
		id, err := arm.ParseResourceID(owner.ARMID)
		if err != nil {
			return nil, eris.Wrap(err, "failed to parse resource ID from VaultKey owner")
		}

		return id, nil
	default:
		return nil, eris.Errorf("unexpected owner type of VaultKey, type: %s", owner.Result)
	}
}
