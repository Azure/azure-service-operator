/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package tests

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	batchstorage "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestObjAsOriginalVersion_WorksWhenNoPivotNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	err := batch.AddToScheme(scheme)
	g.Expect(err).To(Succeed())
	err = batchstorage.AddToScheme(scheme)
	g.Expect(err).To(Succeed())

	account := &batch.BatchAccount{}

	rsrc, err := genruntime.ObjAsOriginalVersion(account, scheme)
	g.Expect(err).To(Succeed())
	g.Expect(rsrc).NotTo(BeNil())
}
