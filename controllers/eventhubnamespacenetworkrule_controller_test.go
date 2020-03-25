// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all eventhubnamespacenetworkrule

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestEventHubNamespaceNetworkRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the Eventhub namespace object as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Create EventhubNamespace network rule for this namespace but with a non existent RG

	eventhubNamespaceNetRuleInstance := &azurev1alpha1.EventhubNamespaceNetworkRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ehns-netrule",
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceNetworkRule{
			Namespace:     eventhubNamespaceName,
			ResourceGroup: "gone",
			DefaultAction: "allow",
		},
	}

	// Check that we get the RG not found error
	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceNetRuleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, eventhubNamespaceNetRuleInstance)

	// Delete eventhubnamespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
}

func TestEventHubNamespaceNetworkRuleControllerNoNamespace(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create EventhubNamespace network rule for this namespace but with a non existent RG

	eventhubNamespaceNetRuleInstance := &azurev1alpha1.EventhubNamespaceNetworkRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ehns-netrule",
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceNetworkRule{
			Namespace:     eventhubNamespaceName,
			ResourceGroup: rgName,
			DefaultAction: "allow",
		},
	}

	// Check that we get the ParentNotfound error
	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceNetRuleInstance, errhelp.ParentNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, eventhubNamespaceNetRuleInstance)

}

func TestEventHubNamespaceNetworkRuleControllerBasicNamespace(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the Eventhub namespace object as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.EventhubNamespaceSku{
				Name:     "Basic",
				Tier:     "Basic",
				Capacity: 1,
			},
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Create EventhubNamespace network rule for this namespace and expect error

	eventhubNamespaceNetRuleInstance := &azurev1alpha1.EventhubNamespaceNetworkRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ehns-netrule",
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceNetworkRule{
			Namespace:     eventhubNamespaceName,
			ResourceGroup: rgName,
			DefaultAction: "allow",
		},
	}

	// Check that we get the  error
	EnsureInstanceWithResult(ctx, t, tc, eventhubNamespaceNetRuleInstance, errhelp.BadRequest, false)

	EnsureDelete(ctx, t, tc, eventhubNamespaceNetRuleInstance)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)

}

func TestEventHubNamespaceNetworkRuleControllerHappy(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the Eventhub namespace object as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.EventhubNamespaceSku{
				Name:     "Standard",
				Tier:     "Standard",
				Capacity: 1,
			},
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Create EventhubNamespace network rule for this namespace and expect success
	eventhubNamespaceNetRuleInstance := &azurev1alpha1.EventhubNamespaceNetworkRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ehns-netrule",
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceNetworkRule{
			Namespace:     eventhubNamespaceName,
			ResourceGroup: rgName,
			DefaultAction: "deny",
		},
	}

	// Check that we get success
	EnsureInstanceW(ctx, t, tc, eventhubNamespaceNetRuleInstance)

	//TODO: How do we check if the rule was actually added?

	// Delete network rule
	EnsureDelete(ctx, t, tc, eventhubNamespaceNetRuleInstance)

	// Delete the namespace
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
}
