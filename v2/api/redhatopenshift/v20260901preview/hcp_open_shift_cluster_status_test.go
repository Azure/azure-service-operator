package v20260901preview

import (
	"testing"

	. "github.com/onsi/gomega"

	arm "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v20260901preview/arm"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestHcpOpenShiftClusterPropertiesStatus_PopulateFromARM_PreservesConditions(t *testing.T) {
	g := NewWithT(t)

	typeValue := arm.ConditionType_STATUS("RequirementsValid")
	statusValue := arm.StatusType_STATUS_False
	message := "control plane operators are missing required permissions"

	properties := HcpOpenShiftClusterProperties_STATUS{}
	err := properties.PopulateFromARM(genruntime.ArbitraryOwnerReference{}, arm.HcpOpenShiftClusterProperties_STATUS{
		Status: &arm.ResourceStatus_STATUS{
			Conditions: []arm.Condition_STATUS{{
				Type:    &typeValue,
				Status:  &statusValue,
				Message: &message,
			}},
		},
	})
	g.Expect(err).To(Succeed())
	g.Expect(properties.Status).NotTo(BeNil())
	g.Expect(properties.Status.Conditions).To(HaveLen(1))
	g.Expect(string(*properties.Status.Conditions[0].Type)).To(Equal("RequirementsValid"))
	g.Expect(string(*properties.Status.Conditions[0].Status)).To(Equal("False"))
	g.Expect(*properties.Status.Conditions[0].Message).To(Equal(message))
}
