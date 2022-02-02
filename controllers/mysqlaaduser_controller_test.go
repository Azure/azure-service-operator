// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysqluser
// +build all mysqluser

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

func TestMySQLAADUserWebhook(t *testing.T) {
	// The webhook prevents a user from being created with ALL in roles.
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)

	user := v1alpha2.MySQLAADUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      GenerateTestResourceNameWithRandom("mysqlaaduser", 5),
			Namespace: "default",
		},
		Spec: v1alpha2.MySQLAADUserSpec{
			Server:        "does-not-matter",
			Roles:         []string{"PROCESS", "ALL"},
			ResourceGroup: "also-does-not-matter",
		},
	}
	err := tc.k8sClient.Create(ctx, &user)
	require.NotNil(err)
	require.Contains(err.Error(), "ASO admin user doesn't have privileges to grant ALL at server level")
}
