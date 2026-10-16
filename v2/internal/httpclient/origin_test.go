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
			u, parseErr := url.Parse(c.raw)
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			err := allowed.Validate(u)
			if (err == nil) != c.ok {
				t.Fatalf("allowed=%v, error=%v", c.ok, err)
			}
			if err != nil {
				var rejected *DestinationError
				if !errors.As(err, &rejected) {
					t.Fatalf("unexpected error type: %T", err)
				}
				if strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "?") {
					t.Fatalf("unsafe diagnostic: %v", err)
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
