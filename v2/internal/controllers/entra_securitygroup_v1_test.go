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
			MailNickname:   to.Ptr("asotest-security-group"),
			MembershipType: to.Ptr(entra.SecurityGroupMembershipTypeAssigned),
		},
	}

	tc.CreateResourceAndWait(securityGroup)

	tc.DeleteResourceAndWait(securityGroup)
}
