/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	arm "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301/arm"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_FuzzySetExtensions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vmss := &arm.VirtualMachineScaleSet_Spec{
		Location: to.Ptr("westus"),
		Properties: &arm.VirtualMachineScaleSetProperties{
			ZoneBalance: to.Ptr(false),
		},
	}

	rawVMSSWithExtensions := map[string]any{
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-isodfv/providers/Microsoft.Compute/virtualMachineScaleSets/asotest-vmss-idossk",
		"location": "westus",
		"name":     "asotest-vmss-idossk",
		"properties": map[string]any{
			"doNotRunExtensionsOnOverprovisionedVMs": false,
			"orchestrationMode":                      "Uniform",
			"overprovision":                          true,
			"platformFaultDomainCount":               3,
			"singlePlacementGroup":                   false,
			"upgradePolicy": map[string]any{
				"mode": "Automatic",
			},
			"virtualMachineProfile": map[string]any{
				"extensionProfile": map[string]any{
					"extensions": []any{
						map[string]any{
							"name": "mycustomextension",
							"properties": map[string]any{
								"autoUpgradeMinorVersion": false,
								"publisher":               "Microsoft.Azure.Extensions",
								"settings": map[string]any{
									"commandToExecute": "/bin/bash -c ",
								},
								"type":               "CustomScript",
								"typeHandlerVersion": "2.0",
							},
						},
					},
				},
			},
		},
		"type": "Microsoft.Compute/virtualMachineScaleSets",
	}

	// Note that many of these fields are readonly and will not be present on the PUT

	vmssExtensions, err := getRawChildCollection(rawVMSSWithExtensions, rawChildCollectionPath...)
	g.Expect(err).ToNot(HaveOccurred())

	err = setChildCollection(vmss, vmssExtensions, childCollectionPathARM...)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vmss.Location).To(Equal(to.Ptr("westus")))
	g.Expect(vmss.Properties).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0]).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Name).To(Equal(to.Ptr("mycustomextension")))
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Properties).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Properties.Type).To(Equal(to.Ptr("CustomScript")))
}

func Test_FuzzySetVMSS_ExtensionsMerged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vmss := &arm.VirtualMachineScaleSet_Spec{
		Name:     "asotest-vmss-idossk",
		Location: to.Ptr("westus"),
		Properties: &arm.VirtualMachineScaleSetProperties{
			VirtualMachineProfile: &arm.VirtualMachineScaleSetVMProfile{
				ExtensionProfile: &arm.VirtualMachineScaleSetExtensionProfile{
					Extensions: []arm.VirtualMachineScaleSetExtension{
						{
							Name: to.Ptr("mycustomextension"),
							Properties: &arm.VirtualMachineScaleSetExtensionProperties{
								Publisher:          to.Ptr("Microsoft.Azure.Extensions"),
								Type:               to.Ptr("CustomScript"),
								TypeHandlerVersion: to.Ptr("2.0"),
							},
						},
						{
							Name: to.Ptr("mycustomextension1"),
							Properties: &arm.VirtualMachineScaleSetExtensionProperties{
								Publisher:          to.Ptr("Microsoft.ManagedServices"),
								Type:               to.Ptr("ApplicationHealthLinux"),
								TypeHandlerVersion: to.Ptr("1.0"),
							},
						},
					},
				},
			},
		},
	}

	rawVMSSWithExtensions := map[string]any{
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-isodfv/providers/Microsoft.Compute/virtualMachineScaleSets/asotest-vmss-idossk",
		"location": "westus",
		"name":     "asotest-vmss-idossk",
		"properties": map[string]any{
			"doNotRunExtensionsOnOverprovisionedVMs": false,
			"orchestrationMode":                      "Uniform",
			"overprovision":                          true,
			"platformFaultDomainCount":               3,
			"provisioningState":                      "Succeeded",
			"singlePlacementGroup":                   false,
			"upgradePolicy": map[string]any{
				"mode": "Automatic",
			},
			"virtualMachineProfile": map[string]any{
				"extensionProfile": map[string]any{
					"extensions": []any{
						map[string]any{
							"name": "mycustomextension1",
							"properties": map[string]any{
								"publisher":          "Microsoft.ManagedServices",
								"type":               "ApplicationHealthLinux",
								"typeHandlerVersion": "1.0",
							},
						},
					},
				},
			},
		},
		"type": "Microsoft.Compute/virtualMachineScaleSets",
	}

	vmssExtension, err := getRawChildCollection(rawVMSSWithExtensions, rawChildCollectionPath...)
	g.Expect(err).ToNot(HaveOccurred())

	err = setChildCollection(vmss, vmssExtension, childCollectionPathARM...)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vmss.Location).To(Equal(to.Ptr("westus")))
	g.Expect(vmss.Properties).ToNot(BeNil())
	g.Expect(vmss.Properties).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions).To(HaveLen(2))
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0]).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Name).To(Equal(to.Ptr("mycustomextension")))
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Properties).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[0].Properties.Type).To(Equal(to.Ptr("CustomScript")))
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[1]).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[1].Name).To(Equal(to.Ptr("mycustomextension1")))
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[1].Properties).ToNot(BeNil())
	g.Expect(vmss.Properties.VirtualMachineProfile.ExtensionProfile.Extensions[1].Properties.Type).To(Equal(to.Ptr("ApplicationHealthLinux")))
}
