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

type RouteTable struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-routetable,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetables,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.routetables.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &RouteTable{}

// Default applies defaults to the RouteTable resource
func (table *RouteTable) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.RouteTable)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/RouteTable, but got %T", obj)
	}
	err := table.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = table
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (table *RouteTable) defaultAzureName(ctx context.Context, obj *v20240301.RouteTable) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the RouteTable resource
func (table *RouteTable) defaultImpl(ctx context.Context, obj *v20240301.RouteTable) error {
	err := table.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-routetable,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetables,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.routetables.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &RouteTable{}

// ValidateCreate validates the creation of the resource
func (table *RouteTable) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.RouteTable)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/RouteTable, but got %T", resource)
	}
	validations := table.createValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.RouteTable]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (table *RouteTable) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.RouteTable)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/RouteTable, but got %T", resource)
	}
	validations := table.deleteValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.RouteTable]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (table *RouteTable) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.RouteTable)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/RouteTable, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.RouteTable)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/RouteTable, but got %T", oldResource)
	}
	validations := table.updateValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.RouteTable]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (table *RouteTable) createValidations() []func(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error){table.validateResourceReferences, table.validateOwnerReference, table.validateSecretDestinations, table.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (table *RouteTable) deleteValidations() []func(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (table *RouteTable) updateValidations() []func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
			return table.validateResourceReferences(ctx, newObj)
		},
		table.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
			return table.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
			return table.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
			return table.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (table *RouteTable) validateConfigMapDestinations(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (table *RouteTable) validateOwnerReference(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (table *RouteTable) validateResourceReferences(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (table *RouteTable) validateSecretDestinations(ctx context.Context, obj *v20240301.RouteTable) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (table *RouteTable) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.RouteTable, newObj *v20240301.RouteTable) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
