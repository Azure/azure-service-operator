/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
)

// Data-plane SDK clients built from these options must reach Azure through the ARM client's
// transport (so recordings capture them) and identify themselves as ASO, like ARM requests do.
func Test_DataPlaneClientOptions(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		userAgent       string
		expectUserAgent string
	}{
		"default user agent":    {expectUserAgent: "aso-controller/"},
		"configured user agent": {userAgent: "custom-agent/1.2.3", expectUserAgent: "custom-agent/1.2.3"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var seen []*http.Request
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				seen = append(seen, r)
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			cfg := cloud.Configuration{
				Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					cloud.ResourceManager: {
						Endpoint: server.URL,
						Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
					},
				},
			}

			client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, &genericarmclient.GenericClientOptions{
				HTTPClient: server.Client(),
				Metrics:    asometrics.NewARMClientMetrics(),
				UserAgent:  c.userAgent,
			})
			g.Expect(err).ToNot(HaveOccurred())

			opts := client.DataPlaneClientOptions()

			// Only the user agent policy comes along; the ARM-specific ones (resource provider
			// registration, ARM metrics) don't apply to data-plane URLs
			g.Expect(opts.PerCallPolicies).To(HaveLen(1))
			g.Expect(opts.Cloud).To(Equal(cfg))
			g.Expect(opts.Retry).To(Equal(client.ClientOptions().Retry))

			// A pipeline built from the options, as any data-plane SDK client would build one,
			// reaches the server through the shared transport and carries the user agent
			pipeline := runtime.NewPipeline("dataplane-test", "v0.0.0", runtime.PipelineOptions{}, &opts)
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, server.URL+"/keys/probe")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := pipeline.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			g.Expect(resp.StatusCode).To(Equal(http.StatusOK))

			g.Expect(seen).To(HaveLen(1))
			g.Expect(seen[0].Header.Get("User-Agent")).To(ContainSubstring(c.expectUserAgent))
			g.Expect(seen[0].Header.Get("User-Agent")).To(ContainSubstring("azsdk-go-dataplane-test/"), "the SDK's own user agent must be preserved")
		})
	}
}
