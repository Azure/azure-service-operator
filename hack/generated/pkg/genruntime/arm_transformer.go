/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"strings"
)

// ArmTransformer is a type which can be converted to/from an Arm object shape.
// Each CRD resource must implement these methods.
type ArmTransformer interface {
	// owningName is the "a/b/c" name for the owner. For example this would be "VNet1" when deploying subnet1
	// or myaccount when creating Batch pool1
	ConvertToArm(owningName string) (interface{}, error)
	PopulateFromArm(owner KnownResourceReference, input interface{}) error
}

// CreateArmResourceNameForDeployment creates a "fully qualified" resource name for use
// in an Azure template deployment.
// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/child-resource-name-type#outside-parent-resource
// for details on the format of the name field in ARM templates.
func CreateArmResourceNameForDeployment(owningName string, name string) string {
	result := owningName + "/" + name
	return result
}

// ExtractKubernetesResourceNameFromArmName extracts the Kubernetes resource name from an ARM name.
// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/child-resource-name-type#outside-parent-resource
// for details on the format of the name field in ARM templates.
func ExtractKubernetesResourceNameFromArmName(armName string) string {
	if len(armName) == 0 {
		return ""
	}

	// TODO: Possibly need to worry about preserving case here, although ARM should be already
	strs := strings.Split(armName, "/")
	return strs[len(strs)-1]
}
