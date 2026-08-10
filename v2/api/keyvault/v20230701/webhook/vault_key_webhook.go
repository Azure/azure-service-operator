/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package webhook

import (
	"context"
	"reflect"

	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	v20230701 "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var _ genruntime.Validator[*v20230701.VaultKey] = &VaultKey{}

// CreateValidations returns validation functions for VaultKey creation.
func (vaultKey *VaultKey) CreateValidations() []func(ctx context.Context, obj *v20230701.VaultKey) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230701.VaultKey) (admission.Warnings, error){
		vaultKey.validateNotExportable,
	}
}

// UpdateValidations returns validation functions for VaultKey updates.
func (vaultKey *VaultKey) UpdateValidations() []func(ctx context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error) {
			return vaultKey.validateNotExportable(ctx, newObj)
		},
		vaultKey.validateIntrinsicallyImmutable,
		vaultKey.validateNotSilentlyIgnored,
	}
}

// DeleteValidations returns validation functions for VaultKey deletion.
func (vaultKey *VaultKey) DeleteValidations() []func(ctx context.Context, obj *v20230701.VaultKey) (admission.Warnings, error) {
	return nil
}

// validateNotExportable rejects any VaultKey whose properties.attributes.exportable is true, on both
// create and update. Exportable keys allow the private key material to leave Key Vault, which this
// operator does not support enabling via a declarative CRD (there is no way to safely audit or gate that
// via Kubernetes RBAC alone).
func (vaultKey *VaultKey) validateNotExportable(_ context.Context, obj *v20230701.VaultKey) (admission.Warnings, error) {
	if obj.Spec.Properties == nil || obj.Spec.Properties.Attributes == nil {
		return nil, nil
	}

	if obj.Spec.Properties.Attributes.Exportable != nil && *obj.Spec.Properties.Attributes.Exportable {
		return nil, eris.Errorf(
			"spec.properties.attributes.exportable=true is not allowed for %s : %s; exportable keys are not supported",
			obj.GetObjectKind().GroupVersionKind(),
			obj.GetName(),
		)
	}

	return nil, nil
}

// validateIntrinsicallyImmutable rejects changes to the key properties that are immutable in Azure
// under ANY mechanism, ARM or data-plane: a key's type, size, and curve are fixed at generation time.
// Changing them describes a different key, not a modification of this one, so such an edit could never
// be applied and is rejected outright. These fields would remain immutable even if this resource ever
// gained data-plane update support (unlike the fields covered by validateNotSilentlyIgnored).
//
// Note: spec.azureName and spec.owner are similarly write-once and are already validated by
// genruntime.ValidateWriteOnceProperties, which every generated webhook runs - we don't duplicate
// that here.
func (vaultKey *VaultKey) validateIntrinsicallyImmutable(_ context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error) {
	if !genruntime.IsResourceCreatedSuccessfully(oldObj) {
		// No ARM resource ID stamped yet - no immutability concerns apply. (The ID is stamped when
		// the operator first claims the resource, before the initial PUT to Azure, so this gate
		// closes at claim time, not on confirmed successful creation.)
		return nil, nil
	}

	type intrinsic struct {
		kty       *v20230701.KeyProperties_Kty
		keySize   *int
		curveName *v20230701.KeyProperties_CurveName
	}
	extract := func(obj *v20230701.VaultKey) intrinsic {
		props := obj.Spec.Properties
		if props == nil {
			return intrinsic{}
		}
		return intrinsic{kty: props.Kty, keySize: props.KeySize, curveName: props.CurveName}
	}

	if !reflect.DeepEqual(extract(oldObj), extract(newObj)) {
		return nil, eris.Errorf(
			"spec.properties.kty, keySize and curveName are immutable for %s : %s - they are fixed at "+
				"key generation time, so changing them describes a different key rather than a modification "+
				"of this one; delete and recreate the resource to change them (if the resource never "+
				"successfully created in Azure, deletion is not blocked - delete and re-apply the corrected "+
				"spec)",
			newObj.GetObjectKind().GroupVersionKind(),
			newObj.GetName(),
		)
	}

	return nil, nil
}

// validateNotSilentlyIgnored rejects changes to the remaining spec fields (keyOps, attributes,
// release_policy, rotationPolicy and tags). These are mutable on the key itself via the Key Vault
// data plane, but this resource operates purely on the ARM control plane, whose only write operation
// (CreateIfNotExist) is a no-op against an existing key. Letting such an edit through would leave the
// resource reporting Ready with a spec that silently diverges from the real state in Azure - a worse
// failure mode than a clear rejection. If data-plane update support is ever added, this validator -
// and only this validator - can be relaxed.
//
// The comparison covers the whole spec.properties subtree except the fields owned by
// validateIntrinsicallyImmutable, so a field added to KeyProperties by a future regeneration is
// automatically covered until it is deliberately reviewed.
func (vaultKey *VaultKey) validateNotSilentlyIgnored(_ context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error) {
	if !genruntime.IsResourceCreatedSuccessfully(oldObj) {
		// No ARM resource ID stamped yet - no immutability concerns apply. (The ID is stamped when
		// the operator first claims the resource, before the initial PUT to Azure, so this gate
		// closes at claim time, not on confirmed successful creation.)
		return nil, nil
	}

	// Mask out the intrinsically-immutable fields; those are validated, with a more precise error,
	// by validateIntrinsicallyImmutable.
	stripIntrinsic := func(props *v20230701.KeyProperties) *v20230701.KeyProperties {
		if props == nil {
			return nil
		}
		stripped := *props
		stripped.Kty = nil
		stripped.KeySize = nil
		stripped.CurveName = nil
		return &stripped
	}

	if !reflect.DeepEqual(stripIntrinsic(oldObj.Spec.Properties), stripIntrinsic(newObj.Spec.Properties)) {
		return nil, eris.Errorf(
			"spec.properties is immutable after creation for %s : %s (keyOps, attributes, release_policy "+
				"and rotationPolicy cannot be changed through this resource: the ARM API's only write "+
				"operation is create-if-not-exist, so the edit would be accepted but silently ignored by "+
				"Azure); delete and recreate the resource to change them (if the resource never "+
				"successfully created in Azure, deletion is not blocked - delete and re-apply the corrected "+
				"spec)",
			newObj.GetObjectKind().GroupVersionKind(),
			newObj.GetName(),
		)
	}

	if !reflect.DeepEqual(oldObj.Spec.Tags, newObj.Spec.Tags) {
		return nil, eris.Errorf(
			"spec.tags is immutable after creation for %s : %s (the ARM API's only write operation is "+
				"create-if-not-exist, so the edit would be accepted but silently ignored by Azure); delete "+
				"and recreate the resource to change it",
			newObj.GetObjectKind().GroupVersionKind(),
			newObj.GetName(),
		)
	}

	return nil, nil
}
