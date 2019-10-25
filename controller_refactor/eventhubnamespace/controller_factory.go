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

package eventhubnamespace

import (
	"context"
	"github.com/Azure/azure-service-operator/controller_refactor"
	resourcegrouphelpers "github.com/Azure/azure-service-operator/controller_refactor/resourcegroup"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"strings"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
)

type ControllerFactory struct {
	EventHubNamespaceManager eventhubs.EventHubNamespaceManager
	Scheme                   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubnamespaces/status,verbs=get;update;patch

const ResourceKind = "EventhubNamespace"
const EventRecorderName = "EventhubNamespace-controller"
const FinalizerName = "eventhubnamespace.finalizers.azure.microsoft.com"
const ManagedResourceGroupAnnotation = "eventhubnamespace.azure.microsoft.com/managed-resource-group"

func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager, parameters controller_refactor.Parameters) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.EventhubNamespace{}).
		Complete(factory.create(mgr.GetClient(),
			ctrl.Log.WithName("controllers").WithName(ResourceKind),
			mgr.GetEventRecorderFor(EventRecorderName), parameters))
}

func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder, parameters controller_refactor.Parameters) *controller_refactor.GenericController {
	resourceManagerClient := &ResourceManagerClient{
		logger:                   logger,
		eventHubNamespaceManager: factory.EventHubNamespaceManager,
	}
	return &controller_refactor.GenericController{
		Parameters:            parameters,
		ResourceKind:          ResourceKind,
		KubeClient:            kubeClient,
		Log:                   logger,
		Recorder:              recorder,
		Scheme:                factory.Scheme,
		ResourceManagerClient: resourceManagerClient,
		DefinitionManager:     &definitionManager{},
		FinalizerName:         FinalizerName,
		PostProvisionFactory:  nil,
	}
}

type definitionManager struct{}

func (dm *definitionManager) GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *controller_refactor.ResourceDefinition {
	return &controller_refactor.ResourceDefinition{
		InitialInstance: &v1alpha1.EventhubNamespace{},
		StatusAccessor:  getStatus,
		StatusUpdater:   updateStatus,
	}
}

func (dm *definitionManager) GetDependencies(ctx context.Context, thisInstance runtime.Object) (*controller_refactor.DependencyDefinitions, error) {
	ehnInstance, err := convertInstance(thisInstance)
	if err != nil {
		return nil, err
	}

	// get the metadata annotation to check if the eventhubnamespace belongs to a resourcegroup that is
	// managed by Kubernetes, and if so, make this resourcegroup a dependency
	managedResourceGroup := ehnInstance.Annotations[ManagedResourceGroupAnnotation]

	// defaults to true
	isManaged := strings.ToLower(managedResourceGroup) != "false"

	var owner *controller_refactor.Dependency = nil
	if isManaged {
		owner = &controller_refactor.Dependency{
			InitialInstance: &v1alpha1.ResourceGroup{},
			NamespacedName: types.NamespacedName{
				Namespace: ehnInstance.Namespace,
				Name:      ehnInstance.Spec.ResourceGroup,
			},
			StatusAccessor: resourcegrouphelpers.GetStatus,
		}
	}

	return &controller_refactor.DependencyDefinitions{
		Dependencies: []*controller_refactor.Dependency{},
		Owner:        owner,
	}, nil
}
