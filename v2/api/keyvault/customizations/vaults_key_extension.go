/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	storage "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

// detachAckAnnotation, when set to "true" on a VaultsKey object, acknowledges that the underlying Azure
// key will be left ENABLED and live in Key Vault when the object is detached (reconcile-policy=detach-on-delete).
// Without this acknowledgment, detach requests are not honored and the operator falls back to its
// managed-delete behavior, which blocks deletion of the Kubernetes resource until the key no longer
// exists in Azure (see Delete below).
const detachAckAnnotation = "vaultskey.aso.io/acknowledge-key-retained-enabled"

// vaultsKeyARMAPIVersion is the (fixed) ARM API version for VaultsKey. VaultsKey currently has only a
// single API version, and this mirrors the same hardcoded literal used in the generated
// VaultsKey.GetAPIVersion() method (v2/api/keyvault/v20230701/vaults_key_types_gen.go), which is only
// available on the versioned type - not on the storage/hub type this extension operates on.
const vaultsKeyARMAPIVersion = "2023-07-01"

// Reasons used for the Ready condition set (indirectly, via ReadyConditionImpactingError) while
// blocking deletion of a VaultsKey.
var (
	// ReasonKeyDeletionBlocked indicates the key still exists in Azure. VaultsKey has no ARM operation
	// capable of modifying or deleting an existing key, so deletion of the Kubernetes resource remains
	// blocked until the key is removed via the Key Vault data-plane, or until detach is acknowledged.
	ReasonKeyDeletionBlocked = conditions.MakeReason("KeyDeletionBlocked", retry.Slow)

	// ReasonKeyDeletionStatusUnknown indicates the operator could not determine whether the key still
	// exists in Azure (a GET failed for a reason other than NotFound). Deletion remains blocked and the
	// operation is retried.
	ReasonKeyDeletionStatusUnknown = conditions.MakeReason("KeyDeletionStatusUnknown", retry.Slow)
)

var _ extensions.DetachAcknowledgementRequirer = &VaultsKeyExtension{}

// RequireDetachAcknowledgement implements extensions.DetachAcknowledgementRequirer.
// VaultsKey retains the Azure key in an ENABLED, live state when detached (there is no ARM DELETE
// operation for keys - see Delete below), so a detach must be explicitly acknowledged per-object
// before the operator will honor it.
func (e *VaultsKeyExtension) RequireDetachAcknowledgement(obj genruntime.MetaObject) error {
	if obj.GetAnnotations()[detachAckAnnotation] == "true" {
		return nil
	}

	return errors.New("reconcile-policy=detach-on-delete is in effect (possibly via a namespace-wide policy), " +
		"but VaultsKey retains the Azure key in an ENABLED, live state when detached. To detach, set annotation " +
		detachAckAnnotation + `="true" on THIS object, or override with object-level reconcile-policy=manage ` +
		"(which blocks deletion until the key is removed via the data-plane). Falling back to managed delete " +
		"(block until key is removed) until then.")
}

var _ extensions.Deleter = &VaultsKeyExtension{}

// Delete implements extensions.Deleter.
//
// There is no ARM management-plane DELETE (or UPDATE) operation for Key Vault keys (see
// VaultsKey.GetSupportedOperations, which only returns Get and Put; the Put operation
// - Keys_CreateIfNotExist - performs no write operations when the key already exists, i.e. it is
// create-only). Modifying or disabling a key's attributes is exclusively a Key Vault data-plane
// concept, outside the ARM control plane this resource operates on. Because of this, and because a key
// holds live cryptographic material, this extension deliberately never allows the operator to make the
// underlying Azure key inaccessible on its own initiative: it always blocks removal of the Kubernetes
// finalizer while the key still exists. The only ways to actually get rid of the Kubernetes object are:
//   - the key has already been removed/purged from Key Vault out of band (data-plane purge), in which
//     case the GET below returns NotFound and we allow the Kubernetes delete to proceed, or
//   - RequireDetachAcknowledgement above is satisfied and reconcile-policy=detach-on-delete is used, in
//     which case this Delete function is never called at all (see generic_reconciler.go).
func (e *VaultsKeyExtension) Delete(
	ctx context.Context,
	log logr.Logger,
	resolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	obj genruntime.ARMMetaObject,
	next extensions.DeleteFunc,
) (ctrl.Result, error) {
	typedObj, ok := obj.(*storage.VaultsKey)
	if !ok {
		return ctrl.Result{}, eris.Errorf("cannot run on unknown resource type %T, expected *storage.VaultsKey", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = typedObj

	resourceID := genruntime.GetResourceIDOrDefault(typedObj)
	if resourceID == "" {
		// Never created in Azure to begin with, nothing to block on.
		log.V(1).Info("VaultsKey has no ARM resource ID; nothing to delete")
		return ctrl.Result{}, nil
	}

	// GET the current state of the key from ARM to determine whether it still exists.
	raw := make(map[string]any)
	_, err := armClient.GetByID(ctx, resourceID, vaultsKeyARMAPIVersion, &raw)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			// The key is already gone from Azure (e.g. purged out of band). Nothing further to do;
			// allow the Kubernetes delete of this object to proceed.
			log.V(1).Info("VaultsKey no longer exists in Azure, allowing deletion to proceed")
			return ctrl.Result{}, nil
		}

		// We couldn't determine whether the key still exists. Do not allow deletion to proceed - retry.
		return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
			eris.Wrapf(err, "failed to get current state of key %q before delete", typedObj.Name),
			conditions.ConditionSeverityWarning,
			ReasonKeyDeletionStatusUnknown,
		)
	}

	// The key still exists in Azure. There is no ARM operation available to modify or delete it (see
	// the function doc comment above), so deletion of this Kubernetes resource remains blocked. We
	// return an error to force a custom Reason/Message to stick on the Ready condition (set indirectly
	// via ReadyConditionImpactingError - see generic_reconciler.go, which unconditionally overwrites any
	// condition we might set directly here on a non-error return).
	msg := eris.Errorf(
		"Microsoft.KeyVault/vaults/keys cannot be modified or deleted through the ARM control plane (the "+
			"API supports only create and read; it has no update or delete operation). ASO will not remove "+
			"this Kubernetes resource while the key still exists in Azure. To proceed: (a) set "+
			"reconcile-policy=detach-on-delete plus annotation %s=\"true\" to detach the CR while leaving "+
			"the key intact in Azure, or (b) disable/destroy the key via the Key Vault data-plane "+
			"(CLI/Portal/az keyvault key set-attributes|delete), then re-attempt deletion",
		detachAckAnnotation,
	)
	return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
		msg,
		conditions.ConditionSeverityInfo,
		ReasonKeyDeletionBlocked,
	)
}
