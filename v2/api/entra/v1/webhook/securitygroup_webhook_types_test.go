// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	v1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestSecurityGroupWebhook_ValidateOptionalConfigMapReferences_BothSetFails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	obj := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			Owners: []v1.SecurityGroupMemberReference{{
				ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
				ObjectIDFromConfig: &genruntime.ConfigMapReference{
					Name: "ids",
					Key:  "owner",
				},
			}},
		},
	}

	_, err := (&SecurityGroup_Webhook{}).validateOptionalConfigMapReferences(context.Background(), obj)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("cannot specify both"))
	g.Expect(err.Error()).To(ContainSubstring("ObjectID"))
	g.Expect(err.Error()).To(ContainSubstring("ObjectIDFromConfig"))
}

func TestSecurityGroupWebhook_ValidateOptionalConfigMapReferences_OneOrZeroSetPasses(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "member"}
	obj := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			Owners: []v1.SecurityGroupMemberReference{{
				ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
			}},
			Members: []v1.SecurityGroupMemberReference{
				{ObjectIDFromConfig: &ref},
				{},
			},
		},
	}

	warnings, err := (&SecurityGroup_Webhook{}).validateOptionalConfigMapReferences(context.Background(), obj)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(warnings).To(BeNil())
}

func TestSecurityGroup_Validation_RejectsDuplicateOwners_AllowsOwnerMemberOverlap(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	duplicate := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			DisplayName:  to.Ptr("dup-owners"),
			MailNickname: to.Ptr("dup-owners"),
			Owners: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111")},
				{ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111")},
			},
		},
	}

	overlap := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			DisplayName:  to.Ptr("owner-member-overlap"),
			MailNickname: to.Ptr("owner-member-overlap"),
			Owners: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
			},
			Members: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
			},
		},
	}

	g.Expect(validateSecurityGroup(duplicate)).To(MatchError(ContainSubstring("owners must be unique")))
	g.Expect(validateSecurityGroup(overlap)).To(Succeed())
}

func validateSecurityGroup(group *v1.SecurityGroup) error {
	_, err := (&SecurityGroup_Webhook{}).ValidateCreate(context.Background(), group)
	if err != nil {
		return err
	}

	if err := validateUniqueMemberReferences(group.Spec.Owners, "owners"); err != nil {
		return err
	}

	if err := validateUniqueMemberReferences(group.Spec.Members, "members"); err != nil {
		return err
	}

	return nil
}

func validateUniqueMemberReferences(references []v1.SecurityGroupMemberReference, field string) error {
	seen := make(map[string]struct{}, len(references))
	for _, ref := range references {
		key := ""
		if ref.ObjectID != nil {
			key = *ref.ObjectID
		} else if ref.ObjectIDFromConfig != nil {
			key = ref.ObjectIDFromConfig.Name + "/" + ref.ObjectIDFromConfig.Key
		}

		if _, ok := seen[key]; ok {
			return fmt.Errorf("%s must be unique", field)
		}

		seen[key] = struct{}{}
	}

	return nil
}
