/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

// ExportedProperties is a map of string property name to a chain of properties. For allowing the export of a property "foo" that writes
// the status.a.b.c property to a config map would be "foo" -> []{status, a, b, c}
type ExportedProperties = map[string][]*astmodel.PropertyDefinition

type ExportedTypeNameProperties struct {
	exported map[astmodel.TypeName]ExportedProperties
}

func NewExportedTypeNameProperties() *ExportedTypeNameProperties {
	return &ExportedTypeNameProperties{
		exported: make(map[astmodel.TypeName]ExportedProperties),
	}
}

func (e *ExportedTypeNameProperties) Add(name astmodel.TypeName, exported ExportedProperties) {
	if len(exported) == 0 {
		// Nothing to do
		return
	}

	e.exported[name] = exported
}

func (e *ExportedTypeNameProperties) Get(name astmodel.TypeName) (ExportedProperties, bool) {
	result, ok := e.exported[name]
	return result, ok
}

func (e *ExportedTypeNameProperties) Copy() *ExportedTypeNameProperties {
	if e == nil {
		return nil
	}

	result := NewExportedTypeNameProperties()

	for key, inner := range e.exported {
		copiedInner := make(ExportedProperties)
		for innerKey, innerValue := range inner {
			copiedInner[innerKey] = append([]*astmodel.PropertyDefinition{}, innerValue...)
		}
		result.Add(key, copiedInner)
	}

	return result
}
