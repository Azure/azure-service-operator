/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package extensions

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// SecretsRetriever is an extension interface which allows code-generated resources to
// include manual implementations querying Azure for secrets, which can then be written
// to Kubernetes (or elsewhere).
type SecretsRetriever interface {
	// RetrieveSecrets is called to retrieve secrets from Azure. The generic controller
	// will write the secrets to their destination.
	RetrieveSecrets(
		ctx context.Context,
		obj genruntime.ARMMetaObject,
		armClient *genericarmclient.GenericClient,
		log logr.Logger) ([]*corev1.Secret, error)
}

type SecretRetrievalFunc = func(obj genruntime.ARMMetaObject) ([]*corev1.Secret, error)

// CreateSecretRetriever creates a function to retrieve secrets from Azure. If the resource
// in question has not been configured with the SecretsRetriever extension, the returned function
// is a no-op.
func CreateSecretRetriever(
	ctx context.Context,
	host genruntime.ResourceExtension,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) SecretRetrievalFunc {

	impl, ok := host.(SecretsRetriever)
	if !ok {
		return func(obj genruntime.ARMMetaObject) ([]*corev1.Secret, error) {
			return nil, nil
		}
	}

	return func(obj genruntime.ARMMetaObject) ([]*corev1.Secret, error) {
		log.V(Info).Info("Retrieving secrets from Azure")
		secrets, err := impl.RetrieveSecrets(ctx, obj, armClient, log)
		if err != nil {
			return secrets, err
		}

		log.V(Info).Info(
			"Successfully retrieved secrets",
			"SecretsToWrite", len(secrets))

		return secrets, nil
	}
}
