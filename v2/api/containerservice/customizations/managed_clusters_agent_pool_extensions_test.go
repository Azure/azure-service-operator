/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

func TestManagedClustersAgentPoolExtension_PreReconcileOwnerCheck_BlocksWhileClusterUpgradeIsInProgress(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &ManagedClustersAgentPoolExtension{}
	cluster := &containerservice.ManagedCluster{
		Spec: containerservice.ManagedCluster_Spec{
			KubernetesVersion: stringPtr("1.36.1"),
		},
		Status: containerservice.ManagedCluster_STATUS{
			CurrentKubernetesVersion: stringPtr("1.35.4"),
		},
	}

	result, err := extension.PreReconcileOwnerCheck(
		context.Background(),
		cluster,
		nil,
		nil,
		logr.Discard(),
		func(
			context.Context,
			genruntime.MetaObject,
			*resolver.Resolver,
			*genericarmclient.GenericClient,
			logr.Logger,
		) (extensions.PreReconcileCheckResult, error) {
			t.Fatal("next should not be called")
			return extensions.PreReconcileCheckResult{}, nil
		},
	)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.BlockReconciliation()).To(BeTrue())
}

func TestManagedClustersAgentPoolExtension_ClassifyError_NodePoolMcVersionIncompatible_RetriesSlowly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &ManagedClustersAgentPoolExtension{}
	cloudError := genericarmclient.NewTestCloudError("NodePoolMcVersionIncompatible", "control plane is upgrading")
	result, err := extension.ClassifyError(
		cloudError,
		"2025-08-01",
		logr.Discard(),
		func(error *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
			return core.CloudErrorDetails{
				Classification: core.ErrorRetryable,
				Retry:          retry.VerySlow,
			}, nil
		},
	)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Classification).To(Equal(core.ErrorRetryable))
	g.Expect(result.Retry).To(Equal(retry.Slow))
}

func TestAgentPoolVersionExceedsControlPlaneVersion(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		agentPoolVersion    *string
		controlPlaneVersion *string
		expected            bool
	}{
		"higher minor version": {
			agentPoolVersion:    stringPtr("1.36.1"),
			controlPlaneVersion: stringPtr("1.35.4"),
			expected:            true,
		},
		"higher patch version": {
			agentPoolVersion:    stringPtr("1.35.5"),
			controlPlaneVersion: stringPtr("1.35.4"),
			expected:            false,
		},
		"invalid agent pool version": {
			agentPoolVersion:    stringPtr("not-a-version"),
			controlPlaneVersion: stringPtr("1.35.4"),
			expected:            false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			agentPool := &containerservice.ManagedClustersAgentPool{
				Spec: containerservice.ManagedClustersAgentPool_Spec{
					OrchestratorVersion: c.agentPoolVersion,
				},
			}
			cluster := &containerservice.ManagedCluster{
				Status: containerservice.ManagedCluster_STATUS{
					CurrentKubernetesVersion: c.controlPlaneVersion,
				},
			}

			if actual := agentPoolVersionExceedsControlPlaneVersion(agentPool, cluster); actual != c.expected {
				t.Errorf("agentPoolVersionExceedsControlPlaneVersion() = %t, want %t", actual, c.expected)
			}
		})
	}
}

func stringPtr(value string) *string {
	return &value
}
