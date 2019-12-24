package server

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
)

type PSQLServerClient struct {
	Log logr.Logger
}

func NewPSQLServerClient(log logr.Logger) *PSQLServerClient {
	return &PSQLServerClient{
		Log: log,
	}
}

func getPSQLServersClient() psql.ServersClient {
	serversClient := psql.NewServersClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func (p *PSQLServerClient) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	p.Log.Info("Inside the Ensure method")

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	client := getPSQLServersClient()

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		labels[k] = &v
	}

	p.Log.Info("name is here", "locaiton", instance.Name)
	p.Log.Info("rg is here", "locaiton", instance.Spec.ResourceGroup)
	p.Log.Info("location is here", "locaiton", instance.Spec.Location)

	future, err := client.Create(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Name,
		psql.ServerForCreate{
			Location: &instance.Spec.Location,
			Tags:     labels,
			Properties: &psql.ServerPropertiesForDefaultCreate{
				AdministratorLogin:         to.StringPtr("adm1nus3r"),
				AdministratorLoginPassword: to.StringPtr("m@#terU$3r"),
				Version:                    psql.ServerVersion("10"),
				SslEnforcement:             psql.SslEnforcementEnumEnabled,
				//StorageProfile: &psql.StorageProfile{},
				CreateMode: psql.CreateModeServerPropertiesForCreate,
			},
			Sku: &psql.Sku{
				Name:     to.StringPtr(instance.Spec.Sku.Name),
				Tier:     psql.SkuTier(instance.Spec.Sku.Tier),
				Capacity: to.Int32Ptr(instance.Spec.Sku.Capacity),
				//Size:     to.StringPtr(instance.Spec.Sku.Size),
				Family: to.StringPtr(instance.Spec.Sku.Family),
			},
		},
	)

	if err != nil {
		return false, err
	}

	instance.Status.State = future.Status()

	server, err := future.Result(client)
	if err != nil {
		return false, err
	}

	instance.Status.State = server.Status

	return true, nil
}

func (p *PSQLServerClient) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	p.Log.Info("Inside the Delete method")
	return false, nil
}

func (p *PSQLServerClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	_, err := p.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{}, nil

}

func (p *PSQLServerClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
