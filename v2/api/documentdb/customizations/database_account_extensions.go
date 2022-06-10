/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.SecretsRetriever = &DatabaseAccountExtension{}

func (ext *DatabaseAccountExtension) RetrieveSecrets(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]*v1.Secret, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*documentdb.DatabaseAccount)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *documentdb.DatabaseAccount", obj)
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

	var keys armcosmos.DatabaseAccountListKeysResult
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := armClient.SubscriptionID()
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var acctClient *armcosmos.DatabaseAccountsClient
		acctClient, err = armcosmos.NewDatabaseAccountsClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new DatabaseAccountClient")
		}

		// TODO: There is a ListReadOnlyKeys API that requires less permissions. We should consider determining
		// TODO: that we don't need to call the ListKeys API and install call the listReadOnlyKeys API.
		var resp armcosmos.DatabaseAccountsClientListKeysResponse
		resp, err = acctClient.ListKeys(ctx, id.ResourceGroupName, obj.AzureName(), nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}

		keys = resp.DatabaseAccountListKeysResult
	}

	secretSlice, err := secretsToWrite(typedObj, keys)
	if err != nil {
		return nil, err
	}

	return secretSlice, nil
}

func secretsSpecified(obj *documentdb.DatabaseAccount) (bool, bool) {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false, false
	}

	specSecrets := obj.Spec.OperatorSpec.Secrets
	hasSecrets := false
	hasEndpoints := false
	if specSecrets.PrimaryMasterKey != nil ||
		specSecrets.SecondaryMasterKey != nil ||
		specSecrets.PrimaryReadonlyMasterKey != nil ||
		specSecrets.SecondaryReadonlyMasterKey != nil {
		hasSecrets = true
	}

	if specSecrets.DocumentEndpoint != nil {
		hasEndpoints = true
	}

	return hasSecrets, hasEndpoints
}

func secretsToWrite(obj *documentdb.DatabaseAccount, accessKeys armcosmos.DatabaseAccountListKeysResult) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewSecretCollector(obj.Namespace)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryMasterKey, to.String(accessKeys.PrimaryMasterKey))
	collector.AddSecretValue(operatorSpecSecrets.SecondaryMasterKey, to.String(accessKeys.SecondaryMasterKey))
	collector.AddSecretValue(operatorSpecSecrets.PrimaryReadonlyMasterKey, to.String(accessKeys.PrimaryReadonlyMasterKey))
	collector.AddSecretValue(operatorSpecSecrets.SecondaryReadonlyMasterKey, to.String(accessKeys.SecondaryReadonlyMasterKey))
	collector.AddSecretValue(operatorSpecSecrets.DocumentEndpoint, to.String(obj.Status.DocumentEndpoint))

	return collector.Secrets(), nil
}
