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

package controllers

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	"github.com/Azure/azure-service-operator/pkg/telemetry"

	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AppInsightsReconciler reconciles the desired Application Insights service reference with k8s
type AppInsightsReconciler struct {
	client.Client
	Log                logr.Logger
	Telemetry          telemetry.PrometheusTelemetry
	Recorder           record.EventRecorder
	Scheme             *runtime.Scheme
	AppInsightsManager appinsights.Manager
	Reconciler         *AsyncReconciler
}

// Reconcile attempts to set the desired state representation for the Application Insights operator
func (r *AppInsightsReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AppInsights{})
}

// SetupWithManager starts the operator in k8s
func (r *AppInsightsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AppInsights{}).
		Complete(r)
}
