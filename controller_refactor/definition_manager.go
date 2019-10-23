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
	// returns ThisResourceDefinitions
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

var NoDependencies = DependencyDefinitions{
	Dependencies: []*CustomResourceDetails{},
	Owner:        nil,
}

type CustomResourceDetails struct {
	Instance       runtime.Object
	Status         *azurev1alpha1.ResourceStatus
	BaseDefinition *azurev1alpha1.ResourceBaseDefinition
}

// modifies the runtime.Object in place
type stateUpdate = func(*azurev1alpha1.ResourceBaseDefinition)

type statusUpdate = func(provisionState *azurev1alpha1.ResourceStatus)
type metaUpdate = func(meta *metav1.ObjectMeta)

// CustomResourceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type CustomResourceUpdater struct {
	UpdateInstance func(*azurev1alpha1.ResourceBaseDefinition)
	metaUpdates    []metaUpdate
	stateUpdates   []stateUpdate
	statusUpdates  []statusUpdate
}

type CustomResourceUpdaters struct {
	UpdateInstance func(*azurev1alpha1.ResourceBaseDefinition)
	stateUpdates   []stateUpdate
}

func (updater *CustomResourceUpdater) AddFinalizer(name string) {
	updateFunc := func(meta *metav1.ObjectMeta) { meta.Finalizers = append(meta.Finalizers, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *CustomResourceUpdater) RemoveFinalizer(name string) {
	updateFunc := func(meta *metav1.ObjectMeta) { meta.Finalizers = append(meta.Finalizers, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *CustomResourceUpdater) SetProvisionState(provisionState azurev1alpha1.ProvisionState) {
	updateFunc := func(s *azurev1alpha1.ResourceStatus) {
		s.ProvisionState = provisionState
		if provisionState == azurev1alpha1.Verifying {
			s.Provisioning = true
		}
		if provisionState == azurev1alpha1.Succeeded {
			s.Provisioned = true
		}
	}
	updater.statusUpdates = append(updater.statusUpdates, updateFunc)
}

func (updater *CustomResourceUpdater) SetOwnerReferences(ownerDetails []*CustomResourceDetails) {
	updateFunc := func(s *azurev1alpha1.ResourceBaseDefinition) {
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
		s.ObjectMeta.SetOwnerReferences(references)
	}
	updater.stateUpdates = append(updater.stateUpdates, updateFunc)
}

func (updater *CustomResourceUpdater) ApplyUpdates(state *azurev1alpha1.ResourceBaseDefinition) {
	for _, f := range updater.stateUpdates {
		f(state)
	}
	updater.UpdateInstance(state)
}

func (updater *CustomResourceUpdater) Clear() {
	updater.stateUpdates = []stateUpdate{}
}
