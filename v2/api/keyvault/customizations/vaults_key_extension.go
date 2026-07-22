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
// managed-delete behavior (disable the key and block deletion).
const detachAckAnnotation = "vaultskey.aso.io/acknowledge-key-retained-enabled"

// vaultsKeyARMAPIVersion is the (fixed) ARM API version for VaultsKey. VaultsKey currently has only a
// single API version, and this mirrors the same hardcoded literal used in the generated
// VaultsKey.GetAPIVersion() method (v2/api/keyvault/v20230701/vaults_key_types_gen.go), which is only
// available on the versioned type - not on the storage/hub type this extension operates on.
const vaultsKeyARMAPIVersion = "2023-07-01"

// Reasons used for the Ready condition set (indirectly, via ReadyConditionImpactingError) while
// blocking deletion of a VaultsKey.
var (
	// ReasonKeyDeletionBlocked indicates the key was successfully disabled in Azure, but deletion
	// of the Kubernetes resource remains blocked until the operator is explicitly told to proceed.
	ReasonKeyDeletionBlocked = conditions.MakeReason("KeyDeletionBlocked", retry.Slow)

	// ReasonKeyDeletionBlockedDisableFailed indicates an attempt to disable the key in Azure failed,
	// so the key may still be ENABLED. Deletion remains blocked.
	ReasonKeyDeletionBlockedDisableFailed = conditions.MakeReason("KeyDeletionBlockedDisableFailed", retry.Slow)

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
		"(which disables the key on delete). Falling back to managed delete (disable + block) until then.")
}

var _ extensions.Deleter = &VaultsKeyExtension{}

// Delete implements extensions.Deleter.
//
// There is no ARM management-plane DELETE operation for Key Vault keys (see
// VaultsKey.GetSupportedOperations, which only returns Get and Put); deleting/disabling keys is a
// data-plane concept. Because of this, and because a key holds live cryptographic material, this
// extension deliberately never allows the operator to make the underlying Azure key inaccessible on
// its own initiative: the best it will do is disable the key (attributes.enabled=false) via an ARM PUT,
// and it always blocks removal of the Kubernetes finalizer afterward. The only ways to actually get rid
// of the Kubernetes object are:
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
		// Never created in Azure to begin with, nothing to disable/block on.
		log.V(1).Info("VaultsKey has no ARM resource ID; nothing to delete")
		return ctrl.Result{}, nil
	}

	// Step 1: GET the current state of the key from ARM.
	//
	// Known, documented, non-blocking limitation: there is a narrow race window between this GET and
	// the PUT below. If the key is purged (data-plane) in that split-second window, this PUT will
	// recreate the key (in a disabled state) rather than leaving it purged. There is no atomic
	// conditional PUT available in this ARM API to close this race entirely. This is considered
	// acceptable: the window is narrow, and the outcome (a disabled key) is still safe.
	raw := make(map[string]any)
	_, err := armClient.GetByID(ctx, resourceID, vaultsKeyARMAPIVersion, &raw)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			// The key is already gone from Azure (e.g. purged out of band). Nothing further to do;
			// allow the Kubernetes delete of this object to proceed.
			log.V(1).Info("VaultsKey no longer exists in Azure, allowing deletion to proceed")
			return ctrl.Result{}, nil
		}

		// We couldn't determine whether the key still exists. Do not proceed with a PUT, and do not
		// allow deletion to proceed either - retry.
		return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
			eris.Wrapf(err, "failed to get current state of key %q before delete", typedObj.Name),
			conditions.ConditionSeverityWarning,
			ReasonKeyDeletionStatusUnknown,
		)
	}

	// Step 2: best-effort disable (not delete) of the key via ARM PUT.
	disableErr := disableKey(ctx, armClient, resourceID, vaultsKeyARMAPIVersion, raw)

	// Step 3/4: regardless of whether disabling succeeded, deletion remains blocked. We distinguish the
	// two outcomes via distinct Reasons/Severities on the Ready condition (set indirectly via the
	// returned ReadyConditionImpactingError - see generic_reconciler.go, which unconditionally overwrites
	// any condition we might set directly here on a non-error return, so an error return is required for
	// a custom Reason/Message to stick).
	if disableErr != nil {
		msg := eris.Wrapf(
			disableErr,
			"failed to disable key %q in Azure; the key MAY STILL BE ENABLED. Deletion of this resource is "+
				"blocked until the key can be confirmed disabled. Retry, or check RBAC/permissions on the key vault",
			typedObj.Name,
		)
		return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
			msg,
			conditions.ConditionSeverityWarning,
			ReasonKeyDeletionBlockedDisableFailed,
		)
	}

	msg := eris.Errorf(
		"key %q has been disabled in Azure but not deleted; VaultsKey does not support deleting keys via "+
			"ARM. Deletion of this resource remains blocked to prevent silently leaving a live/enabled key "+
			"behind. To proceed, either use reconcile-policy=detach-on-delete with the %q=\"true\" "+
			"annotation, or purge the key out of band via the Key Vault data plane",
		typedObj.Name,
		detachAckAnnotation,
	)
	return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
		msg,
		conditions.ConditionSeverityInfo,
		ReasonKeyDeletionBlocked,
	)
}

// disableKey issues a best-effort ARM PUT to set properties.attributes.enabled=false on the raw ARM
// resource representation obtained from a prior GET, without otherwise modifying the resource.
func disableKey(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	resourceID string,
	apiVersion string,
	raw map[string]any,
) error {
	properties, ok := raw["properties"].(map[string]any)
	if !ok {
		properties = map[string]any{}
		raw["properties"] = properties
	}

	attributes, ok := properties["attributes"].(map[string]any)
	if !ok {
		attributes = map[string]any{}
		properties["attributes"] = attributes
	}

	attributes["enabled"] = false

	poller, err := armClient.BeginCreateOrUpdateByID(ctx, resourceID, apiVersion, raw)
	if err != nil {
		return eris.Wrapf(err, "failed to begin disabling key")
	}

	_, err = poller.Poller.PollUntilDone(ctx, nil)
	if err != nil {
		return eris.Wrapf(err, "failed while polling disable operation to completion")
	}

	return nil
}
