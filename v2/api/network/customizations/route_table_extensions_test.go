/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_FuzzySetRoutes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	routeTable := &network.RouteTable_Spec_ARM{
		Location: to.Ptr("westus"),
		Properties: &network.RouteTablePropertiesFormat_ARM{
			DisableBgpRoutePropagation: to.Ptr(true),
		},
	}

	// Note that many of these fields are readonly and will not be present on the PUT
	rawRouteTable := map[string]any{
		"name":     "myroutetable",
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/routeTables/myroutetable",
		"etag":     "W/\"ae3ff66d-e583-4ef9-a83c-a2e00029e2ed\"",
		"type":     "Microsoft.Network/routeTables",
		"location": "westus",
		"properties": map[string]any{
			"provisioningState":          "Succeeded",
			"resourceGuid":               "de9b1f0d-1911-4a5f-9902-74d4c14018d7",
			"disableBgpRoutePropagation": false,
			"routes": []any{
				map[string]any{
					"name": "myroute",
					"id":   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/routeTables/myroutetable/routes/myroute",
					"etag": "W/\"ae3ff66d-e583-4ef9-a83c-a2e00029e2ed\"",
					"properties": map[string]any{
						"provisioningState": "Succeeded",
						"addressPrefix":     "Storage",
						"nextHopType":       "VirtualAppliance",
						"nextHopIpAddress":  "10.0.100.4",
						"hasBgpOverride":    false,
					},
					"type": "Microsoft.Network/routeTables/routes",
				},
			},
		},
	}

	azureRoutes, err := getRawChildCollection(rawRouteTable, "routes")
	g.Expect(err).ToNot(HaveOccurred())

	err = setChildCollection(routeTable, azureRoutes, "Routes")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(routeTable.Location).To(Equal(to.Ptr("westus")))
	g.Expect(routeTable.Properties).ToNot(BeNil())
	g.Expect(routeTable.Properties.DisableBgpRoutePropagation).ToNot(BeNil())
	g.Expect(*routeTable.Properties.DisableBgpRoutePropagation).To(Equal(true))
	g.Expect(routeTable.Properties.Routes).To(HaveLen(1))
	g.Expect(routeTable.Properties.Routes[0].Properties).ToNot(BeNil())
	g.Expect(routeTable.Properties.Routes[0].Name).To(Equal(to.Ptr("myroute")))
	g.Expect(routeTable.Properties.Routes[0].Properties.AddressPrefix).To(Equal(to.Ptr("Storage")))
	g.Expect(routeTable.Properties.Routes[0].Properties.NextHopIpAddress).ToNot(BeNil())
	g.Expect(routeTable.Properties.Routes[0].Properties.NextHopIpAddress).To(Equal(to.Ptr("10.0.100.4")))
}

func Test_FuzzySetRoute(t *testing.T) {
	t.Parallel()

	embeddedType := reflect.TypeOf(network.Route_ARM{})
	properties := gopter.NewProperties(nil)
	arbitraries := arbitrary.DefaultArbitraries()

	properties.Property(
		"all subnet types can be converted between non-embedded and embedded",
		arbitraries.ForAll(
			func(route *network.RouteTables_Route_Spec_ARM) (bool, error) {
				val := reflect.New(embeddedType)
				bytes, err := json.Marshal(route)
				if err != nil {
					return false, err
				}

				raw := make(map[string]any)
				err = json.Unmarshal(bytes, &raw)
				if err != nil {
					return false, err
				}

				err = fuzzySetResource(raw, val)
				return err == nil, err
			}))

	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}
