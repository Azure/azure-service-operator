/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package typo

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/rotisserie/eris"
)

func TestTypoAdvisorErrorf_WhenNoTerms_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor()
	err := advisor.Errorf("typo", "%s went boom!", "typo")
	g.Expect(err.Error()).To(Equal("typo went boom!"))
}

func TestTypoAdvisorErrorf_WhenTermsAvailable_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	err := advisor.Errorf("aleph", "%s went boom!", "aleph")
	g.Expect(err.Error()).To(Equal("aleph went boom! (did you mean alpha?)"))
}

func TestTypoAdvisor_Wrapf_WhenNoError_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	g.Expect(advisor.Wrapf(nil, "beat", "format string")).To(BeNil())
}

func TestTypoAdvisor_Wrapf_WhenErrorButNoTerms_ReturnsOriginalError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor()
	err := eris.New("Boom")
	g.Expect(advisor.Wrapf(err, "beat", "format string")).To(Equal(err))
}

func TestTypoAdvisor_Wrapf_WhenErrorButNoTypo_ReturnsOriginalError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	err := eris.New("Boom")
	g.Expect(advisor.Wrapf(err, "beta", "format string")).To(Equal(err))
}

func TestTypoAdvisor_Wrapf_WhenErrorAndTypo_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	err := eris.New("Boom")
	actual := advisor.Wrapf(err, "beat", "the typo was %s", "beat")

	g.Expect(actual).NotTo(BeNil())
	g.Expect(actual.Error()).To(ContainSubstring(err.Error()))
	g.Expect(actual.Error()).To(ContainSubstring("the typo was beat"))
	g.Expect(actual.Error()).To(ContainSubstring("did you mean beta?"))
}

func TestTypoAdvisor_Suggest_WhenNoTerms_ReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor()
	_, ok := advisor.Suggest("FirstName")
	g.Expect(ok).To(BeFalse())
}

func TestTypoAdvisor_Suggest_WhenTermsAvailable_ReturnsExpectedSuggestion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("LegalName", "KnownAs", "BirthDate")
	suggestion, ok := advisor.Suggest("FamilyName")
	g.Expect(ok).To(BeTrue())
	g.Expect(suggestion).To(Equal("LegalName"))
}

func TestTypoAdvisor_Suggest_IsCaseInsensitive(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("KNOWNAS")
	suggestion, ok := advisor.Suggest("knownas")
	g.Expect(ok).To(BeTrue())
	g.Expect(suggestion).To(Equal("KNOWNAS"))
}

func createTestTypoAdvisor(terms ...string) *Advisor {
	result := NewAdvisor()
	for _, term := range terms {
		result.AddTerm(term)
	}

	return result
}
