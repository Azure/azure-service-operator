package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
)

// injectOriginalVersionFunctionId is the unique identifier for this pipeline stage
const injectOriginalVersionFunctionId = "injectOriginalVersionFunction"

// InjectOriginalVersionFunction injects the function OriginalVersion() into each Spec type
// This function allows us to recover the original version used to create each custom resource, giving the operator the
// information needed to interact with ARM using the correct API version.
func InjectOriginalVersionFunction(idFactory astmodel.IdentifierFactory) Stage {

	return MakeStage(
		injectOriginalVersionFunctionId,
		"Inject the function OriginalVersion() into each Spec type",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			injector := storage.NewFunctionInjector()
			result := types.Copy()

			specs := storage.FindSpecTypes(types)
			for name, def := range specs {
				fn := functions.NewOriginalVersionFunction(idFactory)
				defWithFn, err := injector.Inject(def, fn)
				if err != nil {
					return nil, errors.Wrapf(err, "injecting OriginalVersion() into %s", name)
				}

				result[defWithFn.Name()] = defWithFn
			}

			return result, nil
		})
}
