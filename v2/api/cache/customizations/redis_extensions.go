/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	redis "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.SecretsRetriever = &RedisExtension{}

func (ext *RedisExtension) RetrieveSecrets(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]*v1.Secret, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*redis.Redis)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *redis.Redis", obj)
	}

	// Type assert that we are the hub type. This should fail to compile if
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

	var accessKeys armredis.AccessKeys
	// Only bother calling ListKeys if there are secrets to retrieve
	if hasSecrets {
		subscription := armClient.SubscriptionID()
		// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
		// connection each time through
		redisClient := armredis.NewClient(subscription, armClient.Creds(), armClient.ClientOptions())

		var resp armredis.ClientListKeysResponse
		resp, err = redisClient.ListKeys(ctx, id.ResourceGroupName, obj.AzureName(), nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed listing keys")
		}
		accessKeys = resp.AccessKeys
	}
	secretSlice, err := secretsToWrite(typedObj, accessKeys)
	if err != nil {
		return nil, err
	}

	return secretSlice, nil
}

func secretsSpecified(obj *redis.Redis) (bool, bool) {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false, false
	}

	secrets := obj.Spec.OperatorSpec.Secrets
	hasSecrets := false
	hasEndpoints := false

	if secrets.PrimaryKey != nil ||
		secrets.SecondaryKey != nil {
		hasSecrets = true
	}

	if secrets.HostName != nil ||
		secrets.Port != nil ||
		secrets.SSLPort != nil {
		hasEndpoints = true
	}

	return hasSecrets, hasEndpoints
}

func secretsToWrite(obj *redis.Redis, accessKeys armredis.AccessKeys) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewSecretCollector(obj.Namespace)
	collector.AddSecretValue(operatorSpecSecrets.PrimaryKey, to.String(accessKeys.PrimaryKey))
	collector.AddSecretValue(operatorSpecSecrets.SecondaryKey, to.String(accessKeys.SecondaryKey))
	collector.AddSecretValue(operatorSpecSecrets.HostName, to.String(obj.Status.HostName))
	collector.AddSecretValue(operatorSpecSecrets.Port, intPtrToString(obj.Status.Port))
	collector.AddSecretValue(operatorSpecSecrets.SSLPort, intPtrToString(obj.Status.SslPort))

	return collector.Secrets(), nil
}

func intPtrToString(i *int) string {
	if i == nil {
		return ""
	}

	return strconv.Itoa(*i)
}
