package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var (
	alphaTypeName = MakeTypeName(GenRuntimeReference, "Alpha")
	betaTypeName  = MakeTypeName(GenRuntimeReference, "Beta")
	gammaTypeName = MakeTypeName(GenRuntimeReference, "Gamma")
)

/*
 * Construction Tests
 */

func Test_TypeNameProcessingRegister_AfterConstruction_IsEmpty(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	g.Expect(register.Length()).To(Equal(0))
}

/*
 * Empty Register tests
 */

func Test_EmptyTypeNameRegister_DoesNotContainTypeNames(t *testing.T) {
	g := NewGomegaWithT(t)
	emptyRegister := NewTypeNameProcessingRegister()
	g.Expect(emptyRegister.IsCompleted(alphaTypeName)).To(BeFalse())
	g.Expect(emptyRegister.IsPending(alphaTypeName)).To(BeFalse())
}

/*
 * Requires() tests
 */

func Test_TypeNameProcessingRegister_Requires_IncreasesLengthOfList(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(betaTypeName, "because")
	g.Expect(register.Length()).To(Equal(1))
}

func Test_TypeNameProcessingRegister_Requires_NewNameIsPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(betaTypeName, "because")
	g.Expect(register.IsPending(betaTypeName)).To(BeTrue())
}

func Test_TypeNameProcessingRegister_Requires_NewNameIsNotComplete(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(betaTypeName, "because")
	g.Expect(register.IsCompleted(betaTypeName)).To(BeFalse())
}

func Test_TypeNameProcessingRegister_Requires_DoesNotResetCompletedStatus(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(betaTypeName, "because")
	register.Complete(betaTypeName)
	register.Requires(betaTypeName, "because")
	g.Expect(register.IsPending(betaTypeName)).To(BeFalse())
	g.Expect(register.IsCompleted(betaTypeName)).To(BeTrue())
}

/*
 * Complete() Tests
 */

func Test_TypeNameProcessingRegister_Complete_MarksNameAsCompleted(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(gammaTypeName, "because")
	register.Complete(gammaTypeName)
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeTrue())
}

func Test_TypeNameProcessingRegister_Complete_MarksNameNotPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(gammaTypeName, "because")
	register.Complete(gammaTypeName)
	g.Expect(register.IsPending(gammaTypeName)).To(BeFalse())
}

func Test_TypeNameProcessingRegister_Complete_MakesRegisterLargerIfMissing(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	initialLength := register.Length()
	register.Complete(gammaTypeName)
	g.Expect(register.Length()).To(BeNumerically(">", initialLength))
}

func Test_TypeNameProcessingRegister_Complete_DoesNotIncludeNameInPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Complete(gammaTypeName)
	g.Expect(register.Pending()).NotTo(HaveKey(gammaTypeName))
}

func Test_TypeNameProcessingRegister_Complete_IncludesNameInCompleted(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Complete(gammaTypeName)
	g.Expect(register.Completed()).To(HaveKey(gammaTypeName))
}

/*
 * Reset() Tests
 */

func Test_TypeNameProcessingRegister_Reset_MarksNameAsPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(gammaTypeName, "because")
	register.Complete(gammaTypeName)
	register.Reset(gammaTypeName)
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeFalse())
}

func Test_TypeNameProcessingRegister_Reset_MarksNameAsNotComplete(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Requires(gammaTypeName, "because")
	register.Complete(gammaTypeName)
	register.Reset(gammaTypeName)
	g.Expect(register.IsPending(gammaTypeName)).To(BeTrue())
}

func Test_TypeNameProcessingRegister_Reset_MakesRegisterLargerIfMissing(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	initialLength := register.Length()
	register.Reset(gammaTypeName)
	g.Expect(register.Length()).To(BeNumerically(">", initialLength))
}

func Test_TypeNameProcessingRegister_Reset_IncludesNameInPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Reset(gammaTypeName)
	g.Expect(register.Pending()).To(HaveKey(gammaTypeName))
}

func Test_TypeNameProcessingRegister_Reset_DoesNotIncludeNameInCompleted(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Reset(gammaTypeName)
	g.Expect(register.Completed()).NotTo(HaveKey(gammaTypeName))
}

/*
 * Remove() tests
 */

func Test_TypeNameProcessingRegister_Remove_NameIsNoLongerPending(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Reset(gammaTypeName)
	g.Expect(register.IsPending(gammaTypeName)).To(BeTrue())
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeFalse())
	register.Remove(gammaTypeName)
	g.Expect(register.IsPending(gammaTypeName)).To(BeFalse())
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeFalse())
}

func Test_TypeNameProcessingRegister_Remove_NameIsNoLongerCompleted(t *testing.T) {
	g := NewGomegaWithT(t)
	register := NewTypeNameProcessingRegister()
	register.Complete(gammaTypeName)
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeTrue())
	g.Expect(register.IsPending(gammaTypeName)).To(BeFalse())
	register.Remove(gammaTypeName)
	g.Expect(register.IsCompleted(gammaTypeName)).To(BeFalse())
	g.Expect(register.IsPending(gammaTypeName)).To(BeFalse())
}