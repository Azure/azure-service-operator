/* Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func SimplifySwaggerNames(idFactory astmodel.IdentifierFactory, config *config.Configuration) Stage {
	return MakeLegacyStage(
		"simplifySwaggerNames",
		"Remove redundant components from Swagger names",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			targetCounts := make(map[astmodel.TypeName]int)
			for tn := range types {
				newName := simplifySwaggerName(tn)
				targetCounts[newName]++
			}

			renames := make(map[astmodel.TypeName]astmodel.TypeName)
			for tn := range types {
				newName := simplifySwaggerName(tn)
				// if >1 that means multiple names mapped to the same name, so don’t rename those
				if targetCounts[newName] == 1 {
					renames[tn] = newName
				}
			}

			renamer := astmodel.NewRenamingVisitor(renames)
			return renamer.RenameAll(types)
		})
}

var redundantComponents = regexp.MustCompile("CreateParameters|UpdateParameters")

func simplifySwaggerName(tn astmodel.TypeName) astmodel.TypeName {
	return tn.WithName(redundantComponents.ReplaceAllLiteralString(tn.Name(), ""))
}
