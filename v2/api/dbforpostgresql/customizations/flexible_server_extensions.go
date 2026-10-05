/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesSecretExporter = &FlexibleServerExtension{}

func (ext *FlexibleServerExtension) ExportKubernetesSecrets(
	ctx context.Context,
	obj genruntime.MetaObject,
	additionalSecrets set.Set[string],
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*postgresql.FlexibleServer)
	if !ok {
		return nil, eris.Errorf("cannot run on unknown resource type %T, expected *postgresql.FlexibleServer", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	hasSecrets := secretsSpecified(typedObj)
	if !hasSecrets {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	secretSlice, err := secretsToWrite(typedObj)
	if err != nil {
		return nil, err
	}

	return &genruntime.KubernetesSecretExportResult{
		Objs:       secrets.SliceToClientObjectSlice(secretSlice),
		RawSecrets: nil, // No RawSecrets as all secret values are coming from status (not real secrets)
	}, nil
}

func secretsSpecified(obj *postgresql.FlexibleServer) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	operatorSecrets := obj.Spec.OperatorSpec.Secrets
	return operatorSecrets.FullyQualifiedDomainName != nil
}

func secretsToWrite(obj *postgresql.FlexibleServer) ([]*v1.Secret, error) {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return nil, nil
	}

	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets

	collector := secrets.NewCollector(obj.Namespace)
	collector.AddValue(operatorSpecSecrets.FullyQualifiedDomainName, to.Value(obj.Status.FullyQualifiedDomainName))

	return collector.Values()
}

var _ extensions.PreReconciliationChecker = &FlexibleServerExtension{}

// If the provisioningState of a flexible server is not in this set, it will reject any attempt to PUT the resource
// out of hand; so there's no point in even trying. This is true even if the PUT we're doing will have no effect on
// the state of the server.
// These are all listed lowercase, so we can do a case-insensitive match.
var nonBlockingFlexibleServerStates = set.Make(
	"succeeded",
	"failed",
	"canceled",
	"ready",
)

func (ext *FlexibleServerExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	server, ok := obj.(*postgresql.FlexibleServer)
	if !ok {
		return extensions.PreReconcileCheckResult{},
			eris.Errorf("cannot run on unknown resource type %T, expected *postgresql.FlexibleServer", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated
	var _ conversion.Hub = server

	state := server.Status.State
	if state != nil && flexibleServerStateBlocksReconciliation(*state) {
		return extensions.BlockReconcile(
			fmt.Sprintf(
				"Flexible Server is in state %q",
				*state,
			),
		), nil
	}

	return next(ctx, obj, resourceResolver, armClient, log)
}

func flexibleServerStateBlocksReconciliation(state string) bool {
	return !nonBlockingFlexibleServerStates.Contains(strings.ToLower(state))
}

var _ extensions.ARMResourceModifier = &FlexibleServerExtension{}

const (
	createModeReplica = "Replica"
	createModeUpdate  = "Update"
)

// ModifyARMResource sends createMode Update in place of Replica once the replica exists, as the PostgreSQL
// resource provider rejects any further Replica PUT with ReadReplicaAlreadyExistWithGivenName.
// See https://github.com/Azure/azure-service-operator/issues/5086.
func (ext *FlexibleServerExtension) ModifyARMResource(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	server, ok := obj.(*postgresql.FlexibleServer)
	if !ok {
		return nil, eris.Errorf("cannot run on unknown resource type %T, expected *postgresql.FlexibleServer", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = server

	if to.Value(server.Spec.CreateMode) != createModeReplica {
		return armObj, nil
	}

	exists, _, err := armClient.CheckExistenceWithGetByID(ctx, armObj.GetID(), armObj.Spec().GetAPIVersion())
	if err != nil {
		return nil, eris.Wrapf(err, "checking whether replica %s exists", armObj.GetID())
	}

	if !exists {
		return armObj, nil
	}

	log.V(Verbose).Info("Replica already exists, sending createMode Update", "resourceID", armObj.GetID())

	err = reflecthelpers.SetProperty(armObj.Spec(), "Properties.CreateMode", to.Ptr(createModeUpdate))
	if err != nil {
		return nil, eris.Wrapf(err, "setting CreateMode to %s", createModeUpdate)
	}

	return armObj, nil
}
