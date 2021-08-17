/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package common

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// TODO: State Annotation parameter should be removed once the interface for Status determined and promoted
// TODO: to genruntime. Same for errorAnnotation
type KubeGlobalContext struct {
	TestContext

	useEnvTest bool

	namespace string
}

func (ctx KubeGlobalContext) Namespace() string {
	return ctx.namespace
}

func NewKubeContext(
	useEnvTest bool,
	recordReplay bool,
	namespace string,
	region string) KubeGlobalContext {
	return KubeGlobalContext{
		TestContext: NewTestContext(region, recordReplay),
		useEnvTest:  useEnvTest,
		namespace:   namespace,
	}
}

type KubeBaseTestContext struct {
	PerTestContext

	KubeConfig *rest.Config
}

func AsOwner(obj metav1.ObjectMeta) genruntime.KnownResourceReference {
	return genruntime.KnownResourceReference{
		Name: obj.Name,
	}
}
