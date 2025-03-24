// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210601 "github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Workspace struct {
}

// +kubebuilder:webhook:path=/mutate-synapse-azure-com-v1api20210601-workspace,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=synapse.azure.com,resources=workspaces,verbs=create;update,versions=v1api20210601,name=default.v1api20210601.workspaces.synapse.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Workspace{}

// Default applies defaults to the Workspace resource
func (workspace *Workspace) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210601.Workspace)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/Workspace, but got %T", obj)
	}
	err := workspace.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = workspace
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (workspace *Workspace) defaultAzureName(ctx context.Context, obj *v20210601.Workspace) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Workspace resource
func (workspace *Workspace) defaultImpl(ctx context.Context, obj *v20210601.Workspace) error {
	err := workspace.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-synapse-azure-com-v1api20210601-workspace,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=synapse.azure.com,resources=workspaces,verbs=create;update,versions=v1api20210601,name=validate.v1api20210601.workspaces.synapse.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Workspace{}

// ValidateCreate validates the creation of the resource
func (workspace *Workspace) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.Workspace)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/Workspace, but got %T", resource)
	}
	validations := workspace.createValidations()
	var temp any = workspace
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.Workspace]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (workspace *Workspace) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.Workspace)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/Workspace, but got %T", resource)
	}
	validations := workspace.deleteValidations()
	var temp any = workspace
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.Workspace]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (workspace *Workspace) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210601.Workspace)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/Workspace, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210601.Workspace)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/Workspace, but got %T", oldResource)
	}
	validations := workspace.updateValidations()
	var temp any = workspace
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.Workspace]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (workspace *Workspace) createValidations() []func(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error){workspace.validateResourceReferences, workspace.validateOwnerReference, workspace.validateSecretDestinations, workspace.validateConfigMapDestinations, workspace.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (workspace *Workspace) deleteValidations() []func(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (workspace *Workspace) updateValidations() []func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
			return workspace.validateResourceReferences(ctx, newObj)
		},
		workspace.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
			return workspace.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
			return workspace.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
			return workspace.validateConfigMapDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
			return workspace.validateOptionalConfigMapReferences(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (workspace *Workspace) validateConfigMapDestinations(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (workspace *Workspace) validateOptionalConfigMapReferences(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (workspace *Workspace) validateOwnerReference(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (workspace *Workspace) validateResourceReferences(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (workspace *Workspace) validateSecretDestinations(ctx context.Context, obj *v20210601.Workspace) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (workspace *Workspace) validateWriteOnceProperties(ctx context.Context, oldObj *v20210601.Workspace, newObj *v20210601.Workspace) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
