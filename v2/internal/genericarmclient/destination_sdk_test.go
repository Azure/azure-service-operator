/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"

	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

func Test_DestinationSDK_BackupInstanceResume(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	factory, err := armdataprotection.NewClientFactory("subscription", cred, client.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())
	state := map[string]any{
		"asyncURL": "https://attacker.example.invalid/poll",
		"state":    "InProgress",
	}
	poller, err := factory.NewBackupInstancesClient().BeginSyncBackupInstance(
		context.Background(), "rg", "vault", "instance", armdataprotection.SyncBackupInstanceRequest{},
		&armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions{
			ResumeToken: boundaryToken(t, "BackupInstancesClientSyncBackupInstanceResponse", state),
		},
	)
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := poller.Poll(context.Background())
	if resp != nil {
		g.Expect(resp.Body.Close()).To(Succeed())
	}
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(BeZero())
	g.Expect(cred.calls.Load()).To(BeZero())
}

func Test_DestinationSDK_CloudBoundariesAreIndependent(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint string
	}{
		"public": {
			endpoint: "https://management.azure.com",
		},
		"China": {
			endpoint: "https://management.chinacloudapi.cn",
		},
		"Government": {
			endpoint: "https://management.usgovcloudapi.net",
		},
		"custom": {
			endpoint: "https://custom.example.invalid:8443/base",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, c.endpoint, func(r *http.Request) (*http.Response, error) {
				calls++
				return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
			})
			pl, err := armruntime.NewPipeline("boundary-test", "v0.0.0", cred, runtime.PipelineOptions{}, client.ClientOptions())
			g.Expect(err).ToNot(HaveOccurred())
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, c.endpoint+"/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := pl.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			other, err := runtime.NewRequest(context.Background(), http.MethodGet, "https://different.example.invalid/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err = pl.Do(other)
			if resp != nil {
				g.Expect(resp.Body.Close()).To(Succeed())
			}
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(calls).To(Equal(1))
		})
	}
}

func Test_DestinationSDK_ClientsRejectEachOthersOrigins(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const firstEndpoint = "https://management.azure.com"
	const secondEndpoint = "https://management.chinacloudapi.cn"
	firstCalls := 0
	firstClient, firstCred := boundaryClient(t, firstEndpoint, func(r *http.Request) (*http.Response, error) {
		firstCalls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	secondCalls := 0
	secondClient, secondCred := boundaryClient(t, secondEndpoint, func(r *http.Request) (*http.Response, error) {
		secondCalls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	firstPipeline, err := armruntime.NewPipeline("boundary-test", "v0.0.0", firstCred, runtime.PipelineOptions{}, firstClient.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())
	secondPipeline, err := armruntime.NewPipeline("boundary-test", "v0.0.0", secondCred, runtime.PipelineOptions{}, secondClient.ClientOptions())
	g.Expect(err).ToNot(HaveOccurred())

	cases := map[string]struct {
		endpoint      string
		otherEndpoint string
		pipeline      runtime.Pipeline
		calls         *int
	}{
		firstEndpoint: {
			endpoint:      firstEndpoint,
			otherEndpoint: secondEndpoint,
			pipeline:      firstPipeline,
			calls:         &firstCalls,
		},
		secondEndpoint: {
			endpoint:      secondEndpoint,
			otherEndpoint: firstEndpoint,
			pipeline:      secondPipeline,
			calls:         &secondCalls,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, c.endpoint+"/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := c.pipeline.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			g.Expect(*c.calls).To(Equal(1))

			other, err := runtime.NewRequest(context.Background(), http.MethodGet, c.otherEndpoint+"/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err = c.pipeline.Do(other)
			if resp != nil {
				g.Expect(resp.Body.Close()).To(Succeed())
			}
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(*c.calls).To(Equal(1))
		})
	}
}
