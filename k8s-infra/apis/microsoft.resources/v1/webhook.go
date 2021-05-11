/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var resourcegrouplog = logf.Log.WithName("resourcegroup-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-resources-infra-azure-com-v1-resourcegroup,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=create;update,versions=v1,name=default.resourcegroup.infra.azure.com

var _ webhook.Defaulter = &ResourceGroup{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ResourceGroup) Default() {
	resourcegrouplog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-resources-infra-azure-com-v1-resourcegroup,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=microsoft.resources.infra.azure.com,resources=resourcegroups,versions=v1,name=validation.resourcegroup.infra.azure.com

var _ webhook.Validator = &ResourceGroup{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ResourceGroup) ValidateCreate() error {
	resourcegrouplog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ResourceGroup) ValidateUpdate(old runtime.Object) error {
	resourcegrouplog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ResourceGroup) ValidateDelete() error {
	resourcegrouplog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
