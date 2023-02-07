/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
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
	sub := rpt.Addf("%s: %s", name.Name(), tcr.asShortNameForType(definition.Type(), name.PackageReference, parentTypes))
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
	// Generate a subreport for each kind of type
	// We switch on exact types because we don't want to accidentally unwrap a detail we need
	switch t := t.(type) {
	case *astmodel.ObjectType:
		tcr.writeObjectType(rpt, t, currentPackage, parentTypes)
	case *astmodel.InterfaceType:
		tcr.writeInterfaceType(rpt, t, currentPackage, parentTypes)
	case *astmodel.ResourceType:
		tcr.writeResourceType(rpt, t, currentPackage, parentTypes)
	case *astmodel.EnumType:
		tcr.writeEnumType(rpt, t, currentPackage, parentTypes)
	case *astmodel.OptionalType:
		tcr.writeType(rpt, t.Element(), currentPackage, parentTypes)
	case *astmodel.OneOfType:
		tcr.writeOneOfType(rpt, t, currentPackage, parentTypes)
	case *astmodel.AllOfType:
		tcr.writeAllOfType(rpt, t, currentPackage, parentTypes)
	case *astmodel.ErroredType:
		tcr.writeErroredType(rpt, t, currentPackage, parentTypes)
	case *astmodel.ValidatedType:
		tcr.writeValidatedType(rpt, t, currentPackage, parentTypes)
	case astmodel.MetaType:
		tcr.writeType(rpt, t.Unwrap(), currentPackage, parentTypes)
	default:
		// We don't need to write anything for simple types
	}
}

// writeResource writes the resource to the debug report.
// rpt is the debug report to write to.
// name is the name of the resource.
// resource is the resource to write.
// currentPackage is the package that the resource is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeResourceType(
	rpt *StructureReport,
	resource *astmodel.ResourceType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	// Write the expected owner of the resource, if we have one
	if owner := resource.Owner(); owner != nil {
		// We don't use asShortNameForType here because we don't want to inline the owner
		rpt.Addf("Owner: %s", astmodel.DebugDescription(owner, currentPackage))
	}

	for _, prop := range resource.Properties().AsSlice() {
		tcr.writeProperty(rpt, prop, currentPackage, parentTypes)
	}

	if tcr.optionIncludeFunctions {
		for _, fn := range resource.Functions() {
			tcr.writeFunction(rpt, fn)
		}
	}
}

// writeObjectType writes the object to the debug report.
// rpt is the debug report to write to.
// obj is the object to write.
// currentPackage is the package that the object is defined in (used to simplify type descriptions).
// parentTypes is the set of types that have already been written (used to avoid infinite recursion).
func (tcr *TypeCatalogReport) writeObjectType(
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
	sub := rpt.Addf(
		"%s: %s",
		prop.PropertyName(),
		tcr.asShortNameForType(prop.PropertyType(), currentPackage, parentTypes))

	if def, ok := tcr.asDefinitionToInline(prop.PropertyType(), parentTypes); ok && tcr.inlinedTypes.Contains(def.Name()) {
		pt := parentTypes.Copy()
		pt.Add(def.Name())
		tcr.writeType(sub, def.Type(), currentPackage, pt)
	}

	tcr.writeType(sub, prop.PropertyType(), currentPackage, parentTypes)
}

// writeInterfaceType writes the interface to the debug report.
// rpt is the debug report to write to.
// i is the interface to write.
// currentPackage is the package that the interface is defined in (used to simplify type descriptions).
// parentTypes is the set of types that have already been written (used to avoid infinite recursion).
func (tcr *TypeCatalogReport) writeInterfaceType(
	rpt *StructureReport,
	i *astmodel.InterfaceType,
	_ astmodel.PackageReference,
	_ astmodel.TypeNameSet,
) {
	if tcr.optionIncludeFunctions {
		for _, fn := range i.Functions() {
			tcr.writeFunction(rpt, fn)
		}
	}
}

func (tcr *TypeCatalogReport) writeComplexType(
	rpt *StructureReport,
	propertyType astmodel.Type,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet) {

	// If we have a complex type, we may need to write it out in detail
	switch t := propertyType.(type) {
	case *astmodel.ObjectType,
		*astmodel.InterfaceType,
		*astmodel.ResourceType,
		*astmodel.EnumType,
		*astmodel.OneOfType,
		*astmodel.AllOfType,
		*astmodel.ValidatedType:
		tcr.writeType(rpt, t, currentPackage, parentTypes)
	case *astmodel.OptionalType:
		tcr.writeComplexType(rpt, t.Element(), currentPackage, parentTypes)
	}
}
func (tcr *TypeCatalogReport) writeErroredType(
	rpt *StructureReport,
	et *astmodel.ErroredType,
	currentPackage astmodel.PackageReference,
	types astmodel.TypeNameSet,
) {
	for _, err := range et.Errors() {
		rpt.Addf("Error: %s", err)
	}

	for _, warn := range et.Warnings() {
		rpt.Addf("Warning: %s", warn)
	}

	tcr.writeType(rpt, et.InnerType(), currentPackage, types)
}

