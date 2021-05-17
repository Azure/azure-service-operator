/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestPackageGroupVersion_IncludesGeneratorVersion(t *testing.T) {
	g := NewGomegaWithT(t)

	pkg := NewPackageDefinition("longest-johns", "santiana", "latest-version")
	pkgs := make(map[PackageReference]*PackageDefinition)
	pkgs[makeTestLocalPackageReference(pkg.GroupName, pkg.PackageName)] = pkg

	destDir, err := ioutil.TempDir("", "package_definition_test")
	g.Expect(err).To(BeNil())
	defer os.RemoveAll(destDir)

	fileCount, err := pkg.EmitDefinitions(destDir, pkgs)
	g.Expect(err).To(BeNil())
	g.Expect(fileCount).To(Equal(0))

	gvFile := filepath.Join(destDir, "groupversion_info_gen.go")
	data, err := ioutil.ReadFile(gvFile)
	g.Expect(err).To(BeNil())
	g.Expect(string(data)).To(ContainSubstring("// Generator version: latest-version"))
}

func Test_PropertyDefinitionWithTag_GivenTag_DoesNotModifyOriginal(t *testing.T) {
	g := NewGomegaWithT(t)
	original := NewPropertyDefinition("FullName", "fullName", StringType)
	original = original.WithTag("a", "b")
	field := original.WithTag("c", "d")

	g.Expect(field.tags).NotTo(Equal(original.tags))
}
