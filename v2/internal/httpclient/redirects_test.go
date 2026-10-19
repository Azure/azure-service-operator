/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package httpclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRedirects_OriginBoundary(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	cases := map[string]struct {
		target string
		ok     bool
	}{
		"relative": {
			target: "/done",
			ok:     true,
		},
		"same origin": {
			target: "https://management.example.invalid:443/done",
			ok:     true,
		},
		"other host": {
			target: "https://other.example.invalid/done?secret=private",
			ok:     false,
		},
		"subdomain": {
			target: "https://child.management.example.invalid/done",
			ok:     false,
		},
		"other port": {
			target: "https://management.example.invalid:444/done",
			ok:     false,
		},
		"downgrade": {
			target: "http://management.example.invalid/done",
			ok:     false,
		},
		"userinfo": { //nolint:gosec // Dummy URL credentials test userinfo rejection.
			target: "https://user:private@management.example.invalid/done",
			ok:     false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			calls := 0
			base := &http.Client{
				Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					calls++
					status := http.StatusOK
					h := make(http.Header)
					if calls == 1 {
						status = http.StatusFound
						h.Set("Location", c.target)
					}
					return &http.Response{
						StatusCode: status,
						Header:     h,
						Request:    r,
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				}),
			}
			client := NewRedirectGuardedClient(base, allowed)
			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://management.example.invalid/start?secret=private", nil)
			g.Expect(err).ToNot(HaveOccurred())
			resp, err := client.Do(req)
			if resp != nil {
				g.Expect(resp.Body.Close()).To(Succeed())
			}
			if c.ok {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(calls).To(Equal(2))
			} else {
				var rejected *DestinationError
				g.Expect(errors.As(err, &rejected)).To(BeTrue())
				g.Expect(calls).To(Equal(1))
				g.Expect(err).ToNot(MatchError(ContainSubstring("private")))
				g.Expect(err).ToNot(MatchError(ContainSubstring("?")))
			}
		})
	}
}

func TestRedirects_ClientSettingsAndCallbacks(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	callbackErr := errors.New("caller rejected redirect")
	calls := 0
	jar, err := cookiejar.New(nil)
	g.Expect(err).ToNot(HaveOccurred())
	base := &http.Client{
		Timeout:   time.Minute,
		Jar:       jar,
		Transport: &http.Transport{},
		CheckRedirect: func(*http.Request, []*http.Request) error {
			calls++
			return callbackErr
		},
	}
	client := NewRedirectGuardedClient(base, allowed)
	g.Expect(client.client).ToNot(BeIdenticalTo(base))
	g.Expect(client.client.Timeout).To(Equal(base.Timeout))
	g.Expect(client.client.Transport == base.Transport).To(BeTrue())
	g.Expect(client.client.Jar == base.Jar).To(BeTrue())
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://management.example.invalid/done", nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(client.client.CheckRedirect(req, nil)).To(MatchError(callbackErr))
	g.Expect(calls).To(Equal(1))
	g.Expect(base.CheckRedirect(req, nil)).To(MatchError(callbackErr))
	g.Expect(calls).To(Equal(2))
	base.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	client = NewRedirectGuardedClient(base, allowed)
	g.Expect(client.client.CheckRedirect(req, nil)).To(MatchError(http.ErrUseLastResponse))
	base.CheckRedirect = func(r *http.Request, _ []*http.Request) error {
		r.URL.Host = "other.example.invalid"
		return nil
	}
	client = NewRedirectGuardedClient(base, allowed)
	g.Expect(client.client.CheckRedirect(req, nil)).To(HaveOccurred())
}

func TestRedirects_DefaultLimit(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	calls := 0
	base := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			calls++
			return &http.Response{
				StatusCode: http.StatusFound,
				Request:    r,
				Header:     http.Header{"Location": []string{"/loop"}},
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		}),
	}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://management.example.invalid/loop", nil)
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		g.Expect(resp.Body.Close()).To(Succeed())
	}
	g.Expect(err).To(HaveOccurred())
	g.Expect(calls).To(Equal(10))
}

func TestRedirects_NilBase(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	client := NewRedirectGuardedClient(nil, allowed)
	g.Expect(client.client).ToNot(BeIdenticalTo(http.DefaultClient))
	g.Expect(client.client.Timeout).To(Equal(http.DefaultClient.Timeout))
	g.Expect(client.client.Transport == http.DefaultClient.Transport).To(BeTrue())
	g.Expect(client.client.Jar == http.DefaultClient.Jar).To(BeTrue())
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://other.example.invalid/done", nil)
	g.Expect(err).ToNot(HaveOccurred())
	var rejected *DestinationError
	g.Expect(errors.As(client.client.CheckRedirect(req, nil), &rejected)).To(BeTrue())
}

func TestRedirects_UseLastResponse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	calls := 0
	callbackCalls := 0
	base := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			calls++
			return &http.Response{
				StatusCode: http.StatusFound,
				Request:    r,
				Header:     http.Header{"Location": []string{"/done"}},
				Body:       io.NopCloser(strings.NewReader("last response")),
			}, nil
		}),
		CheckRedirect: func(*http.Request, []*http.Request) error {
			callbackCalls++
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://management.example.invalid/start", nil)
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(calls).To(Equal(1))
	g.Expect(callbackCalls).To(Equal(1))
	g.Expect(resp.StatusCode).To(Equal(http.StatusFound))
	body, err := io.ReadAll(resp.Body)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(string(body)).To(Equal("last response"))
}

func TestRedirects_InitialRequestIsNotValidated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	calls := 0
	base := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			calls++
			return &http.Response{
				StatusCode: http.StatusOK,
				Request:    r,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("{}")),
			}, nil
		}),
	}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://other.example.invalid/start", nil)
	g.Expect(err).ToNot(HaveOccurred())
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(calls).To(Equal(1))
	g.Expect(resp.StatusCode).To(Equal(http.StatusOK))
}
