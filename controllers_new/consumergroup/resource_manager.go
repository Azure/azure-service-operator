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
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
)

type ResourceManager struct {
	Logger               logr.Logger
	ConsumerGroupManager eventhubs.ConsumerGroupManager
}

func CreateResourceManagerClient(consumerGroupManager eventhubs.ConsumerGroupManager, logger logr.Logger) ResourceManager {
	return ResourceManager{
		Logger:               logger,
		ConsumerGroupManager: consumerGroupManager,
	}
}

func (client *ResourceManager) Create(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	cg, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.ApplyError, err
	}
	spec := cg.Spec

	client.Logger.Info("ConsumerGroup " + cg.Name + " creating on Azure. Please be patient.")
	_, err = client.ConsumerGroupManager.CreateConsumerGroup(ctx, spec.ResourceGroupName, spec.NamespaceName, spec.EventhubName, getAzureName(cg))
	client.Logger.Info("ConsumerGroup " + cg.Name + " finished creating on Azure.")
	if err != nil {
		return reconciler.ApplyError, err
	}
	return reconciler.ApplySucceeded, nil
}

func (client *ResourceManager) Update(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	return reconciler.ApplyError, fmt.Errorf("ConsumerGroup updating not supported")
}

func (client *ResourceManager) Verify(ctx context.Context, r reconciler.ResourceSpec) (reconciler.VerifyResponse, error) {
	cg, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.VerifyError, err
	}
	spec := cg.Spec

	client.Logger.Info("Fetching ConsumerGroup " + cg.Name + " from Azure.")
	consumerGroup, err := client.ConsumerGroupManager.GetConsumerGroup(ctx, spec.ResourceGroupName, spec.NamespaceName, spec.EventhubName, getAzureName(cg))
	if consumerGroup.Response.Response == nil {
		return reconciler.VerifyError, fmt.Errorf("ConsumerGroup verify was nil for %s", cg.Name)
	} else if consumerGroup.Response.StatusCode == http.StatusNotFound {
		return reconciler.VerifyMissing, nil
	} else if err != nil {
		return reconciler.VerifyError, err
	} else if consumerGroup.Response.StatusCode == http.StatusOK {
		return reconciler.VerifyReady, nil
	}

	// we ideally shouldn't get to this point - all cases should be handled explicitly
	return reconciler.VerifyMissing, nil
}

func (client *ResourceManager) Delete(ctx context.Context, r reconciler.ResourceSpec) (reconciler.DeleteResult, error) {
	cg, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.DeleteError, err
	}
	spec := cg.Spec

	client.Logger.Info("ConsumerGroup " + cg.Name + " deleting on Azure.")
	resp, err := client.ConsumerGroupManager.DeleteConsumerGroup(ctx, spec.ResourceGroupName, spec.NamespaceName, spec.EventhubName, getAzureName(cg))
	if resp.Response == nil {
		return reconciler.DeleteError, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return reconciler.DeleteAlreadyDeleted, nil
	} else if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		return reconciler.DeleteSucceeded, nil
	}

	// TODO: handle all other cases
	return reconciler.DeleteSucceeded, nil
}

func getAzureName(cg *v1alpha1.ConsumerGroup) string {
	azureName := cg.Spec.AzureConsumerGroupName
	if azureName == "" {
		azureName = cg.Name
	}
	return azureName
}
