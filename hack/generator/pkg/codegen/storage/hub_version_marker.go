/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// HubVersionMarker is a utility for marking resource types as "hub" versions
type HubVersionMarker struct {
	// visitor is used to do the actual marking
	visitor astmodel.TypeVisitor
}

// NewHubVersionMarker returns a new hub version marker for flagging resource types
func NewHubVersionMarker() *HubVersionMarker {
	result := &HubVersionMarker{}

	result.visitor = astmodel.TypeVisitorBuilder{
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
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, _ interface{}) (astmodel.Type, error) {
	return rt.MarkAsStorageVersion(), nil
}
