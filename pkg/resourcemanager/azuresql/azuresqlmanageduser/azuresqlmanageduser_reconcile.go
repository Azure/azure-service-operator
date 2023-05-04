// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlmanageduser

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/helpers"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"

	_ "github.com/denisenkom/go-mssqldb"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure that user exists
func (s *AzureSqlManagedUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	requestedUsername := instance.Spec.ManagedIdentityName
	if len(requestedUsername) == 0 {
		requestedUsername = instance.Name
	}

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := s.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	subscriptionID := instance.Spec.SubscriptionID
	_, err = s.GetDB(ctx, subscriptionID, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		requeueErrors := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(requeueErrors, azerr.Type) {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		// if this is an unmarshall error - ignore and continue, otherwise report error and requeue
		if !strings.Contains(errorString, "cannot unmarshal array into Go struct field serviceError2.details") {
			return false, err
		}
	}

	db, err := s.ConnectToSqlDbAsCurrentUser(ctx, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		instance.Status.SetFailedProvisioning(instance.Status.Message)
		return false, nil
	}
	defer db.Close()

	userExists, err := s.UserExists(ctx, db, requestedUsername)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %s", err)
		return false, nil
	}

	if !userExists {
		instance.Status.SetProvisioning("")

		err = s.EnableUser(ctx, requestedUsername, instance.Spec.ManagedIdentityClientId, db)
		if err != nil {
			instance.Status.Message = fmt.Sprintf("failed enabling managed identity user, err: %s", err)
			if strings.Contains(err.Error(), "The login already has an account under a different user name") {
				instance.Status.SetFailedProvisioning(instance.Status.Message)
				return true, nil
			}
			return false, err
		}
	}

	// apply roles to user
	if len(instance.Spec.Roles) == 0 {
		instance.Status.Message = "No roles specified for user"
		return false, fmt.Errorf("no roles specified for database user")
	}

	err = s.GrantUserRoles(ctx, requestedUsername, instance.Spec.Roles, db)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("GrantUserRoles failed: %s", err)
		return false, errors.Wrap(err, "GrantUserRoles failed")
	}

	err = s.UpdateSecret(ctx, instance, secretClient)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Updating secret failed: %s", err)
		return false, fmt.Errorf("updating secret failed")
	}

	instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
	instance.Status.State = "Succeeded"
	return true, nil
}

// Delete deletes a user
func (s *AzureSqlManagedUserManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	secretClient := s.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	requestedUsername := instance.Spec.ManagedIdentityName
	if len(requestedUsername) == 0 {
		requestedUsername = instance.Name
	}

	// short circuit connection if database doesn't exist
	_, err = s.GetDB(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// Best case deletion of secrets
			s.DeleteSecrets(ctx, instance, secretClient)

			return false, nil
		}
		return false, err
	}

	db, err := s.ConnectToSqlDbAsCurrentUser(ctx, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		return true, nil
	}
	defer db.Close()

	userExists, err := s.UserExists(ctx, db, requestedUsername)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %s", err)
		return false, nil
	}

	if userExists {
		err = s.DropUser(ctx, db, requestedUsername)
		if err != nil {
			instance.Status.Message = "failed dropping managed identity user, err: " + err.Error()
			return false, err
		}
	}

	// Best case deletion of secrets
	s.DeleteSecrets(ctx, instance, secretClient)

	instance.Status.Message = fmt.Sprintf("Delete AzureSqlManagedUser succeeded")

	return false, nil
}

// GetParents gets the parents of the user
func (s *AzureSqlManagedUserManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.DbName,
			},
			Target: &v1alpha1.AzureSqlDatabase{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &v1alpha1.AzureSqlServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus gets the status
func (s *AzureSqlManagedUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (s *AzureSqlManagedUserManager) convert(obj runtime.Object) (*v1alpha1.AzureSQLManagedUser, error) {
	local, ok := obj.(*v1alpha1.AzureSQLManagedUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
