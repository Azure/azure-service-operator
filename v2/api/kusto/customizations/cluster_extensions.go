/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	kusto "github.com/Azure/azure-service-operator/v2/api/kusto/v1api20230815/storage"
)

var _ extensions.PreReconciliationChecker = &ClusterExtension{}

// clusterTerminalStates is the set of provisioning states that are considered terminal for a cluster.
// These are all listed lowercase, so we can do a case-insensitive match.
var clusterTerminalStates = set.Make(
	"succeeded",
	"failed",
	"canceled",
)

// PreReconcileCheck is called before the reconciliation of the resource to see if the cluster
// is in a state that will allow reconciliation to proceed.
// If a cluster has a provisioningState not in the set above, it will reject any attempt to PUT a
// new state out of hand; so there's no point in even trying. This is true even if the PUT we're
// doing will have no effect on the state of the cluster.
func (ext *ClusterExtension) PreReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ genruntime.MetaObject,
	_ *resolver.Resolver,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	cluster, ok := obj.(*kusto.Cluster)
	if !ok {
		return extensions.PreReconcileCheckResult{},
			eris.Errorf("cannot run on unknown resource type %T, expected *kusto.Cluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = cluster

	// If the cluster is in a state that will reject any PUT, then we should skip reconciliation
	// as there's no point in even trying.
	// This allows us to "play nice with others" and not use up request quota attempting to make changes when we
	// already know those attempts will fail.
	state := cluster.Status.ProvisioningState
	if state != nil && clusterProvisioningStateBlocksReconciliation(state) {
		return extensions.BlockReconcile(
				fmt.Sprintf("Cluster is in provisioning state %q", *state)),
			nil
	}

	return extensions.ProceedWithReconcile(), nil
}

func clusterProvisioningStateBlocksReconciliation(provisioningState *string) bool {
	if provisioningState == nil {
		return false
	}

	return !clusterTerminalStates.Contains(strings.ToLower(*provisioningState))
}

var _ extensions.PostReconciliationChecker = &ClusterExtension{}

func (extension *ClusterExtension) PostReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ genruntime.MetaObject,
	_ *resolver.Resolver,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	cluster, ok := obj.(*kusto.Cluster)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			eris.Errorf("cannot run on unknown resource type %T, expected *kusto.Cluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = cluster

	// We can't claim we've successfully completed reconciliation until the cluster is in a terminal state.
	if cluster.Status.ProvisioningState == nil {
		return extensions.PostReconcileCheckResult{},
			eris.Errorf("no provisioning state found for cluster %s", cluster.Name)
	}

	provisioningState := *cluster.Status.ProvisioningState
	if !clusterTerminalStates.Contains(strings.ToLower(provisioningState)) {
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf("Waiting for cluster to be in a terminal state, current state is %q", provisioningState),
		), nil
	}

	return extensions.PostReconcileCheckResultSuccess(), nil
}
