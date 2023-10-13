// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var azuresqluserlog = logf.Log.WithName("azuresqluser-resource")

func (r *AzureSQLUser) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-azure-microsoft-com-v1alpha1-azuresqluser,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=azuresqlusers,versions=v1alpha1,name=vazuresqluser.kb.io,sideEffects=none,webhookVersions=v1,admissionReviewVersions=v1;v1beta1

func ValidateAzureSQLDBName(name string) error {
	if name == "master" {
		return errors.Errorf("'master' is a reserved database name and cannot be used")
	}

	return nil
}

var _ webhook.Validator = &AzureSQLUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLUser) ValidateCreate() (admission.Warnings, error) {
	azuresqluserlog.Info("validate create", "name", r.Name)

	return nil, ValidateAzureSQLDBName(r.Spec.DbName)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLUser) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	azuresqluserlog.Info("validate update", "name", r.Name)

	return nil, ValidateAzureSQLDBName(r.Spec.DbName)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AzureSQLUser) ValidateDelete() (admission.Warnings, error) {
	azuresqluserlog.Info("validate delete", "name", r.Name)
	return nil, nil
}
