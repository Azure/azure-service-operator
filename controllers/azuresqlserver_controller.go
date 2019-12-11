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
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/sethvargo/go-password/password"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// AzureSqlServerReconciler reconciles an AzureSqlServer object
type AzureSqlServerReconciler struct {
	client.Client
	Log                   logr.Logger
	Recorder              record.EventRecorder
	Scheme                *runtime.Scheme
	AzureSqlServerManager sql.SqlServerManager
	SecretClient          secrets.SecretClient
}

// Constants
const usernameLength = 8
const passwordLength = 16
const minUsernameAllowedLength = 8
const maxUsernameAllowedLength = 63
const minPasswordAllowedLength = 8
const maxPasswordAllowedLength = 128

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlservers/status,verbs=get;update;patch

func (r *AzureSqlServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("azuresqlserver", req.NamespacedName)
	var instance azurev1alpha1.AzureSqlServer

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql-server resource", "err", err.Error())
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
		if helpers.HasFinalizer(&instance, AzureSQLServerFinalizerName) {
			if err := r.deleteExternal(ctx, &instance); err != nil {
				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				azerr := errhelp.NewAzureErrorAzureError(err)
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}

				instance.Status.Message = fmt.Sprintf("Delete AzureSqlServer failed with %v", err)
				log.Info(instance.Status.Message)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, AzureSQLServerFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, AzureSQLServerFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			instance.Status.Message = fmt.Sprintf("Adding AzureSqlServer finalizer failed with error %v", err)
			log.Info(instance.Status.Message)
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		if err := r.reconcileExternal(&instance); err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
				errhelp.InvalidServerName,
				errhelp.AlreadyExists,
			}
			azerr := errhelp.NewAzureErrorAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {

				if azerr.Type == errhelp.AlreadyExists {
					// This error could happen in two cases - when the SQL server
					// exists in some other resource group or when this is a repeat
					// call to the reconcile loop for an update of the resource. So
					// we call a Get to check if this is the current resource and if
					// yes, we let the call go through instead of ending the reconcile loop
					_, err := r.ResourceClient.GetServer(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
					if err != nil {
						// This means that the Server exists elsewhere and we should
						// terminate the reconcile loop
						instance.Status.Message = "Server Already exists"
						instance.Status.Provisioning = false
						r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", instance.Status.Message)
						return ctrl.Result{Requeue: false}, nil
					}
				}

				if azerr.Type == errhelp.InvalidServerName {
					instance.Status.Message = "Invalid Server Name"
					r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", instance.Status.Message)
					return ctrl.Result{Requeue: false}, nil
				}
				instance.Status.Message = fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
				log.Info(instance.Status.Message)
				return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
			}

			return ctrl.Result{}, fmt.Errorf("error reconciling sql server in azure: %v", err)
		}
		// give azure some time to catch up
		log.Info("waiting for provision to take effect")
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	if err := r.verifyExternal(ctx, &instance); err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
			log.Info(instance.Status.Message)
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}

		return ctrl.Result{}, fmt.Errorf("error verifying sql server in azure: %v", err)
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "azuresqlserver "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *AzureSqlServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlServer{}).
		Complete(r)
}

func (r *AzureSqlServerReconciler) reconcileExternal(instance *azurev1alpha1.AzureSqlServer) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	//get owner instance of ResourceGroup
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner ResourceGroup instance")
	var ownerInstance azurev1alpha1.ResourceGroup

	// Get resource group
	resourceGroupNamespacedName := types.NamespacedName{Name: groupName, Namespace: instance.Namespace}
	err := r.Get(ctx, resourceGroupNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get owner instance of ReourceGroup")
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Got owner instance of Resource Group and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to set controller reference to ResourceGroup")
		}
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if err := r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// Check to see if secret already exists for admin username/password
	secret, _ := r.GetOrPrepareSecret(ctx, instance)
	azureSqlServerProperties := sql.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(string(secret["username"])),
		AdministratorLoginPassword: to.StringPtr(string(secret["password"])),
	}

	// create the sql server
	instance.Status.Provisioning = true
	if _, err := r.AzureSqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, location, name, azureSqlServerProperties); err != nil {
		if !strings.Contains(err.Error(), "not complete") {
			msg := fmt.Sprintf("CreateOrUpdateSQLServer not complete: %v", err)
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return err
		}
		if strings.Contains(err.Error(), errhelp.InvalidServerName) {
			msg := fmt.Sprintf("Invalid Server Name: %v", err)
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return errhelp.NewAzureError(err)
		}
	} else {
		msg := "Resource request successfully submitted to Azure"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeNormal, "Provisioned", msg)
	}

	// create or update the secret
	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	err = r.SecretClient.Upsert(
		ctx,
		key,
		secret,
		secrets.WithOwner(instance),
		secrets.WithScheme(r.Scheme),
	)
	if err != nil {
		return err
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *AzureSqlServerReconciler) verifyExternal(ctx context.Context, instance *azurev1alpha1.AzureSqlServer) error {
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	serv, err := r.AzureSqlServerManager.GetServer(ctx, groupName, name)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound {
			return azerr
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = *serv.State
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Checking", fmt.Sprintf("instance in %s state", instance.Status.State))

	if instance.Status.State == "Ready" {
		instance.Status.Message = "AzureSqlServer successfully provisioned"
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
	}

	return err
}

func (r *AzureSqlServerReconciler) deleteExternal(ctx context.Context, instance *azurev1alpha1.AzureSqlServer) error {
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	_, err := r.AzureSqlServerManager.DeleteSQLServer(ctx, groupName, name)
	if err != nil {
		msg := fmt.Sprintf("Couldn't delete resource in Azure: %v", err)
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		return err
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	return nil
}

func (r *AzureSqlServerReconciler) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureSqlServer) (map[string][]byte, error) {
	name := instance.ObjectMeta.Name

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := r.SecretClient.Get(ctx, key); err == nil {
		r.Log.Info("secret already exists, pulling creds now")
		return stored, nil
	}

	r.Log.Info("secret not found, generating values for new secret")

	randomUsername, err := generateRandomUsername(usernameLength)
	if err != nil {
		return secret, err
	}

	randomPassword, err := generateRandomPassword(passwordLength)
	if err != nil {
		return secret, err
	}

	secret["username"] = []byte(randomUsername)
	secret["fullyqualifiedusername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, name))
	secret["password"] = []byte(randomPassword)
	secret["azuresqlservername"] = []byte(name)
	secret["fullyqualifiedservername"] = []byte(name + ".database.windows.net")

	return secret, nil
}

// helper function to generate random username for sql server
func generateRandomUsername(n int) (string, error) {

	// Generate a username that is n characters long, with n/2 digits and 0 symbols (not allowed),
	// allowing only lower case letters (upper case not allowed), and disallowing repeat characters.
	res, err := password.Generate(n, (n / 2), 0, true, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

// helper function to generate random password for sql server
func generateRandomPassword(n int) (string, error) {

	// Math - Generate a password where: 1/3 of the # of chars are digits, 1/3 of the # of chars are symbols,
	// and the remaining 1/3 is a mix of upper- and lower-case letters
	digits := n / 3
	symbols := n / 3

	// Generate a password that is n characters long, with # of digits and symbols described above,
	// allowing upper and lower case letters, and disallowing repeat characters.
	res, err := password.Generate(n, digits, symbols, false, false)
	if err != nil {
		return "", err
	}

	return res, nil
}
