// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"context"

	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	v1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

type ServicePrincipal_Webhook struct{}

// +kubebuilder:webhook:path=/mutate-entra-azure-com-v1-serviceprincipal,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=serviceprincipals,verbs=create;update,versions=v1,name=default.v1.serviceprincipals.entra.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServicePrincipal_Webhook{}

func (*ServicePrincipal_Webhook) Default(_ context.Context, obj runtime.Object) error {
	sp, ok := obj.(*v1.ServicePrincipal)
	if !ok {
		return eris.Errorf("cannot default resource of type %T as ServicePrincipal", obj)
	}
	if sp.Spec.OperatorSpec == nil {
		sp.Spec.OperatorSpec = &v1.ServicePrincipalOperatorSpec{}
	}
	if sp.Spec.OperatorSpec.CreationMode == nil {
		sp.Spec.OperatorSpec.CreationMode = to.Ptr(v1.AdoptOrCreate)
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-entra-azure-com-v1-serviceprincipal,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=serviceprincipals,verbs=create;update,versions=v1,name=validate.v1.serviceprincipals.entra.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServicePrincipal_Webhook{}

func (*ServicePrincipal_Webhook) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	if _, ok := obj.(*v1.ServicePrincipal); !ok {
		return nil, eris.Errorf("cannot validate resource of type %T as ServicePrincipal", obj)
	}
	return nil, nil
}

func (*ServicePrincipal_Webhook) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	oldSP, ok := oldObj.(*v1.ServicePrincipal)
	if !ok {
		return nil, eris.Errorf("cannot validate resource of type %T as ServicePrincipal", oldObj)
	}
	newSP, ok := newObj.(*v1.ServicePrincipal)
	if !ok {
		return nil, eris.Errorf("cannot validate resource of type %T as ServicePrincipal", newObj)
	}
	if (oldSP.Spec.AppId == nil) != (newSP.Spec.AppId == nil) ||
		(oldSP.Spec.AppId != nil && newSP.Spec.AppId != nil && *oldSP.Spec.AppId != *newSP.Spec.AppId) {
		return nil, eris.New("spec.appId cannot be changed after creation")
	}
	return nil, nil
}

func (*ServicePrincipal_Webhook) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	if _, ok := obj.(*v1.ServicePrincipal); !ok {
		return nil, eris.Errorf("cannot validate resource of type %T as ServicePrincipal", obj)
	}
	return nil, nil
}
