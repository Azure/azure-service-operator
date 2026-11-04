// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"testing"

	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	. "github.com/onsi/gomega"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestReconcilerBaseAsResource(t *testing.T) {
	t.Parallel()

	app := &asoentra.Application{}
	cases := map[string]struct {
		obj           genruntime.MetaObject
		expected      *asoentra.Application
		expectedError string
	}{
		"matching type": {
			obj:      app,
			expected: app,
		},
		"mismatched type": {
			obj:           &asoentra.SecurityGroup{},
			expectedError: "cannot modify resource that is not of type *v1.Application. Type is *v1.SecurityGroup",
		},
		"nil interface": {
			expectedError: "cannot modify resource that is not of type *v1.Application. Type is <nil>",
		},
		"typed nil": {
			obj: (*asoentra.Application)(nil),
		},
	}

	base := &reconcilerBase[*asoentra.Application, msgraphmodels.Applicationable]{}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			result, err := base.asResource(c.obj)
			if c.expectedError != "" {
				g.Expect(err).To(MatchError(c.expectedError))
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}

			if c.expected == nil {
				g.Expect(result).To(BeNil())
			} else {
				g.Expect(result).To(BeIdenticalTo(c.expected))
			}
		})
	}
}

func TestReconcilerBaseCreationModes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		spec      *asoentra.ApplicationOperatorSpec
		canAdopt  bool
		canCreate bool
	}{
		"nil spec": {
			canAdopt:  true,
			canCreate: true,
		},
		"default mode": {
			spec:      &asoentra.ApplicationOperatorSpec{},
			canAdopt:  true,
			canCreate: true,
		},
		"AdoptOnly": {
			spec:      &asoentra.ApplicationOperatorSpec{CreationMode: new(asoentra.AdoptOnly)},
			canAdopt:  true,
			canCreate: false,
		},
		"AdoptOrCreate": {
			spec:      &asoentra.ApplicationOperatorSpec{CreationMode: new(asoentra.AdoptOrCreate)},
			canAdopt:  true,
			canCreate: true,
		},
		"AlwaysCreate": {
			spec:      &asoentra.ApplicationOperatorSpec{CreationMode: new(asoentra.AlwaysCreate)},
			canAdopt:  false,
			canCreate: true,
		},
	}

	base := &reconcilerBase[*asoentra.Application, msgraphmodels.Applicationable]{}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			g.Expect(base.canAdopt(c.spec)).To(Equal(c.canAdopt))
			g.Expect(base.canCreate(c.spec)).To(Equal(c.canCreate))
		})
	}
}
