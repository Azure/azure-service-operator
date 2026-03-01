# Project: RAPID Testing

Overall goal is to migrate fom the existing use of gopter tests to instead use the rapid library. Because this is a big change, we're going to tackle it step by step. This file is our overall plan, including reference information, and should be updated as a part of each step.

The rapid library is available from https://github.com/flyingmutant/rapid

The gopter library we're replacing is github.com/leanovate/gopter

## Phase I: Research gopter

Read up on the gopter library and how it's used the ASO code generator.

The pipeline stage is in json_serialization_test_cases.go

A sample of the generated gopter code:
v2/api/batch/v20240701/batch_account_types_gen_test.go

The code generator is found at v2/tools/generator. It's complex, so don't try to read all of it in one go.

Add notes to this section detailing what you find.

## Phase II: Skeleton

Add a new code generator pipeline stage that executes immediately after the existing stage. 

Initially this will do nothing due to a deny list that blocks it for all gropus - use the similar deny list in v2/internal/reconcilers/arm/azure_generic_arm_reconciler_instance.go as a model.

Additionaly, modify the gopter generator in json_serialization_test_cases to use the same deny list, but inverted so that we get either gopter or rapid tests for each group.

## Phase III: Implementation

Expand the skeleton from Phase II so that it fully implements the required tests using rapid. Enable just the "batch" group for this new testing style.

