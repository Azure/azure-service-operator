// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

func TestLoadByDisplayNamePagination(t *testing.T) {
	t.Parallel()

	for _, resource := range []string{"groups", "applications"} {
		t.Run(resource, func(t *testing.T) {
			t.Parallel()

			cases := map[string]struct {
				pages      [][]string
				fail       bool
				noResponse bool
			}{
				"collects every page": {pages: [][]string{{"first", "second"}, {}, {"third"}}},
				"no matches":          {pages: [][]string{{}}},
				"first page error":    {fail: true},
				"first nil response":  {noResponse: true},
				"later page error":    {pages: [][]string{{"first", "second"}}, fail: true},
				"later nil response":  {pages: [][]string{{"first", "second"}}, noResponse: true},
			}
			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					g := NewWithT(t)
					const displayName = "Team's resources"
					requests := 0
					var wantIDs []string
					for _, page := range c.pages {
						wantIDs = append(wantIDs, page...)
					}
					adapter := &servicePrincipalTestAdapter{}
					adapter.get = func(request *abstractions.RequestInformation) (serialization.Parsable, error) {
						uri, err := request.GetUri()
						g.Expect(err).NotTo(HaveOccurred())
						g.Expect(request.Method).To(Equal(abstractions.GET))
						g.Expect(uri.Path).To(Equal("/beta/" + resource))
						if requests == 0 {
							g.Expect(uri.Query().Get("$filter")).To(Equal("displayName eq 'Team''s resources'"))
						} else {
							g.Expect(uri.String()).To(Equal(fmt.Sprintf("https://graph.microsoft.com/beta/%s?$skiptoken=page%d", resource, requests)))
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
						var next *string
						if requests < len(c.pages) || c.fail || c.noResponse {
							next = new(fmt.Sprintf("https://graph.microsoft.com/beta/%s?$skiptoken=page%d", resource, requests))
						}
						if resource == "groups" {
							page := msgraphmodels.NewGroupCollectionResponse()
							var values []msgraphmodels.Groupable
							for _, id := range c.pages[index] {
								value := msgraphmodels.NewGroup()
								value.SetId(&id)
								values = append(values, value)
							}
							page.SetValue(values)
							page.SetOdataNextLink(next)
							return page, nil
						}
						page := msgraphmodels.NewApplicationCollectionResponse()
						var values []msgraphmodels.Applicationable
						for _, id := range c.pages[index] {
							value := msgraphmodels.NewApplication()
							value.SetId(&id)
							values = append(values, value)
						}
						page.SetValue(values)
						page.SetOdataNextLink(next)
						return page, nil
					}
					graph := msgraphsdk.NewGraphServiceClient(adapter)
					var ids []string
					var err error
					if resource == "groups" {
						var values []msgraphmodels.Groupable
						values, err = (&EntraSecurityGroupReconciler{}).loadGroupsByDisplayName(t.Context(), displayName, graph)
						for _, value := range values {
							ids = append(ids, *value.GetId())
						}
					} else {
						var values []msgraphmodels.Applicationable
						values, err = (&EntraApplicationReconciler{}).loadApplicationsByDisplayName(t.Context(), displayName, graph)
						for _, value := range values {
							ids = append(ids, *value.GetId())
						}
					}
					if c.fail || c.noResponse {
						g.Expect(err).To(HaveOccurred())
						g.Expect(err.Error()).To(ContainSubstring(displayName))
						g.Expect(ids).To(BeNil())
						g.Expect(requests).To(Equal(len(c.pages) + 1))
						return
					}
					g.Expect(err).NotTo(HaveOccurred())
					g.Expect(ids).To(Equal(wantIDs))
					g.Expect(requests).To(Equal(len(c.pages)))
				})
			}
		})
	}
}
