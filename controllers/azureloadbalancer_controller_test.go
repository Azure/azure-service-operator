// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azureloadbalancer
// +build all azureloadbalancer

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLoadBalancerNonExistingSubResources(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	lbName := GenerateTestResourceNameWithRandom("lb", 10)
	pipName := GenerateTestResourceNameWithRandom("pip", 10)
	bpName := "test"
	natName := "test"
	portRangeStart := 100
	portRangeEnd := 200
	backendPort := 300

	// Create an LB instance
	lbInstance := &azurev1alpha1.AzureLoadBalancer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      lbName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureLoadBalancerSpec{
			Location:               tc.resourceGroupLocation,
			ResourceGroup:          tc.resourceGroupName,
			PublicIPAddressName:    pipName,
			BackendAddressPoolName: bpName,
			InboundNatPoolName:     natName,
			FrontendPortRangeStart: portRangeStart,
			FrontendPortRangeEnd:   portRangeEnd,
			BackendPort:            backendPort,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, lbInstance, errhelp.InvalidResourceReference, false)

	EnsureDelete(ctx, t, tc, lbInstance)
}

func TestLaodBalancerHappyPathWithPublicIPAddress(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Create a Public IP Address
	pipName := GenerateTestResourceNameWithRandom("lb-pip1", 10)
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 tc.resourceGroupLocation,
			ResourceGroup:            tc.resourceGroupName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	// Create a Load Balancer
	lbName := GenerateTestResourceNameWithRandom("lb1", 10)
	bpName := "test"
	natName := "test"
	portRangeStart := 100
	portRangeEnd := 200
	backendPort := 300
	lbInstance := &azurev1alpha1.AzureLoadBalancer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      lbName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureLoadBalancerSpec{
			Location:               tc.resourceGroupLocation,
			ResourceGroup:          tc.resourceGroupName,
			PublicIPAddressName:    pipName,
			BackendAddressPoolName: bpName,
			InboundNatPoolName:     natName,
			FrontendPortRangeStart: portRangeStart,
			FrontendPortRangeEnd:   portRangeEnd,
			BackendPort:            backendPort,
		},
	}

	EnsureInstance(ctx, t, tc, lbInstance)

	EnsureDelete(ctx, t, tc, lbInstance)

	EnsureDelete(ctx, t, tc, pipInstance)
}

func TestLoadBalancerControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("lb-rand-rg", 10)

	// Create a Public IP Address
	pipName := GenerateTestResourceNameWithRandom("lb-pip2", 10)
	publicIPAllocationMethod := "Static"
	idleTimeoutInMinutes := 10
	publicIPAddressVersion := "IPv4"
	skuName := "Basic"
	pipInstance := &azurev1alpha1.AzurePublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pipName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzurePublicIPAddressSpec{
			Location:                 tc.resourceGroupLocation,
			ResourceGroup:            tc.resourceGroupName,
			PublicIPAllocationMethod: publicIPAllocationMethod,
			IdleTimeoutInMinutes:     idleTimeoutInMinutes,
			PublicIPAddressVersion:   publicIPAddressVersion,
			SkuName:                  skuName,
		},
	}

	EnsureInstance(ctx, t, tc, pipInstance)

	// Create a Load Balancer
	lbName := GenerateTestResourceNameWithRandom("lb2", 10)
	bpName := "test"
	natName := "test"
	portRangeStart := 100
	portRangeEnd := 200
	backendPort := 300
	lbInstance := &azurev1alpha1.AzureLoadBalancer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      lbName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureLoadBalancerSpec{
			Location:               tc.resourceGroupLocation,
			ResourceGroup:          rgName,
			PublicIPAddressName:    pipName,
			BackendAddressPoolName: bpName,
			InboundNatPoolName:     natName,
			FrontendPortRangeStart: portRangeStart,
			FrontendPortRangeEnd:   portRangeEnd,
			BackendPort:            backendPort,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, lbInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, lbInstance)

	EnsureDelete(ctx, t, tc, pipInstance)
}
