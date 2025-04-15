/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestValidateWriteOnceProperties_ReturnsExpectedErrors(t *testing.T) {
	t.Parallel()

	testSub := uuid.New().String()
	resourceGroup := createResourceGroup(
		"rg",
		testSub,
	)

	otherOwner := &genruntime.KnownResourceReference{
		Name: "other-owner",
	}

	cases := map[string]struct {
		modifyOriginal          func(*batch.BatchAccount)
		modifyUpdate            func(*batch.BatchAccount)
		expectedErrorSubstrings []string
	}{
		"WhenNotYetCreated_CanBeModified": {
			modifyOriginal: removeResourceIDAnnotation,
		},
		"WhenNoChange_CanBeModified": {},
		"WhenOriginalHasNoOwner_CannotAddOwner": {
			modifyOriginal: setOwner(nil),
			expectedErrorSubstrings: []string{
				"adding an owner",
				"already created resource",
				"is not allowed",
			},
		},
		"WhenUpdateHasNoOwner_CannotRemoveOwner": {
			modifyUpdate: setOwner(nil),
			expectedErrorSubstrings: []string{
				"removing 'spec.owner'",
				"is not allowed",
			},
		},
		"WhenUpdateHasDifferentOwner_CannotChangeOwner": {
			modifyUpdate: setOwner(otherOwner),
			expectedErrorSubstrings: []string{
				"updating 'spec.owner.name'",
				"is not allowed",
			},
		},
		"WhenUpdateHasDifferentAzureName_CannotChangeAzureName": {
			modifyUpdate: setAzureName("new-name"),
			expectedErrorSubstrings: []string{
				"updating 'spec.azureName'",
				"is not allowed",
			},
		},
		"WhenOriginalHasNoAzureName_CanSetAzureName": {
			modifyOriginal: setAzureName(""),
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			originalAccount := createBatchAccount("acc", resourceGroup)
			updatedAccount := originalAccount.DeepCopy()

			if c.modifyOriginal != nil {
				c.modifyOriginal(originalAccount)
			}

			if c.modifyUpdate != nil {
				c.modifyUpdate(updatedAccount)
			}

			_, err := genruntime.ValidateWriteOnceProperties(originalAccount, updatedAccount)

			if len(c.expectedErrorSubstrings) == 0 {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(Not(BeNil()))
				for _, s := range c.expectedErrorSubstrings {
					g.Expect(err.Error()).To(ContainSubstring(s))
				}
			}
		})
	}
}

func removeResourceIDAnnotation(acc *batch.BatchAccount) {
	delete(acc.Annotations, genruntime.ResourceIDAnnotation)
}

func setAzureName(name string) func(acc *batch.BatchAccount) {
	return func(acc *batch.BatchAccount) {
		acc.Spec.AzureName = name
	}
}

func setOwner(ref *genruntime.KnownResourceReference) func(acc *batch.BatchAccount) {
	return func(acc *batch.BatchAccount) {
		acc.Spec.Owner = ref
	}
}

func createResourceGroup(
	name string,
	namespace string,
) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       resolver.ResourceGroupKind,
			APIVersion: resources.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location:  to.Ptr("West US"),
			AzureName: name,
		},
	}
}

func createBatchAccount(
	name string,
	owner genruntime.ARMMetaObject,
) *batch.BatchAccount {
	return &batch.BatchAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BatchAccount",
			APIVersion: batch.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: owner.GetNamespace(),
			Annotations: map[string]string{
				genruntime.ResourceIDAnnotation: fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", name),
			},
		},
		Spec: batch.BatchAccount_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: owner.GetName(),
			},
			AzureName: name,
		},
	}
}
