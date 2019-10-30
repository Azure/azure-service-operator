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

package eventhub

import (
	"context"
	"github.com/Azure/azure-service-operator/controllers_new/shared"
	"strings"

	eventhubnamespaceaccessors "github.com/Azure/azure-service-operator/controllers_new/eventhubnamespace"
	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"
)

type ControllerFactory struct {
	ClientCreator   func(eventhubs.EventHubManager, logr.Logger, record.EventRecorder) ResourceManagerClient
	EventHubManager eventhubs.EventHubManager
	Scheme          *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs/status,verbs=get;update;patch

const ResourceKind = "Eventhub"
const FinalizerName = "eventhub.finalizers.azure.microsoft.com"

func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager, parameters reconciler.ReconcileParameters) error {
	gc, err := factory.create(mgr.GetClient(),
		ctrl.Log.WithName("controllers").WithName(ResourceKind),
		mgr.GetEventRecorderFor(ResourceKind+"-controller"), parameters)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Eventhub{}).
		Complete(gc)
}

func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder, parameters reconciler.ReconcileParameters) (*reconciler.GenericController, error) {
	resourceManagerClient := factory.ClientCreator(factory.EventHubManager, logger, recorder)

	// create a PostProvisionHandler for writing secret
	secretsWriterFactory := func(c *reconciler.GenericController) reconciler.PostProvisionHandler {
		return &secretsWriter{
			GenericController: c,
			eventHubManager:   resourceManagerClient.EventHubManager,
		}
	}

	return reconciler.CreateGenericFactory(parameters, ResourceKind, kubeClient, logger, recorder, factory.Scheme, &resourceManagerClient, &definitionManager{}, FinalizerName, secretsWriterFactory)
}

type definitionManager struct{}

func (dm *definitionManager) GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *reconciler.ResourceDefinition {
	return &reconciler.ResourceDefinition{
		InitialInstance: &v1alpha1.Eventhub{},
		StatusAccessor:  GetStatus,
		StatusUpdater:   updateStatus,
	}
}

func (dm *definitionManager) GetDependencies(ctx context.Context, thisInstance runtime.Object) (*reconciler.DependencyDefinitions, error) {
	ehnInstance, err := convertInstance(thisInstance)
	if err != nil {
		return nil, err
	}

	managedEventhubNamespace := ehnInstance.Annotations[shared.ManagedParentAnnotation]

	// defaults to true
	isManaged := strings.ToLower(managedEventhubNamespace) != "false"

	var owner *reconciler.Dependency = nil

	if isManaged {
		owner = &reconciler.Dependency{
			InitialInstance: &v1alpha1.EventhubNamespace{},
			NamespacedName: types.NamespacedName{
				Namespace: ehnInstance.Namespace,
				Name:      ehnInstance.Spec.Namespace,
			},
			StatusAccessor: eventhubnamespaceaccessors.GetStatus,
		}
	}

	return &reconciler.DependencyDefinitions{
		Dependencies: []*reconciler.Dependency{},
		Owner:        owner,
	}, nil
}
