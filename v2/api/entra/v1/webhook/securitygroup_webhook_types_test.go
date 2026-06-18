// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"context"
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
