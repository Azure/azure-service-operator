/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ genruntime.KubernetesExporter = &FlexibleServerExtension{}

func (ext *FlexibleServerExtension) ExportKubernetesResources(
	_ context.Context,
	obj genruntime.MetaObject,
	_ *genericarmclient.GenericClient,
	log logr.Logger) ([]client.Object, error) {

	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*postgresql.FlexibleServer)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *postgresql.FlexibleServer", obj)
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

	return secrets.SliceToClientObjectSlice(secretSlice), nil
}

func secretsSpecified(obj *postgresql.FlexibleServer) bool {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return false
	}

	operatorSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSecrets.FullyQualifiedDomainName != nil {
		return true
	}

	return false
}

func secretsToWrite(obj *postgresql.FlexibleServer) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := secrets.NewCollector(obj.Namespace)
	collector.AddValue(operatorSpecSecrets.FullyQualifiedDomainName, to.String(obj.Status.FullyQualifiedDomainName))

	return collector.Values()
}

var _ extensions.PreReconciliationChecker = &FlexibleServerExtension{}

var blockingFlexibleServerStates = set.Make(
	"starting",
	"stopping",
	"updating",
)

func (ext *FlexibleServerExtension) PreReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	flexibleServer, ok := obj.(*postgresql.FlexibleServer)
	if !ok {
		return extensions.SkipReconcile("Expected Flexible Server"),
			errors.Errorf("cannot run on unknown resource type %T, expected *postgresql.FlexibleServer", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = flexibleServer

	if state := flexibleServer.Status.State; state != nil {
		if blockingFlexibleServerStates.Contains(strings.ToLower(*state)) {
			return extensions.SkipReconcile(
					fmt.Sprintf("Flexible Server is in provisioning state %q", *state)),
				nil
		}
	}

	return extensions.ProceedWithReconcile(), nil
}
