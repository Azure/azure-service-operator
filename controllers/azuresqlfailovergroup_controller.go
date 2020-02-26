/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.

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

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	azuresqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

const azureSQLFailoverGroupFinalizerName = finalizerName

// AzureSqlFailoverGroupReconciler reconciles a AzureSqlFailoverGroup object
type AzureSqlFailoverGroupReconciler struct {
	client.Client
	Log                          logr.Logger
	Recorder                     record.EventRecorder
	Scheme                       *runtime.Scheme
	AzureSqlFailoverGroupManager azuresqlfailovergroup.SqlFailoverGroupManager
	SecretClient                 secrets.SecretClient
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfailovergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfailovergroups/status,verbs=get;update;patch

func (r *AzureSqlFailoverGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("AzureSqlFailoverGroup", req.NamespacedName)
	var instance azurev1alpha1.AzureSqlFailoverGroup

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql failover group resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	defer func() {
		if !helpers.IsBeingDeleted(&instance) {
			if err := r.Status().Update(ctx, &instance); err != nil {
				r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
			}
		}
	}()

	if helpers.IsBeingDeleted(&instance) {
		if HasFinalizer(&instance, azureSQLFailoverGroupFinalizerName) {
			if err := r.deleteExternal(ctx, &instance); err != nil {
				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				if azerr, ok := err.(*errhelp.AzureError); ok {
					if helpers.ContainsString(catch, azerr.Type) {
						log.Info("Got ignorable error", "type", azerr.Type)
						return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
					}
				}
				msg := fmt.Sprintf("Delete AzureFailoverGroup failed with %s", err.Error())
				log.Info(msg)
				instance.Status.Message = msg
				return ctrl.Result{}, err
			}

			RemoveFinalizer(&instance, azureSQLFailoverGroupFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !HasFinalizer(&instance, azureSQLFailoverGroupFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			msg := fmt.Sprintf("Adding AzureFailoverGroup finalizer failed with error %s", err.Error())
			log.Info(msg)
			instance.Status.Message = msg
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		if err := r.reconcileExternal(ctx, &instance); err != nil {
			instance.Status.Message = err.Error()
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
				errhelp.InvalidServerName,
				errhelp.ResourceNotFound,
				errhelp.FailoverGroupBusy,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					msg := fmt.Sprintf("Got ignorable error of type %v", azerr.Type)
					log.Info(msg)
					instance.Status.Message = msg

					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql failover group in azure: %v", err)
		}

		return ctrl.Result{}, nil
	}

	_, err := r.AzureSqlFailoverGroupManager.GetFailoverGroup(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.GetName())
	if err != nil {
		if azerr := errhelp.NewAzureErrorAzureError(err); azerr.Type == errhelp.ResourceNotFound {
			log.Info("waiting for failovergroup to come up")
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}
		return ctrl.Result{}, err
	}

	// Check to see if secret already exists for server names and paths
	r.Log.Info("retrieving or generating secret values")
	secret, _ := r.GetOrPrepareSecret(ctx, &instance)

	// create or update the secret
	r.Log.Info("persisting failovergroup secrets")
	key := types.NamespacedName{Name: instance.ObjectMeta.Name, Namespace: instance.Namespace}
	err = r.SecretClient.Upsert(
		ctx,
		key,
		secret,
		secrets.WithOwner(&instance),
		secrets.WithScheme(r.Scheme),
	)
	if err != nil {
		return ctrl.Result{}, err
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "AzureSqlFailoverGroup "+instance.ObjectMeta.Name+" provisioned ")
	instance.Status.Message = successMsg
	instance.Status.State = "done"
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return ctrl.Result{}, nil
}

func (r *AzureSqlFailoverGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlFailoverGroup{}).
		Complete(r)
}

func (r *AzureSqlFailoverGroupReconciler) reconcileExternal(ctx context.Context, instance *azurev1alpha1.AzureSqlFailoverGroup) error {
	failoverGroupName := instance.ObjectMeta.Name
	failoverPolicy := instance.Spec.FailoverPolicy
	failoverGracePeriod := instance.Spec.FailoverGracePeriod
	secondaryServer := instance.Spec.SecondaryServerName
	secondaryResourceGroup := instance.Spec.SecondaryServerResourceGroup
	databaseList := instance.Spec.DatabaseList
	server := instance.Spec.Server
	groupName := instance.Spec.ResourceGroup
	servername := instance.Spec.Server

	//get owner instance of AzureSqlServer
	r.Recorder.Event(instance, v1.EventTypeNormal, "UpdatingOwner", "Updating owner AzureSqlServer instance")
	var ownerInstance azurev1alpha1.AzureSqlServer

	azureSQLServerNamespacedName := types.NamespacedName{Name: server, Namespace: instance.Namespace}
	err := r.Get(ctx, azureSQLServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		msg := "Unable to get owner instance of AzureSqlServer"
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
	} else {
		msg := "Got owner instance of Sql Server and assigning controller reference now"
		r.Recorder.Event(instance, v1.EventTypeNormal, "OwnerAssign", msg)

		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			msg := "Unable to set controller reference to AzureSqlServer"
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
			instance.Status.Message = msg
		}
		successmsg := "Owner instance assigned successfully"
		r.Recorder.Event(instance, v1.EventTypeNormal, "OwnerAssign", successmsg)
		instance.Status.Message = successmsg
	}

	// write information back to instance
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// Create Failover Group properties struct
	sqlFailoverGroupProperties := azuresqlshared.SQLFailoverGroupProperties{
		FailoverPolicy:               failoverPolicy,
		FailoverGracePeriod:          failoverGracePeriod,
		SecondaryServerName:          secondaryServer,
		SecondaryServerResourceGroup: secondaryResourceGroup,
		DatabaseList:                 databaseList,
	}

	_, err = r.AzureSqlFailoverGroupManager.CreateOrUpdateFailoverGroup(ctx, groupName, servername, failoverGroupName, sqlFailoverGroupProperties)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Log.Info("Async operation not complete or group not found")
			instance.Status.Provisioning = true
		}

		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = true

	_, err = r.AzureSqlFailoverGroupManager.GetFailoverGroup(ctx, groupName, servername, failoverGroupName)
	if err != nil {
		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = "Provisioned failover group"

	return nil
}

func (r *AzureSqlFailoverGroupReconciler) deleteExternal(ctx context.Context, instance *azurev1alpha1.AzureSqlFailoverGroup) error {
	name := instance.ObjectMeta.Name
	servername := instance.Spec.Server
	groupName := instance.Spec.ResourceGroup

	response, err := r.AzureSqlFailoverGroupManager.DeleteFailoverGroup(ctx, groupName, servername, name)
	if err != nil {
		msg := fmt.Sprintf("Couldn't delete resource in Azure: %v", err)
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		return errhelp.NewAzureError(err)
	}
	if response.StatusCode == 200 {
		r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	}

	return nil
}

func (r *AzureSqlFailoverGroupReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlFailoverGroup) error {
	AddFinalizer(instance, azureSQLFailoverGroupFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", azureSQLFailoverGroupFinalizerName))
	return nil
}

func (r *AzureSqlFailoverGroupReconciler) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureSqlFailoverGroup) (map[string][]byte, error) {
	failovergroupname := instance.ObjectMeta.Name
	azuresqlprimaryservername := instance.Spec.Server
	azuresqlsecondaryservername := instance.Spec.SecondaryServerName

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: failovergroupname, Namespace: instance.Namespace}

	if stored, err := r.SecretClient.Get(ctx, key); err == nil {
		r.Log.Info("secret already exists, pulling stored values")
		return stored, nil
	}

	r.Log.Info("secret not found, generating values for new secret")

	secret["azureSqlPrimaryServerName"] = []byte(azuresqlprimaryservername)
	secret["readWriteListenerEndpoint"] = []byte(failovergroupname + ".database.windows.net")
	secret["azureSqlSecondaryServerName"] = []byte(azuresqlsecondaryservername)
	secret["readOnlyListenerEndpoint"] = []byte(failovergroupname + ".secondary.database.windows.net")

	return secret, nil
}
