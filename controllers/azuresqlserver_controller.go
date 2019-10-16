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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// AzureSqlServerReconciler reconciles an AzureSqlServer object
type AzureSqlServerReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
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

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, AzureSQLServerFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				if azerr, ok := err.(*errhelp.AzureError); ok {
					if helpers.ContainsString(catch, azerr.Type) {
						log.Info("Got ignorable error", "type", azerr.Type)
						return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
					}
				}
				msg := fmt.Sprintf("Delete AzureSqlServer failed with %s", err.Error())
				log.Info(msg)
				instance.Status.Message = msg
				if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
					r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
				}

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
			msg := fmt.Sprintf("Adding AzureSqlServer finalizer failed with error %s", err.Error())
			log.Info(msg)
			instance.Status.Message = msg
			if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
				r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
			}
			return ctrl.Result{}, err
		}
	}

	// Re-create secret if server is provisioned but secret doesn't exist
	if instance.IsProvisioned() {
		name := instance.ObjectMeta.Name

		secret := &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: instance.Namespace,
			},
			Type: "Opaque",
		}

		if err := r.Get(context.Background(), types.NamespacedName{Name: name, Namespace: instance.Namespace}, secret); err != nil {
			r.Log.Info("Error", "ReconcileSecret", "Server exists but secret does not, recreating now")

			// Add admin credentials to "data" block in secret

			// CreateOrUpdate secret
		}

	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		// TODO: Add error handling for cases where username or password are invalid:
		// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response
		if err := r.reconcileExternal(&instance); err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
				errhelp.InvalidServerName,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					if azerr.Type == errhelp.InvalidServerName {
						msg := "Invalid Server Name"
						r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", msg)
						instance.Status.Message = msg
						if updateerr := r.Status().Update(context.Background(), &instance); updateerr != nil {
							r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
						}
						return ctrl.Result{Requeue: false}, nil
					}
					msg := fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
					log.Info(msg)
					instance.Status.Message = msg
					if updateerr := r.Status().Update(context.Background(), &instance); updateerr != nil {
						r.Recorder.Event(&instance, "Warning", "Failed", "Unable to update instance")
					}
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql server in azure: %v", err)
		}
		// give azure some time to catch up
		log.Info("waiting for provision to take effect")
		return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	if err := r.verifyExternal(&instance); err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.AsyncOpIncompleteError,
		}
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				msg := fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
				log.Info(msg)
				instance.Status.Message = msg
				if updateerr := r.Status().Update(context.Background(), &instance); updateerr != nil {
					r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
				}
				return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
			}
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

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

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
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// Check to see if secret already exists for admin username/password
	secret, _ := r.GetOrPrepareSecret(instance)
	azureSqlServerProperties := sql.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(string(secret.Data["username"])),
		AdministratorLoginPassword: to.StringPtr(string(secret.Data["password"])),
	}

	// create the sql server
	instance.Status.Provisioning = true
	if _, err := sdkClient.CreateOrUpdateSQLServer(azureSqlServerProperties); err != nil {
		if !strings.Contains(err.Error(), "not complete") {
			msg := fmt.Sprintf("CreateOrUpdateSQLServer not complete: %v", err)
			instance.Status.Message = msg

			// write information back to instance
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
			}

			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return errhelp.NewAzureError(err)
		}
	} else {
		msg := "Resource request successfully submitted to Azure"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeNormal, "Provisioned", msg)

		// write information back to instance
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	}

	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
		r.Log.Info("Creating or updating secret with SQL Server credentials")
		innerErr := controllerutil.SetControllerReference(instance, secret, r.Scheme)
		if innerErr != nil {
			return innerErr
		}
		return nil
	})
	if createOrUpdateSecretErr != nil {
		return createOrUpdateSecretErr
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *AzureSqlServerReconciler) verifyExternal(instance *azurev1alpha1.AzureSqlServer) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	serv, err := sdkClient.GetServer()
	if err != nil {
		azerr := errhelp.NewAzureError(err).(*errhelp.AzureError)
		if azerr.Type != errhelp.ResourceNotFound {
			return azerr
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = *serv.State
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Checking", fmt.Sprintf("instance in %s state", instance.Status.State))

	if instance.Status.State == "Ready" {
		msg := "AzureSqlServer successfully provisioned"
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = msg
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		return updateerr
	}

	return errhelp.NewAzureError(err)
}

func (r *AzureSqlServerReconciler) deleteExternal(instance *azurev1alpha1.AzureSqlServer) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	location := instance.Spec.Location

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	_, err := sdkClient.DeleteSQLServer()
	if err != nil {
		msg := fmt.Sprintf("Couldn't delete resource in Azure: %v", err)
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return errhelp.NewAzureError(err)
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	return nil
}

func (r *AzureSqlServerReconciler) GetOrPrepareSecret(instance *azurev1alpha1.AzureSqlServer) (*v1.Secret, error) {
	name := instance.ObjectMeta.Name

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: instance.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{
			"username":                 []byte(""),
			"fullyqualifiedusername":   []byte(""),
			"password":                 []byte(""),
			"azuresqlservername":       []byte(""),
			"fullyqualifiedservername": []byte(""),
		},
		Type: "Opaque",
	}

	randomUsername, usernameErr := generateRandomUsername(usernameLength)
	if usernameErr != nil {
		return secret, usernameErr
	}

	randomPassword, passwordErr := generateRandomPassword(passwordLength)
	if passwordErr != nil {
		return secret, passwordErr
	}

	usernameSuffix := "@" + name
	servernameSuffix := ".database.windows.net"
	fullyQualifiedAdminUsername := randomUsername + usernameSuffix // "<username>@<azuresqlservername>""
	fullyQualifiedServername := name + servernameSuffix            // "<azuresqlservername>.database.windows.net"

	secret.Data["username"] = []byte(randomUsername)
	secret.Data["fullyqualifiedusername"] = []byte(fullyQualifiedAdminUsername)
	secret.Data["password"] = []byte(randomPassword)
	secret.Data["azuresqlservername"] = []byte(name)
	secret.Data["fullyqualifiedservername"] = []byte(fullyQualifiedServername)

	if err := r.Get(context.Background(), types.NamespacedName{Name: name, Namespace: instance.Namespace}, secret); err == nil {
		r.Log.Info("secret already exists, pulling creds now")
	}

	return secret, nil
}

// helper function to generate random username for sql server
func generateRandomUsername(n int) (string, error) {
	if n < minUsernameAllowedLength || n > maxUsernameAllowedLength {
		return "", errors.New("Username length should be between 8 and 63 characters.")
	}

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
	if n < minPasswordAllowedLength || n > maxPasswordAllowedLength {
		return "", errors.New("Password length must be between 8 and 128 characters.")
	}

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
