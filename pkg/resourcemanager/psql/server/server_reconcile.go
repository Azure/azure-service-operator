// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
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

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := p.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = p.AddServerCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	// if an error occurs thats ok as it means that it doesn't exist yet
	getServer, err := p.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.State = string(getServer.UserVisibleState)

		// succeeded! so end reconcilliation successfully
		if getServer.UserVisibleState == "Ready" {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *getServer.ID
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}

		// the database exists but has not provisioned yet - so keep waiting
		instance.Status.Message = "Postgres server exists but may not be ready"
		instance.Status.State = string(getServer.UserVisibleState)
		return false, nil
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
	_, err = p.CreateServerIfValid(
		ctx,
		instance.Name,
		instance.Spec.ResourceGroup,
		instance.Spec.Location,
		labels,
		psql.ServerVersion(instance.Spec.ServerVersion),
		psql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
		skuInfo,
		adminlogin,
		adminpassword,
	)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
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
		}

		// handle the errors
		if helpers.ContainsString(catchInProgress, azerr.Type) {
			instance.Status.Message = "Postgres server exists but may not be ready"
			return false, nil
		} else if helpers.ContainsString(catchKnownError, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		} else {

			// serious error occured, end reconcilliation and mark it as failed
			instance.Status.Message = fmt.Sprintf("Error occurred creating the Postgres server: %s", errhelp.StripErrorIDs(err))
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			instance.Status.FailedProvisioning = true
			return true, nil
		}
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
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			return true, err
		}
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
	return &instance.Status, nil
}

func (p *PSQLServerClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
