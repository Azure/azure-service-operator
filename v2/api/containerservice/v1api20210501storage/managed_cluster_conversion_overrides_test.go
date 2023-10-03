// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1api20210501storage_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501storage"
	v20230201s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201storage"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ManagedClusterSpec_RenamedPropertiesMapCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := v20210501s.ManagedCluster_Spec{
		DiskEncryptionSetIDReference: &genruntime.ResourceReference{
			ARMID: "Foo",
		},
	}

	nextSpec := v20230201s.ManagedCluster_Spec{}
	g.Expect(spec.AssignProperties_To_ManagedCluster_Spec(&nextSpec)).To(Succeed())

	g.Expect(nextSpec.DiskEncryptionSetReference).ToNot(BeNil())
	g.Expect(nextSpec.DiskEncryptionSetReference.ARMID).To(Equal("Foo"))
}

func Test_ManagedClusterAgentPoolProfile_RenamedPropertiesMapCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	profile := v20210501s.ManagedClusterAgentPoolProfile{
		VnetSubnetIDReference: &genruntime.ResourceReference{
			ARMID: "Foo",
		},
		PodSubnetIDReference: &genruntime.ResourceReference{
			ARMID: "Bar",
		},
		NodePublicIPPrefixIDReference: &genruntime.ResourceReference{
			ARMID: "Baz",
		},
		ProximityPlacementGroupID: to.Ptr("Qux"),
	}

	nextProfile := v20230201s.ManagedClusterAgentPoolProfile{}
	g.Expect(profile.AssignProperties_To_ManagedClusterAgentPoolProfile(&nextProfile)).To(Succeed())

	g.Expect(nextProfile.VnetSubnetReference).ToNot(BeNil())
	g.Expect(nextProfile.VnetSubnetReference.ARMID).To(Equal("Foo"))

	g.Expect(nextProfile.PodSubnetReference).ToNot(BeNil())
	g.Expect(nextProfile.PodSubnetReference.ARMID).To(Equal("Bar"))

	g.Expect(nextProfile.NodePublicIPPrefixReference).ToNot(BeNil())
	g.Expect(nextProfile.NodePublicIPPrefixReference.ARMID).To(Equal("Baz"))

	g.Expect(nextProfile.ProximityPlacementGroupReference).ToNot(BeNil())
	g.Expect(nextProfile.ProximityPlacementGroupReference.ARMID).To(Equal("Qux"))
}

func Test_ManagedClusterSKU_PropertiesMapCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	basic := string(v20210501.ManagedClusterSKU_Name_Basic)
	paid := string(v20210501.ManagedClusterSKU_Tier_Paid)
	sku := v20210501s.ManagedClusterSKU{
		Name: &basic,
		Tier: &paid,
	}

	nextSKU := v20230201s.ManagedClusterSKU{}
	g.Expect(sku.AssignProperties_To_ManagedClusterSKU(&nextSKU)).To(Succeed())

	g.Expect(nextSKU.Name).ToNot(BeNil())
	g.Expect(*nextSKU.Name).To(Equal("Base"))
	g.Expect(nextSKU.Tier).ToNot(BeNil())
	g.Expect(*nextSKU.Tier).To(Equal("Standard"))
}

func Test_ManagedClusterSKU_STATUS_PropertiesMapCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	basic := string(v20210501.ManagedClusterSKU_Name_Basic)
	paid := string(v20210501.ManagedClusterSKU_Tier_Paid)
	sku := v20210501s.ManagedClusterSKU_STATUS{
		Name: &basic,
		Tier: &paid,
	}

	nextSKU := v20230201s.ManagedClusterSKU_STATUS{}
	g.Expect(sku.AssignProperties_To_ManagedClusterSKU_STATUS(&nextSKU)).To(Succeed())

	g.Expect(nextSKU.Name).ToNot(BeNil())
	g.Expect(*nextSKU.Name).To(Equal("Base"))
	g.Expect(nextSKU.Tier).ToNot(BeNil())
	g.Expect(*nextSKU.Tier).To(Equal("Standard"))
}
