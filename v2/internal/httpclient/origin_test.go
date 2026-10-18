/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package httpclient

import (
	"errors"
	"net/url"
	"strings"
	"testing"
)

func TestOrigin(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name string
		raw  string
		ok   bool
	}{
		{"same", "https://management.example.invalid/subscriptions/s/operations/o?q=x", true},
		{"case", "https://MANAGEMENT.example.invalid/path", true},
		{"default port", "https://management.example.invalid:443/path", true},
		{"numeric port", "https://management.example.invalid:0443/path", true},
		{"other host", "https://other.example.invalid/path", false},
		{"subdomain", "https://child.management.example.invalid/path", false},
		{"suffix", "https://management.example.invalid.evil.invalid/path", false},
		{"trailing dot", "https://management.example.invalid./path", false},
		{"port", "https://management.example.invalid:444/path", false},
		{"downgrade", "http://management.example.invalid/path", false},
		{"userinfo", "https://user:private@management.example.invalid/path", false},
		{"relative", "/operations/1", false},
		{"opaque", "https:management.example.invalid/path", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			u, parseErr := url.Parse(c.raw)
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			err := allowed.Validate(u)
			if (err == nil) != c.ok {
				t.Fatalf("expected allowed=%t, got allowed=%t", c.ok, err == nil)
			}
			if err != nil {
				var rejected *DestinationError
				if !errors.As(err, &rejected) {
					t.Fatalf("unexpected error type: %T", err)
				}
				if strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "?") {
					t.Fatalf("unsafe diagnostic: %s", err)
				}
			}
		})
	}
	if err := allowed.Validate(nil); err == nil {
		t.Fatal("nil URL accepted")
	}
}

func TestOrigin_InvalidEndpoint(t *testing.T) {
	t.Parallel()
	for _, raw := range []string{
		"", "/relative", "http://example.invalid", "https://",
		"https://u:secret@example.invalid", "https://example.invalid:",
		"https://example.invalid:0", "https://example.invalid:65536",
		"https://example.invalid:bad", "https://[broken",
	} {
		if _, err := NewOrigin(raw); err == nil {
			t.Fatalf("accepted invalid endpoint %q", raw)
		}
	}
}

func TestOrigin_CustomPortsAndIPv6(t *testing.T) {
	t.Parallel()
	for _, endpoint := range []string{"https://localhost:8443/base", "https://[::1]:8443/base"} {
		allowed, err := NewOrigin(endpoint)
		if err != nil {
			t.Fatal(err)
		}
		u, err := url.Parse(endpoint + "/operations/1")
		if err != nil {
			t.Fatal(err)
		}
		if err := allowed.Validate(u); err != nil {
			t.Fatal(err)
		}
	}
}

func TestOrigin_ConfiguredDNSLabelsAreNotRestricted(t *testing.T) {
	t.Parallel()
	for _, endpoint := range []string{"https://-bad.example.invalid", "https://bad..example.invalid"} {
		allowed, err := NewOrigin(endpoint)
		if err != nil {
			t.Fatal(err)
		}
		u, err := url.Parse(endpoint + "/operations/1")
		if err != nil {
			t.Fatal(err)
		}
		if err := allowed.Validate(u); err != nil {
			t.Fatal(err)
		}
	}
}

func TestOrigin_TransportHostname(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		endpoint    string
		destination string
		ok          bool
	}{
		{"China dotted I", "management.chinacloudapi.cn", "management.ch\u0130nacloudapi.cn", false},
		{"Government dotted I", "management.usgovcloudapi.net", "management.usgovcloudap\u0130.net", false},
		{"Unicode configured", "bücher.example.invalid", "xn--bcher-kva.example.invalid", true},
		{"punycode configured", "xn--bcher-kva.example.invalid", "bücher.example.invalid", true},
		{"Unicode case", "BÜCHER.example.invalid", "bücher.example.invalid", true},
		{"punycode case", "XN--BCHER-KVA.example.invalid", "bücher.example.invalid", true},
		{"mapped ASCII", "management.example.invalid", "ｍanagement.example.invalid", true},
		{"ignored character in hostname", "management.example.invalid", "manage\u00adment.example.invalid", true},
		{"dotted I configured", "management.ch\u0130nacloudapi.cn", "management.xn--chinacloudapi-5nh.cn", true},
		{"dotted I is not ASCII I", "management.ch\u0130nacloudapi.cn", "management.chinacloudapi.cn", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			allowed, err := NewOrigin("https://" + c.endpoint)
			if err != nil {
				t.Fatal(err)
			}
			u, err := url.Parse("https://" + c.destination + "/poll")
			if err != nil {
				t.Fatal(err)
			}
			err = allowed.Validate(u)
			if (err == nil) != c.ok {
				t.Fatalf("allowed=%t, error=%s", c.ok, err)
			}
			if err != nil {
				var rejected *DestinationError
				if !errors.As(err, &rejected) {
					t.Fatalf("unexpected error type: %T", err)
				}
			}
		})
	}
}

func TestOrigin_EmptyIDNAHostname(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name      string
		authority string
	}{
		{"soft hyphen", "\u00ad"},
		{"multiple ignored characters", "\u00ad\u00ad"},
		{"custom port", "\u00ad:8443"},
	} {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewOrigin("https://" + c.authority)
			if err == nil {
				t.Fatal("accepted an empty canonical hostname")
			}
			var rejected *DestinationError
			if !errors.As(err, &rejected) {
				t.Fatalf("unexpected error type: %T", err)
			}
		})
	}
}

func TestOrigin_InvalidIDNA(t *testing.T) {
	t.Parallel()
	allowed, err := NewOrigin("https://management.example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	for _, host := range []string{
		"bad\u200d.example.invalid",
		"bad\ufffd.example.invalid",
		"bad\xff.example.invalid",
		"-bücher.example.invalid",
	} {
		t.Run(host, func(t *testing.T) {
			t.Parallel()
			if _, err := NewOrigin("https://" + host); err == nil {
				t.Fatal("accepted invalid IDNA endpoint")
			}
			err := allowed.Validate(&url.URL{Scheme: "https", Host: host})
			var rejected *DestinationError
			if !errors.As(err, &rejected) {
				t.Fatalf("expected DestinationError, got %s", err)
			}
		})
	}
}
