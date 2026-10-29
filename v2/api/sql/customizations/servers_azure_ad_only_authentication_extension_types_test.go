/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"errors"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

func Test_ServersAzureADOnlyAuthenticationExtension_ClassifyError(t *testing.T) {
	t.Parallel()

	const code = "InvalidServerAADOnlyAuthNoAADAdminPropertyName"
	cases := map[string]struct {
		code           string
		classification core.ErrorClassification
		retry          retry.Classification
		expected       core.ErrorClassification
		expectedRetry  retry.Classification
	}{
		"documented casing": {code, core.ErrorFatal, retry.None, core.ErrorRetryable, retry.Slow},
		"reported casing":   {"invalidServerAADOnlyAuthNoAADAdminPropertyName", core.ErrorFatal, retry.None, core.ErrorRetryable, retry.Slow},
		"lowercase":         {strings.ToLower(code), core.ErrorFatal, retry.None, core.ErrorRetryable, retry.Slow},
		"uppercase":         {strings.ToUpper(code), core.ErrorFatal, retry.None, core.ErrorRetryable, retry.Slow},
		"already retryable": {code, core.ErrorRetryable, retry.VerySlow, core.ErrorRetryable, retry.Slow},
		"unrelated fatal":   {"InvalidParameter", core.ErrorFatal, retry.None, core.ErrorFatal, retry.None},
		"unrelated retry":   {"Conflict", core.ErrorRetryable, retry.Fast, core.ErrorRetryable, retry.Fast},
		"suffix":            {code + "Other", core.ErrorFatal, retry.None, core.ErrorFatal, retry.None},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			cloudError := genericarmclient.NewTestCloudError(c.code, "AAD Admin must be configured first.")
			called := false
			details, err := (&ServersAzureADOnlyAuthenticationExtension{}).ClassifyError(
				cloudError, "2021-11-01", logr.Discard(),
				func(received *genericarmclient.CloudError) (core.CloudErrorDetails, error) {
					called = true
					g.Expect(received).To(BeIdenticalTo(cloudError))
					return core.CloudErrorDetails{
						Classification: c.classification,
						Retry:          c.retry,
						Code:           received.Code(),
						Message:        received.Message(),
					}, nil
				},
			)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(called).To(BeTrue())
			g.Expect(details).To(Equal(core.CloudErrorDetails{
				Classification: c.expected,
				Retry:          c.expectedRetry,
				Code:           cloudError.Code(),
				Message:        cloudError.Message(),
			}))
		})
	}
}

func Test_ServersAzureADOnlyAuthenticationExtension_ClassifyError_Nil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := core.CloudErrorDetails{Classification: core.ErrorRetryable, Retry: retry.Fast}
	details, err := (&ServersAzureADOnlyAuthenticationExtension{}).ClassifyError(
		nil, "2021-11-01", logr.Discard(),
		func(*genericarmclient.CloudError) (core.CloudErrorDetails, error) {
			return expected, nil
		},
	)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(details).To(Equal(expected))
}

func Test_ServersAzureADOnlyAuthenticationExtension_ClassifyError_NextError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := errors.New("classification failed")
	details, err := (&ServersAzureADOnlyAuthenticationExtension{}).ClassifyError(
		genericarmclient.NewTestCloudError("InvalidServerAADOnlyAuthNoAADAdminPropertyName", ""),
		"2021-11-01", logr.Discard(),
		func(*genericarmclient.CloudError) (core.CloudErrorDetails, error) {
			return core.CloudErrorDetails{}, expected
		},
	)
	g.Expect(err).To(BeIdenticalTo(expected))
	g.Expect(details).To(Equal(core.CloudErrorDetails{}))
}
