/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	storage "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1beta20220501storage"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.SecretsRetriever = &ConfigurationStoreExtension{}

func (ext *ConfigurationStoreExtension) RetrieveSecrets(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]*v1.Secret, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.ConfigurationStore)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *appconfiguration.ConfigurationStore", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(obj)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]armappconfiguration.APIKey)
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := armClient.SubscriptionID()
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var confClient *armappconfiguration.ConfigurationStoresClient
		confClient, err = armappconfiguration.NewConfigurationStoresClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new ConfigurationStoresClient")
		}

		var pager *runtime.Pager[armappconfiguration.ConfigurationStoresClientListKeysResponse]
		var resp armappconfiguration.ConfigurationStoresClientListKeysResponse
		pager = confClient.NewListKeysPager(id.ResourceGroupName, obj.AzureName(), nil)
		resp, err = pager.NextPage(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to retreive response")
		}
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}

		keys = secretsByName(resp.Value)
	}

	secretSlice, err := secretsToWrite(typedObj, keys)
	if err != nil {
		return nil, err
	}

	return secretSlice, nil
}

func secretsSpecified(obj *storage.ConfigurationStore) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	secrets := obj.Spec.OperatorSpec.Secrets

	if secrets.PrimaryConnectionString != nil ||
		secrets.PrimaryKeyID != nil ||
		secrets.PrimaryKey != nil ||
		secrets.SecondaryConnectionString != nil ||
		secrets.SecondaryKey != nil ||
		secrets.SecondaryKeyID != nil ||
		secrets.PrimaryReadOnlyConnectionString != nil ||
		secrets.PrimaryReadOnlyKeyID != nil ||
		secrets.PrimaryReadOnlyKey != nil ||
		secrets.SecondaryReadOnlyConnectionString != nil ||
		secrets.SecondaryReadOnlyKey != nil ||
		secrets.SecondaryReadOnlyKeyID != nil {
		return true
	}

	return false
}

func secretsByName(keys []*armappconfiguration.APIKey) map[string]armappconfiguration.APIKey {
	result := make(map[string]armappconfiguration.APIKey)

	for _, key := range keys {
		if key == nil || key.Name == nil {
			continue
		}
		result[*key.Name] = *key
	}

	return result
}

func secretsToWrite(obj *storage.ConfigurationStore, keys map[string]armappconfiguration.APIKey) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewSecretCollector(obj.Namespace)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryConnectionString, *keys["Primary"].ConnectionString)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryKeyID, *keys["Primary"].ID)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryKey, *keys["Primary"].Value)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryReadOnlyConnectionString, *keys["Primary Read Only"].ConnectionString)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryReadOnlyKeyID, *keys["Primary Read Only"].ID)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryReadOnlyKey, *keys["Primary Read Only"].Value)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryConnectionString, *keys["Secondary"].ConnectionString)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryKeyID, *keys["Secondary"].ID)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryKey, *keys["Secondary"].Value)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryReadOnlyConnectionString, *keys["Secondary Read Only"].ConnectionString)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryReadOnlyKeyID, *keys["Secondary Read Only"].ID)
	collector.AddSecretValue(operatorSpecSecrets.SecondaryReadOnlyKey, *keys["Secondary Read Only"].Value)

	return collector.Secrets(), nil
}
