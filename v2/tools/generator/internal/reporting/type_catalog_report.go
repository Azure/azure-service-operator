/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/pkg/errors"
	"io"
	"os"
	"sort"
)

type TypeCatalogReport struct {
	title                  string
	defs                   astmodel.TypeDefinitionSet
	inlinedTypes           astmodel.TypeNameSet // Set of types that we inline when generating the report
	optionIncludeFunctions bool
}

func NewTypeCatalogReport(defs astmodel.TypeDefinitionSet) *TypeCatalogReport {
	return &TypeCatalogReport{
		defs:         defs,
		inlinedTypes: astmodel.NewTypeNameSet(),
	}
}

// SaveTo writes the report to the specified file
func (tcr *TypeCatalogReport) SaveTo(filePath string) error {

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken
		// state, so remove it
		if r := recover(); r != nil {
			os.Remove(filePath)
			panic(r)
		}
	}()

	err = tcr.WriteTo(file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(filePath)
	}

	return err
}

// IncludeFunctions specifies that the generated report should include functions
func (tcr *TypeCatalogReport) IncludeFunctions() {
	tcr.optionIncludeFunctions = true
}

// InlineTypes specifies that the generated report should inline types where referenced,
// We achieve this by scanning for properties with types we have definitions for
func (tcr *TypeCatalogReport) InlineTypes() {
	for _, def := range tcr.defs {
		if c, ok := astmodel.AsPropertyContainer(def.Type()); ok {
			tcr.inlineTypesFrom(c)
		}
	}
}

// inlineTypesFrom ensures we will inline any candidate types referenced by the property container
func (tcr *TypeCatalogReport) inlineTypesFrom(container astmodel.PropertyContainer) {
	emptySet := astmodel.NewTypeNameSet()
	for _, prop := range container.Properties().AsSlice() {
		// Check to see if this property references a definition that can be inlined
		if def, ok := tcr.asDefinitionToInline(prop.PropertyType(), emptySet); ok {
			tcr.inlinedTypes.Add(def.Name())
		}
	}
}

func (tcr *TypeCatalogReport) WriteTo(writer io.Writer) error {
	packages := tcr.findPackages()
	for _, pkg := range packages {
		rpt := NewStructureReport(pkg.PackagePath())
		tcr.writeDefinitions(rpt, tcr.inPackage(pkg))

		err := rpt.SaveTo(writer)
		if err != nil {
			return errors.Wrapf(err, "failed to create type catalog report for %s", pkg.PackagePath())
		}
	}

	return nil
}

// writeDefinitions writes the definitions to the type catalog report.
// rpt is the debug report to write to.
// definitions is the set of definitions to write.
// Definitions are written in alphabetical order, by case-sensitive sort
func (tcr *TypeCatalogReport) writeDefinitions(
	rpt *StructureReport,
	definitions astmodel.TypeDefinitionSet) {
	defs := definitions.AsSlice()
	sort.Slice(defs, func(i, j int) bool {
		return defs[i].Name().Name() < defs[j].Name().Name()
	})

	for _, d := range defs {
		if !tcr.inlinedTypes.Contains(d.Name()) {
			tcr.writeDefinition(rpt, d)
		}
	}
}

// writeDefinition writes the definition to the debug report.
// rpt is the debug report to write to.
// definition is the definition to write.
func (tcr *TypeCatalogReport) writeDefinition(
	rpt *StructureReport,
	definition astmodel.TypeDefinition,
) {
	name := definition.Name()
	parentTypes := astmodel.NewTypeNameSet(name)
	sub := rpt.Addf("%s: %s", name.Name(), tcr.asShortName(definition))
	tcr.writeType(sub, definition.Type(), name.PackageReference, parentTypes)
}

