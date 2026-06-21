// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import (
	"testing"

	. "github.com/onsi/gomega"

	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"

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
	g.Expect(additional["owners@odata.bind"]).To(Equal([]string{"https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111"}))
	g.Expect(additional["members@odata.bind"]).To(Equal([]string{"https://graph.microsoft.com/v1.0/directoryObjects/22222222-2222-2222-2222-222222222222"}))
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
	g.Expect(additional["owners@odata.bind"]).To(Equal([]string{"https://graph.microsoft.com/v1.0/directoryObjects/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"}))
	g.Expect(additional["members@odata.bind"]).To(Equal([]string{"https://graph.microsoft.com/v1.0/directoryObjects/bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"}))
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

func TestSecurityGroupSpec_ResolveOwnerObjectIDs_DeduplicatesResolvedValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "owner"}
	resolved := genruntime.MakeResolved(map[genruntime.ConfigMapReference]string{ref: "11111111-1111-1111-1111-111111111111"})

	spec := &SecurityGroupSpec{
		Owners: []SecurityGroupMemberReference{
			{ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111")},
			{ObjectIDFromConfig: &ref},
		},
	}

	ids, err := spec.ResolveOwnerObjectIDs(resolved)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ids).To(Equal([]string{"11111111-1111-1111-1111-111111111111"}))
}

func TestSecurityGroupSpec_ResolveMemberObjectIDs_DeduplicatesResolvedValues(t *testing.T) {
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

	ids, err := spec.ResolveMemberObjectIDs(resolved)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(ids).To(Equal([]string{"22222222-2222-2222-2222-222222222222"}))
}
