// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup_test

import (
	"testing"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
)

func toPartnerInfoSlicePtr(info []sql.PartnerInfo) *[]sql.PartnerInfo {
	return &info
}

func TestDoesResourceMatchAzure(t *testing.T) {

	cases := []struct {
		caseName string
		expected sql.FailoverGroup
		actual   sql.FailoverGroup
		match    bool
	}{
		{
			"differ only by readonly fields matches",
			sql.FailoverGroup{
				Location: to.StringPtr("westus"),
				ID:       to.StringPtr("a"),
				Name:     to.StringPtr("a"),
			},
			sql.FailoverGroup{
				Location: to.StringPtr("eastus"),
				ID:       to.StringPtr("b"),
				Name:     to.StringPtr("b"),
			},
			true,
		},
		{
			"differ only by tag values do not match",
			sql.FailoverGroup{
				Tags: map[string]*string{
					"a": to.StringPtr("b"),
				},
			},
			sql.FailoverGroup{
				Tags: map[string]*string{
					"a": to.StringPtr("c"),
				},
			},
			false,
		},
		{
			"differ only by tag keys do not match",
			sql.FailoverGroup{
				Tags: map[string]*string{
					"a": to.StringPtr("a"),
				},
			},
			sql.FailoverGroup{
				Tags: map[string]*string{
					"b": to.StringPtr("a"),
				},
			},
			false,
		},
		{
			"different tag length do not match",
			sql.FailoverGroup{
				Tags: map[string]*string{
					"a": to.StringPtr("a"),
				},
			},
			sql.FailoverGroup{
				Tags: nil,
			},
			false,
		},
		{
			"differ by failovergroup properties presence do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: nil,
			},
			false,
		},
		{
			// TODO: We don't support ReadOnlyEndpoint which is why we don't consider it -- should we?
			"differ by readonly endpoint do match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadOnlyEndpoint: &sql.FailoverGroupReadOnlyEndpoint{
						FailoverPolicy: sql.ReadOnlyEndpointFailoverPolicyDisabled,
					},
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadOnlyEndpoint: &sql.FailoverGroupReadOnlyEndpoint{
						FailoverPolicy: sql.ReadOnlyEndpointFailoverPolicyEnabled,
					},
				},
			},
			true,
		},
		{
			"differ by readwrite endpoint do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
						FailoverPolicy: sql.Automatic,
					},
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
						FailoverPolicy: sql.Manual,
					},
				},
			},
			false,
		},
		{
			"differ by readwrite endpoint grace period do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
						FailoverWithDataLossGracePeriodMinutes: to.Int32Ptr(10),
					},
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
						FailoverWithDataLossGracePeriodMinutes: to.Int32Ptr(17),
					},
				},
			},
			false,
		},
		{
			"differ by databases do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					Databases: to.StringSlicePtr([]string{
						"a", "b", "c",
					}),
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					Databases: to.StringSlicePtr([]string{
						"a", "b",
					}),
				},
			},
			false,
		},
		{
			"same databases do match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					Databases: to.StringSlicePtr([]string{
						"a", "b", "c",
					}),
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					Databases: to.StringSlicePtr([]string{
						"a", "b", "c",
					}),
				},
			},
			true,
		},
		{
			"differ by partnerServers do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: toPartnerInfoSlicePtr([]sql.PartnerInfo{
						{
							ID: to.StringPtr("a"),
						},
					}),
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: nil,
				},
			},
			false,
		},
		{
			"differ by partnerServers do not match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: toPartnerInfoSlicePtr([]sql.PartnerInfo{
						{
							ID: to.StringPtr("a"),
						},
					}),
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: toPartnerInfoSlicePtr([]sql.PartnerInfo{
						{
							ID: to.StringPtr("b"),
						},
					}),
				},
			},
			false,
		},
		{
			"same partnerServers do match",
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: toPartnerInfoSlicePtr([]sql.PartnerInfo{
						{
							ID: to.StringPtr("a"),
						},
					}),
				},
			},
			sql.FailoverGroup{
				FailoverGroupProperties: &sql.FailoverGroupProperties{
					PartnerServers: toPartnerInfoSlicePtr([]sql.PartnerInfo{
						{
							ID: to.StringPtr("a"),
						},
					}),
				},
			},
			true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.caseName, func(t *testing.T) {
			match := azuresqlfailovergroup.DoesResourceMatchAzure(c.expected, c.actual)
			assert.Equal(t, c.match, match)
		})
	}
}
