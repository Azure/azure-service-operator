/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"sort"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"k8s.io/klog/v2"
)

// checkForMissingStatusInformation returns a stage that will check if there
// are any resources with no status information
func checkForMissingStatusInformation() PipelineStage {
	return MakePipelineStage(
		"statusCheck",
		"Check for missing status information",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {

			var missingStatus []astmodel.TypeName

			for typeName, typeDef := range defs {
				if resourceType, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					if resourceType.StatusType() == nil {
						missingStatus = append(missingStatus, typeName)
					}
				}
			}

			sort.Slice(missingStatus, func(i, j int) bool {
				return astmodel.SortTypeName(missingStatus[i], missingStatus[j])
			})

			klog.Errorf("Missing status information for %v types", len(missingStatus))

			for _, typeName := range missingStatus {
				klog.V(2).Infof("No status information found for %v", typeName)
			}

			return defs, nil // TODO: return an error in future if status info is missing
		})
}
