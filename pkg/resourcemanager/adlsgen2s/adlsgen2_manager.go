package adlsgen2s

import (
	"context"
	"github.com/Azure/go-autorest/autorest"	
)

type AdlsGen2Manager interface {
	CreateAdlsGen2(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string, accountName string) (*autorest.Response, error)

	GetAdlsGen2()

	DeleteAdlsGen2()
}
