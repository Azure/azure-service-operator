/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20250101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SQL Managed Instance requires a delegated subnet (with its own network security group and route table) and takes
// several hours to provision when it is the first instance in an empty subnet (4-6 hours observed).
// See https://learn.microsoft.com/azure/azure-sql/managed-instance/connectivity-architecture-overview
func Test_SQL_ManagedInstance_v20250101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr("westus2")

	secretName := "sqlmisecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	rg := tc.CreateTestResourceGroupAndWait()

	_, subnet := newManagedInstanceNetwork(tc, rg, "mi", "10.0.0.0/16", "10.0.0.0/26")

	configMapName := "miconfig"
	configMapKey := "fqdn"

	mi := &sql.ManagedInstance{
		ObjectMeta: tc.MakeObjectMeta("sqlmi"),
		Spec: sql.ManagedInstance_Spec{
			Location:                   tc.AzureRegion,
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.Ptr("miadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			SubnetReference:            tc.MakeReferenceFromResource(subnet),
			Sku: &sql.Sku{
				Name: to.Ptr("GP_Gen5"),
				Tier: to.Ptr("GeneralPurpose"),
			},
			VCores:                           to.Ptr(4),
			StorageSizeInGB:                  to.Ptr(32),
			LicenseType:                      to.Ptr(sql.ManagedInstanceLicenseType_LicenseIncluded),
			PublicDataEndpointEnabled:        to.Ptr(false),
			RequestedBackupStorageRedundancy: to.Ptr(sql.BackupStorageRedundancy_Geo),
			OperatorSpec: &sql.ManagedInstanceOperatorSpec{
				ConfigMaps: &sql.ManagedInstanceOperatorConfigMaps{
					FullyQualifiedDomainName: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  configMapKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)

	tc.Expect(mi.Status.Id).ToNot(BeNil())
	tc.Expect(mi.Status.FullyQualifiedDomainName).ToNot(BeNil())
	tc.Expect(mi.Status.RequestedBackupStorageRedundancy).ToNot(BeNil())
	tc.Expect(string(*mi.Status.RequestedBackupStorageRedundancy)).To(Equal("Geo"))
	tc.ExpectConfigMapHasKeysAndValues(
		configMapName,
		configMapKey,
		*mi.Status.FullyQualifiedDomainName,
	)

	armId := *mi.Status.Id

	// Update the instance - changing the tags exercises PUT without triggering a long running scale operation
	old := mi.DeepCopy()
	mi.Spec.Tags = map[string]string{"cheese": "blue"}
	tc.PatchResourceAndWait(old, mi)
	tc.Expect(mi.Status.Tags).To(HaveKeyWithValue("cheese", "blue"))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL ManagedInstance Database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_ManagedInstance_Database_v20250101_CRUD(tc, mi)
			},
		},
	)

	tc.DeleteResourceAndWait(mi)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func SQL_ManagedInstance_Database_v20250101_CRUD(tc *testcommon.KubePerTestContext, mi *sql.ManagedInstance) {
	db := &sql.ManagedInstancesDatabase{
		ObjectMeta: tc.MakeObjectMeta("midb"),
		Spec: sql.ManagedInstancesDatabase_Spec{
			Owner:     testcommon.AsOwner(mi),
			Location:  tc.AzureRegion,
			Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		},
	}

	tc.CreateResourceAndWait(db)
	tc.Expect(db.Status.Id).ToNot(BeNil())
	armId := *db.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL ManagedInstance Database TransparentDataEncryption CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				// TDE can't be created or deleted, only updated on an existing database (it's named "current")
				tde := &sql.ManagedInstancesDatabasesTransparentDataEncryption{
					ObjectMeta: tc.MakeObjectMeta("tde"),
					Spec: sql.ManagedInstancesDatabasesTransparentDataEncryption_Spec{
						Owner: testcommon.AsOwner(db),
						State: to.Ptr(sql.TransparentDataEncryptionState_Enabled),
					},
				}

				tc.CreateResourceAndWait(tde)
				tc.Expect(tde.Status.Id).ToNot(BeNil())
				tc.Expect(tde.Status.State).ToNot(BeNil())
			},
		},
	)

	tc.DeleteResourceAndWait(db)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// An instance failover group links a primary and a secondary managed instance in different regions.
