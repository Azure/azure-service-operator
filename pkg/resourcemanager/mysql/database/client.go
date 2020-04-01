// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"
	"fmt"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type MySQLDatabaseClient struct {
}

func NewMySQLDatabaseClient() *MySQLDatabaseClient {
	return &MySQLDatabaseClient{}
}

func getMySQLDatabasesClient() mysql.DatabasesClient {
	databasesClient := mysql.NewDatabasesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	databasesClient.Authorizer = a
	databasesClient.AddToUserAgent(config.UserAgent())
	return databasesClient
}

func getMySQLCheckNameAvailabilityClient() mysql.CheckNameAvailabilityClient {
	nameavailabilityClient := mysql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (m *MySQLDatabaseClient) CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error) {

	client := getMySQLCheckNameAvailabilityClient()

	resourceType := "database"

	nameAvailabilityRequest := mysql.NameAvailabilityRequest{
		Name: &databasename,
		Type: &resourceType,
	}
	_, err := client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}
func (m *MySQLDatabaseClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	client := getMySQLDatabasesClient()

	instance.Status.Provisioning = true
	// Check if this database already exists. This is required
	// to overcome the issue with the lack of idempotence of the Create call

	_, err = m.GetDatabase(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
		return true, nil
	}
	future, err := m.CreateDatabaseIfValid(
		ctx,
		instance.Name,
		instance.Spec.Server,
		instance.Spec.ResourceGroup,
	)

	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceNotFound,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	_, err = future.Result(client)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.SubscriptionDoesNotHaveServer,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			case errhelp.SubscriptionDoesNotHaveServer:
				instance.Status.Message = fmt.Sprintf("The PostgreSQL Server %s has not been provisioned yet. ", instance.Spec.Server)
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (m *MySQLDatabaseClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := m.DeleteDatabase(ctx, instance.Name, instance.Spec.Server, instance.Spec.ResourceGroup)
	if err != nil {
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			return true, err
		}
	}

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}

func (m *MySQLDatabaseClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (m *MySQLDatabaseClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *MySQLDatabaseClient) convert(obj runtime.Object) (*v1alpha1.MySQLDatabase, error) {
	local, ok := obj.(*v1alpha1.MySQLDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (m *MySQLDatabaseClient) CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (future mysql.DatabasesCreateOrUpdateFuture, err error) {

	client := getMySQLDatabasesClient()

	// Check if name is valid if this is the first create call
	valid, err := m.CheckDatabaseNameAvailability(ctx, databasename)
	if valid == false {
		return future, err
	}

	dbParameters := mysql.Database{}

	future, err = client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		databasename,
		dbParameters,
	)

	return future, err
}

func (m *MySQLDatabaseClient) DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (status string, err error) {

	client := getMySQLDatabasesClient()

	_, err = client.Get(ctx, resourcegroup, servername, databasename)
	if err == nil { // db present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, databasename)
		return future.Status(), err
	}
	// db not present so return success anyway
	return "db not present", nil

}

func (m *MySQLDatabaseClient) GetDatabase(ctx context.Context, resourcegroup string, servername string, databasename string) (db mysql.Database, err error) {

	client := getMySQLDatabasesClient()

	return client.Get(ctx, resourcegroup, servername, databasename)
}
