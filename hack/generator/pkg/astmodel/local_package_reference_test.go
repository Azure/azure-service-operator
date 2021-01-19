/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMakeLocalPackageReference_GivenGroupAndPackage_ReturnsInstanceWithProperties(t *testing.T) {
	cases := []struct {
		name  string
		group string
		pkg   string
	}{
		{"Networking", "microsoft.networking", "v20200901"},
		{"Batch (new)", "microsoft.batch", "v20200901"},
		{"Batch (old)", "microsoft.batch", "v20150101"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeLocalPackageReference(c.group, c.pkg)
			grp := ref.Group()
			pkg := ref.PackageName()

			g.Expect(grp).To(Equal(c.group))
			g.Expect(pkg).To(Equal(c.pkg))
		})
	}
}

func TestLocalPackageReferences_ReturnExpectedProperties(t *testing.T) {
	cases := []struct {
		name         string
		group        string
		pkg          string
		expectedPath string
	}{
		{
			"Networking",
			"microsoft.networking",
			"v20200901",
			"github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.networking/v20200901",
		},
		{
			"Batch (new)",
			"microsoft.batch",
			"v20200901",
			"github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v20200901",
		},
		{
			"Batch (old)",
			"microsoft.batch",
			"v20150101",
			"github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v20150101",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ref := MakeLocalPackageReference(c.group, c.pkg)
			grp := ref.Group()
			_, ok := ref.AsLocalPackage()

			g.Expect(ok).To(BeTrue())
			g.Expect(ref.PackageName()).To(Equal(c.pkg))
			g.Expect(ref.PackagePath()).To(Equal(c.expectedPath))
			g.Expect(ref.String()).To(Equal(c.expectedPath))
			g.Expect(grp).To(Equal(c.group))
		})
	}
}

func TestLocalPackageReferences_Equals_GivesExpectedResults(t *testing.T) {

	batchRef := MakeLocalPackageReference("microsoft.batch", "v20200901")
	olderRef := MakeLocalPackageReference("microsoft.batch", "v20150101")
	networkingRef := MakeLocalPackageReference("microsoft.networking", "v20200901")
	fmtRef := MakeExternalPackageReference("fmt")

	cases := []struct {
		name     string
		this     LocalPackageReference
		other    PackageReference
		areEqual bool
	}{
		{"Equal self", batchRef, batchRef, true},
		{"Equal self", olderRef, olderRef, true},
		{"Not equal other library name", batchRef, networkingRef, false},
		{"Not equal other library name", networkingRef, batchRef, false},
		{"Not equal other library version", batchRef, olderRef, false},
		{"Not equal other library version", olderRef, batchRef, false},
		{"Not equal other kind", batchRef, fmtRef, false},
		{"Not equal other kind", networkingRef, fmtRef, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.this.Equals(c.other)
			g.Expect(areEqual).To(Equal(c.areEqual))
		})
	}
}
