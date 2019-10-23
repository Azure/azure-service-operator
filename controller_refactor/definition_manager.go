package controller_refactor

import (
	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"

	// "k8s.io/apimachinery/pkg/api/meta"

	"k8s.io/apimachinery/pkg/runtime"
)

// DefinitionManager is used to retrieve the required custom resource definitions
// and convert them into a state that can be consumed and updated (where applicable) generically
type DefinitionManager interface {
	// returns ResourceDefinition
	GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *ResourceDefinition
	// if any dependency is not found, should return nil and a NotFound api error
	GetDependencies(ctx context.Context, thisInstance runtime.Object) (*DependencyDefinitions, error)
}

// Details of the current resource being reconciled
type ResourceDefinition struct {
	InitialInstance runtime.Object
	StatusGetter    StatusGetter
	StatusUpdater   StatusUpdater
}

type Dependency struct {
	InitialInstance runtime.Object
	NamespacedName  types.NamespacedName
	StatusGetter    StatusGetter
}

// Details of the owner and the dependencies of the resource
type DependencyDefinitions struct {
	Owner        *Dependency
	Dependencies []*Dependency
}

var NoDependencies = DependencyDefinitions{
	Dependencies: []*Dependency{},
	Owner:        nil,
}

// updates the status of the instance of runtime.Object with status
type StatusGetter = func(instance runtime.Object) (*azurev1alpha1.ResourceStatus, error)
type StatusUpdater = func(instance runtime.Object, status *azurev1alpha1.ResourceStatus) error
