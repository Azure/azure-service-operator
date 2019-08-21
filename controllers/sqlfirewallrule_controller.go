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
	"fmt"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
)

// SqlFirewallRuleReconciler reconciles a SqlFirewallRule object
type SqlFirewallRuleReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlfirewallrules/status,verbs=get;update;patch

func (r *SqlFirewallRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlfirewallrule", req.NamespacedName)

	// your logic here
	var instance azurev1.SqlFirewallRule

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql-firewall-rule resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *SqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlFirewallRule{}).
		Complete(r)
}

func (r *SqlFirewallRuleReconciler) reconcileExternal(instance *azurev1.SqlFirewallRule) error {
	// ctx := context.Background()
	// location := instance.Spec.Location
	// name := instance.ObjectMeta.Name
	// groupName := instance.Spec.ResourceGroupName

	// // write information back to instance
	// instance.Status.Provisioning = true

	// if err := r.Status().Update(ctx, instance); err != nil {
	// 	r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	// }

	//err = CreateResource()
	// if err != nil {

	// }

	// instance.Status.Provisioning = false
	// instance.Status.Provisioned = true

	// if err = r.Status().Update(ctx, instance); err != nil {
	// 	r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	// }

	return nil
}

func (r *SqlFirewallRuleReconciler) deleteExternal(instance *azurev1.SqlFirewallRule) error {
	// ctx := context.Background()
	name := instance.ObjectMeta.Name
	// groupName := instance.Spec.ResourceGroupName
	// // delete resource

	r.Recorder.Event(instance, "Normal", "Deleted", name+" deleted")
	return nil
}

func (r *SqlFirewallRuleReconciler) addFinalizer(instance *azurev1.SqlFirewallRule) error {
	helpers.AddFinalizer(instance, SQLServerFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", SQLServerFinalizerName))
	return nil
}
