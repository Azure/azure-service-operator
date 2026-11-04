// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	. "github.com/onsi/gomega"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestReconcilerBaseLoadByID(t *testing.T) {
	t.Parallel()

	app := msgraphmodels.NewApplication()
	notFound := makeODataError(http.StatusNotFound, nil)
	forbidden := makeODataError(http.StatusForbidden, nil)
	failure := fmt.Errorf("request failed")
	cases := map[string]struct {
		response      msgraphmodels.Applicationable
		err           error
		expected      msgraphmodels.Applicationable
		expectedError error
	}{
		"found": {
			response: app,
			expected: app,
		},
		"not found": {
			err: notFound,
		},
		"other Graph error": {
			err:           forbidden,
			expectedError: forbidden,
		},
		"other error": {
			err:           failure,
			expectedError: failure,
		},
		"empty response": {},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			ctx := t.Context()
			const id = "application-object-id"
			calls := 0
			fetch := func(fetchContext context.Context, fetchID string) (msgraphmodels.Applicationable, error) {
				g.Expect(fetchContext).To(Equal(ctx))
				g.Expect(fetchID).To(Equal(id))
				calls++
				return c.response, c.err
			}

			base := &reconcilerBase[*asoentra.Application, msgraphmodels.Applicationable]{}
			result, err := base.loadByID(ctx, id, fetch)
			if c.expectedError == nil {
				g.Expect(err).NotTo(HaveOccurred())
			} else {
				g.Expect(err).To(BeIdenticalTo(c.expectedError))
			}

			if c.expected == nil {
				g.Expect(result).To(BeNil())
			} else {
				g.Expect(result).To(BeIdenticalTo(c.expected))
			}

			g.Expect(calls).To(Equal(1))
		})
	}
}

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

func TestReconcilerBaseQuery(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		pages      [][]string
		fail       bool
		noResponse bool
	}{
		"all pages including an empty page": {
			pages: [][]string{
				{"first", "second"},
				{},
				{"third"},
			},
		},
		"no matches": {
			pages: [][]string{
				{},
			},
		},
		"first page error": {
			fail: true,
		},
		"first page missing response": {
			noResponse: true,
		},
		"later page error": {
			pages: [][]string{{"first"}},
			fail:  true,
		},
		"later page missing response": {
			pages:      [][]string{{"first"}},
			noResponse: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			ctx := t.Context()
			const value = "Team's resources"
			requests := 0
			var expectedIDs []string
			for _, page := range c.pages {
				expectedIDs = append(expectedIDs, page...)
			}

			fetchPage := func(
				pageContext context.Context,
				filter string,
				nextLink string,
			) (*msgraphmodels.ApplicationCollectionResponse, error) {
				g.Expect(pageContext).To(Equal(ctx))
				g.Expect(filter).To(Equal("displayName eq 'Team''s resources'"))
				if requests == 0 {
					g.Expect(nextLink).To(BeEmpty())
				} else {
					g.Expect(nextLink).To(Equal(fmt.Sprintf("https://graph.microsoft.com/beta/applications?$skiptoken=page%d", requests)))
				}

				index := requests
				requests++
				if index == len(c.pages) {
					if c.fail {
						return nil, fmt.Errorf("page unavailable")
					}

					if c.noResponse {
						return nil, nil
					}
				}

				g.Expect(index).To(BeNumerically("<", len(c.pages)))
				page := msgraphmodels.NewApplicationCollectionResponse()
				var values []msgraphmodels.Applicationable
				for _, id := range c.pages[index] {
					app := msgraphmodels.NewApplication()
					app.SetId(&id)
					values = append(values, app)
				}

				page.SetValue(values)
				if requests < len(c.pages) || c.fail || c.noResponse {
					page.SetOdataNextLink(new(fmt.Sprintf("https://graph.microsoft.com/beta/applications?$skiptoken=page%d", requests)))
				}

				return page, nil
			}

			readPage := func(
				page *msgraphmodels.ApplicationCollectionResponse,
			) ([]msgraphmodels.Applicationable, *string) {
				return page.GetValue(), page.GetOdataNextLink()
			}

			base := &reconcilerBase[*asoentra.Application, msgraphmodels.Applicationable]{}
			result, err := base.query(ctx, "displayName", value, fetchPage, readPage)
			if c.fail || c.noResponse {
				g.Expect(result).To(BeNil())
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring("*v1.Application"))
				g.Expect(err.Error()).To(ContainSubstring("displayName"))
				g.Expect(err.Error()).To(ContainSubstring(value))
				if c.fail {
					g.Expect(err.Error()).To(ContainSubstring("page unavailable"))
				} else {
					g.Expect(err.Error()).To(ContainSubstring("no response"))
				}

				g.Expect(requests).To(Equal(len(c.pages) + 1))
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			var ids []string
			for _, app := range result {
				ids = append(ids, *app.GetId())
			}

			g.Expect(ids).To(Equal(expectedIDs))
			g.Expect(requests).To(Equal(len(c.pages)))
		})
	}
}
