// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type NetworkSecurityGroup struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-networksecuritygroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=networksecuritygroups,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.networksecuritygroups.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &NetworkSecurityGroup{}

// Default applies defaults to the NetworkSecurityGroup resource
func (group *NetworkSecurityGroup) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.NetworkSecurityGroup)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkSecurityGroup, but got %T", obj)
	}
	err := group.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = group
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (group *NetworkSecurityGroup) defaultAzureName(ctx context.Context, obj *v20240301.NetworkSecurityGroup) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the NetworkSecurityGroup resource
func (group *NetworkSecurityGroup) defaultImpl(ctx context.Context, obj *v20240301.NetworkSecurityGroup) error {
	err := group.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-networksecuritygroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=networksecuritygroups,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.networksecuritygroups.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &NetworkSecurityGroup{}

// ValidateCreate validates the creation of the resource
func (group *NetworkSecurityGroup) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.NetworkSecurityGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkSecurityGroup, but got %T", resource)
	}
	validations := group.createValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkSecurityGroup]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (group *NetworkSecurityGroup) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.NetworkSecurityGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkSecurityGroup, but got %T", resource)
	}
	validations := group.deleteValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkSecurityGroup]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (group *NetworkSecurityGroup) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.NetworkSecurityGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkSecurityGroup, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.NetworkSecurityGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkSecurityGroup, but got %T", oldResource)
	}
	validations := group.updateValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkSecurityGroup]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (group *NetworkSecurityGroup) createValidations() []func(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error){group.validateResourceReferences, group.validateOwnerReference, group.validateSecretDestinations, group.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (group *NetworkSecurityGroup) deleteValidations() []func(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (group *NetworkSecurityGroup) updateValidations() []func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
			return group.validateResourceReferences(ctx, newObj)
		},
		group.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
			return group.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
			return group.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
			return group.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (group *NetworkSecurityGroup) validateConfigMapDestinations(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (group *NetworkSecurityGroup) validateOwnerReference(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (group *NetworkSecurityGroup) validateResourceReferences(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (group *NetworkSecurityGroup) validateSecretDestinations(ctx context.Context, obj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (group *NetworkSecurityGroup) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.NetworkSecurityGroup, newObj *v20240301.NetworkSecurityGroup) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
