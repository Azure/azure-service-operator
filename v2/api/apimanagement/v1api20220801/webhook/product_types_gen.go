// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220801 "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Product struct {
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-product,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=products,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.products.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Product{}

// Default applies defaults to the Product resource
func (product *Product) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220801.Product)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/Product, but got %T", obj)
	}
	err := product.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = product
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (product *Product) defaultAzureName(ctx context.Context, obj *v20220801.Product) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Product resource
func (product *Product) defaultImpl(ctx context.Context, obj *v20220801.Product) error {
	err := product.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-product,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=products,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.products.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Product{}

// ValidateCreate validates the creation of the resource
func (product *Product) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220801.Product)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/Product, but got %T", resource)
	}
	validations := product.createValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.Product]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (product *Product) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220801.Product)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/Product, but got %T", resource)
	}
	validations := product.deleteValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.Product]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (product *Product) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220801.Product)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/Product, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220801.Product)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/Product, but got %T", oldResource)
	}
	validations := product.updateValidations()
	var temp any = product
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.Product]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (product *Product) createValidations() []func(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error){product.validateResourceReferences, product.validateOwnerReference, product.validateSecretDestinations, product.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (product *Product) deleteValidations() []func(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (product *Product) updateValidations() []func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
			return product.validateResourceReferences(ctx, newObj)
		},
		product.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
			return product.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
			return product.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
			return product.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (product *Product) validateConfigMapDestinations(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (product *Product) validateOwnerReference(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (product *Product) validateResourceReferences(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (product *Product) validateSecretDestinations(ctx context.Context, obj *v20220801.Product) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (product *Product) validateWriteOnceProperties(ctx context.Context, oldObj *v20220801.Product, newObj *v20220801.Product) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
