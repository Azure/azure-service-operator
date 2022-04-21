// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var azuresqlmanageduserlog = logf.Log.WithName("azuresqlmanageduser-resource")

func (r *AzureSQLManagedUser) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-azure-microsoft-com-v1alpha1-azuresqlmanageduser,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=azuresqlmanagedusers,versions=v1alpha1,name=vazuresqlmanageduser.kb.io,sideEffects=none,webhookVersions=v1,admissionReviewVersions=v1;v1beta1

var _ webhook.Validator = &AzureSQLManagedUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLManagedUser) ValidateCreate() error {
	azuresqlmanageduserlog.Info("validate create", "name", r.Name)

	return ValidateAzureSQLDBName(r.Spec.DbName)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLManagedUser) ValidateUpdate(old runtime.Object) error {
	azuresqlmanageduserlog.Info("validate update", "name", r.Name)

	return ValidateAzureSQLDBName(r.Spec.DbName)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLManagedUser) ValidateDelete() error {
	azuresqlmanageduserlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
