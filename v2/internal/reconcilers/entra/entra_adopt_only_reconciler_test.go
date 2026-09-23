// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"context"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

func TestApplicationAdoptOnlyMissingReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	mode := asoentra.AdoptOnly
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(uri.Path).To(gomega.Equal("/beta/applications"))
		return msgraphmodels.NewApplicationCollectionResponse(), nil
	}
	reconciler := &EntraApplicationReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.Application{
		ObjectMeta: metav1.ObjectMeta{Name: "missing-application"},
		Spec: asoentra.ApplicationSpec{
			DisplayName:  stringPtr("missing application"),
			OperatorSpec: &asoentra.ApplicationOperatorSpec{CreationMode: &mode},
		},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).To(gomega.MatchError(gomega.ContainSubstring("application missing-application not found for adoption")))
}

func TestSecurityGroupAdoptOnlyMissingReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	mode := asoentra.AdoptOnly
	adapter := &servicePrincipalTestAdapter{}
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(uri.Path).To(gomega.Equal("/beta/groups"))
		return msgraphmodels.NewGroupCollectionResponse(), nil
	}
	reconciler := &EntraSecurityGroupReconciler{
		EntraClientFactory: servicePrincipalTestFactory(adapter),
	}
	obj := &asoentra.SecurityGroup{
		ObjectMeta: metav1.ObjectMeta{Name: "missing-group"},
		Spec: asoentra.SecurityGroupSpec{
			DisplayName:  stringPtr("missing group"),
			MailNickname: stringPtr("missing-group"),
			OperatorSpec: &asoentra.SecurityGroupOperatorSpec{CreationMode: &mode},
		},
	}

	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).To(gomega.MatchError(gomega.ContainSubstring("security group missing-group not found for adoption")))
}

func TestApplicationAdoptOnlyDoesNotModifyOrDelete(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const id = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	adapter := &servicePrincipalTestAdapter{}
	calls := 0
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		calls++
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		result := msgraphmodels.NewApplication()
		result.SetId(stringPtr(id))
		return result, nil
	}
	reconciler := &EntraApplicationReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	mode := asoentra.AdoptOnly
	obj := &asoentra.Application{
		Spec: asoentra.ApplicationSpec{
			DisplayName:  stringPtr("unchanged"),
			OperatorSpec: &asoentra.ApplicationOperatorSpec{CreationMode: &mode},
		},
	}
	setEntraID(obj, id)
	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(1))
}

func TestSecurityGroupAdoptOnlyDoesNotModifyOrDelete(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)
	const id = "58f30b77-0736-4ef3-8d0c-51e78c1d42b7"
	adapter := &servicePrincipalTestAdapter{}
	calls := 0
	adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
		calls++
		g.Expect(request.Method).To(gomega.Equal(abstractions.GET))
		uri, err := request.GetUri()
		g.Expect(err).NotTo(gomega.HaveOccurred())
		if strings.HasSuffix(uri.Path, "/owners") || strings.HasSuffix(uri.Path, "/members") {
			return msgraphmodels.NewDirectoryObjectCollectionResponse(), nil
		}
		result := msgraphmodels.NewGroup()
		result.SetId(stringPtr(id))
		return result, nil
	}
	reconciler := &EntraSecurityGroupReconciler{EntraClientFactory: servicePrincipalTestFactory(adapter)}
	mode := asoentra.AdoptOnly
	obj := &asoentra.SecurityGroup{
		Spec: asoentra.SecurityGroupSpec{
			DisplayName:  stringPtr("unchanged"),
			MailNickname: stringPtr("unchanged"),
			OperatorSpec: &asoentra.SecurityGroupOperatorSpec{CreationMode: &mode},
		},
	}
	setEntraID(obj, id)
	_, err := reconciler.CreateOrUpdate(context.Background(), logr.Discard(), nil, obj, annotations.ResolvedReconcilePolicies{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	_, err = reconciler.Delete(context.Background(), logr.Discard(), nil, obj)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(calls).To(gomega.Equal(3))
}
