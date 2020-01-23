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
)

const apiMSFinalizerName = "apimanagementservice.finalizers.com"

func (r *ApiManagementServiceReconciler) addFinalizer(instance *azurev1.ApiManagementService) error {
	instance.AddFinalizer(apiMSFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", apiMSFinalizerName))
	return nil
}

func (r *ApiManagementServiceReconciler) handleFinalizer(instance *azurev1.ApiManagementService) error {
	if instance.HasFinalizer(apiMSFinalizerName) {
		// our finalizer is present, so lets handle our external dependency
		if err := r.deleteExternalDependency(instance); err != nil {
			return err
		}

		instance.RemoveFinalizer(apiMSFinalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return err
		}
	}
	// Our finalizer has finished, so the reconciler can do nothing.
	return nil
}

func (r *ApiManagementServiceReconciler) deleteExternalDependency(instance *azurev1.ApiManagementService) error {

	return nil
	//return r.deleteAPIManagement(instance)
}
