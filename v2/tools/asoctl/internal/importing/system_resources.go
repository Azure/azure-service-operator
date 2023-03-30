/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

type skippedResource struct {
	resourceType string
	name         string
	because      string
}

// skippedResources is a list of resources that we don't want to import, usually because they're automatically
// created by the Azure platform, and they can't be (re)defined by the user.
// New additions to this list should be sorted by resource type, then by name.
var skippedResources = []skippedResource{
	{
		resourceType: "Microsoft.DBforPostgreSQL/flexibleServers/databases",
		name:         "azure_maintenance",
		because:      "access by users is not allowed",
	},
	{
		resourceType: "Microsoft.DBforPostgreSQL/flexibleServers/databases",
		name:         "azure_sys",
		because:      "built in databases cannot be managed by ARM",
	},
	{
		resourceType: "Microsoft.DBforPostgreSQL/flexibleServers/databases",
		name:         "postgres",
		because:      "built in databases cannot be managed by ARM",
	},
}

func IsSystemResource(armID *arm.ResourceID) (string, bool) {
	for _, skipped := range skippedResources {
		if armID.ResourceType.String() == skipped.resourceType && armID.Name == skipped.name {
			return skipped.because, true
		}
	}

	return "", false
}
