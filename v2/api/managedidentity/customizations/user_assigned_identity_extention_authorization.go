/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"

	v20230131s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ genruntime.KubernetesExporter = &UserAssignedIdentityExtension{}

// ExportKubernetesResources defines a resource which can create other resources in Kubernetes.
func (identity *UserAssignedIdentityExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) ([]client.Object, error) {
	typedObj, ok := obj.(*v20230131s.UserAssignedIdentity)
	if !ok {
		return nil, fmt.Errorf(
			"cannot run on unknown resource type %T, expected *v20230131s.UserAssignedIdentity", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec.Secrets is empty")
		return nil, nil
	}

	collector := secrets.NewCollector(typedObj.Namespace)
	if typedObj.Spec.OperatorSpec != nil && typedObj.Spec.OperatorSpec.Secrets != nil {
		if typedObj.Status.ClientId != nil {
			collector.AddValue(typedObj.Spec.OperatorSpec.Secrets.ClientId, *typedObj.Status.ClientId)
		}
	}
	if typedObj.Spec.OperatorSpec != nil && typedObj.Spec.OperatorSpec.Secrets != nil {
		if typedObj.Status.PrincipalId != nil {
			collector.AddValue(typedObj.Spec.OperatorSpec.Secrets.PrincipalId, *typedObj.Status.PrincipalId)
		}
	}
	if typedObj.Spec.OperatorSpec != nil && typedObj.Spec.OperatorSpec.Secrets != nil {
		if typedObj.Status.TenantId != nil {
			collector.AddValue(typedObj.Spec.OperatorSpec.Secrets.TenantId, *typedObj.Status.TenantId)
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return secrets.SliceToClientObjectSlice(result), nil
}

func secretsSpecified(obj *v20230131s.UserAssignedIdentity) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	specSecrets := obj.Spec.OperatorSpec.Secrets
	hasSecrets := false
	if specSecrets.ClientId != nil ||
		specSecrets.PrincipalId != nil ||
		specSecrets.TenantId != nil {
		hasSecrets = true
	}

	return hasSecrets
}
