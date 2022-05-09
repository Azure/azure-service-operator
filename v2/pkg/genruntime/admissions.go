/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// Validator is similar to controller-runtime/pkg/webhook/admission Validator. Implementing this interface
// allows you to hook into the code generated validations and add custom handcrafted validations.
type Validator interface {
	// CreateValidations returns validation functions that should be run on create.
	CreateValidations() []func() error
	// UpdateValidations returns validation functions that should be run on update.
	UpdateValidations() []func(old runtime.Object) error
	// DeleteValidations returns validation functions that should be run on delete.
	DeleteValidations() []func() error
}

// Defaulter is similar to controller-runtime/pkg/webhook/admission Defaulter. Implementing this interface
// allows you to hook into the code generated defaults and add custom handcrafted defaults.
type Defaulter interface {
	// CustomDefault performs custom defaults that are run in addition to the code generated defaults.
	CustomDefault()
}

// ValidateWriteOnceProperties function validates the update on WriteOnce properties.
func ValidateWriteOnceProperties(oldObj MetaObject, newObj MetaObject) error {
	var errs []error

	if !IsResourceCreatedSuccessfully(newObj) {
		return nil
	}

	if oldObj.AzureName() != newObj.AzureName() {
		errs = append(errs, errors.Errorf("updating 'AzureName' is not allowed for '%s : %s", oldObj.GetObjectKind().GroupVersionKind(), oldObj.GetName()))
	}

	// Allow ResourceGroup update only if resource is not created successfully
	if oldObj.Owner().Name != newObj.Owner().Name {
		errs = append(errs, errors.Errorf("updating 'Owner.Name' is not allowed for '%s : %s", oldObj.GetObjectKind().GroupVersionKind(), oldObj.GetName()))
	}

	return kerrors.NewAggregate(errs)
}

func IsResourceCreatedSuccessfully(obj MetaObject) bool {
	return GetResourceIDOrDefault(obj) != ""
}
