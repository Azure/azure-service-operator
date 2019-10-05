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
	"time"

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
	"github.com/Azure/azure-service-operator/pkg/errhelp"
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
		if helpers.HasFinalizer(&instance, SQLFirewallRuleFinalizerName) {
			if err := r.deleteExternal(instance); err != nil {
				log.Info("Delete Sqluser failed with ", "error", err.Error())
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
			log.Info("Adding SqlUser finalizer failed with ", "error", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, "Normal", "Submitting", "starting resource reconciliation for SqlUser")
		if err := r.reconcileExternal(instance); err != nil {

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
			return ctrl.Result{}, fmt.Errorf("error reconciling sql user in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "sqluser "+instance.ObjectMeta.Name+" provisioned ")

	return ctrl.Result{}, nil
}

func (r *SqlUserReconciler) deleteExternal(instance azurev1.SqlUser) error {
	// Add DB call to drop user and delete secret
	return nil
}

// Reconcile user sql request
func (r *SqlUserReconciler) reconcileExternal(instance azurev1.SqlUser) error {
	ctx := context.Background()

	// get admin credentials to connect to db
	secret := r.GetOrPrepareSecret(&instance, instance.Spec.Server)
	var user = string(secret.Data[SecretUsernameKey])
	var password = string(secret.Data[SecretPasswordKey])
	connString := r.getConnectionString(instance.Spec.Server, user, password, SqlServerPort, instance.Spec.DbName)

	db, _ := sql.Open(DriverName, connString)
	err := db.Ping()
	if err != nil {
		log.Error(err, "Unable to ping server")
		return err
	}

	// create or get new user secret
	secret = r.GetOrPrepareSecret(&instance, instance.ObjectMeta.Name)
	user, err = createUser(ctx, secret, db)
	if err != nil {
		log.Error(err, "Unable to create user")
		return err
	}

	// apply roles to user
	roles := instance.Spec.Roles
	if len(roles) == 0 {
		log.Info("No roles specified for user")
	} else {
		grantUserRoles(ctx, user, roles, db)
	}

	// create and publish secret
	secret = r.GetOrPrepareSecret(&instance, user)

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
		tsql := fmt.Sprintf("sp_addrolemember '%s', %s", role, user)
		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			log.Info("Error executing", "err", err.Error())
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

	log.Info("User created", "username", newUser)
	return newUser, nil
}

// Builds connection string to connect to database
func (r *SqlUserReconciler) getConnectionString(server string, user string, password string, port int, database string) string {
	fullServerAddress := fmt.Sprintf("%s.database.windows.net", server)
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		fullServerAddress, user, password, port, database)
}

// GetOrPrepareSecret gets or creates a secret
func (r *SqlUserReconciler) GetOrPrepareSecret(instance *azurev1.SqlUser, name string) *v1.Secret {

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: instance.Namespace,
		},
		Data: map[string][]byte{
			"username":           []byte(generateRandomString(8)),
			"password":           []byte(generateRandomString(16)),
			"sqlservernamespace": []byte(instance.Namespace),
			"sqlservername":      []byte(name),
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
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", SQLFirewallRuleFinalizerName))
	return nil
}
