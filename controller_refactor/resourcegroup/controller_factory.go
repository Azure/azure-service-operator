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

package resourcegroup

import (
	"context"
	"github.com/Azure/azure-service-operator/controller_refactor"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
)

type ControllerFactory struct {
	ResourceGroupManager resourcegroups.ResourceGroupManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

const LogName = "ResourceGroup"
const EventRecorderName = "ResourceGroup-controller"
const FinalizerName = "resourcegroup.finalizers.azure.microsoft.com"

func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager, parameters controller_refactor.Parameters) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ResourceGroup{}).
		Complete(factory.create(mgr.GetClient(),
			ctrl.Log.WithName("controllers").WithName(LogName),
			mgr.GetEventRecorderFor(EventRecorderName), parameters))
}

func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder, parameters controller_refactor.Parameters) *controller_refactor.GenericController {
	resourceManagerClient := &ResourceManagerClient{
		resourceGroupManager: factory.ResourceGroupManager,
	}
	return &controller_refactor.GenericController{
		Parameters:            parameters,
		KubeClient:            kubeClient,
		Log:                   logger,
		Recorder:              recorder,
		ResourceManagerClient: resourceManagerClient,
		DefinitionManager:     &definitionManager{},
		FinalizerName:         FinalizerName,
		PostProvisionHandler:  nil,
	}
}

type definitionManager struct{}

func (dm *definitionManager) GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *controller_refactor.ResourceDefinition {
	return &controller_refactor.ResourceDefinition{
		InitialInstance: &v1alpha1.ResourceGroup{},
		StatusGetter:    GetStatus,
		StatusUpdater:   updateStatus,
	}
}

func (dm *definitionManager) GetDependencies(context.Context, runtime.Object) (*controller_refactor.DependencyDefinitions, error) {
	return &controller_refactor.NoDependencies, nil
}
