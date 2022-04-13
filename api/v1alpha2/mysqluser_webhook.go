// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
)

// log is for logging in this package.
var mysqluserlog = logf.Log.WithName("mysqluser-resource")

func (r *MySQLUser) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:verbs=create;update,path=/validate-azure-microsoft-com-v1alpha2-mysqluser,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=mysqlusers,versions=v1alpha2,name=vmysqluser.kb.io,sideEffects=none,webhookVersions=v1,admissionReviewVersions=v1;v1beta1

func ensureNoSQLAll(privileges []string) error {
	for _, privilege := range privileges {
		if helpers.IsSQLAll(privilege) {
			return errors.Errorf("ASO admin user doesn't have privileges to grant ALL at server level")
		}
	}
	return nil
}

var _ webhook.Validator = &MySQLUser{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLUser) ValidateCreate() error {
	mysqluserlog.Info("validate create", "name", r.Name)

	return ensureNoSQLAll(r.Spec.Roles)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLUser) ValidateUpdate(old runtime.Object) error {
	mysqluserlog.Info("validate update", "name", r.Name)

	return ensureNoSQLAll(r.Spec.Roles)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *MySQLUser) ValidateDelete() error {
	mysqluserlog.Info("validate delete", "name", r.Name)
	return nil
}
