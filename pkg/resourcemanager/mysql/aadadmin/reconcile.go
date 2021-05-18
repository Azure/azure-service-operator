// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package aadadmin

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/gofrs/uuid"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

type MySQLServerAdministratorManager struct {
	creds config.Credentials
}

func NewMySQLServerAdministratorManager(creds config.Credentials) *MySQLServerAdministratorManager {
	return &MySQLServerAdministratorManager{
		creds: creds,
	}
}

func newMySQLServerAdministratorsClient(creds config.Credentials) (mysql.ServerAdministratorsClient, error) {
	adminClient := mysql.NewServerAdministratorsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return mysql.ServerAdministratorsClient{}, err
	}
	adminClient.Authorizer = a
	err = adminClient.AddToUserAgent(config.UserAgent())
	if err != nil {
		return mysql.ServerAdministratorsClient{}, err
	}
	return adminClient, nil
}

func (m *MySQLServerAdministratorManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	instance.Status.SetProvisioning("")
	client, err := newMySQLServerAdministratorsClient(m.creds)
	if err != nil {
		return true, err
	}

	_, err = client.Get(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)
	if err == nil {
		instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
		instance.Status.State = "Succeeded"
		return true, nil
	}

	tenantUUID, err := uuid.FromString(instance.Spec.TenantId)
	if err != nil {
		instance.Status.SetFailedProvisioning(fmt.Sprintf("failed to parse tenant UUID %s", instance.Spec.TenantId))
		instance.Status.State = "Failed"
		return true, nil
	}

	sidUUID, err := uuid.FromString(instance.Spec.Sid)
	if err != nil {
		instance.Status.SetFailedProvisioning(fmt.Sprintf("failed to parse sid UUID %s", instance.Spec.Sid))
		instance.Status.State = "Failed"
		return true, nil
	}

	future, err := client.CreateOrUpdate(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.Server,
		mysql.ServerAdministratorResource{
			ServerAdministratorProperties: &mysql.ServerAdministratorProperties{
				AdministratorType: to.StringPtr(string(instance.Spec.AdministratorType)),
				Login:             to.StringPtr(instance.Spec.Login),
				Sid:               &sidUUID,
				TenantID:          &tenantUUID,
			},
		})

	if err == nil {
		_, err = future.Result(client)
	}

	if err != nil {
		// let the user know what happened
		instance.Status.Message = errhelp.StripErrorIDs(err)

		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceNotFound,
		}

		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			// reconciliation is not done but error is acceptable
			return false, nil
		}

		instance.Status.SetFailedProvisioning(instance.Status.Message) // Preserve the message
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
	instance.Status.State = "Succeeded"
	return true, nil
}

func (m *MySQLServerAdministratorManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	client, err := newMySQLServerAdministratorsClient(m.creds)
	if err != nil {
		return true, err
	}

	// Check to see if the admin exists - if it doesn't short circuit
	_, err = client.Get(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// Resource doesn't exist
			return false, nil
		}

		// Ignoring other errors and just proceeding to delete
	}

	future, err := client.Delete(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)

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
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// Op is incomplete, continue checking
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}

	operationStatus := future.Status()
	instance.Status.State = operationStatus

	// Assume it exists by default
	return true, nil
}

func (m *MySQLServerAdministratorManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.MySQLServer{},
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

func (m *MySQLServerAdministratorManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *MySQLServerAdministratorManager) convert(obj runtime.Object) (*v1alpha1.MySQLServerAdministrator, error) {
	local, ok := obj.(*v1alpha1.MySQLServerAdministrator)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
