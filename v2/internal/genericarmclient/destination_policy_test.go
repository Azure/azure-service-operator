/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

type boundaryCredential struct {
	calls atomic.Int32
}

func (c *boundaryCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	c.calls.Add(1)
	return azcore.AccessToken{
		Token:     "boundary-test-token",
		ExpiresOn: time.Now().Add(time.Hour),
	}, nil
}

type boundaryRoundTripper func(*http.Request) (*http.Response, error)

func (f boundaryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func boundaryResponse(r *http.Request, status int, headers http.Header, body string) *http.Response {
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: status,
		Header:     headers,
		Request:    r,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func boundaryClient(
	t *testing.T,
	endpoint string,
	send func(*http.Request) (*http.Response, error),
) (*genericarmclient.GenericClient, *boundaryCredential) {
	t.Helper()
	g := NewGomegaWithT(t)
	cred := &boundaryCredential{}
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: endpoint,
				Audience: "https://management.core.windows.net/",
			},
		},
	}
	client, err := genericarmclient.NewGenericClient(cfg, cred, &genericarmclient.GenericClientOptions{
		HTTPClient: &http.Client{
			Transport: boundaryRoundTripper(send),
		},
	})
	g.Expect(err).ToNot(HaveOccurred())
	return client, cred
}

func boundaryToken(t *testing.T, resultType string, state map[string]any) string {
	t.Helper()
	b, err := json.Marshal(map[string]any{
		"type":  resultType,
		"token": state,
	})
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func Test_DestinationPolicy_RejectsBeforeAuthentication(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	pl, err := armruntime.NewPipeline("boundary-test", "v0.0.0", cred, runtime.PipelineOptions{}, client.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())
	req, err := runtime.NewRequest(context.Background(), http.MethodGet, "https://attacker.example.invalid/path?secret=private")
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := pl.Do(req)
	if resp != nil {
		g.Expect(resp.Body.Close()).To(Succeed())
	}
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(cred.calls.Load()).To(BeZero())
	g.Expect(calls).To(BeZero())
	g.Expect(err.Error()).ToNot(ContainSubstring("private"))
}

func Test_DestinationPolicy_InvalidEndpoint(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: "http://management.example.invalid",
				Audience: "https://management.core.windows.net/",
			},
		},
	}
	_, err := genericarmclient.NewGenericClient(cfg, &boundaryCredential{}, nil)
	g.Expect(err).To(HaveOccurred())
}

func Test_DestinationPolicy_RedirectDiagnosticIsSafe(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, _ := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusFound,
			http.Header{"Location": []string{"https://attacker.example.invalid/path?secret=private"}}, ""), nil
	})
	var result genericarmclient.GenericResource
	_, err := client.GetByID(context.Background(),
		"/subscriptions/s/resourceGroups/rg/providers/Microsoft.Fake/resources/r", "2020-01-01", &result)
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(Equal(1))
	g.Expect(err.Error()).ToNot(ContainSubstring("private"))
	g.Expect(err.Error()).ToNot(ContainSubstring("boundary-test-token"))
}

func Test_DestinationPolicy_RPRegistrationRedirectIsRejected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const resourcePath = "/subscriptions/s/resourceGroups/rg/providers/Microsoft.Fake/resources/r"
	const registrationPath = "/subscriptions/s/providers/Microsoft.Fake/register"
	var requests []string
	redirectedCalls := 0
	client, _ := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		requests = append(requests, r.Method+" "+r.URL.Host+r.URL.Path)
		if r.URL.Host != "management.example.invalid" {
			redirectedCalls++
		}
		switch r.URL.Path {
		case resourcePath:
			return boundaryResponse(r, http.StatusConflict, nil, rpNotRegisteredError), nil
		case registrationPath:
			return boundaryResponse(r, http.StatusTemporaryRedirect,
				http.Header{"Location": []string{"https://attacker.example.invalid/redirected?secret=private"}}, ""), nil
		default:
			return boundaryResponse(r, http.StatusInternalServerError, nil, "{}"), nil
		}
	})

	_, err := client.BeginCreateOrUpdateByID(t.Context(), resourcePath, "2020-01-01", map[string]any{})
	g.Expect(redirectedCalls).To(BeZero())
	g.Expect(requests).To(Equal([]string{
		"PUT management.example.invalid" + resourcePath,
		"POST management.example.invalid" + registrationPath,
	}))
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(err.Error()).ToNot(ContainSubstring("private"))
	g.Expect(err.Error()).ToNot(ContainSubstring("boundary-test-token"))
}
