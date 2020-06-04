/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

const (
	// VirtualMachineScaleSetScaleInRuleDefault is the default
	VirtualMachineScaleSetScaleInRuleDefault VirtualMachineScaleSetScaleInRules = "Default"
	// VirtualMachineScaleSetScaleInRuleNewestVM chooses the newest VM
	VirtualMachineScaleSetScaleInRuleNewestVM VirtualMachineScaleSetScaleInRules = "NewestVM"
	// VirtualMachineScaleSetScaleInRuleOldestVM chooses the oldest VM
	VirtualMachineScaleSetScaleInRuleOldestVM VirtualMachineScaleSetScaleInRules = "OldestVM"
)

type (
	Sku struct {
		// Name - The sku name.
		Name string `json:"name"`
		// Tier - Specifies the tier of virtual machines in a scale set. Possible Values: **Standard** **Basic**
		// +kubebuilder:validation:Enum=Standard;Basic
		// +kubebuilder:default=Standard
		Tier string `json:"tier"`
		// Capacity - Specifies the number of virtual machines in the scale set.
		Capacity int64 `json:"capacity"`
	}

	VirtualMachineScaleSetProperties struct {
		// UpgradePolicy - The upgrade policy.
		UpgradePolicy UpgradePolicy `json:"upgradePolicy"`
		// AutomaticRepairsPolicy - Policy for automatic repairs.
		AutomaticRepairsPolicy *AutomaticRepairsPolicy `json:"automaticRepairsPolicy,omitempty"`
		// VirtualMachineProfile - The virtual machine profile.
		VirtualMachineProfile VirtualMachineScaleSetVMProfile `json:"virtualMachineProfile"`
		// ProvisioningState - READ-ONLY; The provisioning state, which only appears in the response.
		ProvisioningState *string `json:"provisioningState,omitempty"`
		// Overprovision - Specifies whether the Virtual Machine Scale Set should be overprovisioned.
		Overprovision *bool `json:"overprovision,omitempty"`
		// DoNotRunExtensionsOnOverprovisionedVMs - When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
		DoNotRunExtensionsOnOverprovisionedVMs *bool `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
		// SinglePlacementGroup - When true this limits the scale set to a single placement group, of max size 100 virtual machines. NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may not be modified to true.
		SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty"`
		// ZoneBalance - Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
		ZoneBalance *bool `json:"zoneBalance,omitempty"`
		// PlatformFaultDomainCount - Fault Domain count for each placement group.
		PlatformFaultDomainCount *int32 `json:"platformFaultDomainCount,omitempty"`
		// AdditionalCapabilities - Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
		AdditionalCapabilities *AdditionalCapabilities `json:"additionalCapabilities,omitempty"`
		// ScaleInPolicy - Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
		ScaleInPolicy *ScaleInPolicy `json:"scaleInPolicy,omitempty"`
	}

	// +kubebuilder:validation:Enum=Default;NewestVM;OldestVM
	VirtualMachineScaleSetScaleInRules string

	ScaleInPolicy struct {
		// Rules - The rules to be followed when scaling-in a virtual machine scale set. Possible values are: **Default** When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the virtual machines chosen for removal will be the newest ones that are not protected from scale-in. **OldestVM** When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal. **NewestVM** When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.
		Rules []VirtualMachineScaleSetScaleInRules `json:"rules,omitempty"`
	}

	AdditionalCapabilities struct {
		// UltraSSDEnabled - The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual machine or virtual machine scale set only if this property is enabled.
		UltraSSDEnabled *bool `json:"ultraSSDEnabled,omitempty"`
	}

	UpgradePolicy struct {
		// Mode - Specifies the mode of an upgrade to virtual machines in the scale set. Possible values are: **Manual** - You  control the application of updates to virtual machines in the scale set. You do this by using the manualUpgrade action. **Automatic** - All virtual machines in the scale set are  automatically updated at the same time.
		// +kubebuilder:validation:Enum=Automatic;Manual;Rolling
		// +kubebuilder:default=Manual
		Mode *string `json:"mode,omitempty"`
		// RollingUpgradePolicy - The configuration parameters used while performing a rolling upgrade.
		RollingUpgradePolicy *RollingUpgradePolicy `json:"rollingUpgradePolicy,omitempty"`
		// AutomaticOSUpgradePolicy - Configuration parameters used for performing automatic OS Upgrade.
		AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicy `json:"automaticOSUpgradePolicy,omitempty"`
	}

	RollingUpgradePolicy struct {
		// MaxBatchInstancePercent - The maximum percent of total virtual machine instances that will be upgraded simultaneously by the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.
		MaxBatchInstancePercent *int32 `json:"maxBatchInstancePercent,omitempty"`
		// MaxUnhealthyInstancePercent - The maximum percentage of the total virtual machine instances in the scale set that can be simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch. The default value for this parameter is 20%.
		MaxUnhealthyInstancePercent *int32 `json:"maxUnhealthyInstancePercent,omitempty"`
		// MaxUnhealthyUpgradedInstancePercent - The maximum percentage of upgraded virtual machine instances that can be found to be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the rolling update aborts. The default value for this parameter is 20%.
		MaxUnhealthyUpgradedInstancePercent *int32 `json:"maxUnhealthyUpgradedInstancePercent,omitempty"`
		// PauseTimeBetweenBatches - The wait time between completing the update for all virtual machines in one batch and starting the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).
		PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches,omitempty"`
	}

	AutomaticOSUpgradePolicy struct {
		// EnableAutomaticOSUpgrade - Indicates whether OS upgrades should automatically be applied to scale set instances in a rolling fashion when a newer version of the OS image becomes available. Default value is false.  If this is set to true for Windows based scale sets, [enableAutomaticUpdates](https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet) is automatically set to false and cannot be set to true.
		EnableAutomaticOSUpgrade *bool `json:"enableAutomaticOSUpgrade,omitempty"`
		// DisableAutomaticRollback - Whether OS image rollback feature should be disabled. Default value is false.
		DisableAutomaticRollback *bool `json:"disableAutomaticRollback,omitempty"`
	}

	AutomaticRepairsPolicy struct {
		// Enabled - Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is false.
		Enabled *bool `json:"enabled,omitempty"`
		// GracePeriod - The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default value. The maximum allowed grace period is 90 minutes (PT90M).
		GracePeriod *string `json:"gracePeriod,omitempty"`
	}

	VirtualMachineScaleSetVMProfile struct {
		// OsProfile - Specifies the operating system settings for the virtual machines in the scale set.
		OsProfile VirtualMachineScaleSetOSProfile `json:"osProfile"`
		// StorageProfile - Specifies the storage settings for the virtual machine disks.
		StorageProfile VirtualMachineScaleSetStorageProfile `json:"storageProfile"`
		// NetworkProfile - Specifies properties of the network interfaces of the virtual machines in the scale set.
		NetworkProfile VirtualMachineScaleSetNetworkProfile `json:"networkProfile"`
		// DiagnosticsProfile - Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
		DiagnosticsProfile *DiagnosticsProfile `json:"diagnosticsProfile,omitempty"`
		// Priority - Specifies the priority for the virtual machines in the scale set. Minimum api-version: 2017-10-30-preview. Possible values include: 'Regular', 'Low', 'Spot'
		// +kubebuilder:validation:Enum=Regular;Low;Spot
		Priority *string `json:"priority,omitempty"`
		// EvictionPolicy - Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. For Azure Spot virtual machines, the only supported value is 'Deallocate' and the minimum api-version is 2019-03-01. For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview. Possible values include: 'Deallocate', 'Delete'
		// +kubebuilder:validation:Enum=Deallocate;Delete
		EvictionPolicy *string `json:"evictionPolicy,omitempty"`
		// BillingProfile - Specifies the billing related details of a Azure Spot VMSS. Minimum api-version: 2019-03-01.
		BillingProfile *BillingProfile `json:"billingProfile,omitempty"`
		// ScheduledEventsProfile - Specifies Scheduled Event related configurations.
		ScheduledEventsProfile *ScheduledEventsProfile `json:"scheduledEventsProfile,omitempty"`
	}

	ScheduledEventsProfile struct {
		// TerminateNotificationProfile - Specifies Terminate Scheduled Event related configurations.
		TerminateNotificationProfile *TerminateNotificationProfile `json:"terminateNotificationProfile,omitempty"`
	}

	TerminateNotificationProfile struct {
		// NotBeforeTimeout - Configurable length of time a Virtual Machine being deleted will have to potentially approve the Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601 format, the default value is 5 minutes (PT5M)
		NotBeforeTimeout *string `json:"notBeforeTimeout,omitempty"`
		// Enable - Specifies whether the Terminate Scheduled event is enabled or disabled.
		Enable *bool `json:"enable,omitempty"`
	}

	BillingProfile struct {
		// MaxPrice - Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.  This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current Azure Spot price.  The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the maxPrice after creation of VM/VMSS.  Possible values are:  - Any decimal value greater than zero. Example: 0.01538  -1 – indicates default price to be up-to on-demand.  You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also, the default max price is -1 if it is not provided by you. Minimum api-version: 2019-03-01.
		// +optional
		// +kubebuilder:validation:pattern="^[0-9]+(\.[0-9]+)?$"
		MaxPrice *string `json:"maxPrice,omitempty"`
	}

	VirtualMachineScaleSetNetworkProfile struct {
		// HealthProbe - A reference to a load balancer probe used to determine the health of an instance in the virtual machine scale set. The reference will be in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}'.
		HealthProbeRef *azcorev1.KnownTypeReference `json:"healthProbeRef,omitempty"`
		// NetworkInterfaceConfigurations - The list of network configurations.
		NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration `json:"networkInterfaceConfigurations,omitempty"`
	}

	VirtualMachineScaleSetNetworkConfiguration struct {
		// Name - The network configuration name.
		Name       *string                                               `json:"name,omitempty"`
		Properties *VirtualMachineScaleSetNetworkConfigurationProperties `json:"properties,omitempty"`
	}

	VirtualMachineScaleSetIPConfiguration struct {
		Name       *string                                          `json:"name,omitempty"`
		Properties *VirtualMachineScaleSetIPConfigurationProperties `json:"properties,omitempty"`
	}

	VirtualMachineScaleSetIPConfigurationProperties struct {
		// Subnet - Specifies the identifier of the subnet.
		SubnetRef *azcorev1.KnownTypeReference `json:"subnetRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"Subnet"`
		// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
		Primary *bool `json:"primary,omitempty"`
		// PublicIPAddressConfiguration - The publicIPAddressConfiguration.
		PublicIPAddressConfiguration *VirtualMachineScaleSetPublicIPAddressConfiguration `json:"publicIPAddressConfiguration,omitempty"`
		// +kubebuilder:validation:Enum=IPv4;IPv6
		PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty"`
		// LoadBalancerBackendAddressPools - Specifies an array of references to backend address pools of load balancers. A scale set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer.
		LoadBalancerBackendAddressPoolRefs []azcorev1.KnownTypeReference `json:"loadBalancerBackendAddressPoolRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"BackendAddressPool"`
		// LoadBalancerInboundNatPools - Specifies an array of references to inbound Nat pools of the load balancers. A scale set can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer
		LoadBalancerInboundNatPoolRefs []azcorev1.KnownTypeReference `json:"loadBalancerInboundNatPoolRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"InboundNatPool"`
	}

	VirtualMachineScaleSetPublicIPAddressConfiguration struct {
		// Name - The publicIP address configuration name.
		Name       *string                                                       `json:"name,omitempty"`
		Properties *VirtualMachineScaleSetPublicIPAddressConfigurationProperties `json:"properties,omitempty"`
	}

	VirtualMachineScaleSetPublicIPAddressConfigurationProperties struct {
		// IdleTimeoutInMinutes - The idle timeout of the public IP address.
		IdleTimeoutInMinutes *int32 `json:"idleTimeoutInMinutes,omitempty"`
		// DNSSettings - The dns settings to be applied on the publicIP addresses .
		DNSSettings *VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings `json:"dnsSettings,omitempty"`
		// IPTags - The list of IP tags associated with the public IP address.
		IPTags []VirtualMachineScaleSetIPTag `json:"ipTags,omitempty"`
		// +kubebuilder:validation:Enum=IPv4;IPv6
		PublicIPAddressVersion string `json:"publicIPAddressVersion,omitempty"`
	}

	VirtualMachineScaleSetIPTag struct {
		// IPTagType - IP tag type. Example: FirstPartyUsage.
		IPTagType *string `json:"ipTagType,omitempty"`
		// Tag - IP tag associated with the public IP. Example: SQL, Storage etc.
		Tag *string `json:"tag,omitempty"`
	}

	VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings struct {
		// DomainNameLabel - The Domain name label.The concatenation of the domain name label and vm index will be the domain name labels of the PublicIPAddress resources that will be created
		DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	}

	VirtualMachineScaleSetNetworkConfigurationProperties struct {
		// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
		Primary *bool `json:"primary,omitempty"`
		// EnableAcceleratedNetworking - Specifies whether the network interface is accelerated networking-enabled.
		EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`
		// NetworkSecurityGroup - The network security group.
		NetworkSecurityGroupRef *azcorev1.KnownTypeReference `json:"networkSecurityGroupRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"NetworkSecurityGroup"`
		// DNSSettings - The dns settings to be applied on the network interfaces.
		DNSSettings *VirtualMachineScaleSetNetworkConfigurationDNSSettings `json:"dnsSettings,omitempty"`
		// IPConfigurations - Specifies the IP configurations of the network interface.
		IPConfigurations []VirtualMachineScaleSetIPConfiguration `json:"ipConfigurations,omitempty"`
		// EnableIPForwarding - Whether IP forwarding enabled on this NIC.
		EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`
	}

	VirtualMachineScaleSetNetworkConfigurationDNSSettings struct {
		// DNSServers - List of DNS servers IP addresses
		DNSServers []string `json:"dnsServers,omitempty"`
	}

	VirtualMachineScaleSetStorageProfile struct {
		// ImageReference - Specifies information about the image to use. You can specify information about platform images, marketplace images, or virtual machine images. This element is required when you want to use a platform image, marketplace image, or virtual machine image, but is not used in other creation operations.
		ImageReference *ImageReference `json:"imageReference,omitempty"`
		// OsDisk - Specifies information about the operating system disk used by the virtual machines in the scale set.  For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
		OsDisk *VirtualMachineScaleSetOSDisk `json:"osDisk,omitempty"`
		// DataDisks - Specifies the parameters that are used to add data disks to the virtual machines in the scale set.  For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
		DataDisks []VirtualMachineScaleSetDataDisk `json:"dataDisks,omitempty"`
	}

	VirtualMachineScaleSetDataDisk struct {
		// Name - The disk name.
		Name *string `json:"name,omitempty"`
		// Lun - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
		Lun *int32 `json:"lun,omitempty"`
		// Caching - Specifies the caching requirements.
		// +kubebuilder:validation:Enum=None;ReadOnly;ReadWrite
		// +kubebuilder:default=None
		Caching string `json:"caching,omitempty"`
		// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
		WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
		// CreateOption - Specifies how the virtual machine should be created.
		// +kubebuilder:validation:Enum=FromImage;Empty;Attach
		CreateOption string `json:"createOption,omitempty"`
		// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image.  This value cannot be larger than 1023 GB
		DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
		// ManagedDisk - The managed disk parameters.
		ManagedDisk *VirtualMachineScaleSetManagedDiskParameters `json:"managedDisk,omitempty"`
		// DiskIOPSReadWrite - Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.
		DiskIOPSReadWrite *int64 `json:"diskIOPSReadWrite,omitempty"`
		// DiskMBpsReadWrite - Specifies the bandwidth in MB per second for the managed disk. Should be used only when StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.
		DiskMBpsReadWrite *int64 `json:"diskMBpsReadWrite,omitempty"`
	}

	VirtualMachineScaleSetOSDisk struct {
		// Name - The disk name.
		Name *string `json:"name,omitempty"`
		// Caching - Specifies the caching requirements.
		// +kubebuilder:validation:Enum=None;ReadOnly;ReadWrite
		// +kubebuilder:default=None
		Caching string `json:"caching,omitempty"`
		// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
		WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
		// CreateOption - Specifies how the virtual machine should be created.
		// +kubebuilder:validation:Enum=FromImage;Empty;Attach
		CreateOption string `json:"createOption,omitempty"`
		// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image. This value cannot be larger than 1023 GB
		// +kubebuilder:validation:Maximum=1023
		DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
		// OsType - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD.
		// +kubebuilder:validation:Enum=Windows;Linux
		OsType *string `json:"osType,omitempty"`
		// Image - Specifies information about the unmanaged user image to base the scale set on.
		Image *VirtualHardDisk `json:"image,omitempty"`
		// VhdContainers - Specifies the container urls that are used to store operating system disks for the scale set.
		VhdContainers []string `json:"vhdContainers,omitempty"`
		// ManagedDisk - The managed disk parameters.
		ManagedDisk *VirtualMachineScaleSetManagedDiskParameters `json:"managedDisk,omitempty"`
	}

	VirtualMachineScaleSetManagedDiskParameters struct {
		// StorageAccountType - Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk.
		// +kubebuilder:validation:Enum=Premium_LRS;StandardSSD_LRS;Standard_LRS;UltraSSD_LRS
		StorageAccountType *string `json:"storageAccountType,omitempty"`
	}

	VirtualMachineScaleSetOSProfile struct {
		// ComputerNamePrefix - Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name prefixes must be 1 to 15 characters long.
		ComputerNamePrefix string `json:"computerNamePrefix"`
		// AdminUsername - Specifies the name of the administrator account.  **Windows-only restriction:** Cannot end in "."  **Disallowed values:** "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5".  **Minimum-length (Linux):** 1  character  **Max-length (Linux):** 64 characters  **Max-length (Windows):** 20 characters   For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json) For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
		AdminUsername string `json:"adminUsername"`
		// AdminPassword - Specifies the password of the administrator account.  **Minimum-length (Windows):** 8 characters  **Minimum-length (Linux):** 6 characters  **Max-length (Windows):** 123 characters  **Max-length (Linux):** 72 characters  **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled  Has lower characters Has upper characters  Has a digit  Has a special character (Regex match [\W_])  **Disallowed values:** "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!"  For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json)  For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
		AdminPassword *string `json:"adminPassword,omitempty"`
		// CustomData - Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.  For using cloud-init for your VM, see [Using cloud-init to customize a Linux VM during creation](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
		CustomData *string `json:"customData,omitempty"`
		// LinuxConfiguration - Specifies the Linux operating system settings on the virtual machine. For a list of supported Linux distributions, see [Linux on Azure-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)  For running non-endorsed distributions, see [Information for Non-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
		LinuxConfiguration *LinuxConfiguration `json:"linuxConfiguration,omitempty"`
	}

	VirtualMachineScaleSetIdentity struct {
		// Type - The type of identity used for the virtual machine. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
		// +kubebuilder:validation:Enum=SystemAssigned;UserAssigned;None
		Type string `json:"type"`
		// UserAssignedIdentities - The list of user identities associated with the Virtual Machine. The user identity references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
		UserAssignedIdentities []string `json:"userAssignedIdentities"`
	}

	// VirtualMachineScaleSetSpec defines the desired state of VirtualMachineScaleSet
	VirtualMachineScaleSetSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion"`
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`
		// Sku - The virtual machine scale set sku.
		Sku Sku `json:"sku"`
		// Plan - Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
		Plan       *Plan                            `json:"plan,omitempty"`
		Properties VirtualMachineScaleSetProperties `json:"properties"`
		// Identity - The identity of the virtual machine scale set, if configured.
		Identity *VirtualMachineScaleSetIdentity `json:"identity,omitempty"`
		// Zones - The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
		Zones []string `json:"zones,omitempty"`
		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`
		// Tags - Resource tags
		Tags map[string]string `json:"tags,omitempty"`
	}

	// VirtualMachineScaleSetStatus defines the observed state of VirtualMachineScaleSet
	VirtualMachineScaleSetStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// VirtualMachineScaleSet is the Schema for the virtualmachinescalesets API
	VirtualMachineScaleSet struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   VirtualMachineScaleSetSpec   `json:"spec,omitempty"`
		Status VirtualMachineScaleSetStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// VirtualMachineScaleSetList contains a list of VirtualMachineScaleSet
	VirtualMachineScaleSetList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []VirtualMachineScaleSet `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&VirtualMachineScaleSet{}, &VirtualMachineScaleSetList{})
}
