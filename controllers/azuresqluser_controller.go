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
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	_ "github.com/denisenkom/go-mssqldb"
)

// SqlServerPort is the default server port for sql server
const SqlServerPort = 1433

// DriverName is driver name for db connection
const DriverName = "sqlserver"

// SecretUsernameKey is the username key in secret
const SecretUsernameKey = "username"

// SecretPasswordKey is the password key in secret
const SecretPasswordKey = "password"

// AzureSQLUserFinalizerName is the name of the finalizer
const AzureSQLUserFinalizerName = "azuresqluser.finalizers.azure.com"

// AzureAzureSQLUserReconciler reconciles a AzureSQLUser object
type AzureSQLUserReconciler struct {
	client.Client
	Log                 logr.Logger
	Recorder            record.EventRecorder
	Scheme              *runtime.Scheme
	AzureSqlUserManager sqlclient.SqlUserManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=AzureSQLUsers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=AzureSQLUsers/status,verbs=get;update;patch

func (r *AzureSQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("AzureSQLUser", req.NamespacedName)

	var instance azurev1alpha1.AzureSQLUser

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve AzureSQLUser resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	defer func() {
		if err := r.Status().Update(ctx, &instance); err != nil {
			r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	}()

	// if the admin credentials haven't been set, default admin credentials to servername
	if instance.Spec.AdminSecret == "" {
		instance.Spec.AdminSecret = instance.Spec.Server
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, AzureSQLUserFinalizerName) {
			if err := r.deleteExternal(instance); err != nil {
				instance.Status.Message = fmt.Sprintf("Delete external failed with %s", err.Error())
				log.Info(instance.Status.Message)
			}
			helpers.RemoveFinalizer(&instance, AzureSQLUserFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, AzureSQLUserFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding AzureSQLUser finalizer failed with ", "error", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation for AzureSQLUser")
		if err := r.reconcileExternal(instance); err != nil {
			msg := fmt.Sprintf("Reconcile external failed with %s", err.Error())
			log.Info(msg)
			instance.Status.Message = msg
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}
		instance.Status.Message = fmt.Sprintf("Create AzureSqlUser Successful")
		log.Info(instance.Status.Message)
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", fmt.Sprintf("AzureSQLUser %s provisioned", instance.ObjectMeta.Name))
	return ctrl.Result{}, nil
}

func (r *AzureSQLUserReconciler) deleteExternal(instance azurev1alpha1.AzureSQLUser) error {
	ctx := context.Background()

	// get admin credentials to connect to db
	adminSecret := &v1.Secret{}
	if r.SecretExists(&instance, instance.Spec.AdminSecret) {
		adminSecret = r.GetOrPrepareSecret(&instance, instance.Spec.AdminSecret, instance.Spec.AdminSecret)
	} else {
		return fmt.Errorf("admin secret : %s, not found", instance.Spec.AdminSecret)
	}

	var user = string(adminSecret.Data[SecretUsernameKey])
	var password = string(adminSecret.Data[SecretPasswordKey])
	connString := r.getConnectionString(instance.Spec.Server, user, password, SqlServerPort, instance.Spec.DbName)

	db, err := sql.Open(DriverName, connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DBSecret := r.GetOrPrepareSecret(&instance, instance.ObjectMeta.Name, instance.ObjectMeta.Name)
	user = string(DBSecret.Data[SecretUsernameKey])

	err = r.AzureSqlUserManager.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete AzureSqlUser failed with %s", err.Error())
		log.Info(instance.Status.Message)
		return err
	}

	instance.Status.Message = fmt.Sprintf("Delete AzureSqlUser succeeded")
	log.Info(instance.Status.Message)
	return nil
}

// Reconcile user sql request
func (r *AzureSQLUserReconciler) reconcileExternal(instance azurev1alpha1.AzureSQLUser) error {
	ctx := context.Background()

	// set DB as owner of user
	r.Recorder.Event(&instance, v1.EventTypeNormal, "UpdatingOwner", "Updating owner SqlDatabase instance")
	var ownerInstance azurev1alpha1.AzureSqlDatabase
	sqlDatabaseNamespaceName := types.NamespacedName{Name: instance.Spec.DbName, Namespace: instance.Namespace}
	err := r.Get(ctx, sqlDatabaseNamespaceName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to get owner instance of SqlDatabase")
	} else {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "OwnerAssign", "Got owner instance of Sql Database and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, &instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to set controller reference to SqlDatabase")
		}
		r.Recorder.Event(&instance, v1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if updateerr := r.Update(ctx, &instance); updateerr != nil {
		r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// get admin credentials to connect to db
	adminSecret := &v1.Secret{}
	if r.SecretExists(&instance, instance.Spec.AdminSecret) {
		adminSecret = r.GetOrPrepareSecret(&instance, instance.Spec.AdminSecret, instance.Spec.AdminSecret)
	} else {
		return fmt.Errorf("admin secret : %s, not found", instance.Spec.AdminSecret)
	}

	var user = string(adminSecret.Data[SecretUsernameKey])
	var password = string(adminSecret.Data[SecretPasswordKey])
	connString := r.getConnectionString(instance.Spec.Server, user, password, SqlServerPort, instance.Spec.DbName)

	db, err := sql.Open(DriverName, connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	// create or get new user secret
	user = fmt.Sprintf("%s-%s", instance.ObjectMeta.Name, uuid.New())
	DBSecret := r.GetOrPrepareSecret(&instance, instance.ObjectMeta.Name, user)
	containsUser, err := r.AzureSqlUserManager.ContainsUser(ctx, db, string(DBSecret.Data[SecretUsernameKey]))
	if err != nil {
		log.Info("Couldn't find user", "err:", err.Error())
	}

	if !containsUser {
		user, err = r.AzureSqlUserManager.CreateUser(ctx, DBSecret, db)
		if err != nil {
			return err
		}
	}

	// apply roles to user
	roles := instance.Spec.Roles
	if len(roles) == 0 {
		log.Info("No roles specified for user")
	} else {
		r.AzureSqlUserManager.GrantUserRoles(ctx, user, roles, db)
	}

	// publish user secret
	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, DBSecret, func() error {
		r.Log.Info("mutating secret bundle")
		innerErr := controllerutil.SetControllerReference(&instance, DBSecret, r.Scheme)
		if innerErr != nil {
			log.Info("inner err", "err", err.Error())
		}
		return nil
	})
	if createOrUpdateSecretErr != nil {
		log.Info("createOrUpdateSecretErr", "err", createOrUpdateSecretErr.Error())
		return err
	}
	return nil
}

// SetupWithManager runs reconcile loop with manager
func (r *AzureSQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSQLUser{}).
		Complete(r)
}

// add finalizer
func (r *AzureSQLUserReconciler) addFinalizer(instance *azurev1alpha1.AzureSQLUser) error {
	helpers.AddFinalizer(instance, AzureSQLUserFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", AzureSQLUserFinalizerName))
	return nil
}

// GetOrPrepareSecret gets or creates a secret
func (r *AzureSQLUserReconciler) GetOrPrepareSecret(instance *azurev1alpha1.AzureSQLUser, secretname string, username string) *v1.Secret {
	pw, _ := generateRandomPassword(16)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretname,
			Namespace: instance.Namespace,
		},
		Data: map[string][]byte{
			"username":           []byte(username),
			"password":           []byte(pw),
			"sqlservernamespace": []byte(instance.Namespace),
			"sqlservername":      []byte(instance.Spec.Server),
		},
		Type: "Opaque",
	}

	if err := r.Get(context.Background(), types.NamespacedName{Name: secretname, Namespace: instance.Namespace}, secret); err == nil {
		r.Log.Info("secret already exists, pulling creds now")
	}

	return secret
}

// Checks if secret exists
func (r *AzureSQLUserReconciler) SecretExists(instance *azurev1alpha1.AzureSQLUser, secretname string) bool {
	secret := &v1.Secret{}
	if err := r.Get(context.Background(), types.NamespacedName{Name: secretname, Namespace: instance.Namespace}, secret); err == nil {
		return true
	}

	return false
}

// Builds connection string to connect to database
func (r *AzureSQLUserReconciler) getConnectionString(server string, user string, password string, port int, database string) string {
	fullServerAddress := fmt.Sprintf("%s.database.windows.net", server)
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		fullServerAddress, user, password, port, database)
}
