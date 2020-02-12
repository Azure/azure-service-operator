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

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	v1 "k8s.io/api/core/v1"
)

const AzureSQLServerFinalizerName = finalizerName

func (r *AzureSqlServerReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlServer) error {
	helpers.AddFinalizer(instance, AzureSQLServerFinalizerName)
	if updateerr := r.Update(context.Background(), instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Failed to update finalizer")
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", AzureSQLServerFinalizerName))
	return nil
}