// writeType writes the type to the debug report.
// rpt is the debug report to write to.
// t is the type to write.
// currentPackage is the package that the type is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
// Only complex types where astmodel.DebugDescription is insufficient are written.
func (tcr *TypeCatalogReport) writeType(
	rpt *StructureReport,
	t astmodel.Type,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	if rsrc, ok := astmodel.AsResourceType(t); ok {
		tcr.writeResource(rpt, rsrc, currentPackage, parentTypes)
	}

	if obj, ok := astmodel.AsObjectType(t); ok {
		tcr.writeObject(rpt, obj, currentPackage, parentTypes)
	}

	if obj, ok := astmodel.AsEnumType(t); ok {
		tcr.writeEnum(rpt, obj, currentPackage, parentTypes)
	}

	if one, ok := astmodel.AsOneOfType(t); ok {
		tcr.writeOneOf(rpt, one, currentPackage, parentTypes)
	}

	if all, ok := astmodel.AsAllOfType(t); ok {
		tcr.writeAllOf(rpt, all, currentPackage, parentTypes)
	}
}

// writeResource writes the resource to the debug report.
// rpt is the debug report to write to.
// name is the name of the resource.
// resource is the resource to write.
// currentPackage is the package that the resource is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeResource(
	rpt *StructureReport,
	resource *astmodel.ResourceType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	for _, prop := range resource.Properties().AsSlice() {
		tcr.writeProperty(rpt, prop, currentPackage, parentTypes)
	}

	if tcr.optionIncludeFunctions {
		for _, fn := range resource.Functions() {
			tcr.writeFunction(rpt, fn)
		}
	}
}

// writeObject writes the object to the debug report..
// rpt is the debug report to write to.
// obj is the object to write.
// currentPackage is the package that the object is defined in (used to simplify type descriptions).
// parentTypes is the set of types that have already been written (used to avoid infinite recursion).
func (tcr *TypeCatalogReport) writeObject(
	rpt *StructureReport,
	obj *astmodel.ObjectType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	for _, prop := range obj.Properties().AsSlice() {
		tcr.writeProperty(rpt, prop, currentPackage, parentTypes)
	}

	if tcr.optionIncludeFunctions {
		for _, fn := range obj.Functions() {
			tcr.writeFunction(rpt, fn)
		}
	}
}

// writeProperty writes an individual property to the debug report, potentially inlining it's type
// rpt is the (sub)report we're writing to.
// prop is the property to write.
// currentPackage is the package that the property is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are parents of the property (used to detect cycles).
func (tcr *TypeCatalogReport) writeProperty(
	rpt *StructureReport,
	prop *astmodel.PropertyDefinition,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	if def, ok := tcr.asDefinitionToInline(prop.PropertyType(), parentTypes); ok && tcr.inlinedTypes.Contains(def.Name()) {
		// When inlining the type, we use a shortname to avoid the type name being repeated
		pt := parentTypes.Copy()
		pt.Add(def.Name())
		sub := rpt.Addf("%s: %s", prop.PropertyName(), tcr.asShortName(*def))
		tcr.writeType(sub, def.Type(), currentPackage, pt)
		return
	}

	// Base case, give the name of the property and a description of the type
	rpt.Addf(
		"%s: %s",
		prop.PropertyName(),
		astmodel.DebugDescription(prop.PropertyType(), currentPackage))
}

// asDefinitionToInline returns the definition to inline, if any.
// t is the type we're considering inlining.
// parentTypes is a set of all the types we're already inlining (to avoid infinite recursion).
func (tcr *TypeCatalogReport) asDefinitionToInline(
	t astmodel.Type,
	parentTypes astmodel.TypeNameSet,
) (*astmodel.TypeDefinition, bool) {

	// We can inline a typename if we have a definition for it, and if it's not already inlined
	if n, ok := astmodel.AsTypeName(t); ok {
		if parentTypes.Contains(n) {
			return nil, false
		}

		if def, ok := tcr.defs[n]; ok {
			return &def, true
		}
	}

	if m, ok := astmodel.AsMapType(t); ok {
		// We can inline the value of a map if we have a definition for it
		def, ok := tcr.asDefinitionToInline(m.ValueType(), parentTypes)
		return def, ok
	}

	if a, ok := astmodel.AsArrayType(t); ok {
		// We can inline the element of an array if we have a definition for it
		def, ok := tcr.asDefinitionToInline(a.Element(), parentTypes)
		return def, ok
	}

	return nil, false
}

