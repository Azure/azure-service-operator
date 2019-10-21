package controller_refactor

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

// DefinitionManager is used to retrieve the required custom resource definitions
// and convert them into a state that can be consumed and updated (where applicable) generically
type DefinitionManager interface {
	GetThis(ctx context.Context, req ctrl.Request) (*ThisResourceDefinitions, error)
	// if any dependency is not found, should return nil and a NotFound api error
	GetDependencies(ctx context.Context, thisInstance runtime.Object) (*DependencyDefinitions, error)
}

// Details of the current resource being reconciled
type ThisResourceDefinitions struct {
	Details *CustomResourceDetails
	Updater *CustomResourceUpdater
}

// Details of the owner and the dependencies of the resource
type DependencyDefinitions struct {
	Owner        *CustomResourceDetails
	Dependencies []*CustomResourceDetails
}

type CustomResourceDetails struct {
	Name           string
	ProvisionState azurev1alpha1.ProvisionState
	Instance       runtime.Object
	BaseDefinition *azurev1alpha1.ResourceBaseDefinition
	IsBeingDeleted bool
}

// CustomResourceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type CustomResourceUpdater struct {
	CustomResourceDetails *CustomResourceDetails
	UpdateInstance        func(*azurev1alpha1.ResourceBaseDefinition)
}

func (updater *CustomResourceUpdater) AddFinalizer(name string) {
	baseState := updater.CustomResourceDetails.BaseDefinition
	baseState.AddFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) RemoveFinalizer(name string) {
	baseState := updater.CustomResourceDetails.BaseDefinition
	baseState.RemoveFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) SetProvisionState(provisionState azurev1alpha1.ProvisionState) {
	state := updater.CustomResourceDetails.BaseDefinition
	state.Status.ProvisionState = provisionState
	if provisionState == azurev1alpha1.Verifying {
		state.Status.Provisioning = true
	}
	if provisionState == azurev1alpha1.Succeeded {
		state.Status.Provisioned = true
	}
	updater.UpdateInstance(state)
}

func (updater *CustomResourceUpdater) SetOwnerReferences(ownerDetails []*CustomResourceDetails) {
	references := make([]metav1.OwnerReference, len(ownerDetails))
	for i, o := range ownerDetails {
		ownerBase := o.BaseDefinition
		references[i] = metav1.OwnerReference{
			APIVersion: "v1",
			Kind:       ownerBase.Kind,
			Name:       ownerBase.Name,
			UID:        ownerBase.GetUID(),
		}
	}
	state := updater.CustomResourceDetails.BaseDefinition
	state.ObjectMeta.SetOwnerReferences(references)
	updater.UpdateInstance(state)
}
