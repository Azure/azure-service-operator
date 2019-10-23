package controller_refactor

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"k8s.io/apimachinery/pkg/runtime"

	// "k8s.io/apimachinery/pkg/api/meta"

	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// modifies the runtime.Object in place
type statusUpdate = func(provisionState *azurev1alpha1.ResourceStatus)
type metaUpdate = func(meta metav1.Object)

// customResourceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type customResourceUpdater struct {
	StatusUpdater
	metaUpdates   []metaUpdate
	statusUpdates []statusUpdate
}

func (updater *customResourceUpdater) addFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.AddFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *customResourceUpdater) removeFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.RemoveFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *customResourceUpdater) setProvisionState(provisionState azurev1alpha1.ProvisionState) {
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

func (updater *customResourceUpdater) setOwnerReferences(owners []runtime.Object) {
	updateFunc := func(s metav1.Object) {
		references := make([]metav1.OwnerReference, len(owners))
		for i, o := range owners {
			meta, _ := apimeta.Accessor(o)
			references[i] = metav1.OwnerReference{
				APIVersion: "v1",
				Kind:       o.GetObjectKind().GroupVersionKind().Kind,
				Name:       meta.GetName(),
				UID:        meta.GetUID(),
			}
		}
		s.SetOwnerReferences(references)
	}
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *customResourceUpdater) applyUpdates(instance runtime.Object, status *azurev1alpha1.ResourceStatus) error {
	for _, f := range updater.statusUpdates {
		f(status)
	}
	err := updater.StatusUpdater(instance, status)
	m, _ := apimeta.Accessor(instance)
	for _, f := range updater.metaUpdates {
		f(m)
	}
	return err
}

func (updater *customResourceUpdater) clear() {
	updater.metaUpdates = []metaUpdate{}
	updater.statusUpdates = []statusUpdate{}
}
