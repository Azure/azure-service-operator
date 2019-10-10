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
	"github.com/Azure/azure-service-operator/pkg/helpers"
)

const consumerGroupFinalizerName = "consumergroup.finalizers.com"

func (r *ConsumerGroupReconciler) addFinalizer(instance *azurev1.ConsumerGroup) error {
	helpers.AddFinalizer(consumerGroupFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", consumerGroupFinalizerName))
	return nil
}

func (r *ConsumerGroupReconciler) handleFinalizer(instance *azurev1.ConsumerGroup) error {
	if helpers.HasFinalizer(consumerGroupFinalizerName) {
		// our finalizer is present, so lets handle our external dependency
		if err := r.deleteExternalDependency(instance); err != nil {
			return err
		}

		helpers.RemoveFinalizer(consumerGroupFinalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return err
		}
	}
	// Our finalizer has finished, so the reconciler can do nothing.
	return nil
}

func (r *ConsumerGroupReconciler) deleteExternalDependency(instance *azurev1.ConsumerGroup) error {

	return r.deleteConsumerGroup(instance)
}
