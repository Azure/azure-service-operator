/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/google/uuid"
)

type sampleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Reference         genruntime.ResourceReference
	SubscriptionID    string
	TenantID          string
}

var _ genruntime.ARMMetaObject = &sampleResource{}

func Test_SamplesTester_UpdatesSubscriptionReferences(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		azureSubscription: uuid.New().String(),
	}

	sample := &sampleResource{
		Reference: genruntime.CreateResourceReferenceFromARMID("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm"),
	}

	err := tester.updateFieldsForTest(sample)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(sample.Reference.ARMID).To(ContainSubstring(tester.azureSubscription))
}

func Test_SamplesTester_UpdatesSubscriptionID(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		azureSubscription: uuid.New().String(),
	}

	sample := &sampleResource{
		SubscriptionID: emptyGuid,
	}

	err := tester.updateFieldsForTest(sample)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(sample.SubscriptionID).To(Equal(tester.azureSubscription))
}

func Test_SamplesTester_UpdatesTenantID(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		azureTenant: uuid.New().String(),
	}

	sample := &sampleResource{
		TenantID: emptyGuid,
	}

	err := tester.updateFieldsForTest(sample)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(sample.TenantID).To(Equal(tester.azureTenant))
}

// GetConditions implements genruntime.ARMMetaObject.
func (*sampleResource) GetConditions() conditions.Conditions {
	panic("unimplemented")
}

// GetResourceScope implements genruntime.ARMMetaObject.
func (*sampleResource) GetResourceScope() genruntime.ResourceScope {
	panic("unimplemented")
}

// GetSpec implements genruntime.ARMMetaObject.
func (*sampleResource) GetSpec() genruntime.ConvertibleSpec {
	panic("unimplemented")
}

// GetStatus implements genruntime.ARMMetaObject.
func (*sampleResource) GetStatus() genruntime.ConvertibleStatus {
	panic("unimplemented")
}

// GetType implements genruntime.ARMMetaObject.
func (*sampleResource) GetType() string {
	panic("unimplemented")
}

// NewEmptyStatus implements genruntime.ARMMetaObject.
func (*sampleResource) NewEmptyStatus() genruntime.ConvertibleStatus {
	panic("unimplemented")
}

// Owner implements genruntime.ARMMetaObject.
func (*sampleResource) Owner() *genruntime.ResourceReference {
	panic("unimplemented")
}

// SetConditions implements genruntime.ARMMetaObject.
func (*sampleResource) SetConditions(conditions conditions.Conditions) {
	panic("unimplemented")
}

// SetStatus implements genruntime.ARMMetaObject.
func (*sampleResource) SetStatus(status genruntime.ConvertibleStatus) error {
	panic("unimplemented")
}

func (s *sampleResource) AzureName() string {
	return "azure"
}

func (s *sampleResource) DeepCopyObject() runtime.Object {
	panic("not implemented")
}

func (s *sampleResource) GetAPIVersion() string {
	panic("not implemented")
}
