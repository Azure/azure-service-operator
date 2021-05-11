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
)

func Example_inferNameFromURLPath() {
	group, name, _ := inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}")
	fmt.Printf("%s: %s", group, name)
	// Output: Microsoft.GroupName: ResourceName
}

func Example_inferNameFromURLPath_ChildResources() {
	group, name, _ := inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}/someChild/{childId}")
	fmt.Printf("%s: %s", group, name)
	// Output: Microsoft.GroupName: ResourceNameSomeChild
}

func Test_InferNameFromURLPath_FailsWithMultipleParametersInARow(t *testing.T) {
	g := NewGomegaWithT(t)

	_, _, err := inferNameFromURLPath("/Microsoft.GroupName/resourceName/{resourceId}/{anotherParameter}")
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("multiple parameters"))
}

func Test_InferNameFromURLPath_FailsWithNoGroupName(t *testing.T) {
	g := NewGomegaWithT(t)

	_, _, err := inferNameFromURLPath("/resourceName/{resourceId}/{anotherParameter}")
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("no group name"))
}

func Test_InferNameFromURLPath_SkipsDefault(t *testing.T) {
	g := NewGomegaWithT(t)

	group, name, err := inferNameFromURLPath("Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}")
	g.Expect(err).To(BeNil())
	g.Expect(group).To(Equal("Microsoft.Storage"))
	g.Expect(name).To(Equal("StorageAccountsBlobServicesContainers"))
}

func Test_expandEnumsInPath_ExpandsAnEnum(t *testing.T) {
	g := NewGomegaWithT(t)

	paths := expandEnumsInPath("/some/{value}", []spec.Parameter{
		{
			CommonValidations: spec.CommonValidations{
				Enum: []interface{}{"yes", "no"},
			},
			ParamProps: spec.ParamProps{
				In:       "path",
				Name:     "value",
				Required: true,
			},
		},
	})

	g.Expect(paths).To(ContainElements("/some/yes", "/some/no"))
}

func Test_expandEnumsInPath_ExpandsMultipleEnums(t *testing.T) {
	g := NewGomegaWithT(t)

	paths := expandEnumsInPath("/some/{value1}/{value2}", []spec.Parameter{
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
		{
			CommonValidations: spec.CommonValidations{
				Enum: []interface{}{"orange", "blue"},
			},
			ParamProps: spec.ParamProps{
				In:       "path",
				Name:     "value2",
				Required: true,
			},
		},
	})

	g.Expect(paths).To(ContainElements(
		"/some/yes/blue",
		"/some/no/blue",
		"/some/yes/orange",
		"/some/no/orange",
	))
}
