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

package reconciler

import (
	"context"

	"k8s.io/apimachinery/pkg/types"

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
	// A function to get the status of the kubernetes object
	StatusAccessor StatusAccessor
	// A function to update the status of a kubernetes object instance
	StatusUpdater StatusUpdater
}

// The information required to pull the resource definition of a dependency from kubernetes
// and determine whether it has been successfully applied or not
type Dependency struct {
	// This can be an empty resource definition object of the required Kind
	InitialInstance runtime.Object
	NamespacedName  types.NamespacedName
	// A function to return whether the object has been successfully applied. The current object will only
	// continue once this returns true for all dependencies
	SucceededAccessor SucceededAccessor
}

// Details of the owner and the dependencies of the resource
type DependencyDefinitions struct {
	Owner        *Dependency
	Dependencies []*Dependency
}

// A shortcut for objects with no dependencies
var NoDependencies = DependencyDefinitions{
	Dependencies: []*Dependency{},
	Owner:        nil,
}

// fetches the status of the instance of runtime.Object
type StatusAccessor = func(instance runtime.Object) (*Status, error)

// updates the status of the instance of runtime.Object with status
type StatusUpdater = func(instance runtime.Object, status *Status) error

// unfortunately if doesn't seem possible to make this an extension method with a receiver
// unfortunately if doesn't seem possible to make this an extension method with a receiver
func AsSuccessAccessor(s StatusAccessor) SucceededAccessor {
	return func(instance runtime.Object) (bool, error) {
		status, err := s(instance)
		if err != nil {
			return false, err
		}
		if status == nil {
			return false, nil
		}
		return status.IsSucceeded(), nil
	}
}

// fetches a boolean flag to indicate whether the resource is in a succeeded (i.e. ready) state
type SucceededAccessor = func(instance runtime.Object) (bool, error)
