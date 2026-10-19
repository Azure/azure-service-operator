/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync/atomic"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

func Test_DestinationPoller_TransportHostname(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint    string
		destination string
		dialAddress string
		allowed     bool
	}{
		"China dotted I": {
			endpoint:    "management.chinacloudapi.cn",
			destination: "management.ch\u0130nacloudapi.cn",
			dialAddress: "management.xn--chinacloudapi-5nh.cn:443",
			allowed:     false,
		},
		"Government dotted I": {
			endpoint:    "management.usgovcloudapi.net",
			destination: "management.usgovcloudap\u0130.net",
			dialAddress: "management.xn--usgovcloudapi-goh.net:443",
			allowed:     false,
		},
		"Unicode configured": {
			endpoint:    "bücher.example.invalid",
			destination: "xn--bcher-kva.example.invalid",
			dialAddress: "xn--bcher-kva.example.invalid:443",
			allowed:     true,
		},
		"punycode configured": {
			endpoint:    "xn--bcher-kva.example.invalid",
			destination: "bücher.example.invalid",
			dialAddress: "xn--bcher-kva.example.invalid:443",
			allowed:     true,
		},
		"Unicode case": {
			endpoint:    "BÜCHER.example.invalid",
			destination: "bücher.example.invalid",
			dialAddress: "xn--bcher-kva.example.invalid:443",
			allowed:     true,
		},
		"mapped ASCII": {
			endpoint:    "management.example.invalid",
			destination: "ｍanagement.example.invalid",
			dialAddress: "management.example.invalid:443",
			allowed:     true,
		},
		"leading hyphen": {
			endpoint:    "-bad.example.invalid",
			destination: "-bad.example.invalid",
			dialAddress: "-bad.example.invalid:443",
			allowed:     true,
		},
		"empty DNS label": {
			endpoint:    "bad..example.invalid",
			destination: "bad..example.invalid",
			dialAddress: "bad..example.invalid:443",
			allowed:     true,
		},
	}
	modes := map[string]struct {
		guarded bool
	}{
		"transport control": {
			guarded: false,
		},
		"resumed poller": {
			guarded: true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for name, mode := range modes {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					g := NewGomegaWithT(t)
					blocked := errors.New("test dial blocked: no network")
					var dials atomic.Int32
					var address atomic.Value
					address.Store("")
					transport := &http.Transport{
						DialContext: func(_ context.Context, _, addr string) (net.Conn, error) {
							dials.Add(1)
							address.Store(addr)
							return nil, blocked
						},
					}
					t.Cleanup(transport.CloseIdleConnections)
					var sends atomic.Int32
					var authenticated atomic.Bool
					send := func(r *http.Request) (*http.Response, error) {
						sends.Add(1)
						authenticated.Store(r.Header.Get("Authorization") == "Bearer boundary-test-token")
						return transport.RoundTrip(r)
					}

					var err error
					var tokenCalls int32
					destination := "https://" + c.destination + "/poll"
					if mode.guarded {
						client, cred := boundaryClient(t, "https://"+c.endpoint, send)
						state := map[string]any{
							"asyncURL": destination,
							"state":    "InProgress",
						}
						resumed := client.ResumeCreatePoller(genericarmclient.CreatePollerID)
						err = resumed.Resume(t.Context(), client, boundaryToken(t, "GenericResource", state))
						tokenCalls = cred.calls.Load()
					} else {
						// Independently observe Go's destination, without the origin validator or SDK.
						client := &http.Client{
							Transport: boundaryRoundTripper(send),
						}
						req, requestErr := http.NewRequestWithContext(t.Context(), http.MethodGet, destination, nil)
						g.Expect(requestErr).ToNot(HaveOccurred())
						var resp *http.Response
						resp, err = client.Do(req)
						if resp != nil {
							g.Expect(resp.Body.Close()).To(Succeed())
						}
					}
					t.Logf(
						"tokens=%d sends=%d bearer=%t dials=%d address=%s",
						tokenCalls,
						sends.Load(),
						authenticated.Load(),
						dials.Load(),
						address.Load(),
					)
					if mode.guarded && !c.allowed {
						var rejected *httpclient.DestinationError
						g.Expect(errors.As(err, &rejected)).To(BeTrue())
						g.Expect(tokenCalls).To(BeZero())
						g.Expect(sends.Load()).To(BeZero())
						g.Expect(dials.Load()).To(BeZero())
						g.Expect(authenticated.Load()).To(BeFalse())
						return
					}
					g.Expect(err).To(MatchError(blocked))
					g.Expect(sends.Load()).To(Equal(int32(1)))
					g.Expect(dials.Load()).To(Equal(int32(1)))
					g.Expect(address.Load()).To(Equal(c.dialAddress))
					g.Expect(authenticated.Load()).To(Equal(mode.guarded))
					if mode.guarded {
						g.Expect(tokenCalls).To(Equal(int32(1)))
					}
				})
			}
		})
	}
}
