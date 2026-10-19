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
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRedirects_OriginBoundary(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		target string
		ok     bool
	}{
		{"/done", true},
		{"https://management.example.invalid:443/done", true},
		{"https://other.example.invalid/done?secret=private", false},
		{"https://child.management.example.invalid/done", false},
		{"https://management.example.invalid:444/done", false},
		{"http://management.example.invalid/done", false},
		{"https://user:private@management.example.invalid/done", false},
	} {
		t.Run(c.target, func(t *testing.T) {
			t.Parallel()
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
			if err != nil {
				t.Fatal(err)
			}
			resp, err := client.Do(req)
			if resp != nil {
				_ = resp.Body.Close()
			}
			if c.ok {
				if err != nil || calls != 2 {
					t.Fatalf("valid redirect failed: calls=%d error=%s", calls, err)
				}
			} else {
				var rejected *DestinationError
				if !errors.As(err, &rejected) || calls != 1 {
					t.Fatalf("redirect not blocked: calls=%d error=%s", calls, err)
				}
				if strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "?") {
					t.Fatalf("unsafe outer error: %s", err)
				}
			}
		})
	}
}

func TestRedirects_ClientSettingsAndCallbacks(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	callbackErr := errors.New("caller rejected redirect")
	calls := 0
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
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
	if client.client == base || client.client.Timeout != base.Timeout || client.client.Transport != base.Transport || client.client.Jar != base.Jar {
		t.Fatal("client was not copied with settings intact")
	}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://management.example.invalid/done", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := client.client.CheckRedirect(req, nil); !errors.Is(err, callbackErr) || calls != 1 {
		t.Fatalf("caller callback lost: %s", err)
	}
	if err := base.CheckRedirect(req, nil); !errors.Is(err, callbackErr) || calls != 2 {
		t.Fatal("base client callback was modified")
	}
	base.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	client = NewRedirectGuardedClient(base, allowed)
	if err := client.client.CheckRedirect(req, nil); !errors.Is(err, http.ErrUseLastResponse) {
		t.Fatalf("ErrUseLastResponse not preserved: %s", err)
	}
	base.CheckRedirect = func(r *http.Request, _ []*http.Request) error {
		r.URL.Host = "other.example.invalid"
		return nil
	}
	client = NewRedirectGuardedClient(base, allowed)
	if err := client.client.CheckRedirect(req, nil); err == nil {
		t.Fatal("callback mutation bypassed origin check")
	}
}

func TestRedirects_DefaultLimit(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
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
	if err != nil {
		t.Fatal(err)
	}
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		_ = resp.Body.Close()
	}
	if err == nil || calls != 10 {
		t.Fatalf("redirect limit changed: calls=%d error=%s", calls, err)
	}
}

func TestRedirects_NilBase(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	client := NewRedirectGuardedClient(nil, allowed)
	if client.client == http.DefaultClient || client.client.Timeout != http.DefaultClient.Timeout ||
		client.client.Transport != http.DefaultClient.Transport || client.client.Jar != http.DefaultClient.Jar {
		t.Fatal("default client was not copied with settings intact")
	}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://other.example.invalid/done", nil)
	if err != nil {
		t.Fatal(err)
	}
	var rejected *DestinationError
	if err := client.client.CheckRedirect(req, nil); !errors.As(err, &rejected) {
		t.Fatalf("default client redirect not guarded: %s", err)
	}
}

func TestRedirects_UseLastResponse(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
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
	if err != nil {
		t.Fatal(err)
	}
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil || resp == nil || calls != 1 || callbackCalls != 1 || resp.StatusCode != http.StatusFound {
		t.Fatalf("last response not preserved: response=%p calls=%d callbackCalls=%d error=%s", resp, calls, callbackCalls, err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil || string(body) != "last response" {
		t.Fatalf("last response body not preserved: body=%q error=%s", body, err)
	}
}

func TestRedirects_InitialRequestIsNotValidated(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
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
	if err != nil {
		t.Fatal(err)
	}
	resp, err := NewRedirectGuardedClient(base, allowed).Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil || resp == nil || calls != 1 || resp.StatusCode != http.StatusOK {
		t.Fatalf("initial request was guarded: response=%p calls=%d error=%s", resp, calls, err)
	}
}
