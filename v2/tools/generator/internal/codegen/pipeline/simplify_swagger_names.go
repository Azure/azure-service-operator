/* Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func SimplifySwaggerNames(idFactory astmodel.IdentifierFactory, config *config.Configuration) Stage {
	return MakeLegacyStage(
		"simplifySwaggerNames",
		"Remove redundant components from Swagger names",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			// build reverse lookup so we can see how many names would be renamed to each target name
			renamedTo := make(map[astmodel.TypeName]astmodel.StringSet)
			for tn := range types {
				newName := simplifySwaggerName(tn)

				set, ok := renamedTo[newName]
				if !ok {
					set = make(astmodel.StringSet)
					renamedTo[newName] = set
				}

				set.Add(tn.Name())
			}

			logged := make(astmodel.StringSet)

			renames := make(map[astmodel.TypeName]astmodel.TypeName)
			for tn := range types {
				newName := simplifySwaggerName(tn)
				// if >1 that means multiple names mapped to the same name, so don’t rename those
				if len(renamedTo[newName]) == 1 {
					renames[tn] = newName
				} else {
					if !logged.Contains(newName.Name()) {
						logged.Add(newName.Name())

						clashers := []string{}
						for v := range renamedTo[newName] {
							clashers = append(clashers, v)
						}

						klog.Infof("Won't rename anything to %s, %d names would clash: %s\n", newName.Name(), len(clashers), strings.Join(clashers, ", "))
					}
				}
			}

			renamer := astmodel.NewRenamingVisitor(renames)
			return renamer.RenameAll(types)
		})
}

type replacement struct{ from, to string }

// this is a slice instead of a map because replacement
// order matters here; CreateUpdateParameters must be replaced
// before UpdateParameters, for example.
var replacements = []replacement{
	{from: "CreateUpdateParameters", to: ""},
	{from: "UpdateParameters", to: ""},
	{from: "CreateParameters", to: ""},

	{from: "CreateUpdateProperties", to: "Properties"},
	{from: "UpdateProperties", to: "Properties"},
	{from: "CreateProperties", to: "Properties"},
}

func simplifySwaggerName(tn astmodel.TypeName) astmodel.TypeName {
	name := tn.Name()

	for _, replace := range replacements {
		name = strings.ReplaceAll(name, replace.from, replace.to)
	}

	return tn.WithName(name)
}
