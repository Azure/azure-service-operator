/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestInterleaveStringSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        []string
		b        []string
		expected []string
	}{
		{
			name:     "Slices of same length are interleaved correctly",
			a:        []string{"1", "2", "3", "4"},
			b:        []string{"alpha", "beta", "delta", "gamma"},
			expected: []string{"1", "alpha", "2", "beta", "3", "delta", "4", "gamma"},
		},
		{
			name:     "Shorter first slice interleaved correctly",
			a:        []string{"1", "2"},
			b:        []string{"alpha", "beta", "delta", "gamma"},
			expected: []string{"1", "alpha", "2", "beta", "delta", "gamma"},
		},
		{
			name:     "Shorter second slice interleaved correctly",
			a:        []string{"1", "2", "3", "4"},
			b:        []string{"alpha", "beta"},
			expected: []string{"1", "alpha", "2", "beta", "3", "4"},
		},
		{
			name:     "Empty first slice",
			a:        nil,
			b:        []string{"alpha", "beta", "delta", "gamma"},
			expected: []string{"alpha", "beta", "delta", "gamma"},
		},
		{
			name:     "Empty second slice",
			a:        []string{"1", "2", "3", "4"},
			b:        nil,
			expected: []string{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := InterleaveStrSlice(tt.a, tt.b)
			g.Expect(result).To(Equal(tt.expected))
		})
	}
}
