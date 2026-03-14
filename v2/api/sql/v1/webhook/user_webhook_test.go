// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_UserWebhook_ValidateIsLocalOrAAD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	webhook := &User_Webhook{}
	ctx := context.Background()

	tests := []struct {
		name    string
		user    *v1.User
		wantErr string
	}{
		{
			name: "valid local user",
			user: &v1.User{
				Spec: v1.UserSpec{
					LocalUser: &v1.LocalUserSpec{
						ServerAdminUsername: "admin",
						ServerAdminPassword: &genruntime.SecretReference{Name: "secret", Key: "password"},
						Password:            &genruntime.SecretReference{Name: "secret", Key: "userpass"},
					},
				},
			},
			wantErr: "",
		},
		{
			name: "valid AAD user",
			user: &v1.User{
				Spec: v1.UserSpec{
					AADUser: &v1.AADUserSpec{
						ServerAdminUsername: "admin",
					},
				},
			},
			wantErr: "",
		},
		{
			name:    "neither specified",
			user:    &v1.User{Spec: v1.UserSpec{}},
			wantErr: "exactly one of spec.localUser or spec.aadUser must be set",
		},
		{
			name: "both specified",
			user: &v1.User{
				Spec: v1.UserSpec{
					LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"},
					AADUser:   &v1.AADUserSpec{ServerAdminUsername: "admin"},
				},
			},
			wantErr: "exactly one of spec.localUser or spec.aadUser must be set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := webhook.validateIsLocalOrAAD(ctx, tt.user)
			if tt.wantErr == "" {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).ToNot(BeNil())
				g.Expect(err.Error()).To(ContainSubstring(tt.wantErr))
			}
		})
	}
}

func Test_UserWebhook_ValidateUserTypeNotChanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	webhook := &User_Webhook{}
	ctx := context.Background()

	tests := []struct {
		name    string
		oldUser *v1.User
		newUser *v1.User
		wantErr string
	}{
		{
			name:    "local to local - allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}},
			newUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin2"}}},
			wantErr: "",
		},
		{
			name:    "aad to aad - allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{ServerAdminUsername: "admin"}}},
			newUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{ServerAdminUsername: "admin2"}}},
			wantErr: "",
		},
		{
			name:    "local to aad - not allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}},
			newUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{ServerAdminUsername: "admin"}}},
			wantErr: "cannot change from local user to AAD user",
		},
		{
			name:    "aad to local - not allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{ServerAdminUsername: "admin"}}},
			newUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}},
			wantErr: "cannot change from AAD user to local user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := webhook.validateUserTypeNotChanged(ctx, tt.oldUser, tt.newUser)
			if tt.wantErr == "" {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).ToNot(BeNil())
				g.Expect(err.Error()).To(ContainSubstring(tt.wantErr))
			}
		})
	}
}

func Test_UserWebhook_ValidateUserAADAliasNotChanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	webhook := &User_Webhook{}
	ctx := context.Background()

	tests := []struct {
		name    string
		oldUser *v1.User
		newUser *v1.User
		wantErr string
	}{
		{
			name:    "alias unchanged - allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{Alias: "myalias"}}},
			newUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{Alias: "myalias"}}},
			wantErr: "",
		},
		{
			name:    "alias changed - not allowed",
			oldUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{Alias: "oldalias"}}},
			newUser: &v1.User{Spec: v1.UserSpec{AADUser: &v1.AADUserSpec{Alias: "newalias"}}},
			wantErr: "cannot change AAD user 'alias'",
		},
		{
			name:    "local user - skipped",
			oldUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}},
			newUser: &v1.User{Spec: v1.UserSpec{LocalUser: &v1.LocalUserSpec{ServerAdminUsername: "admin"}}},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := webhook.validateUserAADAliasNotChanged(ctx, tt.oldUser, tt.newUser)
			if tt.wantErr == "" {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).ToNot(BeNil())
				g.Expect(err.Error()).To(ContainSubstring(tt.wantErr))
			}
		})
	}
}

func Test_UserWebhook_ValidateCreate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	webhook := &User_Webhook{}
	ctx := context.Background()

	user := &v1.User{
		ObjectMeta: metav1.ObjectMeta{Name: "test-user", Namespace: "default"},
		Spec: v1.UserSpec{
			AzureName: "my-managed-identity",
			AADUser:   &v1.AADUserSpec{ServerAdminUsername: "admin"},
		},
	}

	_, err := webhook.ValidateCreate(ctx, user)
	g.Expect(err).To(BeNil())
}
