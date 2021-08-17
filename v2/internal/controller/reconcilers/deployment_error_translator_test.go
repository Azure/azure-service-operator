/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/controller/reconcilers"
)

func Test_RequestErrorWithoutDetails_IsTranslatedCorrectly(t *testing.T) {
	g := NewGomegaWithT(t)

	err := &azure.RequestError{
		ServiceError: &azure.ServiceError{
			Code:    "TestCode",
			Message: "TestMessage",
		},
	}

	translatedError := reconcilers.TranslateAzureErrorToDeploymentError(err)

	g.Expect(translatedError.Code).To(Equal("TestCode"))
	g.Expect(translatedError.Message).To(Equal("TestMessage"))

	g.Expect(translatedError.Details).To(HaveLen(0))
}

func Test_RequestErrorWithDetails_IsTranslatedCorrectly(t *testing.T) {
	g := NewGomegaWithT(t)

	innerCode := "InvalidParameter"
	innerMessage := "Resource 'asotest-vmss-arcvwr' has invalid parameters. Details: The value of parameter imageReference.publisher is invalid."
	innerTarget := "imageReference.publisher"

	err := &azure.RequestError{
		ServiceError: &azure.ServiceError{
			Code:    "TestCode",
			Message: "TestMessage",
			Target:  to.StringPtr("TestTarget"),
			Details: []map[string]interface{}{
				{
					"code":    innerCode,
					"message": innerMessage,
					"target":  innerTarget,
				},
			},
		},
	}

	translatedError := reconcilers.TranslateAzureErrorToDeploymentError(err)

	g.Expect(translatedError.Code).To(Equal("TestCode"))
	g.Expect(translatedError.Message).To(Equal("TestMessage"))
	g.Expect(translatedError.Target).To(Equal("TestTarget"))

	g.Expect(translatedError.Details).To(HaveLen(1))

	detail := translatedError.Details[0]
	g.Expect(detail.Code).To(Equal(innerCode))
	g.Expect(detail.Message).To(Equal(innerMessage))
	g.Expect(detail.Target).To(Equal(innerTarget))
	g.Expect(detail.Details).To(BeNil())
}

func Test_RequestErrorWithEmptyDetails_IsTranslatedCorrectly(t *testing.T) {
	g := NewGomegaWithT(t)

	err := &azure.RequestError{
		ServiceError: &azure.ServiceError{
			Code:    "TestCode",
			Message: "TestMessage",
			Target:  to.StringPtr("TestTarget"),
			Details: []map[string]interface{}{
				{},
			},
		},
	}

	translatedError := reconcilers.TranslateAzureErrorToDeploymentError(err)

	g.Expect(translatedError.Code).To(Equal("TestCode"))
	g.Expect(translatedError.Message).To(Equal("TestMessage"))
	g.Expect(translatedError.Target).To(Equal("TestTarget"))

	g.Expect(translatedError.Details).To(HaveLen(1))

	detail := translatedError.Details[0]
	g.Expect(detail.Code).To(Equal(""))
	g.Expect(detail.Message).To(Equal(""))
	g.Expect(detail.Target).To(Equal(""))
	g.Expect(detail.Details).To(BeNil())
}

func Test_RequestErrorWithUnexpectedDetailsType_IsTranslatedCorrectly(t *testing.T) {
	g := NewGomegaWithT(t)

	innerMessage := "This message should work"

	err := &azure.RequestError{
		ServiceError: &azure.ServiceError{
			Code:    "TestCode",
			Message: "TestMessage",
			Details: []map[string]interface{}{
				{
					"code":       7,
					"message":    innerMessage,
					"otherField": struct{}{},
					"target":     struct{}{},
				},
			},
		},
	}

	translatedError := reconcilers.TranslateAzureErrorToDeploymentError(err)

	g.Expect(translatedError.Code).To(Equal("TestCode"))
	g.Expect(translatedError.Message).To(Equal("TestMessage"))
	g.Expect(translatedError.Target).To(Equal(""))

	g.Expect(translatedError.Details).To(HaveLen(1))

	detail := translatedError.Details[0]
	g.Expect(detail.Code).To(Equal(""))
	g.Expect(detail.Message).To(Equal(innerMessage))
	g.Expect(detail.Target).To(Equal(""))
	g.Expect(detail.Details).To(BeNil())
}
