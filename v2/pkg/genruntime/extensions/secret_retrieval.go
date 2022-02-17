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
		obj genruntime.MetaObject,
		armClient *genericarmclient.GenericClient,
		log logr.Logger) ([]*corev1.Secret, error)
}
