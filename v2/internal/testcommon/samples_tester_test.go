/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type sampleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Reference         genruntime.ResourceReference
	SubscriptionID    string
	TenantID          string
}

var _ genruntime.ARMMetaObject = &sampleResource{}

func Test_SamplesTester_UpdatesFieldsForTest(t *testing.T) {
	t.Parallel()

	subscription := uuid.New().String()
	tenant := uuid.New().String()
	rgName := "somerandomrg"

	cases := []struct {
		name          string
		useRandomName bool
		rgName        string
		subscription  string
		tenant        string
		nameRenames   map[referenceKey]string
		sample        *sampleResource
		check         func(g Gomega, sample *sampleResource)
	}{
		{
			name:         "updates ARMID reference with subscription",
			subscription: subscription,
			sample: &sampleResource{
				Reference: genruntime.CreateResourceReferenceFromARMID("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm"),
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.Reference.ARMID).To(ContainSubstring(subscription))
			},
		},
		{
			name:         "updates subscription-only ARMID reference",
			subscription: subscription,
			sample: &sampleResource{
				Reference: genruntime.CreateResourceReferenceFromARMID("subscriptions/00000000-0000-0000-0000-000000000000"),
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.Reference.ARMID).To(ContainSubstring(subscription))
			},
		},
		{
			name:         "updates SubscriptionID field",
			subscription: subscription,
			sample: &sampleResource{
				SubscriptionID: emptyGUID,
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.SubscriptionID).To(Equal(subscription))
			},
		},
		{
			name:   "updates TenantID field",
			tenant: tenant,
			sample: &sampleResource{
				TenantID: emptyGUID,
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.TenantID).To(Equal(tenant))
			},
		},
		{
			name:   "updates ResourceGroup reference name",
			rgName: rgName,
			sample: &sampleResource{
				Reference: genruntime.ResourceReference{
					Group: "resources.azure.com",
					Kind:  "ResourceGroup",
					Name:  "aso-sample-rg",
				},
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.Reference.Name).To(Equal(rgName))
			},
		},
		{
			name:          "updates reference name when useRandomNames is true",
			useRandomName: true,
			nameRenames: map[referenceKey]string{
				{Kind: "VirtualNetwork", Name: "my-virtual-network"}: "randomname123",
			},
			sample: &sampleResource{
				Reference: genruntime.ResourceReference{
					Group: "network.azure.com",
					Kind:  "VirtualNetwork",
					Name:  "my-virtual-network",
				},
			},
			check: func(g Gomega, sample *sampleResource) {
				g.Expect(sample.Reference.Name).To(Equal("randomname123"))
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			tester := NewSamplesTester(
				ResourceNamer{},
				nil,
				"",
				"",
				c.useRandomName,
				c.rgName,
				c.subscription,
				c.tenant,
			)

			for k, v := range c.nameRenames {
				tester.nameRenames[k] = v
			}

			err := tester.updateFieldsForTest(c.sample)
			g.Expect(err).ToNot(HaveOccurred())

			c.check(g, c.sample)
		})
	}
}

func (s *sampleResource) GetSupportedOperations() []genruntime.ResourceOperation {
	panic("unimplemented")
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
