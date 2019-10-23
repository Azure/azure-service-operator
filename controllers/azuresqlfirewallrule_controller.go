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
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
)

const azureSQLFirewallRuleFinalizerName = "azuresqlfirewallrule.finalizers.azure.com"

// AzureSqlFirewallRuleReconciler reconciles a AzureSqlFirewallRule object
type AzureSqlFirewallRuleReconciler struct {
	client.Client
	Telemetry telemetry.PrometheusTelemetry
	Recorder  record.EventRecorder
	Scheme    *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules/status,verbs=get;update;patch

func (r *AzureSqlFirewallRuleReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
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

	// init the resource manager for this reconcilliation
	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: instance.Spec.ResourceGroup,
		ServerName:        instance.Spec.Server,
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, azureSQLFirewallRuleFinalizerName) {
			if err = r.deleteExternal(&instance, sdkClient); err != nil {
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
		if err := r.reconcileExternal(&instance, sdkClient); err != nil {
			instance.Status.Message = fmt.Sprintf("Reconcile external failed with %s", err.Error())
			r.Telemetry.LogError("Reconcile external failed", err)
			return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "azuresqlfirewallrule "+instance.ObjectMeta.Name+" provisioned ")
	instance.Status.Message = fmt.Sprintf("AzureSqlFirewallrule %s successfully provisioned", instance.ObjectMeta.Name)

	return ctrl.Result{}, nil
}

func (r *AzureSqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlFirewallRule{}).
		Complete(r)
}

func (r *AzureSqlFirewallRuleReconciler) reconcileExternal(instance *azurev1alpha1.AzureSqlFirewallRule, sdk sql.GoSDKClient) error {
	ctx := context.Background()
	ruleName := instance.ObjectMeta.Name
	startIP := instance.Spec.StartIPAddress
	endIP := instance.Spec.EndIPAddress

	r.Telemetry.LogTrace(
		"Status",
		"Calling CreateOrUpdate Azure SQL firewall rule")

	//get owner instance of AzureSqlServer
	r.Recorder.Event(instance, v1.EventTypeNormal, "UpdatingOwner", "Updating owner AzureSqlServer instance")
	var ownerInstance azurev1alpha1.AzureSqlServer
	azureSQLServerNamespacedName := types.NamespacedName{Name: sdk.ServerName, Namespace: instance.Namespace}
	err := r.Get(ctx, azureSQLServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		msg := "Unable to get owner instance of AzureSqlServer"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
	} else {
		msg := "Got owner instance of Sql Server and assigning controller reference now"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeNormal, "OwnerAssign", msg)

		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			msg := "Unable to set controller reference to AzureSqlServer"
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		}
		successmsg := "Owner instance assigned successfully"
		instance.Status.Message = successmsg
		r.Recorder.Event(instance, v1.EventTypeNormal, "OwnerAssign", successmsg)
	}

	// write information back to instance
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	_, err = sdk.CreateOrUpdateSQLFirewallRule(ruleName, startIP, endIP)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Telemetry.LogInfo(
				"IgnorableError",
				"Async operation not complete or group not found")
			instance.Status.Provisioning = true
		}

		return errhelp.NewAzureError(err)
	}

	_, err = sdk.GetSQLFirewallRule(ruleName)
	if err != nil {
		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return nil
}

func (r *AzureSqlFirewallRuleReconciler) deleteExternal(instance *azurev1alpha1.AzureSqlFirewallRule, sdk sql.GoSDKClient) error {
	ruleName := instance.ObjectMeta.Name

	r.Telemetry.LogTrace(
		"Status",
		fmt.Sprintf("deleting external resource: group/%s/server/%s/firewallrule/%s", sdk.ResourceGroupName, sdk.ServerName, ruleName))
	err := sdk.DeleteSQLFirewallRule(ruleName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, v1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}
		msg := "Couldn't delete resouce in azure"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		return err
	}
	msg := fmt.Sprintf("Deleted %s", ruleName)
	instance.Status.Message = msg
	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", msg)

	return nil
}

func (r *AzureSqlFirewallRuleReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	helpers.AddFinalizer(instance, azureSQLFirewallRuleFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Failed to update finalizer: %v", err)
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", azureSQLFirewallRuleFinalizerName))
	return nil
}
