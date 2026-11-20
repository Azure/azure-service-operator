// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	v1 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1"
)

func localUser() *v1.User {
	return &v1.User{Spec: v1.UserSpec{AzureName: "app", LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}}
}

func aadUser(objectID string, principalType string) *v1.User {
	return &v1.User{Spec: v1.UserSpec{
		AzureName: "app-identity",
		AADUser:   &v1.AADUserSpec{ServerAdminUsername: "admin-group", ObjectId: objectID, PrincipalType: principalType},
	}}
}

func TestUserWebhook_ValidateCreate(t *testing.T) {
	t.Parallel()

	neither := &v1.User{Spec: v1.UserSpec{AzureName: "app"}}
	both := localUser()
	both.Spec.AADUser = &v1.AADUserSpec{ServerAdminUsername: "admin-group"}

	cases := map[string]struct {
		user    *v1.User
		wantErr string
	}{
		"local user":                       {user: localUser()},
		"AAD user by name":                 {user: aadUser("", "")},
		"AAD user by object ID":            {user: aadUser("00000000-0000-0000-0000-000000000001", "service")},
		"neither local nor AAD":            {user: neither, wantErr: "exactly one of spec.localUser or spec.aadUser must be set"},
		"both local and AAD":               {user: both, wantErr: "exactly one of spec.localUser or spec.aadUser must be set"},
		"object ID without principal type": {user: aadUser("00000000-0000-0000-0000-000000000001", ""), wantErr: "must be set together"},
		"principal type without object ID": {user: aadUser("", "group"), wantErr: "must be set together"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := (&User_Webhook{}).ValidateCreate(context.Background(), c.user)
			if c.wantErr == "" {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				g.Expect(err).To(MatchError(ContainSubstring(c.wantErr)))
			}
		})
	}
}

func TestUserWebhook_ValidateUpdate(t *testing.T) {
	t.Parallel()

	const oid = "00000000-0000-0000-0000-000000000001"

	cases := map[string]struct {
		old     *v1.User
		new     *v1.User
		wantErr string
	}{
		"AAD user unchanged":        {old: aadUser(oid, "service"), new: aadUser(oid, "service")},
		"local user to AAD user":    {old: localUser(), new: aadUser("", ""), wantErr: "cannot change from local user to AAD user"},
		"AAD user to local user":    {old: aadUser("", ""), new: localUser(), wantErr: "cannot change from AAD User to local user"},
		"object ID changed":         {old: aadUser(oid, "service"), new: aadUser("00000000-0000-0000-0000-000000000002", "service"), wantErr: "cannot change AAD user 'objectId'"},
		"object ID added":           {old: aadUser("", ""), new: aadUser(oid, "service"), wantErr: "cannot change AAD user 'objectId'"},
		"principal type changed":    {old: aadUser(oid, "service"), new: aadUser(oid, "group"), wantErr: "cannot change AAD user 'principalType'"},
		"new spec is neither":       {old: localUser(), new: &v1.User{Spec: v1.UserSpec{AzureName: "app"}}, wantErr: "exactly one of"},
		"new spec half a principal": {old: aadUser(oid, "service"), new: aadUser(oid, ""), wantErr: "must be set together"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := (&User_Webhook{}).ValidateUpdate(context.Background(), c.old, c.new)
			if c.wantErr == "" {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				g.Expect(err).To(MatchError(ContainSubstring(c.wantErr)))
			}
		})
	}
}
