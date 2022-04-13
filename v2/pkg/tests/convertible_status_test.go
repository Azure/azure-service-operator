/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package tests

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestGetVersionedStatus_WorksWhenNoPivotNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := batch.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	account := &batch.BatchAccount{}

	rsrc, err := genruntime.GetVersionedStatus(account, scheme)
	g.Expect(err).To(Succeed())
	g.Expect(rsrc).NotTo(BeNil())
}

//TODO: once we have multiple versions of a resource, we should test that the pivot works too
//func TestGetVersionedStatus_WorksWhenPivotNeeded(t *testing.T) {

func TestNewEmptyVersionedStatus_WorksWhenNoPivotNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := batch.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	account := &batch.BatchAccount{}

	rsrc, err := genruntime.NewEmptyVersionedStatus(account, scheme)
	g.Expect(err).To(Succeed())
	g.Expect(rsrc).NotTo(BeNil())
}

//TODO: once we have multiple versions of a resource, we should test that the pivot works too
//func TestNewEmptyVersionedStatus_WorksWhenPivotNeeded(t *testing.T) {
