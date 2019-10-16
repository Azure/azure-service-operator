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
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
)

type ResourceGroupControllerFactory struct {
	ResourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

func (factory *ResourceGroupControllerFactory) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ResourceGroup{}).
		Complete(factory.create(mgr.GetClient(),
			ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
			mgr.GetEventRecorderFor("ResourceGroup-controller")))
}

func (factory *ResourceGroupControllerFactory) create(kubeClient client.Client, logger logr.Logger, recorder record.EventRecorder) *AzureController {
	return &AzureController{
		KubeClient: kubeClient,
		Log:        logger,
		Recorder:   recorder,
		ResourceClient: &ResourceGroupClient{
			ResourceGroupManager: factory.ResourceGroupManager,
		},
		CRDFetcher:           &ResourceGroupDefinitionFetcher{},
		FinalizerName:        "resourcegroup.finalizers.com",
		PostProvisionHandler: nil,
	}
}
