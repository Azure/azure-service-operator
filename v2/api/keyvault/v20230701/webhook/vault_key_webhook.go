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
		vaultKey.validateImmutable,
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

// validateImmutable rejects any change to VaultKey's key-material-affecting spec fields after creation.
// A VaultKey's cryptographic properties are immutable by design: changing them (e.g. kty, keySize) is
// not really a "modification" of the same key in Azure semantics, and silently attempting to PUT such a
// change could have surprising/unsafe results. Users must delete and recreate the resource to change any
// of these fields.
//
// Note: spec.azureName and spec.owner are already validated as immutable (once the resource is created)
// by genruntime.ValidateWriteOnceProperties, which every generated webhook already runs - we don't
// duplicate that here.
func (vaultKey *VaultKey) validateImmutable(_ context.Context, oldObj *v20230701.VaultKey, newObj *v20230701.VaultKey) (admission.Warnings, error) {
	if !genruntime.IsResourceCreatedSuccessfully(oldObj) {
		// Not created yet in Azure - no immutability concerns apply.
		return nil, nil
	}

	if !reflect.DeepEqual(oldObj.Spec.Properties, newObj.Spec.Properties) {
		return nil, eris.Errorf(
			"spec.properties is immutable after creation for %s : %s (this includes kty, keySize, curveName, "+
				"keyOps, attributes, release_policy, and rotationPolicy); delete and recreate the resource to "+
				"change it",
			newObj.GetObjectKind().GroupVersionKind(),
			newObj.GetName(),
		)
	}

	if !reflect.DeepEqual(oldObj.Spec.Tags, newObj.Spec.Tags) {
		return nil, eris.Errorf(
			"spec.tags is immutable after creation for %s : %s; delete and recreate the resource to change it",
			newObj.GetObjectKind().GroupVersionKind(),
			newObj.GetName(),
		)
	}

	return nil, nil
}
