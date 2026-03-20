/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_AsNewStylePackageReference(t *testing.T) {
	t.Parallel()

	// A legacy-style (v1api prefix) local package reference
	legacyRef := astmodel.MakeVersionedLocalPackageReference(test.GoModulePrefix, "batch", "20200101").
		WithVersionPrefix("v1api")

	// A new-style (v prefix) local package reference
	newStyleRef := astmodel.MakeVersionedLocalPackageReference(test.GoModulePrefix, "compute", "20200101").
		WithVersionPrefix("v")

	cases := map[string]struct {
		ref         astmodel.PackageReference
		expectOk    bool
		expectEqual astmodel.InternalPackageReference // expected result when ok is true
	}{
		"ExternalPackageReference returns nil, false": {
			ref:      astmodel.MakeExternalPackageReference("fmt"),
			expectOk: false,
		},
		"LocalPackageReference with v1api prefix returns converted reference": {
			ref:         legacyRef,
			expectOk:    true,
			expectEqual: legacyRef.WithVersionPrefix("v"),
		},
		"LocalPackageReference without v1api prefix returns nil, false": {
			ref:      newStyleRef,
			expectOk: false,
		},
		"SubPackageReference with convertible parent returns new sub with converted parent": {
			ref:         astmodel.MakeSubPackageReference("customizations", legacyRef),
			expectOk:    true,
			expectEqual: astmodel.MakeSubPackageReference("customizations", legacyRef.WithVersionPrefix("v")),
		},
		"SubPackageReference with non-convertible parent returns nil, false": {
			ref:      astmodel.MakeSubPackageReference("customizations", newStyleRef),
			expectOk: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result, ok := asNewStylePackageReference(c.ref)
			g.Expect(ok).To(Equal(c.expectOk))
			if c.expectOk {
				g.Expect(result.Equals(c.expectEqual)).To(BeTrue())
			} else {
				g.Expect(result).To(BeNil())
			}
		})
	}
}
