/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	storage "github.com/Azure/azure-service-operator/v2/api/botservice/v20220915/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

const migrationToken = "migrationToken"

var _ genruntime.KubernetesSecretExporter = &BotServiceExtension{}

func (extension *BotServiceExtension) ExportKubernetesSecrets(
	ctx context.Context,
	obj genruntime.MetaObject,
	additionalSecrets set.Set[string],
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
	typedObj, ok := obj.(*storage.BotService)
	if !ok {
		return nil, eris.Errorf("cannot run on unknown resource type %T, expected *storage.BotService", obj)
	}

	var _ conversion.Hub = typedObj

	var destination *genruntime.SecretDestination
	if typedObj.Spec.OperatorSpec != nil && typedObj.Spec.OperatorSpec.Secrets != nil {
		destination = typedObj.Spec.OperatorSpec.Secrets.MigrationToken
	}

	requestedSecrets := set.Make[string]()
	if destination != nil {
		requestedSecrets.Add(migrationToken)
	}
	requestedSecrets = set.Union(requestedSecrets, additionalSecrets)
	if len(requestedSecrets) == 0 {
		log.V(1).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	resourceID, hasResourceID := genruntime.GetResourceID(typedObj)
	if !hasResourceID {
		return nil, eris.Errorf("cannot retrieve migration token for BotService %s without an ARM resource ID", typedObj.Name)
	}

	var response struct {
		Properties struct {
			MigrationToken *string `json:"migrationToken"`
		} `json:"properties"`
	}
	_, err := armClient.GetByID(ctx, resourceID, typedObj.GetAPIVersion(), &response)
	if err != nil {
		return nil, eris.Wrap(err, "retrieving BotService migration token")
	}

	token := ""
	if response.Properties.MigrationToken != nil {
		token = *response.Properties.MigrationToken
	}

	collector := secrets.NewCollector(typedObj.Namespace)
	collector.AddValue(destination, token)
	secretObjects, err := collector.Values()
	if err != nil {
		return nil, err
	}

	retrievedSecrets := make(map[string]string)
	if token != "" {
		retrievedSecrets[migrationToken] = token
	}

	return &genruntime.KubernetesSecretExportResult{
		Objs:       secrets.SliceToClientObjectSlice(secretObjects),
		RawSecrets: secrets.SelectSecrets(additionalSecrets, retrievedSecrets),
	}, nil
}
