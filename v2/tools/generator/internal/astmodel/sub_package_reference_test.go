/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewSubPackageReference_GivenParentAndName_ReturnsExpectedProperties(t *testing.T) {
	t.Parallel()

	parent := makeTestLocalPackageReference("group", "version")
	batch := makeTestLocalPackageReference("microsoft.batch", "2020-09-01")

	cases := []struct {
		name                string
		parent              PackageReference
		expectedPackageName string
		expectedPackagePath string
	}{
		{
			"simple",
			parent,
			"simple",
			"github.com/Azure/azure-service-operator/v2/api/group/version/simple",
		},
		{
			"sub",
			batch,
			"sub",
			"github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v20200901/sub",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			subPackage := MakeSubPackageReference(c.name, c.parent)

			g.Expect(subPackage.name).To(Equal(c.name))
			g.Expect(subPackage.parent).To(Equal(c.parent))
			g.Expect(subPackage.PackageName()).To(Equal(c.expectedPackageName))
			g.Expect(subPackage.PackagePath()).To(Equal(c.expectedPackagePath))
		})
	}
}

func Test_SubPackageReference_GivenParent_InheritsParentsProperties(t *testing.T) {
	t.Parallel()

	network := makeTestLocalPackageReference("microsoft.network", "2023-07-01")
	networkPreview := makeTestLocalPackageReference("microsoft.network", "2023-07-01-preview")

	cases := []struct {
		name      string
		parent    PackageReference
		group     string
		version   string
		isPreview bool
	}{
		{
			"Network",
			network,
			"microsoft.network",
			"v20230701",
			false,
		},
		{
			"NetworkPreview",
			networkPreview,
			"microsoft.network",
			"v20230701preview",
			true,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			sub := MakeSubPackageReference("sub", c.parent)

			grp, ver := sub.GroupVersion()
			g.Expect(grp).To(Equal(c.group))
			g.Expect(ver).To(Equal(c.version))
			g.Expect(sub.IsPreview()).To(Equal(c.isPreview))

			grp, ver = c.parent.GroupVersion()
			g.Expect(grp).To(Equal(c.group))
			g.Expect(ver).To(Equal(c.version))
			g.Expect(c.parent.IsPreview()).To(Equal(c.isPreview))
		})
	}
}
