/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline/recursivetypefixer"
)

const UnrollRecursiveTypesStageID = "unrollRecursiveTypes"

// UnrollRecursiveTypes finds types that reference themselves and "unrolls" the reference.
func UnrollRecursiveTypes() *Stage {
	/*
	 * So a type that looks like:
	 * type Error struct {
	 *     code    string
	 *     message string
	 *     errors []Error
	 * }
	 * gets unrolled to look like:
	 * type Error struct {
	 *     code    string
	 *     message string
	 *     errors []Error_Unrolled
	 * }
	 * where Error_Unrolled looks like:
	 * type Error_Unrolled struct {
	 *     code    string
	 *     message string
	 * }
	 * The recursive references must be removed because
	 * controller-tools doesn't support generating "references" (JSON $ref) so it can't support recursive types today.
	 * See https://github.com/kubernetes-sigs/controller-tools/issues/489 for more information.
	 * Unrolling these types, while required for controller-tools to function, means
	 * THAT THE TYPES WE GENERATE DON'T EXACTLY CONFORM TO THE PUBLISHED API!
	 * In practice, we think this is ok for Error types because our observation is that Errors that reference themselves only return
	 * a depth of 1 in practice.
	 * If we were to unroll all loops (rather than just types that directly reference themselves like we're doing here) that
	 * "in practice" observation may no longer hold, so we avoid doing it here (it's also more complicated to do that).
	 */

	return NewStage(
		UnrollRecursiveTypesStageID,
		"Unroll directly recursive types since they are not supported by controller-gen",
		func(ctx context.Context, state *State) (*State, error) {
			result := make(astmodel.TypeDefinitionSet)

			// Find object types that reference themselves
			fixer := recursivetypefixer.NewSimpleRecursiveTypeFixer()

			for _, def := range state.Definitions() {
				updatedDef, err := fixer.Fix(def)
				if err != nil {
					return nil, err
				}

				// Confirm that all the changed types have Error in their name. If the type doesn't have Error
				// we can't be sure it's safe to unroll, so we raise an error. This is fatal because controller-gen
				// will error on recursive types when we run it later and we want to fail fast.
				if !astmodel.TypeEquals(def.Type(), updatedDef.Type()) && !strings.Contains(def.Name().Name(), "Error") {
					return nil, errors.Errorf("%q is directly recursive and cannot be unrolled because it doesn't contain 'Error'", def.Name())
				}

				result.Add(updatedDef)
			}

			// Add new unrolled types
			result.AddTypes(fixer.Types())

			return state.WithDefinitions(result), nil
		})
}
