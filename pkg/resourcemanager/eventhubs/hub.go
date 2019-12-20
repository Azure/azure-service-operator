/*
Copyright 2019 microsoft.

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

package eventhubs

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/types"
)

type azureEventHubManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func getHubsClient() eventhub.EventHubsClient {
	hubClient := eventhub.NewEventHubsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	hubClient.Authorizer = auth
	hubClient.AddToUserAgent(config.UserAgent())
	return hubClient
}

func NewEventhubClient(secretClient secrets.SecretClient, scheme *runtime.Scheme) *azureEventHubManager {
	return &azureEventHubManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func (_ *azureEventHubManager) DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error) {
	hubClient := getHubsClient()
	return hubClient.Delete(ctx,
		resourceGroupName,
		namespaceName,
		eventHubName)

}

func (_ *azureEventHubManager) CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, MessageRetentionInDays int32, PartitionCount int32, captureDescription *eventhub.CaptureDescription) (eventhub.Model, error) {
	hubClient := getHubsClient()

	// MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	if MessageRetentionInDays < 1 || MessageRetentionInDays > 7 {
		return eventhub.Model{}, fmt.Errorf("MessageRetentionInDays is invalid")
	}

	// PartitionCount - Number of partitions created for the Event Hub, allowed values are from 2 to 32 partitions.
	if PartitionCount < 2 || PartitionCount > 32 {
		return eventhub.Model{}, fmt.Errorf("PartitionCount is invalid")
	}

	properties := eventhub.Properties{
		PartitionCount:         to.Int64Ptr(int64(PartitionCount)),
		MessageRetentionInDays: to.Int64Ptr(int64(MessageRetentionInDays)),
		CaptureDescription:     captureDescription,
	}

	return hubClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		eventhub.Model{
			Properties: &properties,
		},
	)
}

func (_ *azureEventHubManager) GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.Get(ctx, resourceGroupName, namespaceName, eventHubName)
}

func (_ *azureEventHubManager) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdateAuthorizationRule(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters)
}

func (_ *azureEventHubManager) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AccessKeys, err error) {
	hubClient := getHubsClient()
	return hubClient.ListKeys(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName)
}

func (e *azureEventHubManager) createOrUpdateAccessPolicyEventHub(resourcegroup string, eventhubNamespace string, eventhubName string, instance *azurev1alpha1.Eventhub) error {

	var err error
	ctx := context.Background()

	authorizationRuleName := instance.Spec.AuthorizationRule.Name
	accessRights := make([]model.AccessRights, len(instance.Spec.AuthorizationRule.Rights))
	for i, v := range instance.Spec.AuthorizationRule.Rights {
		accessRights[i] = model.AccessRights(v)
	}
	//accessRights := r.toAccessRights(instance.Spec.AuthorizationRule.Rights)
	parameters := model.AuthorizationRule{
		AuthorizationRuleProperties: &model.AuthorizationRuleProperties{
			Rights: &accessRights,
		},
	}
	_, err = e.CreateOrUpdateAuthorizationRule(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName, parameters)
	if err != nil {
		return err
	}
	return nil
}

func (e *azureEventHubManager) createEventhubSecrets(ctx context.Context, secretName string, instance *azurev1alpha1.Eventhub, data map[string][]byte) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	err := e.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(e.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *azureEventHubManager) listAccessKeysAndCreateSecrets(resourcegroup string, eventhubNamespace string, eventhubName string, secretName string, authorizationRuleName string, instance *azurev1alpha1.Eventhub) error {

	var err error
	var result model.AccessKeys
	ctx := context.Background()

	result, err = e.ListKeys(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName)
	if err != nil {
		//log error and kill it
		return err
	} else {
		//create secrets in the k8s with the listed keys
		data := map[string][]byte{
			"primaryconnectionstring":   []byte(*result.PrimaryConnectionString),
			"secondaryconnectionstring": []byte(*result.SecondaryConnectionString),
			"primaryKey":                []byte(*result.PrimaryKey),
			"secondaryKey":              []byte(*result.SecondaryKey),
			"sharedaccesskey":           []byte(authorizationRuleName),
			"eventhubnamespace":         []byte(eventhubNamespace),
			"eventhubName":              []byte(eventhubName),
		}
		err = e.createEventhubSecrets(
			ctx,
			secretName,
			instance,
			data,
		)
		if err != nil {
			return err
		}
	}
	return nil

}

func (e *azureEventHubManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := e.convert(obj)
	if err != nil {
		return false, err
	}

	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	partitionCount := instance.Spec.Properties.PartitionCount
	messageRetentionInDays := instance.Spec.Properties.MessageRetentionInDays
	captureDescription := instance.Spec.Properties.CaptureDescription
	secretName := instance.Spec.SecretName

	if len(secretName) == 0 {
		secretName = eventhubName
		instance.Spec.SecretName = eventhubName
	}

	// write information back to instance
	instance.Status.Provisioning = true

	capturePtr := getCaptureDescriptionPtr(captureDescription)

	resp, err := e.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName, messageRetentionInDays, partitionCount, capturePtr)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()

		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			if azerr.Type != errhelp.AsyncOpIncompleteError {
				instance.Status.Provisioning = false
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		} else {
			// reconciliation not done and we don't know what happened
			instance.Status.Provisioning = false
			return false, err
		}
	}

	err = e.createOrUpdateAccessPolicyEventHub(resourcegroup, eventhubNamespace, eventhubName, instance)
	if err != nil {
		return false, err
	}

	err = e.listAccessKeysAndCreateSecrets(resourcegroup, eventhubNamespace, eventhubName, secretName, instance.Spec.AuthorizationRule.Name, instance)
	if err != nil {
		return false, err
	}

	// write information back to instance
	instance.Status.State = string(resp.Status)
	instance.Status.Message = "Success"
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	// reconciliation done and everything looks ok
	return true, nil
}

func (e *azureEventHubManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := e.convert(obj)
	if err != nil {
		return true, err
	}

	eventhubName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup

	resp, err := e.DeleteHub(ctx, resourcegroup, namespaceName, eventhubName)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = err.Error()
			return false, nil
		}
		return false, err
	}

	if resp.StatusCode == http.StatusNoContent {
		return false, nil
	}

	return true, nil
}

func (e *azureEventHubManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := e.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Namespace,
			},
			Target: &azurev1alpha1.EventhubNamespace{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil

}

func (e *azureEventHubManager) convert(obj runtime.Object) (*azurev1alpha1.Eventhub, error) {
	local, ok := obj.(*azurev1alpha1.Eventhub)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

const storageAccountResourceFmt = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s"

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
