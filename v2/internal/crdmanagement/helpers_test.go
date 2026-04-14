// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/entra"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/generic"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

type schemer struct {
	scheme *runtime.Scheme
}

func (s *schemer) GetScheme() *runtime.Scheme {
	return s.scheme
}

// This test requires that the task target `bundle-crds` has been run
func Test_AllCRDsReady_NoneAreFiltered(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testData := testSetup(t)
	g.Expect(testcommon.CheckBundledCRDsDirectory(testData.crdPath)).To(Succeed())

	// load crds
	goalCRDs, err := testData.crdManager.LoadOperatorCRDs(testData.crdPath, testData.namespace, testData.crdManager.BuildCRDFileFilter("*", nil))
	g.Expect(err).ToNot(HaveOccurred())

	readyResources := crdmanagement.MakeCRDMap(goalCRDs)

	knownTypes, err := testData.getKnownStorageTypes()
	g.Expect(err).ToNot(HaveOccurred())

	// Filter the types to register
	objs, err := crdmanagement.FilterStorageTypesByReadyCRDs(testData.logger, testData.s.GetScheme(), readyResources, knownTypes)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(objs).To(Equal(knownTypes))
}

// This test requires that the task target `bundle-crds` has been run
func Test_FiveCRDsReady_AllOthersAreFiltered(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testData := testSetup(t)
	g.Expect(testcommon.CheckBundledCRDsDirectory(testData.crdPath)).To(Succeed())

	// load crds
	goalCRDs, err := testData.crdManager.LoadOperatorCRDs(testData.crdPath, testData.namespace, testData.crdManager.BuildCRDFileFilter("*", nil))
	g.Expect(err).ToNot(HaveOccurred())

	// Filter all but the first 5 CRDs
	filteredGoalCRDs := make([]apiextensions.CustomResourceDefinition, 0, 5)
	for i, goalCRD := range goalCRDs {
		if i >= 5 {
			break
		}
		filteredGoalCRDs = append(filteredGoalCRDs, goalCRD)
	}

	readyResources := crdmanagement.MakeCRDMap(filteredGoalCRDs)

	knownTypes, err := testData.getKnownStorageTypes()
	g.Expect(err).ToNot(HaveOccurred())

	// Filter the types to register
	objs, err := crdmanagement.FilterStorageTypesByReadyCRDs(testData.logger, testData.s.GetScheme(), readyResources, knownTypes)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(objs).To(HaveLen(5))
}

// This test requires that the task target `bundle-crds` has been run
func Test_ResourceGroupPattern_AllOthersAreFiltered(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testData := testSetup(t)
	g.Expect(testcommon.CheckBundledCRDsDirectory(testData.crdPath)).To(Succeed())

	// load crds
	shouldLoad := testData.crdManager.BuildCRDFileFilter("resources.azure.com/ResourceGroup", nil)
	goalCRDs, err := testData.crdManager.LoadOperatorCRDs(testData.crdPath, testData.namespace, shouldLoad)
	g.Expect(err).ToNot(HaveOccurred())

	readyResources := crdmanagement.MakeCRDMap(goalCRDs)

	knownTypes, err := testData.getKnownStorageTypes()
	g.Expect(err).ToNot(HaveOccurred())

	// Filter the types to register
	objs, err := crdmanagement.FilterStorageTypesByReadyCRDs(testData.logger, testData.s.GetScheme(), readyResources, knownTypes)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(objs).To(HaveLen(1))
}

// This test requires that the task target `bundle-crds` has been run
func Test_AllCRDFilenames_MatchExpectedPattern(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testData := testSetup(t)
	g.Expect(testcommon.CheckBundledCRDsDirectory(testData.crdPath)).To(Succeed())

	entries, err := os.ReadDir(testData.crdPath)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(entries).ToNot(BeEmpty())

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()

		// Every file should have the expected prefix and .yaml suffix
		g.Expect(filename).To(HavePrefix(crdmanagement.CRDFilePrefix), "file %q missing expected prefix", filename)
		g.Expect(filename).To(HaveSuffix(".yaml"), "file %q missing .yaml suffix", filename)
	}
}

/*
 * Helpers
 */

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

type testData struct {
	cfg        config.Values
	s          controllers.Schemer
	kubeClient kubeclient.Client
	logger     logr.Logger
	crdManager *crdmanagement.Manager
	crdPath    string
	namespace  string
}

func testSetup(t *testing.T) *testData {
	asoScheme := api.CreateScheme()
	s := &schemer{scheme: asoScheme}

	kubeClient := NewFakeKubeClient(asoScheme)
	logger := testcommon.NewTestLogger(t)
	cfg := config.Values{}

	crdManager := crdmanagement.NewManager(logger, kubeClient, nil)

	return &testData{
		cfg:        cfg,
		s:          s,
		kubeClient: kubeClient,
		logger:     logger,
		crdManager: crdManager,
		crdPath:    "../../out/crds",
		namespace:  "azureserviceoperator-system",
	}
}

func (t *testData) getKnownStorageTypes() ([]*registration.StorageType, error) {
	clientsProvider := &controllers.ClientsProvider{
		ARMConnectionFactory: func(ctx context.Context, obj genruntime.ARMMetaObject) (arm.Connection, error) {
			return nil, nil
		},
		EntraConnectionFactory: func(ctx context.Context, obj genruntime.EntraMetaObject) (entra.Connection, error) {
			return nil, nil
		},
	}

	return controllers.GetKnownStorageTypes(
		t.s,
		clientsProvider,
		nil, // Not used for this test
		nil, // Not used for this test
		nil, // Not used for this test
		generic.Options{})
}
