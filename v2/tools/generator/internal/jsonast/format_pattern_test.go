/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"encoding/base64"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
)

// The base64url pattern once lacked a quantifier, so the CRD accepted exactly one character and no
// real blob could pass admission. Realistic values, padded or not, must match.
func Test_FormatToPattern_Base64URL(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pattern := formatToPattern("base64url", logr.Discard())
	g.Expect(pattern).ToNot(BeNil())

	policy := []byte(`{"anyOf":[{"allOf":[{"claim":"x-ms-sgx-is-debuggable","equals":"false"}]}]}`)
	g.Expect(pattern.MatchString(base64.RawURLEncoding.EncodeToString(policy))).To(BeTrue())
	g.Expect(pattern.MatchString(base64.URLEncoding.EncodeToString(policy))).To(BeTrue())
	g.Expect(pattern.MatchString("")).To(BeTrue())

	// Standard base64 alphabet characters and whitespace are not base64url
	g.Expect(pattern.MatchString("abc+/def")).To(BeFalse())
	g.Expect(pattern.MatchString("abc def")).To(BeFalse())
	g.Expect(pattern.MatchString("abc===")).To(BeFalse())
}
