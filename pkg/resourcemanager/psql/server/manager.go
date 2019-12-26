package server

import (
	"context"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
)

type PostgreSQLServerManager interface {
	Ensure(ctx context.Context, obj runtime.Object) (bool, error)
	Delete(ctx context.Context, obj runtime.Object) (bool, error)
	GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error)
	convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
