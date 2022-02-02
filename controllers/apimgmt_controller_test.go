// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || apimgmt
// +build all apimgmt

package controllers

import (
	"context"
	"os"
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
	rgName := "AzureOperatorsTest"
	apimServiceName := "AzureOperatorsTestAPIM"

	// Read the pre-created APIM service name and RG name from Environment variable if it exists
	rgFromEnv := os.Getenv("TEST_APIM_RG")
	if len(rgFromEnv) != 0 {
		rgName = rgFromEnv
	}
	nameFromEnv := os.Getenv("TEST_APIM_NAME")
	if len(nameFromEnv) != 0 {
		apimServiceName = nameFromEnv
	}

	// Create an instance of Azure APIMgmnt
	apiMgmtInstance := &azurev1alpha1.APIMgmtAPI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiMgmtName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.APIMgmtSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			APIService:    apimServiceName,
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
