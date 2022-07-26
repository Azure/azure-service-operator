/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

// LocatableResource represents a resource with a location.
// TODO: The generated types should impl this. Currently only ResourceGroup has this interface and it's written manually
type LocatableResource interface {
	Location() string
}

// TenantResource is a marker interface that indicates that the implementing resource is scoped at the tenant level.
// Examples are: Subscriptions, Management Groups
type TenantResource interface {
	// Tenant is a marker method that indicates the resource in question is a tenant resource
	Tenant()
}
