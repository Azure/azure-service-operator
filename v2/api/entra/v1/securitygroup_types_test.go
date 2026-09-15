// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestSecurityGroupSpec_AssignODataBindOnCreate_UsesInlineObjectID(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{{
			ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		}},
		Members: []SecurityGroupMemberReference{{
			ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		}},
	}

	group := msgraphmodels.NewGroup()
	err := spec.AssignODataBindOnCreate(group, genruntime.MakeResolved[genruntime.ConfigMapReference, string](nil))
	g.Expect(err).ToNot(HaveOccurred())

	additional := group.GetAdditionalData()
	g.Expect(additional).To(HaveKey("owners@odata.bind"))
	g.Expect(additional).To(HaveKey("members@odata.bind"))
	g.Expect(additional["owners@odata.bind"]).To(Equal([]string{
		"https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111",
	}))
	g.Expect(additional["members@odata.bind"]).To(Equal([]string{
		"https://graph.microsoft.com/v1.0/directoryObjects/22222222-2222-2222-2222-222222222222",
	}))
}

func TestSecurityGroupSpec_AssignODataBindOnCreate_UsesObjectIDFromConfig(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ownerRef := genruntime.ConfigMapReference{Name: "ids", Key: "owner"}
	memberRef := genruntime.ConfigMapReference{Name: "ids", Key: "member"}
	resolved := genruntime.MakeResolved(map[genruntime.ConfigMapReference]string{
		ownerRef:  "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		memberRef: "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
	})

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{{
			ObjectIDFromConfig: &ownerRef,
		}},
		Members: []SecurityGroupMemberReference{{
			ObjectIDFromConfig: &memberRef,
		}},
	}

	group := msgraphmodels.NewGroup()
	err := spec.AssignODataBindOnCreate(group, resolved)
	g.Expect(err).ToNot(HaveOccurred())

	additional := group.GetAdditionalData()
	g.Expect(additional["owners@odata.bind"]).To(Equal([]string{
		"https://graph.microsoft.com/v1.0/directoryObjects/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}))
	g.Expect(additional["members@odata.bind"]).To(Equal([]string{
		"https://graph.microsoft.com/v1.0/directoryObjects/bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
	}))
}

func TestSecurityGroupSpec_AssignODataBindOnCreate_ErrorsWhenNeitherSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{{}},
	}

	group := msgraphmodels.NewGroup()
	err := spec.AssignODataBindOnCreate(group, genruntime.MakeResolved[genruntime.ConfigMapReference, string](nil))
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("owners[0]")))
}

func TestSecurityGroupSpec_AssignODataBindOnCreate_ErrorsWhenConfigLookupFails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "missing"}
	spec := &SecurityGroupSpec{
		Members: []SecurityGroupMemberReference{{
			ObjectIDFromConfig: &ref,
		}},
	}

	group := msgraphmodels.NewGroup()
	err := spec.AssignODataBindOnCreate(group, genruntime.MakeResolved[genruntime.ConfigMapReference, string](nil))
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("members[0]")))
	g.Expect(err).To(MatchError(ContainSubstring("objectIDFromConfig")))
}

func TestSecurityGroupSpec_ResolveOwnerObjectIDs_ErrorsOnDuplicateResolvedValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "owner"}
	resolved := genruntime.MakeResolved(map[genruntime.ConfigMapReference]string{ref: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"})

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{
			{ObjectID: to.Ptr("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA")},
			{ObjectIDFromConfig: &ref},
		},
	}

	_, err := spec.ResolveOwnerObjectIDs(resolved)
	g.Expect(err).To(MatchError(ContainSubstring("owners[1] resolves to the same object id as owners[0]")))
	g.Expect(err).To(MatchError(ContainSubstring("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")))
}

func TestSecurityGroupSpec_ResolveMemberObjectIDs_ErrorsOnDuplicateResolvedValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "member"}
	resolved := genruntime.MakeResolved(map[genruntime.ConfigMapReference]string{ref: "22222222-2222-2222-2222-222222222222"})

	spec := &SecurityGroupSpec{
		Members: []SecurityGroupMemberReference{
			{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
			{ObjectIDFromConfig: &ref},
		},
	}

	_, err := spec.ResolveMemberObjectIDs(resolved)
	g.Expect(err).To(MatchError(ContainSubstring("members[1] resolves to the same object id as members[0]")))
	g.Expect(err).To(MatchError(ContainSubstring("22222222-2222-2222-2222-222222222222")))
}

func TestSecurityGroupSpec_ResolveOwnerObjectIDs_PreservesInputOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{
			{ObjectID: to.Ptr("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA")},
			{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
		},
	}

	ids, err := spec.ResolveOwnerObjectIDs(genruntime.MakeResolved[genruntime.ConfigMapReference, string](nil))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ids).To(Equal([]uuid.UUID{
		uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	}))
}

func TestSecurityGroupSpec_HasDynamicMembership(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		membershipType *SecurityGroupMembershipType
		expected       bool
	}{
		"default assigned":       {expected: false},
		"assigned":               {membershipType: to.Ptr(SecurityGroupMembershipTypeAssigned), expected: false},
		"assigned Microsoft 365": {membershipType: to.Ptr(SecurityGroupMembershipTypeAssignedM365), expected: false},
		"dynamic":                {membershipType: to.Ptr(SecurityGroupMembershipTypeDynamic), expected: true},
		"dynamic Microsoft 365":  {membershipType: to.Ptr(SecurityGroupMembershipTypeDynamicM365), expected: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			spec := SecurityGroupSpec{MembershipType: test.membershipType}

			g.Expect(spec.HasDynamicMembership()).To(Equal(test.expected))
		})
	}
}
