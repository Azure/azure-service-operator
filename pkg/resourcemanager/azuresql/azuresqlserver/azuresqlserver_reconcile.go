// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlserver

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

const usernameLength = 8
const passwordLength = 16

// Ensure creates an AzureSqlServer
func (s *AzureSqlServerManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := s.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	// convert kube labels to expected tag format
	tags := helpers.LabelsToTags(instance.GetLabels())

	// Check to see if secret already exists for admin username/password
	// create or update the secret
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	secret, err := secretClient.Get(ctx, secretKey)
	if err != nil {
		if instance.Status.Provisioned {
			instance.Status.Message = err.Error()
			return false, fmt.Errorf("secret missing for provisioned server: %s", secretKey)
		}

		// Assure that the requested name is available and assume the secret exists
		checkNameResult, err := CheckNameAvailability(ctx, s.Creds, instance.Name)

		if err != nil {
			instance.Status.Provisioning = false
			return false, err
		}
		if *checkNameResult.Available != true {
			instance.Status.Provisioning = false
			if _, err := s.GetServer(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Name); err != nil {
				instance.Status.Message = "SQL server already exists somewhere else"
				return true, nil
			}

			err = errors.Wrapf(err, "SQL server already exists and the credentials could not be found")
			instance.Status.Message = err.Error()
			return false, nil
		}

		secret, err = NewSecret(instance.Name)
		if err != nil {
			return false, err
		}
		err = secretClient.Upsert(
			ctx,
			secretKey,
			secret,
			secrets.WithOwner(instance),
			secrets.WithScheme(s.Scheme),
		)
		if err != nil {
			instance.Status.Message = err.Error()
			return false, err
		}
	}

	azureSQLServerProperties := azuresqlshared.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(string(secret["username"])),
		AdministratorLoginPassword: to.StringPtr(string(secret["password"])),
	}

	// set a spec hash if one hasn't been set
	hash := helpers.Hash256(instance.Spec)
	specHashWasEmpty := false
	if instance.Status.SpecHash == "" {
		instance.Status.SpecHash = hash
		specHashWasEmpty = true
	}

	// early exit if hashes match and this server has been provisioned / failed provisioning
	if !specHashWasEmpty &&
		instance.Status.SpecHash == hash &&
		(instance.Status.Provisioned || instance.Status.FailedProvisioning) {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	// if we've already started provisioning then try to get the server,
	// 	or if the hash wasnt empty and the Spec Hash matches the calculated hash
	// 	this indicates that we've already issued and update and its worth
	// 	checking to see if the server is there

	if instance.Status.Provisioning ||
		(!specHashWasEmpty && instance.Status.SpecHash == hash) {

		serv, err := s.GetServer(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			azerr := errhelp.NewAzureError(err)

			// handle failures in the async operation
			if instance.Status.PollingURL != "" {
				pClient := pollclient.NewPollClient(s.Creds)
				res, err := pClient.Get(ctx, instance.Status.PollingURL)
				if err != nil {
					return false, err
				}

				if res.Status == pollclient.LongRunningOperationPollStatusFailed {
					instance.Status.Message = res.Error.Error()
					instance.Status.Provisioning = false
					return true, nil
				}
			}

			if azerr.Type == errhelp.ResourceGroupNotFoundErrorCode {
				return false, nil
			}

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
			instance.Status.FailedProvisioning = false
			instance.Status.ResourceId = *serv.ID
			instance.Status.SpecHash = hash
			instance.Status.PollingURL = ""
			return true, nil
		}

		// server not done provisioning
		return false, nil
	}

	// create the sql server
	instance.Status.Provisioning = true
	pollURL, _, err := s.CreateOrUpdateSQLServer(
		ctx,
		instance.Spec.SubscriptionID,
		instance.Spec.ResourceGroup,
		instance.Spec.Location,
		instance.Name,
		tags,
		azureSQLServerProperties,
		false)

	if err != nil {
		instance.Status.Message = err.Error()

		// check for our known errors
		azerr := errhelp.NewAzureError(err)

		switch azerr.Type {
		case errhelp.AsyncOpIncompleteError:
			// the first successful call to create the server should result in this type of error
			// we save the credentials here
			instance.Status.Message = "Resource request successfully submitted to Azure"
			instance.Status.PollingURL = pollURL
			return false, nil
		case errhelp.AlreadyExists:
			// SQL Server names are globally unique so if a server with this name exists we
			// need to see if it meets our criteria (ie. same rg/sub)
			// see if server exists in correct rg
			if serv, err := s.GetServer(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Name); err == nil {
				// mismatched location
				if *serv.Location != instance.Spec.Location {
					instance.Status.Provisioned = false
					instance.Status.Provisioning = false
					instance.Status.Message = fmt.Sprintf("%s does not match location of existing sql server, %s", instance.Spec.Location, *serv.Location)
					return true, nil
				}

				// should be good, let the next reconcile finish
				return false, nil
			}
			// this server likely belongs to someone else
			instance.Status.Provisioning = false
			instance.Status.RequestedAt = nil
			return true, nil
		case errhelp.LocationNotAvailableForResourceType,
			errhelp.RequestDisallowedByPolicy,
			errhelp.RegionDoesNotAllowProvisioning,
			errhelp.InvalidResourceLocation,
			errhelp.QuotaExceeded:

			instance.Status.Message = "Unable to provision Azure SQL Server due to error: " + errhelp.StripErrorIDs(err)
			instance.Status.Provisioning = false
			instance.Status.Provisioned = false
			instance.Status.FailedProvisioning = true
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

	secretClient := s.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	subscriptionID := instance.Spec.SubscriptionID
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	// if the resource is in a failed state it was never created or could never be verified
	// so we skip attempting to delete the resrouce from Azure
	if instance.Status.FailedProvisioning || strings.Contains(instance.Status.Message, "credentials could not be found") {
		return false, nil
	}

	_, err = s.DeleteSQLServer(ctx, subscriptionID, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

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
			secretClient.Delete(ctx, secretKey)
			return false, nil
		}

		return false, err
	}

	//Best effort deletion of secrets
	secretClient.Delete(ctx, secretKey)
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
	st := v1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (s *AzureSqlServerManager) convert(obj runtime.Object) (*v1beta1.AzureSqlServer, error) {
	local, ok := obj.(*v1beta1.AzureSqlServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// NewSecret generates a new sqlserver secret
func NewSecret(serverName string) (map[string][]byte, error) {

	secret := map[string][]byte{}

	randomUsername := helpers.GenerateRandomUsername(usernameLength)
	randomPassword := helpers.NewPassword()

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, serverName))
	secret["password"] = []byte(randomPassword)
	secret["azureSqlServerName"] = []byte(serverName)
	secret["fullyQualifiedServerName"] = []byte(serverName + "." + config.Environment().SQLDatabaseDNSSuffix)

	return secret, nil
}
