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
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const SQLFirewallRuleFinalizerName = "sqlfirewallrule.finalizers.azure.com"

// SqlFirewallRuleReconciler reconciles a SqlFirewallRule object
type SqlFirewallRuleReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlfirewallrules/status,verbs=get;update;patch

func (r *SqlFirewallRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlfirewallrule", req.NamespacedName)

	// your logic here
	var instance azurev1alpha1.SqlFirewallRule

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql-firewall-rule resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, SQLFirewallRuleFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete SqlFirewallRule failed with ", "error", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, SQLFirewallRuleFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, SQLFirewallRuleFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding SqlFirewallRule finalizer failed with ", "error", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, "Normal", "Submitting", "starting resource reconciliation for SqlFirewallRule")
		if err := r.reconcileExternal(&instance); err != nil {

			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql firewall rule in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "sqlfirewallrule "+instance.ObjectMeta.Name+" provisioned ")

	return ctrl.Result{}, nil
}

func (r *SqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.SqlFirewallRule{}).
		Complete(r)
}

func (r *SqlFirewallRuleReconciler) reconcileExternal(instance *azurev1alpha1.SqlFirewallRule) error {
	ctx := context.Background()
	ruleName := instance.ObjectMeta.Name
	server := instance.Spec.Server
	startIP := instance.Spec.StartIPAddress
	endIP := instance.Spec.EndIPAddress

	r.Log.Info("Calling createorupdate SQL firewall rule")

	// get owner instance of SqlServer
	r.Recorder.Event(instance, "Normal", "UpdatingOwner", "Updating owner SqlServer instance")
	var ownerInstance azurev1alpha1.SqlServer
	sqlServerNamespacedName := types.NamespacedName{Name: server, Namespace: instance.Namespace}
	err := r.Get(ctx, sqlServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to get owner instance of SqlServer when setting firewall parent")
		return err
	}

	// set owner instance to the SqlServer instance
	r.Recorder.Event(instance, "Normal", "OwnerAssign", "Got owner instance of Sql Server and assigning controller reference now")
	err = controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to set controller reference to SqlServer")
		return err
	}
	r.Recorder.Event(instance, "Normal", "OwnerAssign", "Owner instance assigned successfully")

	// write information back to instance
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	// actually add the firewall rule
	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: ownerInstance.Spec.ResourceGroup,
		ServerName:        server,
	}
	_, err = sdkClient.CreateOrUpdateSQLFirewallRule(ruleName, startIP, endIP)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Log.Info("Async operation not complete or group not found")
			instance.Status.Provisioning = true
			if errup := r.Status().Update(ctx, instance); errup != nil {
				r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
			}
		}

		return errhelp.NewAzureError(err)
	}

	_, err = sdkClient.GetSQLFirewallRule(ruleName)
	if err != nil {
		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if err = r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	return nil
}

func (r *SqlFirewallRuleReconciler) deleteExternal(instance *azurev1alpha1.SqlFirewallRule) error {
	ctx := context.Background()
	ruleName := instance.ObjectMeta.Name
	server := instance.Spec.Server

	//get owner instance of SqlServer
	var ownerInstance azurev1alpha1.SqlServer
	sqlServerNamespacedName := types.NamespacedName{Name: server, Namespace: instance.Namespace}
	err := r.Get(ctx, sqlServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, "Warning", "DoesNotExist", "Unable to get owner instance of SqlServer when removing firewall rule (OK)")
		return nil
	}
	groupName := ownerInstance.Spec.ResourceGroup

	// create the Go SDK client with relevant info
	sdk := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
	}

	r.Log.Info(fmt.Sprintf("deleting external resource: group/%s/server/%s/firewallrule/%s"+groupName, server, ruleName))
	err = sdk.DeleteSQLFirewallRule(ruleName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}
	r.Recorder.Event(instance, "Normal", "Deleted", ruleName+" deleted")
	return nil
}

func (r *SqlFirewallRuleReconciler) addFinalizer(instance *azurev1alpha1.SqlFirewallRule) error {
	helpers.AddFinalizer(instance, SQLFirewallRuleFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", SQLFirewallRuleFinalizerName))
	return nil
}
