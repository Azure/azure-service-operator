// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

// CreationMode specifies whether ASO will create or adopt the Entra resource.
// +kubebuilder:validation:Enum=AdoptOrCreate;AlwaysCreate;AdoptOnly
type CreationMode string

const (
	// AlwaysCreate means that ASO will always attempt to create the resource,
	// without first checking to see whether it already exists.
	AlwaysCreate CreationMode = "AlwaysCreate"

	// AdoptOrCreate means that ASO will try to adopt an existing resource if it exists,
	// and can be uniquely identified.
	// If multiple matches are found, the resource condition will show an error.
	// If it does not exist, ASO will create a new resource.
	AdoptOrCreate CreationMode = "AdoptOrCreate"

	// AdoptOnly requires an existing resource and never creates one.
	AdoptOnly CreationMode = "AdoptOnly"
)

// AllowsCreation checks if the creation mode allows ASO to create a new resource.
func (cm CreationMode) AllowsCreation() bool {
	return cm == AlwaysCreate || cm == AdoptOrCreate
}

// AllowsAdoption checks if the creation mode allows ASO to adopt an existing resource.
func (cm CreationMode) AllowsAdoption() bool {
	return cm == AdoptOrCreate || cm == AdoptOnly
}
