/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_DestinationFailure_ClearsLROForNormalRecovery(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		deleting bool
	}{
		"create": {
			deleting: false,
		},
		"delete": {
			deleting: true,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			scheme := runtime.NewScheme()
			g.Expect(resources.AddToScheme(scheme)).To(Succeed())
			obj := &resources.ResourceGroup{}
			obj.SetGroupVersionKind(resources.GroupVersion.WithKind("ResourceGroup"))
			SetPollerResumeToken(obj, genericarmclient.CreatePollerID, "saved-operation")
			r := &azureDeploymentReconcilerInstance{
				ARMOwnedResourceReconcilerCommon: reconcilers.ARMOwnedResourceReconcilerCommon{
					ResourceResolver: resolver.NewResolver(NewFakeKubeClient(scheme)),
				},
				Obj: obj,
				Log: logr.Discard(),
			}
			allowed, err := httpclient.NewOrigin("https://management.example.invalid")
			g.Expect(err).ToNot(HaveOccurred())
			target, err := url.Parse("https://attacker.example.invalid/path?secret=private")
			g.Expect(err).ToNot(HaveOccurred())
			err = allowed.Validate(target)
			g.Expect(err).To(HaveOccurred())
			if c.deleting {
				err = r.handleDeleteFailed(err)
			} else {
				err = r.handleCreateOrUpdateFailed(err)
			}
			conditionError, ok := conditions.AsReadyConditionImpactingError(err)
			g.Expect(ok).To(BeTrue())
			g.Expect(conditionError.Severity).To(Equal(conditions.ConditionSeverityWarning))
			g.Expect(err).ToNot(MatchError(ContainSubstring("private")))
			_, _, hasToken := GetPollerResumeToken(obj)
			g.Expect(hasToken).To(BeFalse())
			g.Expect(obj.GetAnnotations()).ToNot(HaveKey(reconcilers.PollerResumeTokenAnnotation))
			g.Expect(obj.GetAnnotations()).ToNot(HaveKey(reconcilers.PollerResumeIDAnnotation))
		})
	}
}
