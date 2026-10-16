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
	state := map[string]any{"asyncURL": "https://attacker.example.invalid/poll", "state": "InProgress"}
	poller, err := factory.NewBackupInstancesClient().BeginSyncBackupInstance(
		context.Background(), "rg", "vault", "instance", armdataprotection.SyncBackupInstanceRequest{},
		&armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions{
			ResumeToken: boundaryToken(t, "BackupInstancesClientSyncBackupInstanceResponse", state),
		},
	)
	g.Expect(err).ToNot(HaveOccurred())
	_, err = poller.Poll(context.Background())
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(BeZero())
	g.Expect(cred.calls.Load()).To(BeZero())
}

func Test_DestinationSDK_CloudBoundariesAreIndependent(t *testing.T) {
	t.Parallel()
	endpoints := []string{
		"https://management.azure.com", "https://management.chinacloudapi.cn",
		"https://management.usgovcloudapi.net", "https://custom.example.invalid:8443/base",
	}
	for _, endpoint := range endpoints {
		t.Run(endpoint, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, endpoint, func(r *http.Request) (*http.Response, error) {
				calls++
				return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
			})
			pl, err := armruntime.NewPipeline("boundary-test", "v0.0.0", cred, runtime.PipelineOptions{}, client.ClientOptions())
			g.Expect(err).ToNot(HaveOccurred())
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, endpoint+"/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := pl.Do(req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(resp.Body.Close()).To(Succeed())
			other, err := runtime.NewRequest(context.Background(), http.MethodGet, "https://different.example.invalid/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			_, err = pl.Do(other)
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(calls).To(Equal(1))
		})
	}
}
