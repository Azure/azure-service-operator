/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package template_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/internal/template"
)

func Test_Apply(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		template       string
		crdPattern     string
		expected       string
		expectedErrMsg string
	}{
		{
			name:           "no crd-pattern",
			template:       "a template without crd-pattern",
			crdPattern:     "*",
			expectedErrMsg: "failed to inject crd-pattern",
		},
		{
			name: "simple crd-pattern",
			template: `
	parameters:
		- --other-parameter=12
		- --crd-pattern=
`,
			crdPattern: "*",
			expected: `
	parameters:
		- --other-parameter=12
		- --crd-pattern=*
`,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c := c
			t.Parallel()

			g := NewGomegaWithT(t)

			result, err := template.Apply(c.template, c.crdPattern)
			if c.expectedErrMsg != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedErrMsg)))
			} else {
				g.Expect(c.expected).To(Equal(result))
			}
		})
	}
}

func Test_GetFromLocal_ReturnsCorrectly(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()

	data, err := template.Get(ctx, "testdata/local.yaml")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(data).ToNot(BeEmpty())
}

func Test_GetFromRemote_ReturnsCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte("this is a test"))).ToNot(BeZero())
			return
		}

		g.Fail(fmt.Sprintf("unknown request attempted. Method: %s, URL: %s", r.Method, r.URL))
	}))
	defer server.Close()

	ctx := context.Background()

	data, err := template.Get(ctx, server.URL)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(data).ToNot(BeEmpty())
}
