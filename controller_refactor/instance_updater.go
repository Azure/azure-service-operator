/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

func (updater *instanceUpdater) setProvisionState(state azurev1alpha1.ProvisionState) {
	updateFunc := func(s *azurev1alpha1.ASOStatus) {
		s.State = string(state)
		s.Provisioned = state == azurev1alpha1.Succeeded
		s.Provisioning = state != azurev1alpha1.Pending &&
			state != azurev1alpha1.Succeeded &&
			state != azurev1alpha1.Failed
	}
	updater.statusUpdate = &updateFunc
}

func (updater *instanceUpdater) setAnnotation(name string, value string) {
	updateFunc := func(meta metav1.Object) {
		annotations := meta.GetAnnotations()
		if annotations == nil {
			annotations = map[string]string{}
		}
		annotations[name] = value
		meta.SetAnnotations(annotations)
	}
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
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
