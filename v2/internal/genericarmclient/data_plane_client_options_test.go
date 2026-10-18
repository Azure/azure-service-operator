/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
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

			var mu sync.Mutex
			var userAgents []string
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				userAgents = append(userAgents, r.Header.Get("User-Agent"))
				mu.Unlock()
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			cfg := cloud.Configuration{
				Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					cloud.ResourceManager: {
						Endpoint: "https://management.example.invalid",
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
			// The operator relies on the SDK not retrying by itself; its reconcile loop is the retry
			g.Expect(opts.Retry.MaxRetries).To(Equal(int32(-1)))

			// A pipeline built from the options, as any data-plane SDK client would build one,
			// reaches the server through the shared transport and carries the user agent
			pipeline := runtime.NewPipeline("dataplane-test", "v0.0.0", runtime.PipelineOptions{}, &opts)
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, server.URL+"/keys/probe")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := pipeline.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			g.Expect(resp.StatusCode).To(Equal(http.StatusOK))

			mu.Lock()
			defer mu.Unlock()
			g.Expect(userAgents).To(HaveLen(1))
			g.Expect(userAgents[0]).To(ContainSubstring(c.expectUserAgent))
			g.Expect(userAgents[0]).To(ContainSubstring("azsdk-go-dataplane-test/"), "the SDK's own user agent must be preserved")
		})
	}
}

func Test_DataPlaneClientOptions_PreservesRedirects(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const initialURL = "https://vault.example.invalid/keys/probe"
	const finalURL = "https://regional.example.invalid/keys/probe"
	var requests []string
	callbackCalls := 0
	httpClient := &http.Client{
		Transport: boundaryRoundTripper(func(r *http.Request) (*http.Response, error) {
			requests = append(requests, r.URL.String())
			switch r.URL.String() {
			case initialURL:
				return boundaryResponse(r, http.StatusFound, http.Header{"Location": []string{finalURL}}, ""), nil
			case finalURL:
				return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
			default:
				return boundaryResponse(r, http.StatusInternalServerError, nil, "{}"), nil
			}
		}),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			callbackCalls++
			g.Expect(req.URL.String()).To(Equal(finalURL))
			g.Expect(via).To(HaveLen(1))
			g.Expect(via[0].URL.String()).To(Equal(initialURL))
			return nil
		},
	}
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: "https://management.example.invalid",
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}
	client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{},
		&genericarmclient.GenericClientOptions{HTTPClient: httpClient})
	g.Expect(err).ToNot(HaveOccurred())
	opts := client.DataPlaneClientOptions()
	pipeline := runtime.NewPipeline("dataplane-test", "v0.0.0", runtime.PipelineOptions{}, &opts)
	req, err := runtime.NewRequest(t.Context(), http.MethodGet, initialURL)
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := pipeline.Do(req)
	if resp != nil {
		g.Expect(resp.Body.Close()).To(Succeed())
	}
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(http.StatusOK))
	g.Expect(requests).To(Equal([]string{initialURL, finalURL}))
	g.Expect(callbackCalls).To(Equal(1))
}

// Settings the ARM client is given must reach the data-plane client too, not just the ones ASO sets
// today; ClientOptions() hands out the live options, so adjusting them stands in for a differently
// configured ARM client.
func Test_DataPlaneClientOptions_CarriesLoggingAndTelemetry(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: "https://management.example.invalid",
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, nil)
	g.Expect(err).ToNot(HaveOccurred())

	client.ClientOptions().Logging.IncludeBody = true
	client.ClientOptions().Logging.AllowedHeaders = []string{"x-ms-request-id"}
	client.ClientOptions().Telemetry.ApplicationID = "aso-test"

	opts := client.DataPlaneClientOptions()
	g.Expect(opts.Logging.IncludeBody).To(BeTrue())
	g.Expect(opts.Logging.AllowedHeaders).To(ConsistOf("x-ms-request-id"))
	g.Expect(opts.Telemetry.ApplicationID).To(Equal("aso-test"))
}
