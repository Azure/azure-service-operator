/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package labels_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/labels"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label     string
		wantKey   string
		wantValue string
		wantErr   bool
	}{
		{"example.com/label=value", "example.com/label", "value", false},
		{"example.com/label=", "example.com/label", "", false},
		{"=value", "", "", true},
		{"example.com/label", "", "", true},
		{"example.com/test/label", "", "", true},
		{"thisisaverylonglabelname_solonginfactthatitisgoingtocauseanerror", "", "", true},
		{"", "", "", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, err := labels.Parse(tt.label)

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
