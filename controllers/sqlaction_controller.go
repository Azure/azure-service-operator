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
	//"github.com/Azure/azure-service-operator/pkg/errhelp"
	//sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	//"github.com/Azure/go-autorest/autorest/to"
	//v1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/types"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

// SqlActionReconciler reconciles a SqlAction object
type SqlActionReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions/status,verbs=get;update;patch

func (r *SqlActionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlaction", req.NamespacedName)

	var instance azurev1.SqlAction
	// your logic here

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("the instance", "instance", instance)

	return ctrl.Result{}, nil
}

func (r *SqlActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlAction{}).
		Complete(r)
}

func RollCreds(length int) string {
	// logic to generate new password and return it
	newPassword := generateRandomString(length)
	return newPassword
}
