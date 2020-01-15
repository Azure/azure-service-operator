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
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AzureSqlFirewallRuleReconciler reconciles a AzureSqlFirewallRule object
type AzureSqlFirewallRuleReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureSqlFirewallRuleReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
<<<<<<< HEAD
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureSqlFirewallRule{})
=======
	ctx := context.Background()

	// your logic here
	var instance azurev1alpha1.AzureSqlFirewallRule

	// log operator start
	r.Telemetry.LogStart()

	defer func() {

		// log failure / success
		if err != nil {
			r.Telemetry.LogError(
				"Failure occured during reconcilliation",
				err)
			r.Telemetry.LogFailure()
		} else if result.Requeue {
			r.Telemetry.LogFailure()
		} else {
			r.Telemetry.LogSuccess()
		}

		if errUpdate := r.Status().Update(ctx, &instance); errUpdate != nil {
			r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	}()

	if err = r.Get(ctx, req.NamespacedName, &instance); err != nil {
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, azureSQLFirewallRuleFinalizerName) {
			if err = r.deleteExternal(ctx, &instance); err != nil {
				instance.Status.Message = fmt.Sprintf("Delete AzureSqlFirewallRule failed with %s", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, azureSQLFirewallRuleFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, azureSQLFirewallRuleFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			instance.Status.Message = fmt.Sprintf("Adding AzureSqlFirewallRule finalizer failed with %s", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation for AzureSqlFirewallRule")
		if err := r.reconcileExternal(ctx, &instance); err != nil {
			instance.Status.Message = err.Error()
			instance.Status.Provisioning = false
			r.Telemetry.LogError("Reconcile external failed", err)
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "azuresqlfirewallrule "+instance.ObjectMeta.Name+" provisioned ")
	instance.Status.Message = successMsg

	return ctrl.Result{}, nil
>>>>>>> 07fc82e12f504c7df9e5463cca82763db8af5bd8
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureSqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlFirewallRule{}).
		Complete(r)
}
