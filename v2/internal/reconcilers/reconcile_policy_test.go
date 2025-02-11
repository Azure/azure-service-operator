package reconcilers

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

func TestParseReconcilePolicy(t *testing.T) {
	testingMap := map[string]annotations.ReconcilePolicyValue{
		"manage":           annotations.ReconcilePolicyManage,
		"skip":             annotations.ReconcilePolicySkip,
		"detach-on-delete": annotations.ReconcilePolicyDetachOnDelete,
	}

	t.Parallel()
	g := NewGomegaWithT(t)

	// test nominal case
	for policyString, policyValue := range testingMap {
		returnedPolicy, err := ParseReconcilePolicy(policyString, string(policyValue))
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))
	}

	// test default value
	returnedPolicy, err := ParseReconcilePolicy("", string(annotations.ReconcilePolicySkip))
	g.Expect(err).ToNot((HaveOccurred()))
	g.Expect(returnedPolicy).Should(Equal(annotations.ReconcilePolicySkip))

	// test error in case of any other value
	returnedPolicy, err = ParseReconcilePolicy("whatever", string(annotations.ReconcilePolicySkip))
	g.Expect(err).Should(HaveOccurred())

}
