// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220120p "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type FlexibleServersConfiguration struct {
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversconfiguration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1api20220120preview,name=default.v1api20220120preview.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &FlexibleServersConfiguration{}

// Default applies defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220120p.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/FlexibleServersConfiguration, but got %T", obj)
	}
	err := configuration.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = configuration
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (configuration *FlexibleServersConfiguration) defaultAzureName(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) defaultImpl(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) error {
	err := configuration.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1api20220120preview-flexibleserversconfiguration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1api20220120preview,name=validate.v1api20220120preview.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &FlexibleServersConfiguration{}

// ValidateCreate validates the creation of the resource
func (configuration *FlexibleServersConfiguration) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220120p.FlexibleServersConfiguration)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/FlexibleServersConfiguration, but got %T", resource)
	}
	validations := configuration.createValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220120p.FlexibleServersConfiguration]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220120p.FlexibleServersConfiguration)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/FlexibleServersConfiguration, but got %T", resource)
	}
	validations := configuration.deleteValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220120p.FlexibleServersConfiguration]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (configuration *FlexibleServersConfiguration) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220120p.FlexibleServersConfiguration)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/FlexibleServersConfiguration, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220120p.FlexibleServersConfiguration)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/FlexibleServersConfiguration, but got %T", oldResource)
	}
	validations := configuration.updateValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220120p.FlexibleServersConfiguration]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (configuration *FlexibleServersConfiguration) createValidations() []func(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error){configuration.validateResourceReferences, configuration.validateOwnerReference, configuration.validateSecretDestinations, configuration.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) deleteValidations() []func(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (configuration *FlexibleServersConfiguration) updateValidations() []func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
			return configuration.validateResourceReferences(ctx, newObj)
		},
		configuration.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
			return configuration.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
			return configuration.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
			return configuration.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (configuration *FlexibleServersConfiguration) validateConfigMapDestinations(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (configuration *FlexibleServersConfiguration) validateOwnerReference(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (configuration *FlexibleServersConfiguration) validateResourceReferences(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (configuration *FlexibleServersConfiguration) validateSecretDestinations(ctx context.Context, obj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (configuration *FlexibleServersConfiguration) validateWriteOnceProperties(ctx context.Context, oldObj *v20220120p.FlexibleServersConfiguration, newObj *v20220120p.FlexibleServersConfiguration) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
