// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

var _ = Describe("MySQLAADUser", func() {
	Context("Conversion", func() {
		It("can upgrade to v1alpha2", func() {

			v1 := &MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					DBName:        "mydb",
					ResourceGroup: "foo-group",
					Roles:         []string{"role1", "role2", "role3"},
					AADID:         "aadid",
					Username:      "username",
				},
			}

			var v2 v1alpha2.MySQLAADUser
			Expect(v1.ConvertTo(&v2)).To(Succeed())
			Expect(v2).To(Equal(v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{},
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
					AADID:    "aadid",
					Username: "username",
				},
			}))
		})

		It("can downgrade to v1alpha1 if it's the right shape", func() {
			v2 := v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         nil,
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
					AADID:    "aadid",
					Username: "username",
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					DBName:        "mydb",
					ResourceGroup: "foo-group",
					Roles:         []string{"role1", "role2", "role3"},
					AADID:         "aadid",
					Username:      "username",
				},
			}))
		})

		It("can downgrade with roles for multiple databases", func() {
			v2 := v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DatabaseRoles: map[string][]string{
						"mydb":    {"role1", "role2", "role3"},
						"otherdb": {"otherrole"},
					},
					AADID:    "aadid",
					Username: "username",
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					// Might need to account for ordering here -
					// compare deserialised instead.
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"],"otherdb":["otherrole"]}}`,
					},
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DBName:        "mydb",
					Roles:         []string{"role1", "role2", "role3"},
					AADID:         "aadid",
					Username:      "username",
				},
			}))
		})

		It("can downgrade with roles for no databases", func() {
			v2 := v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DatabaseRoles: nil,
					AADID:         "aadid",
					Username:      "username",
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{}`,
					},
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DBName:        "",
					Roles:         nil,
					AADID:         "aadid",
					Username:      "username",
				},
			}))
		})

		It("can downgrade with server roles", func() {
			v2 := v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{"somekindofsuperuser"},
					DatabaseRoles: map[string][]string{
						"mydb": {"role1", "role2", "role3"},
					},
					AADID:    "aadid",
					Username: "username",
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(Succeed())
			Expect(v1).To(Equal(MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"]},"roles":["somekindofsuperuser"]}`,
					},
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DBName:        "mydb",
					Roles:         []string{"role1", "role2", "role3"},
					AADID:         "aadid",
					Username:      "username",
				},
			}))
		})

		It("can incorporate annotations when upgrading", func() {
			// Server-level roles should just be carried forwards.
			// Database roles need to be applied to the ones in the annotation.
			v1 := MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
					Annotations: map[string]string{
						"azure.microsoft.com/convert-stash": `{"databaseRoles":{"mydb":["role1","role2","role3"],"otherdb":["anotherrole"]},"roles":["somekindofsuperuser"]}`,
					},
				},
				Spec: MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DBName:        "mydb",
					Roles:         []string{"role3", "role4"},
					AADID:         "aadid",
					Username:      "username",
				},
			}
			var v2 v1alpha2.MySQLAADUser
			Expect(v1.ConvertTo(&v2)).To(Succeed())
			Expect(v2).To(Equal(v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					Roles:         []string{"somekindofsuperuser"},
					DatabaseRoles: map[string][]string{
						"mydb":    {"role3", "role4"},
						"otherdb": {"anotherrole"},
					},
					AADID:    "aadid",
					Username: "username",
				},
			}))
		})

	})

})
