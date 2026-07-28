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
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
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
			KubernetesVersion: to.Ptr("1.36.1"),
		},
		Status: containerservice.ManagedCluster_STATUS{
			CurrentKubernetesVersion: to.Ptr("1.35.4"),
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
	g.Expect(result.Message()).To(ContainSubstring("has not reached Kubernetes version"))
	g.Expect(result.Message()).To(ContainSubstring("1.36.1"))
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
				Message:        "control plane is upgrading",
			}, nil
		},
	)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Classification).To(Equal(core.ErrorRetryable))
	g.Expect(result.Retry).To(Equal(retry.Slow))
	g.Expect(result.Message).To(ContainSubstring("control plane is upgrading"))
}

func TestManagedClusterUpgradeInProgress(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		specVersion    *string
		currentVersion *string
		expected       bool
	}{
		"upgrade in progress (minor bump)": {
			specVersion:    to.Ptr("1.36.1"),
			currentVersion: to.Ptr("1.35.4"),
			expected:       true,
		},
		"patch-only difference does not block": {
			specVersion:    to.Ptr("1.35.5"),
			currentVersion: to.Ptr("1.35.4"),
			expected:       false,
		},
		"downgrade does not block": {
			specVersion:    to.Ptr("1.34.0"),
			currentVersion: to.Ptr("1.35.4"),
			expected:       false,
		},
		"versions equal does not block": {
			specVersion:    to.Ptr("1.35.4"),
			currentVersion: to.Ptr("1.35.4"),
			expected:       false,
		},
		"nil spec version does not block": {
			specVersion:    nil,
			currentVersion: to.Ptr("1.35.4"),
			expected:       false,
		},
		"nil current version does not block": {
			specVersion:    to.Ptr("1.36.0"),
			currentVersion: nil,
			expected:       false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			cluster := &containerservice.ManagedCluster{
				Spec: containerservice.ManagedCluster_Spec{
					KubernetesVersion: c.specVersion,
				},
				Status: containerservice.ManagedCluster_STATUS{
					CurrentKubernetesVersion: c.currentVersion,
				},
			}

			g.Expect(managedClusterUpgradeInProgress(cluster)).To(Equal(c.expected))
		})
	}
}

func TestAgentPoolVersionExceedsControlPlaneVersion(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		agentPoolVersion    *string
		controlPlaneVersion *string
		expected            bool
	}{
		"higher minor version": {
			agentPoolVersion:    to.Ptr("1.36.1"),
			controlPlaneVersion: to.Ptr("1.35.4"),
			expected:            true,
		},
		"higher patch version": {
			agentPoolVersion:    to.Ptr("1.35.5"),
			controlPlaneVersion: to.Ptr("1.35.4"),
			expected:            false,
		},
		"invalid agent pool version": {
			agentPoolVersion:    to.Ptr("not-a-version"),
			controlPlaneVersion: to.Ptr("1.35.4"),
			expected:            false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

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

			g.Expect(agentPoolVersionExceedsControlPlaneVersion(agentPool, cluster)).To(Equal(c.expected))
		})
	}
}
