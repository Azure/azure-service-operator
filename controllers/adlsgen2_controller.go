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
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
)

// AdlsGen2Reconciler reconciles a AdlsGen2 object
type AdlsGen2Reconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=adlsgen2s,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=adlsgen2s/status,verbs=get;update;patch

func (r *AdlsGen2Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("adlsgen2", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *AdlsGen2Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.AdlsGen2{}).
		Complete(r)
}
