// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all appinsights

package controllers

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func TestAppInsightsController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	assert := assert.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	appInsightsName := GenerateTestResourceName("appinsights")

	// Create an instance of Azure AppInsights
	instance := &azurev1alpha1.AppInsights{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appInsightsName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AppInsightsSpec{
			Kind:            "web",
			Location:        rgLocation,
			ResourceGroup:   rgName,
			ApplicationType: "other",
		},
	}

	EnsureInstance(ctx, t, tc, instance)

	// Make sure the secret is created
	var keyName string
	if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		keyName = fmt.Sprintf("appinsights-%s-%s", instance.Spec.ResourceGroup, instance.Name)
	} else {
		keyName = instance.Name
	}

	secretKey := secrets.SecretKey{Name: keyName, Namespace: instance.Namespace, Kind: "appinsights"}

	// Secret is created after reconciliation is in provisioned state
	var secret map[string][]byte
	assert.Eventually(func() bool {
		var err error
		secret, err = tc.secretClient.Get(ctx, secretKey)
		return err == nil
	}, tc.timeoutFast, tc.retry, "should be able to get secret")

	assert.NotEmpty(string(secret["instrumentationKey"]))

	EnsureDelete(ctx, t, tc, instance)
}
