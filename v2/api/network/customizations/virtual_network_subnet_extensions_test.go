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

	cases := map[string]struct {
		errorCode    string
		errorMessage string
		mockReturns  core.ErrorClassification
		expected     core.ErrorClassification
	}{
		"Fatal error is still fatal": {
			errorCode:    "BadRequest",
			errorMessage: "Invalid parameter value",
			mockReturns:  core.ErrorFatal,
			expected:     core.ErrorFatal,
		},
		"Other error is still retryable": {
			errorCode:    "OtherError",
			errorMessage: "Some other error occurred",
			mockReturns:  core.ErrorRetryable,
			expected:     core.ErrorRetryable,
		},
		"SubnetRangeOutsideVnet error is retryable": {
			errorCode:    "NetcfgSubnetRangeOutsideVnet",
			errorMessage: "Subnet 'testsubnet' is not valid because its IP address range is outside the IP address range of virtual network 'testvnet'.",
			mockReturns:  core.ErrorRetryable,
			expected:     core.ErrorRetryable,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			extension := &VirtualNetworksSubnetExtension{}

			// Create a test error with the specific code
			err := genericarmclient.NewTestCloudError(c.errorCode, c.errorMessage)

			// Mock the next classifier function - this would normally classify as fatal
			mockCalled := false
			next := func(cloudError *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
				mockCalled = true
				return core.CloudErrorDetails{
					Classification: c.mockReturns,
					Code:           cloudError.Code(),
					Message:        cloudError.Message(),
				}, nil
			}

			details, classifyErr := extension.ClassifyError(err, "2024-03-01", logr.Discard(), next)

			g.Expect(classifyErr).ToNot(HaveOccurred())
			g.Expect(details.Classification).To(Equal(c.expected))
			g.Expect(details.Code).To(Equal(c.errorCode))
			g.Expect(details.Message).To(ContainSubstring(c.errorMessage))
			g.Expect(mockCalled).To(BeTrue())
		})
	}
}
