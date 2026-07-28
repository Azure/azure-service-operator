/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/api/containerservice/customizations"
	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

func TestManagedClustersAgentPoolUpgrade_WaitsForControlPlane(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &customizations.ManagedClustersAgentPoolExtension{}
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
			t.Fatal("agent pool reconciliation should wait for the control plane")
			return extensions.PreReconcileCheckResult{}, nil
		})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.BlockReconciliation()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring("has not reached Kubernetes version"))
	g.Expect(result.Message()).To(ContainSubstring("1.36.1"))
}
