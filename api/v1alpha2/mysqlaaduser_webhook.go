// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var mysqlaaduserlog = logf.Log.WithName("mysqlaaduser-resource")

func (r *MySQLAADUser) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:verbs=create;update,path=/validate-azure-microsoft-com-v1alpha2-mysqlaaduser,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=mysqlaadusers,versions=v1alpha2,name=vmysqlaaduser.kb.io,sideEffects=none,webhookVersions=v1,admissionReviewVersions=v1;v1beta1

var _ webhook.Validator = &MySQLAADUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLAADUser) ValidateCreate() error {
	mysqlaaduserlog.Info("validate create", "name", r.Name)

	return ensureNoSQLAll(r.Spec.Roles)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLAADUser) ValidateUpdate(old runtime.Object) error {
	mysqlaaduserlog.Info("validate update", "name", r.Name)

	return ensureNoSQLAll(r.Spec.Roles)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLAADUser) ValidateDelete() error {
	mysqlaaduserlog.Info("validate delete", "name", r.Name)
	return nil
}
