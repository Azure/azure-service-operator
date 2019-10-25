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
type statusUpdate = func(provisionState *azurev1alpha1.ASOStatus)
type metaUpdate = func(meta metav1.Object)

// instanceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type instanceUpdater struct {
	StatusUpdater
	metaUpdates  []metaUpdate
	statusUpdate *statusUpdate
}

func (updater *instanceUpdater) addFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.AddFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *instanceUpdater) removeFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.RemoveFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *instanceUpdater) setProvisionState(provisionState azurev1alpha1.ProvisionState) {
	updateFunc := func(s *azurev1alpha1.ASOStatus) {
		s.State = string(provisionState)
		if provisionState == azurev1alpha1.Verifying {
			s.Provisioning = true
		}
		if provisionState == azurev1alpha1.Succeeded {
			s.Provisioned = true
		}
	}
	updater.statusUpdate = &updateFunc
}

func (updater *instanceUpdater) setOwnerReferences(owners []runtime.Object) {
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

func (updater *instanceUpdater) applyUpdates(instance runtime.Object, status *azurev1alpha1.ASOStatus) error {
	if updater.statusUpdate != nil {
		(*updater.statusUpdate)(status)
	}
	err := updater.StatusUpdater(instance, status)
	m, _ := apimeta.Accessor(instance)
	for _, f := range updater.metaUpdates {
		f(m)
	}
	return err
}

func (updater *instanceUpdater) clear() {
	updater.metaUpdates = []metaUpdate{}
	updater.statusUpdate = nil
}

func (updater *instanceUpdater) hasUpdates() bool {
	return len(updater.metaUpdates) > 0 || updater.statusUpdate != nil
}
