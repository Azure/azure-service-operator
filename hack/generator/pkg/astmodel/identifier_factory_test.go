/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CreateIdentifier_GivenName_ReturnsExpectedIdentifier(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{"name", "Name"},
		{"Name", "Name"},
		{"$schema", "Schema"},
		{"my_important_name", "MyImportantName"},
		{"MediaServices_liveEvents_liveOutputs_childResource", "MediaServicesLiveEventsLiveOutputsChildResource"},
	}

	idfactory := NewIdentifierFactory()

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			identifier := idfactory.CreateIdentifier(c.name)
			g.Expect(identifier).To(Equal(c.expected))
		})
	}
}

func Test_SliceIntoWords_GivenIdentifier_ReturnsExpectedSlice(t *testing.T) {
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
