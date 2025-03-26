/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var _ genruntime.Reconciler = &EntraSecurityGroupReconciler{}

type EntraSecurityGroupReconciler struct {
	reconcilers.ReconcilerCommon
	ResourceResolver *resolver.Resolver
	Config           config.Values
}

func NewEntraSecurityGroupReconciler(
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
) *EntraSecurityGroupReconciler {
	return &EntraSecurityGroupReconciler{
		ResourceResolver: resourceResolver,
		Config:           cfg,
		ReconcilerCommon: reconcilers.ReconcilerCommon{
			KubeClient:         kubeClient,
			PositiveConditions: positiveConditions,
		},
	}
}

func (r *EntraSecurityGroupReconciler) asSecurityGroup(obj genruntime.MetaObject) (*asoentra.SecurityGroup, error) {
	typedObj, ok := obj.(*asoentra.SecurityGroup)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.SecurityGroup. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *EntraSecurityGroupReconciler) CreateOrUpdate(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	/*
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

		log.V(Status).Info("Creating PostgreSql user")

		password, err := secrets.LookupFromPtr(user.Spec.LocalUser.Password)
		if err != nil {
			return ctrl.Result{}, eris.Wrap(err, "failed to look up .spec.localUser.Password")
		}

		// Create or update the user. Note that this updates password if it has changed
		username := user.Spec.AzureName

		sqlUser, err := postgresqlutil.FindUserIfExist(ctx, db, username)
		if err != nil {
			return ctrl.Result{}, eris.Wrap(err, "failed to find user")
		}
		if sqlUser == nil {
			sqlUser, err = postgresqlutil.CreateUser(ctx, db, username, password)
			if err != nil {
				return ctrl.Result{}, eris.Wrap(err, "failed to create user")
			}
		} else {
			err = postgresqlutil.UpdateUser(ctx, db, *sqlUser, password)
			if err != nil {
				return ctrl.Result{}, eris.Wrap(err, "failed to update user")
			}
		}
		// TODO integrate in create and update user?
		if user.Spec.RoleOptions == nil {
			return ctrl.Result{}, eris.Wrap(err, "failed to look up .spec.roleOptions")
		}
		roleOptions := postgresqlutil.RoleOptions(*user.Spec.RoleOptions)
		// Ensure that the user role options are set
		err = postgresqlutil.ReconcileUserRoleOptions(ctx, db, *sqlUser, roleOptions)
		if err != nil {
			return ctrl.Result{}, eris.Wrap(err, "ensuring user role options")
		}

		// Ensure that the roles are set
		err = postgresqlutil.ReconcileUserServerRoles(ctx, db, *sqlUser, user.Spec.Roles)
		if err != nil {
			return ctrl.Result{}, eris.Wrap(err, "ensuring server roles")
		}

		log.V(Status).Info("Successfully reconciled PostgreSqlUser")
	*/

	return ctrl.Result{}, nil
}

func (r *EntraSecurityGroupReconciler) Delete(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	/*
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
			if eris.As(err, &typedErr) {
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
		err = postgresqlutil.DropUser(ctx, db, user.Spec.AzureName)
		if err != nil {
			return ctrl.Result{}, err
		}

	*/

	return ctrl.Result{}, nil
}

func (r *EntraSecurityGroupReconciler) Claim(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	/*
		user, err := r.asUser(obj)
		if err != nil {
			return err
		}

		err = r.ARMOwnedResourceReconcilerCommon.ClaimResource(ctx, log, user)
		if err != nil {
			return err
		}
	*/

	return nil
}

func (r *EntraSecurityGroupReconciler) UpdateStatus(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	/*
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

		exists, err := postgresqlutil.DoesUserExist(ctx, db, user.Spec.AzureName)
		if err != nil {
			return err
		}

		if !exists {
			err = eris.Errorf("user %s does not exist", user.Spec.AzureName)
			err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonAzureResourceNotFound)
			return err
		}
	*/

	return nil
}
