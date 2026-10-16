/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"errors"
	"maps"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

func Test_DestinationPoller_RejectsEveryPollingForm(t *testing.T) {
	t.Parallel()
	states := map[string]map[string]any{
		"async":     {"asyncURL": "https://attacker.example.invalid/poll?secret=private", "state": "InProgress"},
		"operation": {"oplocURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"location":  {"type": "loc", "pollURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"body":      {"type": "body", "pollURL": "https://attacker.example.invalid/poll", "state": "InProgress"},
		"fake":      {"fakeURL": "https://attacker.example.invalid/poll", "apiName": "test", "status": "InProgress"},
	}
	for name, state := range states {
		t.Run(name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
				calls++
				return boundaryResponse(r, http.StatusOK, nil, `{"status":"Succeeded"}`), nil
			})
			token := boundaryToken(t, "GenericResource", state)
			resumed := client.ResumeCreatePoller(genericarmclient.CreatePollerID)
			err := resumed.Resume(context.Background(), client, token)
			var rejected *httpclient.DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
			g.Expect(cred.calls.Load()).To(BeZero())
			g.Expect(calls).To(BeZero())
			g.Expect(err.Error()).ToNot(ContainSubstring("private"))
		})
	}
}

func Test_DestinationPoller_FinalAndUnusedURLs(t *testing.T) {
	t.Parallel()
	const endpoint = "https://management.example.invalid"
	const offOrigin = "https://regional.example.invalid/result"
	cases := []struct {
		name      string
		state     map[string]any
		response  string
		rejected  bool
		wantCalls int
	}{
		{"async original", map[string]any{"asyncURL": endpoint + "/poll", "origURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, true, 1},
		{"async location", map[string]any{"asyncURL": endpoint + "/poll", "locURL": offOrigin, "method": "POST"}, `{"status":"Succeeded"}`, true, 1},
		{"operation original", map[string]any{"oplocURL": endpoint + "/poll", "origURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, true, 1},
		{"operation location", map[string]any{"oplocURL": endpoint + "/poll", "locURL": offOrigin, "method": "POST", "finalState": "location"}, `{"status":"Succeeded"}`, true, 1},
		{"operation response location", map[string]any{"oplocURL": endpoint + "/poll", "method": "POST"}, `{"status":"Succeeded","resourceLocation":"https://regional.example.invalid/result"}`, true, 1},
		{"unused regional location", map[string]any{"asyncURL": endpoint + "/poll", "origURL": endpoint + "/resource", "locURL": offOrigin, "method": "PUT"}, `{"status":"Succeeded"}`, false, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			calls := 0
			client, cred := boundaryClient(t, endpoint, func(r *http.Request) (*http.Response, error) {
				calls++
				g.Expect(r.URL.Host).To(Equal("management.example.invalid"))
				body := c.response
				if r.URL.Path == "/resource" {
					body = "{}"
				}
				return boundaryResponse(r, http.StatusOK, nil, body), nil
			})
			state := maps.Clone(c.state)
			state["state"] = "InProgress"
			resumed := client.ResumeCreatePoller(genericarmclient.CreatePollerID)
			err := resumed.Resume(context.Background(), client, boundaryToken(t, "GenericResource", state))
			if c.rejected {
				var rejected *httpclient.DestinationError
				g.Expect(errors.As(err, &rejected)).To(BeTrue())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(resumed.Poller.Done()).To(BeTrue())
			}
			g.Expect(calls).To(Equal(c.wantCalls))
			g.Expect(cred.calls.Load()).To(Equal(int32(1)))
		})
	}
}

func Test_DestinationPoller_ResponseCannotExpandLocationOrigin(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, _ := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusAccepted,
			http.Header{"Location": []string{"https://attacker.example.invalid/next"}}, "{}"), nil
	})
	state := map[string]any{"type": "loc", "pollURL": "https://management.example.invalid/poll", "state": "InProgress"}
	resumed := client.ResumeDeletePoller(genericarmclient.DeletePollerID)
	g.Expect(resumed.Resume(context.Background(), client, boundaryToken(t, "GenericDeleteResponse", state))).To(Succeed())
	token, err := resumed.Poller.ResumeToken()
	g.Expect(err).ToNot(HaveOccurred())
	next := client.ResumeDeletePoller(genericarmclient.DeletePollerID)
	err = next.Resume(context.Background(), client, token)
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(Equal(1))
}

func Test_DestinationPoller_ActionUsesSameBoundary(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	calls := 0
	client, cred := boundaryClient(t, "https://management.example.invalid", func(r *http.Request) (*http.Response, error) {
		calls++
		return boundaryResponse(r, http.StatusOK, nil, "{}"), nil
	})
	state := map[string]any{"asyncURL": "https://attacker.example.invalid/poll", "state": "InProgress"}
	resumed := client.ResumeActionPoller(genericarmclient.ActionPollerID)
	err := resumed.Resume(context.Background(), client, boundaryToken(t, "GenericActionResponse", state))
	var rejected *httpclient.DestinationError
	g.Expect(errors.As(err, &rejected)).To(BeTrue())
	g.Expect(calls).To(BeZero())
	g.Expect(cred.calls.Load()).To(BeZero())
}
