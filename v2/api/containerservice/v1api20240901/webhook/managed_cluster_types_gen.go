// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240901 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ManagedCluster struct {
}

// +kubebuilder:webhook:path=/mutate-containerservice-azure-com-v1api20240901-managedcluster,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclusters,verbs=create;update,versions=v1api20240901,name=default.v1api20240901.managedclusters.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ManagedCluster{}

// Default applies defaults to the ManagedCluster resource
func (cluster *ManagedCluster) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240901.ManagedCluster)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/ManagedCluster, but got %T", obj)
	}
	err := cluster.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = cluster
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (cluster *ManagedCluster) defaultAzureName(ctx context.Context, obj *v20240901.ManagedCluster) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ManagedCluster resource
func (cluster *ManagedCluster) defaultImpl(ctx context.Context, obj *v20240901.ManagedCluster) error {
	err := cluster.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-containerservice-azure-com-v1api20240901-managedcluster,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclusters,verbs=create;update,versions=v1api20240901,name=validate.v1api20240901.managedclusters.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ManagedCluster{}

// ValidateCreate validates the creation of the resource
func (cluster *ManagedCluster) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240901.ManagedCluster)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/ManagedCluster, but got %T", resource)
	}
	validations := cluster.createValidations()
	var temp any = cluster
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240901.ManagedCluster]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (cluster *ManagedCluster) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240901.ManagedCluster)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/ManagedCluster, but got %T", resource)
	}
	validations := cluster.deleteValidations()
	var temp any = cluster
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240901.ManagedCluster]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (cluster *ManagedCluster) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240901.ManagedCluster)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/ManagedCluster, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240901.ManagedCluster)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/ManagedCluster, but got %T", oldResource)
	}
	validations := cluster.updateValidations()
	var temp any = cluster
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240901.ManagedCluster]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (cluster *ManagedCluster) createValidations() []func(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error){cluster.validateResourceReferences, cluster.validateOwnerReference, cluster.validateSecretDestinations, cluster.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (cluster *ManagedCluster) deleteValidations() []func(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (cluster *ManagedCluster) updateValidations() []func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
			return cluster.validateResourceReferences(ctx, newObj)
		},
		cluster.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
			return cluster.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
			return cluster.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
			return cluster.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (cluster *ManagedCluster) validateConfigMapDestinations(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.ConfigMapDestination
	if obj.Spec.OperatorSpec.ConfigMaps != nil {
		toValidate = []*genruntime.ConfigMapDestination{
			obj.Spec.OperatorSpec.ConfigMaps.OIDCIssuerProfile,
		}
	}
	return configmaps.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (cluster *ManagedCluster) validateOwnerReference(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (cluster *ManagedCluster) validateResourceReferences(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (cluster *ManagedCluster) validateSecretDestinations(ctx context.Context, obj *v20240901.ManagedCluster) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.AdminCredentials,
			obj.Spec.OperatorSpec.Secrets.UserCredentials,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (cluster *ManagedCluster) validateWriteOnceProperties(ctx context.Context, oldObj *v20240901.ManagedCluster, newObj *v20240901.ManagedCluster) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
