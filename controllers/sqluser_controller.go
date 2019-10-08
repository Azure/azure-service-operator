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
	"strings"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
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

// SQLUserFinalizerName is the name of the finalizer
const SQLUserFinalizerName = "sqluser.finalizers.azure.com"

// SqlUserReconciler reconciles a SqlUser object
type SqlUserReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlusers/status,verbs=get;update;patch

func (r *SqlUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqluser", req.NamespacedName)

	var instance azurev1.SqlUser

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sqluser resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, SQLUserFinalizerName) {
			if err := r.deleteExternal(instance); err != nil {
				log.Info("Delete Sqluser failed with ", "error", err.Error())
			}
			helpers.RemoveFinalizer(&instance, SQLUserFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, SQLUserFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding SqlUser finalizer failed with ", "error", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, "Normal", "Submitting", "starting resource reconciliation for SqlUser")
		if err := r.reconcileExternal(instance); err != nil {
			return ctrl.Result{}, fmt.Errorf("error reconciling sql user in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", fmt.Sprintf("sqluser %s provisioned", instance.ObjectMeta.Name))

	return ctrl.Result{}, nil
}

func (r *SqlUserReconciler) deleteExternal(instance azurev1.SqlUser) error {
	ctx := context.Background()

	// get admin credentials to connect to db
	secret := r.GetOrPrepareSecret(&instance, instance.Spec.AdminSecret)
	var user = string(secret.Data[SecretUsernameKey])
	var password = string(secret.Data[SecretPasswordKey])
	connString := r.getConnectionString(instance.Spec.Server, user, password, SqlServerPort, instance.Spec.DbName)

	db, _ := sql.Open(DriverName, connString)
	err := db.Ping()
	if err != nil {
		log.Error(err, "Unable to ping server")
	}
	err = dropUser(ctx, db, instance.ObjectMeta.Name)
	if err != nil {
		log.Error(err, fmt.Sprintf("Unable to drop user %s", instance.ObjectMeta.Name))
	}
	return nil
}

// Drops user from db
func dropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := fmt.Sprintf("DROP USER \"%s\"", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// Reconcile user sql request
func (r *SqlUserReconciler) reconcileExternal(instance azurev1.SqlUser) error {
	ctx := context.Background()

	// set DB as owner of user
	r.Recorder.Event(&instance, "Normal", "UpdatingOwner", "Updating owner SqlDatabase instance")
	var ownerInstance azurev1.SqlDatabase
	sqlDatabaseNamespaceName := types.NamespacedName{Name: instance.Spec.DbName, Namespace: instance.Namespace}
	err := r.Get(ctx, sqlDatabaseNamespaceName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(&instance, "Warning", "Failed", "Unable to get owner instance of SqlDatabase")
	} else {
		r.Recorder.Event(&instance, "Normal", "OwnerAssign", "Got owner instance of Sql Database and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, &instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(&instance, "Warning", "Failed", "Unable to set controller reference to SqlDatabase")
		}
		r.Recorder.Event(&instance, "Normal", "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if updateerr := r.Update(ctx, &instance); updateerr != nil {
		r.Recorder.Event(&instance, "Warning", "Failed", "Unable to update instance")
	}

	// get admin credentials to connect to db
	secret := r.GetOrPrepareSecret(&instance, instance.Spec.AdminSecret)
	var user = string(secret.Data[SecretUsernameKey])
	var password = string(secret.Data[SecretPasswordKey])
	connString := r.getConnectionString(instance.Spec.Server, user, password, SqlServerPort, instance.Spec.DbName)

	db, _ := sql.Open(DriverName, connString)
	err = db.Ping()
	if err != nil {
		log.Error(err, "Unable to ping server")
		return err
	}

	// create or get new user secret
	secret = r.GetOrPrepareSecret(&instance, instance.ObjectMeta.Name)
	containsUser, err := ContainsUser(ctx, db, instance.ObjectMeta.Name)
	if err != nil {
		log.Info("Error or couldn't find user", "err:", err.Error())
	}

	if !containsUser {
		user, err = createUser(ctx, secret, db)
		if err != nil {
			return nil
		}
	}

	// apply roles to user
	roles := instance.Spec.Roles
	if len(roles) == 0 {
		log.Info("No roles specified for user")
	} else {
		grantUserRoles(ctx, user, roles, db)
	}

	// publish user secret
	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
		r.Log.Info("mutating secret bundle")
		innerErr := controllerutil.SetControllerReference(&instance, secret, r.Scheme)
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

// Grants roles to a user for a given database
func grantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := fmt.Sprintf("sp_addrolemember \"%s\", \"%s\"", role, user)
		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			log.Info("Error executing add role", "err", err.Error())
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// Creates user with secret credentials
func createUser(ctx context.Context, secret *v1.Secret, db *sql.DB) (string, error) {
	newUser := string(secret.Data[SecretUsernameKey])
	newPassword := string(secret.Data[SecretPasswordKey])
	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)

	// Execute non-query with named parameters
	_, err := db.ExecContext(ctx, tsql)
	if err != nil {
		log.Error("Error executing", "err", err.Error())
		return newUser, err
	}
	return newUser, nil
}

// Builds connection string to connect to database
func (r *SqlUserReconciler) getConnectionString(server string, user string, password string, port int, database string) string {
	fullServerAddress := fmt.Sprintf("%s.database.windows.net", server)
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		fullServerAddress, user, password, port, database)
}

// ContainsUser checks if db contains user
func ContainsUser(ctx context.Context, db *sql.DB, username string) (bool, error) {
	tsql := fmt.Sprintf("SELECT * FROM sysusers WHERE NAME='%s'", username)
	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// GetOrPrepareSecret gets or creates a secret
func (r *SqlUserReconciler) GetOrPrepareSecret(instance *azurev1.SqlUser, name string) *v1.Secret {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: instance.Namespace,
		},
		Data: map[string][]byte{
			"username":           []byte(name),
			"password":           []byte(generateRandomString(16)),
			"sqlservernamespace": []byte(instance.Namespace),
			"sqlservername":      []byte(instance.Spec.Server),
		},
		Type: "Opaque",
	}

	if err := r.Get(context.Background(), types.NamespacedName{Name: name, Namespace: instance.Namespace}, secret); err == nil {
		r.Log.Info("secret already exists, pulling creds now")
	}

	return secret
}

// SetupWithManager runs reconcile loop with manager
func (r *SqlUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlUser{}).
		Complete(r)
}

// add finalizer
func (r *SqlUserReconciler) addFinalizer(instance *azurev1.SqlUser) error {
	helpers.AddFinalizer(instance, SQLUserFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", SQLUserFinalizerName))
	return nil
}
