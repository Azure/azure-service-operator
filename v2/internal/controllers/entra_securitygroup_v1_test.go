/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Entra_SecurityGroup_v1_CRUD(t *testing.T) {
	t.Parallel()

	// g := NewGomegaWithT(t)
	// ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	securityGroup := &entra.SecurityGroup{
		ObjectMeta: tc.MakeObjectMeta("securitygroup"),
		Spec: entra.SecurityGroupSpec{
			DisplayName:    to.Ptr("ASO Test Security Group"),
			MailNickname:   to.Ptr("asotest-security-group"),
			MembershipType: to.Ptr(entra.SecurityGroupMembershipTypeAssigned),
		},
	}

	// Create the resource and wait for it to be ready
	tc.CreateResourceAndWait(securityGroup)

	// Save an update
	old := securityGroup.DeepCopy()
	securityGroup.Spec.DisplayName = to.Ptr("ASO Test Security Group Updated")
	tc.PatchResourceAndWait(old, securityGroup)

	// Delete the resource and wait for it to be gone
	tc.DeleteResourceAndWait(securityGroup)
}
