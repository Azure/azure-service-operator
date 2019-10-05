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

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"

	//sqlclient "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	_ "github.com/denisenkom/go-mssqldb"
)

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
	var instance azurev1.SqlUser
	log := r.Log.WithValues("sqluser", req.NamespacedName)
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql user resource", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	secret := r.GetOrPrepareSecret(&instance, instance.Spec.Server)

	log.Info("GOT SECRET", "secret", string(secret.Data["username"]))
	log.Info("GOT SECRET", "secret", string(secret.Data["password"]))

	//location := instance.Spec.Location
	server := fmt.Sprintf("%s.database.windows.net", instance.Spec.Server)
	//groupName := instance.Spec.ResourceGroup
	database := "sqldatabase-aso-sql"
	var port = 1433
	var user = string(secret.Data["username"])
	var password = string(secret.Data["password"])

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, _ := sql.Open("sqlserver", connString)
	err := db.Ping()
	if err != nil {
		log.Info("Unable to ping server", "err", err.Error())
	} else {
		log.Info("succesfully logged in")
	}

	newUser := "usernametest0"
	newPassword := "pass@word1"
	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)

	// Execute non-query with named parameters
	_, err = db.ExecContext(ctx, tsql)
	if err != nil {
		log.Info("Error executing", "err", err.Error())
	} else {
		log.Info("Successfully executed user create")
	}

	tsql = fmt.Sprintf("sp_addrolemember '%s', %s", "db_owner", newUser)
	_, err = db.ExecContext(ctx, tsql)
	if err != nil {
		log.Info("Error executing", "err", err.Error())
	} else {
		log.Info("Successfully gave user owner creds")
	}

	secret = r.GetOrPrepareSecret(&instance, newUser)
	log.Info("GOT NEW SECRET", "secret", string(secret.Data["username"]))
	log.Info("GOT NEW SECRET", "secret", string(secret.Data["password"]))

	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
		r.Log.Info("mutating secret bundle")
		innerErr := controllerutil.SetControllerReference(&instance, secret, r.Scheme)
		if innerErr != nil {
			log.Info("inner err", "err", err.Error())
		}
		return nil
	})
	if createOrUpdateSecretErr != nil {
		log.Info("createOrUpdateSecretErr", "err", err.Error())
	}

	// TODO: write the secret
	return ctrl.Result{}, nil
}

// Get or prepare secret
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

func (r *SqlUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlUser{}).
		Complete(r)
}
