/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
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
	return pipeline.NewLegacyStage(
		id, "Stage "+id, func(ctx context.Context, defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			return defs, nil
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

/*
 * verifyPipeline Tests
 */

func TestVerifyPipeline_GivenNoPrerequisites_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			bazStage,
		},
	}

	g.Expect(gen.verifyPipeline()).To(BeNil())
}

func TestVerifyPipeline_GivenSatisfiedPrerequisites_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPrerequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			stage,
			bazStage,
		},
	}

	g.Expect(gen.verifyPipeline()).To(BeNil())
}

func TestVerifyPipeline_GivenUnsatisfiedPrerequisites_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPrerequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			stage,
			bazStage,
		},
	}

	err := gen.verifyPipeline()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(stage.Id()))
	g.Expect(err.Error()).To(ContainSubstring(barStage.Id()))
}

func TestVerifyPipeline_GivenOutOfOrderPrerequisites_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPrerequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			stage,
			barStage,
			bazStage,
		},
	}

	err := gen.verifyPipeline()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(stage.Id()))
	g.Expect(err.Error()).To(ContainSubstring(barStage.Id()))
}

func TestVerifyPipeline_GivenSatisfiedPostrequisites_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPostrequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			stage,
			barStage,
			bazStage,
		},
	}

	err := gen.verifyPipeline()
	g.Expect(err).To(BeNil())
}

func TestVerifyPipeline_GivenUnsatisfiedPostrequisites_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPrerequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			stage,
			bazStage,
		},
	}

	err := gen.verifyPipeline()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(stage.Id()))
	g.Expect(err.Error()).To(ContainSubstring(barStage.Id()))
}

func TestVerifyPipeline_GivenOutOfOrderPostrequisites_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	stage := MakeFakePipelineStage("stage").WithRequiredPostrequisites(barStage.Id())

	gen := &CodeGenerator{
		pipeline: []*pipeline.Stage{
			fooStage,
			barStage,
			stage,
			bazStage,
		},
	}

	err := gen.verifyPipeline()
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(stage.Id()))
	g.Expect(err.Error()).To(ContainSubstring(barStage.Id()))
}
