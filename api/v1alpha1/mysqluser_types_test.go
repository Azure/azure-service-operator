// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

var _ = Describe("MySQLUser", func() {
	Context("Conversion", func() {
		It("can upgrade to v1alpha2", func() {

			v1 := &MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: MySQLUserSpec{
					Server:                 "myserver",
					DbName:                 "mydb",
					ResourceGroup:          "foo-group",
					Roles:                  []string{"role1", "role2", "role3"},
					AdminSecret:            "adminsecret",
					AdminSecretKeyVault:    "adminsecretkeyvault",
					Username:               "username",
					KeyVaultToStoreSecrets: "keyvaulttostoresecrets",
				},
			}

			var v2 v1alpha2.MySQLUser
			Expect(v1.ConvertTo(&v2)).To(Succeed())
			Expect(v2).To(Equal(v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{},
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
					AdminSecret:            "adminsecret",
					AdminSecretKeyVault:    "adminsecretkeyvault",
					Username:               "username",
					KeyVaultToStoreSecrets: "keyvaulttostoresecrets",
				},
			}))
		})

		It("can downgrade to v1alpha1 if it's the right shape", func() {
			v2 := v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         nil,
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
					AdminSecret:            "adminsecret",
					AdminSecretKeyVault:    "adminsecretkeyvault",
					Username:               "username",
					KeyVaultToStoreSecrets: "keyvaulttostoresecrets",
				},
			}
			var v1 MySQLUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: MySQLUserSpec{
					Server:                 "myserver",
					DbName:                 "mydb",
					ResourceGroup:          "foo-group",
					Roles:                  []string{"role1", "role2", "role3"},
					AdminSecret:            "adminsecret",
					AdminSecretKeyVault:    "adminsecretkeyvault",
					Username:               "username",
					KeyVaultToStoreSecrets: "keyvaulttostoresecrets",
				},
			}))
		})

		It("can downgrade with roles for multiple databases", func() {
			v2 := v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DatabaseRoles: map[string][]string{
						"mydb":    {"role1", "role2", "role3"},
						"otherdb": {"otherrole"},
					},
				},
			}
			var v1 MySQLUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					// Might need to account for ordering here -
					// compare deserialised instead.
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"],"otherdb":["otherrole"]}}`,
					},
				},
				Spec: MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DbName:        "mydb",
					Roles:         []string{"role1", "role2", "role3"},
				},
			}))
		})

		It("can downgrade with roles for no databases", func() {
			v2 := v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DatabaseRoles: nil,
				},
			}
			var v1 MySQLUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{}`,
					},
				},
				Spec: MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DbName:        "",
					Roles:         nil,
				},
			}))
		})

		It("can downgrade with server roles", func() {
			v2 := v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{"somekindofsuperuser"},
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
				},
			}
			var v1 MySQLUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"]},"roles":["somekindofsuperuser"]}`,
					},
				},
				Spec: MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DbName:        "mydb",
					Roles:         []string{"role1", "role2", "role3"},
				},
			}))
		})

		It("can incorporate annotations when upgrading", func() {
			// Server-level roles should just be carried forwards.
			// Database roles need to be applied to the ones in the annotation.
			v1 := MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"],"otherdb":["anotherrole"]},"roles":["somekindofsuperuser"]}`,
					},
				},
				Spec: MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DbName:        "mydb",
					Roles:         []string{"role3", "role4"},
				},
			}
			var v2 v1alpha2.MySQLUser
			Expect(v1.ConvertTo(&v2)).To(Succeed())
			Expect(v2).To(Equal(v1alpha2.MySQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{"somekindofsuperuser"},
					DatabaseRoles: map[string][]string{
						"mydb":    {"role3", "role4"},
						"otherdb": {"anotherrole"},
					},
				},
			}))
		})

	})

})
