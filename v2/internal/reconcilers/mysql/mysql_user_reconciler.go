/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"context"
	"database/sql"

	"github.com/go-logr/logr"
	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	dbformysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501storage"
	asomysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta1"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	mysqlutil "github.com/Azure/azure-service-operator/v2/internal/util/mysql"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

var _ genruntime.Reconciler = &MySQLUserReconciler{}

type MySQLUserReconciler struct {
	reconcilers.ARMOwnedResourceReconcilerCommon
	ResourceResolver *resolver.Resolver
	Config           config.Values
}

func NewMySQLUserReconciler(
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values) *MySQLUserReconciler {

	return &MySQLUserReconciler{
		ResourceResolver: resourceResolver,
		Config:           cfg,
		ARMOwnedResourceReconcilerCommon: reconcilers.ARMOwnedResourceReconcilerCommon{
			ResourceResolver: resourceResolver,
			ReconcilerCommon: reconcilers.ReconcilerCommon{
				KubeClient:         kubeClient,
				PositiveConditions: positiveConditions,
			},
		},
	}
}

func (r *MySQLUserReconciler) asUser(obj genruntime.MetaObject) (*asomysql.User, error) {
	typedObj, ok := obj.(*asomysql.User)
	if !ok {
		return nil, errors.Errorf("cannot modify resource that is not of type *asomysql.User. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *MySQLUserReconciler) CreateOrUpdate(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) (ctrl.Result, error) {
	user, err := r.asUser(obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Augment Log
	log = log.WithValues("azureName", user.AzureName())

	// Resolve the secrets
	secrets, err := r.ResourceResolver.ResolveResourceSecretReferences(ctx, user)
	if err != nil {
		return ctrl.Result{}, reconcilers.ClassifyResolverError(err)
	}

	db, err := r.connectToDB(ctx, log, user, secrets)
	if err != nil {
		return ctrl.Result{}, err
	}
	defer db.Close()

	log.V(Status).Info("Creating MySQL user")

	password, err := secrets.LookupFromPtr(user.Spec.LocalUser.Password)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to look up .spec.localUser.Password")
	}

	// Create or update the user. Note that this updates password if it has changed
	username := user.Spec.AzureName
	err = mysqlutil.CreateOrUpdateUser(ctx, db, user.Spec.AzureName, user.Spec.Hostname, password)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to create user")
	}

	// Ensure that the privileges are set
	err = mysqlutil.ReconcileUserServerPrivileges(ctx, db, username, user.Spec.Hostname, user.Spec.Privileges)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "ensuring server roles")
	}

	err = mysqlutil.ReconcileUserDatabasePrivileges(ctx, db, username, user.Spec.Hostname, user.Spec.DatabasePrivileges)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "ensuring database roles")
	}

	log.V(Status).Info("Successfully reconciled MySQLUser")

	return ctrl.Result{}, nil
}

func (r *MySQLUserReconciler) Delete(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) (ctrl.Result, error) {
	user, err := r.asUser(obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Augment Log
	log = log.WithValues("azureName", user.AzureName())

	log.V(Status).Info("Starting delete of resource")

	// Check that this objects owner still exists
	// This is an optimization to avoid excess requests to Azure.
	_, err = r.ResourceResolver.ResolveOwner(ctx, user)
	if err != nil {
		var typedErr *core.ReferenceNotFound
		if errors.As(err, &typedErr) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	secrets, err := r.ResourceResolver.ResolveResourceSecretReferences(ctx, user)
	if err != nil {
		return ctrl.Result{}, err
	}

	db, err := r.connectToDB(ctx, log, user, secrets)
	if err != nil {
		return ctrl.Result{}, err
	}
	defer db.Close()

	// TODO: There's still probably some ways that this user can be deleted but that we don't detect (and
	// TODO: so might cause an error triggering the resource to get stuck).
	// TODO: We check for owner not existing above, but cases where the server is in the process of being
	// TODO: deleted (or all system tables have been wiped?) might also exist...
	err = mysqlutil.DropUser(ctx, db, user.Spec.AzureName)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *MySQLUserReconciler) Claim(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) error {
	user, err := r.asUser(obj)
	if err != nil {
		return err
	}

	err = r.ARMOwnedResourceReconcilerCommon.ClaimResource(ctx, log, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLUserReconciler) UpdateStatus(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) error {
	user, err := r.asUser(obj)
	if err != nil {
		return err
	}

	secrets, err := r.ResourceResolver.ResolveResourceSecretReferences(ctx, user)
	if err != nil {
		return err
	}

	db, err := r.connectToDB(ctx, log, user, secrets)
	if err != nil {
		return err
	}
	defer db.Close()

	exists, err := mysqlutil.DoesUserExist(ctx, db, user.Spec.AzureName)
	if err != nil {
		return err
	}

	if !exists {
		err = errors.Errorf("user %s does not exist", user.Spec.AzureName)
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonAzureResourceNotFound)
		return err
	}

	return nil
}

func (r *MySQLUserReconciler) connectToDB(ctx context.Context, _ logr.Logger, user *asomysql.User, secrets genruntime.Resolved[genruntime.SecretReference]) (*sql.DB, error) {
	// Get the owner - at this point it must exist
	owner, err := r.ResourceResolver.ResolveOwner(ctx, user)
	if err != nil {
		return nil, err
	}

	flexibleServer, ok := owner.(*dbformysql.FlexibleServer)
	if !ok {
		return nil, errors.Errorf("owner was not type FlexibleServer, instead: %T", owner)
	}
	// Magical assertion to ensure that this is still the storage type
	var _ ctrlconversion.Hub = &dbformysql.FlexibleServer{}

	if flexibleServer.Status.FullyQualifiedDomainName == nil {
		// This possibly means that the server hasn't finished deploying yet
		err = errors.Errorf("owning Flexibleserver %q '.status.fullyQualifiedDomainName' not set. Has the server been provisioned successfully?", flexibleServer.Name)
		return nil, conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonWaitingForOwner)
	}
	serverFQDN := *flexibleServer.Status.FullyQualifiedDomainName

	adminPassword, err := secrets.LookupFromPtr(user.Spec.LocalUser.ServerAdminPassword)
	if err != nil {
		err = errors.Wrap(err, "failed to look up .spec.localUser.ServerAdminPassword")
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonSecretNotFound)
		return nil, err
	}

	// Admin User
	adminUser := user.Spec.LocalUser.ServerAdminUsername

	// Connect to the DB
	db, err := mysqlutil.ConnectToDB(ctx, serverFQDN, mysqlutil.SystemDatabase, mysqlutil.ServerPort, adminUser, adminPassword)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to connect database. Server: %s, Database: %s, Port: %d, AdminUser: %s",
			serverFQDN,
			mysqlutil.SystemDatabase,
			mysqlutil.ServerPort,
			adminUser)
	}

	return db, nil
}
