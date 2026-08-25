/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
)

const (
	postActionResourceID = "/subscriptions/12345/resourceGroups/myrg/providers/Microsoft.Fake/fakeResource/fake"
	postActionAPIVersion = "2019-01-01"
)

func Test_BeginPostActionByID_GivenAction_PostsToTheActionURL(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	var path, query string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		g.Expect(r.Method).To(Equal(http.MethodPost))
		path = r.URL.Path
		query = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte("{}"))).ToNot(BeZero())
	}))
	defer server.Close()

	client := postActionClient(g, server)

	poller, err := client.BeginPostActionByID(ctx, postActionResourceID, "start", postActionAPIVersion)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(poller.Poller.Done()).To(BeTrue())
	g.Expect(path).To(Equal(postActionResourceID + "/start"))
	g.Expect(query).To(Equal("api-version=" + postActionAPIVersion))
}

func Test_BeginPostActionByID_GivenLongRunningAction_ReturnsWithoutWaiting(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	var polled bool
	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/operations/1" {
			polled = true
		}

		w.Header().Set("Azure-AsyncOperation", server.URL+"/operations/1")
		w.WriteHeader(http.StatusAccepted)
		// ARM answers with the resource, whose provisioningState would end the poll before it began
		g.Expect(w.Write([]byte(`{"properties":{"status":"Stopped","provisioningState":"Succeeded"}}`))).ToNot(BeZero())
	}))
	defer server.Close()

	client := postActionClient(g, server)

	poller, err := client.BeginPostActionByID(ctx, postActionResourceID, "start", postActionAPIVersion)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(polled).To(BeFalse())
	g.Expect(poller.Poller.Done()).To(BeFalse())

	// The caller has to be able to put the operation down and pick it up on a later reconcile
	g.Expect(poller.Poller.ResumeToken()).ToNot(BeEmpty())
}

// An action reports failure through its operation, so only a caller that resumes it ever learns
func Test_ResumeActionPoller_GivenFailedOperation_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	var accepted bool
	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/operations/1" {
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(`{"status":"Failed","error":{"code":"WatcherStartFailed","message":"No."}}`))).ToNot(BeZero())
			return
		}

		accepted = true
		w.Header().Set("Azure-AsyncOperation", server.URL+"/operations/1")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	client := postActionClient(g, server)

	poller, err := client.BeginPostActionByID(ctx, postActionResourceID, "start", postActionAPIVersion)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(accepted).To(BeTrue())

	token, err := poller.Poller.ResumeToken()
	g.Expect(err).ToNot(HaveOccurred())

	resumed := client.ResumeActionPoller(genericarmclient.ActionPollerID)
	err = resumed.Resume(ctx, client, token)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("WatcherStartFailed"))
}

func Test_BeginPostActionByID_GivenRejectedAction_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		g.Expect(w.Write([]byte(`{"error":{"code":"WatcherStartFailedDueToNoTargets","message":"No targets."}}`))).ToNot(BeZero())
	}))
	defer server.Close()

	client := postActionClient(g, server)

	_, err := client.BeginPostActionByID(ctx, postActionResourceID, "start", postActionAPIVersion)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("WatcherStartFailedDueToNoTargets"))
}

func Test_BeginPostActionByID_GivenMissingParameter_ReturnsErrorWithoutCallingAzure(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		resourceID string
		action     string
	}{
		"No resource ID": {
			resourceID: "",
			action:     "start",
		},
		"No action": {
			resourceID: postActionResourceID,
			action:     "",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				g.Fail(fmt.Sprintf("unexpected request. Method: %s, URL: %s", r.Method, r.URL))
			}))
			defer server.Close()

			client := postActionClient(g, server)

			_, err := client.BeginPostActionByID(context.Background(), c.resourceID, c.action, postActionAPIVersion)
			g.Expect(err).To(HaveOccurred())
		})
	}
}

func postActionClient(g *WithT, server *httptest.Server) *genericarmclient.GenericClient {
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	options := &genericarmclient.GenericClientOptions{
		HTTPClient: server.Client(),
		Metrics:    asometrics.NewARMClientMetrics(),
	}

	client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, options)
	g.Expect(err).ToNot(HaveOccurred())

	return client
}
