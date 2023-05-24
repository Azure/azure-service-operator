/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package typo

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
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
	err := errors.New("Boom")
	g.Expect(advisor.Wrapf(err, "beat", "format string")).To(Equal(err))
}

func TestTypoAdvisor_Wrapf_WhenErrorButNoTypo_ReturnsOriginalError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	err := errors.New("Boom")
	g.Expect(advisor.Wrapf(err, "beta", "format string")).To(Equal(err))
}

func TestTypoAdvisor_Wrapf_WhenErrorAndTypo_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	advisor := createTestTypoAdvisor("alpha", "beta", "gamma", "delta")
	err := errors.New("Boom")
	actual := advisor.Wrapf(err, "beat", "the typo was %s", "beat")

	g.Expect(actual).NotTo(BeNil())
	g.Expect(actual.Error()).To(ContainSubstring(err.Error()))
	g.Expect(actual.Error()).To(ContainSubstring("the typo was beat"))
	g.Expect(actual.Error()).To(ContainSubstring("did you mean beta?"))
}

func createTestTypoAdvisor(terms ...string) *Advisor {
	result := NewAdvisor()
	for _, term := range terms {
		result.AddTerm(term)
	}

	return result
}
