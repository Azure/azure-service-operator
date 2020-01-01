package server

import (
	"context"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLServerManager interface {
	//convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error)

	//CheckServerNameAvailability(ctx context.Context, servername string) (bool, error)
	CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion psql.ServerVersion, sslenforcement psql.SslEnforcementEnum, skuInfo psql.Sku, adminlogin string, adminpassword string) (psql.ServersCreateFuture, error)
	DeleteServer(ctx context.Context, resourcegroup string, servername string) (string, error)
	GetServer(ctx context.Context, resourcegroup string, servername string) (psql.Server, error)
	//AddServerCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.PostgreSQLServer) error

	// also embed async client methods
	resourcemanager.ARMClient
}
