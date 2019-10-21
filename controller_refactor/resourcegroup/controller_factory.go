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

func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ResourceGroup{}).
		Complete(factory.create(mgr.GetClient(),
			ctrl.Log.WithName("controllers").WithName(LogName),
			mgr.GetEventRecorderFor(EventRecorderName)))
}

func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder) *controller_refactor.AzureController {
	resourceManagerClient := &ResourceManagerClient{
		ResourceGroupManager: factory.ResourceGroupManager,
	}
	return &controller_refactor.AzureController{
		KubeClient:            kubeClient,
		Log:                   logger,
		Recorder:              recorder,
		ResourceManagerClient: resourceManagerClient,
		DefinitionManager: &DefinitionManager{
			kubeClient: kubeClient,
		},
		FinalizerName:        FinalizerName,
		PostProvisionHandler: nil,
	}
}

type DefinitionManager struct {
	kubeClient client.Client
}

func (fetcher *DefinitionManager) GetThis(ctx context.Context, req ctrl.Request) (*controller_refactor.ThisResourceDefinitions, error) {
	var instance v1alpha1.ResourceGroup
	err := fetcher.kubeClient.Get(ctx, req.NamespacedName, &instance)
	crdInfo := fetcher.getDefinition(&instance)
	return &controller_refactor.ThisResourceDefinitions{
		Details: crdInfo,
		Updater: fetcher.getUpdater(&instance, crdInfo),
	}, err
}

func (_ *DefinitionManager) GetDependencies(ctx context.Context, req ctrl.Request) (*controller_refactor.DependencyDefinitions, error) {
	return &controller_refactor.DependencyDefinitions{
		Dependencies: []*controller_refactor.CustomResourceDetails{},
		Owner:        nil,
	}, nil
}

func (_ *DefinitionManager) getDefinition(instance *v1alpha1.ResourceGroup) *controller_refactor.CustomResourceDetails {
	return &controller_refactor.CustomResourceDetails{
		ProvisionState: instance.Status.ProvisionState,
		Name:           instance.Name,
		Parameters:     instance.Spec.Parameters,
		Instance:       instance,
		BaseDefinition: &instance.ResourceBaseDefinition,
		IsBeingDeleted: !instance.ObjectMeta.DeletionTimestamp.IsZero(),
	}
}

func (_ *DefinitionManager) getUpdater(instance *v1alpha1.ResourceGroup, crDetails *controller_refactor.CustomResourceDetails) *controller_refactor.CustomResourceUpdater {
	return &controller_refactor.CustomResourceUpdater{
		UpdateInstance: func(state *v1alpha1.ResourceBaseDefinition) {
			instance.ResourceBaseDefinition = *state
		},
		CustomResourceDetails: crDetails,
	}
}
