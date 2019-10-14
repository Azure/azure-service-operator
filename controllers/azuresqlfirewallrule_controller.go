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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const azureSQLFirewallRuleFinalizerName = "azuresqlfirewallrule.finalizers.azure.com"

// AzureSqlFirewallRuleReconciler reconciles a AzureSqlFirewallRule object
type AzureSqlFirewallRuleReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules/status,verbs=get;update;patch

func (r *AzureSqlFirewallRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("azuresqlfirewallrule", req.NamespacedName)

	// your logic here
	var instance azurev1alpha1.AzureSqlFirewallRule

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve azure-sql-firewall-rule resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, azureSQLFirewallRuleFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete AzureSqlFirewallRule failed with ", "error", err.Error())
				instance.Status.Message = fmt.Sprintf("Delete AzureSqlFirewallRule failed with %s", err.Error())
				if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
					r.Recorder.Event(&instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
				}
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
			log.Info("Adding AzureSqlFirewallRule finalizer failed with ", "error", err.Error())
			instance.Status.Message = fmt.Sprintf("Adding AzureSqlFirewallRule finalizer failed with error %s", err.Error())
			if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
				r.Recorder.Event(&instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, corev1.EventTypeNormal, "Submitting", "starting resource reconciliation for AzureSqlFirewallRule")
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
					instance.Status.Message = fmt.Sprintf("Got ignorable error of type %s", azerr.Type)
					if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
						r.Recorder.Event(&instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
					}
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling azure sql firewall rule in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, corev1.EventTypeNormal, "Provisioned", "azuresqlfirewallrule "+instance.ObjectMeta.Name+" provisioned ")
	instance.Status.Message = fmt.Sprintf("AzureSqlFirewallrule%s successfully provisioned", instance.ObjectMeta.Name)
	if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
		r.Recorder.Event(&instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return ctrl.Result{}, nil
}

func (r *AzureSqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlFirewallRule{}).
		Complete(r)
}

func (r *AzureSqlFirewallRuleReconciler) reconcileExternal(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name
	startIP := instance.Spec.StartIPAddress
	endIP := instance.Spec.EndIPAddress

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
	}

	r.Log.Info("Calling createorupdate Azure SQL firewall rule")

	//get owner instance of AzureSqlServer
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner AzureSqlServer instance")
	var ownerInstance azurev1alpha1.AzureSqlServer
	azureSqlServerNamespacedName := types.NamespacedName{Name: server, Namespace: instance.Namespace}
	err := r.Get(ctx, azureSqlServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get owner instance of AzureSqlServer")
		instance.Status.Message = "Unable to get owner instance of AzureSqlServer"
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Got owner instance of Sql Server and assigning controller reference now")
		instance.Status.Message = "Grabbed Owner instance. Assigning controller Reference."
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to set controller reference to AzureSqlServer")
			instance.Status.Message = "Unable to set controller reference to AzureSqlServer"
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
		}
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
		instance.Status.Message = "Owner instance assigned successfully"
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	}

	// write information back to instance
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	_, err = sdkClient.CreateOrUpdateSQLFirewallRule(ruleName, startIP, endIP)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Log.Info("Async operation not complete or group not found")
			instance.Status.Message = "Async operation not complete or group not found"
			instance.Status.Provisioning = true
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
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
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *AzureSqlFirewallRuleReconciler) deleteExternal(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name

	// create the Go SDK client with relevant info
	sdk := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
	}

	r.Log.Info(fmt.Sprintf("deleting external resource: group/%s/server/%s/firewallrule/%s"+groupName, server, ruleName))
	err := sdk.DeleteSQLFirewallRule(ruleName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't delete resouce in azure")
		instance.Status.Message = "Couldn't delete resource in Azure"
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return err
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Deleted", ruleName+" deleted")
	instance.Status.Message = fmt.Sprintf("Deleted %s", ruleName)
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}
	return nil
}

func (r *AzureSqlFirewallRuleReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	helpers.AddFinalizer(instance, azureSQLFirewallRuleFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Failed to update finalizer: %v", err)
		if updateerr := r.Status().Update(context.Background(), instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", azureSQLFirewallRuleFinalizerName))
	return nil
}
