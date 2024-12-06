/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"github.com/go-openapi/spec"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func Test_InferNameFromPath_GivenURL_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		path     string
		group    string
		resource string
		name     string
		err      string
	}{
		"ParentResource": {
			path:     "/Microsoft.GroupName/resourceName/{resourceId}",
			group:    "Microsoft.GroupName",
			resource: "resourceName",
			name:     "ResourceName",
		},
		"ChildResources": {
			path:     "/Microsoft.GroupName/resourceName/{resourceId}/someChild/{childId}",
			group:    "Microsoft.GroupName",
			resource: "resourceName/someChild",
			name:     "ResourceName_SomeChild",
		},
		"FailsWithMultipleParametersInARow": {
			path: "/Microsoft.GroupName/resourceName/{resourceId}/{anotherParameter}",
			err:  "multiple parameters",
		},
		"FailsWithNoGroupName": {
			path: "/resourceName/{resourceId}/{anotherParameter}",
			err:  "no group name",
		},
		"SkipsDefault": {
			path:     "Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}",
			group:    "Microsoft.Storage",
			resource: "storageAccounts/blobServices/containers",
			name:     "StorageAccounts_BlobServices_Container",
		},
		"SkipsWeb": {
			path:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web",
			group:    "Microsoft.Web",
			resource: "sites/slots/sourcecontrols",
			name:     "Sites_Slots_Sourcecontrol",
		},
		"ExtensionResource": {
			path:     "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}",
			group:    "Microsoft.Authorization",
			resource: "roleAssignments",
			name:     "RoleAssignment",
		},
	}

	extractor := &SwaggerTypeExtractor{
		idFactory: astmodel.NewIdentifierFactory(),
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			group, resource, name, err := extractor.inferNameFromURLPath(c.path)
			t.Logf("%s/%s: %s", group, resource, name)

			if c.err != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(c.err))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(group).To(Equal(c.group))
				g.Expect(resource).To(Equal(c.resource))
				g.Expect(name).To(Equal(c.name))
			}
		})
	}
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

func TestCategorizeResourceScope(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		path     string
		expected astmodel.ResourceScope
	}{
		"virtual machine": {
			path:     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}",
			expected: astmodel.ResourceScopeResourceGroup,
		},
		"role assignment": {
			path:     "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}",
			expected: astmodel.ResourceScopeExtension,
		},
		"diagnostic setting": {
			path:     "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/diagnosticSettings/{name}",
			expected: astmodel.ResourceScopeLocation,
		},
		"diagnostic setting extension": {
			path:     "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}",
			expected: astmodel.ResourceScopeExtension,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			scope := categorizeResourceScope(c.path)
			g.Expect(scope).To(Equal(c.expected))
		})
	}
}
