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

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
)

type ResourceGroupControllerFactory struct {
	ResourceGroupManager resourcegroups.ResourceGroupManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

const ResourceGroupLogName = "ResourceGroup"
const EventRecorderName = "ResourceGroup-controller"
const ResourceGroupFinalizerBaseName = "resourcegroup.finalizers.azure.microsoft.com"

func (factory *ResourceGroupControllerFactory) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ResourceGroup{}).
		Complete(factory.create(mgr.GetClient(),
			ctrl.Log.WithName("controllers").WithName(ResourceGroupLogName),
			mgr.GetEventRecorderFor(EventRecorderName)))
}

func (factory *ResourceGroupControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder) *AzureController {
	resourceManagerClient := &ResourceGroupClient{
		ResourceGroupManager: factory.ResourceGroupManager,
	}
	return &AzureController{
		KubeClient:            kubeClient,
		Log:                   logger,
		Recorder:              recorder,
		ResourceManagerClient: resourceManagerClient,
		DefinitionManager: &ResourceGroupDefinitionManager{
			kubeClient: kubeClient,
		},
		FinalizerName:        ResourceGroupFinalizerBaseName,
		PostProvisionHandler: nil,
	}
}

type ResourceGroupDefinitionManager struct {
	kubeClient client.Client
}

func (fetcher *ResourceGroupDefinitionManager) GetThis(ctx context.Context, req ctrl.Request) (*ThisResourceDefinitions, error) {
	var instance v1alpha1.ResourceGroup
	err := fetcher.kubeClient.Get(ctx, req.NamespacedName, &instance)
	crdInfo := fetcher.getDefinition(&instance)
	return &ThisResourceDefinitions{
		Details: crdInfo,
		Updater: fetcher.getUpdater(&instance, crdInfo),
	}, err
}

func (_ *ResourceGroupDefinitionManager) GetDependencies(ctx context.Context, req ctrl.Request) (*DependencyDefinitions, error) {
	return &DependencyDefinitions{
		Dependencies: []*CustomResourceDetails{},
		Owner:        nil,
	}, nil
}

func (_ *ResourceGroupDefinitionManager) getDefinition(instance *v1alpha1.ResourceGroup) *CustomResourceDetails {
	return &CustomResourceDetails{
		ProvisionState: instance.Status.ProvisionState,
		Name:           instance.Name,
		Parameters:     instance.Spec.Parameters,
		Instance:       instance,
		BaseDefinition: &instance.ResourceBaseDefinition,
		IsBeingDeleted: !instance.ObjectMeta.DeletionTimestamp.IsZero(),
	}
}

func (_ *ResourceGroupDefinitionManager) getUpdater(instance *v1alpha1.ResourceGroup, crDetails *CustomResourceDetails) *CustomResourceUpdater {
	return &CustomResourceUpdater{
		UpdateInstance: func(state *v1alpha1.ResourceBaseDefinition) {
			instance.ResourceBaseDefinition = *state
		},
		CustomResourceDetails: crDetails,
	}
}
