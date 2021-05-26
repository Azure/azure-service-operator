/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// Validator is similar to controller-runtime/pkg/webhook/admission Validator. Implementing this interface
// allows you to hook into the code generated validations and add custom handcrafted validations.
type Validator interface {
	// CreateValidations returns validation functions that should be run on create.
	CreateValidations() []func() error
	// UpdateValidations returns validation functions that should be run on update.
	UpdateValidations() []func(old runtime.Object) error
	// DeleteValidations returns validation functions taht should be run on delete.
	DeleteValidations() []func() error
}

// Defaulter is similar to controller-runtime/pkg/webhook/admission Defaulter. Implementing this interface
// allows you to hook into the code generated defaults and add custom handcrafted defaults.
type Defaulter interface {
	// CustomDefault performs custom defaults that are run in addition to the code generated defaults.
	CustomDefault()
}
