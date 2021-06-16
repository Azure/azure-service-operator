package test

import (
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

func CreateFileDefinition(definitions ...astmodel.TypeDefinition) *astmodel.FileDefinition {
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

	// Use the package reference of the first definition for the whole file
	ref, err := astmodel.PackageAsLocalPackage(definitions[0].Name().PackageReference)
	if err != nil {
		panic(errors.Wrap(err, "Expected first definition to have a local package reference - fix your test!"))
	}

	pkgDefinition := astmodel.NewPackageDefinition(ref.Group(), ref.PackageName(), ref.Version())
	for _, def := range definitions {
		pkgDefinition.AddDefinition(def)
	}

	packages[ref] = pkgDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewFileDefinition(ref, definitions, packages)
	return fileDef
}
