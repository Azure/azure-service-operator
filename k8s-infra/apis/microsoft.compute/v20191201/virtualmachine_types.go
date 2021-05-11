/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191201

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

type (
	Plan struct {
		Name string `json:"name,omitempty"`
		// Publisher - The publisher ID.
		Publisher string `json:"publisher,omitempty"`
		// Product - Specifies the product of the image from the marketplace. This is the same value as Offer under the imageReference element.
		Product string `json:"product,omitempty"`
		// PromotionCode - The promotion code.
		PromotionCode string `json:"promotionCode,omitempty"`
	}

	HardwareProfile struct {
		// VMSize - Specifies the size of the virtual machine. For more information about virtual machine sizes, see [Sizes for virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-sizes?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). The available VM sizes depend on region and availability set. For a list of available sizes use these APIs:  [List all available virtual machine sizes in an availability set](https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes) [List all available virtual machine sizes in a region](https://docs.microsoft.com/rest/api/compute/virtualmachinesizes/list) [List all available virtual machine sizes for resizing](https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes). Possible values include: 'BasicA0', 'BasicA1', 'BasicA2', 'BasicA3', 'BasicA4', 'StandardA0', 'StandardA1', 'StandardA2', 'StandardA3', 'StandardA4', 'StandardA5', 'StandardA6', 'StandardA7', 'StandardA8', 'StandardA9', 'StandardA10', 'StandardA11', 'StandardA1V2', 'StandardA2V2', 'StandardA4V2', 'StandardA8V2', 'StandardA2mV2', 'StandardA4mV2', 'StandardA8mV2', 'StandardB1s', 'StandardB1ms', 'StandardB2s', 'StandardB2ms', 'StandardB4ms', 'StandardB8ms', 'StandardD1', 'StandardD2', 'StandardD3', 'StandardD4', 'StandardD11', 'StandardD12', 'StandardD13', 'StandardD14', 'StandardD1V2', 'StandardD2V2', 'StandardD3V2', 'StandardD4V2', 'StandardD5V2', 'StandardD2V3', 'StandardD4V3', 'StandardD8V3', 'StandardD16V3', 'StandardD32V3', 'StandardD64V3', 'StandardD2sV3', 'StandardD4sV3', 'StandardD8sV3', 'StandardD16sV3', 'StandardD32sV3', 'StandardD64sV3', 'StandardD11V2', 'StandardD12V2', 'StandardD13V2', 'StandardD14V2', 'StandardD15V2', 'StandardDS1', 'StandardDS2', 'StandardDS3', 'StandardDS4', 'StandardDS11', 'StandardDS12', 'StandardDS13', 'StandardDS14', 'StandardDS1V2', 'StandardDS2V2', 'StandardDS3V2', 'StandardDS4V2', 'StandardDS5V2', 'StandardDS11V2', 'StandardDS12V2', 'StandardDS13V2', 'StandardDS14V2', 'StandardDS15V2', 'StandardDS134V2', 'StandardDS132V2', 'StandardDS148V2', 'StandardDS144V2', 'StandardE2V3', 'StandardE4V3', 'StandardE8V3', 'StandardE16V3', 'StandardE32V3', 'StandardE64V3', 'StandardE2sV3', 'StandardE4sV3', 'StandardE8sV3', 'StandardE16sV3', 'StandardE32sV3', 'StandardE64sV3', 'StandardE3216V3', 'StandardE328sV3', 'StandardE6432sV3', 'StandardE6416sV3', 'StandardF1', 'StandardF2', 'StandardF4', 'StandardF8', 'StandardF16', 'StandardF1s', 'StandardF2s', 'StandardF4s', 'StandardF8s', 'StandardF16s', 'StandardF2sV2', 'StandardF4sV2', 'StandardF8sV2', 'StandardF16sV2', 'StandardF32sV2', 'StandardF64sV2', 'StandardF72sV2', 'StandardG1', 'StandardG2', 'StandardG3', 'StandardG4', 'StandardG5', 'StandardGS1', 'StandardGS2', 'StandardGS3', 'StandardGS4', 'StandardGS5', 'StandardGS48', 'StandardGS44', 'StandardGS516', 'StandardGS58', 'StandardH8', 'StandardH16', 'StandardH8m', 'StandardH16m', 'StandardH16r', 'StandardH16mr', 'StandardL4s', 'StandardL8s', 'StandardL16s', 'StandardL32s', 'StandardM64s', 'StandardM64ms', 'StandardM128s', 'StandardM128ms', 'StandardM6432ms', 'StandardM6416ms', 'StandardM12864ms', 'StandardM12832ms', 'StandardNC6', 'StandardNC12', 'StandardNC24', 'StandardNC24r', 'StandardNC6sV2', 'StandardNC12sV2', 'StandardNC24sV2', 'StandardNC24rsV2', 'StandardNC6sV3', 'StandardNC12sV3', 'StandardNC24sV3', 'StandardNC24rsV3', 'StandardND6s', 'StandardND12s', 'StandardND24s', 'StandardND24rs', 'StandardNV6', 'StandardNV12', 'StandardNV24'
		VMSize string `json:"vmSize,omitempty"`
	}

	StorageProfile struct {
		// ImageReference - Specifies information about the image to use. You can specify information about platform images, marketplace images, or virtual machine images. This element is required when you want to use a platform image, marketplace image, or virtual machine image, but is not used in other creation operations.
		ImageReference *ImageReference `json:"imageReference,omitempty"`
		// OsDisk - Specifies information about the operating system disk used by the virtual machine. For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
		OsDisk *OSDisk `json:"osDisk,omitempty"`
		// DataDisks - Specifies the parameters that are used to add a data disk to a virtual machine. For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
		DataDisks *[]DataDisk `json:"dataDisks,omitempty"`
	}

	DataDisk struct {
		// Lun - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
		Lun *int32 `json:"lun,omitempty"`
		// Name - The disk name.
		Name *string `json:"name,omitempty"`
		// Vhd - The virtual hard disk.
		Vhd *VirtualHardDisk `json:"vhd,omitempty"`
		// Image - The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
		Image *VirtualHardDisk `json:"image,omitempty"`
		// Caching - Specifies the caching requirements.
		// +kubebuilder:validation:Enum=None;ReadOnly;ReadWrite
		// +kubebuilder:default=None
		Caching string `json:"caching,omitempty"`
		// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
		WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
		// CreateOption - Specifies how the disk should be created.
		// +kubebuilder:validation:Enum=FromImage;Empty;Attach
		CreateOption string `json:"createOption,omitempty"`
		// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image. This value cannot be larger than 1023 GB
		// +kubebuilder:validation:Maximum=1023
		DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
		// ManagedDisk - The managed disk parameters.
		ManagedDisk *ManagedDiskParameters `json:"managedDisk,omitempty"`
		// ToBeDetached - Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset
		ToBeDetached *bool `json:"toBeDetached,omitempty"`
	}

	ImageReference struct {
		// Publisher - The image publisher.
		Publisher *string `json:"publisher,omitempty"`
		// Offer - Specifies the offer of the platform image or marketplace image used to create the virtual machine.
		Offer *string `json:"offer,omitempty"`
		// Sku - The image SKU.
		Sku *string `json:"sku,omitempty"`
		// Version - Specifies the version of the platform image or marketplace image used to create the virtual machine. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available.
		Version *string `json:"version,omitempty"`
		// ID - Resource Id
		ID *string `json:"id,omitempty"`
	}

	OSDisk struct {
		// OsType - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD.
		// +kubebuilder:validation:Enum=Windows;Linux
		OsType *string `json:"osType,omitempty"`
		// Name - The disk name.
		Name *string `json:"name,omitempty"`
		// Vhd - The virtual hard disk.
		Vhd *VirtualHardDisk `json:"vhd,omitempty"`
		// Image - The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
		Image *VirtualHardDisk `json:"image,omitempty"`
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
		// ManagedDisk - The managed disk parameters.
		ManagedDisk *ManagedDiskParameters `json:"managedDisk,omitempty"`
	}

	ManagedDiskParameters struct {
		// StorageAccountType - Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk.
		// +kubebuilder:validation:Enum=Premium_LRS;StandardSSD_LRS;Standard_LRS;UltraSSD_LRS
		StorageAccountType *string `json:"storageAccountType,omitempty"`
		// ID - Resource Id
		ID *string `json:"id,omitempty"`
	}

	VirtualHardDisk struct {
		// URI - Specifies the virtual hard disk's uri.
		URI *string `json:"uri,omitempty"`
	}

	OSProfile struct {
		// ComputerName - Specifies the host OS name of the virtual machine. This name cannot be updated after the VM is created. **Max-length (Windows):** 15 characters **Max-length (Linux):** 64 characters. For naming conventions and restrictions see [Azure infrastructure services implementation guidelines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions).
		ComputerName *string `json:"computerName,omitempty"`
		// AdminUsername - Specifies the name of the administrator account. This property cannot be updated after the VM is created. **Windows-only restriction:** Cannot end in "." **Disallowed values:** "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". **Minimum-length (Linux):** 1  character **Max-length (Linux):** 64 characters **Max-length (Windows):** 20 characters   For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json) For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
		AdminUsername *string `json:"adminUsername,omitempty"`
		// AdminPassword - Specifies the password of the administrator account. **Minimum-length (Windows):** 8 characters **Minimum-length (Linux):** 6 characters **Max-length (Windows):** 123 characters **Max-length (Linux):** 72 characters **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled  Has lower characters Has upper characters  Has a digit  Has a special character (Regex match [\W_]) **Disallowed values:** "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!" For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
		AdminPassword *string `json:"adminPassword,omitempty"`
		// CustomData - Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes. **Note: Do not pass any secrets or passwords in customData property** This property cannot be updated after the VM is created. customData is passed to the VM to be saved as a file, for more information see [Custom Data on Azure VMs](https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/) For using cloud-init for your Linux VM, see [Using cloud-init to customize a Linux VM during creation](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
		CustomData *string `json:"customData,omitempty"`
		// LinuxConfiguration - Specifies the Linux operating system settings on the virtual machine. For a list of supported Linux distributions, see [Linux on Azure-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json) For running non-endorsed distributions, see [Information for Non-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
		LinuxConfiguration *LinuxConfiguration `json:"linuxConfiguration,omitempty"`
		// AllowExtensionOperations - Specifies whether extension operations should be allowed on the virtual machine. This may only be set to False when no extensions are present on the virtual machine.
		AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty"`
		// RequireGuestProvisionSignal - Specifies whether the guest provision signal is required to infer provision success of the virtual machine.
		RequireGuestProvisionSignal *bool `json:"requireGuestProvisionSignal,omitempty"`
	}

	LinuxConfiguration struct {
		// DisablePasswordAuthentication - Specifies whether password authentication should be disabled.
		DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty"`
		// SSH - Specifies the ssh key configuration for a Linux OS.
		SSH *SSHConfiguration `json:"ssh,omitempty"`
		// ProvisionVMAgent - Indicates whether virtual machine agent should be provisioned on the virtual machine. When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
		ProvisionVMAgent *bool `json:"provisionVMAgent,omitempty"`
	}

	SSHConfiguration struct {
		// PublicKeys - The list of SSH public keys used to authenticate with linux based VMs.
		PublicKeys *[]SSHPublicKey `json:"publicKeys,omitempty"`
	}

	SSHPublicKey struct {
		// Path - Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the specified key is appended to the file. Example: /home/user/.ssh/authorized_keys
		Path *string `json:"path,omitempty"`
		// KeyData - SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit and in ssh-rsa format. For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
		KeyData *string `json:"keyData,omitempty"`
	}

	NetworkProfile struct {
		// NetworkInterfaces - Specifies the list of resource Ids for the network interfaces associated with the virtual machine.
		NetworkInterfaceRefs *[]NetworkInterfaceReference `json:"networkInterfaceRefs,omitempty"`
	}

	NetworkInterfaceReference struct {
		*NetworkInterfaceReferenceProperties `json:"properties,omitempty"`
		azcorev1.KnownTypeReference          `json:""`
	}

	NetworkInterfaceReferenceProperties struct {
		Primary *bool `json:"primary,omitempty"`
	}

	DiagnosticsProfile struct {
		// BootDiagnostics - Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to diagnose VM status. You can easily view the output of your console log. Azure also enables you to see a screenshot of the VM from the hypervisor.
		BootDiagnostics *BootDiagnostics `json:"bootDiagnostics,omitempty"`
	}

	BootDiagnostics struct {
		// Enabled - Whether boot diagnostics should be enabled on the Virtual Machine.
		Enabled *bool `json:"enabled,omitempty"`
		// StorageURI - Uri of the storage account to use for placing the console output and screenshot.
		StorageURI *string `json:"storageUri,omitempty"`
	}

	VirtualMachineProperties struct {
		// HardwareProfile - Specifies the hardware settings for the virtual machine.
		HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`
		// StorageProfile - Specifies the storage settings for the virtual machine disks.
		StorageProfile *StorageProfile `json:"storageProfile,omitempty"`
		// OsProfile - Specifies the operating system settings for the virtual machine.
		OsProfile *OSProfile `json:"osProfile,omitempty"`
		// NetworkProfile - Specifies the network interfaces of the virtual machine.
		NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
		// DiagnosticsProfile - Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
		DiagnosticsProfile *DiagnosticsProfile `json:"diagnosticsProfile,omitempty"`
		// LicenseType - Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. Possible values are: Windows_Client Windows_Server If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) Minimum api-version: 2015-06-15
		LicenseType *string `json:"licenseType,omitempty"`
	}

	VirtualMachineIdentity struct {
		// Type - The type of identity used for the virtual machine. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
		// +kubebuilder:validation:Enum=SystemAssigned;UserAssigned;None
		Type string `json:"type"`
		// UserAssignedIdentities - The list of user identities associated with the Virtual Machine. The user identity references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
		UserAssignedIdentities []string `json:"userAssignedIdentities"`
	}

	// VirtualMachineSpec defines the desired state of VirtualMachine
	VirtualMachineSpec struct {
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef"`
		// Plan specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images.
		Plan *Plan `json:"plan,omitempty"`
		// Zones - The virtual machine zones.
		Zones *[]string `json:"zones,omitempty"`
		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`
		// Identity - The identity of the virtual machine, if configured.
		Identity *VirtualMachineIdentity `json:"identity,omitempty"`
		// Tags - Resource tags
		Tags map[string]string `json:"tags,omitempty"`
		// Properties contains details which describe the virtual machine
		Properties *VirtualMachineProperties `json:"properties,omitempty"`
	}

	// VirtualMachineStatus defines the observed state of VirtualMachine
	VirtualMachineStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// VirtualMachine is the Schema for the virtualmachines API
	VirtualMachine struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   VirtualMachineSpec   `json:"spec,omitempty"`
		Status VirtualMachineStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// VirtualMachineList contains a list of VirtualMachine
	VirtualMachineList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []VirtualMachine `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
