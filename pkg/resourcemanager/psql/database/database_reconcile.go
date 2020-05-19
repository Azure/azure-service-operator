// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure ensures a Postgres database exists
func (p *PSQLDatabaseClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	getDB, err := p.GetDatabase(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err == nil {

		// succeeded! so end reconcilliation successfully
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *getDB.ID
		instance.Status.State = getDB.Status
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false
	resp, err := p.CreateDatabaseIfValid(
		ctx,
		instance.Name,
		instance.Spec.Server,
		instance.Spec.ResourceGroup,
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
			errhelp.ResourceNotFound,
		}

		// assertion that a 404 error implies that the Postgres server hasn't been provisioned yet
		if resp != nil && resp.StatusCode == 404 {
			instance.Status.Message = fmt.Sprintf("Waiting for Postgres server %s to provision", instance.Spec.Server)
			instance.Status.Provisioning = false
			return false, nil
		}

		// handle the errors
		if helpers.ContainsString(catchInProgress, azerr.Type) {
			instance.Status.Message = "Postgres database exists but may not be ready"
			instance.Status.Provisioning = true
			return false, nil
		} else if helpers.ContainsString(catchKnownError, azerr.Type) {
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

// Delete removes the Postgres database
func (p *PSQLDatabaseClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteDatabase(ctx, instance.Name, instance.Spec.Server, instance.Spec.ResourceGroup)
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

	if status != "InProgress" {
		return false, nil
	}

	return true, nil
}

// GetParents gets the database's parents
func (p *PSQLDatabaseClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.PostgreSQLServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus gets the status
func (p *PSQLDatabaseClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (p *PSQLDatabaseClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLDatabase, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
