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
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ResourceGroupReconciler reconciles a ResourceGroup object
type ResourceGroupReconciler struct {
	client.Client
	Log                  logr.Logger
	Recorder             record.EventRecorder
	Reconciler           *AsyncReconciler
	ResourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *ResourceGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.ResourceGroup{})
}

// SetupWithManager function sets up the functions with the controller
func (r *ResourceGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
}
