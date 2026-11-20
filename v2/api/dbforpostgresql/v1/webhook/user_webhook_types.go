// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	v1 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type User_Webhook struct{}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1-user,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=users,verbs=create;update,versions=v1,name=default.v1.users.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &User_Webhook{}

func (webhook *User_Webhook) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v1.User)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1/User, but got %T", obj)
	}
	err := webhook.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = webhook
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (webhook *User_Webhook) defaultAzureName(_ context.Context, user *v1.User) error {
	if user.Spec.AzureName == "" {
		user.Spec.AzureName = user.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the FlexibleServer resource
func (webhook *User_Webhook) defaultImpl(ctx context.Context, user *v1.User) error {
	err := webhook.defaultAzureName(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1-user,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=users,verbs=create;update,versions=v1,name=validate.v1.users.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &User_Webhook{}

// ValidateCreate validates the creation of the resource
func (webhook *User_Webhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	resource, ok := obj.(*v1.User)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1/User, but got %T", obj)
	}
	validations := webhook.createValidations()
	var temp any = webhook
	if runtimeValidator, ok := temp.(genruntime.Validator[*v1.User]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, resource, validations)
}

// ValidateDelete validates the deletion of the resource
func (webhook *User_Webhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	resource, ok := obj.(*v1.User)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1/User, but got %T", obj)
	}
	validations := webhook.deleteValidations()
	var temp any = webhook
	if runtimeValidator, ok := temp.(genruntime.Validator[*v1.User]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, resource, validations)
}

// ValidateUpdate validates an update of the resource
func (webhook *User_Webhook) ValidateUpdate(ctx context.Context, oldObj runtime.Object, newObj runtime.Object) (admission.Warnings, error) {
	newResource, ok := newObj.(*v1.User)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1/User, but got %T", newObj)
	}
	oldResource, ok := oldObj.(*v1.User)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1/User, but got %T", oldObj)
	}
	validations := webhook.updateValidations()
	var temp any = webhook
	if runtimeValidator, ok := temp.(genruntime.Validator[*v1.User]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldResource,
		newResource,
		validations,
	)
}

// createValidations validates the creation of the resource
func (webhook *User_Webhook) createValidations() []func(ctx context.Context, obj *v1.User) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v1.User) (admission.Warnings, error){
		webhook.validateIsLocalOrAAD,
		webhook.validateAADPrincipal,
	}
}

// deleteValidations validates the deletion of the resource
func (webhook *User_Webhook) deleteValidations() []func(ctx context.Context, obj *v1.User) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (webhook *User_Webhook) updateValidations() []func(ctx context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error) {
			return webhook.validateIsLocalOrAAD(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error) {
			return webhook.validateAADPrincipal(ctx, newObj)
		},
		webhook.validateUserTypeNotChanged,
		webhook.validateAADPrincipalNotChanged,
	}
}

func (webhook *User_Webhook) validateIsLocalOrAAD(_ context.Context, obj *v1.User) (admission.Warnings, error) {
	if (obj.Spec.LocalUser == nil) == (obj.Spec.AADUser == nil) {
		return nil, eris.Errorf("exactly one of spec.localUser or spec.aadUser must be set")
	}

	return nil, nil
}

func (webhook *User_Webhook) validateAADPrincipal(_ context.Context, obj *v1.User) (admission.Warnings, error) {
	if obj.Spec.AADUser == nil {
		return nil, nil
	}

	if (obj.Spec.AADUser.ObjectId == "") != (obj.Spec.AADUser.PrincipalType == "") {
		return nil, eris.Errorf("spec.aadUser.objectId and spec.aadUser.principalType must be set together")
	}

	return nil, nil
}

func (webhook *User_Webhook) validateUserTypeNotChanged(_ context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error) {
	// Prevent change from AAD -> Local
	if oldObj.Spec.AADUser != nil && newObj.Spec.AADUser == nil {
		return nil, eris.Errorf("cannot change from AAD User to local user")
	}

	// Prevent change from Local -> AAD
	if oldObj.Spec.LocalUser != nil && newObj.Spec.LocalUser == nil {
		return nil, eris.Errorf("cannot change from local user to AAD user")
	}

	return nil, nil
}

// validateAADPrincipalNotChanged prevents changing the principal an AAD user is bound to: the role is bound once, when
// it is created.
func (webhook *User_Webhook) validateAADPrincipalNotChanged(_ context.Context, oldObj *v1.User, newObj *v1.User) (admission.Warnings, error) {
	if oldObj.Spec.AADUser == nil || newObj.Spec.AADUser == nil {
		return nil, nil
	}

	if oldObj.Spec.AADUser.ObjectId != newObj.Spec.AADUser.ObjectId {
		return nil, eris.Errorf("cannot change AAD user 'objectId' from %q to %q", oldObj.Spec.AADUser.ObjectId, newObj.Spec.AADUser.ObjectId)
	}

	if oldObj.Spec.AADUser.PrincipalType != newObj.Spec.AADUser.PrincipalType {
		return nil, eris.Errorf("cannot change AAD user 'principalType' from %q to %q", oldObj.Spec.AADUser.PrincipalType, newObj.Spec.AADUser.PrincipalType)
	}

	return nil, nil
}
