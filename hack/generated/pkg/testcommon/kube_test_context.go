/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	batch "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v20170901"
	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v20200601"
	storage "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.storage/v20190401"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
)

// TODO: State Annotation parameter should be removed once the interface for Status determined and promoted
// TODO: to genruntime. Same for errorAnnotation
type KubeGlobalContext struct {
	TestContext

	useEnvTest bool

	namespace       string
	stateAnnotation string
	errorAnnotation string
}

func (ctx KubeGlobalContext) Namespace() string {
	return ctx.namespace
}

func NewKubeContext(
	useEnvTest bool,
	recordReplay bool,
	namespace string,
	region string,
	stateAnnotation string,
	errorAnnotation string) KubeGlobalContext {
	return KubeGlobalContext{
		TestContext:     NewTestContext(region, recordReplay),
		useEnvTest:      useEnvTest,
		namespace:       namespace,
		stateAnnotation: stateAnnotation,
		errorAnnotation: errorAnnotation,
	}
}

func (ctx KubeGlobalContext) ForTest(t *testing.T) (KubePerTestContext, error) {
	perTestContext, err := ctx.TestContext.ForTest(t)
	if err != nil {
		return KubePerTestContext{}, err
	}

	var baseCtx *KubeBaseTestContext
	if ctx.useEnvTest {
		baseCtx, err = createEnvtestContext(perTestContext)
	} else {
		baseCtx, err = createRealKubeContext(perTestContext)
	}

	if err != nil {
		return KubePerTestContext{}, err
	}

	clientOptions := client.Options{Scheme: CreateScheme()}
	kubeClient, err := client.New(baseCtx.KubeConfig, clientOptions)
	if err != nil {
		return KubePerTestContext{}, err
	}

	ensure := NewEnsure(
		kubeClient,
		ctx.stateAnnotation,
		ctx.errorAnnotation)

	match := NewKubeMatcher(ensure)

	result := KubePerTestContext{
		KubeGlobalContext:   &ctx,
		KubeBaseTestContext: *baseCtx,
		KubeClient:          kubeClient,
		Ensure:              ensure,
		Match:               match,
	}

	err = result.createTestNamespace()
	if err != nil {
		return KubePerTestContext{}, err
	}

	return result, nil
}

type KubeBaseTestContext struct {
	PerTestContext

	KubeConfig *rest.Config
}

type KubePerTestContext struct {
	*KubeGlobalContext
	KubeBaseTestContext

	KubeClient client.Client
	Ensure     *Ensure
	Match      *KubeMatcher
}

func (tc KubePerTestContext) createTestNamespace() error {
	ctx := context.Background()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.namespace,
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

func (tc KubePerTestContext) MakeObjectMeta(prefix string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      tc.Namer.GenerateName(prefix),
		Namespace: tc.namespace,
	}
}

func (tc KubePerTestContext) MakeObjectMetaWithName(name string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      name,
		Namespace: tc.namespace,
	}
}

func (tc KubePerTestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: tc.MakeObjectMeta("rg"),
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			// This tag is used for cleanup optimization
			Tags: CreateTestResourceGroupDefaultTags(),
		},
	}
}

func CreateTestResourceGroupDefaultTags() map[string]string {
	return map[string]string{"CreatedAt": time.Now().UTC().Format(time.RFC3339)}
}

// TODO: Code generate this (I think I already do in another branch)
func CreateScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()

	_ = clientgoscheme.AddToScheme(scheme)
	_ = batch.AddToScheme(scheme)
	_ = storage.AddToScheme(scheme)
	_ = resources.AddToScheme(scheme)

	return scheme
}

type WaitCondition bool

const (
	WaitForCreation WaitCondition = true
	DoNotWait       WaitCondition = false
)

// CreateNewTestResourceGroup creates a new randomly-named resource group
// and registers it to be deleted up when the context is cleaned up
func (tc KubePerTestContext) CreateNewTestResourceGroup(wait WaitCondition) (*resources.ResourceGroup, error) {
	ctx := context.Background()

	rg := tc.NewTestResourceGroup()

	log.Printf("Creating test resource group %q", rg.Name)
	err := tc.KubeClient.Create(ctx, rg)
	if err != nil {
		return nil, errors.Wrapf(err, "creating resource group")
	}

	// register the RG for cleanup
	// important to do this before waiting for it, so that
	// we delete it even if we time out
	tc.T.Cleanup(func() {
		ctx := context.Background()
		log.Printf("Deleting test resource group %q", rg.Name)
		err := tc.KubeClient.Delete(ctx, rg)
		if err != nil {
			// don't error out, just warn
			log.Printf("Unable to delete resource group: %s", err.Error())
		}
	})

	if wait {
		err = WaitFor(ctx, 2*time.Minute, func(ctx context.Context) (bool, error) {
			return tc.Ensure.Provisioned(ctx, rg)
		})

		if err != nil {
			return nil, errors.Wrapf(err, "waiting for resource group creation")
		}
	}

	return rg, nil
}

func AsOwner(obj metav1.ObjectMeta) genruntime.KnownResourceReference {
	return genruntime.KnownResourceReference{
		Name: obj.Name,
	}
}
