/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// WellknownResourceReference is a variation of ResourceReference that also permits the referenced resource to be
// identified by a well known name.
type WellknownResourceReference struct {
	ResourceReference `json:",inline"`

	WellknownName string `json:"wellknownName,omitempty"`
}

// Copy makes an independent copy of the WellknownResourceReference
func (ref *WellknownResourceReference) Copy() WellknownResourceReference {
	return *ref
}

// CreateResourceReferenceFromARMID creates a new ResourceReference from a string representing an ARM ID
func CreateWellknownResourceReferenceFromARMID(armID string) WellknownResourceReference {
	return WellknownResourceReference{
		ResourceReference: ResourceReference{
			ARMID: armID,
		},
	}
}
