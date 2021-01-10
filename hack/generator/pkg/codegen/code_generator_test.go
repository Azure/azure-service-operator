/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

// RemoveStages will remove all stages from the pipeline with the given ids.
// Only available for test builds.
// Will panic if you specify an unknown id.
func (generator *CodeGenerator) RemoveStages(stageIds ...string) {
	stagesToRemove := make(map[string]bool)
	for _, s := range stageIds {
		stagesToRemove[s] = false
	}

	var pipeline []PipelineStage

	for _, stage := range generator.pipeline {
		if _, ok := stagesToRemove[stage.id]; ok {
			stagesToRemove[stage.id] = true
			continue
		}

		pipeline = append(pipeline, stage)
	}

	for stage, removed := range stagesToRemove {
		if !removed {
			panic(fmt.Sprintf("Expected to remove stage %s from pipeline, but it wasn't found.", stage))
		}
	}

	generator.pipeline = pipeline
}

// ReplaceStage replaces all uses of an existing stage with another one.
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) ReplaceStage(existingStage string, stage PipelineStage) {
	replaced := false
	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			generator.pipeline[i] = stage
			replaced = true
		}
	}

	if !replaced {
		panic(fmt.Sprintf("Expected to replace stage %s but it wasn't found", existingStage))
	}
}

// InjectStageAfter injects a new stage immediately after the first occurrence of an existing stage
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) InjectStageAfter(existingStage string, stage PipelineStage) {
	injected := false

	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			var p []PipelineStage
			p = append(p, generator.pipeline[:i+1]...)
			p = append(p, stage)
			p = append(p, generator.pipeline[i+1:]...)
			generator.pipeline = p
			injected = true
			break
		}
	}

	if !injected {
		panic(fmt.Sprintf("Expected to inject stage %s but %s wasn't found", stage.id, existingStage))
	}
}

// HasStage tests whether the pipeline has a stage with the given id
// Only available for test builds.
func (generator *CodeGenerator) HasStage(id string) bool {
	return generator.IndexOfStage(id) != -1
}

// IndexOfStage returns the index of the stage, if present, or -1 if not
// Only available for test builds.
func (generator *CodeGenerator) IndexOfStage(id string) int {
	for i, s := range generator.pipeline {
		if s.HasId(id) {
			return i
		}
	}

	return -1
}

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
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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

func MakeFakePipelineStage(id string) PipelineStage {
	return MakePipelineStage(
		id, "Stage "+id, func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			return types, nil
		})
}

/*
 * ReplaceStageTests
 */

func TestReplaceStage_ReplacesSpecifiedStage(t *testing.T) {
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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
	g := NewGomegaWithT(t)

	gen := &CodeGenerator{
		pipeline: []PipelineStage{
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
