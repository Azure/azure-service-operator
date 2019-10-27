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

package consumergroup

import (
	"context"
	"github.com/Azure/azure-service-operator/controller_refactor"
	eventhubaccessors "github.com/Azure/azure-service-operator/controller_refactor/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"
)

type ControllerFactory struct {
	ClientCreator        func(eventhubs.ConsumerGroupManager, logr.Logger) ResourceManagerClient
	ConsumerGroupManager eventhubs.ConsumerGroupManager
	Scheme               *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=consumergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=consumergroups/status,verbs=get;update;patch

const ResourceKind = "ConsumerGroup"
const FinalizerName = "consumergroup.finalizers.azure.microsoft.com"
const ManagedEventhubAnnotation = "consumergroup.azure.microsoft.com/managed-eventhub"

func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager, parameters controller_refactor.ReconcileParameters) error {
	gc, err := factory.create(mgr.GetClient(),
		ctrl.Log.WithName("controllers").WithName(ResourceKind),
		mgr.GetEventRecorderFor(ResourceKind+"-controller"), parameters)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ConsumerGroup{}).
		Complete(gc)
}

func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder, parameters controller_refactor.ReconcileParameters) (*controller_refactor.GenericController, error) {
	resourceManagerClient := factory.ClientCreator(factory.ConsumerGroupManager, logger)
	return controller_refactor.CreateGenericFactory(parameters, ResourceKind, kubeClient, logger, recorder, factory.Scheme, &resourceManagerClient, &definitionManager{}, FinalizerName, nil)
}

type definitionManager struct{}

func (dm *definitionManager) GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *controller_refactor.ResourceDefinition {
	return &controller_refactor.ResourceDefinition{
		InitialInstance: &v1alpha1.ConsumerGroup{},
		StatusAccessor:  getStatus,
		StatusUpdater:   updateStatus,
	}
}

func (dm *definitionManager) GetDependencies(ctx context.Context, thisInstance runtime.Object) (*controller_refactor.DependencyDefinitions, error) {
	ehnInstance, err := convertInstance(thisInstance)
	if err != nil {
		return nil, err
	}

	managedEventhub := ehnInstance.Annotations[ManagedEventhubAnnotation]
	// defaults to true
	isManaged := strings.ToLower(managedEventhub) != "false"

	var owner *controller_refactor.Dependency = nil
	if isManaged {
		owner = &controller_refactor.Dependency{
			InitialInstance: &v1alpha1.Eventhub{},
			NamespacedName: types.NamespacedName{
				Namespace: ehnInstance.Namespace,
				Name:      ehnInstance.Spec.EventhubName,
			},
			StatusAccessor: eventhubaccessors.GetStatus,
		}
	}

	return &controller_refactor.DependencyDefinitions{
		Dependencies: []*controller_refactor.Dependency{},
		Owner:        owner,
	}, nil

}