// asShortName returns a short name for the type, for use when it's inlined
func (tcr *TypeCatalogReport) asShortName(def astmodel.TypeDefinition) string {
	if _, r := astmodel.AsResourceType(def.Type()); r {
		return "Resource"
	}

	if _, ok := astmodel.AsEnumType(def.Type()); ok {
		return "enum"
	}

	if m, ok := astmodel.AsMapType(def.Type()); ok {
		return fmt.Sprintf(
			"map[%s]",
			astmodel.DebugDescription(m.KeyType(), def.Name().PackageReference))
	}

	if a, ok := astmodel.AsArrayType(def.Type()); ok {
		return fmt.Sprintf(
			"%s[]",
			astmodel.DebugDescription(a.Element(), def.Name().PackageReference))
	}

	if one, ok := astmodel.AsOneOfType(def.Type()); ok {
		return fmt.Sprintf(
			"oneOf[%d]",
			one.Types().Len())
	}

	if all, ok := astmodel.AsAllOfType(def.Type()); ok {
		return fmt.Sprintf(
			"allOf[%d]",
			all.Types().Len())
	}

	return ""
}

func (tcr *TypeCatalogReport) writeFunction(
	rpt *StructureReport,
	fn astmodel.Function,
) {
	rpt.Addf("%s()", fn.Name())
}

// writeEnum writes an enum to the report
// rpt is the report to write to.
// enum is the enum to write.
// currentPackage is the package that the enum is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeEnum(
	rpt *StructureReport,
	enum *astmodel.EnumType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	tcr.writeType(rpt, enum.BaseType(), currentPackage, parentTypes)
	for _, v := range enum.Options() {
		rpt.Addf("%s", v.Value)
	}
}

// writeOneOf writes a oneof to the report.
// rpt is the report to write to.
// oneOf is the oneof to write.
// currentPackage is the package that the oneof is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeOneOf(
	rpt *StructureReport,
	oneOfType *astmodel.OneOfType,
	currentPackage astmodel.PackageReference,
	types astmodel.TypeNameSet,
) {
	oneOfType.Types().ForEach(func(t astmodel.Type, _ int) {
		sub := rpt.Addf("%s", astmodel.DebugDescription(t, currentPackage))
		tcr.writeType(sub, t, currentPackage, types)
	})
}

// writeAllOf writes an allof to the report.
// rpt is the report to write to.
// allOfType is the allof to write.
// currentPackage is the package that the allof is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeAllOf(
	rpt *StructureReport,
	allOfType *astmodel.AllOfType,
	currentPackage astmodel.PackageReference,
	types astmodel.TypeNameSet,
) {
	allOfType.Types().ForEach(func(t astmodel.Type, _ int) {
		sub := rpt.Addf("%s", astmodel.DebugDescription(t, currentPackage))
		tcr.writeType(sub, t, currentPackage, types)
	})
}

func (tcr *TypeCatalogReport) findPackages() []astmodel.PackageReference {
	packages := astmodel.NewPackageReferenceSet()
	for _, def := range tcr.defs {
		packages.AddReference(def.Name().PackageReference)
	}

	result := packages.AsSortedSlice(
		func(left astmodel.PackageReference, right astmodel.PackageReference) bool {
			return astmodel.ComparePathAndVersion(left.PackagePath(), right.PackagePath())
		})

	return result
}

func (tcr *TypeCatalogReport) inPackage(
	ref astmodel.PackageReference,
) astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)
	for _, def := range tcr.defs {
		if def.Name().PackageReference == ref {
			result.Add(def)
		}
	}

	return result
}
