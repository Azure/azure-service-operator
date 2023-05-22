/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	storage "github.com/Azure/azure-service-operator/v2/api/devices/v1api20210702storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesExporter = &IotHubExtension{}

func (ext *IotHubExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]client.Object, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.IotHub)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *devices.IotHub", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(typedObj)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]armiothub.SharedAccessSignatureAuthorizationRule)
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := id.SubscriptionID
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var resClient *armiothub.ResourceClient
		resClient, err = armiothub.NewResourceClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new DevicesClient")
		}

		var pager *runtime.Pager[armiothub.ResourceClientListKeysResponse]
		var resp armiothub.ResourceClientListKeysResponse
		pager = resClient.NewListKeysPager(id.ResourceGroupName, typedObj.AzureName(), nil)
		for pager.More() {
			resp, err = pager.NextPage(ctx)
			addSecretsToMap(resp.Value, keys)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "failed to retreive response")
		}
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}

	}

	secretSlice, err := secretsToWrite(typedObj, keys)
	if err != nil {
		return nil, err
	}

	return secrets.SliceToClientObjectSlice(secretSlice), nil
}

func secretsSpecified(obj *storage.IotHub) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	secrets := obj.Spec.OperatorSpec.Secrets

	if secrets.IotHubOwnerPrimaryKey != nil ||
		secrets.IotHubOwnerSecondaryKey != nil ||
		secrets.ServicePrimaryKey != nil ||
		secrets.ServiceSecondaryKey != nil ||
		secrets.RegistryReadWritePrimaryKey != nil ||
		secrets.RegistryReadWriteSecondaryKey != nil ||
		secrets.RegistryReadPrimaryKey != nil ||
		secrets.RegistryReadSecondaryKey != nil ||
		secrets.DevicePrimaryKey != nil ||
		secrets.DeviceSecondaryKey != nil {
		return true
	}

	return false
}

func addSecretsToMap(keys []*armiothub.SharedAccessSignatureAuthorizationRule, result map[string]armiothub.SharedAccessSignatureAuthorizationRule) {
	for _, key := range keys {
		if key == nil || key.KeyName == nil {
			continue
		}
		result[*key.KeyName] = *key
	}
}

func secretsToWrite(obj *storage.IotHub, keys map[string]armiothub.SharedAccessSignatureAuthorizationRule) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewCollector(obj.Namespace)
	iothubowner, ok := keys["iothubowner"]
	if ok {
		collector.AddValue(operatorSpecSecrets.IotHubOwnerPrimaryKey, *iothubowner.PrimaryKey)
		collector.AddValue(operatorSpecSecrets.IotHubOwnerSecondaryKey, *iothubowner.SecondaryKey)
	}

	service, ok := keys["service"]
	if ok {
		collector.AddValue(operatorSpecSecrets.ServicePrimaryKey, *service.PrimaryKey)
		collector.AddValue(operatorSpecSecrets.ServiceSecondaryKey, *service.SecondaryKey)
	}

	device, ok := keys["device"]
	if ok {
		collector.AddValue(operatorSpecSecrets.DevicePrimaryKey, *device.PrimaryKey)
		collector.AddValue(operatorSpecSecrets.DeviceSecondaryKey, *device.SecondaryKey)
	}

	registryRead, ok := keys["registryRead"]
	if ok {
		collector.AddValue(operatorSpecSecrets.RegistryReadPrimaryKey, *registryRead.PrimaryKey)
		collector.AddValue(operatorSpecSecrets.RegistryReadSecondaryKey, *registryRead.SecondaryKey)
	}

	registryReadWrite, ok := keys["registryReadWrite"]
	if ok {
		collector.AddValue(operatorSpecSecrets.RegistryReadWritePrimaryKey, *registryReadWrite.PrimaryKey)
		collector.AddValue(operatorSpecSecrets.RegistryReadWriteSecondaryKey, *registryReadWrite.SecondaryKey)
	}

	return collector.Values()
}
