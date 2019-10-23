package controller_refactor

import (
	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/meta"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

// DefinitionManager is used to retrieve the required custom resource definitions
// and convert them into a state that can be consumed and updated (where applicable) generically
type DefinitionManager interface {
	// returns ThisResourceDefinitions
	GetThis(ctx context.Context, req ctrl.Request) (*ThisResourceDefinitions, error)
	// if any dependency is not found, should return nil and a NotFound api error
	GetDependencies(ctx context.Context, thisInstance runtime.Object) (*DependencyDefinitions, error)
}

// Details of the current resource being reconciled
type ThisResourceDefinitions struct {
	Details       *CustomResourceDetails
	StatusUpdater StatusUpdater
}

// Details of the owner and the dependencies of the resource
type DependencyDefinitions struct {
	Owner        *CustomResourceDetails
	Dependencies []*CustomResourceDetails
}

var NoDependencies = DependencyDefinitions{
	Dependencies: []*CustomResourceDetails{},
	Owner:        nil,
}

type CustomResourceDetails struct {
	Instance runtime.Object
	Status   *azurev1alpha1.ResourceStatus
}

// updates the status of the instance of runtime.Object with status
type StatusUpdater = func(instance runtime.Object, status azurev1alpha1.ResourceStatus) error
