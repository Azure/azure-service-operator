package adlsgen2s

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/go-autorest/autorest"
)

type AdlsGen2Manager interface {
	CreateAdlsGen2(ctx context.Context, groupName string,
		datalakeName string,
		location string,
		sku azurev1alpha1.StorageSku,
		tags map[string]*string,
		accessTier azurev1alpha1.StorageAccessTier,
		enableHTTPSTrafficOnly *bool) (*storage.Account, error)

	GetAdlsGen2(ctx context.Context, groupName string, datalakeName string) (result storage.Account, err error)

	DeleteAdlsGen2(ctx context.Context, groupName string, datalakeName string) (result autorest.Response, err error)
}