func (tcr *TypeCatalogReport) writeValidatedType(
	rpt *StructureReport,
	vt *astmodel.ValidatedType,
	_ astmodel.PackageReference,
	_ astmodel.TypeNameSet,
) {
	for index, rule := range vt.Validations().ToKubeBuilderValidations() {
		rpt.Addf("Rule %d: %s", index, rule)
	}
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

// asShortNameForType returns a short name for the type, relative to the current package.
// t is the type to get a name for.
// currentPackage is the package that we're currently processing (used to simplify type names).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) asShortNameForType(
	t astmodel.Type,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) string {
	// We switch on exact types because we don't want to accidentally unwrap a detail we need
	switch t := t.(type) {
	case astmodel.TypeName:
		// If an inlined type, we use what it points to, otherwise we use the name
		if tcr.inlinedTypes.Contains(t) && !parentTypes.Contains(t) {
			def := tcr.defs[t]
			return tcr.asShortNameForType(def.Type(), currentPackage, parentTypes)
		}

		return astmodel.DebugDescription(t, currentPackage)
	case *astmodel.OptionalType:
		return fmt.Sprintf(
			"*%s",
			tcr.asShortNameForType(t.Element(), currentPackage, parentTypes))
	case *astmodel.ArrayType:
		return fmt.Sprintf(
			"%s[]",
			tcr.asShortNameForType(t.Element(), currentPackage, parentTypes))
	case *astmodel.MapType:
		return fmt.Sprintf(
			"map[%s]%s",
			tcr.asShortNameForType(t.KeyType(), currentPackage, parentTypes),
			tcr.asShortNameForType(t.ValueType(), currentPackage, parentTypes))
	case *astmodel.ResourceType:
		return "Resource"
	case *astmodel.EnumType:
		return fmt.Sprintf(
			"Enum (%s)",
			tcr.formatCount(len(t.Options()), "value", "values"))
	case *astmodel.ObjectType:
		return fmt.Sprintf(
			"Object (%s)",
			tcr.formatCount(t.Properties().Len(), "property", "properties"))
	case *astmodel.OneOfType:
		return fmt.Sprintf(
			"OneOf (%s, %s)",
			tcr.formatCount(len(t.PropertyObjects()), "object", "objects"),
			tcr.formatCount(t.Types().Len(), "option", "options"))
	case *astmodel.AllOfType:
		return fmt.Sprintf(
			"AllOf (%s)",
			tcr.formatCount(t.Types().Len(), "choice", "choices"))
	case *astmodel.ValidatedType:
		return fmt.Sprintf(
			"Validated<%s> (%s)",
			tcr.asShortNameForType(t.Unwrap(), currentPackage, parentTypes),
			tcr.formatCount(len(t.Validations().ToKubeBuilderValidations()), "rule", "rules"))
	case astmodel.MetaType:
		return tcr.asShortNameForType(t.Unwrap(), currentPackage, parentTypes)
	default:
		return astmodel.DebugDescription(t, currentPackage)
	}
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
func (tcr *TypeCatalogReport) writeEnumType(
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

// writeOneOfType writes a oneof to the report.
// rpt is the report to write to.
// oneOf is the oneof to write.
// currentPackage is the package that the oneof is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeOneOfType(
	rpt *StructureReport,
	oneOf *astmodel.OneOfType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	if oneOf.HasDiscriminatorProperty() {
		rpt.Addf("discriminator: %s", oneOf.DiscriminatorProperty())
	}

	if oneOf.HasDiscriminatorValue() {
		rpt.Addf("discriminator value: %s", oneOf.DiscriminatorValue())
	}

	if propertyObjects := oneOf.PropertyObjects(); len(propertyObjects) > 0 {
		// We expect the order of PropertyObjects() to be consistent, so no need to sort
		sub := rpt.Addf("Property Objects (%s):", tcr.formatCount(len(propertyObjects), "object", "objects"))
		for _, t := range propertyObjects {
			subsub := sub.Addf("%s", tcr.asShortNameForType(t, currentPackage, parentTypes))
			tcr.writeComplexType(subsub, t, currentPackage, parentTypes)
		}
	}

	// The order of entries in oneOf.Types() can vary but isn't significant
	// So we pull them out and write them in sorted order
	options := oneOf.Types()
	if options.Len() > 0 {
		optionTypes := make([]astmodel.Type, 0, options.Len())
		typesToNames := make(map[astmodel.Type]string, options.Len())
		options.ForEach(func(t astmodel.Type, _ int) {
			name := tcr.asShortNameForType(t, currentPackage, parentTypes)
			optionTypes = append(optionTypes, t)
			typesToNames[t] = name
		})

		sort.Slice(optionTypes, func(i, j int) bool {
			iname := typesToNames[optionTypes[i]]
			jname := typesToNames[optionTypes[j]]
			return iname < jname
		})

		sub := rpt.Addf("Options (%s):", tcr.formatCount(len(optionTypes), "option", "options"))
		for index, t := range optionTypes {
			subsub := sub.Addf("Option %d: %s", index, tcr.asShortNameForType(t, currentPackage, parentTypes))
			tcr.writeComplexType(subsub, t, currentPackage, parentTypes)
		}
	}
}

// writeAllOfType writes an allof to the report.
// rpt is the report to write to.
// allOfType is the allof to write.
// currentPackage is the package that the allof is defined in (used to simplify type descriptions).
// parentTypes is the set of types that are currently being written (used to detect cycles).
func (tcr *TypeCatalogReport) writeAllOfType(
	rpt *StructureReport,
	allOf *astmodel.AllOfType,
	currentPackage astmodel.PackageReference,
	parentTypes astmodel.TypeNameSet,
) {
	allOf.Types().ForEach(func(t astmodel.Type, index int) {
		sub := rpt.Addf("option %d: %s", index, tcr.asShortNameForType(t, currentPackage, parentTypes))
		tcr.writeType(sub, t, currentPackage, parentTypes)
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

func (tcr *TypeCatalogReport) formatCount(value int, singular string, plural string) string {
	if value == 1 {
		return fmt.Sprintf("%d %s", value, singular)
	}

	return fmt.Sprintf("%d %s", value, plural)
}