// Requirements, from https://learn.microsoft.com/azure/azure-sql/managed-instance/failover-group-configure-sql-mi:
//   - The two instances must be in different regions, with non-overlapping address ranges.
//   - The virtual networks must be connected (here, with global peering in both directions).
//   - The secondary instance must be created in the same DNS zone as the primary (dnsZonePartner).
//
// Note that this test is expensive: it creates two managed instances, each taking several hours to provision
// (4-6 hours observed), and the second cannot start until the first is ready.
func Test_SQL_InstanceFailoverGroup_v20250101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	primaryRegion := to.Ptr("westus2")
	secondaryRegion := to.Ptr("eastus2")
	tc.AzureRegion = primaryRegion

	secretName := "sqlmisecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	rg := tc.CreateTestResourceGroupAndWait()

	primaryVnet, primarySubnet := newManagedInstanceNetwork(tc, rg, "primary", "10.1.0.0/16", "10.1.0.0/26")
	secondaryVnet, secondarySubnet := newManagedInstanceNetworkInRegion(tc, rg, secondaryRegion, "secondary", "10.2.0.0/16", "10.2.0.0/26")

	// Connect the two networks in both directions
	peerings := []*network.VirtualNetworksVirtualNetworkPeering{
		newManagedInstancePeering(tc, "primarytosecondary", primaryVnet, secondaryVnet),
		newManagedInstancePeering(tc, "secondarytoprimary", secondaryVnet, primaryVnet),
	}
	tc.CreateResourceAndWait(peerings[0])
	tc.CreateResourceAndWait(peerings[1])

	newInstance := func(name string, region *string, subnet *network.VirtualNetworksSubnet) *sql.ManagedInstance {
		return &sql.ManagedInstance{
			ObjectMeta: tc.MakeObjectMeta(name),
			Spec: sql.ManagedInstance_Spec{
				Location:                   region,
				Owner:                      testcommon.AsOwner(rg),
				AdministratorLogin:         to.Ptr("miadmin"),
				AdministratorLoginPassword: &adminPasswordSecretRef,
				SubnetReference:            tc.MakeReferenceFromResource(subnet),
				Sku: &sql.Sku{
					Name: to.Ptr("GP_Gen5"),
					Tier: to.Ptr("GeneralPurpose"),
				},
				VCores:                    to.Ptr(4),
				StorageSizeInGB:           to.Ptr(32),
				LicenseType:               to.Ptr(sql.ManagedInstanceLicenseType_LicenseIncluded),
				PublicDataEndpointEnabled: to.Ptr(false),
			},
		}
	}

	primary := newInstance("primarymi", primaryRegion, primarySubnet)
	tc.CreateResourceAndWait(primary)

	// The secondary must join the primary's DNS zone
	secondary := newInstance("secondarymi", secondaryRegion, secondarySubnet)
	secondary.Spec.DnsZonePartnerReference = tc.MakeReferenceFromResource(primary)
	tc.CreateResourceAndWait(secondary)

	// An instance failover group has the ARM type Microsoft.Sql/locations/instanceFailoverGroups, so its Azure ID needs
	// both the location and the group name. The generated resource has a single AzureName, and a resource group owner
	// leaves ASO with one name for two type segments ("had 1 azureNames and 2 resourceTypes"). Giving the owner as the
	// ARM ID of the location supplies the "locations/<region>" part, and AzureName is then just the group name.
	failoverGroupOwnerARMID := fmt.Sprintf("%s/providers/Microsoft.Sql/locations/%s", *rg.Status.Id, *primaryRegion)
	failoverGroup := &sql.InstanceFailoverGroup{
		ObjectMeta: tc.MakeObjectMeta("fog"),
		Spec: sql.InstanceFailoverGroup_Spec{
			Owner:         &genruntime.KnownResourceReference{ARMID: failoverGroupOwnerARMID},
			AzureName:     tc.Namer.GenerateName("fog"),
			SecondaryType: to.Ptr(sql.SecondaryInstanceType_Geo),
			PartnerRegions: []sql.PartnerRegionInfo{
				{Location: secondaryRegion},
			},
			ManagedInstancePairs: []sql.ManagedInstancePairInfo{
				{
					PrimaryManagedInstanceReference: tc.MakeReferenceFromResource(primary),
					PartnerManagedInstanceReference: tc.MakeReferenceFromResource(secondary),
				},
			},
			ReadWriteEndpoint: &sql.InstanceFailoverGroupReadWriteEndpoint{
				FailoverPolicy:                         to.Ptr(sql.ReadWriteEndpointFailoverPolicy_Automatic),
				FailoverWithDataLossGracePeriodMinutes: to.Ptr(60),
			},
		},
	}

	tc.CreateResourceAndWait(failoverGroup)

	tc.Expect(failoverGroup.Status.Id).ToNot(BeNil())
	armId := *failoverGroup.Status.Id

	tc.DeleteResourceAndWait(failoverGroup)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// newManagedInstanceNetwork creates a virtual network and a subnet delegated to SQL Managed Instance, in the test's
// default region. See newManagedInstanceNetworkInRegion for what the subnet needs.
func newManagedInstanceNetwork(
	tc *testcommon.KubePerTestContext,
	rg genruntime.ARMMetaObject,
	name string,
	vnetPrefix string,
	subnetPrefix string,
) (*network.VirtualNetwork, *network.VirtualNetworksSubnet) {
	return newManagedInstanceNetworkInRegion(tc, rg, tc.AzureRegion, name, vnetPrefix, subnetPrefix)
}

// managedInstanceNSGRule describes one rule of the network security group attached to a managed instance subnet.
type managedInstanceNSGRule struct {
	name      string
	priority  int
	direction network.SecurityRuleDirection
	protocol  network.SecurityRulePropertiesFormat_Protocol
	source    string
	dest      string
	destPort  string
}

