/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	containerinstance "github.com/Azure/azure-service-operator/v2/api/containerinstance/v1beta20211001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ContainerInstance_ContainerGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// The test refers to the quick-start-template from https://github.com/Azure/azure-quickstart-templates/tree/master/quickstarts/microsoft.containerinstance/aci-linuxcontainer-public-ip
	name := tc.NoSpaceNamer.GenerateName("containergroup")
	protocol := containerinstance.ContainerPortProtocolTCP
	osType := containerinstance.ContainerGroupsSpecPropertiesOsTypeLinux
	restartPolicy := containerinstance.ContainerGroupsSpecPropertiesRestartPolicyAlways
	ipAddressType := containerinstance.IpAddressTypePublic
	portProtocol := containerinstance.PortProtocolTCP

	// Create a ContainerGroup
	cg := containerinstance.ContainerGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: containerinstance.ContainerGroups_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Containers: []containerinstance.ContainerGroups_Spec_Properties_Containers{
				{
					Name:  &name,
					Image: to.StringPtr("mcr.microsoft.com/azuredocs/aci-helloworld"),
					Ports: []containerinstance.ContainerPort{
						{
							Port:     to.IntPtr(80),
							Protocol: &protocol,
						},
					},
					Resources: &containerinstance.ResourceRequirements{
						Requests: &containerinstance.ResourceRequests{
							Cpu:        to.Float64Ptr(1),
							MemoryInGB: to.Float64Ptr(2),
						},
					},
				},
			},
			OsType:        &osType,
			RestartPolicy: &restartPolicy,
			IpAddress: &containerinstance.IpAddress{
				Type: &ipAddressType,
				Ports: []containerinstance.Port{
					{
						Port:     to.IntPtr(80),
						Protocol: &portProtocol,
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(&cg)
	defer tc.DeleteResourcesAndWait(&cg)

	// Perform some assertions on the resources we just created
	tc.Expect(cg.Status.Id).ToNot(BeNil())
}
