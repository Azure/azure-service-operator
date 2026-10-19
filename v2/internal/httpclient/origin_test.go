/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package httpclient

import (
	"errors"
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

func TestOrigin(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	cases := map[string]struct {
		raw string
		ok  bool
	}{
		"same": {
			raw: "https://management.example.invalid/subscriptions/s/operations/o?q=x",
			ok:  true,
		},
		"case": {
			raw: "https://MANAGEMENT.example.invalid/path",
			ok:  true,
		},
		"default port": {
			raw: "https://management.example.invalid:443/path",
			ok:  true,
		},
		"numeric port": {
			raw: "https://management.example.invalid:0443/path",
			ok:  true,
		},
		"other host": {
			raw: "https://other.example.invalid/path",
			ok:  false,
		},
		"subdomain": {
			raw: "https://child.management.example.invalid/path",
			ok:  false,
		},
		"suffix": {
			raw: "https://management.example.invalid.evil.invalid/path",
			ok:  false,
		},
		"trailing dot": {
			raw: "https://management.example.invalid./path",
			ok:  false,
		},
		"port": {
			raw: "https://management.example.invalid:444/path",
			ok:  false,
		},
		"downgrade": {
			raw: "http://management.example.invalid/path",
			ok:  false,
		},
		"userinfo": { //nolint:gosec // Dummy URL credentials test userinfo rejection.
			raw: "https://user:private@management.example.invalid/path",
			ok:  false,
		},
		"relative": {
			raw: "/operations/1",
			ok:  false,
		},
		"opaque": {
			raw: "https:management.example.invalid/path",
			ok:  false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			u, parseErr := url.Parse(c.raw)
			g.Expect(parseErr).ToNot(HaveOccurred())
			err := allowed.Validate(u)
			if c.ok {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				var rejected *DestinationError
				g.Expect(errors.As(err, &rejected)).To(BeTrue())
				g.Expect(err).ToNot(MatchError(ContainSubstring("private")))
				g.Expect(err).ToNot(MatchError(ContainSubstring("?")))
			}
		})
	}
	g.Expect(allowed.Validate(nil)).To(HaveOccurred())
}

func TestOrigin_InvalidEndpoint(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint string
	}{
		"empty": {
			endpoint: "",
		},
		"relative": {
			endpoint: "/relative",
		},
		"HTTP": {
			endpoint: "http://example.invalid",
		},
		"missing host": {
			endpoint: "https://",
		},
		"userinfo": { //nolint:gosec // Dummy URL credentials test userinfo rejection.
			endpoint: "https://user:private@example.invalid",
		},
		"empty port": {
			endpoint: "https://example.invalid:",
		},
		"zero port": {
			endpoint: "https://example.invalid:0",
		},
		"port out of range": {
			endpoint: "https://example.invalid:65536",
		},
		"nonnumeric port": {
			endpoint: "https://example.invalid:bad",
		},
		"malformed IPv6": {
			endpoint: "https://[broken",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			_, err := NewOrigin(c.endpoint)
			g.Expect(err).To(HaveOccurred())
		})
	}
}

func TestOrigin_CustomPortsAndIPv6(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint string
	}{
		"localhost": {
			endpoint: "https://localhost:8443/base",
		},
		"IPv6": {
			endpoint: "https://[::1]:8443/base",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			allowed, err := NewOrigin(c.endpoint)
			g.Expect(err).ToNot(HaveOccurred())
			u, err := url.Parse(c.endpoint + "/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(allowed.Validate(u)).To(Succeed())
		})
	}
}

func TestOrigin_ConfiguredDNSLabelsAreNotRestricted(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint string
	}{
		"leading hyphen": {
			endpoint: "https://-bad.example.invalid",
		},
		"empty DNS label": {
			endpoint: "https://bad..example.invalid",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			allowed, err := NewOrigin(c.endpoint)
			g.Expect(err).ToNot(HaveOccurred())
			u, err := url.Parse(c.endpoint + "/operations/1")
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(allowed.Validate(u)).To(Succeed())
		})
	}
}

func TestOrigin_TransportHostname(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		endpoint    string
		destination string
		ok          bool
	}{
		"China dotted I": {
			endpoint:    "management.chinacloudapi.cn",
			destination: "management.ch\u0130nacloudapi.cn",
			ok:          false,
		},
		"Government dotted I": {
			endpoint:    "management.usgovcloudapi.net",
			destination: "management.usgovcloudap\u0130.net",
			ok:          false,
		},
		"Unicode configured": {
			endpoint:    "bücher.example.invalid",
			destination: "xn--bcher-kva.example.invalid",
			ok:          true,
		},
		"punycode configured": {
			endpoint:    "xn--bcher-kva.example.invalid",
			destination: "bücher.example.invalid",
			ok:          true,
		},
		"Unicode case": {
			endpoint:    "BÜCHER.example.invalid",
			destination: "bücher.example.invalid",
			ok:          true,
		},
		"punycode case": {
			endpoint:    "XN--BCHER-KVA.example.invalid",
			destination: "bücher.example.invalid",
			ok:          true,
		},
		"mapped ASCII": {
			endpoint:    "management.example.invalid",
			destination: "ｍanagement.example.invalid",
			ok:          true,
		},
		"ignored character in hostname": {
			endpoint:    "management.example.invalid",
			destination: "manage\u00adment.example.invalid",
			ok:          true,
		},
		"dotted I configured": {
			endpoint:    "management.ch\u0130nacloudapi.cn",
			destination: "management.xn--chinacloudapi-5nh.cn",
			ok:          true,
		},
		"dotted I is not ASCII I": {
			endpoint:    "management.ch\u0130nacloudapi.cn",
			destination: "management.chinacloudapi.cn",
			ok:          false,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			allowed, err := NewOrigin("https://" + c.endpoint)
			g.Expect(err).ToNot(HaveOccurred())
			u, err := url.Parse("https://" + c.destination + "/poll")
			g.Expect(err).ToNot(HaveOccurred())
			err = allowed.Validate(u)
			if c.ok {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				var rejected *DestinationError
				g.Expect(errors.As(err, &rejected)).To(BeTrue())
			}
		})
	}
}

func TestOrigin_EmptyIDNAHostname(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		authority string
	}{
		"soft hyphen": {
			authority: "\u00ad",
		},
		"multiple ignored characters": {
			authority: "\u00ad\u00ad",
		},
		"custom port": {
			authority: "\u00ad:8443",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			_, err := NewOrigin("https://" + c.authority)
			var rejected *DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
		})
	}
}

func TestOrigin_InvalidIDNA(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	allowed, err := NewOrigin("https://management.example.invalid")
	g.Expect(err).ToNot(HaveOccurred())
	cases := map[string]struct {
		host string
	}{
		"joiner": {
			host: "bad\u200d.example.invalid",
		},
		"replacement character": {
			host: "bad\ufffd.example.invalid",
		},
		"invalid UTF-8": {
			host: "bad\xff.example.invalid",
		},
		"leading hyphen with Unicode": {
			host: "-bücher.example.invalid",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			_, err := NewOrigin("https://" + c.host)
			g.Expect(err).To(HaveOccurred())
			err = allowed.Validate(&url.URL{
				Scheme: "https",
				Host:   c.host,
			})
			var rejected *DestinationError
			g.Expect(errors.As(err, &rejected)).To(BeTrue())
		})
	}
}
