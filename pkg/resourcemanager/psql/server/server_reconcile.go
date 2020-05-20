// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates the Postgres server
func (p *PSQLServerClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		p.SecretClient = options.SecretClient
	}

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	createmode := string(psql.CreateModeDefault)
	if len(instance.Spec.CreateMode) != 0 {
		createmode = instance.Spec.CreateMode
	}

	// If a replica is requested, ensure that source server is specified
	if strings.EqualFold(createmode, string(psql.CreateModeReplica)) {
		if len(instance.Spec.ReplicaProperties.SourceServerId) == 0 {
			instance.Status.Message = "Replica requested but source server unspecified"
			return true, nil
		}
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := p.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}

	// Update secret with the fully qualified server name
	err = p.AddServerCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	hash := ""
	// if an error occurs thats ok as it means that it doesn't exist yet
	getServer, err := p.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.State = string(getServer.UserVisibleState)

		hash = helpers.Hash256(instance.Spec)
		if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
			instance.Status.RequestedAt = nil
			return true, nil
		}

		// succeeded! so end reconcilliation successfully
		if getServer.UserVisibleState == psql.ServerStateReady {

			// Update the secret with fully qualified server name. Ignore error as we have the admin creds which is critical.
			p.UpdateSecretWithFullServerName(ctx, instance.Name, secret, instance, *getServer.FullyQualifiedDomainName)

			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *getServer.ID
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.FailedProvisioning = false
			instance.Status.SpecHash = hash
			return true, nil
		}

		// the database exists but has not provisioned yet - so keep waiting
		instance.Status.Message = "Postgres server exists but may not be ready"
		return false, nil
	} else {
		// handle failures in the async operation
		if instance.Status.PollingURL != "" {
			pClient := pollclient.NewPollClient()
			res, err := pClient.Get(ctx, instance.Status.PollingURL)
			if err != nil {
				instance.Status.Provisioning = false
				return false, err
			}

			if res.Status == "Failed" {
				instance.Status.Provisioning = false
				instance.Status.RequestedAt = nil
				ignore := []string{
					errhelp.SubscriptionDoesNotHaveServer,
					errhelp.ServiceBusy,
				}
				if !helpers.ContainsString(ignore, res.Error.Code) {
					instance.Status.Message = res.Error.Error()
					return true, nil
				}
			}
		}
	}

	// setup variables for create call
	labels := helpers.LabelsToTags(instance.GetLabels())
	adminlogin := string(secret["username"])
	adminpassword := string(secret["password"])
	skuInfo := psql.Sku{
		Name:     to.StringPtr(instance.Spec.Sku.Name),
		Tier:     psql.SkuTier(instance.Spec.Sku.Tier),
		Capacity: to.Int32Ptr(instance.Spec.Sku.Capacity),
		Size:     to.StringPtr(instance.Spec.Sku.Size),
		Family:   to.StringPtr(instance.Spec.Sku.Family),
	}

	// create the server
	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false
	pollURL, _, err := p.CreateServerIfValid(
		ctx,
		*instance,
		labels,
		skuInfo,
		adminlogin,
		adminpassword,
		createmode,
	)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		azerr := errhelp.NewAzureErrorAzureError(err)

		catchInProgress := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
		}
		catchKnownError := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ServiceBusy,
			errhelp.InternalServerError,
		}

		// handle the errors
		if helpers.ContainsString(catchInProgress, azerr.Type) {
			if azerr.Type == errhelp.AsyncOpIncompleteError {
				instance.Status.PollingURL = pollURL
			}
			instance.Status.Message = "Postgres server exists but may not be ready"
			instance.Status.Provisioning = true
			return false, nil
		}

		if helpers.ContainsString(catchKnownError, azerr.Type) {
			return false, nil
		}

		// serious error occured, end reconcilliation and mark it as failed
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioned = false
		instance.Status.FailedProvisioning = true
		return true, nil

	}

	return false, nil
}

// Delete deletes the Postgres server
func (p *PSQLServerClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		p.SecretClient = options.SecretClient
	}

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err != nil {
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		gone := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}
	instance.Status.State = status

	if err == nil {
		if status != "InProgress" {
			// Best case deletion of secrets
			key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			p.SecretClient.Delete(ctx, key)
			return false, nil
		}
	}

	return true, nil
}

// GetParents gets the resource's parents
func (p *PSQLServerClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus returns the status
func (p *PSQLServerClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}
	st := v1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (p *PSQLServerClient) convert(obj runtime.Object) (*v1alpha2.PostgreSQLServer, error) {
	local, ok := obj.(*v1alpha2.PostgreSQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
