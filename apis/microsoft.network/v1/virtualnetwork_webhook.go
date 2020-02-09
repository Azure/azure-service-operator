/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var virtualnetworklog = logf.Log.WithName("virtualnetwork-resource")

func (r *VirtualNetwork) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-virtualnetwork,mutating=true,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=virtualnetworks,verbs=create;update,versions=v1,name=mvirtualnetwork.kb.io

var _ webhook.Defaulter = &VirtualNetwork{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VirtualNetwork) Default() {
	virtualnetworklog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-virtualnetwork,mutating=false,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=virtualnetworks,versions=v1,name=vvirtualnetwork.kb.io

var _ webhook.Validator = &VirtualNetwork{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateCreate() error {
	virtualnetworklog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateUpdate(old runtime.Object) error {
	virtualnetworklog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualNetwork) ValidateDelete() error {
	virtualnetworklog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
