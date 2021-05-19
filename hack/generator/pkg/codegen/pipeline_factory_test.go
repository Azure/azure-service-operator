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

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"github.com/sebdah/goldie/v2"

	. "github.com/onsi/gomega"
)

func TestNewARMCodeGeneratorFromConfigCreatesRightPipeline(t *testing.T) {
	gold := goldie.New(t)
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	configuration := config.NewConfiguration()
	configuration.Pipeline = config.GenerationPipelineAzure

	codegen, err := NewCodeGeneratorFromConfig(configuration, idFactory)
	g.Expect(err).To(BeNil())

	result := writePipeline("Expected Pipeline Stages for ARM Code Generation", codegen)

	gold.Assert(t, "ARMCodeGeneratorPipeline", result)
}

func TestNewTestCodeGeneratorCreatesRightPipeline(t *testing.T) {
	gold := goldie.New(t)
	g := NewGomegaWithT(t)

	cfg := makeDefaultTestConfig()
	codegen, err := NewTestCodeGenerator("Sample", "path", t, cfg, config.GenerationPipelineAzure)
	g.Expect(err).To(BeNil())

	result := writePipeline("Expected Pipeline Stages for Test Code Generation", codegen)

	gold.Assert(t, "TestCodeGeneratorPipeline", result)
}

func writePipeline(title string, codegen *CodeGenerator) []byte {
	var b bytes.Buffer

	fmt.Fprintln(&b, title)
	fmt.Fprintln(&b, strings.Repeat("-", len(title)))

	for _, s := range codegen.pipeline {
		targets := ""
		for _, t := range s.targets {
			if len(targets) > 0 {
				targets = targets + "; "
			}

			targets = targets + t.String()
		}

		fmt.Fprintf(&b, "%-35s %-10s %s\n", s.id, targets, s.description)
	}

	return b.Bytes()
}
