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

//import (
//	"context"
//	"github.com/Azure/azure-service-operator/controllers_new/shared"
//
//	"github.com/Azure/azure-service-operator/pkg/reconciler"
//	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apimachinery/pkg/types"
//
//	"k8s.io/client-go/tools/record"
//	ctrl "sigs.k8s.io/controller-runtime"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//
//	"github.com/Azure/azure-service-operator/api/v1alpha1"
//	"github.com/go-logr/logr"
//)

//type ControllerFactory struct {
//	ClientCreator        func(resourcegroups.ResourceGroupManager, logr.Logger) ResourceManager
//	ResourceGroupManager resourcegroups.ResourceGroupManager
//	Scheme               *runtime.Scheme
//}
//
//// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
//// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch
//
//const ResourceKind = "ResourceGroup"
//const FinalizerName = "resourcegroup.finalizers.com"
//
//func (factory *ControllerFactory) SetupWithManager(mgr ctrl.Manager, parameters reconciler.ReconcileParameters) error {
//	gc, err := factory.create(mgr.GetClient(),
//		ctrl.Log.WithName("controllers").WithName(ResourceKind),
//		mgr.GetEventRecorderFor(ResourceKind+"-controller"), parameters)
//	if err != nil {
//		return err
//	}
//	return ctrl.NewControllerManagedBy(mgr).
//		For(&v1alpha1.ResourceGroup{}).
//		Complete(gc)
//}
//
//func (factory *ControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder, parameters reconciler.ReconcileParameters) (*reconciler.GenericController, error) {
//	resourceManagerClient := factory.ClientCreator(factory.ResourceGroupManager, logger)
//	return reconciler.CreateGenericFactory(parameters, ResourceKind, kubeClient, logger, recorder, factory.Scheme, &resourceManagerClient, &definitionManager{}, FinalizerName, shared.AnnotationBaseName, nil)
//}
//
//type definitionManager struct{}
//
//func (dm *definitionManager) GetDefinition(ctx context.Context, namespacedName types.NamespacedName) *reconciler.ResourceDefinition {
//	return &reconciler.ResourceDefinition{
//		InitialInstance: &v1alpha1.ResourceGroup{},
//		StatusAccessor:  GetStatus,
//		StatusUpdater:   updateStatus,
//	}
//}
//
//func (dm *definitionManager) GetDependencies(context.Context, runtime.Object) (*reconciler.DependencyDefinitions, error) {
//	return &reconciler.NoDependencies, nil
//}
