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
	cases := []struct {
		name        string
		endpoint    string
		destination string
		dialAddress string
		allowed     bool
	}{
		{"China dotted I", "management.chinacloudapi.cn", "management.ch\u0130nacloudapi.cn", "management.xn--chinacloudapi-5nh.cn:443", false},
		{"Government dotted I", "management.usgovcloudapi.net", "management.usgovcloudap\u0130.net", "management.xn--usgovcloudapi-goh.net:443", false},
		{"Unicode configured", "bücher.example.invalid", "xn--bcher-kva.example.invalid", "xn--bcher-kva.example.invalid:443", true},
		{"punycode configured", "xn--bcher-kva.example.invalid", "bücher.example.invalid", "xn--bcher-kva.example.invalid:443", true},
		{"Unicode case", "BÜCHER.example.invalid", "bücher.example.invalid", "xn--bcher-kva.example.invalid:443", true},
		{"mapped ASCII", "management.example.invalid", "ｍanagement.example.invalid", "management.example.invalid:443", true},
		{"leading hyphen", "-bad.example.invalid", "-bad.example.invalid", "-bad.example.invalid:443", true},
		{"empty DNS label", "bad..example.invalid", "bad..example.invalid", "bad..example.invalid:443", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			for _, guarded := range []bool{false, true} {
				name := "transport control"
				if guarded {
					name = "resumed poller"
				}
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
					if guarded {
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
					if guarded && !c.allowed {
						var rejected *httpclient.DestinationError
						g.Expect(errors.As(err, &rejected)).To(BeTrue(), "expected DestinationError, got %s", err)
						g.Expect(tokenCalls).To(BeZero())
						g.Expect(sends.Load()).To(BeZero())
						g.Expect(dials.Load()).To(BeZero())
						g.Expect(authenticated.Load()).To(BeFalse())
						return
					}
					g.Expect(errors.Is(err, blocked)).To(BeTrue(), "expected blocked dial, got %s", err)
					g.Expect(sends.Load()).To(Equal(int32(1)))
					g.Expect(dials.Load()).To(Equal(int32(1)))
					g.Expect(address.Load()).To(Equal(c.dialAddress))
					g.Expect(authenticated.Load()).To(Equal(guarded))
					if guarded {
						g.Expect(tokenCalls).To(Equal(int32(1)))
					}
				})
			}
		})
	}
}
