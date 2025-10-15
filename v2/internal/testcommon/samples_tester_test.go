/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

func Test_SamplesTester_UpdatesSubscriptionOnlyReference(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		azureSubscription: uuid.New().String(),
	}

	sample := &sampleResource{
		Reference: genruntime.CreateResourceReferenceFromARMID("subscriptions/00000000-0000-0000-0000-000000000000"),
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
		SubscriptionID: emptyGUID,
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
		TenantID: emptyGUID,
	}

	err := tester.updateFieldsForTest(sample)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(sample.TenantID).To(Equal(tester.azureTenant))
}

func Test_SamplesTester_UpdatesResourceGroupName(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		rgName: "somerandomrg",
	}

	sample := &sampleResource{
		Reference: genruntime.ResourceReference{
			Group: "resources.azure.com",
			Kind:  "ResourceGroup",
			Name:  "aso-sample-rg",
		},
	}

	err := tester.updateFieldsForTest(sample)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(sample.Reference.Name).To(Equal(tester.rgName))
}

func Test_SamplesTester_HandlesDirectARMReferenceOwner(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		useRandomName: true,
		rgName:        "test-rg",
	}

	// Create a mock resource with a direct ARM reference owner (like the Quota sample)
	mockOwner := &genruntime.ArbitraryOwnerReference{
		ARMID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-sample-rg",
		// Note: Kind and Group are empty, simulating the Quota sample scenario
	}

	mockResource := &mockResourceWithOwner{
		owner: mockOwner,
	}

	samples := map[string]client.Object{
		"TestResource": mockResource,
	}
	refs := map[string]client.Object{}

	// This should not return an error anymore with our fix
	err := tester.setSamplesOwnershipAndReferences(samples, refs)
	g.Expect(err).ToNot(HaveOccurred())
}

func Test_SamplesTester_FailsForMissingKubernetesOwner(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	tester := &SamplesTester{
		useRandomName: true,
		rgName:        "test-rg",
	}

	// Create a mock resource with a Kubernetes reference owner that doesn't exist
	mockOwner := &genruntime.ArbitraryOwnerReference{
		Kind:  "NonExistentResource",
		Group: "test.azure.com",
		Name:  "some-name",
		// No ARMID - this is a Kubernetes reference
	}

	mockResource := &mockResourceWithOwner{
		owner: mockOwner,
	}

	samples := map[string]client.Object{
		"TestResource": mockResource,
	}
	refs := map[string]client.Object{}

	// This should still return an error for missing Kubernetes resources
	err := tester.setSamplesOwnershipAndReferences(samples, refs)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("owner: NonExistentResource, does not exist for resource 'TestResource'"))
}

// mockResourceWithOwner is a test helper that implements genruntime.ARMMetaObject
// and has an owner field for testing owner resolution logic
type mockResourceWithOwner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	owner             *genruntime.ArbitraryOwnerReference
}

var _ genruntime.ARMMetaObject = &mockResourceWithOwner{}
var _ client.Object = &mockResourceWithOwner{}

func (m *mockResourceWithOwner) Owner() *genruntime.ResourceReference {
	if m.owner == nil {
		return nil
	}
	return m.owner.AsResourceReference()
}

func (m *mockResourceWithOwner) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{}
}

func (m *mockResourceWithOwner) GetConditions() conditions.Conditions {
	return conditions.Conditions{}
}

func (m *mockResourceWithOwner) SetConditions(conditions conditions.Conditions) {
	// No-op for testing
}

func (m *mockResourceWithOwner) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

func (m *mockResourceWithOwner) GetSpec() genruntime.ConvertibleSpec {
	return nil
}

func (m *mockResourceWithOwner) GetStatus() genruntime.ConvertibleStatus {
	return nil
}

func (m *mockResourceWithOwner) SetStatus(status genruntime.ConvertibleStatus) error {
	return nil
}

func (m *mockResourceWithOwner) GetType() string {
	return "Microsoft.Test/mockResources"
}

func (m *mockResourceWithOwner) NewEmptyStatus() genruntime.ConvertibleStatus {
	return nil
}

func (m *mockResourceWithOwner) AzureName() string {
	return "mock-resource"
}

func (m *mockResourceWithOwner) GetAPIVersion() string {
	return "2023-01-01"
}

func (m *mockResourceWithOwner) DeepCopyObject() runtime.Object {
	return &mockResourceWithOwner{
		TypeMeta:   m.TypeMeta,
		ObjectMeta: m.ObjectMeta,
		owner:      m.owner,
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
