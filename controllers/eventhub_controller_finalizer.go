/*
Copyright 2019 microsoft.

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

package controllers

import (
	"context"
	"fmt"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	v1 "k8s.io/api/core/v1"
)

const eventhubFinalizerName = "eventhub.finalizers.com"

func (r *EventhubReconciler) addFinalizer(instance *azurev1.Eventhub) error {
	instance.AddFinalizer(eventhubFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", eventhubFinalizerName))
	return nil
}

func (r *EventhubReconciler) handleFinalizer(instance *azurev1.Eventhub) error {
	if instance.HasFinalizer(eventhubFinalizerName) {
		// our finalizer is present, so lets handle our external dependency
		if err := r.deleteExternalDependency(instance); err != nil {

			return errhelp.NewAzureError(err)
		}

		instance.RemoveFinalizer(eventhubFinalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return err
		}
	}
	// Our finalizer has finished, so the reconciler can do nothing.
	return nil
}

func (r *EventhubReconciler) deleteExternalDependency(instance *azurev1.Eventhub) error {

	return r.deleteEventhub(instance)
}
