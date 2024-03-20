/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesExporter = &NamespacesAuthorizationRuleExtension{}

func (ext *NamespacesAuthorizationRuleExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) ([]client.Object, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.NamespacesAuthorizationRule)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *eventhub.NamespacesAuthorizationRule", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := namespacesAuthorizationRuleSecretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(typedObj)
	if err != nil {
		return nil, err
	}

	// Only bother calling ListKeys if there are secrets to retrieve
	var res armeventhub.NamespacesClientListKeysResponse
	if hasSecrets {
		subscription := id.SubscriptionID
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		var confClient *armeventhub.NamespacesClient
		confClient, err = armeventhub.NewNamespacesClient(subscription, armClient.Creds(), armClient.ClientOptions())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new NamespaceClient")
		}

		res, err = confClient.ListKeys(ctx, id.ResourceGroupName, id.Parent.Name, typedObj.AzureName(), nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to retreive response")
		}

	}

	secretSlice, err := namespacesAuthorizationRuleSecretsToWrite(typedObj, res.AccessKeys)
	if err != nil {
		return nil, err
	}

	return secrets.SliceToClientObjectSlice(secretSlice), nil
}

func namespacesAuthorizationRuleSecretsSpecified(obj *storage.NamespacesAuthorizationRule) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	secrets := obj.Spec.OperatorSpec.Secrets

	if secrets.PrimaryKey != nil ||
		secrets.SecondaryKey != nil ||
		secrets.PrimaryConnectionString != nil ||
		secrets.SecondaryConnectionString != nil {
		return true
	}

	return false
}

func namespacesAuthorizationRuleSecretsToWrite(obj *storage.NamespacesAuthorizationRule, keys armeventhub.AccessKeys) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewCollector(obj.Namespace)

	collector.AddValue(operatorSpecSecrets.PrimaryKey, *keys.PrimaryKey)
	collector.AddValue(operatorSpecSecrets.SecondaryKey, *keys.SecondaryKey)
	collector.AddValue(operatorSpecSecrets.PrimaryConnectionString, *keys.PrimaryConnectionString)
	collector.AddValue(operatorSpecSecrets.SecondaryConnectionString, *keys.SecondaryConnectionString)

	return collector.Values()
}
