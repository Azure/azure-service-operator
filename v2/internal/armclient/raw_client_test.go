/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestParseHttpDate(t *testing.T) {
	g := NewGomegaWithT(t)

	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}

	expected := time.Date(1994, time.November, 6, 8, 49, 37, 0, gmt)

	// date format examples yanked from: https://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.3
	// see also: https://tools.ietf.org/html/rfc7231#section-7.1.1.1
	for _, dateFormat := range []string{
		"Sun, 06 Nov 1994 08:49:37 GMT",  // IMF-fixdate
		"Sunday, 06-Nov-94 08:49:37 GMT", // obs-date (rfc850-date)
		"Sun Nov  6 08:49:37 1994",       // obs-date (asctime-date)
		"Sun Nov 06 08:49:37 1994",       // modified version of last to have 2 digits
	} {
		parsed, err := parseHttpDate(dateFormat)
		g.Expect(err).ToNot(HaveOccurred())
		// can't use gomega assertion here as .Equal must be called
		g.Expect(parsed.Equal(expected)).To(BeTrue())
	}
}
