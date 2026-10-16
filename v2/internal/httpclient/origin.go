/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package httpclient

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

// DestinationError identifies a rejected destination without retaining a full URL.
type DestinationError struct {
	Allowed     string
	Destination string
}

func (e *DestinationError) Error() string {
	return fmt.Sprintf("request destination %s is not allowed; expected %s", e.Destination, e.Allowed)
}

// Origin is an immutable HTTPS hostname and effective port.
type Origin struct {
	host string
	port string
}

// NewOrigin validates a configured HTTPS endpoint.
func NewOrigin(endpoint string) (Origin, error) {
	u, err := url.Parse(endpoint)
	if err == nil {
		if origin, ok := originOf(u); ok {
			return origin, nil
		}
	}
	return Origin{}, &DestinationError{
		Allowed: "a valid HTTPS origin", Destination: "<invalid HTTPS origin>",
	}
}

func (o Origin) String() string {
	return "https://" + net.JoinHostPort(o.host, o.port)
}

// Validate checks a destination without comparing its path or query.
func (o Origin) Validate(destination *url.URL) error {
	candidate, ok := originOf(destination)
	if ok && o.host != "" && candidate == o {
		return nil
	}
	safe := "<invalid HTTPS origin>"
	if ok {
		safe = candidate.String()
	}
	return &DestinationError{Allowed: o.String(), Destination: safe}
}

func originOf(u *url.URL) (Origin, bool) {
	if u == nil || !strings.EqualFold(u.Scheme, "https") || u.Opaque != "" || u.User != nil || u.Host == "" {
		return Origin{}, false
	}
	authority, err := url.Parse("https://" + u.Host)
	if err != nil || authority.Host != u.Host || authority.User != nil {
		return Origin{}, false
	}
	host := authority.Hostname()
	if host == "" || strings.ContainsAny(host, "%\\/?#@ \t\r\n") {
		return Origin{}, false
	}
	if strings.Contains(host, ":") && net.ParseIP(host) == nil {
		return Origin{}, false
	}
	// Match net/http's IDNA lookup conversion and ASCII fast path, without
	// folding Unicode characters into a different transport destination.
	if strings.ContainsFunc(host, func(r rune) bool { return r >= utf8.RuneSelf }) {
		host, err = idna.Lookup.ToASCII(host)
		if err != nil {
			return Origin{}, false
		}
	}
	port := authority.Port()
	if port == "" {
		if strings.HasSuffix(authority.Host, ":") {
			return Origin{}, false
		}
		port = "443"
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return Origin{}, false
	}
	return Origin{host: strings.ToLower(host), port: strconv.Itoa(n)}, true
}
