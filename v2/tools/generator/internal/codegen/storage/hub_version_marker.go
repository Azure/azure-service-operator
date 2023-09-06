/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

// HubVersionMarker is a utility for marking resource definitions as "hub" versions
type HubVersionMarker struct {
	// visitor is used to do the actual marking
	visitor astmodel.TypeVisitor[any]
}

// NewHubVersionMarker returns a new hub version marker for flagging resource definitions
func NewHubVersionMarker() *HubVersionMarker {
	result := &HubVersionMarker{}

	result.visitor = astmodel.TypeVisitorBuilder[any]{
		VisitResourceType: result.markResourceAsStorageVersion,
	}.Build()

	return result
}

// MarkAsStorageVersion marks the supplied type definition as the storage version
func (m *HubVersionMarker) MarkAsStorageVersion(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	return m.visitor.VisitDefinition(def, nil)
}

// markResourceAsStorageVersion marks the supplied resource as the canonical hub (storage) version
func (m *HubVersionMarker) markResourceAsStorageVersion(
	_ *astmodel.TypeVisitor[any], rt *astmodel.ResourceType, _ interface{}) (astmodel.Type, error) {
	return rt.MarkAsStorageVersion(), nil
}