// managedInstanceNSGRules are the rules a managed instance subnet needs. Azure adds its own mandatory management rules
// (priorities 100-105), so these start at 200. The last three allow geo-replication between the two instances of a
// failover group (TCP 5022 and 11000-11999 between their subnets).
func managedInstanceNSGRules() []managedInstanceNSGRule {
	inbound := network.SecurityRuleDirection_Inbound
	outbound := network.SecurityRuleDirection_Outbound
	tcp := network.SecurityRulePropertiesFormat_Protocol_Tcp
	anyProtocol := network.SecurityRulePropertiesFormat_Protocol_Star

	return []managedInstanceNSGRule{
		{"allowtdsinbound", 200, inbound, tcp, "VirtualNetwork", "*", "1433"},
		{"allowredirectinbound", 210, inbound, tcp, "VirtualNetwork", "*", "11000-11999"},
		{"allowmanagementinbound", 220, inbound, tcp, "SqlManagement", "*", "9000-9003"},
		{"allowhealthprobeinbound", 230, inbound, anyProtocol, "AzureLoadBalancer", "*", "*"},
		{"allowgeoreplicationinbound", 240, inbound, tcp, "VirtualNetwork", "*", "5022"},
		{"allowhttpsoutbound", 200, outbound, tcp, "*", "*", "443"},
		{"allowmanagementoutbound", 210, outbound, tcp, "*", "*", "12000"},
		{"allowgeoreplicationoutbound", 220, outbound, tcp, "*", "VirtualNetwork", "5022"},
		{"allowredirectoutbound", 230, outbound, tcp, "*", "VirtualNetwork", "11000-11999"},
	}
}

// newManagedInstanceNetworkInRegion creates a virtual network and a subnet for a managed instance.
// A managed instance subnet must be delegated to Microsoft.Sql/managedInstances, and (as the subnets of existing
// instances show) it also needs its own network security group and an empty route table. Do not attach a route table
// that sends traffic through a firewall, as that breaks the managed instance management plane.
func newManagedInstanceNetworkInRegion(
	tc *testcommon.KubePerTestContext,
	rg genruntime.ARMMetaObject,
	region *string,
	name string,
	vnetPrefix string,
	subnetPrefix string,
) (*network.VirtualNetwork, *network.VirtualNetworksSubnet) {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMeta(name + "vnet"),
		Spec: network.VirtualNetwork_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: region,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{vnetPrefix},
			},
		},
	}

	routeTable := &network.RouteTable{
		ObjectMeta: tc.MakeObjectMeta(name + "rt"),
		Spec: network.RouteTable_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: region,
		},
	}

	nsg := &network.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMeta(name + "nsg"),
		Spec: network.NetworkSecurityGroup_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: region,
		},
	}

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta(name + "subnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr(subnetPrefix),
			Delegations: []network.Delegation{
				{
					Name:        to.Ptr("managedinstancedelegation"),
					ServiceName: to.Ptr("Microsoft.Sql/managedInstances"),
				},
			},
			NetworkSecurityGroup: &network.NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(nsg),
			},
			RouteTable: &network.RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(routeTable),
			},
		},
	}

	objs := []client.Object{vnet, routeTable, nsg}
	allow := network.SecurityRuleAccess_Allow
	for _, r := range managedInstanceNSGRules() {
		objs = append(objs, &network.NetworkSecurityGroupsSecurityRule{
			ObjectMeta: tc.MakeObjectMeta(name + r.name),
			Spec: network.NetworkSecurityGroupsSecurityRule_Spec{
				Owner:                    testcommon.AsOwner(nsg),
				Priority:                 to.Ptr(r.priority),
				Direction:                to.Ptr(r.direction),
				Access:                   to.Ptr(allow),
				Protocol:                 to.Ptr(r.protocol),
				SourceAddressPrefix:      to.Ptr(r.source),
				SourcePortRange:          to.Ptr("*"),
				DestinationAddressPrefix: to.Ptr(r.dest),
				DestinationPortRange:     to.Ptr(r.destPort),
			},
		})
	}
	objs = append(objs, subnet)

	tc.CreateResourcesAndWait(objs...)

	return vnet, subnet
}

func newManagedInstancePeering(
	tc *testcommon.KubePerTestContext,
	name string,
	from *network.VirtualNetwork,
	remote *network.VirtualNetwork,
) *network.VirtualNetworksVirtualNetworkPeering {
	return &network.VirtualNetworksVirtualNetworkPeering{
		ObjectMeta: tc.MakeObjectMeta(name),
		Spec: network.VirtualNetworksVirtualNetworkPeering_Spec{
			Owner:                     testcommon.AsOwner(from),
			AllowVirtualNetworkAccess: to.Ptr(true),
			AllowForwardedTraffic:     to.Ptr(true),
			RemoteVirtualNetwork: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(remote),
			},
		},
	}
}
