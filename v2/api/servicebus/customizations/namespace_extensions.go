/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesExporter = &NamespaceExtension{}

func (ext *NamespaceExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger) ([]client.Object, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*servicebus.Namespace)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *servicebus.Namespace", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasEndpoints := secretsSpecified(typedObj)
	if !hasEndpoints {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	// TODO: Get other secrets in the future, if needed. They're omitted for now in the hopes that
	// TODO: users use managed identity which is the recommended way.

	secretSlice, err := secretsToWrite(typedObj)
	if err != nil {
		return nil, err
	}

	return secrets.SliceToClientObjectSlice(secretSlice), nil
}

func secretsSpecified(obj *servicebus.Namespace) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	specSecrets := obj.Spec.OperatorSpec.Secrets
	hasEndpoints := false

	if specSecrets.Endpoint != nil {
		hasEndpoints = true
	}

	return hasEndpoints
}

func secretsToWrite(obj *servicebus.Namespace) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewCollector(obj.Namespace)
	collector.AddValue(operatorSpecSecrets.Endpoint, to.String(obj.Status.ServiceBusEndpoint))

	return collector.Values()
}
