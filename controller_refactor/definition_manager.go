/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	// returns a ResourceDefinition
	GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *ResourceDefinition
	// returns the dependencies for a resource
	GetDependencies(ctx context.Context, thisInstance runtime.Object) (*DependencyDefinitions, error)
}

// Details of the current resource being reconciled
type ResourceDefinition struct {
	// This can be an empty resource definition object of the required Kind
	InitialInstance runtime.Object
	StatusAccessor  StatusAccessor
	StatusUpdater   StatusUpdater
}

// The information required to pull the resource definition of a dependency from kubernetes
type Dependency struct {
	InitialInstance runtime.Object
	NamespacedName  types.NamespacedName
	StatusAccessor  StatusAccessor
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

// fetches the status of the instance of runtime.Object
type StatusAccessor = func(instance runtime.Object) (*azurev1alpha1.ASOStatus, error)

// updates the status of the instance of runtime.Object with status
type StatusUpdater = func(instance runtime.Object, status *azurev1alpha1.ASOStatus) error
