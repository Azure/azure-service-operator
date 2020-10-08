// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlmanageduser

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"

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

	if options.SecretClient != nil {
		s.SecretClient = options.SecretClient
	}

	_, err = s.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
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

		// Other failures are terminal
		instance.Status.SetFailedProvisioning(instance.Status.Message)
		return true, nil
	}

	userExists, err := s.UserExists(ctx, db, requestedUsername)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %v", err)
		return false, nil
	}

	if !userExists {
		instance.Status.SetProvisioning("")

		err = s.EnableUser(ctx, requestedUsername, instance.Spec.ManagedIdentityClientId, db)
		if err != nil {
			instance.Status.Message = fmt.Sprintf("failed enabling managed identity user, err: %v", err)
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
		instance.Status.Message = fmt.Sprintf("GrantUserRoles failed: %v", err)
		return false, fmt.Errorf("GrantUserRoles failed: %v", err)
	}

	err = s.UpdateSecret(ctx, instance, s.SecretClient)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Updating secret failed: %v", err)
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

	requestedUsername := instance.Spec.ManagedIdentityName
	if len(requestedUsername) == 0 {
		requestedUsername = instance.Name
	}

	// short circuit connection if database doesn't exist
	_, err = s.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
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
			s.DeleteSecrets(ctx, instance, s.SecretClient)

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

	userExists, err := s.UserExists(ctx, db, requestedUsername)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %v", err)
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
	s.DeleteSecrets(ctx, instance, s.SecretClient)

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
