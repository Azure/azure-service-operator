// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuredisk

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDiskControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := GenerateTestResourceNameWithRandom("rg", 10)
	diskName := GenerateTestResourceNameWithRandom("vm", 10)
	diskCreateOption := "Empty"
	var diskSize int32 = 10

	// Create a VM
	diskInstance := &azurev1alpha1.AzureDisk{
		ObjectMeta: metav1.ObjectMeta{
			Name:      diskName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDiskSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: rgName,
			CreateOption:  diskCreateOption,
			DiskSizeGB:    diskSize,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, diskInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, diskInstance)
}

func TestDiskControllerUnsupportedCreateOption(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := GenerateTestResourceNameWithRandom("rg", 10)
	diskName := GenerateTestResourceNameWithRandom("vm", 10)
	diskCreateOption := "FromImage" // Using option that is not supported.
	var diskSize int32 = 10

	// Create a VM
	diskInstance := &azurev1alpha1.AzureDisk{
		ObjectMeta: metav1.ObjectMeta{
			Name:      diskName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDiskSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: rgName,
			CreateOption:  diskCreateOption,
			DiskSizeGB:    diskSize,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, diskInstance, errhelp.UnsupportedDiskCreateOption, false)

	EnsureDelete(ctx, t, tc, diskInstance)
}

func TestDiskHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	diskCreateOption := "Empty"
	var diskSize int32 = 10

	// Create a Disk
	diskInstance := &azurev1alpha1.AzureDisk{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-disk",
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDiskSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			CreateOption:  diskCreateOption,
			DiskSizeGB:    diskSize,
		},
	}

	EnsureInstance(ctx, t, tc, diskInstance)

	EnsureDelete(ctx, t, tc, diskInstance)
}
