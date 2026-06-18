// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	v1 "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestSecurityGroupWebhook_ValidateOptionalConfigMapReferences_BothSetFails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	obj := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			Owners: []v1.SecurityGroupMemberReference{{
				ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
				ObjectIDFromConfig: &genruntime.ConfigMapReference{
					Name: "ids",
					Key:  "owner",
				},
			}},
		},
	}

	_, err := (&SecurityGroup_Webhook{}).validateOptionalConfigMapReferences(context.Background(), obj)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("cannot specify both"))
	g.Expect(err.Error()).To(ContainSubstring("ObjectID"))
	g.Expect(err.Error()).To(ContainSubstring("ObjectIDFromConfig"))
}

func TestSecurityGroupWebhook_ValidateOptionalConfigMapReferences_OneOrZeroSetPasses(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ConfigMapReference{Name: "ids", Key: "member"}
	obj := &v1.SecurityGroup{
		Spec: v1.SecurityGroupSpec{
			Owners: []v1.SecurityGroupMemberReference{{
				ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111"),
			}},
			Members: []v1.SecurityGroupMemberReference{
				{ObjectIDFromConfig: &ref},
				{},
			},
		},
	}

	warnings, err := (&SecurityGroup_Webhook{}).validateOptionalConfigMapReferences(context.Background(), obj)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(warnings).To(BeNil())
}

func TestSecurityGroup_Validation_AdmissionSchemaUniqueness(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	kubeClient, namespace, stop := startSecurityGroupValidationEnvtest(t)
	t.Cleanup(stop)

	duplicateOwners := &v1.SecurityGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dup-owners",
			Namespace: namespace,
		},
		Spec: v1.SecurityGroupSpec{
			DisplayName:  to.Ptr("dup-owners"),
			MailNickname: to.Ptr("dup-owners"),
			Owners: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111")},
				{ObjectID: to.Ptr("11111111-1111-1111-1111-111111111111")},
			},
		},
	}
	err := kubeClient.Create(ctx, duplicateOwners)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("owners must be unique"))

	duplicateMembers := &v1.SecurityGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dup-members",
			Namespace: namespace,
		},
		Spec: v1.SecurityGroupSpec{
			DisplayName:  to.Ptr("dup-members"),
			MailNickname: to.Ptr("dup-members"),
			Members: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("33333333-3333-3333-3333-333333333333")},
				{ObjectID: to.Ptr("33333333-3333-3333-3333-333333333333")},
			},
		},
	}
	err = kubeClient.Create(ctx, duplicateMembers)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("members must be unique"))

	overlap := &v1.SecurityGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "owner-member-overlap",
			Namespace: namespace,
		},
		Spec: v1.SecurityGroupSpec{
			DisplayName:  to.Ptr("owner-member-overlap"),
			MailNickname: to.Ptr("owner-member-overlap"),
			Owners: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
			},
			Members: []v1.SecurityGroupMemberReference{
				{ObjectID: to.Ptr("22222222-2222-2222-2222-222222222222")},
			},
		},
	}
	err = kubeClient.Create(ctx, overlap)
	g.Expect(err).ToNot(HaveOccurred())
}

func startSecurityGroupValidationEnvtest(t *testing.T) (client.Client, string, func()) {
	t.Helper()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	g.Expect(v1core.AddToScheme(scheme)).To(Succeed())
	g.Expect(v1.AddToScheme(scheme)).To(Succeed())

	crd := securityGroupValidationTestCRD()

	env := &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			CRDs: []*apiextensions.CustomResourceDefinition{
				&crd,
			},
			Scheme: scheme,
		},
		Scheme: scheme,
	}

	cfg, err := env.Start()
	g.Expect(err).ToNot(HaveOccurred())

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	g.Expect(err).ToNot(HaveOccurred())

	namespace := "securitygroup-validation"
	ns := &v1core.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	err = kubeClient.Create(context.Background(), ns)
	g.Expect(err).ToNot(HaveOccurred())

	stop := func() {
		g.Expect(env.Stop()).To(Succeed())
	}

	return kubeClient, namespace, stop
}

func securityGroupValidationTestCRD() apiextensions.CustomResourceDefinition {
	const objectIDUniquenessRule = "size(self.filter(x, has(x.objectID)).map(x, x.objectID).distinct()) == size(self.filter(x, has(x.objectID)))"
	const objectIDFromConfigUniquenessRule = "size(self.filter(x, has(x.objectIDFromConfig)).map(x, x.objectIDFromConfig).distinct()) == size(self.filter(x, has(x.objectIDFromConfig)))"

	memberReferenceSchema := apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"objectID": {
				Type:      "string",
				MaxLength: to.Ptr[int64](36),
			},
			"objectIDFromConfig": {
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"name": {
						Type:      "string",
						MaxLength: to.Ptr[int64](253),
					},
					"key": {
						Type:      "string",
						MaxLength: to.Ptr[int64](253),
					},
				},
			},
		},
	}

	return apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "securitygroups.entra.azure.com",
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: "entra.azure.com",
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind:     "SecurityGroup",
				ListKind: "SecurityGroupList",
				Singular: "securitygroup",
				Plural:   "securitygroups",
			},
			Scope: apiextensions.NamespaceScoped,
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type: "object",
							Properties: map[string]apiextensions.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]apiextensions.JSONSchemaProps{
										"displayName": {
											Type: "string",
										},
										"mailNickname": {
											Type: "string",
										},
										"owners": {
											Type:     "array",
											MaxItems: to.Ptr[int64](20),
											Items: &apiextensions.JSONSchemaPropsOrArray{
												Schema: &memberReferenceSchema,
											},
											XValidations: apiextensions.ValidationRules{
												{
													Rule:    objectIDUniquenessRule,
													Message: "owners must be unique",
												},
												{
													Rule:    objectIDFromConfigUniquenessRule,
													Message: "owners must be unique",
												},
											},
										},
										"members": {
											Type:     "array",
											MaxItems: to.Ptr[int64](20),
											Items: &apiextensions.JSONSchemaPropsOrArray{
												Schema: &memberReferenceSchema,
											},
											XValidations: apiextensions.ValidationRules{
												{
													Rule:    objectIDUniquenessRule,
													Message: "members must be unique",
												},
												{
													Rule:    objectIDFromConfigUniquenessRule,
													Message: "members must be unique",
												},
											},
										},
									},
									Required: []string{
										"displayName",
										"mailNickname",
									},
								},
							},
							Required: []string{"spec"},
						},
					},
				},
			},
		},
	}
}
