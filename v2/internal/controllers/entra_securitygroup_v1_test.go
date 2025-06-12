/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
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
			OperatorSpec: &entra.SecurityGroupOperatorSpec{
				Secrets: &entra.SecurityGroupOperatorSecrets{
					EntraID: &genruntime.SecretDestination{
						Name: "entra-secret",
						Key:  "entra-id",
					},
				},
			},
		},
	}

	// Create the resource and wait for it to be ready
	tc.CreateResourceAndWait(securityGroup)

	// Make sure the secret was created
	secretList := &v1.SecretList{}
	tc.ListResources(secretList, client.InNamespace(tc.Namespace))
	tc.Expect(secretList.Items).To(HaveLen(1))

	// Save an update
	old := securityGroup.DeepCopy()
	securityGroup.Spec.DisplayName = to.Ptr("ASO Test Security Group Updated")
	tc.PatchResourceAndWait(old, securityGroup)

	// Delete the resource and wait for it to be gone
	tc.DeleteResourceAndWait(securityGroup)
}
