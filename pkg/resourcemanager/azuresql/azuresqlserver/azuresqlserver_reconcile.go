// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const usernameLength = 8
const passwordLength = 16

// Ensure creates an AzureSqlServer
func (s *AzureSqlServerManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		s.SecretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	// set a spec hash if one hasn't been set
	hash := helpers.Hash256(instance.Spec)
	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	// Check to see if secret already exists for admin username/password
	// create or update the secret
	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	secret, err := s.SecretClient.Get(ctx, key)
	if err != nil {
		if instance.Status.Provisioned {
			instance.Status.Message = err.Error()
			return false, fmt.Errorf("Secret missing for provisioned server: %s", key.String())
		}

		checkNameResult, _ := CheckNameAvailability(ctx, instance.Name)
		if checkNameResult.Reason == sql.AlreadyExists {
			instance.Status.Provisioning = false
			if _, err := s.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name); err != nil {
				instance.Status.Message = "SQL server already exists somewhere else"
				return true, nil
			}

			instance.Status.Message = fmt.Sprintf(
				`SQL server already exists and the credentials could not be found. 
				If using kube secrets a secret should exist at '%s' for keyvault it should be '%s'`,
				key.String(),
				fmt.Sprintf("%s-%s", key.Namespace, key.Name),
			)

			return false, nil
		}

		secret, err = NewSecret(instance.Name)
		if err != nil {
			return false, err
		}
		err = s.SecretClient.Upsert(
			ctx,
			key,
			secret,
			secrets.WithOwner(instance),
			secrets.WithScheme(s.Scheme),
		)
		if err != nil {
			return false, err
		}
	}

	azureSQLServerProperties := azuresqlshared.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(string(secret["username"])),
		AdministratorLoginPassword: to.StringPtr(string(secret["password"])),
	}

	if instance.Status.SpecHash == "" {
		instance.Status.SpecHash = hash
	}

	if instance.Status.Provisioning {

		serv, err := s.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			azerr := errhelp.NewAzureErrorAzureError(err)
			// @Todo: ResourceNotFound should be handled if the time since the last PUT is unreasonable
			if azerr.Type != errhelp.ResourceNotFound {
				return false, err
			}

			// the first minute or so after a PUT to create a server will result in failed GETs
			instance.Status.State = "NotReady"
		} else {
			instance.Status.State = *serv.State
		}

		if instance.Status.State == "Ready" {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.ResourceId = *serv.ID
			return true, nil
		}

		// server not done provisioning
		return false, nil
	}

	// create the sql server
	instance.Status.Provisioning = true
	if _, err := s.CreateOrUpdateSQLServer(ctx, instance.Spec.ResourceGroup, instance.Spec.Location, instance.Name, azureSQLServerProperties, false); err != nil {
		instance.Status.Message = err.Error()

		azerr := errhelp.NewAzureErrorAzureError(err)

		// the first successful call to create the server should result in this type of error
		// we save the credentials here
		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.Message = "Resource request successfully submitted to Azure"
			return false, nil
		}

		// SQL Server names are globally unique so if a server with this name exists we need to see if it meets our criteria (ie. same rg/sub)
		if azerr.Type == errhelp.AlreadyExists {
			// see if server exists in correct rg
			if _, err := s.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name); err == nil {
				// should be good, let the next reconcile finish
				return false, nil
			}

			// this server likely belongs to someone else
			instance.Status.Provisioning = false
			instance.Status.RequestedAt = nil

			return true, nil
		}

		// these errors are expected for recoverable states
		// ignore them and try again after some time
		ignore := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(ignore, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}

		// these errors can't be recovered from without a change
		// to the server resource's manifest, which will cause a new event/reconciliation
		drop := []string{
			errhelp.InvalidServerName,
		}
		if helpers.ContainsString(drop, azerr.Type) {
			return true, nil
		}

		return false, err
	}

	return true, nil
}

// Delete handles idempotent deletion of a sql server
func (s *AzureSqlServerManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		s.SecretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}

	// if the resource is in a failed state it was never creatred or could never be verified
	// so we skip attempting to delete the resrouce from Azure
	if instance.Status.FailedProvisioning || strings.Contains(instance.Status.Message, "credentials could not be found") {
		return false, nil
	}

	_, err = s.DeleteSQLServer(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		// these errors are expected
		ignore := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.ConflictingServerOperation,
		}

		// this means the thing doesn't exist
		finished := []string{
			errhelp.ResourceNotFound,
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return true, nil
		}

		if helpers.ContainsString(finished, azerr.Type) {
			//Best effort deletion of secrets
			s.SecretClient.Delete(ctx, key)
			return false, nil
		}

		return false, err
	}

	//Best effort deletion of secrets
	s.SecretClient.Delete(ctx, key)
	return false, nil
}

// GetParents returns the parents of AzureSqlDatabase
func (s *AzureSqlServerManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}

	rgKey := types.NamespacedName{Name: instance.Spec.ResourceGroup, Namespace: instance.Namespace}

	return []resourcemanager.KubeParent{
		{Key: rgKey, Target: &v1alpha1.ResourceGroup{}},
	}, nil
}

func (g *AzureSqlServerManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (s *AzureSqlServerManager) convert(obj runtime.Object) (*v1alpha1.AzureSqlServer, error) {
	local, ok := obj.(*v1alpha1.AzureSqlServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// NewSecret generates a new sqlserver secret
func NewSecret(serverName string) (map[string][]byte, error) {

	secret := map[string][]byte{}

	randomUsername, err := helpers.GenerateRandomUsername(usernameLength, (usernameLength / 2))
	if err != nil {
		return secret, err
	}

	randomPassword, err := helpers.GenerateRandomPassword(passwordLength)
	if err != nil {
		return secret, err
	}

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, serverName))
	secret["password"] = []byte(randomPassword)
	secret["azureSqlServerName"] = []byte(serverName)
	secret["fullyQualifiedServerName"] = []byte(serverName + ".database.windows.net")

	return secret, nil
}
