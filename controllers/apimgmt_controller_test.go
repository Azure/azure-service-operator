// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all apimgmt

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAPIMgmtController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// rgName := tc.resourceGroupName
	rgLocation := "southcentralus"
	apiMgmtName := "t-apimgmt-test" + helpers.RandomString(10)

	// Create an instance of Azure APIMgmnt
	apiMgmtInstance := &azurev1alpha1.APIMgmtAPI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiMgmtName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.APIMgmtSpec{
			Location:      rgLocation,
			ResourceGroup: "resourcegroup-azure-operators",
			APIService:    "aso-test-apimgmt",
			APIId:         "apiId0",
			Properties: azurev1alpha1.APIProperties{
				IsCurrent:             true,
				IsOnline:              true,
				DisplayName:           apiMgmtName,
				Description:           "API description",
				APIVersionDescription: "API version description",
				Path:                  "/api/test",
				Protocols:             []string{"http"},
				SubscriptionRequired:  false,
			},
		},
	}
	EnsureInstance(ctx, t, tc, apiMgmtInstance)

	EnsureDelete(ctx, t, tc, apiMgmtInstance)
}
