/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"

	. "github.com/onsi/gomega"
)

func TestGolden_NewARMCodeGeneratorFromConfigCreatesRightPipeline(t *testing.T) {
	t.Parallel()
	gold := goldie.New(t)
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	configuration := config.NewConfiguration()

	codegen, err := NewTargetedCodeGeneratorFromConfig(configuration, idFactory, pipeline.ARMTarget, logr.Discard())
	g.Expect(err).To(Succeed())

	result := writePipeline("Expected Pipeline Stages for ARM Code Generation", codegen)

	// When reviewing changes to the golden file, ensure they make sense in the context of an operator built to work
	// against Azure ARM - we don't want to see any Crossplane specific stages showing up there.
	gold.Assert(t, t.Name(), result)
}

func TestGolden_NewCrossplaneCodeGeneratorFromConfigCreatesRightPipeline(t *testing.T) {
	t.Parallel()
	gold := goldie.New(t)
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	configuration := config.NewConfiguration()

	codegen, err := NewTargetedCodeGeneratorFromConfig(configuration, idFactory, pipeline.CrossplaneTarget, logr.Discard())
	g.Expect(err).To(Succeed())

	result := writePipeline("Expected Pipeline Stages for ARM Code Generation", codegen)

	// When reviewing changes to the golden file, ensure they make sense in the context of an operator built to work
	// with Crossplane - we don't want to see any Azure ARM specific stages showing up there.
	gold.Assert(t, t.Name(), result)
}

func TestGolden_NewTestCodeGeneratorCreatesRightPipeline(t *testing.T) {
	t.Parallel()
	gold := goldie.New(t)
	g := NewGomegaWithT(t)

	cfg := makeDefaultTestConfig()
	codegen, err := NewTestCodeGenerator("Sample", "path", t, cfg, config.GenerationPipelineAzure)
	g.Expect(err).To(BeNil())

	result := writePipeline("Expected Pipeline Stages for Test Code Generation", codegen)

	// When reviewing changes to the golden file, ensure they make sense in the context of the tests we are running
	// of the entire pipeline; you may need to explicity exclude some stages.
	gold.Assert(t, t.Name(), result)
}

func writePipeline(title string, codegen *CodeGenerator) []byte {
	var b bytes.Buffer

	fmt.Fprintln(&b, title)
	fmt.Fprintln(&b, strings.Repeat("-", len(title)))

	idWidth := 0
	for _, s := range codegen.pipeline {
		if len(s.Id()) > idWidth {
			idWidth = len(s.Id())
		}
	}

	format := fmt.Sprintf("%%-%ds %%-10s %%s\n", idWidth+4)

	for _, s := range codegen.pipeline {
		targets := ""
		for _, t := range s.Targets() {
			if len(targets) > 0 {
				targets = targets + "; "
			}

			targets = targets + t.String()
		}

		fmt.Fprintf(&b, format, s.Id(), targets, s.Description())
	}

	return b.Bytes()
}
