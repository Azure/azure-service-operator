/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package annotations_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/annotations"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		annotation string
		wantKey    string
		wantValue  string
		wantErr    bool
	}{
		{"example.com/annotation=value", "example.com/annotation", "value", false},
		{"example.com/annotation=", "example.com/annotation", "", false},
		{"=value", "", "", true},
		{"example.com/annotation", "", "", true},
		{"example.com/test/annotation", "", "", true},
		{"thisisaverylongannotationname_solonginfactthatitisgoingtocauseanerror", "", "", true},
		{"", "", "", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.annotation, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, err := annotations.Parse(tt.annotation)

			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}

			g.Expect(actual.Key).To(Equal(tt.wantKey))
			g.Expect(actual.Value).To(Equal(tt.wantValue))
		})
	}
}
