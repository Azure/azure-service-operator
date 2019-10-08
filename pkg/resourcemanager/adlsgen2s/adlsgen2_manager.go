package adlsgen2s

import (
	"context"

	apiv1 "github.com/Azure/azure-service-operator/api/v1"
)

type AdlsGen2Manager interface {
	CreateAdlsGen2()

	GetAdlsGen2()

	DeleteAdlsGen2()
}