/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"fmt"
	"testing"

	"github.com/go-openapi/spec"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func Example_inferNameFromURLPath() {
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, _ := extractor.inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}")
	fmt.Printf("%s/%s: %s", group, resource, name)
	// Output: Microsoft.GroupName/resourceName: ResourceName
}

func Example_inferNameFromURLPath_ChildResources() {
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, _ := extractor.inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}/someChild/{childId}")
	fmt.Printf("%s/%s: %s", group, resource, name)
	// Output: Microsoft.GroupName/resourceName/someChild: ResourceName_SomeChild
}

func Test_InferNameFromURLPath_FailsWithMultipleParametersInARow(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	_, _, _, err := extractor.inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}/{anotherParameter}")
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("multiple parameters"))
}

func Test_InferNameFromURLPath_FailsWithNoGroupName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	_, _, _, err := extractor.inferNameFromURLPath("/resourceName/{resourceId}/{anotherParameter}")
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("no group name"))
}

func Test_InferNameFromURLPath_SkipsDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, err := extractor.inferNameFromURLPath("Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}")
	g.Expect(err).To(BeNil())
	g.Expect(group).To(Equal("Microsoft.Storage"))
	g.Expect(resource).To(Equal("storageAccounts/blobServices/containers"))
	g.Expect(name).To(Equal("StorageAccounts_BlobServices_Container"))
}

func Test_extractLastPathParam_ExtractsParameter(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := SwaggerTypeExtractor{}

	param := spec.Parameter{
		CommonValidations: spec.CommonValidations{
			Enum: []interface{}{"yes", "no"},
		},
		ParamProps: spec.ParamProps{
			In:       "path",
			Name:     "value",
			Required: true,
		},
	}

	lastParam, ok := extractor.extractLastPathParam("/some/{value}", []spec.Parameter{param})
	g.Expect(ok).To(BeTrue())
	g.Expect(lastParam).To(Equal(param))
}

func Test_extractLastPathParam_ExtractsParameterFromMultipleParameters(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := SwaggerTypeExtractor{}
	param := spec.Parameter{
		CommonValidations: spec.CommonValidations{
			Enum: []interface{}{"orange", "blue"},
		},
		ParamProps: spec.ParamProps{
			In:       "path",
			Name:     "value2",
			Required: true,
		},
	}

	lastParam, ok := extractor.extractLastPathParam("/some/{value1}/{value2}", []spec.Parameter{
		{
			CommonValidations: spec.CommonValidations{
				Enum: []interface{}{"yes", "no"},
			},
			ParamProps: spec.ParamProps{
				In:       "path",
				Name:     "value1",
				Required: true,
			},
		},
		param,
	})

	g.Expect(ok).To(BeTrue())
	g.Expect(lastParam).To(Equal(param))
}

func Test_extractLastPathParam_StaticParameterName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := SwaggerTypeExtractor{}
	param := spec.Parameter{
		SimpleSchema: spec.SimpleSchema{
			Type: "string",
		},
		ParamProps: spec.ParamProps{
			Name:        "name",
			In:          "path",
			Required:    true,
			Description: "The name",
		},
		CommonValidations: spec.CommonValidations{
			Enum: []interface{}{
				"default",
			},
		},
	}

	lastParam, ok := extractor.extractLastPathParam("/some/{value1}/default", []spec.Parameter{
		{
			CommonValidations: spec.CommonValidations{
				Enum: []interface{}{"yes", "no"},
			},
			ParamProps: spec.ParamProps{
				In:       "path",
				Name:     "value1",
				Required: true,
			},
		},
		param,
	})

	g.Expect(ok).To(BeTrue())
	g.Expect(lastParam).To(Equal(param))
}
