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

// TODO

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/controller_refactor"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	"net/http"

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
)

const storageAccountResourceFmt = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s"

type resourceManagerClient struct {
	logger          logr.Logger
	recorder        record.EventRecorder
	eventHubManager eventhubs.EventHubManager
}

func (client *resourceManagerClient) Create(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	instance, err := convertInstance(r)
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	partitionCount := instance.Spec.Properties.PartitionCount
	messageRetentionInDays := instance.Spec.Properties.MessageRetentionInDays
	captureDescription := instance.Spec.Properties.CaptureDescription
	capturePtr := getCaptureDescriptionPtr(captureDescription)

	// create the eventhub
	_, err = client.eventHubManager.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName, messageRetentionInDays, partitionCount, capturePtr)
	if err != nil {
		client.recorder.Event(instance, "Warning", "Failed", "unable to create eventhub")
		return controller_refactor.EnsureFailed, errhelp.NewAzureError(err)
	}

	// create or update the authorisation rule
	authorizationRuleName := instance.Spec.AuthorizationRule.Name
	authRuleParams := getAuthRuleParameters(instance)
	_, err = client.eventHubManager.CreateOrUpdateAuthorizationRule(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName, authRuleParams)
	if err != nil {
		client.recorder.Event(instance, "Warning", "Failed", "Unable to createorupdateauthorizationrule")
		return controller_refactor.EnsureFailed, errhelp.NewAzureError(err)
	}

	// eventhub creation is synchronous, can return succeeded straight away
	return controller_refactor.EnsureSucceeded, nil
}

func (client *resourceManagerClient) Update(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	return controller_refactor.EnsureFailed, fmt.Errorf("updating eventhub not currently supported")
}

func (client *resourceManagerClient) Verify(ctx context.Context, r runtime.Object) (controller_refactor.VerifyResult, error) {
	instance, err := convertInstance(r)
	if err != nil {
		return controller_refactor.VerifyError, err
	}
	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourceGroup := instance.Spec.ResourceGroup
	authorizationRuleName := instance.Spec.AuthorizationRule.Name

	eventhub, err := client.eventHubManager.GetHub(ctx, resourceGroup, eventhubNamespace, eventhubName)
	if eventhub.Response.StatusCode == http.StatusNotFound {
		return controller_refactor.VerifyMissing, nil
	}
	if err != nil {
		return controller_refactor.VerifyError, errhelp.NewAzureError(err)
	}
	if eventhub.Response.Response == nil {
		return controller_refactor.VerifyError, errhelp.NewAzureError(fmt.Errorf("nil response received for eventhub get"))
	}

	_, err = client.eventHubManager.ListKeys(ctx, resourceGroup, eventhubNamespace, eventhubName, authorizationRuleName)
	if err != nil {
		client.recorder.Event(instance, "EventHub", "Failed", "Verify failed. If eventhub is created, AuthKeys must be present.")
		return controller_refactor.VerifyError, err
	}

	return controller_refactor.VerifyMissing, nil
}

func (client *resourceManagerClient) Delete(ctx context.Context, r runtime.Object) (controller_refactor.DeleteResult, error) {
	instance, err := convertInstance(r)
	if err != nil {
		return controller_refactor.DeleteError, err
	}
	eventhubName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup

	resp, err := client.eventHubManager.DeleteHub(ctx, resourcegroup, namespaceName, eventhubName)
	if resp.StatusCode == http.StatusNotFound {
		return controller_refactor.DeleteAlreadyDeleted, nil
	}
	if err != nil {
		return controller_refactor.DeleteError, nil
	}
	if resp.StatusCode == http.StatusOK {
		return controller_refactor.DeleteSucceed, nil
	}

	// not sure - check what the other statuses are
	return controller_refactor.DeleteSucceed, nil
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
