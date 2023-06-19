// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20230101 "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	v1api20230101s "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaultsBackupInstanceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *BackupVaultsBackupInstanceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20230101.BackupVaultsBackupInstance{},
		&v1api20230101s.BackupVaultsBackupInstance{}}
}