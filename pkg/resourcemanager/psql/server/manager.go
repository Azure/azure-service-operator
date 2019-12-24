package server

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLServerManager interface {
	// also embed async client methods
	resourcemanager.ARMClient
}
