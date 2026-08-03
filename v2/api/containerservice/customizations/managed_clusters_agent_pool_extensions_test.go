/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

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
