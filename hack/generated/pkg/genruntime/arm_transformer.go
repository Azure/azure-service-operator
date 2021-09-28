/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"strings"
)

// ConvertToARMResolvedDetails contains resolved references and names for use in
// converting a Kubernetes type to an ARM type.
type ConvertToARMResolvedDetails struct {
	// Name is the fully qualified name of the resource in Azure ("a/b/c").
	Name string

	// Scope is the scope the resource is deployed at. This is nil for resources which are not
	// extension resources
	Scope *string

	// ResolvedReferences is a set of references which have been resolved to their ARM IDs.
	ResolvedReferences ResolvedReferences
}

type ToARMConverter interface {
	// ConvertToARM converts this to an ARM resource.
	ConvertToARM(resolved ConvertToARMResolvedDetails) (interface{}, error)
}

type FromARMConverter interface {
	CreateEmptyARMValue() ARMResourceStatus
	PopulateFromARM(owner KnownResourceReference, input interface{}) error
}

// TODO: Consider ArmSpecTransformer and ARMTransformer, so we don't have to pass owningName/name through all the calls
// ARMTransformer is a type which can be converted to/from an Arm object shape.
// Each CRD resource must implement these methods.
type ARMTransformer interface {
	ToARMConverter
	FromARMConverter
}

// ExtractKubernetesResourceNameFromARMName extracts the Kubernetes resource name from an ARM name.
// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/child-resource-name-type#outside-parent-resource
// for details on the format of the name field in ARM templates.
func ExtractKubernetesResourceNameFromARMName(armName string) string {
	if len(armName) == 0 {
		return ""
	}

	// TODO: Possibly need to worry about preserving case here, although ARM should be already
	strs := strings.Split(armName, "/")
	return strs[len(strs)-1]
}
