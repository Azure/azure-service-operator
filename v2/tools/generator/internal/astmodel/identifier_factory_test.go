/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func Test_CreateIdentifier_GivenName_ReturnsExpectedIdentifier(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		visibility Visibility
		expected   string
	}{
		{"name", Exported, "Name"},
		{"Name", Exported, "Name"},
		{"$schema", Exported, "Schema"},
		{"my_important_name", Exported, "My_Important_Name"},
		{"MediaServices_liveEvents_liveOutputs_childResource", Exported, "MediaServices_LiveEvents_LiveOutputs_ChildResource"},
		{"XMLDocument", Exported, "XMLDocument"},
		{"this id has spaces", Exported, "ThisIdHasSpaces"},
		{"this, id, has, spaces", Exported, "ThisIdHasSpaces"},
		{"package", Exported, "Package"},
		{"name", NotExported, "name"},
		{"Name", NotExported, "name"},
		{"$schema", NotExported, "schema"},
		{"my_important_name", NotExported, "my_Important_Name"},
		{"MediaServices_liveEvents_liveOutputs_childResource", NotExported, "mediaServices_LiveEvents_LiveOutputs_ChildResource"},
		{"XMLDocument", NotExported, "xmlDocument"},
		{"this id has spaces", NotExported, "thisIdHasSpaces"},
		{"this, id, has, spaces", NotExported, "thisIdHasSpaces"},
		{"package", NotExported, "pkg"},
		{"item1ARM", NotExported, "item1ARM"},
		{"itemARM", NotExported, "itemARM"},
		{"XML1Document", NotExported, "xml1Document"},
		{"book12was2good", NotExported, "book12Was2Good"},
		{"12monkeys", NotExported, "12Monkeys"},
		{"version20210201", NotExported, "version20210201"},
		{"version2021 02 01", NotExported, "version20210201"},
	}

	idfactory := NewIdentifierFactory()

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			identifier := idfactory.CreateIdentifier(c.name, c.visibility)
			g.Expect(identifier).To(Equal(c.expected))
		})
	}
}

func Test_SliceIntoWords_GivenIdentifier_ReturnsExpectedSlice(t *testing.T) {
	t.Parallel()

	cases := []struct {
		identifier string
		expected   []string
	}{
		// Single name doesn't get split
		{identifier: "Name", expected: []string{"Name"}},
		// Single Acronym doesn't get split
		{identifier: "XML", expected: []string{"XML"}},
		// Splits simple words
		{identifier: "PascalCase", expected: []string{"Pascal", "Case"}},
		{identifier: "XmlDocument", expected: []string{"Xml", "Document"}},
		// Correctly splits all-caps acronyms
		{identifier: "XMLDocument", expected: []string{"XML", "Document"}},
		{identifier: "ResultAsXML", expected: []string{"Result", "As", "XML"}},
		// Correctly splits strings that already have spaces
		{identifier: "AlreadyHas spaces", expected: []string{"Already", "Has", "spaces"}},
		{identifier: "Already   Has  spaces    ", expected: []string{"Already", "Has", "spaces"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.identifier, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			actual := sliceIntoWords(c.identifier)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func Test_TransformToSnakeCase_ReturnsExpectedString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		string   string
		expected string
	}{
		// Single name doesn't get split
		{string: "Name", expected: "name"},
		// Single Acronym doesn't get split
		{string: "XML", expected: "xml"},
		// Splits simple words
		{string: "PascalCase", expected: "pascal_case"},
		{string: "XmlDocument", expected: "xml_document"},
		// Correctly splits all-caps acronyms
		{string: "XMLDocument", expected: "xml_document"},
		{string: "ResultAsXML", expected: "result_as_xml"},
		// Correctly transforms strings with spaces
		{string: "AlreadyHas spaces", expected: "already_has_spaces"},
		{string: "AlreadyHas spaces    ", expected: "already_has_spaces"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.string, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			actual := transformToSnakeCase(c.string)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func Test_SimplifyIdentifier_GivenContextAndName_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := []struct {
		context    string
		identifier string
		expected   string
	}{
		// No change if no overlap
		{context: "Person", identifier: "FullName", expected: "FullName"},
		{context: "Person", identifier: "KnownAs", expected: "KnownAs"},
		{context: "Person", identifier: "FamilyName", expected: "FamilyName"},
		// Removes Single words
		{context: "SecurityPolicy", identifier: "PolicyDefaultDeny", expected: "DefaultDeny"},
		{context: "SecurityPolicy", identifier: "SecurityException", expected: "Exception"},
		// Removes Multiple words, regardless of order
		{context: "SecurityPolicy", identifier: "SecurityPolicyType", expected: "Type"},
		{context: "SecurityPolicy", identifier: "PolicyTypeSecurity", expected: "Type"},
		// Removes words only once
		{context: "SecurityPolicy", identifier: "SecurityPolicyPolicyType", expected: "PolicyType"},
		// Won't return an empty string
		{context: "SecurityPolicy", identifier: "SecurityPolicy", expected: "SecurityPolicy"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.identifier, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			actual := simplifyName(c.context, c.identifier)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func Test_CreateReceiver_GivenTypeName_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected string
	}{
		// Base cases
		{"Address", "address"},
		// Forbidden receiver suffixes
		{"Address_STATUS", "address"},
		{"Address_Spec", "address"},
		{"Address_SpecARM", "address"},
		// Real world examples
		{"EncryptionSettingsCollection", "collection"},
		{"RedisLinkedServer", "server"},
		{"FlexibleServersConfiguration", "configuration"},
		{"CompositePath_STATUS", "path"},
		// Long examples
		{"VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile", "profile"},
		{"VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions", "extensions"},
		{"ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes", "prefixes"},
		{"DatabaseAccountsMongodbDatabasesCollections_Spec", "collections"},
		{"DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec", "settings"},
		// Very short receiver names need more detail
		{"SignalR_SpecARM", "signalR"},
		{"PublicIPAddressSku_STATUS", "addressSku"},
		{"SBSku_STATUS", "sbSku"},
		{"ManagedClusterSKU", "clusterSKU"},
		{"AdvancedFilter_NumberIn", "numberIn"},
		{"AdvancedFilter_NumberNotIn", "notIn"},
		{"DiskSku_STATUS", "diskSku"},
		// Conflicts with reserved words need more detail
		{"BlobRestoreRange_STATUS", "restoreRange"},
	}

	factory := NewIdentifierFactory()
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			actual := factory.CreateReceiver(c.name)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func TestAppendPreservingSuffixPreservesSuffix(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)
	str := "something3455"
	result := AppendPreservingSuffix(str, "11", "3456")
	g.Expect(result).To(gomega.Equal("something123455"))
}

func TestAppendPreservingSuffixWithoutSuffix(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)
	str := "something_suffix"
	result := AppendPreservingSuffix(str, "_new", "_elsewise")
	g.Expect(result).To(gomega.Equal("something_suffix_new"))
}
