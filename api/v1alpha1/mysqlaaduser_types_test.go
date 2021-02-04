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

		It("can't downgrade with roles for multiple databases", func() {
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
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(MatchError("can't convert user \"foo\" to *v1alpha1.MySQLAADUser because it has privileges in 2 databases"))
		})

		It("can't downgrade with roles for no databases", func() {
			v2 := v1alpha2.MySQLAADUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1alpha2.MySQLAADUserSpec{
					Server:        "myserver",
					ResourceGroup: "foo-group",
					DatabaseRoles: nil,
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(MatchError("can't convert user \"foo\" to *v1alpha1.MySQLAADUser because it has privileges in 0 databases"))
		})

		It("can't downgrade with server roles", func() {
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
				},
			}
			var v1 MySQLAADUser
			Expect(v1.ConvertFrom(&v2)).To(MatchError("can't convert user \"foo\" to *v1alpha1.MySQLAADUser because it has server-level roles"))
		})

	})

})
