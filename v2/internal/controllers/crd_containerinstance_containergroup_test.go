/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

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
	image := "mcr.microsoft.com/azuredocs/aci-helloworld"
	port := int(80)
	protocol := containerinstance.ContainerPortProtocolTCP
	cpu := float64(1)
	memoryInGB := float64(2)
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
					Image: &image,
					Ports: []containerinstance.ContainerPort{
						{
							Port:     &port,
							Protocol: &protocol,
						},
					},
					Resources: &containerinstance.ResourceRequirements{
						Requests: &containerinstance.ResourceRequests{
							Cpu:        &cpu,
							MemoryInGB: &memoryInGB,
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
						Port:     &port,
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
