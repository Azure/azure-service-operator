// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230501 "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type SecurityPolicy struct {
}

// +kubebuilder:webhook:path=/mutate-cdn-azure-com-v1api20230501-securitypolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=securitypolicies,verbs=create;update,versions=v1api20230501,name=default.v1api20230501.securitypolicies.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &SecurityPolicy{}

// Default applies defaults to the SecurityPolicy resource
func (policy *SecurityPolicy) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230501.SecurityPolicy)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/SecurityPolicy, but got %T", obj)
	}
	err := policy.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (policy *SecurityPolicy) defaultAzureName(ctx context.Context, obj *v20230501.SecurityPolicy) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the SecurityPolicy resource
func (policy *SecurityPolicy) defaultImpl(ctx context.Context, obj *v20230501.SecurityPolicy) error {
	err := policy.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-cdn-azure-com-v1api20230501-securitypolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=securitypolicies,verbs=create;update,versions=v1api20230501,name=validate.v1api20230501.securitypolicies.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &SecurityPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *SecurityPolicy) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.SecurityPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/SecurityPolicy, but got %T", resource)
	}
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.SecurityPolicy]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *SecurityPolicy) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.SecurityPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/SecurityPolicy, but got %T", resource)
	}
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.SecurityPolicy]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (policy *SecurityPolicy) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230501.SecurityPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/SecurityPolicy, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230501.SecurityPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/SecurityPolicy, but got %T", oldResource)
	}
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.SecurityPolicy]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (policy *SecurityPolicy) createValidations() []func(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference, policy.validateSecretDestinations, policy.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (policy *SecurityPolicy) deleteValidations() []func(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *SecurityPolicy) updateValidations() []func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
			return policy.validateResourceReferences(ctx, newObj)
		},
		policy.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
			return policy.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
			return policy.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
			return policy.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (policy *SecurityPolicy) validateConfigMapDestinations(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (policy *SecurityPolicy) validateOwnerReference(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (policy *SecurityPolicy) validateResourceReferences(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (policy *SecurityPolicy) validateSecretDestinations(ctx context.Context, obj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *SecurityPolicy) validateWriteOnceProperties(ctx context.Context, oldObj *v20230501.SecurityPolicy, newObj *v20230501.SecurityPolicy) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
