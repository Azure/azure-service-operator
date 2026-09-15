/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"fmt"
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

var _ extensions.PreReconciliationChecker = &VaultKeyExtension{}

// PreReconcileCheck implements extensions.PreReconciliationChecker. It gives VaultKey the same
// goal-seeking behaviour Vault's createMode has, and gates adoption of pre-existing keys:
//
//   - If a live key already holds the requested name, it is adopted - but only when the
//     generation-time properties in the spec (kty, keySize, curveName) match the real key.
//     ARM's only write operation is create-if-not-exist, so proceeding with a mismatched spec
//     would report Ready for a key that differs from what the spec describes.
//   - If no live key exists but a soft-deleted one holds the name, operatorSpec.createMode
//     decides: recover it ('recover'/'createOrRecover'), purge it ('purgeThenCreate'), or let
//     ARM surface the conflict ('default').
func (ex *VaultKeyExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	key, ok := obj.(*keys.VaultKey)
	if !ok {
		return extensions.PreReconcileCheckResult{}, eris.Errorf(
			"cannot run VaultKeyExtension.PreReconcileCheck() with unexpected resource type %T",
			obj,
		)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = key

	name := key.AzureName()

	keyClient, err := newKeyClient(ctx, key, resourceResolver, armClient)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			// The owning vault doesn't exist in Azure yet; nothing to check until it does
			return extensions.BlockReconcile(
				fmt.Sprintf("vault owning key %s not found", name),
			), nil
		}

		return extensions.PreReconcileCheckResult{}, err
	}

	liveKey, err := keyClient.GetKey(ctx, name, "" /* latest version */, nil)
	if err == nil {
		// A live key already holds this name; adopt it if the generation-time properties match
		if mismatch := intrinsicMismatch(key, liveKey.Key); mismatch != "" {
			return extensions.BlockReconcile(fmt.Sprintf(
				"cannot adopt existing key %s: %s; the generation-time properties of a key cannot "+
					"be changed, so align the spec with the existing key or remove the key in Azure",
				name,
				mismatch,
			)), nil
		}

		return next(ctx, obj, resourceResolver, armClient, log)
	}
	if !genericarmclient.IsNotFoundError(err) {
		return extensions.PreReconcileCheckResult{}, eris.Wrapf(err, "failed to check for existing key %s", name)
	}

	mode := vaultKeyCreateMode(key)
	if mode == CreateMode_Default {
		// Plain create; if a soft-deleted key blocks the name, ARM will say so
		return next(ctx, obj, resourceResolver, armClient, log)
	}

	deleted, err := keyClient.GetDeletedKey(ctx, name, nil)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			// No soft-deleted key is in the way
			if mode == CreateMode_Recover {
				return extensions.BlockReconcile(fmt.Sprintf(
					"createMode is recover but no soft-deleted key named %s exists", name,
				)), nil
			}

			// createOrRecover or purgeThenCreate with nothing to recover or purge: plain create
			return next(ctx, obj, resourceResolver, armClient, log)
		}

		return extensions.PreReconcileCheckResult{}, eris.Wrapf(err, "failed to check for soft-deleted key %s", name)
	}

	switch mode {
	case CreateMode_Recover, CreateMode_CreateOrRecover:
		// Refuse to recover a key the spec doesn't describe: recovery would resurrect key
		// material with different generation-time properties than requested
		if mismatch := intrinsicMismatch(key, deleted.Key); mismatch != "" {
			return extensions.BlockReconcile(fmt.Sprintf(
				"cannot recover soft-deleted key %s: %s", name, mismatch,
			)), nil
		}

		_, err = keyClient.RecoverDeletedKey(ctx, name, nil)
		if err != nil {
			return extensions.PreReconcileCheckResult{}, eris.Wrapf(err, "failed to recover soft-deleted key %s", name)
		}

		log.Info(
			"Recovered soft-deleted key",
			"key", name,
			"createMode", mode,
		)

		// Recovery is asynchronous; once the recovered key is visible, the next reconcile
		// adopts it via the live-key path above
		return extensions.BlockReconcile(fmt.Sprintf(
			"recovery of soft-deleted key %s initiated; waiting for it to complete", name,
		)), nil
	case CreateMode_PurgeThenCreate:
		_, err = keyClient.PurgeDeletedKey(ctx, name, nil)
		if err != nil {
			return extensions.PreReconcileCheckResult{}, eris.Wrapf(err, "failed to purge soft-deleted key %s", name)
		}

		log.Info(
			"Purged soft-deleted key",
			"key", name,
		)

		return next(ctx, obj, resourceResolver, armClient, log)
	default:
		// The webhook enum validation should make this unreachable
		return extensions.PreReconcileCheckResult{}, eris.Errorf(
			"unexpected operatorSpec.createMode %q for VaultKey %s",
			mode,
			key.Name,
		)
	}
}

// intrinsicMismatch compares the generation-time properties in the spec (kty, keySize, curveName)
// against the actual JSON web key in Azure. It returns a human-readable description of the first
// mismatch, or "" when every property set in the spec matches. Properties the spec leaves unset
// match anything.
func intrinsicMismatch(key *keys.VaultKey, actual *azkeys.JSONWebKey) string {
	props := key.Spec.Properties
	if props == nil || actual == nil {
		return ""
	}

	if props.Kty != nil {
		actualKty := ""
		if actual.Kty != nil {
			actualKty = string(*actual.Kty)
		}

		if actualKty != *props.Kty {
			return fmt.Sprintf("spec.properties.kty is %s but the key in Azure has kty %q", *props.Kty, actualKty)
		}
	}

	// The RSA modulus always has its top bit set, so its length gives the exact key size.
	// Non-RSA keys have no modulus, in which case a spec'd keySize can't be verified (ARM
	// ignores keySize for such keys anyway) and is left to the service to judge.
	if props.KeySize != nil && len(actual.N) > 0 {
		actualSize := len(actual.N) * 8
		if actualSize != *props.KeySize {
			return fmt.Sprintf("spec.properties.keySize is %d but the key in Azure has size %d", *props.KeySize, actualSize)
		}
	}

	if props.CurveName != nil {
		actualCurve := ""
		if actual.Crv != nil {
			actualCurve = string(*actual.Crv)
		}

		if actualCurve != *props.CurveName {
			return fmt.Sprintf("spec.properties.curveName is %s but the key in Azure has curve %q", *props.CurveName, actualCurve)
		}
	}

	return ""
}

// vaultKeyCreateMode returns the effective operatorSpec.createMode, defaulting to default - a
// plain create that surfaces any name conflict with a soft-deleted key as an error from Azure.
func vaultKeyCreateMode(key *keys.VaultKey) string {
	if key.Spec.OperatorSpec != nil && key.Spec.OperatorSpec.CreateMode != nil {
		return *key.Spec.OperatorSpec.CreateMode
	}

	return CreateMode_Default
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
