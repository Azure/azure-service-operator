/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"

	. "github.com/onsi/gomega"
)

/*
 * Shared test data
 */

var (
	fooStage = MakeFakePipelineStage("foo")
	barStage = MakeFakePipelineStage("bar")
	bazStage = MakeFakePipelineStage("baz")
	zooStage = MakeFakePipelineStage("zoo")
)

/*
 * RemoveStagesTests
 */

func TestRemoveStages_RemovesSpecifiedStages(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	gen.RemoveStages("foo", "baz")
	g.Expect(gen.pipeline).To(HaveLen(1))
	g.Expect(gen.pipeline[0].HasId("bar")).To(BeTrue())
}

func TestRemoveStages_PanicsForUnknownStage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	g.Expect(func() {
		gen.RemoveStages("bang")
	},
	).To(Panic())

	gen.RemoveStages("foo", "baz")
}

func MakeFakePipelineStage(id string) *pipeline.Stage {
	return pipeline.NewStage(
		id,
		"Stage "+id,
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			return state, nil
		})
}

/*
 * ReplaceStageTests
 */

func TestReplaceStage_ReplacesSpecifiedStage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	gen.ReplaceStage("foo", zooStage)
	g.Expect(gen.pipeline).To(HaveLen(3))
	g.Expect(gen.HasStage("foo")).To(BeFalse())
	g.Expect(gen.HasStage("zoo")).To(BeTrue())
}

func TestReplaceStage_PanicsForUnknownStage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	g.Expect(func() {
		gen.ReplaceStage("bang", zooStage)
	},
	).To(Panic())
}

/*
 * InjectStageAfterTests
 */

func TestInjectStageAfter_InjectsSpecifiedStage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	gen.InjectStageAfter("foo", zooStage)
	g.Expect(gen.pipeline).To(HaveLen(4))
	g.Expect(gen.IndexOfStage("foo")).To(Equal(0))
	g.Expect(gen.IndexOfStage("zoo")).To(Equal(1))
}

func TestInjectStageAfter_PanicsForUnknownStage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	g.Expect(func() {
		gen.InjectStageAfter("bang", zooStage)
	},
	).To(Panic())
}
