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
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const storageAccountResourceFmt = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s"

type ResourceManager struct {
	Logger          logr.Logger
	Recorder        record.EventRecorder
	EventHubManager eventhubs.EventHubManager
}

func CreateResourceManagerClient(eventHubManager eventhubs.EventHubManager, logger logr.Logger, recorder record.EventRecorder) ResourceManager {
	return ResourceManager{
		Logger:          logger,
		Recorder:        recorder,
		EventHubManager: eventHubManager,
	}
}

func (client *ResourceManager) Create(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	instance, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.ApplyError, err
	}
	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	partitionCount := instance.Spec.Properties.PartitionCount
	messageRetentionInDays := instance.Spec.Properties.MessageRetentionInDays
	captureDescription := instance.Spec.Properties.CaptureDescription
	capturePtr := getCaptureDescriptionPtr(captureDescription)

	// create the eventhub
	var eventhub model.Model
	eventhub, err = client.EventHubManager.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName, messageRetentionInDays, partitionCount, capturePtr)
	if err != nil {
		client.Recorder.Event(instance, "Warning", "Failed", "Unable to create eventhub")
		return reconciler.ApplyError, errhelp.NewAzureError(err)
	}

	// create or update the authorisation rule
	authorizationRuleName := instance.Spec.AuthorizationRule.Name
	authRuleParams := getAuthRuleParameters(instance)
	_, err = client.EventHubManager.CreateOrUpdateAuthorizationRule(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName, authRuleParams)
	if err != nil {
		client.Recorder.Event(instance, "Warning", "Failed", "Unable to createorupdateauthorizationrule")
		return reconciler.ApplyError, errhelp.NewAzureError(err)
	}

	// eventhub creation is synchronous, can return succeeded straight away
	return reconciler.ApplySucceededWithStatus(&eventhub), nil
}

func (client *ResourceManager) Update(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	return reconciler.ApplyError, fmt.Errorf("Updating eventhub not currently supported")
}

func (client *ResourceManager) Verify(ctx context.Context, r reconciler.ResourceSpec) (reconciler.VerifyResponse, error) {
	instance, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.VerifyError, err
	}
	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourceGroup := instance.Spec.ResourceGroup
	authorizationRuleName := instance.Spec.AuthorizationRule.Name

	var eventhub model.Model
	eventhub, err = client.EventHubManager.GetHub(ctx, resourceGroup, eventhubNamespace, eventhubName)
	if eventhub.Response.StatusCode == http.StatusNotFound {
		return reconciler.VerifyMissing, nil
	}
	if err != nil {
		return reconciler.VerifyError, errhelp.NewAzureError(err)
	}
	if eventhub.Response.Response == nil {
		return reconciler.VerifyError, errhelp.NewAzureError(fmt.Errorf("Nil response received for eventhub get"))
	}

	_, err = client.EventHubManager.ListKeys(ctx, resourceGroup, eventhubNamespace, eventhubName, authorizationRuleName)
	if err != nil {
		client.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Verify failed. If eventhub is created, AuthKeys must be present.")
		return reconciler.VerifyError, err
	}

	return reconciler.VerifyReadyWithStatus(&eventhub), nil
}

func (client *ResourceManager) Delete(ctx context.Context, r reconciler.ResourceSpec) (reconciler.DeleteResult, error) {
	instance, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.DeleteError, err
	}
	eventhubName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup

	resp, err := client.EventHubManager.DeleteHub(ctx, resourcegroup, namespaceName, eventhubName)
	if resp.StatusCode == http.StatusNotFound {
		return reconciler.DeleteAlreadyDeleted, nil
	}
	if err != nil {
		return reconciler.DeleteError, nil
	}
	if resp.StatusCode == http.StatusOK {
		return reconciler.DeleteSucceeded, nil
	}

	// not sure - check what the other statuses are
	return reconciler.DeleteSucceeded, nil
}

func getCaptureDescriptionPtr(captureDescription azurev1alpha1.CaptureDescription) *model.CaptureDescription {
	// add capture details
	var capturePtr *model.CaptureDescription

	storage := captureDescription.Destination.StorageAccount
	storageAccountResourceID := fmt.Sprintf(storageAccountResourceFmt, config.SubscriptionID(), storage.ResourceGroup, storage.AccountName)

	if captureDescription.Enabled {
		capturePtr = &model.CaptureDescription{
			Enabled:           to.BoolPtr(true),
			Encoding:          model.Avro,
			IntervalInSeconds: &captureDescription.IntervalInSeconds,
			SizeLimitInBytes:  &captureDescription.SizeLimitInBytes,
			Destination: &model.Destination{
				Name: &captureDescription.Destination.Name,
				DestinationProperties: &model.DestinationProperties{
					StorageAccountResourceID: &storageAccountResourceID,
					BlobContainer:            &captureDescription.Destination.BlobContainer,
					ArchiveNameFormat:        &captureDescription.Destination.ArchiveNameFormat,
				},
			},
			SkipEmptyArchives: to.BoolPtr(true),
		}
	}
	return capturePtr
}

func getAuthRuleParameters(instance *azurev1alpha1.Eventhub) model.AuthorizationRule {
	accessRights := make([]model.AccessRights, len(instance.Spec.AuthorizationRule.Rights))
	for i, v := range instance.Spec.AuthorizationRule.Rights {
		accessRights[i] = model.AccessRights(v)
	}
	parameters := model.AuthorizationRule{
		AuthorizationRuleProperties: &model.AuthorizationRuleProperties{
			Rights: &accessRights,
		},
	}
	return parameters
}
