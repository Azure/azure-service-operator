/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	storage "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.storage/v20190401"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

// If you modify this make sure to modify the cleanup-test-azure-resources target in the Makefile too
const ResourcePrefix = "k8sinfratest"

type TestContext struct {
	Namer        *ResourceNamer
	KubeClient   client.Client
	Ensure       *Ensure
	MatcherMaker *MatcherMaker
	AzureClient  armclient.Applier

	AzureRegion       string
	Namespace         string
	AzureSubscription string
}

// TODO: State Annotation parameter should be removed once the interface for Status determined and promoted
// TODO: to genruntime. Same for errorAnnotation
func NewTestContext(region string, namespace string, stateAnnotation string, errorAnnotation string) (*TestContext, error) {
	scheme := CreateScheme()
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get kubeconfig")
	}

	kubeClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, errors.Wrapf(err, "creating kubeclient")
	}

	armClient, err := armclient.NewAzureTemplateClient()
	if err != nil {
		return nil, errors.Wrapf(err, "creating armclient")
	}

	subscription, ok := os.LookupEnv("AZURE_SUBSCRIPTION_ID")
	if !ok {
		return nil, errors.New("couldn't find AZURE_SUBSCRIPTION_ID")
	}

	ensure := NewEnsure(kubeClient, stateAnnotation, errorAnnotation)

	return &TestContext{
		AzureClient:       armClient,
		AzureRegion:       region,
		AzureSubscription: subscription, // TODO: Do we really need this?
		Ensure:            ensure,
		MatcherMaker:      NewMatcherMaker(ensure),
		KubeClient:        kubeClient,
		Namer:             NewResourceNamer(ResourcePrefix, "-", 6),
		Namespace:         namespace,
	}, nil
}

func (tc *TestContext) CreateTestNamespace() error {
	ctx := context.Background()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.Namespace,
		},
	}
	_, err := controllerutil.CreateOrUpdate(ctx, tc.KubeClient, ns, func() error {
		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "creating namespace")
	}

	return nil
}

func (tc *TestContext) MakeObjectMeta(prefix string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      tc.Namer.GenerateName(prefix),
		Namespace: tc.Namespace,
	}
}

func (tc *TestContext) MakeObjectMetaWithName(name string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      name,
		Namespace: tc.Namespace,
	}
}

func (tc *TestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: tc.MakeObjectMeta("rg"),
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			// This tag is used for cleanup optimization
			Tags: map[string]string{"CreatedAt": time.Now().UTC().Format(time.RFC3339)},
		},
	}
}

// TODO: Code generate this (I think I already do in another branch)
func CreateScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()

	_ = clientgoscheme.AddToScheme(scheme)
	//_ = batch.AddToScheme(scheme)
	_ = storage.AddToScheme(scheme)
	_ = resources.AddToScheme(scheme)

	return scheme
}
