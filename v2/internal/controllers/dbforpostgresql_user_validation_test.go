/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	postgresqlv1 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Test_DBForPostgreSQL_User_ValidationRules exercises the CEL rules on the PostgreSQL User CRD. Every request here is
// either rejected at admission or names an owner that doesn't exist, so nothing reaches Azure.
func Test_DBForPostgreSQL_User_ValidationRules(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	owner := &genruntime.KubernetesOwnerReference{Name: "nonexistent-server"}
	localUser := func() *postgresqlv1.LocalUserSpec {
		return &postgresqlv1.LocalUserSpec{
			ServerAdminUsername: "myadmin",
			ServerAdminPassword: &genruntime.SecretReference{Name: "secret", Key: "adminPassword"},
			Password:            &genruntime.SecretReference{Name: "secret", Key: "password"},
		}
	}
	newUser := func(prefix string, spec postgresqlv1.UserSpec) *postgresqlv1.User {
		spec.Owner = owner
		return &postgresqlv1.User{
			ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName(prefix)),
			Spec:       spec,
		}
	}

	// Creation

	noSpec := &unstructured.Unstructured{}
	noSpec.SetGroupVersionKind(postgresqlv1.GroupVersion.WithKind("User"))
	noSpec.SetName(tc.NoSpaceNamer.GenerateName("nospec"))
	noSpec.SetNamespace(tc.Namespace)
	err := tc.CreateResourceExpectRequestFailure(noSpec)
	tc.Expect(err).To(MatchError(ContainSubstring("spec: Required")))

	createCases := []struct {
		name     string
		spec     postgresqlv1.UserSpec
		expected string
	}{{
		name:     "neither authentication mode",
		spec:     postgresqlv1.UserSpec{},
		expected: "exactly one of spec.localUser or spec.aadUser must be set",
	}, {
		name: "both authentication modes",
		spec: postgresqlv1.UserSpec{
			LocalUser: localUser(),
			AADUser:   &postgresqlv1.AADUserSpec{ServerAdminUsername: "aadadmin"},
		},
		expected: "exactly one of spec.localUser or spec.aadUser must be set",
	}, {
		name: "objectId without principalType",
		spec: postgresqlv1.UserSpec{
			AADUser: &postgresqlv1.AADUserSpec{ServerAdminUsername: "aadadmin", ObjectId: "00000000-0000-0000-0000-000000000001"},
		},
		expected: "spec.aadUser.objectId and spec.aadUser.principalType must be set together",
	}, {
		name: "principalType without objectId",
		spec: postgresqlv1.UserSpec{
			AADUser: &postgresqlv1.AADUserSpec{ServerAdminUsername: "aadadmin", PrincipalType: "service"},
		},
		expected: "spec.aadUser.objectId and spec.aadUser.principalType must be set together",
	}}

	for _, c := range createCases {
		tc.LogSubsectionf("Create with %s is rejected", c.name)
		err = tc.CreateResourceExpectRequestFailure(newUser("user", c.spec))
		tc.Expect(err).To(MatchError(ContainSubstring(c.expected)))
	}

	local := newUser("local", postgresqlv1.UserSpec{LocalUser: localUser()})
	aadByName := newUser("aadname", postgresqlv1.UserSpec{
		AADUser: &postgresqlv1.AADUserSpec{ServerAdminUsername: "aadadmin"},
	})
	aadByObjectID := newUser("aadoid", postgresqlv1.UserSpec{
		AADUser: &postgresqlv1.AADUserSpec{
			ServerAdminUsername: "aadadmin",
			ObjectId:            "00000000-0000-0000-0000-000000000001",
			PrincipalType:       "service",
		},
	})
	tc.CreateResource(local)
	tc.CreateResource(aadByName)
	tc.CreateResource(aadByObjectID)

	// Updates

	updateCases := []struct {
		name     string
		user     *postgresqlv1.User
		mutate   func(user *postgresqlv1.User)
		expected string
	}{{
		name: "switching from local to AAD",
		user: local,
		mutate: func(u *postgresqlv1.User) {
			u.Spec.LocalUser = nil
			u.Spec.AADUser = &postgresqlv1.AADUserSpec{ServerAdminUsername: "aadadmin"}
		},
		expected: "cannot change from local user to AAD user",
	}, {
		name:     "switching from AAD to local",
		user:     aadByName,
		mutate:   func(u *postgresqlv1.User) { u.Spec.AADUser = nil; u.Spec.LocalUser = localUser() },
		expected: "cannot change from AAD user to local user",
	}, {
		name: "adding an object ID binding",
		user: aadByName,
		mutate: func(u *postgresqlv1.User) {
			u.Spec.AADUser.ObjectId = "00000000-0000-0000-0000-000000000001"
			u.Spec.AADUser.PrincipalType = "service"
		},
		expected: "cannot change AAD user objectId",
	}, {
		name:     "removing the object ID binding",
		user:     aadByObjectID,
		mutate:   func(u *postgresqlv1.User) { u.Spec.AADUser.ObjectId = ""; u.Spec.AADUser.PrincipalType = "" },
		expected: "cannot change AAD user objectId",
	}, {
		name:     "changing the object ID",
		user:     aadByObjectID,
		mutate:   func(u *postgresqlv1.User) { u.Spec.AADUser.ObjectId = "00000000-0000-0000-0000-000000000002" },
		expected: "cannot change AAD user objectId",
	}, {
		name:     "changing the principal type",
		user:     aadByObjectID,
		mutate:   func(u *postgresqlv1.User) { u.Spec.AADUser.PrincipalType = "user" },
		expected: "cannot change AAD user principalType",
	}}

	for _, c := range updateCases {
		tc.LogSubsectionf("Update %s is rejected", c.name)
		old := c.user.DeepCopy()
		updated := c.user.DeepCopy()
		c.mutate(updated)
		err = tc.PatchAndExpectError(old, updated)
		tc.Expect(err).To(MatchError(ContainSubstring(c.expected)))
	}

	tc.LogSubsectionf("Update that keeps the principal binding is accepted")
	old := aadByObjectID.DeepCopy()
	aadByObjectID.Spec.Roles = []string{"pg_read_all_stats"}
	aadByObjectID.Spec.AADUser.ServerAdminUsername = "otheradmin"
	tc.Patch(old, aadByObjectID)
}
