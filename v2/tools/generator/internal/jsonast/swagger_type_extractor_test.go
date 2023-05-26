/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/go-openapi/spec"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func Test_InferNameFromURLPath_ParentResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, err := extractor.inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}")
	t.Logf("%s/%s: %s", group, resource, name)
	// Output: Microsoft.GroupName/resourceName: ResourceName
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(group).To(Equal("Microsoft.GroupName"))
	g.Expect(resource).To(Equal("resourceName"))
	g.Expect(name).To(Equal("ResourceName"))
}

func Test_InferNameFromURLPath_ChildResources(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, err := extractor.inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}/someChild/{childId}")
	t.Logf("%s/%s: %s", group, resource, name)
	// Output: Microsoft.GroupName/resourceName/someChild: ResourceName_SomeChild
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(group).To(Equal("Microsoft.GroupName"))
	g.Expect(resource).To(Equal("resourceName/someChild"))
	g.Expect(name).To(Equal("ResourceName_SomeChild"))
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

func Test_InferNameFromURLPath_ExtensionResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, resource, name, err := extractor.inferNameFromURLPath("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}")
	g.Expect(err).To(BeNil())
	g.Expect(group).To(Equal("Microsoft.Authorization"))
	g.Expect(resource).To(Equal("roleAssignments"))
	g.Expect(name).To(Equal("RoleAssignment"))
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
	})

	g.Expect(ok).To(BeTrue())
	g.Expect(lastParam).To(Equal(param))
}

func Test_ExtractResourceSubpath(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}

	group, subpath, err := extractor.extractResourceSubpath("whatever/{whateverParam}/stuff/Microsoft.GroupName/resourceName/{resourceId}/{anotherParameter}")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(group).To(Equal("Microsoft.GroupName"))
	g.Expect(subpath).To(Equal("resourceName/{resourceId}/{anotherParameter}"))
}

func Test_ExpandAndCanonicalizePath_DoesNotExpandSimplePath(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ctx := context.Background()
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}
	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory(), config.NewConfiguration(), logr.Discard())

	parameters := []spec.Parameter{
		makeSubscriptionIDParameter(),
		makeResourceGroupParameter(),
		{
			ParamProps: spec.ParamProps{
				Name: "typeName",
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"string"},
					},
				},
			},
		},
	}

	paths := extractor.expandAndCanonicalizePath(
		ctx,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/{typeName}",
		scanner,
		parameters)

	g.Expect(paths).To(HaveLen(1))
	g.Expect(paths[0]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/{typeName}"))
}

func Test_ExpandAndCanonicalizePath_ExpandsSingleValueEnumInNameLocationWithDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ctx := context.Background()
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}
	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory(), config.NewConfiguration(), logr.Discard())

	parameters := []spec.Parameter{
		makeSubscriptionIDParameter(),
		makeResourceGroupParameter(),
		{
			ParamProps: spec.ParamProps{
				Name: "typeName",
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"string"},
						Enum: []interface{}{
							"default",
						},
					},
				},
			},
		},
	}

	paths := extractor.expandAndCanonicalizePath(
		ctx,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/{typeName}",
		scanner,
		parameters)

	g.Expect(paths).To(HaveLen(1))
	g.Expect(paths[0]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/default"))
}

func Test_ExpandAndCanonicalizePath_DoesNotExpandSingleValueEnumWithoutDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ctx := context.Background()
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}
	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory(), config.NewConfiguration(), logr.Discard())

	parameters := []spec.Parameter{
		makeSubscriptionIDParameter(),
		makeResourceGroupParameter(),
		{
			ParamProps: spec.ParamProps{
				Name: "typeName",
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"string"},
						Enum: []interface{}{
							"current",
						},
					},
				},
			},
		},
	}

	paths := extractor.expandAndCanonicalizePath(
		ctx,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/{typeName}",
		scanner,
		parameters)

	g.Expect(paths).To(HaveLen(1))
	g.Expect(paths[0]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/type/{typeName}"))
}

func Test_ExpandAndCanonicalizePath_ExpandsEnumInResourceTypePath(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ctx := context.Background()
	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}
	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory(), config.NewConfiguration(), logr.Discard())

	parameters := []spec.Parameter{
		makeSubscriptionIDParameter(),
		makeResourceGroupParameter(),
		{
			ParamProps: spec.ParamProps{
				Name: "type",
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"string"},
						Enum: []interface{}{
							"a",
							"b",
							"c",
						},
					},
				},
			},
		},
	}

	paths := extractor.expandAndCanonicalizePath(
		ctx,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/{type}/{typeName}",
		scanner,
		parameters)

	g.Expect(paths).To(HaveLen(3))
	g.Expect(paths[0]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/a/{typeName}"))
	g.Expect(paths[1]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/b/{typeName}"))
	g.Expect(paths[2]).To(Equal("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Group/c/{typeName}"))
}

func makeSubscriptionIDParameter() spec.Parameter {
	return spec.Parameter{
		ParamProps: spec.ParamProps{
			Name: "subscriptionId",
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type: spec.StringOrArray{"string"},
				},
			},
		},
	}
}

func makeResourceGroupParameter() spec.Parameter {
	return spec.Parameter{
		ParamProps: spec.ParamProps{
			Name: "resourceGroup",
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type: spec.StringOrArray{"string"},
				},
			},
		},
	}
}
