/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_VirtualNetworksSubnetExtension_ClassifyError_NetcfgSubnetRangeOutsideVnet_IsRetryable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &VirtualNetworksSubnetExtension{}

	// Create a test error with the specific code
	err := genericarmclient.NewTestCloudError("NetcfgSubnetRangeOutsideVnet", "Subnet 'testsubnet' is not valid because its IP address range is outside the IP address range of virtual network 'testvnet'.")

	// Mock the next classifier function - this would normally classify as fatal
	next := func(cloudError *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
		return core.CloudErrorDetails{
			Classification: core.ErrorFatal, // Default would be fatal
			Code:           cloudError.Code(),
			Message:        cloudError.Message(),
		}, nil
	}

	details, classifyErr := extension.ClassifyError(err, "2024-03-01", logr.Discard(), next)

	g.Expect(classifyErr).ToNot(HaveOccurred())
	g.Expect(details.Classification).To(Equal(core.ErrorRetryable))
	g.Expect(details.Code).To(Equal("NetcfgSubnetRangeOutsideVnet"))
	g.Expect(details.Message).To(ContainSubstring("IP address range is outside"))
}

func Test_VirtualNetworksSubnetExtension_ClassifyError_OtherError_UsesDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &VirtualNetworksSubnetExtension{}

	// Create a test error with a different code
	err := genericarmclient.NewTestCloudError("BadRequest", "Invalid parameter value")

	// Mock the next classifier function - this should remain fatal
	next := func(cloudError *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
		return core.CloudErrorDetails{
			Classification: core.ErrorFatal,
			Code:           cloudError.Code(),
			Message:        cloudError.Message(),
		}, nil
	}

	details, classifyErr := extension.ClassifyError(err, "2024-03-01", logr.Discard(), next)

	g.Expect(classifyErr).ToNot(HaveOccurred())
	g.Expect(details.Classification).To(Equal(core.ErrorFatal)) // Should remain fatal
	g.Expect(details.Code).To(Equal("BadRequest"))
}

func Test_isRetryableSubnetError_NetcfgSubnetRangeOutsideVnet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := genericarmclient.NewTestCloudError("NetcfgSubnetRangeOutsideVnet", "Subnet 'testsubnet' is not valid because its IP address range is outside the IP address range of virtual network 'testvnet'.")
	g.Expect(isRetryableSubnetError(err)).To(BeTrue())
}

func Test_isRetryableSubnetError_OtherError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := genericarmclient.NewTestCloudError("BadRequest", "Invalid parameter value")
	g.Expect(isRetryableSubnetError(err)).To(BeFalse())
}

func Test_isRetryableSubnetError_NilError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(isRetryableSubnetError(nil)).To(BeFalse())
}
