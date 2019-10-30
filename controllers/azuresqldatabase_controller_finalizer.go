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
)

const AzureSQLDatabaseFinalizerName = "azuresqldatabase.finalizers.azure.com"

func (r *AzureSqlDatabaseReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlDatabase) error {
	helpers.AddFinalizer(instance, AzureSQLDatabaseFinalizerName)
	if updateerr := r.Update(context.Background(), instance); updateerr != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Failed to update finalizer")
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", AzureSQLDatabaseFinalizerName))
	return nil
}
