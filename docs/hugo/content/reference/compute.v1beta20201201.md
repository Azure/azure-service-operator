---
title: compute.azure.com/v1beta20201201
---
<h2 id="compute.azure.com/v1beta20201201">compute.azure.com/v1beta20201201</h2>
<div>
<p>Package v1beta20201201 contains API Schema definitions for the compute v1beta20201201 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="compute.azure.com/v1beta20201201.AdditionalCapabilities">AdditionalCapabilities
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalCapabilities">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalCapabilities</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ultraSSDEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UltraSSDEnabled: The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS
storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
machine or virtual machine scale set only if this property is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalCapabilitiesARM">AdditionalCapabilitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalCapabilities">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalCapabilities</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ultraSSDEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UltraSSDEnabled: The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS
storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
machine or virtual machine scale set only if this property is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalCapabilities_Status">AdditionalCapabilities_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ultraSSDEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UltraSSDEnabled: The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS
storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
machine or virtual machine scale set only if this property is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalCapabilities_StatusARM">AdditionalCapabilities_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ultraSSDEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UltraSSDEnabled: The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS
storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
machine or virtual machine scale set only if this property is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContent">AdditionalUnattendContent
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration">WindowsConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalUnattendContent">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalUnattendContent</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentComponentName">
AdditionalUnattendContentComponentName
</a>
</em>
</td>
<td>
<p>ComponentName: The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.</p>
</td>
</tr>
<tr>
<td>
<code>content</code><br/>
<em>
string
</em>
</td>
<td>
<p>Content: Specifies the XML formatted content that is added to the unattend.xml file for the specified path and
component. The XML must be less than 4KB and must include the root element for the setting or feature that is being
inserted.</p>
</td>
</tr>
<tr>
<td>
<code>passName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentPassName">
AdditionalUnattendContentPassName
</a>
</em>
</td>
<td>
<p>PassName: The pass name. Currently, the only allowable value is OobeSystem.</p>
</td>
</tr>
<tr>
<td>
<code>settingName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentSettingName">
AdditionalUnattendContentSettingName
</a>
</em>
</td>
<td>
<p>SettingName: Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and
AutoLogon.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentARM">AdditionalUnattendContentARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfigurationARM">WindowsConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalUnattendContent">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AdditionalUnattendContent</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentComponentName">
AdditionalUnattendContentComponentName
</a>
</em>
</td>
<td>
<p>ComponentName: The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.</p>
</td>
</tr>
<tr>
<td>
<code>content</code><br/>
<em>
string
</em>
</td>
<td>
<p>Content: Specifies the XML formatted content that is added to the unattend.xml file for the specified path and
component. The XML must be less than 4KB and must include the root element for the setting or feature that is being
inserted.</p>
</td>
</tr>
<tr>
<td>
<code>passName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentPassName">
AdditionalUnattendContentPassName
</a>
</em>
</td>
<td>
<p>PassName: The pass name. Currently, the only allowable value is OobeSystem.</p>
</td>
</tr>
<tr>
<td>
<code>settingName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentSettingName">
AdditionalUnattendContentSettingName
</a>
</em>
</td>
<td>
<p>SettingName: Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and
AutoLogon.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentComponentName">AdditionalUnattendContentComponentName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent">AdditionalUnattendContent</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentARM">AdditionalUnattendContentARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Microsoft-Windows-Shell-Setup&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentPassName">AdditionalUnattendContentPassName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent">AdditionalUnattendContent</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentARM">AdditionalUnattendContentARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OobeSystem&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentSettingName">AdditionalUnattendContentSettingName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent">AdditionalUnattendContent</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentARM">AdditionalUnattendContentARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutoLogon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FirstLogonCommands&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusComponentName">AdditionalUnattendContentStatusComponentName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_Status">AdditionalUnattendContent_Status</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_StatusARM">AdditionalUnattendContent_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Microsoft-Windows-Shell-Setup&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusPassName">AdditionalUnattendContentStatusPassName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_Status">AdditionalUnattendContent_Status</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_StatusARM">AdditionalUnattendContent_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OobeSystem&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusSettingName">AdditionalUnattendContentStatusSettingName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_Status">AdditionalUnattendContent_Status</a>, <a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_StatusARM">AdditionalUnattendContent_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutoLogon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FirstLogonCommands&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContent_Status">AdditionalUnattendContent_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_Status">WindowsConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusComponentName">
AdditionalUnattendContentStatusComponentName
</a>
</em>
</td>
<td>
<p>ComponentName: The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.</p>
</td>
</tr>
<tr>
<td>
<code>content</code><br/>
<em>
string
</em>
</td>
<td>
<p>Content: Specifies the XML formatted content that is added to the unattend.xml file for the specified path and
component. The XML must be less than 4KB and must include the root element for the setting or feature that is being
inserted.</p>
</td>
</tr>
<tr>
<td>
<code>passName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusPassName">
AdditionalUnattendContentStatusPassName
</a>
</em>
</td>
<td>
<p>PassName: The pass name. Currently, the only allowable value is OobeSystem.</p>
</td>
</tr>
<tr>
<td>
<code>settingName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusSettingName">
AdditionalUnattendContentStatusSettingName
</a>
</em>
</td>
<td>
<p>SettingName: Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and
AutoLogon.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AdditionalUnattendContent_StatusARM">AdditionalUnattendContent_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">WindowsConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusComponentName">
AdditionalUnattendContentStatusComponentName
</a>
</em>
</td>
<td>
<p>ComponentName: The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.</p>
</td>
</tr>
<tr>
<td>
<code>content</code><br/>
<em>
string
</em>
</td>
<td>
<p>Content: Specifies the XML formatted content that is added to the unattend.xml file for the specified path and
component. The XML must be less than 4KB and must include the root element for the setting or feature that is being
inserted.</p>
</td>
</tr>
<tr>
<td>
<code>passName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusPassName">
AdditionalUnattendContentStatusPassName
</a>
</em>
</td>
<td>
<p>PassName: The pass name. Currently, the only allowable value is OobeSystem.</p>
</td>
</tr>
<tr>
<td>
<code>settingName</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentStatusSettingName">
AdditionalUnattendContentStatusSettingName
</a>
</em>
</td>
<td>
<p>SettingName: Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and
AutoLogon.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiEntityReference">ApiEntityReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ApiEntityReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ApiEntityReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: The ARM resource id in the form of /&#x200b;subscriptions/&#x200b;{SubscriptionId}/&#x200b;resourceGroups/&#x200b;{ResourceGroupName}/&#x200b;&hellip;</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiEntityReferenceARM">ApiEntityReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ApiEntityReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ApiEntityReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiEntityReference_Status">ApiEntityReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">VirtualMachineScaleSetIPConfiguration_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_Status">VirtualMachineScaleSetNetworkProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The ARM resource id in the form of /&#x200b;subscriptions/&#x200b;{SubscriptionId}/&#x200b;resourceGroups/&#x200b;{ResourceGroupName}/&#x200b;&hellip;</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiEntityReference_StatusARM">ApiEntityReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">VirtualMachineScaleSetIPConfigurationProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_StatusARM">VirtualMachineScaleSetNetworkProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The ARM resource id in the form of /&#x200b;subscriptions/&#x200b;{SubscriptionId}/&#x200b;resourceGroups/&#x200b;{ResourceGroupName}/&#x200b;&hellip;</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiErrorBase_Status">ApiErrorBase_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ApiError_Status">ApiError_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The target of the particular error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiErrorBase_StatusARM">ApiErrorBase_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ApiError_StatusARM">ApiError_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The target of the particular error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiError_Status">ApiError_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_Status">AvailablePatchSummary_Status</a>, <a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_Status">LastPatchInstallationSummary_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiErrorBase_Status">
[]ApiErrorBase_Status
</a>
</em>
</td>
<td>
<p>Details: The Api error details</p>
</td>
</tr>
<tr>
<td>
<code>innererror</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InnerError_Status">
InnerError_Status
</a>
</em>
</td>
<td>
<p>Innererror: The Api inner error</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The target of the particular error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ApiError_StatusARM">ApiError_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_StatusARM">AvailablePatchSummary_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_StatusARM">LastPatchInstallationSummary_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiErrorBase_StatusARM">
[]ApiErrorBase_StatusARM
</a>
</em>
</td>
<td>
<p>Details: The Api error details</p>
</td>
</tr>
<tr>
<td>
<code>innererror</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InnerError_StatusARM">
InnerError_StatusARM
</a>
</em>
</td>
<td>
<p>Innererror: The Api inner error</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The target of the particular error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy">AutomaticOSUpgradePolicy
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy">UpgradePolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticOSUpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticOSUpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableAutomaticRollback</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAutomaticRollback: Whether OS image rollback feature should be disabled. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticOSUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticOSUpgrade: Indicates whether OS upgrades should automatically be applied to scale set instances in a
rolling fashion when a newer version of the OS image becomes available. Default value is false.
If this is set to true for Windows based scale sets,
<a href="https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet">enableAutomaticUpdates</a>
is automatically set to false and cannot be set to true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicyARM">AutomaticOSUpgradePolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicyARM">UpgradePolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticOSUpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticOSUpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableAutomaticRollback</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAutomaticRollback: Whether OS image rollback feature should be disabled. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticOSUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticOSUpgrade: Indicates whether OS upgrades should automatically be applied to scale set instances in a
rolling fashion when a newer version of the OS image becomes available. Default value is false.
If this is set to true for Windows based scale sets,
<a href="https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet">enableAutomaticUpdates</a>
is automatically set to false and cannot be set to true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy_Status">AutomaticOSUpgradePolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy_Status">UpgradePolicy_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableAutomaticRollback</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAutomaticRollback: Whether OS image rollback feature should be disabled. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticOSUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticOSUpgrade: Indicates whether OS upgrades should automatically be applied to scale set instances in a
rolling fashion when a newer version of the OS image becomes available. Default value is false.
If this is set to true for Windows based scale sets,
<a href="https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet">enableAutomaticUpdates</a>
is automatically set to false and cannot be set to true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy_StatusARM">AutomaticOSUpgradePolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy_StatusARM">UpgradePolicy_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableAutomaticRollback</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAutomaticRollback: Whether OS image rollback feature should be disabled. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticOSUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticOSUpgrade: Indicates whether OS upgrades should automatically be applied to scale set instances in a
rolling fashion when a newer version of the OS image becomes available. Default value is false.
If this is set to true for Windows based scale sets,
<a href="https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet">enableAutomaticUpdates</a>
is automatically set to false and cannot be set to true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticRepairsPolicy">AutomaticRepairsPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticRepairsPolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticRepairsPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is
false.</p>
</td>
</tr>
<tr>
<td>
<code>gracePeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>GracePeriod: The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time
starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should
be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default
value. The maximum allowed grace period is 90 minutes (PT90M).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticRepairsPolicyARM">AutomaticRepairsPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticRepairsPolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/AutomaticRepairsPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is
false.</p>
</td>
</tr>
<tr>
<td>
<code>gracePeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>GracePeriod: The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time
starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should
be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default
value. The maximum allowed grace period is 90 minutes (PT90M).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticRepairsPolicy_Status">AutomaticRepairsPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is
false.</p>
</td>
</tr>
<tr>
<td>
<code>gracePeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>GracePeriod: The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time
starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should
be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default
value. The maximum allowed grace period is 90 minutes (PT90M).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AutomaticRepairsPolicy_StatusARM">AutomaticRepairsPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is
false.</p>
</td>
</tr>
<tr>
<td>
<code>gracePeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>GracePeriod: The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time
starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should
be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default
value. The maximum allowed grace period is 90 minutes (PT90M).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AvailablePatchSummaryStatusStatus">AvailablePatchSummaryStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_Status">AvailablePatchSummary_Status</a>, <a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_StatusARM">AvailablePatchSummary_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CompletedWithWarnings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;InProgress&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unknown&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AvailablePatchSummary_Status">AvailablePatchSummary_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_Status">VirtualMachinePatchStatus_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>assessmentActivityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AssessmentActivityId: The activity ID of the operation that produced this result. It is used to correlate across CRP and
extension logs.</p>
</td>
</tr>
<tr>
<td>
<code>criticalAndSecurityPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>CriticalAndSecurityPatchCount: The number of critical or security patches that have been detected as available and not
yet installed.</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiError_Status">
ApiError_Status
</a>
</em>
</td>
<td>
<p>Error: The errors that were encountered during execution of the operation. The details array contains the list of them.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>otherPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>OtherPatchCount: The number of all available patches excluding critical and security.</p>
</td>
</tr>
<tr>
<td>
<code>rebootPending</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RebootPending: The overall reboot status of the VM. It will be true when partially installed patches require a reboot to
complete installation but the reboot has not yet occurred.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AvailablePatchSummaryStatusStatus">
AvailablePatchSummaryStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The overall success or failure status of the operation. It remains &ldquo;InProgress&rdquo; until the operation completes.
At that point it will become &ldquo;Unknown&rdquo;, &ldquo;Failed&rdquo;, &ldquo;Succeeded&rdquo;, or &ldquo;CompletedWithWarnings.&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.AvailablePatchSummary_StatusARM">AvailablePatchSummary_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_StatusARM">VirtualMachinePatchStatus_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>assessmentActivityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AssessmentActivityId: The activity ID of the operation that produced this result. It is used to correlate across CRP and
extension logs.</p>
</td>
</tr>
<tr>
<td>
<code>criticalAndSecurityPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>CriticalAndSecurityPatchCount: The number of critical or security patches that have been detected as available and not
yet installed.</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiError_StatusARM">
ApiError_StatusARM
</a>
</em>
</td>
<td>
<p>Error: The errors that were encountered during execution of the operation. The details array contains the list of them.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>otherPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>OtherPatchCount: The number of all available patches excluding critical and security.</p>
</td>
</tr>
<tr>
<td>
<code>rebootPending</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RebootPending: The overall reboot status of the VM. It will be true when partially installed patches require a reboot to
complete installation but the reboot has not yet occurred.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AvailablePatchSummaryStatusStatus">
AvailablePatchSummaryStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The overall success or failure status of the operation. It remains &ldquo;InProgress&rdquo; until the operation completes.
At that point it will become &ldquo;Unknown&rdquo;, &ldquo;Failed&rdquo;, &ldquo;Succeeded&rdquo;, or &ldquo;CompletedWithWarnings.&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BillingProfile">BillingProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BillingProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BillingProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>MaxPrice: Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.
This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time
of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current
Azure Spot price.
The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the
maxPrice after creation of VM/VMSS.
Possible values are:
- Any decimal value greater than zero. Example: 0.01538
-1 – indicates default price to be up-to on-demand.
You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also,
the default max price is -1 if it is not provided by you.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BillingProfileARM">BillingProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BillingProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BillingProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>MaxPrice: Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.
This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time
of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current
Azure Spot price.
The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the
maxPrice after creation of VM/VMSS.
Possible values are:
- Any decimal value greater than zero. Example: 0.01538
-1 – indicates default price to be up-to on-demand.
You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also,
the default max price is -1 if it is not provided by you.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BillingProfile_Status">BillingProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>MaxPrice: Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.
This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time
of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current
Azure Spot price.
The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the
maxPrice after creation of VM/VMSS.
Possible values are:
- Any decimal value greater than zero. Example: 0.01538
-1 – indicates default price to be up-to on-demand.
You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also,
the default max price is -1 if it is not provided by you.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BillingProfile_StatusARM">BillingProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>MaxPrice: Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS. This price is in US Dollars.
This price will be compared with the current Azure Spot price for the VM size. Also, the prices are compared at the time
of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current
Azure Spot price.
The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the
maxPrice after creation of VM/VMSS.
Possible values are:
- Any decimal value greater than zero. Example: 0.01538
-1 – indicates default price to be up-to on-demand.
You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons. Also,
the default max price is -1 if it is not provided by you.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnostics">BootDiagnostics
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile">DiagnosticsProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BootDiagnostics">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BootDiagnostics</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether boot diagnostics should be enabled on the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageUri: Uri of the storage account to use for placing the console output and screenshot.
If storageUri is not specified while enabling boot diagnostics, managed storage will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnosticsARM">BootDiagnosticsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiagnosticsProfileARM">DiagnosticsProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BootDiagnostics">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/BootDiagnostics</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether boot diagnostics should be enabled on the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageUri: Uri of the storage account to use for placing the console output and screenshot.
If storageUri is not specified while enabling boot diagnostics, managed storage will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_Status">BootDiagnosticsInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>consoleScreenshotBlobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConsoleScreenshotBlobUri: The console screenshot blob URI.
NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.</p>
</td>
</tr>
<tr>
<td>
<code>serialConsoleLogBlobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SerialConsoleLogBlobUri: The serial console log blob Uri.
NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Status: The boot diagnostics status information for the VM.
NOTE: It will be set only if there are errors encountered in enabling boot diagnostics.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_StatusARM">BootDiagnosticsInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>consoleScreenshotBlobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConsoleScreenshotBlobUri: The console screenshot blob URI.
NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.</p>
</td>
</tr>
<tr>
<td>
<code>serialConsoleLogBlobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SerialConsoleLogBlobUri: The serial console log blob Uri.
NOTE: This will not be set if boot diagnostics is currently enabled with managed storage.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Status: The boot diagnostics status information for the VM.
NOTE: It will be set only if there are errors encountered in enabling boot diagnostics.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnostics_Status">BootDiagnostics_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_Status">DiagnosticsProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether boot diagnostics should be enabled on the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageUri: Uri of the storage account to use for placing the console output and screenshot.
If storageUri is not specified while enabling boot diagnostics, managed storage will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.BootDiagnostics_StatusARM">BootDiagnostics_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_StatusARM">DiagnosticsProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether boot diagnostics should be enabled on the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageUri: Uri of the storage account to use for placing the console output and screenshot.
If storageUri is not specified while enabling boot diagnostics, managed storage will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Caching_Status">Caching_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_Status">VirtualMachineScaleSetDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_StatusARM">VirtualMachineScaleSetDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.CreateOption_Status">CreateOption_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_Status">VirtualMachineScaleSetDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_StatusARM">VirtualMachineScaleSetDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDisk">DataDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile">StorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DataDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DataDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskCaching">
DataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskCreateOption">
DataDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>detachOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskDetachOption">
DataDiskDetachOption
</a>
</em>
</td>
<td>
<p>DetachOption: Specifies the detach behavior to be used while detaching a disk or which is already in the process of
detachment from the virtual machine. Supported values: ForceDetach.
detachOption: ForceDetach is applicable only for managed data disks. If a previous detachment attempt of the data disk
did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use
force-detach as a last resort option to detach the disk forcibly from the VM. All writes might not have been flushed
when using this detach behavior.
This feature is still in preview mode and is not supported for VirtualMachineScaleSet. To force-detach a data disk
update toBeDetached to &lsquo;true&rsquo; along with setting detachOption: &lsquo;ForceDetach&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk">
VirtualHardDisk
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters">
ManagedDiskParameters
</a>
</em>
</td>
<td>
<p>ManagedDisk: The parameters of a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>toBeDetached</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ToBeDetached: Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk">
VirtualHardDisk
</a>
</em>
</td>
<td>
<p>Vhd: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfileARM">StorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DataDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DataDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskCaching">
DataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskCreateOption">
DataDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>detachOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskDetachOption">
DataDiskDetachOption
</a>
</em>
</td>
<td>
<p>DetachOption: Specifies the detach behavior to be used while detaching a disk or which is already in the process of
detachment from the virtual machine. Supported values: ForceDetach.
detachOption: ForceDetach is applicable only for managed data disks. If a previous detachment attempt of the data disk
did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use
force-detach as a last resort option to detach the disk forcibly from the VM. All writes might not have been flushed
when using this detach behavior.
This feature is still in preview mode and is not supported for VirtualMachineScaleSet. To force-detach a data disk
update toBeDetached to &lsquo;true&rsquo; along with setting detachOption: &lsquo;ForceDetach&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDiskARM">
VirtualHardDiskARM
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersARM">
ManagedDiskParametersARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The parameters of a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>toBeDetached</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ToBeDetached: Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDiskARM">
VirtualHardDiskARM
</a>
</em>
</td>
<td>
<p>Vhd: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDiskCaching">DataDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk">DataDisk</a>, <a href="#compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDiskCreateOption">DataDiskCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk">DataDisk</a>, <a href="#compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDiskDetachOption">DataDiskDetachOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk">DataDisk</a>, <a href="#compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ForceDetach&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_Status">StorageProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>detachOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DetachOption_Status">
DetachOption_Status
</a>
</em>
</td>
<td>
<p>DetachOption: Specifies the detach behavior to be used while detaching a disk or which is already in the process of
detachment from the virtual machine. Supported values: ForceDetach.
detachOption: ForceDetach is applicable only for managed data disks. If a previous detachment attempt of the data disk
did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use
force-detach as a last resort option to detach the disk forcibly from the VM. All writes might not have been flushed
when using this detach behavior.
This feature is still in preview mode and is not supported for VirtualMachineScaleSet. To force-detach a data disk
update toBeDetached to &lsquo;true&rsquo; along with setting detachOption: &lsquo;ForceDetach&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk when StorageAccountType is UltraSSD_LRS. Returned
only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the VirtualMachine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk when StorageAccountType is
UltraSSD_LRS. Returned only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the VirtualMachine
Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_Status">
VirtualHardDisk_Status
</a>
</em>
</td>
<td>
<p>Image: The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_Status">
ManagedDiskParameters_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>toBeDetached</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ToBeDetached: Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_Status">
VirtualHardDisk_Status
</a>
</em>
</td>
<td>
<p>Vhd: The virtual hard disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_StatusARM">StorageProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>detachOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DetachOption_Status">
DetachOption_Status
</a>
</em>
</td>
<td>
<p>DetachOption: Specifies the detach behavior to be used while detaching a disk or which is already in the process of
detachment from the virtual machine. Supported values: ForceDetach.
detachOption: ForceDetach is applicable only for managed data disks. If a previous detachment attempt of the data disk
did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use
force-detach as a last resort option to detach the disk forcibly from the VM. All writes might not have been flushed
when using this detach behavior.
This feature is still in preview mode and is not supported for VirtualMachineScaleSet. To force-detach a data disk
update toBeDetached to &lsquo;true&rsquo; along with setting detachOption: &lsquo;ForceDetach&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk when StorageAccountType is UltraSSD_LRS. Returned
only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the VirtualMachine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk when StorageAccountType is
UltraSSD_LRS. Returned only for VirtualMachine ScaleSet VM disks. Can be updated only via updates to the VirtualMachine
Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">
VirtualHardDisk_StatusARM
</a>
</em>
</td>
<td>
<p>Image: The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_StatusARM">
ManagedDiskParameters_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>toBeDetached</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ToBeDetached: Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">
VirtualHardDisk_StatusARM
</a>
</em>
</td>
<td>
<p>Vhd: The virtual hard disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DetachOption_Status">DetachOption_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ForceDetach&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiagnosticsProfile">DiagnosticsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiagnosticsProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiagnosticsProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnostics">
BootDiagnostics
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiagnosticsProfileARM">DiagnosticsProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiagnosticsProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiagnosticsProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnosticsARM">
BootDiagnosticsARM
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiagnosticsProfile_Status">DiagnosticsProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnostics_Status">
BootDiagnostics_Status
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiagnosticsProfile_StatusARM">DiagnosticsProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnostics_StatusARM">
BootDiagnostics_StatusARM
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskOption_Status">DiffDiskOption_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_Status">DiffDiskSettings_Status</a>, <a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_StatusARM">DiffDiskSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Local&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskPlacement_Status">DiffDiskPlacement_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_Status">DiffDiskSettings_Status</a>, <a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_StatusARM">DiffDiskSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CacheDisk&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ResourceDisk&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettings">DiffDiskSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiffDiskSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiffDiskSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsOption">
DiffDiskSettingsOption
</a>
</em>
</td>
<td>
<p>Option: Specifies the ephemeral disk settings for operating system disk.</p>
</td>
</tr>
<tr>
<td>
<code>placement</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsPlacement">
DiffDiskSettingsPlacement
</a>
</em>
</td>
<td>
<p>Placement: Specifies the ephemeral disk placement for operating system disk.
Possible values are:
CacheDisk
ResourceDisk
Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk is used.
Refer to VM size documentation for Windows VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes</a>
and Linux VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes</a> to check which VM sizes exposes a
cache disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettingsARM">DiffDiskSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiffDiskSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiffDiskSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsOption">
DiffDiskSettingsOption
</a>
</em>
</td>
<td>
<p>Option: Specifies the ephemeral disk settings for operating system disk.</p>
</td>
</tr>
<tr>
<td>
<code>placement</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsPlacement">
DiffDiskSettingsPlacement
</a>
</em>
</td>
<td>
<p>Placement: Specifies the ephemeral disk placement for operating system disk.
Possible values are:
CacheDisk
ResourceDisk
Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk is used.
Refer to VM size documentation for Windows VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes</a>
and Linux VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes</a> to check which VM sizes exposes a
cache disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettingsOption">DiffDiskSettingsOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiffDiskSettings">DiffDiskSettings</a>, <a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsARM">DiffDiskSettingsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Local&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettingsPlacement">DiffDiskSettingsPlacement
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiffDiskSettings">DiffDiskSettings</a>, <a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsARM">DiffDiskSettingsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CacheDisk&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ResourceDisk&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettings_Status">DiffDiskSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskOption_Status">
DiffDiskOption_Status
</a>
</em>
</td>
<td>
<p>Option: Specifies the ephemeral disk settings for operating system disk.</p>
</td>
</tr>
<tr>
<td>
<code>placement</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskPlacement_Status">
DiffDiskPlacement_Status
</a>
</em>
</td>
<td>
<p>Placement: Specifies the ephemeral disk placement for operating system disk.
Possible values are:
CacheDisk
ResourceDisk
Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk is used.
Refer to VM size documentation for Windows VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes</a>
and Linux VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes</a> to check which VM sizes exposes a
cache disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiffDiskSettings_StatusARM">DiffDiskSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskOption_Status">
DiffDiskOption_Status
</a>
</em>
</td>
<td>
<p>Option: Specifies the ephemeral disk settings for operating system disk.</p>
</td>
</tr>
<tr>
<td>
<code>placement</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskPlacement_Status">
DiffDiskPlacement_Status
</a>
</em>
</td>
<td>
<p>Placement: Specifies the ephemeral disk placement for operating system disk.
Possible values are:
CacheDisk
ResourceDisk
Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk is used.
Refer to VM size documentation for Windows VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes</a>
and Linux VM at <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes">https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes</a> to check which VM sizes exposes a
cache disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSetParameters">DiskEncryptionSetParameters
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters">ManagedDiskParameters</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters">VirtualMachineScaleSetManagedDiskParameters</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSetParametersARM">DiskEncryptionSetParametersARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersARM">ManagedDiskParametersARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersARM">VirtualMachineScaleSetManagedDiskParametersARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSettings">DiskEncryptionSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference">
KeyVaultSecretReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Describes a reference to Key Vault Secret</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether disk encryption should be enabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference">
KeyVaultKeyReference
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Describes a reference to Key Vault Key</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSettingsARM">DiskEncryptionSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/DiskEncryptionSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReferenceARM">
KeyVaultSecretReferenceARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Describes a reference to Key Vault Secret</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether disk encryption should be enabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReferenceARM">
KeyVaultKeyReferenceARM
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Describes a reference to Key Vault Key</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSettings_Status">DiskEncryptionSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskInstanceView_Status">DiskInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference_Status">
KeyVaultSecretReference_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Specifies the location of the disk encryption key, which is a Key Vault Secret.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether disk encryption should be enabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference_Status">
KeyVaultKeyReference_Status
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Specifies the location of the key encryption key in Key Vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskEncryptionSettings_StatusARM">DiskEncryptionSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskInstanceView_StatusARM">DiskInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference_StatusARM">
KeyVaultSecretReference_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Specifies the location of the disk encryption key, which is a Key Vault Secret.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Specifies whether disk encryption should be enabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference_StatusARM">
KeyVaultKeyReference_StatusARM
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Specifies the location of the key encryption key in Key Vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskInstanceView_Status">DiskInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_Status">
[]DiskEncryptionSettings_Status
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Specifies the encryption settings for the OS Disk.
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.DiskInstanceView_StatusARM">DiskInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_StatusARM">
[]DiskEncryptionSettings_StatusARM
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Specifies the encryption settings for the OS Disk.
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.EvictionPolicy_Status">EvictionPolicy_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_SpecARM">VirtualMachines_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ExtendedLocation">ExtendedLocation</a>, <a href="#compute.azure.com/v1beta20201201.ExtendedLocationARM">ExtendedLocationARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#compute.azure.com/v1beta20201201.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.GenericExtensionARM">GenericExtensionARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.Extensions.json#/definitions/genericExtension">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.Extensions.json#/definitions/genericExtension</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: Microsoft.Compute/extensions - Publisher</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Microsoft.Compute/extensions - Settings</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Microsoft.Compute/extensions - Type</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Microsoft.Compute/extensions - Type handler version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfile">HardwareProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/HardwareProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/HardwareProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vmSize</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfileVmSize">
HardwareProfileVmSize
</a>
</em>
</td>
<td>
<p>VmSize: Specifies the size of the virtual machine.
The enum data type is currently deprecated and will be removed by December 23rd 2023.
Recommended way to get the list of available sizes is using these APIs:
<a href="https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes">List all available virtual machine sizes in an availability
set</a>
<a href="https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list">List all available virtual machine sizes in a region</a>
<a href="https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes">List all available virtual machine sizes for
resizing</a>. For more information about
virtual machine sizes, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/sizes">Sizes for virtual machines</a>.
The available VM sizes depend on region and availability set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfileARM">HardwareProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/HardwareProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/HardwareProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vmSize</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfileVmSize">
HardwareProfileVmSize
</a>
</em>
</td>
<td>
<p>VmSize: Specifies the size of the virtual machine.
The enum data type is currently deprecated and will be removed by December 23rd 2023.
Recommended way to get the list of available sizes is using these APIs:
<a href="https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes">List all available virtual machine sizes in an availability
set</a>
<a href="https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list">List all available virtual machine sizes in a region</a>
<a href="https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes">List all available virtual machine sizes for
resizing</a>. For more information about
virtual machine sizes, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/sizes">Sizes for virtual machines</a>.
The available VM sizes depend on region and availability set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfileStatusVmSize">HardwareProfileStatusVmSize
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.HardwareProfile_Status">HardwareProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.HardwareProfile_StatusARM">HardwareProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Basic_A0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A7&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A9&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B1ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B2ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B2s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B4ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B8ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D11_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D12_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D13&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D13_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D14_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D15_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D32_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D3_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D5_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D64s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D64_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D8_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS11_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS12_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13-2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13-4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14-4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14-8_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS15_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS3_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS5_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E2s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E2_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32-16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32-8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E4s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E4_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64-16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64-32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E8_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F32s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F64s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F72s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4-4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4-8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5-16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5-8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16m&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16mr&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16r&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H8m&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L16s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L32s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L4s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L8s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128-32ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128-64ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64-16ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64-32ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24r&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24rs_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24rs_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND12s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND24rs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND24s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND6s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfileVmSize">HardwareProfileVmSize
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.HardwareProfile">HardwareProfile</a>, <a href="#compute.azure.com/v1beta20201201.HardwareProfileARM">HardwareProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Basic_A0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Basic_A4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A7&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8m_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A8_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_A9&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B1ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B2ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B2s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B4ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_B8ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D11_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D12_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D13&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D13_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D14_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D15_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D2_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D32_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D3_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D4_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D5_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D64s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D64_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_D8_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS11_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS12_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13-2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13-4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS13_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14-4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14-8_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS14_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS15_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS1_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS2_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS3_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS4_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_DS5_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E2s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E2_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32-16_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32-8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E32_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E4s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E4_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64-16s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64-32s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E64_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E8s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_E8_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F16s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F2s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F32s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F4s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F64s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F72s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_F8s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_G5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4-4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS4-8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5-16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_GS5-8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16m&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16mr&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H16r&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H8&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_H8m&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L16s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L32s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L4s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_L8s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128-32ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128-64ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M128s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64-16ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64-32ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64ms&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_M64s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC12s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24r&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24rs_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24rs_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC24s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6s_v2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NC6s_v3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND12s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND24rs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND24s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ND6s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_NV6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfile_Status">HardwareProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vmSize</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfileStatusVmSize">
HardwareProfileStatusVmSize
</a>
</em>
</td>
<td>
<p>VmSize: Specifies the size of the virtual machine.
The enum data type is currently deprecated and will be removed by December 23rd 2023.
Recommended way to get the list of available sizes is using these APIs:
<a href="https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes">List all available virtual machine sizes in an availability
set</a>
<a href="https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list">List all available virtual machine sizes in a region</a>
<a href="https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes">List all available virtual machine sizes for
resizing</a>. For more information about
virtual machine sizes, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/sizes">Sizes for virtual machines</a>.
The available VM sizes depend on region and availability set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.HardwareProfile_StatusARM">HardwareProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vmSize</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfileStatusVmSize">
HardwareProfileStatusVmSize
</a>
</em>
</td>
<td>
<p>VmSize: Specifies the size of the virtual machine.
The enum data type is currently deprecated and will be removed by December 23rd 2023.
Recommended way to get the list of available sizes is using these APIs:
<a href="https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes">List all available virtual machine sizes in an availability
set</a>
<a href="https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list">List all available virtual machine sizes in a region</a>
<a href="https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes">List all available virtual machine sizes for
resizing</a>. For more information about
virtual machine sizes, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/sizes">Sizes for virtual machines</a>.
The available VM sizes depend on region and availability set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ImageReference">ImageReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile">StorageProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile">VirtualMachineScaleSetStorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ImageReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ImageReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>offer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Offer: Specifies the offer of the platform image or marketplace image used to create the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The image publisher.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sku: The image SKU.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: Specifies the version of the platform image or marketplace image used to create the virtual machine. The
allowed formats are Major.Minor.Build or &lsquo;latest&rsquo;. Major, Minor, and Build are decimal numbers. Specify &lsquo;latest&rsquo; to use
the latest version of an image available at deploy time. Even if you use &lsquo;latest&rsquo;, the VM image will not automatically
update after deploy time even if a new version becomes available.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ImageReferenceARM">ImageReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfileARM">StorageProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfileARM">VirtualMachineScaleSetStorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ImageReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ImageReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>offer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Offer: Specifies the offer of the platform image or marketplace image used to create the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The image publisher.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sku: The image SKU.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: Specifies the version of the platform image or marketplace image used to create the virtual machine. The
allowed formats are Major.Minor.Build or &lsquo;latest&rsquo;. Major, Minor, and Build are decimal numbers. Specify &lsquo;latest&rsquo; to use
the latest version of an image available at deploy time. Even if you use &lsquo;latest&rsquo;, the VM image will not automatically
update after deploy time even if a new version becomes available.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ImageReference_Status">ImageReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_Status">StorageProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_Status">VirtualMachineScaleSetStorageProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exactVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExactVersion: Specifies in decimal numbers, the version of platform image or marketplace image used to create the
virtual machine. This readonly field differs from &lsquo;version&rsquo;, only if the value specified in &lsquo;version&rsquo; field is &lsquo;latest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>offer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Offer: Specifies the offer of the platform image or marketplace image used to create the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The image publisher.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sku: The image SKU.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: Specifies the version of the platform image or marketplace image used to create the virtual machine. The
allowed formats are Major.Minor.Build or &lsquo;latest&rsquo;. Major, Minor, and Build are decimal numbers. Specify &lsquo;latest&rsquo; to use
the latest version of an image available at deploy time. Even if you use &lsquo;latest&rsquo;, the VM image will not automatically
update after deploy time even if a new version becomes available.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ImageReference_StatusARM">ImageReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_StatusARM">StorageProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_StatusARM">VirtualMachineScaleSetStorageProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exactVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExactVersion: Specifies in decimal numbers, the version of platform image or marketplace image used to create the
virtual machine. This readonly field differs from &lsquo;version&rsquo;, only if the value specified in &lsquo;version&rsquo; field is &lsquo;latest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>offer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Offer: Specifies the offer of the platform image or marketplace image used to create the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The image publisher.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sku: The image SKU.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: Specifies the version of the platform image or marketplace image used to create the virtual machine. The
allowed formats are Major.Minor.Build or &lsquo;latest&rsquo;. Major, Minor, and Build are decimal numbers. Specify &lsquo;latest&rsquo; to use
the latest version of an image available at deploy time. Even if you use &lsquo;latest&rsquo;, the VM image will not automatically
update after deploy time even if a new version becomes available.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.InnerError_Status">InnerError_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ApiError_Status">ApiError_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>errordetail</code><br/>
<em>
string
</em>
</td>
<td>
<p>Errordetail: The internal error message or exception dump.</p>
</td>
</tr>
<tr>
<td>
<code>exceptiontype</code><br/>
<em>
string
</em>
</td>
<td>
<p>Exceptiontype: The exception type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.InnerError_StatusARM">InnerError_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ApiError_StatusARM">ApiError_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>errordetail</code><br/>
<em>
string
</em>
</td>
<td>
<p>Errordetail: The internal error message or exception dump.</p>
</td>
</tr>
<tr>
<td>
<code>exceptiontype</code><br/>
<em>
string
</em>
</td>
<td>
<p>Exceptiontype: The exception type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.InstanceViewStatusStatusLevel">InstanceViewStatusStatusLevel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">InstanceViewStatus_Status</a>, <a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">InstanceViewStatus_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Error&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Info&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Warning&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.InstanceViewStatus_Status">InstanceViewStatus_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_Status">BootDiagnosticsInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.DiskInstanceView_Status">DiskInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_Status">VirtualMachineAgentInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_Status">VirtualMachineExtensionHandlerInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_Status">VirtualMachineExtensionInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_Status">VirtualMachineHealthStatus_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_Status">VirtualMachinePatchStatus_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The status code.</p>
</td>
</tr>
<tr>
<td>
<code>displayStatus</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayStatus: The short localizable label for the status.</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatusStatusLevel">
InstanceViewStatusStatusLevel
</a>
</em>
</td>
<td>
<p>Level: The level code.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The detailed status message, including for alerts and error messages.</p>
</td>
</tr>
<tr>
<td>
<code>time</code><br/>
<em>
string
</em>
</td>
<td>
<p>Time: The time of the status.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">InstanceViewStatus_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_StatusARM">BootDiagnosticsInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.DiskInstanceView_StatusARM">DiskInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_StatusARM">VirtualMachineAgentInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_StatusARM">VirtualMachineExtensionHandlerInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_StatusARM">VirtualMachineExtensionInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_StatusARM">VirtualMachineHealthStatus_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_StatusARM">VirtualMachinePatchStatus_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The status code.</p>
</td>
</tr>
<tr>
<td>
<code>displayStatus</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayStatus: The short localizable label for the status.</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatusStatusLevel">
InstanceViewStatusStatusLevel
</a>
</em>
</td>
<td>
<p>Level: The level code.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: The detailed status message, including for alerts and error messages.</p>
</td>
</tr>
<tr>
<td>
<code>time</code><br/>
<em>
string
</em>
</td>
<td>
<p>Time: The time of the status.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultKeyReference">KeyVaultKeyReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings">DiskEncryptionSettings</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultKeyReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultKeyReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: The URL referencing a key encryption key in Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultKeyReferenceARM">KeyVaultKeyReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettingsARM">DiskEncryptionSettingsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultKeyReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultKeyReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: The URL referencing a key encryption key in Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultKeyReference_Status">KeyVaultKeyReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_Status">DiskEncryptionSettings_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: The URL referencing a key encryption key in Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing the key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultKeyReference_StatusARM">KeyVaultKeyReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_StatusARM">DiskEncryptionSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: The URL referencing a key encryption key in Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing the key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultSecretReference">KeyVaultSecretReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings">DiskEncryptionSettings</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultSecretReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultSecretReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: The URL referencing a secret in a Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultSecretReferenceARM">KeyVaultSecretReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettingsARM">DiskEncryptionSettingsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultSecretReference">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/KeyVaultSecretReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: The URL referencing a secret in a Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultSecretReference_Status">KeyVaultSecretReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_Status">DiskEncryptionSettings_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: The URL referencing a secret in a Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing the secret.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.KeyVaultSecretReference_StatusARM">KeyVaultSecretReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_StatusARM">DiskEncryptionSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: The URL referencing a secret in a Key Vault.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing the secret.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LastPatchInstallationSummaryStatusStatus">LastPatchInstallationSummaryStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_Status">LastPatchInstallationSummary_Status</a>, <a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_StatusARM">LastPatchInstallationSummary_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CompletedWithWarnings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;InProgress&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unknown&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LastPatchInstallationSummary_Status">LastPatchInstallationSummary_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_Status">VirtualMachinePatchStatus_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiError_Status">
ApiError_Status
</a>
</em>
</td>
<td>
<p>Error: The errors that were encountered during execution of the operation. The details array contains the list of them.</p>
</td>
</tr>
<tr>
<td>
<code>excludedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExcludedPatchCount: The number of all available patches but excluded explicitly by a customer-specified exclusion list
match.</p>
</td>
</tr>
<tr>
<td>
<code>failedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailedPatchCount: The count of patches that failed installation.</p>
</td>
</tr>
<tr>
<td>
<code>installationActivityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>InstallationActivityId: The activity ID of the operation that produced this result. It is used to correlate across CRP
and extension logs.</p>
</td>
</tr>
<tr>
<td>
<code>installedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>InstalledPatchCount: The count of patches that successfully installed.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowExceeded</code><br/>
<em>
bool
</em>
</td>
<td>
<p>MaintenanceWindowExceeded: Describes whether the operation ran out of time before it completed all its intended actions</p>
</td>
</tr>
<tr>
<td>
<code>notSelectedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>NotSelectedPatchCount: The number of all available patches but not going to be installed because it didn&rsquo;t match a
classification or inclusion list entry.</p>
</td>
</tr>
<tr>
<td>
<code>pendingPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PendingPatchCount: The number of all available patches expected to be installed over the course of the patch
installation operation.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummaryStatusStatus">
LastPatchInstallationSummaryStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The overall success or failure status of the operation. It remains &ldquo;InProgress&rdquo; until the operation completes.
At that point it will become &ldquo;Unknown&rdquo;, &ldquo;Failed&rdquo;, &ldquo;Succeeded&rdquo;, or &ldquo;CompletedWithWarnings.&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LastPatchInstallationSummary_StatusARM">LastPatchInstallationSummary_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_StatusARM">VirtualMachinePatchStatus_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiError_StatusARM">
ApiError_StatusARM
</a>
</em>
</td>
<td>
<p>Error: The errors that were encountered during execution of the operation. The details array contains the list of them.</p>
</td>
</tr>
<tr>
<td>
<code>excludedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExcludedPatchCount: The number of all available patches but excluded explicitly by a customer-specified exclusion list
match.</p>
</td>
</tr>
<tr>
<td>
<code>failedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailedPatchCount: The count of patches that failed installation.</p>
</td>
</tr>
<tr>
<td>
<code>installationActivityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>InstallationActivityId: The activity ID of the operation that produced this result. It is used to correlate across CRP
and extension logs.</p>
</td>
</tr>
<tr>
<td>
<code>installedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>InstalledPatchCount: The count of patches that successfully installed.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowExceeded</code><br/>
<em>
bool
</em>
</td>
<td>
<p>MaintenanceWindowExceeded: Describes whether the operation ran out of time before it completed all its intended actions</p>
</td>
</tr>
<tr>
<td>
<code>notSelectedPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>NotSelectedPatchCount: The number of all available patches but not going to be installed because it didn&rsquo;t match a
classification or inclusion list entry.</p>
</td>
</tr>
<tr>
<td>
<code>pendingPatchCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PendingPatchCount: The number of all available patches expected to be installed over the course of the patch
installation operation.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The UTC timestamp when the operation began.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummaryStatusStatus">
LastPatchInstallationSummaryStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The overall success or failure status of the operation. It remains &ldquo;InProgress&rdquo; until the operation completes.
At that point it will become &ldquo;Unknown&rdquo;, &ldquo;Failed&rdquo;, &ldquo;Succeeded&rdquo;, or &ldquo;CompletedWithWarnings.&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxConfiguration">LinuxConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile">OSProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile">VirtualMachineScaleSetOSProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disablePasswordAuthentication</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisablePasswordAuthentication: Specifies whether password authentication should be disabled.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings">
LinuxPatchSettings
</a>
</em>
</td>
<td>
<p>PatchSettings: Specifies settings related to VM Guest Patching on Linux.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshConfiguration">
SshConfiguration
</a>
</em>
</td>
<td>
<p>Ssh: SSH configuration for Linux based VMs running on Azure</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxConfigurationARM">LinuxConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfileARM">OSProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfileARM">VirtualMachineScaleSetOSProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disablePasswordAuthentication</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisablePasswordAuthentication: Specifies whether password authentication should be disabled.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsARM">
LinuxPatchSettingsARM
</a>
</em>
</td>
<td>
<p>PatchSettings: Specifies settings related to VM Guest Patching on Linux.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshConfigurationARM">
SshConfigurationARM
</a>
</em>
</td>
<td>
<p>Ssh: SSH configuration for Linux based VMs running on Azure</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxConfiguration_Status">LinuxConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_Status">OSProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_Status">VirtualMachineScaleSetOSProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disablePasswordAuthentication</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisablePasswordAuthentication: Specifies whether password authentication should be disabled.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings_Status">
LinuxPatchSettings_Status
</a>
</em>
</td>
<td>
<p>PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching on Linux.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshConfiguration_Status">
SshConfiguration_Status
</a>
</em>
</td>
<td>
<p>Ssh: Specifies the ssh key configuration for a Linux OS.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxConfiguration_StatusARM">LinuxConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_StatusARM">OSProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_StatusARM">VirtualMachineScaleSetOSProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disablePasswordAuthentication</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisablePasswordAuthentication: Specifies whether password authentication should be disabled.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings_StatusARM">
LinuxPatchSettings_StatusARM
</a>
</em>
</td>
<td>
<p>PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching on Linux.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshConfiguration_StatusARM">
SshConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>Ssh: Specifies the ssh key configuration for a Linux OS.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettings">LinuxPatchSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration">LinuxConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxPatchSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxPatchSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsPatchMode">
LinuxPatchSettingsPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
ImageDefault - The virtual machine&rsquo;s default patching configuration is used.
AutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property provisionVMAgent
must be true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettingsARM">LinuxPatchSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfigurationARM">LinuxConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxPatchSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/LinuxPatchSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsPatchMode">
LinuxPatchSettingsPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
ImageDefault - The virtual machine&rsquo;s default patching configuration is used.
AutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property provisionVMAgent
must be true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettingsPatchMode">LinuxPatchSettingsPatchMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings">LinuxPatchSettings</a>, <a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsARM">LinuxPatchSettingsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutomaticByPlatform&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ImageDefault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettingsStatusPatchMode">LinuxPatchSettingsStatusPatchMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings_Status">LinuxPatchSettings_Status</a>, <a href="#compute.azure.com/v1beta20201201.LinuxPatchSettings_StatusARM">LinuxPatchSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutomaticByPlatform&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ImageDefault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettings_Status">LinuxPatchSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_Status">LinuxConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsStatusPatchMode">
LinuxPatchSettingsStatusPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
ImageDefault - The virtual machine&rsquo;s default patching configuration is used.
AutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property provisionVMAgent
must be true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.LinuxPatchSettings_StatusARM">LinuxPatchSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_StatusARM">LinuxConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxPatchSettingsStatusPatchMode">
LinuxPatchSettingsStatusPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
ImageDefault - The virtual machine&rsquo;s default patching configuration is used.
AutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property provisionVMAgent
must be true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.MaintenanceRedeployStatusStatusLastOperationResultCode">MaintenanceRedeployStatusStatusLastOperationResultCode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_Status">MaintenanceRedeployStatus_Status</a>, <a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_StatusARM">MaintenanceRedeployStatus_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;MaintenanceAborted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MaintenanceCompleted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RetryLater&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_Status">MaintenanceRedeployStatus_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>isCustomerInitiatedMaintenanceAllowed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCustomerInitiatedMaintenanceAllowed: True, if customer is allowed to perform Maintenance.</p>
</td>
</tr>
<tr>
<td>
<code>lastOperationMessage</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastOperationMessage: Message returned for the last Maintenance Operation.</p>
</td>
</tr>
<tr>
<td>
<code>lastOperationResultCode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatusStatusLastOperationResultCode">
MaintenanceRedeployStatusStatusLastOperationResultCode
</a>
</em>
</td>
<td>
<p>LastOperationResultCode: The Last Maintenance Operation Result Code.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowEndTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindowEndTime: End Time for the Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowStartTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindowStartTime: Start Time for the Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>preMaintenanceWindowEndTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>PreMaintenanceWindowEndTime: End Time for the Pre Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>preMaintenanceWindowStartTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>PreMaintenanceWindowStartTime: Start Time for the Pre Maintenance Window.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_StatusARM">MaintenanceRedeployStatus_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>isCustomerInitiatedMaintenanceAllowed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCustomerInitiatedMaintenanceAllowed: True, if customer is allowed to perform Maintenance.</p>
</td>
</tr>
<tr>
<td>
<code>lastOperationMessage</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastOperationMessage: Message returned for the last Maintenance Operation.</p>
</td>
</tr>
<tr>
<td>
<code>lastOperationResultCode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatusStatusLastOperationResultCode">
MaintenanceRedeployStatusStatusLastOperationResultCode
</a>
</em>
</td>
<td>
<p>LastOperationResultCode: The Last Maintenance Operation Result Code.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowEndTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindowEndTime: End Time for the Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindowStartTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindowStartTime: Start Time for the Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>preMaintenanceWindowEndTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>PreMaintenanceWindowEndTime: End Time for the Pre Maintenance Window.</p>
</td>
</tr>
<tr>
<td>
<code>preMaintenanceWindowStartTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>PreMaintenanceWindowStartTime: Start Time for the Pre Maintenance Window.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ManagedDiskParameters">ManagedDiskParameters
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk">DataDisk</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ManagedDiskParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ManagedDiskParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersStorageAccountType">
ManagedDiskParametersStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. Managed OS disk storage account type can
only be set when you create the scale set. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with
OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ManagedDiskParametersARM">ManagedDiskParametersARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ManagedDiskParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ManagedDiskParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersStorageAccountType">
ManagedDiskParametersStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. Managed OS disk storage account type can
only be set when you create the scale set. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with
OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ManagedDiskParametersStorageAccountType">ManagedDiskParametersStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters">ManagedDiskParameters</a>, <a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersARM">ManagedDiskParametersARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ManagedDiskParameters_Status">ManagedDiskParameters_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. Managed OS disk storage account type can
only be set when you create the scale set. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with
OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ManagedDiskParameters_StatusARM">ManagedDiskParameters_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. Managed OS disk storage account type can
only be set when you create the scale set. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with
OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkInterfaceReferencePropertiesARM">NetworkInterfaceReferencePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM">VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/NetworkInterfaceReferenceProperties">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/NetworkInterfaceReferenceProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkInterfaceReferenceProperties_StatusARM">NetworkInterfaceReferenceProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.NetworkInterfaceReference_StatusARM">NetworkInterfaceReference_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkInterfaceReference_Status">NetworkInterfaceReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.NetworkProfile_Status">NetworkProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkInterfaceReference_StatusARM">NetworkInterfaceReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.NetworkProfile_StatusARM">NetworkProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkInterfaceReferenceProperties_StatusARM">
NetworkInterfaceReferenceProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkProfile_Status">NetworkProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkInterfaceReference_Status">
[]NetworkInterfaceReference_Status
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: Specifies the list of resource Ids for the network interfaces associated with the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.NetworkProfile_StatusARM">NetworkProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkInterfaceReference_StatusARM">
[]NetworkInterfaceReference_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: Specifies the list of resource Ids for the network interfaces associated with the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDisk">OSDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile">StorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskCaching">
OSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskCreateOption">
OSDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings">
DiffDiskSettings
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Describes the parameters of ephemeral disk settings that can be specified for operating system disk.
NOTE: The ephemeral disk settings can only be specified for managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings">
DiskEncryptionSettings
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Describes a Encryption Settings for a Disk</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk">
VirtualHardDisk
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters">
ManagedDiskParameters
</a>
</em>
</td>
<td>
<p>ManagedDisk: The parameters of a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskOsType">
OSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk">
VirtualHardDisk
</a>
</em>
</td>
<td>
<p>Vhd: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfileARM">StorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskCaching">
OSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskCreateOption">
OSDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsARM">
DiffDiskSettingsARM
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Describes the parameters of ephemeral disk settings that can be specified for operating system disk.
NOTE: The ephemeral disk settings can only be specified for managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettingsARM">
DiskEncryptionSettingsARM
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Describes a Encryption Settings for a Disk</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDiskARM">
VirtualHardDiskARM
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParametersARM">
ManagedDiskParametersARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The parameters of a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskOsType">
OSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDiskARM">
VirtualHardDiskARM
</a>
</em>
</td>
<td>
<p>Vhd: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDiskCaching">OSDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>, <a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDiskCreateOption">OSDiskCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>, <a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDiskOsType">OSDiskOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>, <a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDiskStatusOsType">OSDiskStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_Status">StorageProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_Status">
DiffDiskSettings_Status
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Specifies the ephemeral Disk Settings for the operating system disk used by the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_Status">
DiskEncryptionSettings_Status
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Specifies the encryption settings for the OS Disk.
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_Status">
VirtualHardDisk_Status
</a>
</em>
</td>
<td>
<p>Image: The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_Status">
ManagedDiskParameters_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskStatusOsType">
OSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_Status">
VirtualHardDisk_Status
</a>
</em>
</td>
<td>
<p>Vhd: The virtual hard disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.StorageProfile_StatusARM">StorageProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machine should be created.
Possible values are:
Attach \u2013 This value is used when you are using a specialized disk to create the virtual machine.
FromImage \u2013 This value is used when you are using an image to create the virtual machine. If you are using a
platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also
use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_StatusARM">
DiffDiskSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Specifies the ephemeral Disk Settings for the operating system disk used by the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSettings_StatusARM">
DiskEncryptionSettings_StatusARM
</a>
</em>
</td>
<td>
<p>EncryptionSettings: Specifies the encryption settings for the OS Disk.
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">
VirtualHardDisk_StatusARM
</a>
</em>
</td>
<td>
<p>Image: The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the
virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_StatusARM">
ManagedDiskParameters_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskStatusOsType">
OSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>vhd</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">
VirtualHardDisk_StatusARM
</a>
</em>
</td>
<td>
<p>Vhd: The virtual hard disk.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSProfile">OSProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
This property cannot be updated after the VM is created.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>allowExtensionOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowExtensionOperations: Specifies whether extension operations should be allowed on the virtual machine.
This may only be set to False when no extensions are present on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: Specifies the host OS name of the virtual machine.
This name cannot be updated after the VM is created.
Max-length (Windows): 15 characters
Max-length (Linux): 64 characters.
For naming conventions and restrictions see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions">Azure infrastructure services implementation
guidelines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
Note: Do not pass any secrets or passwords in customData property
This property cannot be updated after the VM is created.
customData is passed to the VM to be saved as a file, for more information see <a href="https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/">Custom Data on Azure
VMs</a>
For using cloud-init for your Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration">
LinuxConfiguration
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>requireGuestProvisionSignal</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireGuestProvisionSignal: Specifies whether the guest provision signal is required to infer provision success of the
virtual machine.  Note: This property is for private testing only, and all customers must not set the property to false.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup">
[]VaultSecretGroup
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration">
WindowsConfiguration
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSProfileARM">OSProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/OSProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
This property cannot be updated after the VM is created.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>allowExtensionOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowExtensionOperations: Specifies whether extension operations should be allowed on the virtual machine.
This may only be set to False when no extensions are present on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: Specifies the host OS name of the virtual machine.
This name cannot be updated after the VM is created.
Max-length (Windows): 15 characters
Max-length (Linux): 64 characters.
For naming conventions and restrictions see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions">Azure infrastructure services implementation
guidelines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
Note: Do not pass any secrets or passwords in customData property
This property cannot be updated after the VM is created.
customData is passed to the VM to be saved as a file, for more information see <a href="https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/">Custom Data on Azure
VMs</a>
For using cloud-init for your Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfigurationARM">
LinuxConfigurationARM
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>requireGuestProvisionSignal</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireGuestProvisionSignal: Specifies whether the guest provision signal is required to infer provision success of the
virtual machine.  Note: This property is for private testing only, and all customers must not set the property to false.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroupARM">
[]VaultSecretGroupARM
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfigurationARM">
WindowsConfigurationARM
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSProfile_Status">OSProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
This property cannot be updated after the VM is created.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>allowExtensionOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowExtensionOperations: Specifies whether extension operations should be allowed on the virtual machine.
This may only be set to False when no extensions are present on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: Specifies the host OS name of the virtual machine.
This name cannot be updated after the VM is created.
Max-length (Windows): 15 characters
Max-length (Linux): 64 characters.
For naming conventions and restrictions see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions">Azure infrastructure services implementation
guidelines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
Note: Do not pass any secrets or passwords in customData property
This property cannot be updated after the VM is created.
customData is passed to the VM to be saved as a file, for more information see <a href="https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/">Custom Data on Azure
VMs</a>
For using cloud-init for your Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_Status">
LinuxConfiguration_Status
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>requireGuestProvisionSignal</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireGuestProvisionSignal: Specifies whether the guest provision signal is required to infer provision success of the
virtual machine.  Note: This property is for private testing only, and all customers must not set the property to false.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_Status">
[]VaultSecretGroup_Status
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_Status">
WindowsConfiguration_Status
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OSProfile_StatusARM">OSProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
This property cannot be updated after the VM is created.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>allowExtensionOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowExtensionOperations: Specifies whether extension operations should be allowed on the virtual machine.
This may only be set to False when no extensions are present on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: Specifies the host OS name of the virtual machine.
This name cannot be updated after the VM is created.
Max-length (Windows): 15 characters
Max-length (Linux): 64 characters.
For naming conventions and restrictions see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions">Azure infrastructure services implementation
guidelines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
Note: Do not pass any secrets or passwords in customData property
This property cannot be updated after the VM is created.
customData is passed to the VM to be saved as a file, for more information see <a href="https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/">Custom Data on Azure
VMs</a>
For using cloud-init for your Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_StatusARM">
LinuxConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>requireGuestProvisionSignal</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireGuestProvisionSignal: Specifies whether the guest provision signal is required to infer provision success of the
virtual machine.  Note: This property is for private testing only, and all customers must not set the property to false.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_StatusARM">
[]VaultSecretGroup_StatusARM
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">
WindowsConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.OrchestrationMode_Status">OrchestrationMode_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Flexible&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Uniform&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettings">PatchSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration">WindowsConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/PatchSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/PatchSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableHotpatching</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableHotpatching: Enables customers to patch their Azure VMs without requiring a reboot. For enableHotpatching, the
&lsquo;provisionVMAgent&rsquo; must be set to true and &lsquo;patchMode&rsquo; must be set to &lsquo;AutomaticByPlatform&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettingsPatchMode">
PatchSettingsPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
Manual - You  control the application of patches to a virtual machine. You do this by applying patches manually inside
the VM. In this mode, automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be
false
AutomaticByOS - The virtual machine will automatically be updated by the OS. The property
WindowsConfiguration.enableAutomaticUpdates must be true.
AutomaticByPlatform - the virtual machine will automatically updated by the platform. The properties provisionVMAgent
and WindowsConfiguration.enableAutomaticUpdates must be true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettingsARM">PatchSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfigurationARM">WindowsConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/PatchSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/PatchSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableHotpatching</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableHotpatching: Enables customers to patch their Azure VMs without requiring a reboot. For enableHotpatching, the
&lsquo;provisionVMAgent&rsquo; must be set to true and &lsquo;patchMode&rsquo; must be set to &lsquo;AutomaticByPlatform&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettingsPatchMode">
PatchSettingsPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
Manual - You  control the application of patches to a virtual machine. You do this by applying patches manually inside
the VM. In this mode, automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be
false
AutomaticByOS - The virtual machine will automatically be updated by the OS. The property
WindowsConfiguration.enableAutomaticUpdates must be true.
AutomaticByPlatform - the virtual machine will automatically updated by the platform. The properties provisionVMAgent
and WindowsConfiguration.enableAutomaticUpdates must be true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettingsPatchMode">PatchSettingsPatchMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.PatchSettings">PatchSettings</a>, <a href="#compute.azure.com/v1beta20201201.PatchSettingsARM">PatchSettingsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutomaticByOS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AutomaticByPlatform&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manual&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettingsStatusPatchMode">PatchSettingsStatusPatchMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.PatchSettings_Status">PatchSettings_Status</a>, <a href="#compute.azure.com/v1beta20201201.PatchSettings_StatusARM">PatchSettings_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AutomaticByOS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AutomaticByPlatform&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manual&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettings_Status">PatchSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_Status">WindowsConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableHotpatching</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableHotpatching: Enables customers to patch their Azure VMs without requiring a reboot. For enableHotpatching, the
&lsquo;provisionVMAgent&rsquo; must be set to true and &lsquo;patchMode&rsquo; must be set to &lsquo;AutomaticByPlatform&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettingsStatusPatchMode">
PatchSettingsStatusPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
Manual - You  control the application of patches to a virtual machine. You do this by applying patches manually inside
the VM. In this mode, automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be
false
AutomaticByOS - The virtual machine will automatically be updated by the OS. The property
WindowsConfiguration.enableAutomaticUpdates must be true.
AutomaticByPlatform - the virtual machine will automatically updated by the platform. The properties provisionVMAgent
and WindowsConfiguration.enableAutomaticUpdates must be true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PatchSettings_StatusARM">PatchSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">WindowsConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableHotpatching</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableHotpatching: Enables customers to patch their Azure VMs without requiring a reboot. For enableHotpatching, the
&lsquo;provisionVMAgent&rsquo; must be set to true and &lsquo;patchMode&rsquo; must be set to &lsquo;AutomaticByPlatform&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>patchMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettingsStatusPatchMode">
PatchSettingsStatusPatchMode
</a>
</em>
</td>
<td>
<p>PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
Possible values are:
Manual - You  control the application of patches to a virtual machine. You do this by applying patches manually inside
the VM. In this mode, automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be
false
AutomaticByOS - The virtual machine will automatically be updated by the OS. The property
WindowsConfiguration.enableAutomaticUpdates must be true.
AutomaticByPlatform - the virtual machine will automatically updated by the platform. The properties provisionVMAgent
and WindowsConfiguration.enableAutomaticUpdates must be true</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Plan">Plan
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Plan">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Plan</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The promotion code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.PlanARM">PlanARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_SpecARM">VirtualMachines_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Plan">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Plan</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The promotion code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Plan_Status">Plan_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The promotion code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Plan_StatusARM">Plan_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The promotion code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Priority_Status">Priority_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Low&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.RollingUpgradePolicy">RollingUpgradePolicy
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy">UpgradePolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/RollingUpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/RollingUpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableCrossZoneUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCrossZoneUpgrade: Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration
the Update Domain and maxBatchInstancePercent to determine the batch size.</p>
</td>
</tr>
<tr>
<td>
<code>maxBatchInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxBatchInstancePercent: The maximum percent of total virtual machine instances that will be upgraded simultaneously by
the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the
percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyInstancePercent: The maximum percentage of the total virtual machine instances in the scale set that can be
simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual
machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyUpgradedInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyUpgradedInstancePercent: The maximum percentage of upgraded virtual machine instances that can be found to
be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the
rolling update aborts. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>pauseTimeBetweenBatches</code><br/>
<em>
string
</em>
</td>
<td>
<p>PauseTimeBetweenBatches: The wait time between completing the update for all virtual machines in one batch and starting
the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).</p>
</td>
</tr>
<tr>
<td>
<code>prioritizeUnhealthyInstances</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PrioritizeUnhealthyInstances: Upgrade all unhealthy instances in a scale set before any healthy instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.RollingUpgradePolicyARM">RollingUpgradePolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicyARM">UpgradePolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/RollingUpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/RollingUpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableCrossZoneUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCrossZoneUpgrade: Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration
the Update Domain and maxBatchInstancePercent to determine the batch size.</p>
</td>
</tr>
<tr>
<td>
<code>maxBatchInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxBatchInstancePercent: The maximum percent of total virtual machine instances that will be upgraded simultaneously by
the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the
percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyInstancePercent: The maximum percentage of the total virtual machine instances in the scale set that can be
simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual
machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyUpgradedInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyUpgradedInstancePercent: The maximum percentage of upgraded virtual machine instances that can be found to
be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the
rolling update aborts. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>pauseTimeBetweenBatches</code><br/>
<em>
string
</em>
</td>
<td>
<p>PauseTimeBetweenBatches: The wait time between completing the update for all virtual machines in one batch and starting
the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).</p>
</td>
</tr>
<tr>
<td>
<code>prioritizeUnhealthyInstances</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PrioritizeUnhealthyInstances: Upgrade all unhealthy instances in a scale set before any healthy instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.RollingUpgradePolicy_Status">RollingUpgradePolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy_Status">UpgradePolicy_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableCrossZoneUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCrossZoneUpgrade: Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration
the Update Domain and maxBatchInstancePercent to determine the batch size.</p>
</td>
</tr>
<tr>
<td>
<code>maxBatchInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxBatchInstancePercent: The maximum percent of total virtual machine instances that will be upgraded simultaneously by
the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the
percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyInstancePercent: The maximum percentage of the total virtual machine instances in the scale set that can be
simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual
machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyUpgradedInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyUpgradedInstancePercent: The maximum percentage of upgraded virtual machine instances that can be found to
be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the
rolling update aborts. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>pauseTimeBetweenBatches</code><br/>
<em>
string
</em>
</td>
<td>
<p>PauseTimeBetweenBatches: The wait time between completing the update for all virtual machines in one batch and starting
the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).</p>
</td>
</tr>
<tr>
<td>
<code>prioritizeUnhealthyInstances</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PrioritizeUnhealthyInstances: Upgrade all unhealthy instances in a scale set before any healthy instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.RollingUpgradePolicy_StatusARM">RollingUpgradePolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy_StatusARM">UpgradePolicy_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableCrossZoneUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCrossZoneUpgrade: Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration
the Update Domain and maxBatchInstancePercent to determine the batch size.</p>
</td>
</tr>
<tr>
<td>
<code>maxBatchInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxBatchInstancePercent: The maximum percent of total virtual machine instances that will be upgraded simultaneously by
the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the
percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyInstancePercent: The maximum percentage of the total virtual machine instances in the scale set that can be
simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual
machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>maxUnhealthyUpgradedInstancePercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxUnhealthyUpgradedInstancePercent: The maximum percentage of upgraded virtual machine instances that can be found to
be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the
rolling update aborts. The default value for this parameter is 20%.</p>
</td>
</tr>
<tr>
<td>
<code>pauseTimeBetweenBatches</code><br/>
<em>
string
</em>
</td>
<td>
<p>PauseTimeBetweenBatches: The wait time between completing the update for all virtual machines in one batch and starting
the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).</p>
</td>
</tr>
<tr>
<td>
<code>prioritizeUnhealthyInstances</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PrioritizeUnhealthyInstances: Upgrade all unhealthy instances in a scale set before any healthy instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicy">ScaleInPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScaleInPolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScaleInPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicyRules">
[]ScaleInPolicyRules
</a>
</em>
</td>
<td>
<p>Rules: The rules to be followed when scaling-in a virtual machine scale set.
Possible values are:
Default When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a
zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the
virtual machines chosen for removal will be the newest ones that are not protected from scale-in.
OldestVM When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal.
NewestVM When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicyARM">ScaleInPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScaleInPolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScaleInPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicyRules">
[]ScaleInPolicyRules
</a>
</em>
</td>
<td>
<p>Rules: The rules to be followed when scaling-in a virtual machine scale set.
Possible values are:
Default When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a
zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the
virtual machines chosen for removal will be the newest ones that are not protected from scale-in.
OldestVM When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal.
NewestVM When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicyRules">ScaleInPolicyRules
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScaleInPolicy">ScaleInPolicy</a>, <a href="#compute.azure.com/v1beta20201201.ScaleInPolicyARM">ScaleInPolicyARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NewestVM&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OldestVM&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicyStatusRules">ScaleInPolicyStatusRules
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScaleInPolicy_Status">ScaleInPolicy_Status</a>, <a href="#compute.azure.com/v1beta20201201.ScaleInPolicy_StatusARM">ScaleInPolicy_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NewestVM&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OldestVM&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicy_Status">ScaleInPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicyStatusRules">
[]ScaleInPolicyStatusRules
</a>
</em>
</td>
<td>
<p>Rules: The rules to be followed when scaling-in a virtual machine scale set.
Possible values are:
Default When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a
zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the
virtual machines chosen for removal will be the newest ones that are not protected from scale-in.
OldestVM When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal.
NewestVM When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScaleInPolicy_StatusARM">ScaleInPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicyStatusRules">
[]ScaleInPolicyStatusRules
</a>
</em>
</td>
<td>
<p>Rules: The rules to be followed when scaling-in a virtual machine scale set.
Possible values are:
Default When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a
zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the
virtual machines chosen for removal will be the newest ones that are not protected from scale-in.
OldestVM When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal.
NewestVM When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from
scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScheduledEventsProfile">ScheduledEventsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScheduledEventsProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScheduledEventsProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>terminateNotificationProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.TerminateNotificationProfile">
TerminateNotificationProfile
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScheduledEventsProfileARM">ScheduledEventsProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScheduledEventsProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/ScheduledEventsProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>terminateNotificationProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.TerminateNotificationProfileARM">
TerminateNotificationProfileARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScheduledEventsProfile_Status">ScheduledEventsProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>terminateNotificationProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.TerminateNotificationProfile_Status">
TerminateNotificationProfile_Status
</a>
</em>
</td>
<td>
<p>TerminateNotificationProfile: Specifies Terminate Scheduled Event related configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.ScheduledEventsProfile_StatusARM">ScheduledEventsProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>terminateNotificationProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.TerminateNotificationProfile_StatusARM">
TerminateNotificationProfile_StatusARM
</a>
</em>
</td>
<td>
<p>TerminateNotificationProfile: Specifies Terminate Scheduled Event related configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfile">SecurityProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SecurityProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SecurityProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EncryptionAtHost: This property can be used by user in the request to enable or disable the Host Encryption for the
virtual machine or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp
disk at host itself.
Default: The Encryption at host will be disabled unless this property is set to true for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileSecurityType">
SecurityProfileSecurityType
</a>
</em>
</td>
<td>
<p>SecurityType: Specifies the SecurityType of the virtual machine. It is set as TrustedLaunch to enable UefiSettings.
Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch.</p>
</td>
</tr>
<tr>
<td>
<code>uefiSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UefiSettings">
UefiSettings
</a>
</em>
</td>
<td>
<p>UefiSettings: Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfileARM">SecurityProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SecurityProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SecurityProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EncryptionAtHost: This property can be used by user in the request to enable or disable the Host Encryption for the
virtual machine or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp
disk at host itself.
Default: The Encryption at host will be disabled unless this property is set to true for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileSecurityType">
SecurityProfileSecurityType
</a>
</em>
</td>
<td>
<p>SecurityType: Specifies the SecurityType of the virtual machine. It is set as TrustedLaunch to enable UefiSettings.
Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch.</p>
</td>
</tr>
<tr>
<td>
<code>uefiSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UefiSettingsARM">
UefiSettingsARM
</a>
</em>
</td>
<td>
<p>UefiSettings: Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfileSecurityType">SecurityProfileSecurityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfile">SecurityProfile</a>, <a href="#compute.azure.com/v1beta20201201.SecurityProfileARM">SecurityProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;TrustedLaunch&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfileStatusSecurityType">SecurityProfileStatusSecurityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfile_Status">SecurityProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.SecurityProfile_StatusARM">SecurityProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;TrustedLaunch&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfile_Status">SecurityProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EncryptionAtHost: This property can be used by user in the request to enable or disable the Host Encryption for the
virtual machine or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp
disk at host itself.
Default: The Encryption at host will be disabled unless this property is set to true for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileStatusSecurityType">
SecurityProfileStatusSecurityType
</a>
</em>
</td>
<td>
<p>SecurityType: Specifies the SecurityType of the virtual machine. It is set as TrustedLaunch to enable UefiSettings.
Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch.</p>
</td>
</tr>
<tr>
<td>
<code>uefiSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UefiSettings_Status">
UefiSettings_Status
</a>
</em>
</td>
<td>
<p>UefiSettings: Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SecurityProfile_StatusARM">SecurityProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EncryptionAtHost: This property can be used by user in the request to enable or disable the Host Encryption for the
virtual machine or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp
disk at host itself.
Default: The Encryption at host will be disabled unless this property is set to true for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileStatusSecurityType">
SecurityProfileStatusSecurityType
</a>
</em>
</td>
<td>
<p>SecurityType: Specifies the SecurityType of the virtual machine. It is set as TrustedLaunch to enable UefiSettings.
Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch.</p>
</td>
</tr>
<tr>
<td>
<code>uefiSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UefiSettings_StatusARM">
UefiSettings_StatusARM
</a>
</em>
</td>
<td>
<p>UefiSettings: Specifies the security settings like secure boot and vTPM used while creating the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Sku">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Sku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Specifies the number of virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Specifies the tier of virtual machines in a scale set.
Possible Values:
Standard
Basic</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Sku">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/Sku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Specifies the number of virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Specifies the tier of virtual machines in a scale set.
Possible Values:
Standard
Basic</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Specifies the number of virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Specifies the tier of virtual machines in a scale set.
Possible Values:
Standard
Basic</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Specifies the number of virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Specifies the tier of virtual machines in a scale set.
Possible Values:
Standard
Basic</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshConfiguration">SshConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration">LinuxConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshPublicKey">
[]SshPublicKey
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with linux based VMs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshConfigurationARM">SshConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfigurationARM">LinuxConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshPublicKeyARM">
[]SshPublicKeyARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with linux based VMs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshConfiguration_Status">SshConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_Status">LinuxConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshPublicKey_Status">
[]SshPublicKey_Status
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with linux based VMs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshConfiguration_StatusARM">SshConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_StatusARM">LinuxConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SshPublicKey_StatusARM">
[]SshPublicKey_StatusARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with linux based VMs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshPublicKey">SshPublicKey
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SshConfiguration">SshConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshPublicKey">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshPublicKey</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit
and in ssh-rsa format.
For creating ssh keys, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Create SSH keys on Linux and Mac for Linux VMs in
Azure</a>.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the
specified key is appended to the file. Example: /home/user/.ssh/authorized_keys</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshPublicKeyARM">SshPublicKeyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SshConfigurationARM">SshConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshPublicKey">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SshPublicKey</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit
and in ssh-rsa format.
For creating ssh keys, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Create SSH keys on Linux and Mac for Linux VMs in
Azure</a>.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the
specified key is appended to the file. Example: /home/user/.ssh/authorized_keys</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshPublicKey_Status">SshPublicKey_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SshConfiguration_Status">SshConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit
and in ssh-rsa format.
For creating ssh keys, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Create SSH keys on Linux and Mac for Linux VMs in
Azure</a>.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the
specified key is appended to the file. Example: /home/user/.ssh/authorized_keys</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SshPublicKey_StatusARM">SshPublicKey_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SshConfiguration_StatusARM">SshConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit
and in ssh-rsa format.
For creating ssh keys, see <a href="https://docs.microsoft.com/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Create SSH keys on Linux and Mac for Linux VMs in
Azure</a>.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the
specified key is appended to the file. Example: /home/user/.ssh/authorized_keys</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.StorageAccountType_Status">StorageAccountType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_Status">ManagedDiskParameters_Status</a>, <a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_StatusARM">ManagedDiskParameters_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status">VirtualMachineScaleSetManagedDiskParameters_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_StatusARM">VirtualMachineScaleSetManagedDiskParameters_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.StorageProfile">StorageProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/StorageProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/StorageProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDisk">
[]DataDisk
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference">
ImageReference
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations. NOTE: Image reference
publisher and offer can only be set when you create the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDisk">
OSDisk
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.StorageProfileARM">StorageProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/StorageProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/StorageProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDiskARM">
[]DataDiskARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReferenceARM">
ImageReferenceARM
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations. NOTE: Image reference
publisher and offer can only be set when you create the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDiskARM">
OSDiskARM
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.StorageProfile_Status">StorageProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDisk_Status">
[]DataDisk_Status
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference_Status">
ImageReference_Status
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDisk_Status">
OSDisk_Status
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.StorageProfile_StatusARM">StorageProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">
[]DataDisk_StatusARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference_StatusARM">
ImageReference_StatusARM
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">
OSDisk_StatusARM
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SubResource">SubResource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference">KeyVaultKeyReference</a>, <a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference">KeyVaultSecretReference</a>, <a href="#compute.azure.com/v1beta20201201.VaultSecretGroup">VaultSecretGroup</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SubResource</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SubResourceARM">SubResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReferenceARM">KeyVaultKeyReferenceARM</a>, <a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReferenceARM">KeyVaultSecretReferenceARM</a>, <a href="#compute.azure.com/v1beta20201201.VaultSecretGroupARM">VaultSecretGroupARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/SubResource</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SubResource_Status">SubResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference_Status">KeyVaultKeyReference_Status</a>, <a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference_Status">KeyVaultSecretReference_Status</a>, <a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_Status">ManagedDiskParameters_Status</a>, <a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_Status">VaultSecretGroup_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">VirtualMachineScaleSetIPConfiguration_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status">VirtualMachineScaleSetManagedDiskParameters_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status">VirtualMachineScaleSetNetworkConfiguration_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">VirtualMachineScaleSetPublicIPAddressConfiguration_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.SubResource_StatusARM">SubResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.KeyVaultKeyReference_StatusARM">KeyVaultKeyReference_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.KeyVaultSecretReference_StatusARM">KeyVaultSecretReference_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.ManagedDiskParameters_StatusARM">ManagedDiskParameters_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_StatusARM">VaultSecretGroup_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">VirtualMachineScaleSetIPConfigurationProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_StatusARM">VirtualMachineScaleSetManagedDiskParameters_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM">VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.TerminateNotificationProfile">TerminateNotificationProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile">ScheduledEventsProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/TerminateNotificationProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/TerminateNotificationProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enable: Specifies whether the Terminate Scheduled event is enabled or disabled.</p>
</td>
</tr>
<tr>
<td>
<code>notBeforeTimeout</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotBeforeTimeout: Configurable length of time a Virtual Machine being deleted will have to potentially approve the
Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601
format, the default value is 5 minutes (PT5M)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.TerminateNotificationProfileARM">TerminateNotificationProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfileARM">ScheduledEventsProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/TerminateNotificationProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/TerminateNotificationProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enable: Specifies whether the Terminate Scheduled event is enabled or disabled.</p>
</td>
</tr>
<tr>
<td>
<code>notBeforeTimeout</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotBeforeTimeout: Configurable length of time a Virtual Machine being deleted will have to potentially approve the
Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601
format, the default value is 5 minutes (PT5M)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.TerminateNotificationProfile_Status">TerminateNotificationProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile_Status">ScheduledEventsProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enable: Specifies whether the Terminate Scheduled event is enabled or disabled.</p>
</td>
</tr>
<tr>
<td>
<code>notBeforeTimeout</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotBeforeTimeout: Configurable length of time a Virtual Machine being deleted will have to potentially approve the
Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601
format, the default value is 5 minutes (PT5M)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.TerminateNotificationProfile_StatusARM">TerminateNotificationProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile_StatusARM">ScheduledEventsProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enable: Specifies whether the Terminate Scheduled event is enabled or disabled.</p>
</td>
</tr>
<tr>
<td>
<code>notBeforeTimeout</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotBeforeTimeout: Configurable length of time a Virtual Machine being deleted will have to potentially approve the
Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601
format, the default value is 5 minutes (PT5M)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UefiSettings">UefiSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfile">SecurityProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UefiSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UefiSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secureBootEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SecureBootEnabled: Specifies whether secure boot should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
<tr>
<td>
<code>vTpmEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>VTpmEnabled: Specifies whether vTPM should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UefiSettingsARM">UefiSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfileARM">SecurityProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UefiSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UefiSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secureBootEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SecureBootEnabled: Specifies whether secure boot should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
<tr>
<td>
<code>vTpmEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>VTpmEnabled: Specifies whether vTPM should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UefiSettings_Status">UefiSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfile_Status">SecurityProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secureBootEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SecureBootEnabled: Specifies whether secure boot should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
<tr>
<td>
<code>vTpmEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>VTpmEnabled: Specifies whether vTPM should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UefiSettings_StatusARM">UefiSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.SecurityProfile_StatusARM">SecurityProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secureBootEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SecureBootEnabled: Specifies whether secure boot should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
<tr>
<td>
<code>vTpmEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>VTpmEnabled: Specifies whether vTPM should be enabled on the virtual machine.
Minimum api-version: 2020-12-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicy">UpgradePolicy
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>automaticOSUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy">
AutomaticOSUpgradePolicy
</a>
</em>
</td>
<td>
<p>AutomaticOSUpgradePolicy: The configuration parameters used for performing automatic OS upgrade.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicyMode">
UpgradePolicyMode
</a>
</em>
</td>
<td>
<p>Mode: Specifies the mode of an upgrade to virtual machines in the scale set.
Possible values are:
Manual - You  control the application of updates to virtual machines in the scale set. You do this by using the
manualUpgrade action.
Automatic - All virtual machines in the scale set are  automatically updated at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>rollingUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.RollingUpgradePolicy">
RollingUpgradePolicy
</a>
</em>
</td>
<td>
<p>RollingUpgradePolicy: The configuration parameters used while performing a rolling upgrade.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicyARM">UpgradePolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UpgradePolicy">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/UpgradePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>automaticOSUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicyARM">
AutomaticOSUpgradePolicyARM
</a>
</em>
</td>
<td>
<p>AutomaticOSUpgradePolicy: The configuration parameters used for performing automatic OS upgrade.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicyMode">
UpgradePolicyMode
</a>
</em>
</td>
<td>
<p>Mode: Specifies the mode of an upgrade to virtual machines in the scale set.
Possible values are:
Manual - You  control the application of updates to virtual machines in the scale set. You do this by using the
manualUpgrade action.
Automatic - All virtual machines in the scale set are  automatically updated at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>rollingUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.RollingUpgradePolicyARM">
RollingUpgradePolicyARM
</a>
</em>
</td>
<td>
<p>RollingUpgradePolicy: The configuration parameters used while performing a rolling upgrade.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicyMode">UpgradePolicyMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy">UpgradePolicy</a>, <a href="#compute.azure.com/v1beta20201201.UpgradePolicyARM">UpgradePolicyARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Automatic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Rolling&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicyStatusMode">UpgradePolicyStatusMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.UpgradePolicy_Status">UpgradePolicy_Status</a>, <a href="#compute.azure.com/v1beta20201201.UpgradePolicy_StatusARM">UpgradePolicy_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Automatic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Rolling&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicy_Status">UpgradePolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>automaticOSUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy_Status">
AutomaticOSUpgradePolicy_Status
</a>
</em>
</td>
<td>
<p>AutomaticOSUpgradePolicy: Configuration parameters used for performing automatic OS Upgrade.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicyStatusMode">
UpgradePolicyStatusMode
</a>
</em>
</td>
<td>
<p>Mode: Specifies the mode of an upgrade to virtual machines in the scale set.
Possible values are:
Manual - You  control the application of updates to virtual machines in the scale set. You do this by using the
manualUpgrade action.
Automatic - All virtual machines in the scale set are  automatically updated at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>rollingUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.RollingUpgradePolicy_Status">
RollingUpgradePolicy_Status
</a>
</em>
</td>
<td>
<p>RollingUpgradePolicy: The configuration parameters used while performing a rolling upgrade.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.UpgradePolicy_StatusARM">UpgradePolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>automaticOSUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticOSUpgradePolicy_StatusARM">
AutomaticOSUpgradePolicy_StatusARM
</a>
</em>
</td>
<td>
<p>AutomaticOSUpgradePolicy: Configuration parameters used for performing automatic OS Upgrade.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicyStatusMode">
UpgradePolicyStatusMode
</a>
</em>
</td>
<td>
<p>Mode: Specifies the mode of an upgrade to virtual machines in the scale set.
Possible values are:
Manual - You  control the application of updates to virtual machines in the scale set. You do this by using the
manualUpgrade action.
Automatic - All virtual machines in the scale set are  automatically updated at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>rollingUpgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.RollingUpgradePolicy_StatusARM">
RollingUpgradePolicy_StatusARM
</a>
</em>
</td>
<td>
<p>RollingUpgradePolicy: The configuration parameters used while performing a rolling upgrade.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultCertificate">VaultCertificate
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VaultSecretGroup">VaultSecretGroup</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultCertificate">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultCertificate</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateStore</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateStore: For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate
should be added. The specified certificate store is implicitly in the LocalMachine account.
For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name
&amp;lt;UppercaseThumbprint&amp;gt;.crt for the X509 certificate file and &amp;lt;UppercaseThumbprint&amp;gt;.prv for private key. Both
of these files are .pem formatted.</p>
</td>
</tr>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultCertificateARM">VaultCertificateARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VaultSecretGroupARM">VaultSecretGroupARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultCertificate">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultCertificate</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateStore</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateStore: For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate
should be added. The specified certificate store is implicitly in the LocalMachine account.
For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name
&amp;lt;UppercaseThumbprint&amp;gt;.crt for the X509 certificate file and &amp;lt;UppercaseThumbprint&amp;gt;.prv for private key. Both
of these files are .pem formatted.</p>
</td>
</tr>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultCertificate_Status">VaultCertificate_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_Status">VaultSecretGroup_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateStore</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateStore: For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate
should be added. The specified certificate store is implicitly in the LocalMachine account.
For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name
&amp;lt;UppercaseThumbprint&amp;gt;.crt for the X509 certificate file and &amp;lt;UppercaseThumbprint&amp;gt;.prv for private key. Both
of these files are .pem formatted.</p>
</td>
</tr>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultCertificate_StatusARM">VaultCertificate_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_StatusARM">VaultSecretGroup_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateStore</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateStore: For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate
should be added. The specified certificate store is implicitly in the LocalMachine account.
For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name
&amp;lt;UppercaseThumbprint&amp;gt;.crt for the X509 certificate file and &amp;lt;UppercaseThumbprint&amp;gt;.prv for private key. Both
of these files are .pem formatted.</p>
</td>
</tr>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultSecretGroup">VaultSecretGroup
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile">OSProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile">VirtualMachineScaleSetOSProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultSecretGroup">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultSecretGroup</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>vaultCertificates</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultCertificate">
[]VaultCertificate
</a>
</em>
</td>
<td>
<p>VaultCertificates: The list of key vault references in SourceVault which contain certificates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultSecretGroupARM">VaultSecretGroupARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfileARM">OSProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfileARM">VirtualMachineScaleSetOSProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultSecretGroup">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VaultSecretGroup</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>vaultCertificates</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultCertificateARM">
[]VaultCertificateARM
</a>
</em>
</td>
<td>
<p>VaultCertificates: The list of key vault references in SourceVault which contain certificates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultSecretGroup_Status">VaultSecretGroup_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_Status">OSProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_Status">VirtualMachineScaleSetOSProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing all of the certificates in VaultCertificates.</p>
</td>
</tr>
<tr>
<td>
<code>vaultCertificates</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultCertificate_Status">
[]VaultCertificate_Status
</a>
</em>
</td>
<td>
<p>VaultCertificates: The list of key vault references in SourceVault which contain certificates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VaultSecretGroup_StatusARM">VaultSecretGroup_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_StatusARM">OSProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_StatusARM">VirtualMachineScaleSetOSProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVault: The relative URL of the Key Vault containing all of the certificates in VaultCertificates.</p>
</td>
</tr>
<tr>
<td>
<code>vaultCertificates</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultCertificate_StatusARM">
[]VaultCertificate_StatusARM
</a>
</em>
</td>
<td>
<p>VaultCertificates: The list of key vault references in SourceVault which contain certificates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualHardDisk">VirtualHardDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk">DataDisk</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk">OSDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualHardDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualHardDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uri</code><br/>
<em>
string
</em>
</td>
<td>
<p>Uri: Specifies the virtual hard disk&rsquo;s uri.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualHardDiskARM">VirtualHardDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDiskARM">DataDiskARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDiskARM">OSDiskARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualHardDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualHardDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uri</code><br/>
<em>
string
</em>
</td>
<td>
<p>Uri: Specifies the virtual hard disk&rsquo;s uri.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualHardDisk_Status">VirtualHardDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_Status">DataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_Status">OSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uri</code><br/>
<em>
string
</em>
</td>
<td>
<p>Uri: Specifies the virtual hard disk&rsquo;s uri.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">VirtualHardDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.DataDisk_StatusARM">DataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.OSDisk_StatusARM">OSDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uri</code><br/>
<em>
string
</em>
</td>
<td>
<p>Uri: Specifies the virtual hard disk&rsquo;s uri.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachine">VirtualMachine
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/resourceDefinitions/virtualMachines">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/resourceDefinitions/virtualMachines</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">
VirtualMachines_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities">
AdditionalCapabilities
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>availabilitySet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile">
BillingProfile
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VM or VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile">
DiagnosticsProfile
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesEvictionPolicy">
VirtualMachinesSpecPropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
<tr>
<td>
<code>hardwareProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfile">
HardwareProfile
</a>
</em>
</td>
<td>
<p>HardwareProfile: Specifies the hardware settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity">
VirtualMachineIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile">
VirtualMachines_Spec_Properties_NetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies the network interfaces of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSProfile">
OSProfile
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings for the virtual machine. Some of the settings cannot be changed once
VM is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan">
Plan
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the scale set logical fault domain into which the Virtual Machine will be created. By
default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across
available fault domains.
<li>This is applicable only if the &lsquo;virtualMachineScaleSet&rsquo; property of this Virtual Machine is set.<li>The Virtual
Machine Scale Set that is referenced, must have &lsquo;platformFaultDomainCount&rsquo; &amp;gt; 1.<li>This property cannot be updated
once the Virtual Machine is created.<li>Fault domain assignment can be viewed in the Virtual Machine Instance View.
Minimum api‐version: 2020‐12‐01</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesPriority">
VirtualMachinesSpecPropertiesPriority
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machine.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile">
SecurityProfile
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security profile settings for the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageProfile">
StorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineScaleSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine zones.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">
VirtualMachine_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_Status">VirtualMachineAgentInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensionHandlers</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_Status">
[]VirtualMachineExtensionHandlerInstanceView_Status
</a>
</em>
</td>
<td>
<p>ExtensionHandlers: The virtual machine extension handler instance view.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>vmAgentVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmAgentVersion: The VM Agent full version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_StatusARM">VirtualMachineAgentInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensionHandlers</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_StatusARM">
[]VirtualMachineExtensionHandlerInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>ExtensionHandlers: The virtual machine extension handler instance view.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>vmAgentVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmAgentVersion: The VM Agent full version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_Status">VirtualMachineExtensionHandlerInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_Status">VirtualMachineAgentInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Status: The extension handler status.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtensionHandlerInstanceView_StatusARM">VirtualMachineExtensionHandlerInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_StatusARM">VirtualMachineAgentInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Status: The extension handler status.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_Status">VirtualMachineExtensionInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineExtension_Status">VirtualMachineExtension_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The virtual machine extension name.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>substatuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Substatuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_StatusARM">VirtualMachineExtensionInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionProperties_StatusARM">VirtualMachineExtensionProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The virtual machine extension name.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>substatuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Substatuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtensionProperties_StatusARM">VirtualMachineExtensionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineExtension_StatusARM">VirtualMachineExtension_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoUpgradeMinorVersion</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
property set to true.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
newer version of the extension available.</p>
</td>
</tr>
<tr>
<td>
<code>forceUpdateTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForceUpdateTag: How the extension handler should be forced to update even if the extension configuration has not changed.</p>
</td>
</tr>
<tr>
<td>
<code>instanceView</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_StatusARM">
VirtualMachineExtensionInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>InstanceView: The virtual machine extension instance view.</p>
</td>
</tr>
<tr>
<td>
<code>protectedSettings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
settings at all.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The name of the extension handler publisher.</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Json formatted public settings for the extension.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtension_Status">VirtualMachineExtension_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoUpgradeMinorVersion</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
property set to true.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
newer version of the extension available.</p>
</td>
</tr>
<tr>
<td>
<code>forceUpdateTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForceUpdateTag: How the extension handler should be forced to update even if the extension configuration has not changed.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>instanceView</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_Status">
VirtualMachineExtensionInstanceView_Status
</a>
</em>
</td>
<td>
<p>InstanceView: The virtual machine extension instance view.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
string
</em>
</td>
<td>
<p>PropertiesType: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>protectedSettings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
settings at all.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The name of the extension handler publisher.</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Json formatted public settings for the extension.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineExtension_StatusARM">VirtualMachineExtension_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionProperties_StatusARM">
VirtualMachineExtensionProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_Status">VirtualMachineHealthStatus_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Status: The health status information for the VM.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_StatusARM">VirtualMachineHealthStatus_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Status: The health status information for the VM.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentity">VirtualMachineIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineIdentity">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineIdentity</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityType">
VirtualMachineIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an
implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the
virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Virtual Machine. The user identity dictionary
key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentityARM">VirtualMachineIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_SpecARM">VirtualMachines_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineIdentity">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineIdentity</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityType">
VirtualMachineIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an
implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the
virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Virtual Machine. The user identity dictionary
key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentityStatusType">VirtualMachineIdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status">VirtualMachineIdentity_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_StatusARM">VirtualMachineIdentity_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentityType">VirtualMachineIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity">VirtualMachineIdentity</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityARM">VirtualMachineIdentityARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status">VirtualMachineIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of virtual machine identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id associated with the virtual machine. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityStatusType">
VirtualMachineIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an
implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the
virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentities">
map[string]./api/compute/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentities
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Virtual Machine. The user identity dictionary
key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentity_StatusARM">VirtualMachineIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of virtual machine identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id associated with the virtual machine. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityStatusType">
VirtualMachineIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an
implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the
virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentitiesARM">
map[string]./api/compute/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentitiesARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Virtual Machine. The user identity dictionary
key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentities">VirtualMachineIdentity_Status_UserAssignedIdentities
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status">VirtualMachineIdentity_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status_UserAssignedIdentitiesARM">VirtualMachineIdentity_Status_UserAssignedIdentitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_StatusARM">VirtualMachineIdentity_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineInstanceViewStatusHyperVGeneration">VirtualMachineInstanceViewStatusHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;V1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;V2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>assignedHost</code><br/>
<em>
string
</em>
</td>
<td>
<p>AssignedHost: Resource id of the dedicated host, on which the virtual machine is allocated through automatic placement,
when the virtual machine is associated with a dedicated host group that has automatic placement enabled.
Minimum api-version: 2020-06-01.</p>
</td>
</tr>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_Status">
BootDiagnosticsInstanceView_Status
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: The computer name assigned to the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>disks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskInstanceView_Status">
[]DiskInstanceView_Status
</a>
</em>
</td>
<td>
<p>Disks: The virtual machine disk information.</p>
</td>
</tr>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_Status">
[]VirtualMachineExtensionInstanceView_Status
</a>
</em>
</td>
<td>
<p>Extensions: The extensions information.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceViewStatusHyperVGeneration">
VirtualMachineInstanceViewStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGeneration Type associated with a resource</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceRedeployStatus</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_Status">
MaintenanceRedeployStatus_Status
</a>
</em>
</td>
<td>
<p>MaintenanceRedeployStatus: The Maintenance Operation status on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osName</code><br/>
<em>
string
</em>
</td>
<td>
<p>OsName: The Operating System running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OsVersion: The version of Operating System running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>patchStatus</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_Status">
VirtualMachinePatchStatus_Status
</a>
</em>
</td>
<td>
<p>PatchStatus: [Preview Feature] The status of virtual machine patch operations.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the fault domain of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>platformUpdateDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformUpdateDomain: Specifies the update domain of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>rdpThumbPrint</code><br/>
<em>
string
</em>
</td>
<td>
<p>RdpThumbPrint: The Remote desktop certificate thumbprint.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>vmAgent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_Status">
VirtualMachineAgentInstanceView_Status
</a>
</em>
</td>
<td>
<p>VmAgent: The VM Agent running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>vmHealth</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_Status">
VirtualMachineHealthStatus_Status
</a>
</em>
</td>
<td>
<p>VmHealth: The health status for the VM.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>assignedHost</code><br/>
<em>
string
</em>
</td>
<td>
<p>AssignedHost: Resource id of the dedicated host, on which the virtual machine is allocated through automatic placement,
when the virtual machine is associated with a dedicated host group that has automatic placement enabled.
Minimum api-version: 2020-06-01.</p>
</td>
</tr>
<tr>
<td>
<code>bootDiagnostics</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BootDiagnosticsInstanceView_StatusARM">
BootDiagnosticsInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to view Console Output and Screenshot to
diagnose VM status.
You can easily view the output of your console log.
Azure also enables you to see a screenshot of the VM from the hypervisor.</p>
</td>
</tr>
<tr>
<td>
<code>computerName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerName: The computer name assigned to the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>disks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskInstanceView_StatusARM">
[]DiskInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>Disks: The virtual machine disk information.</p>
</td>
</tr>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtensionInstanceView_StatusARM">
[]VirtualMachineExtensionInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>Extensions: The extensions information.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceViewStatusHyperVGeneration">
VirtualMachineInstanceViewStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGeneration Type associated with a resource</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceRedeployStatus</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.MaintenanceRedeployStatus_StatusARM">
MaintenanceRedeployStatus_StatusARM
</a>
</em>
</td>
<td>
<p>MaintenanceRedeployStatus: The Maintenance Operation status on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osName</code><br/>
<em>
string
</em>
</td>
<td>
<p>OsName: The Operating System running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OsVersion: The version of Operating System running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>patchStatus</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_StatusARM">
VirtualMachinePatchStatus_StatusARM
</a>
</em>
</td>
<td>
<p>PatchStatus: [Preview Feature] The status of virtual machine patch operations.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the fault domain of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>platformUpdateDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformUpdateDomain: Specifies the update domain of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>rdpThumbPrint</code><br/>
<em>
string
</em>
</td>
<td>
<p>RdpThumbPrint: The Remote desktop certificate thumbprint.</p>
</td>
</tr>
<tr>
<td>
<code>statuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>Statuses: The resource status information.</p>
</td>
</tr>
<tr>
<td>
<code>vmAgent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineAgentInstanceView_StatusARM">
VirtualMachineAgentInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>VmAgent: The VM Agent running on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>vmHealth</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineHealthStatus_StatusARM">
VirtualMachineHealthStatus_StatusARM
</a>
</em>
</td>
<td>
<p>VmHealth: The health status for the VM.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_Status">VirtualMachinePatchStatus_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">VirtualMachineInstanceView_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availablePatchSummary</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_Status">
AvailablePatchSummary_Status
</a>
</em>
</td>
<td>
<p>AvailablePatchSummary: The available patch summary of the latest assessment operation for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>configurationStatuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_Status">
[]InstanceViewStatus_Status
</a>
</em>
</td>
<td>
<p>ConfigurationStatuses: The enablement status of the specified patchMode</p>
</td>
</tr>
<tr>
<td>
<code>lastPatchInstallationSummary</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_Status">
LastPatchInstallationSummary_Status
</a>
</em>
</td>
<td>
<p>LastPatchInstallationSummary: The installation summary of the latest installation operation for the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachinePatchStatus_StatusARM">VirtualMachinePatchStatus_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">VirtualMachineInstanceView_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availablePatchSummary</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AvailablePatchSummary_StatusARM">
AvailablePatchSummary_StatusARM
</a>
</em>
</td>
<td>
<p>AvailablePatchSummary: The available patch summary of the latest assessment operation for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>configurationStatuses</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.InstanceViewStatus_StatusARM">
[]InstanceViewStatus_StatusARM
</a>
</em>
</td>
<td>
<p>ConfigurationStatuses: The enablement status of the specified patchMode</p>
</td>
</tr>
<tr>
<td>
<code>lastPatchInstallationSummary</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LastPatchInstallationSummary_StatusARM">
LastPatchInstallationSummary_StatusARM
</a>
</em>
</td>
<td>
<p>LastPatchInstallationSummary: The installation summary of the latest installation operation for the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">VirtualMachineProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities_StatusARM">
AdditionalCapabilities_StatusARM
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Specifies additional capabilities enabled or disabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>availabilitySet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>AvailabilitySet: Specifies information about the availability set that the virtual machine should be assigned to.
Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For
more information about availability sets, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">Manage the availability of virtual
machines</a>.
For more information on Azure planned maintenance, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">Planned maintenance for virtual machines in
Azure</a>
Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being
added should be under the same resource group as the availability set resource. An existing VM cannot be added to an
availability set.
This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.</p>
</td>
</tr>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile_StatusARM">
BillingProfile_StatusARM
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot virtual machine.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_StatusARM">
DiagnosticsProfile_StatusARM
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.EvictionPolicy_Status">
EvictionPolicy_Status
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
<tr>
<td>
<code>hardwareProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfile_StatusARM">
HardwareProfile_StatusARM
</a>
</em>
</td>
<td>
<p>HardwareProfile: Specifies the hardware settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Host: Specifies information about the dedicated host that the virtual machine resides in.
Minimum api-version: 2018-10-01.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>HostGroup: Specifies information about the dedicated host group that the virtual machine resides in.
Minimum api-version: 2020-06-01.
NOTE: User cannot specify both host and hostGroup properties.</p>
</td>
</tr>
<tr>
<td>
<code>instanceView</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_StatusARM">
VirtualMachineInstanceView_StatusARM
</a>
</em>
</td>
<td>
<p>InstanceView: The virtual machine instance view.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkProfile_StatusARM">
NetworkProfile_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies the network interfaces of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSProfile_StatusARM">
OSProfile_StatusARM
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot
be changed once VM is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the scale set logical fault domain into which the Virtual Machine will be created. By
default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across
available fault domains.
<li>This is applicable only if the &lsquo;virtualMachineScaleSet&rsquo; property of this Virtual Machine is set.<li>The Virtual
Machine Scale Set that is referenced, must have &lsquo;platformFaultDomainCount&rsquo; &amp;gt; 1.<li>This property cannot be updated
once the Virtual Machine is created.<li>Fault domain assignment can be viewed in the Virtual Machine Instance View.
Minimum api‐version: 2020‐12‐01</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Priority_Status">
Priority_Status
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machine.
Minimum api-version: 2019-03-01</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroup: Specifies information about the proximity placement group that the virtual machine should be
assigned to.
Minimum api-version: 2018-04-01.</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile_StatusARM">
SecurityProfile_StatusARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security related profile settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageProfile_StatusARM">
StorageProfile_StatusARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineScaleSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>VirtualMachineScaleSet: Specifies information about the virtual machine scale set that the virtual machine should be
assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to
maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM
cannot be added to a virtual machine scale set.
This property cannot exist along with a non-null properties.availabilitySet reference.
Minimum api‐version: 2019‐03‐01</p>
</td>
</tr>
<tr>
<td>
<code>vmId</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmId: Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS
and can be read using platform BIOS commands.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSet">VirtualMachineScaleSet
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/resourceDefinitions/virtualMachineScaleSets">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/resourceDefinitions/virtualMachineScaleSets</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">
VirtualMachineScaleSets_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities">
AdditionalCapabilities
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>automaticRepairsPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticRepairsPolicy">
AutomaticRepairsPolicy
</a>
</em>
</td>
<td>
<p>AutomaticRepairsPolicy: Specifies the configuration parameters for automatic repairs on the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>doNotRunExtensionsOnOverprovisionedVMs</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
overprovisioned VMs.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity">
VirtualMachineScaleSetIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>orchestrationMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesOrchestrationMode">
VirtualMachineScaleSetsSpecPropertiesOrchestrationMode
</a>
</em>
</td>
<td>
<p>OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>overprovision</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan">
Plan
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomainCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomainCount: Fault Domain count for each placement group.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scaleInPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicy">
ScaleInPolicy
</a>
</em>
</td>
<td>
<p>ScaleInPolicy: Describes a scale-in policy for a virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>singlePlacementGroup</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
not be modified to true.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Describes a virtual machine scale set sku. NOTE: If the new VM SKU is not supported on the hardware the scale set
is currently on, you need to deallocate the VMs in the scale set before you modify the SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>upgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicy">
UpgradePolicy
</a>
</em>
</td>
<td>
<p>UpgradePolicy: Describes an upgrade policy - automatic, manual, or rolling.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile
</a>
</em>
</td>
<td>
<p>VirtualMachineProfile: Describes a virtual machine scale set virtual machine profile.</p>
</td>
</tr>
<tr>
<td>
<code>zoneBalance</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">
VirtualMachineScaleSet_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk">VirtualMachineScaleSetDataDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile">VirtualMachineScaleSetStorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetDataDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetDataDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCaching">
VirtualMachineScaleSetDataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCreateOption">
VirtualMachineScaleSetDataDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: The create option.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is
UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk. Should be used only when
StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters">
VirtualMachineScaleSetManagedDiskParameters
</a>
</em>
</td>
<td>
<p>ManagedDisk: Describes the parameters of a ScaleSet managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskARM">VirtualMachineScaleSetDataDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfileARM">VirtualMachineScaleSetStorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetDataDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetDataDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCaching">
VirtualMachineScaleSetDataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCreateOption">
VirtualMachineScaleSetDataDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: The create option.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is
UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk. Should be used only when
StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersARM">
VirtualMachineScaleSetManagedDiskParametersARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: Describes the parameters of a ScaleSet managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCaching">VirtualMachineScaleSetDataDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk">VirtualMachineScaleSetDataDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskARM">VirtualMachineScaleSetDataDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskCreateOption">VirtualMachineScaleSetDataDiskCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk">VirtualMachineScaleSetDataDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskARM">VirtualMachineScaleSetDataDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_Status">VirtualMachineScaleSetDataDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_Status">VirtualMachineScaleSetStorageProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: The create option.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is
UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk. Should be used only when
StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status">
VirtualMachineScaleSetManagedDiskParameters_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_StatusARM">VirtualMachineScaleSetDataDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_StatusARM">VirtualMachineScaleSetStorageProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: The create option.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is
UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk. Should be used only when
StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_StatusARM">
VirtualMachineScaleSetManagedDiskParameters_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_Status">VirtualMachineScaleSetExtensionProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtension_Status">
[]VirtualMachineScaleSetExtension_Status
</a>
</em>
</td>
<td>
<p>Extensions: The virtual machine scale set child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_StatusARM">VirtualMachineScaleSetExtensionProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtension_StatusARM">
[]VirtualMachineScaleSetExtension_StatusARM
</a>
</em>
</td>
<td>
<p>Extensions: The virtual machine scale set child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProperties_StatusARM">VirtualMachineScaleSetExtensionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtension_StatusARM">VirtualMachineScaleSetExtension_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoUpgradeMinorVersion</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
property set to true.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
newer version of the extension available.</p>
</td>
</tr>
<tr>
<td>
<code>forceUpdateTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForceUpdateTag: If a value is provided and is different from the previous value, the extension handler will be forced to
update even if the extension configuration has not changed.</p>
</td>
</tr>
<tr>
<td>
<code>protectedSettings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
settings at all.</p>
</td>
</tr>
<tr>
<td>
<code>provisionAfterExtensions</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ProvisionAfterExtensions: Collection of extension names after which this extension needs to be provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The name of the extension handler publisher.</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Json formatted public settings for the extension.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtension_Status">VirtualMachineScaleSetExtension_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_Status">VirtualMachineScaleSetExtensionProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoUpgradeMinorVersion</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
property set to true.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpgrade</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
newer version of the extension available.</p>
</td>
</tr>
<tr>
<td>
<code>forceUpdateTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForceUpdateTag: If a value is provided and is different from the previous value, the extension handler will be forced to
update even if the extension configuration has not changed.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extension.</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
string
</em>
</td>
<td>
<p>PropertiesType: Specifies the type of the extension; an example is &ldquo;CustomScriptExtension&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>protectedSettings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
settings at all.</p>
</td>
</tr>
<tr>
<td>
<code>provisionAfterExtensions</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ProvisionAfterExtensions: Collection of extension names after which this extension needs to be provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The name of the extension handler publisher.</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Json formatted public settings for the extension.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Specifies the version of the script handler.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtension_StatusARM">VirtualMachineScaleSetExtension_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_StatusARM">VirtualMachineScaleSetExtensionProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extension.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProperties_StatusARM">
VirtualMachineScaleSetExtensionProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion">VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">VirtualMachineScaleSetIPConfigurationProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">VirtualMachineScaleSetIPConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">VirtualMachineScaleSetIPConfigurationProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_StatusARM">VirtualMachineScaleSetIPConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationGatewayBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: Specifies an array of references to backend address pools of application
gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot
use the same application gateway.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Specifies an array of references to application security group.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: Specifies an array of references to backend address pools of load balancers. A scale
set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the
same basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatPools: Specifies an array of references to inbound Nat pools of the load balancers. A scale set
can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same
basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion">
VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Available from Api-Version 2017-03-30 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM">
VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPAddressConfiguration: The publicIPAddressConfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference_StatusARM">
ApiEntityReference_StatusARM
</a>
</em>
</td>
<td>
<p>Subnet: Specifies the identifier of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">VirtualMachineScaleSetIPConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status">VirtualMachineScaleSetNetworkConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationGatewayBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: Specifies an array of references to backend address pools of application
gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot
use the same application gateway.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Specifies an array of references to application security group.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: Specifies an array of references to backend address pools of load balancers. A scale
set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the
same basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatPools: Specifies an array of references to inbound Nat pools of the load balancers. A scale set
can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same
basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The IP configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion">
VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Available from Api-Version 2017-03-30 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">
VirtualMachineScaleSetPublicIPAddressConfiguration_Status
</a>
</em>
</td>
<td>
<p>PublicIPAddressConfiguration: The publicIPAddressConfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference_Status">
ApiEntityReference_Status
</a>
</em>
</td>
<td>
<p>Subnet: Specifies the identifier of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_StatusARM">VirtualMachineScaleSetIPConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM">VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The IP configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">
VirtualMachineScaleSetIPConfigurationProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity">VirtualMachineScaleSetIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIdentity">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIdentity</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityType">
VirtualMachineScaleSetIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine scale set. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both
an implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from
the virtual machine scale set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityARM">VirtualMachineScaleSetIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIdentity">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIdentity</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityType">
VirtualMachineScaleSetIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine scale set. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both
an implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from
the virtual machine scale set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityStatusType">VirtualMachineScaleSetIdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status">VirtualMachineScaleSetIdentity_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_StatusARM">VirtualMachineScaleSetIdentity_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityType">VirtualMachineScaleSetIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity">VirtualMachineScaleSetIdentity</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityARM">VirtualMachineScaleSetIdentityARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status">VirtualMachineScaleSetIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of virtual machine scale set identity. This property will only be provided for a system
assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id associated with the virtual machine scale set. This property will only be provided for a system
assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityStatusType">
VirtualMachineScaleSetIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine scale set. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both
an implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from
the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentities">
map[string]./api/compute/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentities
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the virtual machine scale set. The user identity
dictionary key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_StatusARM">VirtualMachineScaleSetIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of virtual machine scale set identity. This property will only be provided for a system
assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id associated with the virtual machine scale set. This property will only be provided for a system
assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityStatusType">
VirtualMachineScaleSetIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the virtual machine scale set. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both
an implicitly created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from
the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentitiesARM">
map[string]./api/compute/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentitiesARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the virtual machine scale set. The user identity
dictionary key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentities">VirtualMachineScaleSetIdentity_Status_UserAssignedIdentities
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status">VirtualMachineScaleSetIdentity_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status_UserAssignedIdentitiesARM">VirtualMachineScaleSetIdentity_Status_UserAssignedIdentitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_StatusARM">VirtualMachineScaleSetIdentity_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag">VirtualMachineScaleSetIpTag
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIpTag">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIpTag</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: IP tag type. Example: FirstPartyUsage.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: IP tag associated with the public IP. Example: SQL, Storage etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTagARM">VirtualMachineScaleSetIpTagARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIpTag">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetIpTag</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: IP tag type. Example: FirstPartyUsage.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: IP tag associated with the public IP. Example: SQL, Storage etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag_Status">VirtualMachineScaleSetIpTag_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">VirtualMachineScaleSetPublicIPAddressConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: IP tag type. Example: FirstPartyUsage.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: IP tag associated with the public IP. Example: SQL, Storage etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag_StatusARM">VirtualMachineScaleSetIpTag_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: IP tag type. Example: FirstPartyUsage.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: IP tag associated with the public IP. Example: SQL, Storage etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters">VirtualMachineScaleSetManagedDiskParameters
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk">VirtualMachineScaleSetDataDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetManagedDiskParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetManagedDiskParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersStorageAccountType">
VirtualMachineScaleSetManagedDiskParametersStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersARM">VirtualMachineScaleSetManagedDiskParametersARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskARM">VirtualMachineScaleSetDataDiskARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetManagedDiskParameters">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetManagedDiskParameters</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersStorageAccountType">
VirtualMachineScaleSetManagedDiskParametersStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersStorageAccountType">VirtualMachineScaleSetManagedDiskParametersStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters">VirtualMachineScaleSetManagedDiskParameters</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersARM">VirtualMachineScaleSetManagedDiskParametersARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status">VirtualMachineScaleSetManagedDiskParameters_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_Status">VirtualMachineScaleSetDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_StatusARM">VirtualMachineScaleSetManagedDiskParameters_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_StatusARM">VirtualMachineScaleSetDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings">VirtualMachineScaleSetNetworkConfigurationDnsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetNetworkConfigurationDnsSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetNetworkConfigurationDnsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettingsARM">VirtualMachineScaleSetNetworkConfigurationDnsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetNetworkConfigurationDnsSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetNetworkConfigurationDnsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status">VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status">VirtualMachineScaleSetNetworkConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM">VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM">VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM">VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_StatusARM">VirtualMachineScaleSetNetworkConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM">
VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The dns settings to be applied on the network interfaces.</p>
</td>
</tr>
<tr>
<td>
<code>enableAcceleratedNetworking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAcceleratedNetworking: Specifies whether the network interface is accelerated networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableFpga</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFpga: Specifies whether the network interface is FPGA networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableIPForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableIPForwarding: Whether IP forwarding enabled on this NIC.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_StatusARM">
[]VirtualMachineScaleSetIPConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: Specifies the IP configurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The network security group.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status">VirtualMachineScaleSetNetworkConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_Status">VirtualMachineScaleSetNetworkProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status">
VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status
</a>
</em>
</td>
<td>
<p>DnsSettings: The dns settings to be applied on the network interfaces.</p>
</td>
</tr>
<tr>
<td>
<code>enableAcceleratedNetworking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAcceleratedNetworking: Specifies whether the network interface is accelerated networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableFpga</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFpga: Specifies whether the network interface is FPGA networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableIPForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableIPForwarding: Whether IP forwarding enabled on this NIC.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">
[]VirtualMachineScaleSetIPConfiguration_Status
</a>
</em>
</td>
<td>
<p>IpConfigurations: Specifies the IP configurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The network configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The network security group.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_StatusARM">VirtualMachineScaleSetNetworkConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_StatusARM">VirtualMachineScaleSetNetworkProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The network configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM">
VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_Status">VirtualMachineScaleSetNetworkProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>healthProbe</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference_Status">
ApiEntityReference_Status
</a>
</em>
</td>
<td>
<p>HealthProbe: A reference to a load balancer probe used to determine the health of an instance in the virtual machine
scale set. The reference will be in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;loadBalancers/&#x200b;{loadBalancerName}/&#x200b;probes/&#x200b;{probeName}&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaceConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status">
[]VirtualMachineScaleSetNetworkConfiguration_Status
</a>
</em>
</td>
<td>
<p>NetworkInterfaceConfigurations: The list of network configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_StatusARM">VirtualMachineScaleSetNetworkProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>healthProbe</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference_StatusARM">
ApiEntityReference_StatusARM
</a>
</em>
</td>
<td>
<p>HealthProbe: A reference to a load balancer probe used to determine the health of an instance in the virtual machine
scale set. The reference will be in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;loadBalancers/&#x200b;{loadBalancerName}/&#x200b;probes/&#x200b;{probeName}&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaceConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_StatusARM">
[]VirtualMachineScaleSetNetworkConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkInterfaceConfigurations: The list of network configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile">VirtualMachineScaleSetStorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCaching">
VirtualMachineScaleSetOSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCreateOption">
VirtualMachineScaleSetOSDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machines in the scale set should be created.
The only allowed value is: FromImage \u2013 This value is used when you are using an image to create the virtual
machine. If you are using a platform image, you also use the imageReference element described above. If you are using a
marketplace image, you  also use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings">
DiffDiskSettings
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Describes the parameters of ephemeral disk settings that can be specified for operating system disk.
NOTE: The ephemeral disk settings can only be specified for managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of the operating system disk in gigabytes. This element can be used to overwrite the size
of the disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk">
VirtualHardDisk
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters">
VirtualMachineScaleSetManagedDiskParameters
</a>
</em>
</td>
<td>
<p>ManagedDisk: Describes the parameters of a ScaleSet managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskOsType">
VirtualMachineScaleSetOSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>vhdContainers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>VhdContainers: Specifies the container urls that are used to store operating system disks for the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfileARM">VirtualMachineScaleSetStorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSDisk">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSDisk</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCaching">
VirtualMachineScaleSetOSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCreateOption">
VirtualMachineScaleSetOSDiskCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machines in the scale set should be created.
The only allowed value is: FromImage \u2013 This value is used when you are using an image to create the virtual
machine. If you are using a platform image, you also use the imageReference element described above. If you are using a
marketplace image, you  also use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettingsARM">
DiffDiskSettingsARM
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Describes the parameters of ephemeral disk settings that can be specified for operating system disk.
NOTE: The ephemeral disk settings can only be specified for managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of the operating system disk in gigabytes. This element can be used to overwrite the size
of the disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDiskARM">
VirtualHardDiskARM
</a>
</em>
</td>
<td>
<p>Image: Describes the uri of a disk.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParametersARM">
VirtualMachineScaleSetManagedDiskParametersARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: Describes the parameters of a ScaleSet managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskOsType">
VirtualMachineScaleSetOSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>vhdContainers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>VhdContainers: Specifies the container urls that are used to store operating system disks for the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCaching">VirtualMachineScaleSetOSDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskCreateOption">VirtualMachineScaleSetOSDiskCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskOsType">VirtualMachineScaleSetOSDiskOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">VirtualMachineScaleSetOSDisk</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">VirtualMachineScaleSetOSDiskARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskStatusOsType">VirtualMachineScaleSetOSDiskStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">VirtualMachineScaleSetOSDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_Status">VirtualMachineScaleSetStorageProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machines in the scale set should be created.
The only allowed value is: FromImage \u2013 This value is used when you are using an image to create the virtual
machine. If you are using a platform image, you also use the imageReference element described above. If you are using a
marketplace image, you  also use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_Status">
DiffDiskSettings_Status
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Specifies the ephemeral disk Settings for the operating system disk used by the virtual machine scale
set.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of the operating system disk in gigabytes. This element can be used to overwrite the size
of the disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_Status">
VirtualHardDisk_Status
</a>
</em>
</td>
<td>
<p>Image: Specifies information about the unmanaged user image to base the scale set on.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status">
VirtualMachineScaleSetManagedDiskParameters_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskStatusOsType">
VirtualMachineScaleSetOSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>vhdContainers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>VhdContainers: Specifies the container urls that are used to store operating system disks for the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">VirtualMachineScaleSetOSDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_StatusARM">VirtualMachineScaleSetStorageProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Caching_Status">
Caching_Status
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.CreateOption_Status">
CreateOption_Status
</a>
</em>
</td>
<td>
<p>CreateOption: Specifies how the virtual machines in the scale set should be created.
The only allowed value is: FromImage \u2013 This value is used when you are using an image to create the virtual
machine. If you are using a platform image, you also use the imageReference element described above. If you are using a
marketplace image, you  also use the plan element previously described.</p>
</td>
</tr>
<tr>
<td>
<code>diffDiskSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiffDiskSettings_StatusARM">
DiffDiskSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DiffDiskSettings: Specifies the ephemeral disk Settings for the operating system disk used by the virtual machine scale
set.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeGB: Specifies the size of the operating system disk in gigabytes. This element can be used to overwrite the size
of the disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualHardDisk_StatusARM">
VirtualHardDisk_StatusARM
</a>
</em>
</td>
<td>
<p>Image: Specifies information about the unmanaged user image to base the scale set on.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_StatusARM">
VirtualMachineScaleSetManagedDiskParameters_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managed disk parameters.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The disk name.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskStatusOsType">
VirtualMachineScaleSetOSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
user-image or a specialized VHD.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>vhdContainers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>VhdContainers: Specifies the container urls that are used to store operating system disks for the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>writeAcceleratorEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile">VirtualMachineScaleSetOSProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>computerNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerNamePrefix: Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name
prefixes must be 1 to 15 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
For using cloud-init for your VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration">
LinuxConfiguration
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup">
[]VaultSecretGroup
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration">
WindowsConfiguration
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfileARM">VirtualMachineScaleSetOSProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetOSProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>computerNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerNamePrefix: Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name
prefixes must be 1 to 15 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
For using cloud-init for your VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfigurationARM">
LinuxConfigurationARM
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroupARM">
[]VaultSecretGroupARM
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfigurationARM">
WindowsConfigurationARM
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_Status">VirtualMachineScaleSetOSProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>computerNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerNamePrefix: Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name
prefixes must be 1 to 15 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
For using cloud-init for your VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_Status">
LinuxConfiguration_Status
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_Status">
[]VaultSecretGroup_Status
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_Status">
WindowsConfiguration_Status
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_StatusARM">VirtualMachineScaleSetOSProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length (Windows): 8 characters
Minimum-length (Linux): 6 characters
Max-length (Windows): 123 characters
Max-length (Linux): 72 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;
For resetting the password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">How to reset the Remote Desktop service or its login password in a Windows
VM</a>
For resetting root password, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password">Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess
Extension</a></p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Windows-only restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length (Linux): 1  character
Max-length (Linux): 64 characters
Max-length (Windows): 20 characters
<li> For root access to the Linux VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using root privileges on Linux virtual machines in
Azure</a>
<li> For a list of built-in system users on Linux that should not be used in this field, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Selecting User Names for
Linux on
Azure</a></p>
</td>
</tr>
<tr>
<td>
<code>computerNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>ComputerNamePrefix: Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name
prefixes must be 1 to 15 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>customData</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
For using cloud-init for your VM, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Using cloud-init to customize a Linux VM during
creation</a></p>
</td>
</tr>
<tr>
<td>
<code>linuxConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.LinuxConfiguration_StatusARM">
LinuxConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
For a list of supported Linux distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Linux on Azure-Endorsed
Distributions</a>
For running non-endorsed distributions, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json">Information for Non-Endorsed
Distributions</a>.</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VaultSecretGroup_StatusARM">
[]VaultSecretGroup_StatusARM
</a>
</em>
</td>
<td>
<p>Secrets: Specifies set of certificates that should be installed onto the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>windowsConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">
WindowsConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities_StatusARM">
AdditionalCapabilities_StatusARM
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual
Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data
disks with UltraSSD_LRS storage account type.</p>
</td>
</tr>
<tr>
<td>
<code>automaticRepairsPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticRepairsPolicy_StatusARM">
AutomaticRepairsPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>AutomaticRepairsPolicy: Policy for automatic repairs.</p>
</td>
</tr>
<tr>
<td>
<code>doNotRunExtensionsOnOverprovisionedVMs</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
overprovisioned VMs.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>HostGroup: Specifies information about the dedicated host group that the virtual machine scale set resides in.
Minimum api-version: 2020-06-01.</p>
</td>
</tr>
<tr>
<td>
<code>orchestrationMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OrchestrationMode_Status">
OrchestrationMode_Status
</a>
</em>
</td>
<td>
<p>OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>overprovision</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomainCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomainCount: Fault Domain count for each placement group.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroup: Specifies information about the proximity placement group that the virtual machine scale set
should be assigned to.
Minimum api-version: 2018-04-01.</p>
</td>
</tr>
<tr>
<td>
<code>scaleInPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicy_StatusARM">
ScaleInPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>ScaleInPolicy: Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual
Machine Scale Set is scaled-in.</p>
</td>
</tr>
<tr>
<td>
<code>singlePlacementGroup</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
not be modified to true.</p>
</td>
</tr>
<tr>
<td>
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Specifies the ID which uniquely identifies a Virtual Machine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>upgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicy_StatusARM">
UpgradePolicy_StatusARM
</a>
</em>
</td>
<td>
<p>UpgradePolicy: The upgrade policy.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">
VirtualMachineScaleSetVMProfile_StatusARM
</a>
</em>
</td>
<td>
<p>VirtualMachineProfile: The virtual machine profile.</p>
</td>
</tr>
<tr>
<td>
<code>zoneBalance</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings">VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>domainNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>DomainNameLabel: The Domain name label.The concatenation of the domain name label and vm index will be the domain name
labels of the PublicIPAddress resources that will be created</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsARM">VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>domainNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>DomainNameLabel: The Domain name label.The concatenation of the domain name label and vm index will be the domain name
labels of the PublicIPAddress resources that will be created</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status">VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">VirtualMachineScaleSetPublicIPAddressConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>domainNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>DomainNameLabel: The Domain name label.The concatenation of the domain name label and vm index will be the domain name
labels of the PublicIPAddress resources that will be created</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>domainNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>DomainNameLabel: The Domain name label.The concatenation of the domain name label and vm index will be the domain name
labels of the PublicIPAddress resources that will be created</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationProperties">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetPublicIPAddressConfigurationProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsARM">
VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsARM
</a>
</em>
</td>
<td>
<p>DnsSettings: Describes a virtual machines scale sets network configuration&rsquo;s DNS settings.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: The idle timeout of the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTagARM">
[]VirtualMachineScaleSetIpTagARM
</a>
</em>
</td>
<td>
<p>IpTags: The list of IP tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion">
VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: Available from Api-Version 2019-07-01 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion">VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">VirtualMachineScaleSetPublicIPAddressConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM">VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM">
VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The dns settings to be applied on the publicIP addresses .</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: The idle timeout of the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag_StatusARM">
[]VirtualMachineScaleSetIpTag_StatusARM
</a>
</em>
</td>
<td>
<p>IpTags: The list of IP tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion">
VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: Available from Api-Version 2019-07-01 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The PublicIPPrefix from which to allocate publicIP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status">VirtualMachineScaleSetPublicIPAddressConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status">VirtualMachineScaleSetIPConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status">
VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status
</a>
</em>
</td>
<td>
<p>DnsSettings: The dns settings to be applied on the publicIP addresses .</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: The idle timeout of the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag_Status">
[]VirtualMachineScaleSetIpTag_Status
</a>
</em>
</td>
<td>
<p>IpTags: The list of IP tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The publicIP address configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion">
VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: Available from Api-Version 2019-07-01 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The PublicIPPrefix from which to allocate publicIP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM">VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_StatusARM">VirtualMachineScaleSetIPConfigurationProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The publicIP address configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM">
VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile">VirtualMachineScaleSetStorageProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetStorageProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetStorageProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk">
[]VirtualMachineScaleSetDataDisk
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add data disks to the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference">
ImageReference
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations. NOTE: Image reference
publisher and offer can only be set when you create the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk">
VirtualMachineScaleSetOSDisk
</a>
</em>
</td>
<td>
<p>OsDisk: Describes a virtual machine scale set operating system disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfileARM">VirtualMachineScaleSetStorageProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetStorageProfile">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/VirtualMachineScaleSetStorageProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDiskARM">
[]VirtualMachineScaleSetDataDiskARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add data disks to the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReferenceARM">
ImageReferenceARM
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations. NOTE: Image reference
publisher and offer can only be set when you create the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDiskARM">
VirtualMachineScaleSetOSDiskARM
</a>
</em>
</td>
<td>
<p>OsDisk: Describes a virtual machine scale set operating system disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_Status">VirtualMachineScaleSetStorageProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_Status">
[]VirtualMachineScaleSetDataDisk_Status
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add data disks to the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference_Status">
ImageReference_Status
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_Status">
VirtualMachineScaleSetOSDisk_Status
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_StatusARM">VirtualMachineScaleSetStorageProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetDataDisk_StatusARM">
[]VirtualMachineScaleSetDataDisk_StatusARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add data disks to the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ImageReference_StatusARM">
ImageReference_StatusARM
</a>
</em>
</td>
<td>
<p>ImageReference: Specifies information about the image to use. You can specify information about platform images,
marketplace images, or virtual machine images. This element is required when you want to use a platform image,
marketplace image, or virtual machine image, but is not used in other creation operations.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSDisk_StatusARM">
VirtualMachineScaleSetOSDisk_StatusARM
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machines in the scale set.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">VirtualMachineScaleSetVMProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile_Status">
BillingProfile_Status
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_Status">
DiagnosticsProfile_Status
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.EvictionPolicy_Status">
EvictionPolicy_Status
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_Status">
VirtualMachineScaleSetExtensionProfile_Status
</a>
</em>
</td>
<td>
<p>ExtensionProfile: Specifies a collection of settings for extensions installed on virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_Status">
VirtualMachineScaleSetNetworkProfile_Status
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies properties of the network interfaces of the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_Status">
VirtualMachineScaleSetOSProfile_Status
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings for the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Priority_Status">
Priority_Status
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machines in the scale set.
Minimum api-version: 2017-10-30-preview</p>
</td>
</tr>
<tr>
<td>
<code>scheduledEventsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile_Status">
ScheduledEventsProfile_Status
</a>
</em>
</td>
<td>
<p>ScheduledEventsProfile: Specifies Scheduled Event related configurations.</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile_Status">
SecurityProfile_Status
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security related profile settings for the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_Status">
VirtualMachineScaleSetStorageProfile_Status
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_StatusARM">VirtualMachineScaleSetVMProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">VirtualMachineScaleSetProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile_StatusARM">
BillingProfile_StatusARM
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_StatusARM">
DiagnosticsProfile_StatusARM
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.EvictionPolicy_Status">
EvictionPolicy_Status
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetExtensionProfile_StatusARM">
VirtualMachineScaleSetExtensionProfile_StatusARM
</a>
</em>
</td>
<td>
<p>ExtensionProfile: Specifies a collection of settings for extensions installed on virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkProfile_StatusARM">
VirtualMachineScaleSetNetworkProfile_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies properties of the network interfaces of the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_StatusARM">
VirtualMachineScaleSetOSProfile_StatusARM
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings for the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Priority_Status">
Priority_Status
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machines in the scale set.
Minimum api-version: 2017-10-30-preview</p>
</td>
</tr>
<tr>
<td>
<code>scheduledEventsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile_StatusARM">
ScheduledEventsProfile_StatusARM
</a>
</em>
</td>
<td>
<p>ScheduledEventsProfile: Specifies Scheduled Event related configurations.</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile_StatusARM">
SecurityProfile_StatusARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security related profile settings for the virtual machines in the scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile_StatusARM">
VirtualMachineScaleSetStorageProfile_StatusARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSet_Status">VirtualMachineScaleSet_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet">VirtualMachineScaleSet</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities_Status">
AdditionalCapabilities_Status
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual
Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data
disks with UltraSSD_LRS storage account type.</p>
</td>
</tr>
<tr>
<td>
<code>automaticRepairsPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticRepairsPolicy_Status">
AutomaticRepairsPolicy_Status
</a>
</em>
</td>
<td>
<p>AutomaticRepairsPolicy: Policy for automatic repairs.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>doNotRunExtensionsOnOverprovisionedVMs</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
overprovisioned VMs.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>HostGroup: Specifies information about the dedicated host group that the virtual machine scale set resides in.
Minimum api-version: 2020-06-01.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_Status">
VirtualMachineScaleSetIdentity_Status
</a>
</em>
</td>
<td>
<p>Identity: The identity of the virtual machine scale set, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>orchestrationMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OrchestrationMode_Status">
OrchestrationMode_Status
</a>
</em>
</td>
<td>
<p>OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>overprovision</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan_Status">
Plan_Status
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomainCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomainCount: Fault Domain count for each placement group.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroup: Specifies information about the proximity placement group that the virtual machine scale set
should be assigned to.
Minimum api-version: 2018-04-01.</p>
</td>
</tr>
<tr>
<td>
<code>scaleInPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicy_Status">
ScaleInPolicy_Status
</a>
</em>
</td>
<td>
<p>ScaleInPolicy: Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual
Machine Scale Set is scaled-in.</p>
</td>
</tr>
<tr>
<td>
<code>singlePlacementGroup</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
not be modified to true.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: The virtual machine scale set sku.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Specifies the ID which uniquely identifies a Virtual Machine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>upgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicy_Status">
UpgradePolicy_Status
</a>
</em>
</td>
<td>
<p>UpgradePolicy: The upgrade policy.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetVMProfile_Status">
VirtualMachineScaleSetVMProfile_Status
</a>
</em>
</td>
<td>
<p>VirtualMachineProfile: The virtual machine profile.</p>
</td>
</tr>
<tr>
<td>
<code>zoneBalance</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSet_StatusARM">VirtualMachineScaleSet_StatusARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity_StatusARM">
VirtualMachineScaleSetIdentity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the virtual machine scale set, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan_StatusARM">
Plan_StatusARM
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetProperties_StatusARM">
VirtualMachineScaleSetProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The virtual machine scale set sku.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecAPIVersion">VirtualMachineScaleSetsSpecAPIVersion
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesOrchestrationMode">VirtualMachineScaleSetsSpecPropertiesOrchestrationMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Flexible&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Uniform&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy">VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion">VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority">VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Low&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSet">VirtualMachineScaleSet</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities">
AdditionalCapabilities
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>automaticRepairsPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticRepairsPolicy">
AutomaticRepairsPolicy
</a>
</em>
</td>
<td>
<p>AutomaticRepairsPolicy: Specifies the configuration parameters for automatic repairs on the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>doNotRunExtensionsOnOverprovisionedVMs</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
overprovisioned VMs.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentity">
VirtualMachineScaleSetIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>orchestrationMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesOrchestrationMode">
VirtualMachineScaleSetsSpecPropertiesOrchestrationMode
</a>
</em>
</td>
<td>
<p>OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>overprovision</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan">
Plan
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomainCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomainCount: Fault Domain count for each placement group.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scaleInPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicy">
ScaleInPolicy
</a>
</em>
</td>
<td>
<p>ScaleInPolicy: Describes a scale-in policy for a virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>singlePlacementGroup</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
not be modified to true.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Describes a virtual machine scale set sku. NOTE: If the new VM SKU is not supported on the hardware the scale set
is currently on, you need to deallocate the VMs in the scale set before you modify the SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>upgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicy">
UpgradePolicy
</a>
</em>
</td>
<td>
<p>UpgradePolicy: Describes an upgrade policy - automatic, manual, or rolling.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile
</a>
</em>
</td>
<td>
<p>VirtualMachineProfile: Describes a virtual machine scale set virtual machine profile.</p>
</td>
</tr>
<tr>
<td>
<code>zoneBalance</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIdentityARM">
VirtualMachineScaleSetIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the VM scale set to create or update.</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PlanARM">
PlanARM
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">
VirtualMachineScaleSets_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes the properties of a Virtual Machine Scale Set.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: Describes a virtual machine scale set sku. NOTE: If the new VM SKU is not supported on the hardware the scale set
is currently on, you need to deallocate the VMs in the scale set before you modify the SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_SpecARM">VirtualMachineScaleSets_SpecARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilitiesARM">
AdditionalCapabilitiesARM
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>automaticRepairsPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AutomaticRepairsPolicyARM">
AutomaticRepairsPolicyARM
</a>
</em>
</td>
<td>
<p>AutomaticRepairsPolicy: Specifies the configuration parameters for automatic repairs on the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>doNotRunExtensionsOnOverprovisionedVMs</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
overprovisioned VMs.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>orchestrationMode</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesOrchestrationMode">
VirtualMachineScaleSetsSpecPropertiesOrchestrationMode
</a>
</em>
</td>
<td>
<p>OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>overprovision</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomainCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomainCount: Fault Domain count for each placement group.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scaleInPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScaleInPolicyARM">
ScaleInPolicyARM
</a>
</em>
</td>
<td>
<p>ScaleInPolicy: Describes a scale-in policy for a virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>singlePlacementGroup</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
not be modified to true.</p>
</td>
</tr>
<tr>
<td>
<code>upgradePolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.UpgradePolicyARM">
UpgradePolicyARM
</a>
</em>
</td>
<td>
<p>UpgradePolicy: Describes an upgrade policy - automatic, manual, or rolling.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM
</a>
</em>
</td>
<td>
<p>VirtualMachineProfile: Describes a virtual machine scale set virtual machine profile.</p>
</td>
</tr>
<tr>
<td>
<code>zoneBalance</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec">VirtualMachineScaleSets_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile">
BillingProfile
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VM or VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile">
DiagnosticsProfile
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile
</a>
</em>
</td>
<td>
<p>ExtensionProfile: Describes a virtual machine scale set extension profile.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Describes a virtual machine scale set network profile.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile">
VirtualMachineScaleSetOSProfile
</a>
</em>
</td>
<td>
<p>OsProfile: Describes a virtual machine scale set OS profile.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machines in the scale set.
Minimum api-version: 2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledEventsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfile">
ScheduledEventsProfile
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile">
SecurityProfile
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security profile settings for the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfile">
VirtualMachineScaleSetStorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Describes a virtual machine scale set storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_PropertiesARM">VirtualMachineScaleSets_Spec_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfileARM">
BillingProfileARM
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VM or VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfileARM">
DiagnosticsProfileARM
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM
</a>
</em>
</td>
<td>
<p>ExtensionProfile: Describes a virtual machine scale set extension profile.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Describes a virtual machine scale set network profile.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfileARM">
VirtualMachineScaleSetOSProfileARM
</a>
</em>
</td>
<td>
<p>OsProfile: Describes a virtual machine scale set OS profile.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfilePriority
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machines in the scale set.
Minimum api-version: 2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledEventsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ScheduledEventsProfileARM">
ScheduledEventsProfileARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileARM">
SecurityProfileARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security profile settings for the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetStorageProfileARM">
VirtualMachineScaleSetStorageProfileARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Describes a virtual machine scale set storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions
</a>
</em>
</td>
<td>
<p>Extensions: The virtual machine scale set child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM
</a>
</em>
</td>
<td>
<p>Extensions: The virtual machine scale set child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_Extensions
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extension.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: Microsoft.Compute/extensions - Publisher</p>
</td>
</tr>
<tr>
<td>
<code>settings</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Settings: Microsoft.Compute/extensions - Settings</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Microsoft.Compute/extensions - Type</p>
</td>
</tr>
<tr>
<td>
<code>typeHandlerVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>TypeHandlerVersion: Microsoft.Compute/extensions - Type handler version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfile_ExtensionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_ExtensionProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the extension.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.GenericExtensionARM">
GenericExtensionARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>healthProbe</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference">
ApiEntityReference
</a>
</em>
</td>
<td>
<p>HealthProbe: The API entity reference.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaceConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations
</a>
</em>
</td>
<td>
<p>NetworkInterfaceConfigurations: The list of network configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>healthProbe</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReferenceARM">
ApiEntityReferenceARM
</a>
</em>
</td>
<td>
<p>HealthProbe: The API entity reference.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaceConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM
</a>
</em>
</td>
<td>
<p>NetworkInterfaceConfigurations: The list of network configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings">
VirtualMachineScaleSetNetworkConfigurationDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: Describes a virtual machines scale sets network configuration&rsquo;s DNS settings.</p>
</td>
</tr>
<tr>
<td>
<code>enableAcceleratedNetworking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAcceleratedNetworking: Specifies whether the network interface is accelerated networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableFpga</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFpga: Specifies whether the network interface is FPGA networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableIPForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableIPForwarding: Whether IP forwarding enabled on this NIC.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations
</a>
</em>
</td>
<td>
<p>IpConfigurations: Specifies the IP configurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The network configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The network configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes a virtual machine scale set network profile&rsquo;s IP configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurationsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettingsARM">
VirtualMachineScaleSetNetworkConfigurationDnsSettingsARM
</a>
</em>
</td>
<td>
<p>DnsSettings: Describes a virtual machines scale sets network configuration&rsquo;s DNS settings.</p>
</td>
</tr>
<tr>
<td>
<code>enableAcceleratedNetworking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAcceleratedNetworking: Specifies whether the network interface is accelerated networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableFpga</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFpga: Specifies whether the network interface is FPGA networking-enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableIPForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableIPForwarding: Whether IP forwarding enabled on this NIC.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM">
[]VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: Specifies the IP configurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationGatewayBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: Specifies an array of references to backend address pools of application
gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot
use the same application gateway.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Specifies an array of references to application security group.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: Specifies an array of references to backend address pools of load balancers. A scale
set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the
same basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatPools: Specifies an array of references to inbound Nat pools of the load balancers. A scale set
can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same
basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The IP configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Available from Api-Version 2017-03-30 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration
</a>
</em>
</td>
<td>
<p>PublicIPAddressConfiguration: Describes a virtual machines scale set IP Configuration&rsquo;s PublicIPAddress configuration</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReference">
ApiEntityReference
</a>
</em>
</td>
<td>
<p>Subnet: The API entity reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The IP configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes a virtual machine scale set network profile&rsquo;s IP configuration properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurationsARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationGatewayBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: Specifies an array of references to backend address pools of application
gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot
use the same application gateway.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Specifies an array of references to application security group.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: Specifies an array of references to backend address pools of load balancers. A scale
set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the
same basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatPools</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatPools: Specifies an array of references to inbound Nat pools of the load balancers. A scale set
can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same
basic sku load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion">
VirtualMachineScaleSetsSpecPropertiesVirtualMachineProfileNetworkProfileNetworkInterfaceConfigurationsPropertiesIpConfigurationsPropertiesPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Available from Api-Version 2017-03-30 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressConfiguration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM">
VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM
</a>
</em>
</td>
<td>
<p>PublicIPAddressConfiguration: Describes a virtual machines scale set IP Configuration&rsquo;s PublicIPAddress configuration</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ApiEntityReferenceARM">
ApiEntityReferenceARM
</a>
</em>
</td>
<td>
<p>Subnet: The API entity reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings">
VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: Describes a virtual machines scale sets network configuration&rsquo;s DNS settings.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: The idle timeout of the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetIpTag">
[]VirtualMachineScaleSetIpTag
</a>
</em>
</td>
<td>
<p>IpTags: The list of IP tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The publicIP address configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion">
VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: Available from Api-Version 2019-07-01 onwards, it represents whether the specific
ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: &lsquo;IPv4&rsquo; and &lsquo;IPv6&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_Properties_PublicIPAddressConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM">VirtualMachineScaleSets_Spec_Properties_VirtualMachineProfile_NetworkProfile_NetworkInterfaceConfigurations_Properties_IpConfigurations_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The publicIP address configuration name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM">
VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes a virtual machines scale set IP Configuration&rsquo;s PublicIPAddress configuration</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachine_Status">VirtualMachine_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine">VirtualMachine</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities_Status">
AdditionalCapabilities_Status
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Specifies additional capabilities enabled or disabled on the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>availabilitySet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>AvailabilitySet: Specifies information about the availability set that the virtual machine should be assigned to.
Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For
more information about availability sets, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">Manage the availability of virtual
machines</a>.
For more information on Azure planned maintenance, see <a href="https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json">Planned maintenance for virtual machines in
Azure</a>
Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being
added should be under the same resource group as the availability set resource. An existing VM cannot be added to an
availability set.
This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.</p>
</td>
</tr>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile_Status">
BillingProfile_Status
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot virtual machine.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile_Status">
DiagnosticsProfile_Status
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.EvictionPolicy_Status">
EvictionPolicy_Status
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
<tr>
<td>
<code>hardwareProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfile_Status">
HardwareProfile_Status
</a>
</em>
</td>
<td>
<p>HardwareProfile: Specifies the hardware settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Host: Specifies information about the dedicated host that the virtual machine resides in.
Minimum api-version: 2018-10-01.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>HostGroup: Specifies information about the dedicated host group that the virtual machine resides in.
Minimum api-version: 2020-06-01.
NOTE: User cannot specify both host and hostGroup properties.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_Status">
VirtualMachineIdentity_Status
</a>
</em>
</td>
<td>
<p>Identity: The identity of the virtual machine, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>instanceView</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineInstanceView_Status">
VirtualMachineInstanceView_Status
</a>
</em>
</td>
<td>
<p>InstanceView: The virtual machine instance view.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkProfile_Status">
NetworkProfile_Status
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies the network interfaces of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSProfile_Status">
OSProfile_Status
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot
be changed once VM is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan_Status">
Plan_Status
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the scale set logical fault domain into which the Virtual Machine will be created. By
default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across
available fault domains.
<li>This is applicable only if the &lsquo;virtualMachineScaleSet&rsquo; property of this Virtual Machine is set.<li>The Virtual
Machine Scale Set that is referenced, must have &lsquo;platformFaultDomainCount&rsquo; &amp;gt; 1.<li>This property cannot be updated
once the Virtual Machine is created.<li>Fault domain assignment can be viewed in the Virtual Machine Instance View.
Minimum api‐version: 2020‐12‐01</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Priority_Status">
Priority_Status
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machine.
Minimum api-version: 2019-03-01</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state, which only appears in the response.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroup: Specifies information about the proximity placement group that the virtual machine should be
assigned to.
Minimum api-version: 2018-04-01.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtension_Status">
[]VirtualMachineExtension_Status
</a>
</em>
</td>
<td>
<p>Resources: The virtual machine child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile_Status">
SecurityProfile_Status
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security related profile settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageProfile_Status">
StorageProfile_Status
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineScaleSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>VirtualMachineScaleSet: Specifies information about the virtual machine scale set that the virtual machine should be
assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to
maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM
cannot be added to a virtual machine scale set.
This property cannot exist along with a non-null properties.availabilitySet reference.
Minimum api‐version: 2019‐03‐01</p>
</td>
</tr>
<tr>
<td>
<code>vmId</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmId: Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS
and can be read using platform BIOS commands.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachine_StatusARM">VirtualMachine_StatusARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity_StatusARM">
VirtualMachineIdentity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the virtual machine, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan_StatusARM">
Plan_StatusARM
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineProperties_StatusARM">
VirtualMachineProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineExtension_StatusARM">
[]VirtualMachineExtension_StatusARM
</a>
</em>
</td>
<td>
<p>Resources: The virtual machine child extension resources.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachinesSpecAPIVersion">VirtualMachinesSpecAPIVersion
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesEvictionPolicy">VirtualMachinesSpecPropertiesEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesPriority">VirtualMachinesSpecPropertiesPriority
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Low&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachine">VirtualMachine</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilities">
AdditionalCapabilities
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>availabilitySet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfile">
BillingProfile
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VM or VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfile">
DiagnosticsProfile
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesEvictionPolicy">
VirtualMachinesSpecPropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
<tr>
<td>
<code>hardwareProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfile">
HardwareProfile
</a>
</em>
</td>
<td>
<p>HardwareProfile: Specifies the hardware settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentity">
VirtualMachineIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile">
VirtualMachines_Spec_Properties_NetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies the network interfaces of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSProfile">
OSProfile
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings for the virtual machine. Some of the settings cannot be changed once
VM is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.Plan">
Plan
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the scale set logical fault domain into which the Virtual Machine will be created. By
default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across
available fault domains.
<li>This is applicable only if the &lsquo;virtualMachineScaleSet&rsquo; property of this Virtual Machine is set.<li>The Virtual
Machine Scale Set that is referenced, must have &lsquo;platformFaultDomainCount&rsquo; &amp;gt; 1.<li>This property cannot be updated
once the Virtual Machine is created.<li>Fault domain assignment can be viewed in the Virtual Machine Instance View.
Minimum api‐version: 2020‐12‐01</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesPriority">
VirtualMachinesSpecPropertiesPriority
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machine.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfile">
SecurityProfile
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security profile settings for the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageProfile">
StorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineScaleSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_SpecARM">VirtualMachines_SpecARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachineIdentityARM">
VirtualMachineIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Identity for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>plan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PlanARM">
PlanARM
</a>
</em>
</td>
<td>
<p>Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
programmatically, Get Started -&gt;. Enter any required information and then click Save.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">
VirtualMachines_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes the properties of a Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The virtual machine zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_SpecARM">VirtualMachines_SpecARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalCapabilities</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalCapabilitiesARM">
AdditionalCapabilitiesARM
</a>
</em>
</td>
<td>
<p>AdditionalCapabilities: Enables or disables a capability on the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>availabilitySet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>billingProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.BillingProfileARM">
BillingProfileARM
</a>
</em>
</td>
<td>
<p>BillingProfile: Specifies the billing related details of a Azure Spot VM or VMSS.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>diagnosticsProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.DiagnosticsProfileARM">
DiagnosticsProfileARM
</a>
</em>
</td>
<td>
<p>DiagnosticsProfile: Specifies the boot diagnostic settings state.
Minimum api-version: 2015-06-15.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesEvictionPolicy">
VirtualMachinesSpecPropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
For Azure Spot virtual machines, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is 2019-03-01.
For Azure Spot scale sets, both &lsquo;Deallocate&rsquo; and &lsquo;Delete&rsquo; are supported and the minimum api-version is
2017-10-30-preview.</p>
</td>
</tr>
<tr>
<td>
<code>extensionsTimeBudget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
(PT1H30M).
Minimum api-version: 2020-06-01</p>
</td>
</tr>
<tr>
<td>
<code>hardwareProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.HardwareProfileARM">
HardwareProfileARM
</a>
</em>
</td>
<td>
<p>HardwareProfile: Specifies the hardware settings for the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hostGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
Possible values for Windows Server operating system are:
Windows_Client
Windows_Server
Possible values for Linux Server operating system are:
RHEL_BYOS (for RHEL)
SLES_BYOS (for SUSE)
For more information, see <a href="https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing">Azure Hybrid Use Benefit for Windows
Server</a>
<a href="https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux">Azure Hybrid Use Benefit for Linux
Server</a>
Minimum api-version: 2015-06-15</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfileARM">
VirtualMachines_Spec_Properties_NetworkProfileARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Specifies the network interfaces of the virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>osProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.OSProfileARM">
OSProfileARM
</a>
</em>
</td>
<td>
<p>OsProfile: Specifies the operating system settings for the virtual machine. Some of the settings cannot be changed once
VM is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>platformFaultDomain</code><br/>
<em>
int
</em>
</td>
<td>
<p>PlatformFaultDomain: Specifies the scale set logical fault domain into which the Virtual Machine will be created. By
default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across
available fault domains.
<li>This is applicable only if the &lsquo;virtualMachineScaleSet&rsquo; property of this Virtual Machine is set.<li>The Virtual
Machine Scale Set that is referenced, must have &lsquo;platformFaultDomainCount&rsquo; &amp;gt; 1.<li>This property cannot be updated
once the Virtual Machine is created.<li>Fault domain assignment can be viewed in the Virtual Machine Instance View.
Minimum api‐version: 2020‐12‐01</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachinesSpecPropertiesPriority">
VirtualMachinesSpecPropertiesPriority
</a>
</em>
</td>
<td>
<p>Priority: Specifies the priority for the virtual machine.
Minimum api-version: 2019-03-01.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroup</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SecurityProfileARM">
SecurityProfileARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Specifies the Security profile settings for the virtual machine or virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.StorageProfileARM">
StorageProfileARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachineScaleSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile">VirtualMachines_Spec_Properties_NetworkProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec">VirtualMachines_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces">
[]VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: Specifies the list of resource Ids for the network interfaces associated with the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfileARM">VirtualMachines_Spec_Properties_NetworkProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_PropertiesARM">VirtualMachines_Spec_PropertiesARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM">
[]VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: Specifies the list of resource Ids for the network interfaces associated with the virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces">VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile">VirtualMachines_Spec_Properties_NetworkProfile</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>primary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM">VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfileARM">VirtualMachines_Spec_Properties_NetworkProfileARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.NetworkInterfaceReferencePropertiesARM">
NetworkInterfaceReferencePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes a network interface reference properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMConfiguration">WinRMConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration">WindowsConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>listeners</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListener">
[]WinRMListener
</a>
</em>
</td>
<td>
<p>Listeners: The list of Windows Remote Management listeners</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMConfigurationARM">WinRMConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfigurationARM">WindowsConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>listeners</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListenerARM">
[]WinRMListenerARM
</a>
</em>
</td>
<td>
<p>Listeners: The list of Windows Remote Management listeners</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMConfiguration_Status">WinRMConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_Status">WindowsConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>listeners</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListener_Status">
[]WinRMListener_Status
</a>
</em>
</td>
<td>
<p>Listeners: The list of Windows Remote Management listeners</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMConfiguration_StatusARM">WinRMConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">WindowsConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>listeners</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListener_StatusARM">
[]WinRMListener_StatusARM
</a>
</em>
</td>
<td>
<p>Listeners: The list of Windows Remote Management listeners</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListener">WinRMListener
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMConfiguration">WinRMConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMListener">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMListener</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListenerProtocol">
WinRMListenerProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Specifies the protocol of WinRM listener.
Possible values are:
http
https.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListenerARM">WinRMListenerARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMConfigurationARM">WinRMConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMListener">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WinRMListener</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListenerProtocol">
WinRMListenerProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Specifies the protocol of WinRM listener.
Possible values are:
http
https.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListenerProtocol">WinRMListenerProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMListener">WinRMListener</a>, <a href="#compute.azure.com/v1beta20201201.WinRMListenerARM">WinRMListenerARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Https&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListenerStatusProtocol">WinRMListenerStatusProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMListener_Status">WinRMListener_Status</a>, <a href="#compute.azure.com/v1beta20201201.WinRMListener_StatusARM">WinRMListener_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Https&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListener_Status">WinRMListener_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMConfiguration_Status">WinRMConfiguration_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListenerStatusProtocol">
WinRMListenerStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Specifies the protocol of WinRM listener.
Possible values are:
http
https</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WinRMListener_StatusARM">WinRMListener_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.WinRMConfiguration_StatusARM">WinRMConfiguration_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificateUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificateUrl: This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to
the Key Vault, see <a href="https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add">Add a key or secret to the key
vault</a>. In this case, your certificate needs to
be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8:
{
&ldquo;data&rdquo;:&rdquo;<Base64-encoded-certificate>&rdquo;,
&ldquo;dataType&rdquo;:&ldquo;pfx&rdquo;,
&ldquo;password&rdquo;:&rdquo;<pfx-file-password>&rdquo;
}</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMListenerStatusProtocol">
WinRMListenerStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Specifies the protocol of WinRM listener.
Possible values are:
http
https</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WindowsConfiguration">WindowsConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile">OSProfile</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile">VirtualMachineScaleSetOSProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WindowsConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WindowsConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalUnattendContent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent">
[]AdditionalUnattendContent
</a>
</em>
</td>
<td>
<p>AdditionalUnattendContent: Specifies additional base-64 encoded XML formatted information that can be included in the
Unattend.xml file, which is used by Windows Setup.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpdates</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpdates: Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is
true.
For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettings">
PatchSettings
</a>
</em>
</td>
<td>
<p>PatchSettings: Specifies settings related to VM Guest Patching on Windows.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>timeZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeZone: Specifies the time zone of the virtual machine. e.g. &ldquo;Pacific Standard Time&rdquo;.
Possible values can be
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id">TimeZoneInfo.Id</a> value from
time zones returned by
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones">TimeZoneInfo.GetSystemTimeZones</a>.</p>
</td>
</tr>
<tr>
<td>
<code>winRM</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMConfiguration">
WinRMConfiguration
</a>
</em>
</td>
<td>
<p>WinRM: Describes Windows Remote Management configuration of the VM</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WindowsConfigurationARM">WindowsConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfileARM">OSProfileARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfileARM">VirtualMachineScaleSetOSProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WindowsConfiguration">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Compute.json#/definitions/WindowsConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalUnattendContent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContentARM">
[]AdditionalUnattendContentARM
</a>
</em>
</td>
<td>
<p>AdditionalUnattendContent: Specifies additional base-64 encoded XML formatted information that can be included in the
Unattend.xml file, which is used by Windows Setup.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpdates</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpdates: Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is
true.
For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettingsARM">
PatchSettingsARM
</a>
</em>
</td>
<td>
<p>PatchSettings: Specifies settings related to VM Guest Patching on Windows.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>timeZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeZone: Specifies the time zone of the virtual machine. e.g. &ldquo;Pacific Standard Time&rdquo;.
Possible values can be
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id">TimeZoneInfo.Id</a> value from
time zones returned by
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones">TimeZoneInfo.GetSystemTimeZones</a>.</p>
</td>
</tr>
<tr>
<td>
<code>winRM</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMConfigurationARM">
WinRMConfigurationARM
</a>
</em>
</td>
<td>
<p>WinRM: Describes Windows Remote Management configuration of the VM</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WindowsConfiguration_Status">WindowsConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_Status">OSProfile_Status</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_Status">VirtualMachineScaleSetOSProfile_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalUnattendContent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_Status">
[]AdditionalUnattendContent_Status
</a>
</em>
</td>
<td>
<p>AdditionalUnattendContent: Specifies additional base-64 encoded XML formatted information that can be included in the
Unattend.xml file, which is used by Windows Setup.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpdates</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpdates: Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is
true.
For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettings_Status">
PatchSettings_Status
</a>
</em>
</td>
<td>
<p>PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching on Windows.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>timeZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeZone: Specifies the time zone of the virtual machine. e.g. &ldquo;Pacific Standard Time&rdquo;.
Possible values can be
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id">TimeZoneInfo.Id</a> value from
time zones returned by
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones">TimeZoneInfo.GetSystemTimeZones</a>.</p>
</td>
</tr>
<tr>
<td>
<code>winRM</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMConfiguration_Status">
WinRMConfiguration_Status
</a>
</em>
</td>
<td>
<p>WinRM: Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20201201.WindowsConfiguration_StatusARM">WindowsConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20201201.OSProfile_StatusARM">OSProfile_StatusARM</a>, <a href="#compute.azure.com/v1beta20201201.VirtualMachineScaleSetOSProfile_StatusARM">VirtualMachineScaleSetOSProfile_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalUnattendContent</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.AdditionalUnattendContent_StatusARM">
[]AdditionalUnattendContent_StatusARM
</a>
</em>
</td>
<td>
<p>AdditionalUnattendContent: Specifies additional base-64 encoded XML formatted information that can be included in the
Unattend.xml file, which is used by Windows Setup.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticUpdates</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticUpdates: Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is
true.
For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.</p>
</td>
</tr>
<tr>
<td>
<code>patchSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.PatchSettings_StatusARM">
PatchSettings_StatusARM
</a>
</em>
</td>
<td>
<p>PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching on Windows.</p>
</td>
</tr>
<tr>
<td>
<code>provisionVMAgent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned on the virtual machine.
When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that
VM Agent is installed on the VM so that extensions can be added to the VM later.</p>
</td>
</tr>
<tr>
<td>
<code>timeZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeZone: Specifies the time zone of the virtual machine. e.g. &ldquo;Pacific Standard Time&rdquo;.
Possible values can be
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id">TimeZoneInfo.Id</a> value from
time zones returned by
<a href="https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones">TimeZoneInfo.GetSystemTimeZones</a>.</p>
</td>
</tr>
<tr>
<td>
<code>winRM</code><br/>
<em>
<a href="#compute.azure.com/v1beta20201201.WinRMConfiguration_StatusARM">
WinRMConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>WinRM: Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
