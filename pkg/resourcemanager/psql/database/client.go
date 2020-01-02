package database

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type PSQLDatabaseClient struct {
	Log logr.Logger
}

func NewPSQLDatabaseClient(log logr.Logger) *PSQLDatabaseClient {
	return &PSQLDatabaseClient{
		Log: log,
	}
}

func getPSQLDatabasesClient() psql.DatabasesClient {
	databasesClient := psql.NewDatabasesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	databasesClient.Authorizer = a
	databasesClient.AddToUserAgent(config.UserAgent())
	return databasesClient
}

func getPSQLCheckNameAvailabilityClient() psql.CheckNameAvailabilityClient {
	nameavailabilityClient := psql.NewCheckNameAvailabilityClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (p *PSQLDatabaseClient) CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error) {

	client := getPSQLCheckNameAvailabilityClient()

	resourceType := "database"

	nameAvailabilityRequest := psql.NameAvailabilityRequest{
		Name: &databasename,
		Type: &resourceType,
	}
	_, err := client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}
func (p *PSQLDatabaseClient) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	client := getPSQLDatabasesClient()

	instance.Status.Provisioning = true
	// Check if this database already exists and its state if it does. This is required
	// to overcome the issue with the lack of idempotence of the Create call

	db, err := p.GetDatabase(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err == nil {
		p.Log.Info("Database present", "State=", db.Status)
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.State = db.Status
		return true, nil
	}
	p.Log.Info("Ensure: Database not present, creating")

	future, err := p.CreateDatabaseIfValid(
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

	instance.Status.State = future.Status()

	db, err = future.Result(client)
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
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		p.Log.Info("Ensure", "azerr.Type", azerr.Type)
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

	instance.Status.State = db.Status

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = "Provisioned successfully"
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (p *PSQLDatabaseClient) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteDatabase(ctx, instance.Name, instance.Spec.Server, instance.Spec.ResourceGroup)
	if err != nil {
		p.Log.Info("Delete:", "db Delete returned=", err.Error())
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			p.Log.Info("Error from delete call")
			return true, err
		}
	}
	instance.Status.State = status
	p.Log.Info("Delete", "future.Status=", status)

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}

func (p *PSQLDatabaseClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.PostgreSQLServer{},
		},
	}, nil

}

func (p *PSQLDatabaseClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLDatabase, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (p *PSQLDatabaseClient) CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (future psql.DatabasesCreateOrUpdateFuture, err error) {

	client := getPSQLDatabasesClient()

	// Check if name is valid if this is the first create call
	valid, err := p.CheckDatabaseNameAvailability(ctx, databasename)
	if valid == false {
		p.Log.Info("Database name invalid - cannot create db")
		return future, err
	}

	dbParameters := psql.Database{}

	future, err = client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		databasename,
		dbParameters,
	)

	return future, err
}

func (p *PSQLDatabaseClient) DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (status string, err error) {

	client := getPSQLDatabasesClient()

	_, err = client.Get(ctx, resourcegroup, servername, databasename)
	if err == nil { // db present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, databasename)
		return future.Status(), err
	}
	// db not present so return success anyway
	return "db not present", nil

}

func (p *PSQLDatabaseClient) GetDatabase(ctx context.Context, resourcegroup string, servername string, databasename string) (db psql.Database, err error) {

	client := getPSQLDatabasesClient()

	return client.Get(ctx, resourcegroup, servername, databasename)
}
