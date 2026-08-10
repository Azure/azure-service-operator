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
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

func Test_AfdCustomDomainExtension_ClassifyError(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		errorCode    string
		errorMessage string
		expected     retry.Classification
	}{
		"Route association error is retried slowly": {
			errorCode:    "BadRequest",
			errorMessage: "This resource is still associated with a route. Please delete the association with the route first before deleting this resource.",
			expected:     retry.Slow,
		},
		"Other bad request remains retried very slowly": {
			errorCode:    "BadRequest",
			errorMessage: "Invalid parameter value",
			expected:     retry.VerySlow,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			extension := &AfdCustomDomainExtension{}
			cloudError := genericarmclient.NewTestCloudError(c.errorCode, c.errorMessage)

			details, err := extension.ClassifyError(
				cloudError,
				"2023-05-01",
				logr.Discard(),
				func(err *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
					return core.CloudErrorDetails{
						Classification: core.ErrorRetryable,
						Retry:          retry.VerySlow,
						Code:           err.Code(),
						Message:        err.Message(),
					}, nil
				},
			)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(details.Retry).To(Equal(c.expected))
		})
	}
}
