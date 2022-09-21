/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/pkg/errors"
	"io"
	"os"
	"sort"
)

type TypeCatalogReport struct {
	title        string
	defs         astmodel.TypeDefinitionSet
	inlinedTypes astmodel.TypeNameSet // Set of types that we inline when generating the report
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

// InlineTypes specifies that the generated report should inline types where referenced,
// We achieve this by scanning for properties with types we have definitions for
func (tcr *TypeCatalogReport) InlineTypes() {
	for _, def := range tcr.defs {
		if c, ok := astmodel.AsPropertyContainer(def.Type()); ok {
			tcr.inlineTypesFrom(c)
		}
	}
}

// inlineTypesFrom inlines the types referenced by the property container
func (tcr *TypeCatalogReport) inlineTypesFrom(container astmodel.PropertyContainer) {
	for _, prop := range container.Properties().AsSlice() {
		if n, ok := astmodel.AsTypeName(prop.PropertyType()); ok {
			tcr.inlinedTypes.Add(n)
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
func (tcr *TypeCatalogReport) writeDefinition(rpt *StructureReport, definition astmodel.TypeDefinition) {
	name := definition.Name()
	t := definition.Type()
	sub := rpt.Addf("%s: %s", name.Name(), astmodel.DebugDescription(t, name.PackageReference))
	tcr.writeType(sub, definition.Type(), name.PackageReference)
}

// writeType writes the type to the debug report.
// rpt is the debug report to write to.
// t is the type to write.
// currentPackage is the package that the type is defined in (used to simplify type descriptions).
// Only complex types where astmodel.DebugDescription is insufficient are written.
func (tcr *TypeCatalogReport) writeType(
	rpt *StructureReport,
	t astmodel.Type,
	currentPackage astmodel.PackageReference,
) {
	if rsrc, ok := astmodel.AsResourceType(t); ok {
		tcr.writeResource(rpt, rsrc, currentPackage)
	} else if obj, ok := astmodel.AsObjectType(t); ok {
		tcr.writeObject(rpt, obj, currentPackage)
	} else if obj, ok := astmodel.AsEnumType(t); ok {
		tcr.writeEnum(rpt, obj, currentPackage)
	}
}

// writeResource writes the resource to the debug report.
// rpt is the debug report to write to.
// resource is the resource to write.
// currentPackage is the package that the resource is defined in (used to simplify type descriptions).
func (tcr *TypeCatalogReport) writeResource(
	rpt *StructureReport,
	resource *astmodel.ResourceType,
	currentPackage astmodel.PackageReference,
) {
	for _, prop := range resource.Properties().AsSlice() {
		tcr.writeProperty(rpt, prop, currentPackage)
	}

	for _, fn := range resource.Functions() {
		tcr.writeFunction(rpt, fn)
	}
}

// writeObject writes the object to the debug report.
func (tcr *TypeCatalogReport) writeObject(
	rpt *StructureReport,
	obj *astmodel.ObjectType,
	currentPackage astmodel.PackageReference,
) {
	for _, prop := range obj.Properties().AsSlice() {
		tcr.writeProperty(rpt, prop, currentPackage)
	}

	for _, fn := range obj.Functions() {
		tcr.writeFunction(rpt, fn)
	}
}

func (tcr *TypeCatalogReport) writeProperty(
	rpt *StructureReport,
	prop *astmodel.PropertyDefinition,
	currentPackage astmodel.PackageReference,
) {

	propertyType := prop.PropertyType()

	if n, ok := astmodel.AsTypeName(propertyType); ok && tcr.inlinedTypes.Contains(n) {
		// We're inlining types, so don't bother writing the property type
		sub := rpt.Addf("%s", prop.PropertyName())
		if def, ok := tcr.defs[n]; ok {
			tcr.inlinedTypes.Add(n)
			tcr.writeType(sub, def.Type(), currentPackage)
		} else {
			sub.Addf("No definition found for %s", n)
		}

		return
	}

	// By default, just give the name of the property and a description of the type
	rpt.Addf(
		"%s: %s",
		prop.PropertyName(),
		astmodel.DebugDescription(propertyType, currentPackage))
}

func (tcr *TypeCatalogReport) writeFunction(
	rpt *StructureReport,
	fn astmodel.Function,
) {
	rpt.Addf("%s()", fn.Name())
}

func (tcr *TypeCatalogReport) writeEnum(
	rpt *StructureReport,
	enum *astmodel.EnumType,
	currentPackage astmodel.PackageReference,
) {
	tcr.writeType(rpt, enum.BaseType(), currentPackage)
	for _, v := range enum.Options() {
		rpt.Addf("%s: %s", v.Identifier, v.Value)
	}
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
