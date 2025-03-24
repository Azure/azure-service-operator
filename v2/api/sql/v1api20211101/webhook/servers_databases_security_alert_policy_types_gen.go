// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ServersDatabasesSecurityAlertPolicy struct {
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabasessecurityalertpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasessecurityalertpolicies,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabasessecurityalertpolicies.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServersDatabasesSecurityAlertPolicy{}

// Default applies defaults to the ServersDatabasesSecurityAlertPolicy resource
func (policy *ServersDatabasesSecurityAlertPolicy) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.ServersDatabasesSecurityAlertPolicy)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesSecurityAlertPolicy, but got %T", obj)
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

// defaultImpl applies the code generated defaults to the ServersDatabasesSecurityAlertPolicy resource
func (policy *ServersDatabasesSecurityAlertPolicy) defaultImpl(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabasessecurityalertpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasessecurityalertpolicies,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabasessecurityalertpolicies.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServersDatabasesSecurityAlertPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabasesSecurityAlertPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesSecurityAlertPolicy, but got %T", resource)
	}
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesSecurityAlertPolicy]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabasesSecurityAlertPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesSecurityAlertPolicy, but got %T", resource)
	}
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesSecurityAlertPolicy]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.ServersDatabasesSecurityAlertPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesSecurityAlertPolicy, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.ServersDatabasesSecurityAlertPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesSecurityAlertPolicy, but got %T", oldResource)
	}
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesSecurityAlertPolicy]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) createValidations() []func(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference, policy.validateSecretDestinations, policy.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) deleteValidations() []func(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) updateValidations() []func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
			return policy.validateResourceReferences(ctx, newObj)
		},
		policy.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
			return policy.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
			return policy.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
			return policy.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (policy *ServersDatabasesSecurityAlertPolicy) validateConfigMapDestinations(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (policy *ServersDatabasesSecurityAlertPolicy) validateOwnerReference(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (policy *ServersDatabasesSecurityAlertPolicy) validateResourceReferences(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (policy *ServersDatabasesSecurityAlertPolicy) validateSecretDestinations(ctx context.Context, obj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersDatabasesSecurityAlertPolicy) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.ServersDatabasesSecurityAlertPolicy, newObj *v20211101.ServersDatabasesSecurityAlertPolicy) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
