// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package sqldatabase_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-03-15/documentdb"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/sqldatabase"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
)

func TestResourceDiffer(t *testing.T) {
	testID := "test"
	nonMatchingId := "nonmatching"
	var throughput500 int32 = 500
	var throughput1000 int32 = 1000

	cases := []struct {
		name             string
		expected         documentdb.SQLDatabaseCreateUpdateParameters
		actual           documentdb.SQLDatabaseGetResults
		actualThroughput documentdb.ThroughputSettingsGetResults
		equal            bool
	}{
		{
			name:             "empty types are equal",
			expected:         documentdb.SQLDatabaseCreateUpdateParameters{},
			actual:           documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            true,
		},
		{
			name: "types with equal ids are equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Resource: &documentdb.SQLDatabaseResource{
						ID: &testID,
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				SQLDatabaseGetProperties: &documentdb.SQLDatabaseGetProperties{
					Resource: &documentdb.SQLDatabaseGetPropertiesResource{
						ID: &testID,
					},
				},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            true,
		},
		{
			name: "types with equal empty ids are equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Resource: &documentdb.SQLDatabaseResource{},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				SQLDatabaseGetProperties: &documentdb.SQLDatabaseGetProperties{
					Resource: &documentdb.SQLDatabaseGetPropertiesResource{},
				},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            true,
		},
		{
			name: "types with different ids are not equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Resource: &documentdb.SQLDatabaseResource{
						ID: &testID,
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				SQLDatabaseGetProperties: &documentdb.SQLDatabaseGetProperties{
					Resource: &documentdb.SQLDatabaseGetPropertiesResource{
						ID: &nonMatchingId,
					},
				},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            false,
		},
		{
			name: "expected and actual mismatched throughput is ignored (throughput source of truth is from actualThroughput)",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						Throughput: &throughput500,
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				SQLDatabaseGetProperties: &documentdb.SQLDatabaseGetProperties{},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{
						Throughput: &throughput500,
					},
				},
			},
			equal: true,
		},
		{
			name:     "empty types, partially empty actualThroughput are equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{},
			actual:   documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{},
			},
			equal: true,
		},
		{
			name:     "empty types, partially empty actualThroughput are equal 2",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{},
			actual:   documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{},
				},
			},
			equal: true,
		},
		{
			name:     "empty types, partially empty actualThroughput are equal 3",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{},
			actual:   documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{
						AutoscaleSettings: &documentdb.AutoscaleSettingsResource{},
					},
				},
			},
			equal: true,
		},
		{
			name: "empty throughput equal to nil throughput",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{},
			},
			equal: true,
		},
		{
			name: "throughput not equal to nil throughput",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						Throughput: &throughput500,
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{},
				},
			},
			equal: false,
		},
		{
			name: "throughput not equal to different throughput",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						Throughput: &throughput500,
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{
						Throughput: &throughput1000,
					},
				},
			},
			equal: false,
		},
		{
			name: "empty autoscale equal to nil autoscale",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						AutoscaleSettings: &documentdb.AutoscaleSettings{},
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{},
			},
			equal: true,
		},
		{
			name: "autoscale equal to same autoscale",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						AutoscaleSettings: &documentdb.AutoscaleSettings{
							MaxThroughput: &throughput500,
						},
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{
						AutoscaleSettings: &documentdb.AutoscaleSettingsResource{
							MaxThroughput: &throughput500,
						},
					},
				},
			},
			equal: true,
		},
		{
			name: "autoscale not equal to different autoscale",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						AutoscaleSettings: &documentdb.AutoscaleSettings{
							MaxThroughput: &throughput500,
						},
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{
						AutoscaleSettings: &documentdb.AutoscaleSettingsResource{
							MaxThroughput: &throughput1000,
						},
					},
				},
			},
			equal: false,
		},
		{
			name: "autoscale not equal to nil autoscale",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
					Options: &documentdb.CreateUpdateOptions{
						AutoscaleSettings: &documentdb.AutoscaleSettings{
							MaxThroughput: &throughput500,
						},
					},
				},
			},
			actual: documentdb.SQLDatabaseGetResults{},
			actualThroughput: documentdb.ThroughputSettingsGetResults{
				ThroughputSettingsGetProperties: &documentdb.ThroughputSettingsGetProperties{
					Resource: &documentdb.ThroughputSettingsGetPropertiesResource{},
				},
			},
			equal: false,
		},
		{
			name: "same tags equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				Tags: map[string]*string{
					"a": to.StringPtr("b"),
					"c": to.StringPtr("d"),
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				Tags: map[string]*string{
					"a": to.StringPtr("b"),
					"c": to.StringPtr("d"),
				},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            true,
		},
		{
			name: "different tags not equal",
			expected: documentdb.SQLDatabaseCreateUpdateParameters{
				Tags: map[string]*string{
					"a": to.StringPtr("b"),
					"c": to.StringPtr("b"),
				},
			},
			actual: documentdb.SQLDatabaseGetResults{
				Tags: map[string]*string{
					"a": to.StringPtr("b"),
					"c": to.StringPtr("d"),
				},
			},
			actualThroughput: documentdb.ThroughputSettingsGetResults{},
			equal:            false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			equal := sqldatabase.DoesResourceMatchAzure(c.expected, c.actual, c.actualThroughput)
			g.Expect(equal).To(Equal(c.equal))
		})
	}
}
