/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.SecretsRetriever = &StorageAccountExtension{}

func (ext *StorageAccountExtension) RetrieveSecrets(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]*v1.Secret, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.StorageAccount)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *storage.StorageAccount", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets, hasEndpoints := secretsSpecified(typedObj)
	if !hasSecrets && !hasEndpoints {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(obj)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]string)
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := armClient.SubscriptionID()
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		acctClient := armstorage.NewAccountsClient(subscription, armClient.Creds(), armClient.ClientOptions())

		var resp armstorage.AccountsClientListKeysResponse
		resp, err = acctClient.ListKeys(ctx, id.ResourceGroupName, obj.AzureName(), nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}

		keys = secretsByName(resp.Keys)
	}

	secretSlice, err := secretsToWrite(typedObj, keys)
	if err != nil {
		return nil, err
	}

	return secretSlice, nil
}

func secretsSpecified(obj *storage.StorageAccount) (bool, bool) {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false, false
	}

	hasSecrets := false
	hasEndpoints := false
	secrets := obj.Spec.OperatorSpec.Secrets
	if secrets.Key1 != nil || secrets.Key2 != nil {
		hasSecrets = true
	}

	if secrets.BlobEndpoint != nil ||
		secrets.QueueEndpoint != nil ||
		secrets.TableEndpoint != nil ||
		secrets.FileEndpoint != nil ||
		secrets.WebEndpoint != nil ||
		secrets.DfsEndpoint != nil {
		hasEndpoints = true
	}

	return hasSecrets, hasEndpoints
}

func secretsByName(keys []*armstorage.AccountKey) map[string]string {
	result := make(map[string]string)

	for _, key := range keys {
		if key.KeyName == nil || key.Value == nil {
			continue
		}
		result[*key.KeyName] = *key.Value
	}

	return result
}

func secretsToWrite(obj *storage.StorageAccount, keys map[string]string) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewSecretCollector(obj.Namespace)
	collector.AddSecretValue(operatorSpecSecrets.Key1, keys["key1"])
	collector.AddSecretValue(operatorSpecSecrets.Key2, keys["key2"])
	// There are tons of different endpoints we could write, including secondary endpoints.
	// For now we're just exposing the main ones. See:
	// https://docs.microsoft.com/en-us/rest/api/storagerp/storage-accounts/get-properties for more details
	if obj.Status.PrimaryEndpoints != nil {
		eps := obj.Status.PrimaryEndpoints
		collector.AddSecretValue(operatorSpecSecrets.BlobEndpoint, to.String(eps.Blob))
		collector.AddSecretValue(operatorSpecSecrets.QueueEndpoint, to.String(eps.Queue))
		collector.AddSecretValue(operatorSpecSecrets.TableEndpoint, to.String(eps.Table))
		collector.AddSecretValue(operatorSpecSecrets.FileEndpoint, to.String(eps.File))
		collector.AddSecretValue(operatorSpecSecrets.WebEndpoint, to.String(eps.Web))
		collector.AddSecretValue(operatorSpecSecrets.DfsEndpoint, to.String(eps.Dfs))
	}

	return collector.Secrets(), nil
}
