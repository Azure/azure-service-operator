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

package reconciler

import (
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"k8s.io/apimachinery/pkg/runtime"

	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// modifies the runtime.Object in place
type statusUpdate = func(status *Status)
type metaUpdate = func(meta metav1.Object)

// instanceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type instanceUpdater struct {
	StatusUpdater
	metaUpdates   []metaUpdate
	statusUpdates []statusUpdate
}

func (updater *instanceUpdater) addFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.AddFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *instanceUpdater) removeFinalizer(name string) {
	updateFunc := func(meta metav1.Object) { helpers.RemoveFinalizer(meta, name) }
	updater.metaUpdates = append(updater.metaUpdates, updateFunc)
}

func (updater *instanceUpdater) setStatusPayload(statusPayload interface{}) {
	updateFunc := func(s *Status) {
		s.StatusPayload = statusPayload
	}
	updater.statusUpdates = append(updater.statusUpdates, updateFunc)
}

func (updater *instanceUpdater) setReconcileState(state ReconcileState, message string) {
	updateFunc := func(s *Status) {
		s.State = state
		s.Message = message
	}
	updater.statusUpdates = append(updater.statusUpdates, updateFunc)
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

func (updater *instanceUpdater) applyUpdates(instance runtime.Object, status *Status) error {
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

func (updater *instanceUpdater) clear() {
	updater.metaUpdates = []metaUpdate{}
	updater.statusUpdates = []statusUpdate{}
}

func (updater *instanceUpdater) hasUpdates() bool {
	return len(updater.metaUpdates) > 0 || len(updater.statusUpdates) > 0
}
