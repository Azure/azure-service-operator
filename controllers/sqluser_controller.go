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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

	//location := instance.Spec.Location
	server := "midobrzysqlserver.database.windows.net"
	//groupName := instance.Spec.ResourceGroup
	database := "sqldatabase-aso-sql"
	var port = 1433
	//secret, _ := helpers.GetSecret("name", "namespace")
	var user = ")W_FFhAe"
	var password = "mS_n9M<lfLrqiAo&"

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

	return ctrl.Result{}, nil
}

func (r *SqlUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlUser{}).
		Complete(r)
}
