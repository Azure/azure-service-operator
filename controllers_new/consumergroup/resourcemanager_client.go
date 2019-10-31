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
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
)

type ResourceManagerClient struct {
	Logger               logr.Logger
	ConsumerGroupManager eventhubs.ConsumerGroupManager
}

func CreateResourceManagerClient(consumerGroupManager eventhubs.ConsumerGroupManager, logger logr.Logger) ResourceManagerClient {
	return ResourceManagerClient{
		Logger:               logger,
		ConsumerGroupManager: consumerGroupManager,
	}
}

func (client *ResourceManagerClient) Create(ctx context.Context, r reconciler.ResourceSpec) (reconciler.EnsureResponse, error) {
	cg, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.EnsureError, err
	}
	spec := cg.Spec

	client.Logger.Info("ConsumerGroup " + cg.Name + " creating on Azure. Please be patient.")
	_, err = client.ConsumerGroupManager.CreateConsumerGroup(ctx, spec.ResourceGroupName, spec.NamespaceName, spec.EventhubName, getAzureName(cg))
	client.Logger.Info("ConsumerGroup " + cg.Name + " finished creating on Azure.")
	if err != nil {
		return reconciler.EnsureError, err
	}
	return reconciler.EnsureSucceeded, nil
}

func (client *ResourceManagerClient) Update(ctx context.Context, r reconciler.ResourceSpec) (reconciler.EnsureResponse, error) {
	return reconciler.EnsureError, fmt.Errorf("ConsumerGroup updating not supported")
}

func (client *ResourceManagerClient) Verify(ctx context.Context, r reconciler.ResourceSpec) (reconciler.VerifyResponse, error) {
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

func (client *ResourceManagerClient) Delete(ctx context.Context, r reconciler.ResourceSpec) (reconciler.DeleteResult, error) {
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
		return reconciler.DeleteSucceed, nil
	}

	// TODO: handle all other cases
	return reconciler.DeleteSucceed, nil
}

func getAzureName(cg *v1alpha1.ConsumerGroup) string {
	azureName := cg.Spec.AzureConsumerGroupName
	if azureName == "" {
		azureName = cg.Name
	}
	return azureName
}
