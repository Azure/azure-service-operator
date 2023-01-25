/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ extensions.PreReconciliationChecker = &ManagedClustersAgentPoolExtension{}

// If an agent pool has a provisioningState in this set, it will reject any attempt to PUT a new state out of hand;
// so there's no point in even trying. This is true even if the PUT we're doing will have no effect on the state of the
// cluster.
// These are all listed lowercase, so we can do a case-insensitive match.
var blockingManagedClustersAgentPoolProvisioningStates = set.Make(
	"creating",
	"updating",
	"scaling",
	"deleting",
	"migrating",
	"upgrading",
	"stopping",
	"starting",
	"canceling",
)

func (ext *ManagedClustersAgentPoolExtension) PreReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	agentPool, ok := obj.(*containerservice.ManagedClustersAgentPool)
	if !ok {
		return extensions.PreReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *containerservice.ManagedCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = agentPool

	// Check to see if the owning cluster is in a state that will block us from reconciling
	if owner != nil {
		if managedCluster, ok := owner.(*containerservice.ManagedCluster); ok {
			if provisioningState := managedCluster.Status.ProvisioningState; clusterProvisioningStateBlocksReconciliation(provisioningState) {
				return extensions.BlockReconcile(
						fmt.Sprintf("Managed cluster %q is in provisioning state %q", owner.GetName(), *provisioningState)),
					nil
			}

		}
	}

	// If the agent pool is in a state that will reject any PUT, then we should skip reconciliation
	// as there's no point in even trying.
	// This allows us to "play nice with others" and not use up request quota attempting to make changes when we
	// already know those attempts will fail.
	if provisioningState := agentPool.Status.ProvisioningState; agentPoolProvisioningStateBlocksReconciliation(provisioningState) {
		return extensions.BlockReconcile(
				fmt.Sprintf("Managed cluster agent pool is in provisioning state %q", *provisioningState)),
			nil
	}

	return extensions.ProceedWithReconcile(), nil
}

func agentPoolProvisioningStateBlocksReconciliation(provisioningState *string) bool {
	if provisioningState == nil {
		return false
	}

	return blockingManagedClustersAgentPoolProvisioningStates.Contains(strings.ToLower(*provisioningState))
}
