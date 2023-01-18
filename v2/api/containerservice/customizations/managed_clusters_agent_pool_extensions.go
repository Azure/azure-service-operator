/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"

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

var blockingManagedClustersAgentPoolProvisioningStates = set.Make(
	"Creating",
	"Updating",
	"Scaling",
	"Deleting",
	"Migrating",
	"Upgrading",
	"Stopping",
	"Starting",
	"Canceling",
)

func (ext *ManagedClustersAgentPoolExtension) PreReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	agentPool, ok := obj.(*containerservice.ManagedClustersAgentPool)
	if !ok {
		return extensions.SkipReconcile("Expected Managed Cluster Agent Pool"),
			errors.Errorf("cannot run on unknown resource type %T, expected *containerservice.ManagedCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = agentPool

	if provisioningState := agentPool.Status.ProvisioningState; provisioningState != nil {
		if blockingManagedClustersAgentPoolProvisioningStates.Contains(*provisioningState) {
			return extensions.SkipReconcile(
					fmt.Sprintf("Managed cluster agent pool is in provisioning state %q", provisioningState)),
				nil
		}
	}

	return extensions.ProceedWithReconcile(), nil
}
