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
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	asomysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta1"
	dbformysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501storage"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	mysqlutil "github.com/Azure/azure-service-operator/v2/internal/util/mysql"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var _ genruntime.Reconciler = &MySQLUserReconciler{}

type MySQLUserReconciler struct {
	reconcilers.ARMOwnedResourceReconcilerCommon
	KubeClient         kubeclient.Client
	ResourceResolver   *resolver.Resolver
	PositiveConditions *conditions.PositiveConditionBuilder
	Config             config.Values
}

func NewMySQLUserReconciler(
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values) *MySQLUserReconciler {

	return &MySQLUserReconciler{
		KubeClient:         kubeClient,
		ResourceResolver:   resourceResolver,
		PositiveConditions: positiveConditions,
		Config:             cfg,
		ARMOwnedResourceReconcilerCommon: reconcilers.ARMOwnedResourceReconcilerCommon{
			ResourceResolver: resourceResolver,
			ReconcilerCommon: reconcilers.ReconcilerCommon{
				KubeClient: kubeClient,
			},
		},
	}
}

func (r *MySQLUserReconciler) Reconcile(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject) (ctrl.Result, error) {

	typedObj, ok := obj.(*asomysql.User)
	if !ok {
		return ctrl.Result{}, errors.Errorf("cannot modify resource that is not of type *asomysql.User. Type is %T", obj)
	}

	// Augment Log
	log = log.WithValues("azureName", typedObj.AzureName())

	var result ctrl.Result
	var err error
	if !obj.GetDeletionTimestamp().IsZero() {
		result, err = r.Delete(ctx, log, typedObj)
	} else {
		result, err = r.CreateOrUpdate(ctx, log, typedObj)
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *MySQLUserReconciler) Delete(ctx context.Context, log logr.Logger, user *asomysql.User) (ctrl.Result, error) {
	// TODO: A lot of this is duplicated. See https://azure.github.io/azure-service-operator/design/reconcile-interface for a proposal to fix that
	log.V(Status).Info("Starting delete of resource")

	// If we have no resourceID to begin with, or no finalizer, the Azure resource was never created
	hasFinalizer := controllerutil.ContainsFinalizer(user, reconcilers.GenericControllerFinalizer)
	if !hasFinalizer {
		return ctrl.Result{}, r.RemoveResourceFinalizer(ctx, log, user)
	}

	reconcilePolicy := reconcilers.GetReconcilePolicy(user, log)
	if !reconcilePolicy.AllowsDelete() {
		log.V(Info).Info("Bypassing delete of resource due to policy", "policy", reconcilePolicy)
		return ctrl.Result{}, r.RemoveResourceFinalizer(ctx, log, user)
	}

	// Check that this objects owner still exists
	// This is an optimization to avoid making excess requests to Azure.
	_, err := r.ResourceResolver.ResolveOwner(ctx, user)
	if err != nil {
		var typedErr *resolver.ReferenceNotFound
		if errors.As(err, &typedErr) {
			return ctrl.Result{}, r.RemoveResourceFinalizer(ctx, log, user)
		}
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

	err = r.RemoveResourceFinalizer(ctx, log, user)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *MySQLUserReconciler) CreateOrUpdate(ctx context.Context, log logr.Logger, user *asomysql.User) (ctrl.Result, error) {
	if r.NeedToClaimResource(user) {
		// TODO: Make this synchronous for both ARM resources and here
		// TODO: Will do this in a follow-up PR
		_, err := r.ClaimResource(ctx, log, user)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

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

	// TODO: A lot of this is duplicated. See https://azure.github.io/azure-service-operator/design/reconcile-interface for a proposal to fix that
	reconcilePolicy := reconcilers.GetReconcilePolicy(user, log)
	if !reconcilePolicy.AllowsModify() {
		return r.handleSkipReconcile(ctx, db, log, user)
	}

	log.V(Status).Info("Creating MySQL user")

	password, err := secrets.LookupSecretFromPtr(user.Spec.LocalUser.Password)
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

	conditions.SetCondition(user, r.PositiveConditions.Ready.Succeeded(user.GetGeneration()))
	err = r.CommitUpdate(ctx, log, user)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *MySQLUserReconciler) connectToDB(ctx context.Context, log logr.Logger, user *asomysql.User, secrets genruntime.ResolvedSecrets) (*sql.DB, error) {
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

	adminPassword, err := secrets.LookupSecretFromPtr(user.Spec.LocalUser.ServerAdminPassword)
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

// TODO: A lot of this is duplicated. See https://azure.github.io/azure-service-operator/design/reconcile-interface for a proposal to fix that
func (r *MySQLUserReconciler) handleSkipReconcile(ctx context.Context, db *sql.DB, log logr.Logger, user *asomysql.User) (ctrl.Result, error) {
	reconcilePolicy := reconcilers.GetReconcilePolicy(user, log)
	log.V(Status).Info(
		"Skipping creation of MySQLUser due to policy",
		reconcilers.ReconcilePolicyAnnotation, reconcilePolicy)

	exists, err := mysqlutil.DoesUserExist(ctx, db, user.Spec.AzureName)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !exists {
		err = errors.Errorf("user %s does not exist", user.Spec.AzureName)
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonAzureResourceNotFound)
		return ctrl.Result{}, err
	}

	conditions.SetCondition(user, r.PositiveConditions.Ready.Succeeded(user.GetGeneration()))
	if err := r.CommitUpdate(ctx, log, user); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
