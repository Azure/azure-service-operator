---
title: network.azure.com/v1beta20201101
---
<h2 id="network.azure.com/v1beta20201101">network.azure.com/v1beta20201101</h2>
<div>
<p>Package v1beta20201101 contains API Schema definitions for the network v1beta20201101 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="network.azure.com/v1beta20201101.AddressSpace">AddressSpace
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">VirtualNetworksVirtualNetworkPeerings_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/AddressSpace">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/AddressSpace</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.AddressSpaceARM">AddressSpaceARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">VirtualNetworkPeeringPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/AddressSpace">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/AddressSpace</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.AddressSpace_Status">AddressSpace_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>, <a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>)
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
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.AddressSpace_StatusARM">AddressSpace_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<code>backendAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddress_StatusARM">
[]ApplicationGatewayBackendAddress_StatusARM
</a>
</em>
</td>
<td>
<p>BackendAddresses: Backend addresses.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the backend address pool resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>backendAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddress_Status">
[]ApplicationGatewayBackendAddress_Status
</a>
</em>
</td>
<td>
<p>BackendAddresses: Backend addresses.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: Name of the backend address pool that is unique within an Application Gateway.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the backend address pool resource.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: Name of the backend address pool that is unique within an Application Gateway.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">
ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the application gateway backend address pool.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayBackendAddress_Status">ApplicationGatewayBackendAddress_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: Fully qualified domain name (FQDN).</p>
</td>
</tr>
<tr>
<td>
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: IP address.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayBackendAddress_StatusARM">ApplicationGatewayBackendAddress_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: Fully qualified domain name (FQDN).</p>
</td>
</tr>
<tr>
<td>
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: IP address.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM">ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_StatusARM">ApplicationGatewayIPConfiguration_StatusARM</a>)
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
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the application gateway IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Subnet: Reference to the subnet resource. A subnet from where application gateway gets its private address.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_Status">ApplicationGatewayIPConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: Name of the IP configuration that is unique within an Application Gateway.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the application gateway IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Subnet: Reference to the subnet resource. A subnet from where application gateway gets its private address.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_StatusARM">ApplicationGatewayIPConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: Name of the IP configuration that is unique within an Application Gateway.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM">
ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the application gateway IP configuration.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded">ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM">ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded">BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM">BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM">BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BgpSettings">BgpSettings
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>asn</code><br/>
<em>
uint32
</em>
</td>
<td>
<p>Asn: The BGP speaker&rsquo;s ASN.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress">
[]IPConfigurationBgpPeeringAddress
</a>
</em>
</td>
<td>
<p>BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>peerWeight</code><br/>
<em>
int
</em>
</td>
<td>
<p>PeerWeight: The weight added to routes learned from this BGP speaker.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BgpSettingsARM">BgpSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/BgpSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>asn</code><br/>
<em>
uint32
</em>
</td>
<td>
<p>Asn: The BGP speaker&rsquo;s ASN.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddressARM">
[]IPConfigurationBgpPeeringAddressARM
</a>
</em>
</td>
<td>
<p>BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>peerWeight</code><br/>
<em>
int
</em>
</td>
<td>
<p>PeerWeight: The weight added to routes learned from this BGP speaker.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BgpSettings_Status">BgpSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<code>asn</code><br/>
<em>
uint32
</em>
</td>
<td>
<p>Asn: The BGP speaker&rsquo;s ASN.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress_Status">
[]IPConfigurationBgpPeeringAddress_Status
</a>
</em>
</td>
<td>
<p>BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>peerWeight</code><br/>
<em>
int
</em>
</td>
<td>
<p>PeerWeight: The weight added to routes learned from this BGP speaker.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.BgpSettings_StatusARM">BgpSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>)
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
<code>asn</code><br/>
<em>
uint32
</em>
</td>
<td>
<p>Asn: The BGP speaker&rsquo;s ASN.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.</p>
</td>
</tr>
<tr>
<td>
<code>bgpPeeringAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress_StatusARM">
[]IPConfigurationBgpPeeringAddress_StatusARM
</a>
</em>
</td>
<td>
<p>BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>peerWeight</code><br/>
<em>
int
</em>
</td>
<td>
<p>PeerWeight: The weight added to routes learned from this BGP speaker.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettings">DdosSettings
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ddosCustomPolicy</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>DdosCustomPolicy: The DDoS custom policy associated with the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectedIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProtectedIP: Enables DDoS protection on the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectionCoverage</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettingsProtectionCoverage">
DdosSettingsProtectionCoverage
</a>
</em>
</td>
<td>
<p>ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
ability to be customized.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettingsARM">DdosSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DdosSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ddosCustomPolicy</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>DdosCustomPolicy: The DDoS custom policy associated with the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectedIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProtectedIP: Enables DDoS protection on the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectionCoverage</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettingsProtectionCoverage">
DdosSettingsProtectionCoverage
</a>
</em>
</td>
<td>
<p>ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
ability to be customized.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettingsProtectionCoverage">DdosSettingsProtectionCoverage
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.DdosSettings">DdosSettings</a>, <a href="#network.azure.com/v1beta20201101.DdosSettingsARM">DdosSettingsARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettingsStatusProtectionCoverage">DdosSettingsStatusProtectionCoverage
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.DdosSettings_Status">DdosSettings_Status</a>, <a href="#network.azure.com/v1beta20201101.DdosSettings_StatusARM">DdosSettings_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettings_Status">DdosSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<code>ddosCustomPolicy</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DdosCustomPolicy: The DDoS custom policy associated with the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectedIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProtectedIP: Enables DDoS protection on the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectionCoverage</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettingsStatusProtectionCoverage">
DdosSettingsStatusProtectionCoverage
</a>
</em>
</td>
<td>
<p>ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
ability to be customized.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DdosSettings_StatusARM">DdosSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>)
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
<code>ddosCustomPolicy</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DdosCustomPolicy: The DDoS custom policy associated with the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectedIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ProtectedIP: Enables DDoS protection on the public IP.</p>
</td>
</tr>
<tr>
<td>
<code>protectionCoverage</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettingsStatusProtectionCoverage">
DdosSettingsStatusProtectionCoverage
</a>
</em>
</td>
<td>
<p>ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
ability to be customized.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Delegation_Status">Delegation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>actions</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Actions: The actions permitted to the service upon delegation.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service delegation resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Delegation_StatusARM">Delegation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormat_StatusARM">
ServiceDelegationPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DhGroup_Status">DhGroup_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;DHGroup1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup2048&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP384&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DhcpOptions">DhcpOptions
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions</a></p>
</div>
<table>
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
<p>DnsServers: The list of DNS servers IP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DhcpOptionsARM">DhcpOptionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions</a></p>
</div>
<table>
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
<p>DnsServers: The list of DNS servers IP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DhcpOptions_Status">DhcpOptions_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>)
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
<p>DnsServers: The list of DNS servers IP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.DhcpOptions_StatusARM">DhcpOptions_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM</a>)
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
<p>DnsServers: The list of DNS servers IP addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec">NetworkInterfaces_Spec</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationType">
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
<h3 id="network.azure.com/v1beta20201101.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_SpecARM">LoadBalancers_SpecARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_SpecARM">NetworkInterfaces_SpecARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddresses_SpecARM">PublicIPAddresses_SpecARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_SpecARM">VirtualNetworks_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationType">
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
<h3 id="network.azure.com/v1beta20201101.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ExtendedLocation">ExtendedLocation</a>, <a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">ExtendedLocationARM</a>)
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
<h3 id="network.azure.com/v1beta20201101.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
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
<h3 id="network.azure.com/v1beta20201101.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded">PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded">PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded">PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded">PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded">PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>)
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationType_Status">
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
<h3 id="network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_StatusARM">LoadBalancer_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM">PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM">PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM">PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM">PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_StatusARM">VirtualNetworkGateway_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_StatusARM">VirtualNetwork_StatusARM</a>)
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationType_Status">
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
<h3 id="network.azure.com/v1beta20201101.FlowLog_Status_SubResourceEmbedded">FlowLog_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FlowLog_Status_SubResourceEmbeddedARM">FlowLog_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatARM">FrontendIPConfigurationPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM">LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/FrontendIPConfigurationPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/FrontendIPConfigurationPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion">
FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The Private IP allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the Public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The reference to the Public IP Prefix resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion">FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatARM">FrontendIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">LoadBalancers_Spec_Properties_FrontendIPConfigurations</a>)
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
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod">FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatARM">FrontendIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">LoadBalancers_Spec_Properties_FrontendIPConfigurations</a>)
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
<tbody><tr><td><p>&#34;Dynamic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Static&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM</a>)
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
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>InboundNatPools: An array of references to inbound pools that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>InboundNatRules: An array of references to inbound rules that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: An array of references to load balancing rules that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>OutboundRules: An array of references to outbound rules that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The Private IP allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the frontend IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM">
PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the Public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The reference to the Public IP Prefix resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_LoadBalancer_SubResourceEmbeddedARM">
Subnet_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>InboundNatPools: An array of references to inbound pools that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>InboundNatRules: An array of references to inbound rules that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: An array of references to load balancing rules that use this frontend IP.</p>
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
<p>Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>OutboundRules: An array of references to outbound rules that use this frontend IP.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The Private IP allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the frontend IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded">
PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the Public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The reference to the Public IP Prefix resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_LoadBalancer_SubResourceEmbedded">
Subnet_Status_LoadBalancer_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
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
<p>Type: Type of the resource.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">
FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the load balancer probe.</p>
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
<p>Type: Type of the resource.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPAllocationMethod_Status">IPAllocationMethod_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded">IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM">VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status">VirtualNetworkGatewayIPConfiguration_Status</a>)
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
<tbody><tr><td><p>&#34;Dynamic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Static&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress">IPConfigurationBgpPeeringAddress
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.BgpSettings">BgpSettings</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>customBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>ipconfigurationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpconfigurationId: The ID of IP configuration which belongs to gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddressARM">IPConfigurationBgpPeeringAddressARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.BgpSettingsARM">BgpSettingsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IPConfigurationBgpPeeringAddress</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>customBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>ipconfigurationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpconfigurationId: The ID of IP configuration which belongs to gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress_Status">IPConfigurationBgpPeeringAddress_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.BgpSettings_Status">BgpSettings_Status</a>)
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
<code>customBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>defaultBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DefaultBgpIpAddresses: The list of default BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>ipconfigurationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpconfigurationId: The ID of IP configuration which belongs to gateway.</p>
</td>
</tr>
<tr>
<td>
<code>tunnelIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>TunnelIpAddresses: The list of tunnel public IP addresses which belong to IP configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationBgpPeeringAddress_StatusARM">IPConfigurationBgpPeeringAddress_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.BgpSettings_StatusARM">BgpSettings_StatusARM</a>)
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
<code>customBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>defaultBgpIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DefaultBgpIpAddresses: The list of default BGP peering addresses which belong to IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>ipconfigurationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpconfigurationId: The ID of IP configuration which belongs to gateway.</p>
</td>
</tr>
<tr>
<td>
<code>tunnelIpAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>TunnelIpAddresses: The list of tunnel public IP addresses which belong to IP configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>)
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
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration profile resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration profile resource.</p>
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
<p>Type: Sub Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the IP configuration profile.</p>
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
<p>Type: Sub Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM</a>)
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
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM">
Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>)
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
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded">IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_PublicIPAddress_SubResourceEmbedded">
Subnet_Status_PublicIPAddress_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM">
IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the IP configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded">
PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the IP configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IPVersion_Status">IPVersion_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<h3 id="network.azure.com/v1beta20201101.IkeEncryption_Status">IkeEncryption_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;AES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IkeIntegrity_Status">IkeIntegrity_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MD5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA384&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatARM">InboundNatPoolPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPoolsARM">LoadBalancers_Spec_Properties_InboundNatPoolsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/InboundNatPoolPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/InboundNatPoolPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65534.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatProtocol">
InboundNatPoolPropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the inbound NAT pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatProtocol">InboundNatPoolPropertiesFormatProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatARM">InboundNatPoolPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools">LoadBalancers_Spec_Properties_InboundNatPools</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormat_StatusARM">InboundNatPoolPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.InboundNatPool_StatusARM">InboundNatPool_StatusARM</a>)
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
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65534.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.TransportProtocol_Status">
TransportProtocol_Status
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the inbound NAT pool.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the inbound NAT pool resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatPool_Status">InboundNatPool_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65534.</p>
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
<p>Id: Resource ID.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
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
<p>Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.TransportProtocol_Status">
TransportProtocol_Status
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the inbound NAT pool.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the inbound NAT pool resource.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatPool_StatusARM">InboundNatPool_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormat_StatusARM">
InboundNatPoolPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer inbound nat pool.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbedded">InboundNatRule_Status_LoadBalancer_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM">InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbedded">InboundNatRule_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM">InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpTag">IpTag
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag</a></p>
</div>
<table>
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
<p>IpTagType: The IP tag type. Example: FirstPartyUsage.</p>
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
<p>Tag: The value of the IP tag associated with the public IP. Example: SQL.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpTagARM">IpTagARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpTag</a></p>
</div>
<table>
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
<p>IpTagType: The IP tag type. Example: FirstPartyUsage.</p>
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
<p>Tag: The value of the IP tag associated with the public IP. Example: SQL.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpTag_Status">IpTag_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<p>IpTagType: The IP tag type. Example: FirstPartyUsage.</p>
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
<p>Tag: The value of the IP tag associated with the public IP. Example: SQL.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpTag_StatusARM">IpTag_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>)
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
<p>IpTagType: The IP tag type. Example: FirstPartyUsage.</p>
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
<p>Tag: The value of the IP tag associated with the public IP. Example: SQL.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecEncryption_Status">IpsecEncryption_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;AES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecIntegrity_Status">IpsecIntegrity_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MD5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA256&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dhGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyDhGroup">
IpsecPolicyDhGroup
</a>
</em>
</td>
<td>
<p>DhGroup: The DH Group used in IKE Phase 1 for initial SA.</p>
</td>
</tr>
<tr>
<td>
<code>ikeEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIkeEncryption">
IpsecPolicyIkeEncryption
</a>
</em>
</td>
<td>
<p>IkeEncryption: The IKE encryption algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ikeIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIkeIntegrity">
IpsecPolicyIkeIntegrity
</a>
</em>
</td>
<td>
<p>IkeIntegrity: The IKE integrity algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIpsecEncryption">
IpsecPolicyIpsecEncryption
</a>
</em>
</td>
<td>
<p>IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIpsecIntegrity">
IpsecPolicyIpsecIntegrity
</a>
</em>
</td>
<td>
<p>IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>pfsGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyPfsGroup">
IpsecPolicyPfsGroup
</a>
</em>
</td>
<td>
<p>PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.</p>
</td>
</tr>
<tr>
<td>
<code>saDataSizeKilobytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
to site VPN tunnel.</p>
</td>
</tr>
<tr>
<td>
<code>saLifeTimeSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
to site VPN tunnel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/IpsecPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dhGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyDhGroup">
IpsecPolicyDhGroup
</a>
</em>
</td>
<td>
<p>DhGroup: The DH Group used in IKE Phase 1 for initial SA.</p>
</td>
</tr>
<tr>
<td>
<code>ikeEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIkeEncryption">
IpsecPolicyIkeEncryption
</a>
</em>
</td>
<td>
<p>IkeEncryption: The IKE encryption algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ikeIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIkeIntegrity">
IpsecPolicyIkeIntegrity
</a>
</em>
</td>
<td>
<p>IkeIntegrity: The IKE integrity algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIpsecEncryption">
IpsecPolicyIpsecEncryption
</a>
</em>
</td>
<td>
<p>IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyIpsecIntegrity">
IpsecPolicyIpsecIntegrity
</a>
</em>
</td>
<td>
<p>IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>pfsGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyPfsGroup">
IpsecPolicyPfsGroup
</a>
</em>
</td>
<td>
<p>PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.</p>
</td>
</tr>
<tr>
<td>
<code>saDataSizeKilobytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
to site VPN tunnel.</p>
</td>
</tr>
<tr>
<td>
<code>saLifeTimeSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
to site VPN tunnel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyDhGroup">IpsecPolicyDhGroup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;DHGroup1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup2048&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DHGroup24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP384&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyIkeEncryption">IpsecPolicyIkeEncryption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;AES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyIkeIntegrity">IpsecPolicyIkeIntegrity
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MD5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA384&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyIpsecEncryption">IpsecPolicyIpsecEncryption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;AES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DES3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyIpsecIntegrity">IpsecPolicyIpsecIntegrity
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;GCMAES128&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES192&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GCMAES256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MD5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SHA256&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicyPfsGroup">IpsecPolicyPfsGroup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy">IpsecPolicy</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">IpsecPolicyARM</a>)
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
<tbody><tr><td><p>&#34;ECP256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP384&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS2048&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFSMM&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>)
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
<code>dhGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhGroup_Status">
DhGroup_Status
</a>
</em>
</td>
<td>
<p>DhGroup: The DH Group used in IKE Phase 1 for initial SA.</p>
</td>
</tr>
<tr>
<td>
<code>ikeEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IkeEncryption_Status">
IkeEncryption_Status
</a>
</em>
</td>
<td>
<p>IkeEncryption: The IKE encryption algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ikeIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IkeIntegrity_Status">
IkeIntegrity_Status
</a>
</em>
</td>
<td>
<p>IkeIntegrity: The IKE integrity algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecEncryption_Status">
IpsecEncryption_Status
</a>
</em>
</td>
<td>
<p>IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecIntegrity_Status">
IpsecIntegrity_Status
</a>
</em>
</td>
<td>
<p>IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>pfsGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PfsGroup_Status">
PfsGroup_Status
</a>
</em>
</td>
<td>
<p>PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.</p>
</td>
</tr>
<tr>
<td>
<code>saDataSizeKilobytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
to site VPN tunnel.</p>
</td>
</tr>
<tr>
<td>
<code>saLifeTimeSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
to site VPN tunnel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<code>dhGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhGroup_Status">
DhGroup_Status
</a>
</em>
</td>
<td>
<p>DhGroup: The DH Group used in IKE Phase 1 for initial SA.</p>
</td>
</tr>
<tr>
<td>
<code>ikeEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IkeEncryption_Status">
IkeEncryption_Status
</a>
</em>
</td>
<td>
<p>IkeEncryption: The IKE encryption algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ikeIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IkeIntegrity_Status">
IkeIntegrity_Status
</a>
</em>
</td>
<td>
<p>IkeIntegrity: The IKE integrity algorithm (IKE phase 2).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecEncryption</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecEncryption_Status">
IpsecEncryption_Status
</a>
</em>
</td>
<td>
<p>IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>ipsecIntegrity</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecIntegrity_Status">
IpsecIntegrity_Status
</a>
</em>
</td>
<td>
<p>IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).</p>
</td>
</tr>
<tr>
<td>
<code>pfsGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PfsGroup_Status">
PfsGroup_Status
</a>
</em>
</td>
<td>
<p>PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.</p>
</td>
</tr>
<tr>
<td>
<code>saDataSizeKilobytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
to site VPN tunnel.</p>
</td>
</tr>
<tr>
<td>
<code>saLifeTimeSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
to site VPN tunnel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancer">LoadBalancer
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/loadBalancers</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">
LoadBalancers_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
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
<code>backendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools">
[]LoadBalancers_Spec_Properties_BackendAddressPools
</a>
</em>
</td>
<td>
<p>BackendAddressPools: Collection of backend address pools used by a load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">
[]LoadBalancers_Spec_Properties_FrontendIPConfigurations
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools">
[]LoadBalancers_Spec_Properties_InboundNatPools
</a>
</em>
</td>
<td>
<p>InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">
[]LoadBalancers_Spec_Properties_LoadBalancingRules
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.</p>
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
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules">
[]LoadBalancers_Spec_Properties_OutboundRules
</a>
</em>
</td>
<td>
<p>OutboundRules: The outbound rules.</p>
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
<code>probes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_Probes">
[]LoadBalancers_Spec_Properties_Probes
</a>
</em>
</td>
<td>
<p>Probes: Collection of probe objects used in the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSku">
LoadBalancerSku
</a>
</em>
</td>
<td>
<p>Sku: The load balancer SKU.</p>
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">
LoadBalancer_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerBackendAddressPropertiesFormatARM">LoadBalancerBackendAddressPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM">LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerBackendAddressPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerBackendAddressPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: IP Address belonging to the referenced virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerFrontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>LoadBalancerFrontendIPConfiguration: Reference to the frontend ip address configuration defined in regional loadbalancer.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>Subnet: Reference to an existing subnet.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>VirtualNetwork: Reference to an existing virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_StatusARM">LoadBalancer_StatusARM</a>)
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
<code>backendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM">
[]BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>BackendAddressPools: Collection of backend address pools used by a load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM">
[]FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPool_StatusARM">
[]InboundNatPool_StatusARM
</a>
</em>
</td>
<td>
<p>InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM">
[]InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>InboundNatRules: Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load
balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine
scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to
reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRule_StatusARM">
[]LoadBalancingRule_StatusARM
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.</p>
</td>
</tr>
<tr>
<td>
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRule_StatusARM">
[]OutboundRule_StatusARM
</a>
</em>
</td>
<td>
<p>OutboundRules: The outbound rules.</p>
</td>
</tr>
<tr>
<td>
<code>probes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Probe_StatusARM">
[]Probe_StatusARM
</a>
</em>
</td>
<td>
<p>Probes: Collection of probe objects used in the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the load balancer resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the load balancer resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSku">LoadBalancerSku
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuName">
LoadBalancerSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of a load balancer SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuTier">
LoadBalancerSkuTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a load balancer SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSkuARM">LoadBalancerSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_SpecARM">LoadBalancers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancerSku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuName">
LoadBalancerSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of a load balancer SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuTier">
LoadBalancerSkuTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a load balancer SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSkuName">LoadBalancerSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerSku">LoadBalancerSku</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerSkuARM">LoadBalancerSkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSkuStatusName">LoadBalancerSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerSku_Status">LoadBalancerSku_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerSku_StatusARM">LoadBalancerSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSkuStatusTier">LoadBalancerSkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerSku_Status">LoadBalancerSku_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerSku_StatusARM">LoadBalancerSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Global&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regional&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSkuTier">LoadBalancerSkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerSku">LoadBalancerSku</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerSkuARM">LoadBalancerSkuARM</a>)
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
<tbody><tr><td><p>&#34;Global&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regional&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSku_Status">LoadBalancerSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuStatusName">
LoadBalancerSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of a load balancer SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuStatusTier">
LoadBalancerSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a load balancer SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancerSku_StatusARM">LoadBalancerSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_StatusARM">LoadBalancer_StatusARM</a>)
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
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuStatusName">
LoadBalancerSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of a load balancer SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuStatusTier">
LoadBalancerSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a load balancer SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer">LoadBalancer</a>)
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
<code>backendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded">
[]BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>BackendAddressPools: Collection of backend address pools used by a load balancer.</p>
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">
[]FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPool_Status">
[]InboundNatPool_Status
</a>
</em>
</td>
<td>
<p>InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbedded">
[]InboundNatRule_Status_LoadBalancer_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>InboundNatRules: Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load
balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine
scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to
reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRule_Status">
[]LoadBalancingRule_Status
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRule_Status">
[]OutboundRule_Status
</a>
</em>
</td>
<td>
<p>OutboundRules: The outbound rules.</p>
</td>
</tr>
<tr>
<td>
<code>probes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Probe_Status">
[]Probe_Status
</a>
</em>
</td>
<td>
<p>Probes: Collection of probe objects used in the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the load balancer resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the load balancer resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSku_Status">
LoadBalancerSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The load balancer SKU.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancer_StatusARM">LoadBalancer_StatusARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">
LoadBalancerPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSku_StatusARM">
LoadBalancerSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The load balancer SKU.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancersSpecAPIVersion">LoadBalancersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer">LoadBalancer</a>)
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
<code>backendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools">
[]LoadBalancers_Spec_Properties_BackendAddressPools
</a>
</em>
</td>
<td>
<p>BackendAddressPools: Collection of backend address pools used by a load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">
[]LoadBalancers_Spec_Properties_FrontendIPConfigurations
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools">
[]LoadBalancers_Spec_Properties_InboundNatPools
</a>
</em>
</td>
<td>
<p>InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">
[]LoadBalancers_Spec_Properties_LoadBalancingRules
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.</p>
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
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules">
[]LoadBalancers_Spec_Properties_OutboundRules
</a>
</em>
</td>
<td>
<p>OutboundRules: The outbound rules.</p>
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
<code>probes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_Probes">
[]LoadBalancers_Spec_Properties_Probes
</a>
</em>
</td>
<td>
<p>Probes: Collection of probe objects used in the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSku">
LoadBalancerSku
</a>
</em>
</td>
<td>
<p>Sku: The load balancer SKU.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_SpecARM">LoadBalancers_SpecARM
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">
LoadBalancers_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerSkuARM">
LoadBalancerSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The load balancer SKU.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_SpecARM">LoadBalancers_SpecARM</a>)
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
<code>backendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPoolsARM">
[]LoadBalancers_Spec_Properties_BackendAddressPoolsARM
</a>
</em>
</td>
<td>
<p>BackendAddressPools: Collection of backend address pools used by a load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM">
[]LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>inboundNatPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPoolsARM">
[]LoadBalancers_Spec_Properties_InboundNatPoolsARM
</a>
</em>
</td>
<td>
<p>InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRulesARM">
[]LoadBalancers_Spec_Properties_LoadBalancingRulesARM
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.</p>
</td>
</tr>
<tr>
<td>
<code>outboundRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRulesARM">
[]LoadBalancers_Spec_Properties_OutboundRulesARM
</a>
</em>
</td>
<td>
<p>OutboundRules: The outbound rules.</p>
</td>
</tr>
<tr>
<td>
<code>probes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_ProbesARM">
[]LoadBalancers_Spec_Properties_ProbesARM
</a>
</em>
</td>
<td>
<p>Probes: Collection of probe objects used in the load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools">LoadBalancers_Spec_Properties_BackendAddressPools
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<code>loadBalancerBackendAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses">
[]LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddresses: An array of backend addresses.</p>
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
<p>Location: The location of the backend address pool.</p>
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
<p>Name: The name of the resource that is unique within the set of backend address pools used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPoolsARM">LoadBalancers_Spec_Properties_BackendAddressPoolsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of backend address pools used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM">
LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer backend address pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM">LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPoolsARM">LoadBalancers_Spec_Properties_BackendAddressPoolsARM</a>)
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
<code>loadBalancerBackendAddresses</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM">
[]LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddresses: An array of backend addresses.</p>
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
<p>Location: The location of the backend address pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses">LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools">LoadBalancers_Spec_Properties_BackendAddressPools</a>)
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: IP Address belonging to the referenced virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerFrontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>LoadBalancerFrontendIPConfiguration: Reference to the frontend ip address configuration defined in regional loadbalancer.</p>
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
<p>Name: Name of the backend address.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>Subnet: Reference to an existing subnet.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>VirtualNetwork: Reference to an existing virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM">LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM">LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM</a>)
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
<p>Name: Name of the backend address.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancerBackendAddressPropertiesFormatARM">
LoadBalancerBackendAddressPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer backend address pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">LoadBalancers_Spec_Properties_FrontendIPConfigurations
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<p>Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: The private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion">
FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The Private IP allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the Public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The reference to the Public IP Prefix resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM">LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatARM">
FrontendIPConfigurationPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the load balancer probe.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools">LoadBalancers_Spec_Properties_InboundNatPools
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65535.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPortRangeStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
NICs associated with a load balancer. Acceptable values range between 1 and 65534.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
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
<p>Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatProtocol">
InboundNatPoolPropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the inbound NAT pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPoolsARM">LoadBalancers_Spec_Properties_InboundNatPoolsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatARM">
InboundNatPoolPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer inbound nat pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">LoadBalancers_Spec_Properties_LoadBalancingRules
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
that value 0 enables &ldquo;Any Port&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>disableOutboundSnat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
frontend of the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
Acceptable values are between 0 and 65534. Note that value 0 enables &ldquo;Any Port&rdquo;.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>loadDistribution</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatLoadDistribution">
LoadBalancingRulePropertiesFormatLoadDistribution
</a>
</em>
</td>
<td>
<p>LoadDistribution: The load distribution policy for this rule.</p>
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
<p>Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>probe</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>Probe: The reference to the load balancer probe used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatProtocol">
LoadBalancingRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the load balancing rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRulesARM">LoadBalancers_Spec_Properties_LoadBalancingRulesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatARM">
LoadBalancingRulePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer load balancing rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules">LoadBalancers_Spec_Properties_OutboundRules
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The number of outbound ports to be used for NAT.</p>
</td>
</tr>
<tr>
<td>
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: The Frontend IP addresses of the load balancer.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection.</p>
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
<p>Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatProtocol">
OutboundRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol for the outbound rule in load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRulesARM">LoadBalancers_Spec_Properties_OutboundRulesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatARM">
OutboundRulePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer outbound rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_Probes">LoadBalancers_Spec_Properties_Probes
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec">LoadBalancers_Spec</a>)
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
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
the instance out of rotation. The default value is 15, the minimum value is 5.</p>
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
<p>Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>numberOfProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
in Azure.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatProtocol">
ProbePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol of the end point. If &lsquo;Tcp&rsquo; is specified, a received ACK is required for the probe to be
successful. If &lsquo;Http&rsquo; or &lsquo;Https&rsquo; is specified, a 200 OK response from the specifies URI is required for the probe to be
successful.</p>
</td>
</tr>
<tr>
<td>
<code>requestPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
Otherwise, it is not allowed. There is no default value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_ProbesARM">LoadBalancers_Spec_Properties_ProbesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_PropertiesARM">LoadBalancers_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatARM">
ProbePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer probe.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatARM">LoadBalancingRulePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRulesARM">LoadBalancers_Spec_Properties_LoadBalancingRulesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancingRulePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/LoadBalancingRulePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
that value 0 enables &ldquo;Any Port&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>disableOutboundSnat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
frontend of the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
Acceptable values are between 0 and 65534. Note that value 0 enables &ldquo;Any Port&rdquo;.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>loadDistribution</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatLoadDistribution">
LoadBalancingRulePropertiesFormatLoadDistribution
</a>
</em>
</td>
<td>
<p>LoadDistribution: The load distribution policy for this rule.</p>
</td>
</tr>
<tr>
<td>
<code>probe</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>Probe: The reference to the load balancer probe used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatProtocol">
LoadBalancingRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the load balancing rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatLoadDistribution">LoadBalancingRulePropertiesFormatLoadDistribution
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">LoadBalancers_Spec_Properties_LoadBalancingRules</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatARM">LoadBalancingRulePropertiesFormatARM</a>)
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
</tr><tr><td><p>&#34;SourceIP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SourceIPProtocol&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatProtocol">LoadBalancingRulePropertiesFormatProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">LoadBalancers_Spec_Properties_LoadBalancingRules</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatARM">LoadBalancingRulePropertiesFormatARM</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatStatusLoadDistribution">LoadBalancingRulePropertiesFormatStatusLoadDistribution
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">LoadBalancingRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRule_Status">LoadBalancingRule_Status</a>)
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
</tr><tr><td><p>&#34;SourceIP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SourceIPProtocol&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">LoadBalancingRulePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancingRule_StatusARM">LoadBalancingRule_StatusARM</a>)
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
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
that value 0 enables &ldquo;Any Port&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>disableOutboundSnat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
frontend of the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
Acceptable values are between 0 and 65534. Note that value 0 enables &ldquo;Any Port&rdquo;.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>loadDistribution</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatStatusLoadDistribution">
LoadBalancingRulePropertiesFormatStatusLoadDistribution
</a>
</em>
</td>
<td>
<p>LoadDistribution: The load distribution policy for this rule.</p>
</td>
</tr>
<tr>
<td>
<code>probe</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Probe: The reference to the load balancer probe used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.TransportProtocol_Status">
TransportProtocol_Status
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the load balancing rule resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRule_Status">LoadBalancingRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>backendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
that value 0 enables &ldquo;Any Port&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>disableOutboundSnat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
frontend of the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>enableFloatingIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFloatingIP: Configures a virtual machine&rsquo;s endpoint for the floating IP capability required to configure a SQL
AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
This setting can&rsquo;t be changed after you create the endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>FrontendIPConfiguration: A reference to frontend IP addresses.</p>
</td>
</tr>
<tr>
<td>
<code>frontendPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
Acceptable values are between 0 and 65534. Note that value 0 enables &ldquo;Any Port&rdquo;.</p>
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
<p>Id: Resource ID.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
default value is 4 minutes. This element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>loadDistribution</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatStatusLoadDistribution">
LoadBalancingRulePropertiesFormatStatusLoadDistribution
</a>
</em>
</td>
<td>
<p>LoadDistribution: The load distribution policy for this rule.</p>
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
<p>Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>probe</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Probe: The reference to the load balancer probe used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.TransportProtocol_Status">
TransportProtocol_Status
</a>
</em>
</td>
<td>
<p>Protocol: The reference to the transport protocol used by the load balancing rule.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the load balancing rule resource.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.LoadBalancingRule_StatusARM">LoadBalancingRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">
LoadBalancingRulePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer load balancing rule.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NatGatewaySkuStatusName">NatGatewaySkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NatGatewaySku_Status">NatGatewaySku_Status</a>, <a href="#network.azure.com/v1beta20201101.NatGatewaySku_StatusARM">NatGatewaySku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NatGatewaySku_Status">NatGatewaySku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbedded">NatGateway_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<a href="#network.azure.com/v1beta20201101.NatGatewaySkuStatusName">
NatGatewaySkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of Nat Gateway SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NatGatewaySku_StatusARM">NatGatewaySku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM">NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM</a>)
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
<a href="#network.azure.com/v1beta20201101.NatGatewaySkuStatusName">
NatGatewaySkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of Nat Gateway SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbedded">NatGateway_Status_PublicIPAddress_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NatGatewaySku_Status">
NatGatewaySku_Status
</a>
</em>
</td>
<td>
<p>Sku: The nat gateway SKU.</p>
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
<p>Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM">NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NatGatewaySku_StatusARM">
NatGatewaySku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The nat gateway SKU.</p>
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
<p>Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterface">NetworkInterface
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkInterfaces">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkInterfaces</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec">
NetworkInterfaces_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
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
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings">
NetworkInterfaceDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: The DNS settings in network interface.</p>
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
<p>EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.</p>
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
<p>EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">
[]NetworkInterfaces_Spec_Properties_IpConfigurations
</a>
</em>
</td>
<td>
<p>IpConfigurations: A list of IPConfigurations of the network interface.</p>
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
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">
NetworkInterface_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings">NetworkInterfaceDnsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec">NetworkInterfaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings</a></p>
</div>
<table>
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
<p>DnsServers: List of DNS servers IP addresses. Use &lsquo;AzureProvidedDNS&rsquo; to switch to azure provided DNS resolution.
&lsquo;AzureProvidedDNS&rsquo; value cannot be combined with other IPs, it must be the only value in dnsServers collection.</p>
</td>
</tr>
<tr>
<td>
<code>internalDnsNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceDnsSettingsARM">NetworkInterfaceDnsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_PropertiesARM">NetworkInterfaces_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings</a></p>
</div>
<table>
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
<p>DnsServers: List of DNS servers IP addresses. Use &lsquo;AzureProvidedDNS&rsquo; to switch to azure provided DNS resolution.
&lsquo;AzureProvidedDNS&rsquo; value cannot be combined with other IPs, it must be the only value in dnsServers collection.</p>
</td>
</tr>
<tr>
<td>
<code>internalDnsNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings_Status">NetworkInterfaceDnsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>appliedDnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppliedDnsServers: If the VM that uses this NIC is part of an Availability Set, then this list will have the union of
all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of
those VMs.</p>
</td>
</tr>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses. Use &lsquo;AzureProvidedDNS&rsquo; to switch to azure provided DNS resolution.
&lsquo;AzureProvidedDNS&rsquo; value cannot be combined with other IPs, it must be the only value in dnsServers collection.</p>
</td>
</tr>
<tr>
<td>
<code>internalDnsNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
network.</p>
</td>
</tr>
<tr>
<td>
<code>internalDomainNameSuffix</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDomainNameSuffix: Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of
the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.</p>
</td>
</tr>
<tr>
<td>
<code>internalFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalFqdn: Fully qualified DNS name supporting internal communications between VMs in the same virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings_StatusARM">NetworkInterfaceDnsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<code>appliedDnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppliedDnsServers: If the VM that uses this NIC is part of an Availability Set, then this list will have the union of
all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of
those VMs.</p>
</td>
</tr>
<tr>
<td>
<code>dnsServers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DnsServers: List of DNS servers IP addresses. Use &lsquo;AzureProvidedDNS&rsquo; to switch to azure provided DNS resolution.
&lsquo;AzureProvidedDNS&rsquo; value cannot be combined with other IPs, it must be the only value in dnsServers collection.</p>
</td>
</tr>
<tr>
<td>
<code>internalDnsNameLabel</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
network.</p>
</td>
</tr>
<tr>
<td>
<code>internalDomainNameSuffix</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalDomainNameSuffix: Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of
the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.</p>
</td>
</tr>
<tr>
<td>
<code>internalFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>InternalFqdn: Fully qualified DNS name supporting internal communications between VMs in the same virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status">NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>fqdns</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Fqdns: List of FQDNs for current private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID for current private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMemberName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequiredMemberName: The required member name for current private link connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM">NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<code>fqdns</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Fqdns: List of FQDNs for current private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID for current private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMemberName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequiredMemberName: The required member name for current private link connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatARM">NetworkInterfaceIPConfigurationPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurationsARM">NetworkInterfaces_Spec_Properties_IpConfigurationsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceIPConfigurationPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceIPConfigurationPropertiesFormat</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Application security groups in which the IP configuration is included.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.</p>
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
<p>Primary: Whether this is a primary customer address on the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion">
NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: Public IP address bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>Subnet: Subnet bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkTaps</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkTaps: The reference to Virtual Network Taps.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion">NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatARM">NetworkInterfaceIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">NetworkInterfaces_Spec_Properties_IpConfigurations</a>)
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
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod">NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatARM">NetworkInterfaceIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">NetworkInterfaces_Spec_Properties_IpConfigurations</a>)
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
<tbody><tr><td><p>&#34;Dynamic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Static&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM">
[]ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM">
[]ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Application security groups in which the IP configuration is included.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM">
[]BackendAddressPool_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM">
[]InboundNatRule_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.</p>
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
<p>Primary: Whether this is a primary customer address on the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkConnectionProperties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM">
NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_StatusARM
</a>
</em>
</td>
<td>
<p>PrivateLinkConnectionProperties: PrivateLinkConnection properties for the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network interface IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM">
PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: Public IP address bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_NetworkInterface_SubResourceEmbeddedARM">
Subnet_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Subnet: Subnet bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkTaps</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM">
[]VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkTaps: The reference to Virtual Network Taps.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">
[]ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded">
[]ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Application security groups in which the IP configuration is included.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">
[]BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbedded">
[]InboundNatRule_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
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
<p>Primary: Whether this is a primary customer address on the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkConnectionProperties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status">
NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status
</a>
</em>
</td>
<td>
<p>PrivateLinkConnectionProperties: PrivateLinkConnection properties for the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network interface IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded">
PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PublicIPAddress: Public IP address bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_NetworkInterface_SubResourceEmbedded">
Subnet_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>Subnet: Subnet bound to the IP configuration.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkTaps</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded">
[]VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>VirtualNetworkTaps: The reference to Virtual Network Taps.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">
NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Properties: Network interface IP configuration properties.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusMigrationPhase">NetworkInterfacePropertiesFormatStatusMigrationPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Abort&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Commit&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Committed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Prepare&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusNicType">NetworkInterfacePropertiesFormatStatusNicType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Elastic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings_StatusARM">
NetworkInterfaceDnsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The DNS settings in network interface.</p>
</td>
</tr>
<tr>
<td>
<code>dscpConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DscpConfiguration: A reference to the dscp configuration to which the network interface is linked.</p>
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
<p>EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.</p>
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
<p>EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.</p>
</td>
</tr>
<tr>
<td>
<code>hostedWorkloads</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>HostedWorkloads: A list of references to linked BareMetal resources.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM">
[]NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: A list of IPConfigurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>macAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>MacAddress: The MAC address of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>migrationPhase</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusMigrationPhase">
NetworkInterfacePropertiesFormatStatusMigrationPhase
</a>
</em>
</td>
<td>
<p>MigrationPhase: Migration phase of Network Interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM">
NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>nicType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusNicType">
NetworkInterfacePropertiesFormatStatusNicType
</a>
</em>
</td>
<td>
<p>NicType: Type of Network Interface resource.</p>
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
<p>Primary: Whether this is a primary network interface on a virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM">
PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: A reference to the private endpoint to which the network interface is linked.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkService</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM">
PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateLinkService: Privatelinkservice of the network interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the network interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>tapConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM">
[]NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>TapConfigurations: A list of TapConfigurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachine</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>VirtualMachine: The reference to a virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface">NetworkInterface</a>)
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
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings_Status">
NetworkInterfaceDnsSettings_Status
</a>
</em>
</td>
<td>
<p>DnsSettings: The DNS settings in network interface.</p>
</td>
</tr>
<tr>
<td>
<code>dscpConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DscpConfiguration: A reference to the dscp configuration to which the network interface is linked.</p>
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
<p>EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.</p>
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
<p>EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>hostedWorkloads</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>HostedWorkloads: A list of references to linked BareMetal resources.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">
[]NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>IpConfigurations: A list of IPConfigurations of the network interface.</p>
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
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>macAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>MacAddress: The MAC address of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>migrationPhase</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusMigrationPhase">
NetworkInterfacePropertiesFormatStatusMigrationPhase
</a>
</em>
</td>
<td>
<p>MigrationPhase: Migration phase of Network Interface resource.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded">
NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>nicType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormatStatusNicType">
NetworkInterfacePropertiesFormatStatusNicType
</a>
</em>
</td>
<td>
<p>NicType: Type of Network Interface resource.</p>
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
<p>Primary: Whether this is a primary network interface on a virtual machine.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded">
PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: A reference to the private endpoint to which the network interface is linked.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkService</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded">
PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateLinkService: Privatelinkservice of the network interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network interface resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the network interface resource.</p>
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
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>tapConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded">
[]NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>TapConfigurations: A list of TapConfigurations of the network interface.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
<tr>
<td>
<code>virtualMachine</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>VirtualMachine: The reference to a virtual machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterface_Status_NetworkInterface_SubResourceEmbeddedARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">
NetworkInterfacePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the network interface.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfacesSpecAPIVersion">NetworkInterfacesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaces_Spec">NetworkInterfaces_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface">NetworkInterface</a>)
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
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceDnsSettings">
NetworkInterfaceDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: The DNS settings in network interface.</p>
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
<p>EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.</p>
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
<p>EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">
[]NetworkInterfaces_Spec_Properties_IpConfigurations
</a>
</em>
</td>
<td>
<p>IpConfigurations: A list of IPConfigurations of the network interface.</p>
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
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaces_SpecARM">NetworkInterfaces_SpecARM
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the network interface.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_PropertiesARM">
NetworkInterfaces_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the network interface.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaces_Spec_PropertiesARM">NetworkInterfaces_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_SpecARM">NetworkInterfaces_SpecARM</a>)
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
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceDnsSettingsARM">
NetworkInterfaceDnsSettingsARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The DNS settings in network interface.</p>
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
<p>EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.</p>
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
<p>EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurationsARM">
[]NetworkInterfaces_Spec_Properties_IpConfigurationsARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: A list of IPConfigurations of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">NetworkInterfaces_Spec_Properties_IpConfigurations
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec">NetworkInterfaces_Spec</a>)
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
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroups: Application security groups in which the IP configuration is included.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerBackendAddressPools</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerInboundNatRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
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
<p>Primary: Whether this is a primary customer address on the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP address of the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion">
NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion
</a>
</em>
</td>
<td>
<p>PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPAddress: Public IP address bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>Subnet: Subnet bound to the IP configuration.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkTaps</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>VirtualNetworkTaps: The reference to Virtual Network Taps.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurationsARM">NetworkInterfaces_Spec_Properties_IpConfigurationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_PropertiesARM">NetworkInterfaces_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatARM">
NetworkInterfaceIPConfigurationPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Network interface IP configuration properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup">NetworkSecurityGroup
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroups_Spec">
NetworkSecurityGroups_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">
NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM</a>)
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
<code>defaultSecurityRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">
[]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>DefaultSecurityRules: The default security rules of network security group.</p>
</td>
</tr>
<tr>
<td>
<code>flowLogs</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FlowLog_Status_SubResourceEmbeddedARM">
[]FlowLog_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>FlowLogs: A collection of references to flow log resources.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">
[]NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: A collection of references to network interfaces.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network security group resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the network security group resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">
[]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>SecurityRules: A collection of security rules of the network security group.</p>
</td>
</tr>
<tr>
<td>
<code>subnets</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">
[]Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>Subnets: A collection of references to subnets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup">NetworkSecurityGroup</a>)
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
<code>defaultSecurityRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded">
[]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>DefaultSecurityRules: The default security rules of network security group.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>flowLogs</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.FlowLog_Status_SubResourceEmbedded">
[]FlowLog_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>FlowLogs: A collection of references to flow log resources.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded">
[]NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>NetworkInterfaces: A collection of references to network interfaces.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the network security group resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the network security group resource.</p>
</td>
</tr>
<tr>
<td>
<code>securityRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded">
[]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>SecurityRules: A collection of security rules of the network security group.</p>
</td>
</tr>
<tr>
<td>
<code>subnets</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded">
[]Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>Subnets: A collection of references to subnets.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">
NetworkSecurityGroupPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the network security group.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded">NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRule">NetworkSecurityGroupsSecurityRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">
NetworkSecurityGroupsSecurityRules_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>access</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatAccess">
SecurityRulePropertiesFormatAccess
</a>
</em>
</td>
<td>
<p>Access: The network traffic is allowed or denied.</p>
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
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: A description for this rule. Restricted to 140 chars.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk &lsquo;*&rsquo; can also be used to
match all source IPs. Default tags such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>destinationApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>DestinationApplicationSecurityGroups: The application security group specified as destination.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used
to match all ports.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationPortRanges: The destination port ranges.</p>
</td>
</tr>
<tr>
<td>
<code>direction</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatDirection">
SecurityRulePropertiesFormatDirection
</a>
</em>
</td>
<td>
<p>Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.</p>
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
reference to a network.azure.com/NetworkSecurityGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int
</em>
</td>
<td>
<p>Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
rule in the collection. The lower the priority number, the higher the priority of the rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatProtocol">
SecurityRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Network protocol this rule applies to.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceAddressPrefix: The CIDR or source IP range. Asterisk &lsquo;*&rsquo; can also be used to match all source IPs. Default tags
such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used. If this is an ingress rule, specifies
where network traffic originates from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourceAddressPrefixes: The CIDR or source IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>SourceApplicationSecurityGroups: The application security group specified as source.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used to match
all ports.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourcePortRanges: The source port ranges.</p>
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">
SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRulesSpecAPIVersion">NetworkSecurityGroupsSecurityRulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">NetworkSecurityGroupsSecurityRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRule">NetworkSecurityGroupsSecurityRule</a>)
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
<code>access</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatAccess">
SecurityRulePropertiesFormatAccess
</a>
</em>
</td>
<td>
<p>Access: The network traffic is allowed or denied.</p>
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
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: A description for this rule. Restricted to 140 chars.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk &lsquo;*&rsquo; can also be used to
match all source IPs. Default tags such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>destinationApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>DestinationApplicationSecurityGroups: The application security group specified as destination.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used
to match all ports.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationPortRanges: The destination port ranges.</p>
</td>
</tr>
<tr>
<td>
<code>direction</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatDirection">
SecurityRulePropertiesFormatDirection
</a>
</em>
</td>
<td>
<p>Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.</p>
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
reference to a network.azure.com/NetworkSecurityGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int
</em>
</td>
<td>
<p>Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
rule in the collection. The lower the priority number, the higher the priority of the rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatProtocol">
SecurityRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Network protocol this rule applies to.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceAddressPrefix: The CIDR or source IP range. Asterisk &lsquo;*&rsquo; can also be used to match all source IPs. Default tags
such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used. If this is an ingress rule, specifies
where network traffic originates from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourceAddressPrefixes: The CIDR or source IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>SourceApplicationSecurityGroups: The application security group specified as source.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used to match
all ports.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourcePortRanges: The source port ranges.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_SpecARM">NetworkSecurityGroupsSecurityRules_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">
SecurityRulePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the security rule.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroupsSpecAPIVersion">NetworkSecurityGroupsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroups_Spec">NetworkSecurityGroups_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup">NetworkSecurityGroup</a>)
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.NetworkSecurityGroups_SpecARM">NetworkSecurityGroups_SpecARM
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
<p>Name: Name of the resource</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRulePropertiesFormatARM">OutboundRulePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRulesARM">LoadBalancers_Spec_Properties_OutboundRulesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/OutboundRulePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/OutboundRulePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The number of outbound ports to be used for NAT.</p>
</td>
</tr>
<tr>
<td>
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: The Frontend IP addresses of the load balancer.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatProtocol">
OutboundRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol for the outbound rule in load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRulePropertiesFormatProtocol">OutboundRulePropertiesFormatProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules">LoadBalancers_Spec_Properties_OutboundRules</a>, <a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatARM">OutboundRulePropertiesFormatARM</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRulePropertiesFormatStatusProtocol">OutboundRulePropertiesFormatStatusProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormat_StatusARM">OutboundRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.OutboundRule_Status">OutboundRule_Status</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRulePropertiesFormat_StatusARM">OutboundRulePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.OutboundRule_StatusARM">OutboundRule_StatusARM</a>)
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
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The number of outbound ports to be used for NAT.</p>
</td>
</tr>
<tr>
<td>
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: The Frontend IP addresses of the load balancer.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatStatusProtocol">
OutboundRulePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol for the outbound rule in load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the outbound rule resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRule_Status">OutboundRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The number of outbound ports to be used for NAT.</p>
</td>
</tr>
<tr>
<td>
<code>backendAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
IPs.</p>
</td>
</tr>
<tr>
<td>
<code>enableTcpReset</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
element is only used when the protocol is set to TCP.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>frontendIPConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>FrontendIPConfigurations: The Frontend IP addresses of the load balancer.</p>
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
<p>Id: Resource ID.</p>
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
<p>IdleTimeoutInMinutes: The timeout for the TCP idle connection.</p>
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
<p>Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatStatusProtocol">
OutboundRulePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol for the outbound rule in load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the outbound rule resource.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.OutboundRule_StatusARM">OutboundRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormat_StatusARM">
OutboundRulePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer outbound rule.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PfsGroup_Status">PfsGroup_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">IpsecPolicy_Status</a>, <a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">IpsecPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;ECP256&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ECP384&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS2048&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFS24&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PFSMM&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded">PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM">PrivateEndpoint_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded">PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded">PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM">PrivateLinkService_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the load balancer.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ProbePropertiesFormatARM">ProbePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_ProbesARM">LoadBalancers_Spec_Properties_ProbesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ProbePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ProbePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
the instance out of rotation. The default value is 15, the minimum value is 5.</p>
</td>
</tr>
<tr>
<td>
<code>numberOfProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
in Azure.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatProtocol">
ProbePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol of the end point. If &lsquo;Tcp&rsquo; is specified, a received ACK is required for the probe to be
successful. If &lsquo;Http&rsquo; or &lsquo;Https&rsquo; is specified, a 200 OK response from the specifies URI is required for the probe to be
successful.</p>
</td>
</tr>
<tr>
<td>
<code>requestPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
Otherwise, it is not allowed. There is no default value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ProbePropertiesFormatProtocol">ProbePropertiesFormatProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_Probes">LoadBalancers_Spec_Properties_Probes</a>, <a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatARM">ProbePropertiesFormatARM</a>)
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
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ProbePropertiesFormatStatusProtocol">ProbePropertiesFormatStatusProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ProbePropertiesFormat_StatusARM">ProbePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.Probe_Status">Probe_Status</a>)
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
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ProbePropertiesFormat_StatusARM">ProbePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Probe_StatusARM">Probe_StatusARM</a>)
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
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
the instance out of rotation. The default value is 15, the minimum value is 5.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: The load balancer rules that use this probe.</p>
</td>
</tr>
<tr>
<td>
<code>numberOfProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
in Azure.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatStatusProtocol">
ProbePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol of the end point. If &lsquo;Tcp&rsquo; is specified, a received ACK is required for the probe to be
successful. If &lsquo;Http&rsquo; or &lsquo;Https&rsquo; is specified, a 200 OK response from the specifies URI is required for the probe to be
successful.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the probe resource.</p>
</td>
</tr>
<tr>
<td>
<code>requestPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
Otherwise, it is not allowed. There is no default value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Probe_Status">Probe_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
the instance out of rotation. The default value is 15, the minimum value is 5.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingRules</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>LoadBalancingRules: The load balancer rules that use this probe.</p>
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
<p>Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>numberOfProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
in Azure.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormatStatusProtocol">
ProbePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: The protocol of the end point. If &lsquo;Tcp&rsquo; is specified, a received ACK is required for the probe to be
successful. If &lsquo;Http&rsquo; or &lsquo;Https&rsquo; is specified, a 200 OK response from the specifies URI is required for the probe to be
successful.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the probe resource.</p>
</td>
</tr>
<tr>
<td>
<code>requestPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
Otherwise, it is not allowed. There is no default value.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Probe_StatusARM">Probe_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProbePropertiesFormat_StatusARM">
ProbePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of load balancer probe.</p>
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
<p>Type: Type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ProvisioningState_Status">ProvisioningState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">ApplicationGatewayBackendAddressPoolPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded">ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM">ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_Status">ApplicationGatewayIPConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.Delegation_Status">Delegation_Status</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded">IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormat_StatusARM">InboundNatPoolPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPool_Status">InboundNatPool_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerPropertiesFormat_StatusARM">LoadBalancerPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancer_Status">LoadBalancer_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">LoadBalancingRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRule_Status">LoadBalancingRule_Status</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormat_StatusARM">OutboundRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.OutboundRule_Status">OutboundRule_Status</a>, <a href="#network.azure.com/v1beta20201101.ProbePropertiesFormat_StatusARM">ProbePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.Probe_Status">Probe_Status</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.ResourceNavigationLinkFormat_StatusARM">ResourceNavigationLinkFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.ResourceNavigationLink_Status">ResourceNavigationLink_Status</a>, <a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.ServiceAssociationLinkPropertiesFormat_StatusARM">ServiceAssociationLinkPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.ServiceAssociationLink_Status">ServiceAssociationLink_Status</a>, <a href="#network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormat_StatusARM">ServiceDelegationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_Status">ServiceEndpointPropertiesFormat_Status</a>, <a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_StatusARM">ServiceEndpointPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM">VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status">VirtualNetworkGatewayIPConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>, <a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificatePropertiesFormat_StatusARM">VpnClientRevokedCertificatePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificate_Status">VpnClientRevokedCertificate_Status</a>, <a href="#network.azure.com/v1beta20201101.VpnClientRootCertificatePropertiesFormat_StatusARM">VpnClientRootCertificatePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VpnClientRootCertificate_Status">VpnClientRootCertificate_Status</a>)
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
<tbody><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress">PublicIPAddress
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/publicIPAddresses">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/publicIPAddresses</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">
PublicIPAddresses_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
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
<code>ddosSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettings">
DdosSettings
</a>
</em>
</td>
<td>
<p>DdosSettings: The DDoS protection custom policy associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressDnsSettings">
PublicIPAddressDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: The FQDN of the DNS record associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: The IP address associated with the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpTag">
[]IpTag
</a>
</em>
</td>
<td>
<p>IpTags: The list of tags associated with the public IP address.</p>
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
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAddressVersion">
PublicIPAddressPropertiesFormatPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: The public IP address version.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAllocationMethod">
PublicIPAddressPropertiesFormatPublicIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PublicIPAllocationMethod: The public IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku">
PublicIPAddressSku
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">
PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressDnsSettings">PublicIPAddressDnsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings</a></p>
</div>
<table>
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
<p>DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
record is created for the public IP in the Microsoft Azure DNS system.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
the domainNameLabel and the regionalized DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>reverseFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
to the reverse FQDN.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressDnsSettingsARM">PublicIPAddressDnsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressDnsSettings</a></p>
</div>
<table>
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
<p>DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
record is created for the public IP in the Microsoft Azure DNS system.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
the domainNameLabel and the regionalized DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>reverseFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
to the reverse FQDN.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressDnsSettings_Status">PublicIPAddressDnsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<p>DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
record is created for the public IP in the Microsoft Azure DNS system.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
the domainNameLabel and the regionalized DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>reverseFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
to the reverse FQDN.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressDnsSettings_StatusARM">PublicIPAddressDnsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>)
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
<p>DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
record is created for the public IP in the Microsoft Azure DNS system.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
the domainNameLabel and the regionalized DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>reverseFqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
to the reverse FQDN.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_SpecARM">PublicIPAddresses_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ddosSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettingsARM">
DdosSettingsARM
</a>
</em>
</td>
<td>
<p>DdosSettings: The DDoS protection custom policy associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressDnsSettingsARM">
PublicIPAddressDnsSettingsARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The FQDN of the DNS record associated with the public IP address.</p>
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: The IP address associated with the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpTagARM">
[]IpTagARM
</a>
</em>
</td>
<td>
<p>IpTags: The list of tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAddressVersion">
PublicIPAddressPropertiesFormatPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: The public IP address version.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAllocationMethod">
PublicIPAddressPropertiesFormatPublicIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PublicIPAllocationMethod: The public IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAddressVersion">PublicIPAddressPropertiesFormatPublicIPAddressVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
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
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAllocationMethod">PublicIPAddressPropertiesFormatPublicIPAllocationMethod
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
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
<tbody><tr><td><p>&#34;Dynamic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Static&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatStatusMigrationPhase">PublicIPAddressPropertiesFormatStatusMigrationPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Abort&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Commit&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Committed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Prepare&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM</a>)
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
<code>ddosSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettings_StatusARM">
DdosSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DdosSettings: The DDoS protection custom policy associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressDnsSettings_StatusARM">
PublicIPAddressDnsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>DnsSettings: The FQDN of the DNS record associated with the public IP address.</p>
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: The IP address associated with the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM">
IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>IpConfiguration: The IP configuration associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpTag_StatusARM">
[]IpTag_StatusARM
</a>
</em>
</td>
<td>
<p>IpTags: The list of tags associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>migrationPhase</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatStatusMigrationPhase">
PublicIPAddressPropertiesFormatStatusMigrationPhase
</a>
</em>
</td>
<td>
<p>MigrationPhase: Migration phase of Public IP Address.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM">
NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>NatGateway: The NatGateway for the Public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: The public IP address version.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PublicIPAllocationMethod: The public IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the public IP address resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSku">PublicIPAddressSku
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuName">
PublicIPAddressSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of a public IP address SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuTier">
PublicIPAddressSkuTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a public IP address SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSkuARM">PublicIPAddressSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddresses_SpecARM">PublicIPAddresses_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/PublicIPAddressSku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuName">
PublicIPAddressSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of a public IP address SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuTier">
PublicIPAddressSkuTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a public IP address SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSkuName">PublicIPAddressSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressSku">PublicIPAddressSku</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuARM">PublicIPAddressSkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSkuStatusName">PublicIPAddressSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">PublicIPAddressSku_Status</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">PublicIPAddressSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSkuStatusTier">PublicIPAddressSkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">PublicIPAddressSku_Status</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">PublicIPAddressSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Global&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regional&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSkuTier">PublicIPAddressSkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddressSku">PublicIPAddressSku</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuARM">PublicIPAddressSkuARM</a>)
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
<tbody><tr><td><p>&#34;Global&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Regional&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSku_Status">PublicIPAddressSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded">PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded">PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuStatusName">
PublicIPAddressSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of a public IP address SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuStatusTier">
PublicIPAddressSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a public IP address SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">PublicIPAddressSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM">PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM">PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>)
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
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuStatusName">
PublicIPAddressSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of a public IP address SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuStatusTier">
PublicIPAddressSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Tier of a public IP address SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded">PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">
PublicIPAddressSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM">PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">
PublicIPAddressSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded">PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">
PublicIPAddressSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM">PublicIPAddress_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">
PublicIPAddressSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress">PublicIPAddress</a>)
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
<code>ddosSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettings_Status">
DdosSettings_Status
</a>
</em>
</td>
<td>
<p>DdosSettings: The DDoS protection custom policy associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressDnsSettings_Status">
PublicIPAddressDnsSettings_Status
</a>
</em>
</td>
<td>
<p>DnsSettings: The FQDN of the DNS record associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: The IP address associated with the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded">
IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>IpConfiguration: The IP configuration associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpTag_Status">
[]IpTag_Status
</a>
</em>
</td>
<td>
<p>IpTags: The list of tags associated with the public IP address.</p>
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
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>migrationPhase</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatStatusMigrationPhase">
PublicIPAddressPropertiesFormatStatusMigrationPhase
</a>
</em>
</td>
<td>
<p>MigrationPhase: Migration phase of Public IP Address.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbedded">
NatGateway_Status_PublicIPAddress_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>NatGateway: The NatGateway for the Public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPVersion_Status">
IPVersion_Status
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: The public IP address version.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PublicIPAllocationMethod: The public IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">
PublicIPAddressSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">
PublicIPAddressPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Public IP address properties.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">
PublicIPAddressSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded">IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_Status">
PublicIPAddressSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku_StatusARM">
PublicIPAddressSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddressesSpecAPIVersion">PublicIPAddressesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.PublicIPAddress">PublicIPAddress</a>)
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
<code>ddosSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DdosSettings">
DdosSettings
</a>
</em>
</td>
<td>
<p>DdosSettings: The DDoS protection custom policy associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>dnsSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressDnsSettings">
PublicIPAddressDnsSettings
</a>
</em>
</td>
<td>
<p>DnsSettings: The FQDN of the DNS record associated with the public IP address.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<code>ipAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddress: The IP address associated with the public IP address resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipTags</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpTag">
[]IpTag
</a>
</em>
</td>
<td>
<p>IpTags: The list of tags associated with the public IP address.</p>
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
<code>publicIPAddressVersion</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAddressVersion">
PublicIPAddressPropertiesFormatPublicIPAddressVersion
</a>
</em>
</td>
<td>
<p>PublicIPAddressVersion: The public IP address version.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatPublicIPAllocationMethod">
PublicIPAddressPropertiesFormatPublicIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PublicIPAllocationMethod: The public IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPPrefix</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSku">
PublicIPAddressSku
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.PublicIPAddresses_SpecARM">PublicIPAddresses_SpecARM
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the public ip address.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">
PublicIPAddressPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Public IP address properties.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PublicIPAddressSkuARM">
PublicIPAddressSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The public IP address SKU.</p>
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
<p>Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RadiusServer">RadiusServer
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The address of this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerScore</code><br/>
<em>
int
</em>
</td>
<td>
<p>RadiusServerScore: The initial score assigned to this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The secret used for this radius server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RadiusServerARM">RadiusServerARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/RadiusServer</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The address of this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerScore</code><br/>
<em>
int
</em>
</td>
<td>
<p>RadiusServerScore: The initial score assigned to this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The secret used for this radius server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RadiusServer_Status">RadiusServer_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>)
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
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The address of this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerScore</code><br/>
<em>
int
</em>
</td>
<td>
<p>RadiusServerScore: The initial score assigned to this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The secret used for this radius server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RadiusServer_StatusARM">RadiusServer_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The address of this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerScore</code><br/>
<em>
int
</em>
</td>
<td>
<p>RadiusServerScore: The initial score assigned to this radius server.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The secret used for this radius server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ResourceNavigationLinkFormat_StatusARM">ResourceNavigationLinkFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ResourceNavigationLink_StatusARM">ResourceNavigationLink_StatusARM</a>)
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
<code>link</code><br/>
<em>
string
</em>
</td>
<td>
<p>Link: Link to the external resource.</p>
</td>
</tr>
<tr>
<td>
<code>linkedResourceType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedResourceType: Resource type of the linked resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the resource navigation link resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ResourceNavigationLink_Status">ResourceNavigationLink_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource navigation link identifier.</p>
</td>
</tr>
<tr>
<td>
<code>link</code><br/>
<em>
string
</em>
</td>
<td>
<p>Link: Link to the external resource.</p>
</td>
</tr>
<tr>
<td>
<code>linkedResourceType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedResourceType: Resource type of the linked resource.</p>
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
<p>Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the resource navigation link resource.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ResourceNavigationLink_StatusARM">ResourceNavigationLink_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource navigation link identifier.</p>
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
<p>Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ResourceNavigationLinkFormat_StatusARM">
ResourceNavigationLinkFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Resource navigation link properties format.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded">RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRuleAccess_Status">SecurityRuleAccess_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRuleDirection_Status">SecurityRuleDirection_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Inbound&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Outbound&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">SecurityRulePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_SpecARM">NetworkSecurityGroupsSecurityRules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SecurityRulePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SecurityRulePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>access</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatAccess">
SecurityRulePropertiesFormatAccess
</a>
</em>
</td>
<td>
<p>Access: The network traffic is allowed or denied.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: A description for this rule. Restricted to 140 chars.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk &lsquo;*&rsquo; can also be used to
match all source IPs. Default tags such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>destinationApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>DestinationApplicationSecurityGroups: The application security group specified as destination.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used
to match all ports.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationPortRanges: The destination port ranges.</p>
</td>
</tr>
<tr>
<td>
<code>direction</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatDirection">
SecurityRulePropertiesFormatDirection
</a>
</em>
</td>
<td>
<p>Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int
</em>
</td>
<td>
<p>Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
rule in the collection. The lower the priority number, the higher the priority of the rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatProtocol">
SecurityRulePropertiesFormatProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Network protocol this rule applies to.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceAddressPrefix: The CIDR or source IP range. Asterisk &lsquo;*&rsquo; can also be used to match all source IPs. Default tags
such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used. If this is an ingress rule, specifies
where network traffic originates from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourceAddressPrefixes: The CIDR or source IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>SourceApplicationSecurityGroups: The application security group specified as source.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used to match
all ports.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourcePortRanges: The source port ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormatAccess">SecurityRulePropertiesFormatAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">NetworkSecurityGroupsSecurityRules_Spec</a>, <a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">SecurityRulePropertiesFormatARM</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormatDirection">SecurityRulePropertiesFormatDirection
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">NetworkSecurityGroupsSecurityRules_Spec</a>, <a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">SecurityRulePropertiesFormatARM</a>)
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
<tbody><tr><td><p>&#34;Inbound&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Outbound&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormatProtocol">SecurityRulePropertiesFormatProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">NetworkSecurityGroupsSecurityRules_Spec</a>, <a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">SecurityRulePropertiesFormatARM</a>)
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
<tbody><tr><td><p>&#34;Ah&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Esp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Icmp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;*&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormatStatusProtocol">SecurityRulePropertiesFormatStatusProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Ah&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Esp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Icmp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;*&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">SecurityRulePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM</a>)
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
<code>access</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRuleAccess_Status">
SecurityRuleAccess_Status
</a>
</em>
</td>
<td>
<p>Access: The network traffic is allowed or denied.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: A description for this rule. Restricted to 140 chars.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk &lsquo;*&rsquo; can also be used to
match all source IPs. Default tags such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>destinationApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM">
[]ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>DestinationApplicationSecurityGroups: The application security group specified as destination.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used
to match all ports.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationPortRanges: The destination port ranges.</p>
</td>
</tr>
<tr>
<td>
<code>direction</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRuleDirection_Status">
SecurityRuleDirection_Status
</a>
</em>
</td>
<td>
<p>Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int
</em>
</td>
<td>
<p>Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
rule in the collection. The lower the priority number, the higher the priority of the rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatStatusProtocol">
SecurityRulePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Network protocol this rule applies to.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the security rule resource.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceAddressPrefix: The CIDR or source IP range. Asterisk &lsquo;*&rsquo; can also be used to match all source IPs. Default tags
such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used. If this is an ingress rule, specifies
where network traffic originates from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourceAddressPrefixes: The CIDR or source IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM">
[]ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>SourceApplicationSecurityGroups: The application security group specified as source.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used to match
all ports.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourcePortRanges: The source port ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRule">NetworkSecurityGroupsSecurityRule</a>)
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
<code>access</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRuleAccess_Status">
SecurityRuleAccess_Status
</a>
</em>
</td>
<td>
<p>Access: The network traffic is allowed or denied.</p>
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
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: A description for this rule. Restricted to 140 chars.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk &lsquo;*&rsquo; can also be used to
match all source IPs. Default tags such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used.</p>
</td>
</tr>
<tr>
<td>
<code>destinationAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>destinationApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">
[]ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>DestinationApplicationSecurityGroups: The application security group specified as destination.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used
to match all ports.</p>
</td>
</tr>
<tr>
<td>
<code>destinationPortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DestinationPortRanges: The destination port ranges.</p>
</td>
</tr>
<tr>
<td>
<code>direction</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRuleDirection_Status">
SecurityRuleDirection_Status
</a>
</em>
</td>
<td>
<p>Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int
</em>
</td>
<td>
<p>Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
rule in the collection. The lower the priority number, the higher the priority of the rule.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatStatusProtocol">
SecurityRulePropertiesFormatStatusProtocol
</a>
</em>
</td>
<td>
<p>Protocol: Network protocol this rule applies to.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the security rule resource.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceAddressPrefix: The CIDR or source IP range. Asterisk &lsquo;*&rsquo; can also be used to match all source IPs. Default tags
such as &lsquo;VirtualNetwork&rsquo;, &lsquo;AzureLoadBalancer&rsquo; and &lsquo;Internet&rsquo; can also be used. If this is an ingress rule, specifies
where network traffic originates from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceAddressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourceAddressPrefixes: The CIDR or source IP ranges.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApplicationSecurityGroups</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded">
[]ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>SourceApplicationSecurityGroups: The application security group specified as source.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk &lsquo;*&rsquo; can also be used to match
all ports.</p>
</td>
</tr>
<tr>
<td>
<code>sourcePortRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SourcePortRanges: The source port ranges.</p>
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
<p>Type: The type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM">SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormat_StatusARM">
SecurityRulePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the security rule.</p>
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
<p>Type: The type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceAssociationLinkPropertiesFormat_StatusARM">ServiceAssociationLinkPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ServiceAssociationLink_StatusARM">ServiceAssociationLink_StatusARM</a>)
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
<code>allowDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowDelete: If true, the resource can be deleted.</p>
</td>
</tr>
<tr>
<td>
<code>link</code><br/>
<em>
string
</em>
</td>
<td>
<p>Link: Link to the external resource.</p>
</td>
</tr>
<tr>
<td>
<code>linkedResourceType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedResourceType: Resource type of the linked resource.</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service association link resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceAssociationLink_Status">ServiceAssociationLink_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>allowDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowDelete: If true, the resource can be deleted.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>link</code><br/>
<em>
string
</em>
</td>
<td>
<p>Link: Link to the external resource.</p>
</td>
</tr>
<tr>
<td>
<code>linkedResourceType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedResourceType: Resource type of the linked resource.</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
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
<p>Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service association link resource.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceAssociationLink_StatusARM">ServiceAssociationLink_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceAssociationLinkPropertiesFormat_StatusARM">
ServiceAssociationLinkPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Resource navigation link properties format.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormatARM">ServiceDelegationPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_DelegationsARM">VirtualNetworksSubnets_Spec_Properties_DelegationsARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM">VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceDelegationPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceDelegationPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormat_StatusARM">ServiceDelegationPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Delegation_StatusARM">Delegation_StatusARM</a>)
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
<code>actions</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Actions: The actions permitted to the service upon delegation.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service delegation resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded">ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: Kind of service endpoint policy. This is metadata used for the Azure portal experience.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: Kind of service endpoint policy. This is metadata used for the Azure portal experience.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat">ServiceEndpointPropertiesFormat
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec">VirtualNetworksSubnets_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>Service: The type of the endpoint service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormatARM">ServiceEndpointPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_PropertiesARM">VirtualNetworksSubnets_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_PropertiesARM">VirtualNetworks_Spec_Properties_Subnets_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>Service: The type of the endpoint service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_Status">ServiceEndpointPropertiesFormat_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service endpoint resource.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>Service: The type of the endpoint service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_StatusARM">ServiceEndpointPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>)
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
<code>locations</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Locations: A list of locations.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the service endpoint resource.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>Service: The type of the endpoint service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubResource">SubResource
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.DdosSettings">DdosSettings</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses">LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations">LoadBalancers_Spec_Properties_FrontendIPConfigurations</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools">LoadBalancers_Spec_Properties_InboundNatPools</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules">LoadBalancers_Spec_Properties_LoadBalancingRules</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules">LoadBalancers_Spec_Properties_OutboundRules</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec">NetworkInterfaces_Spec</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_Properties_IpConfigurations">NetworkInterfaces_Spec_Properties_IpConfigurations</a>, <a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupsSecurityRules_Spec">NetworkSecurityGroupsSecurityRules_Spec</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddresses_Spec">PublicIPAddresses_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations">VirtualNetworkGateways_Spec_Properties_IpConfigurations</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec">VirtualNetworksSubnets_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">VirtualNetworksVirtualNetworkPeerings_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource</a></p>
</div>
<table>
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
<p>Reference: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubResourceARM">SubResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.DdosSettingsARM">DdosSettingsARM</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormatARM">FrontendIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormatARM">InboundNatPoolPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancerBackendAddressPropertiesFormatARM">LoadBalancerBackendAddressPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormatARM">LoadBalancingRulePropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormatARM">NetworkInterfaceIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfaces_Spec_PropertiesARM">NetworkInterfaces_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormatARM">OutboundRulePropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormatARM">PublicIPAddressPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.SecurityRulePropertiesFormatARM">SecurityRulePropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatARM">VirtualNetworkGatewayIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">VirtualNetworkPeeringPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_PropertiesARM">VirtualNetworksSubnets_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_PropertiesARM">VirtualNetworks_Spec_Properties_Subnets_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource</a></p>
</div>
<table>
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
<h3 id="network.azure.com/v1beta20201101.SubResource_Status">SubResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_Status">ApplicationGatewayIPConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.DdosSettings_Status">DdosSettings_Status</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPool_Status">InboundNatPool_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRule_Status">LoadBalancingRule_Status</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded">NetworkInterface_Status_NetworkInterface_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.OutboundRule_Status">OutboundRule_Status</a>, <a href="#network.azure.com/v1beta20201101.Probe_Status">Probe_Status</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded">PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status">VirtualNetworkGatewayIPConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubResource_StatusARM">SubResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM">ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.DdosSettings_StatusARM">DdosSettings_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormat_StatusARM">InboundNatPoolPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">LoadBalancingRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.NetworkInterfacePropertiesFormat_StatusARM">NetworkInterfacePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.OutboundRulePropertiesFormat_StatusARM">OutboundRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.ProbePropertiesFormat_StatusARM">ProbePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.PublicIPAddressPropertiesFormat_StatusARM">PublicIPAddressPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM">VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies">SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies">SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded</a>)
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
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">SubnetPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM</a>)
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
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>applicationGatewayIpConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_StatusARM">
[]ApplicationGatewayIPConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>ApplicationGatewayIpConfigurations: Application gateway IP configurations of virtual network resource.</p>
</td>
</tr>
<tr>
<td>
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Delegation_StatusARM">
[]Delegation_StatusARM
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurationProfiles</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
[]IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>IpConfigurationProfiles: Array of IP configuration profiles which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
[]IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: An array of references to the network interface IP configurations using subnet.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies">
SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies
</a>
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
[]PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: An array of references to private endpoints.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies">
SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the subnet resource.</p>
</td>
</tr>
<tr>
<td>
<code>purpose</code><br/>
<em>
string
</em>
</td>
<td>
<p>Purpose: A read-only string identifying the intention of use for this subnet based on delegations and other user-defined
properties.</p>
</td>
</tr>
<tr>
<td>
<code>resourceNavigationLinks</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ResourceNavigationLink_StatusARM">
[]ResourceNavigationLink_StatusARM
</a>
</em>
</td>
<td>
<p>ResourceNavigationLinks: An array of references to the external resources using subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAssociationLinks</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceAssociationLink_StatusARM">
[]ServiceAssociationLink_StatusARM
</a>
</em>
</td>
<td>
<p>ServiceAssociationLinks: An array of references to services injecting into this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">
[]ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_StatusARM">
[]ServiceEndpointPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_LoadBalancer_SubResourceEmbedded">Subnet_Status_LoadBalancer_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded">FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_LoadBalancer_SubResourceEmbeddedARM">Subnet_Status_LoadBalancer_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM">FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_NetworkInterface_SubResourceEmbedded">Subnet_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_NetworkInterface_SubResourceEmbeddedARM">Subnet_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded">Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded">NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM">Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkSecurityGroupPropertiesFormat_StatusARM">NetworkSecurityGroupPropertiesFormat_StatusARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_PublicIPAddress_SubResourceEmbedded">Subnet_Status_PublicIPAddress_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded">IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM">Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM">IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnet">VirtualNetworksSubnet</a>)
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
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>applicationGatewayIpConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ApplicationGatewayIPConfiguration_Status">
[]ApplicationGatewayIPConfiguration_Status
</a>
</em>
</td>
<td>
<p>ApplicationGatewayIpConfigurations: Application gateway IP configurations of virtual network resource.</p>
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
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Delegation_Status">
[]Delegation_Status
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurationProfiles</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded">
[]IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>IpConfigurationProfiles: Array of IP configuration profiles which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded">
[]IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>IpConfigurations: An array of references to the network interface IP configurations using subnet.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded">
NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies">
SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies
</a>
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded">
[]PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: An array of references to private endpoints.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies">
SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the subnet resource.</p>
</td>
</tr>
<tr>
<td>
<code>purpose</code><br/>
<em>
string
</em>
</td>
<td>
<p>Purpose: A read-only string identifying the intention of use for this subnet based on delegations and other user-defined
properties.</p>
</td>
</tr>
<tr>
<td>
<code>resourceNavigationLinks</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ResourceNavigationLink_Status">
[]ResourceNavigationLink_Status
</a>
</em>
</td>
<td>
<p>ResourceNavigationLinks: An array of references to the external resources using subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded">
RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAssociationLinks</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceAssociationLink_Status">
[]ServiceAssociationLink_Status
</a>
</em>
</td>
<td>
<p>ServiceAssociationLinks: An array of references to services injecting into this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded">
[]ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat_Status">
[]ServiceEndpointPropertiesFormat_Status
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM">Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubnetPropertiesFormat_StatusARM">
SubnetPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.TransportProtocol_Status">TransportProtocol_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.InboundNatPoolPropertiesFormat_StatusARM">InboundNatPoolPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.InboundNatPool_Status">InboundNatPool_Status</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRulePropertiesFormat_StatusARM">LoadBalancingRulePropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.LoadBalancingRule_Status">LoadBalancingRule_Status</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tcp&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Udp&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetwork">VirtualNetwork
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">
VirtualNetworks_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>addressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.</p>
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
<code>bgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities">
VirtualNetworkBgpCommunities
</a>
</em>
</td>
<td>
<p>BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>ddosProtectionPlan</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>DdosProtectionPlan: The DDoS protection plan associated with the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>dhcpOptions</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhcpOptions">
DhcpOptions
</a>
</em>
</td>
<td>
<p>DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>enableDdosProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
requires a DDoS protection plan associated with the resource.</p>
</td>
</tr>
<tr>
<td>
<code>enableVmProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this VNET.</p>
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">
VirtualNetwork_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities">VirtualNetworkBgpCommunities
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">VirtualNetworksVirtualNetworkPeerings_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>virtualNetworkCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>VirtualNetworkCommunity: The BGP community associated with the virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkBgpCommunitiesARM">VirtualNetworkBgpCommunitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">VirtualNetworkPeeringPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>virtualNetworkCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>VirtualNetworkCommunity: The BGP community associated with the virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_Status">VirtualNetworkBgpCommunities_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status</a>)
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
<code>regionalCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>RegionalCommunity: The BGP community associated with the region of the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>VirtualNetworkCommunity: The BGP community associated with the virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_StatusARM">VirtualNetworkBgpCommunities_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM</a>)
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
<code>regionalCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>RegionalCommunity: The BGP community associated with the region of the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkCommunity</code><br/>
<em>
string
</em>
</td>
<td>
<p>VirtualNetworkCommunity: The BGP community associated with the virtual network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateway">VirtualNetworkGateway
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworkGateways</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">
VirtualNetworkGateways_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>activeActive</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ActiveActive: ActiveActive flag.</p>
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
<code>bgpSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BgpSettings">
BgpSettings
</a>
</em>
</td>
<td>
<p>BgpSettings: Virtual network gateway&rsquo;s BGP speaker settings.</p>
</td>
</tr>
<tr>
<td>
<code>customRoutes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
the customer for virtual network gateway and VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>enableBgp</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBgp: Whether BGP is enabled for this virtual network gateway or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableDnsForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDnsForwarding: Whether dns forwarding is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateIpAddress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayDefaultSite</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
routes. Assign Null value in case of removing existing default site setting.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesGatewayType">
VirtualNetworkGatewaysSpecPropertiesGatewayType
</a>
</em>
</td>
<td>
<p>GatewayType: The type of this virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations">
[]VirtualNetworkGateways_Spec_Properties_IpConfigurations
</a>
</em>
</td>
<td>
<p>IpConfigurations: IP configurations for virtual network gateway.</p>
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
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku">
VirtualNetworkGatewaySku
</a>
</em>
</td>
<td>
<p>Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
gateway.</p>
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
<code>vNetExtendedLocationResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VNetExtendedLocationResourceReference: MAS FIJI customer vnet resource id. VirtualNetworkGateway of type local gateway
is associated with the customer vnet.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkExtendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>VirtualNetworkExtendedLocation: The extended location of type local virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">
VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration
</a>
</em>
</td>
<td>
<p>VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
configurations.</p>
</td>
</tr>
<tr>
<td>
<code>vpnGatewayGeneration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration">
VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration
</a>
</em>
</td>
<td>
<p>VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.</p>
</td>
</tr>
<tr>
<td>
<code>vpnType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnType">
VirtualNetworkGatewaysSpecPropertiesVpnType
</a>
</em>
</td>
<td>
<p>VpnType: The type of this virtual network gateway.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">
VirtualNetworkGateway_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatARM">VirtualNetworkGatewayIPConfigurationPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM">VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewayIPConfigurationPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewayIPConfigurationPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod">VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatARM">VirtualNetworkGatewayIPConfigurationPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations">VirtualNetworkGateways_Spec_Properties_IpConfigurations</a>)
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
<tbody><tr><td><p>&#34;Dynamic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Static&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM">VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_StatusARM">VirtualNetworkGatewayIPConfiguration_StatusARM</a>)
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
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP Address for this gateway.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network gateway IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status">VirtualNetworkGatewayIPConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateIPAddress: Private IP Address for this gateway.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IPAllocationMethod_Status">
IPAllocationMethod_Status
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network gateway IP configuration resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_StatusARM">VirtualNetworkGatewayIPConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM">
VirtualNetworkGatewayIPConfigurationPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network gateway ip configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusGatewayType">VirtualNetworkGatewayPropertiesFormatStatusGatewayType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<tbody><tr><td><p>&#34;ExpressRoute&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LocalGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Vpn&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration">VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<tbody><tr><td><p>&#34;Generation1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Generation2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnType">VirtualNetworkGatewayPropertiesFormatStatusVpnType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<tbody><tr><td><p>&#34;PolicyBased&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RouteBased&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_StatusARM">VirtualNetworkGateway_StatusARM</a>)
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
<code>activeActive</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ActiveActive: ActiveActive flag.</p>
</td>
</tr>
<tr>
<td>
<code>bgpSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BgpSettings_StatusARM">
BgpSettings_StatusARM
</a>
</em>
</td>
<td>
<p>BgpSettings: Virtual network gateway&rsquo;s BGP speaker settings.</p>
</td>
</tr>
<tr>
<td>
<code>customRoutes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_StatusARM">
AddressSpace_StatusARM
</a>
</em>
</td>
<td>
<p>CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
the customer for virtual network gateway and VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>enableBgp</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBgp: Whether BGP is enabled for this virtual network gateway or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableDnsForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDnsForwarding: Whether dns forwarding is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateIpAddress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayDefaultSite</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
routes. Assign Null value in case of removing existing default site setting.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusGatewayType">
VirtualNetworkGatewayPropertiesFormatStatusGatewayType
</a>
</em>
</td>
<td>
<p>GatewayType: The type of this virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>inboundDnsForwardingEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>InboundDnsForwardingEndpoint: The IP address allocated by the gateway to which dns requests can be sent.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_StatusARM">
[]VirtualNetworkGatewayIPConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: IP configurations for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network gateway resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the virtual network gateway resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_StatusARM">
VirtualNetworkGatewaySku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vNetExtendedLocationResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>VNetExtendedLocationResourceId: Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated
with the customer vnet.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">
VpnClientConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
configurations.</p>
</td>
</tr>
<tr>
<td>
<code>vpnGatewayGeneration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration">
VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration
</a>
</em>
</td>
<td>
<p>VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.</p>
</td>
</tr>
<tr>
<td>
<code>vpnType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnType">
VirtualNetworkGatewayPropertiesFormatStatusVpnType
</a>
</em>
</td>
<td>
<p>VpnType: The type of this virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySku">VirtualNetworkGatewaySku
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuName">
VirtualNetworkGatewaySkuName
</a>
</em>
</td>
<td>
<p>Name: Gateway SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuTier">
VirtualNetworkGatewaySkuTier
</a>
</em>
</td>
<td>
<p>Tier: Gateway SKU tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuARM">VirtualNetworkGatewaySkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkGatewaySku</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuName">
VirtualNetworkGatewaySkuName
</a>
</em>
</td>
<td>
<p>Name: Gateway SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuTier">
VirtualNetworkGatewaySkuTier
</a>
</em>
</td>
<td>
<p>Tier: Gateway SKU tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuName">VirtualNetworkGatewaySkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku">VirtualNetworkGatewaySku</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuARM">VirtualNetworkGatewaySkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;HighPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5AZ&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusName">VirtualNetworkGatewaySkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_Status">VirtualNetworkGatewaySku_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_StatusARM">VirtualNetworkGatewaySku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;HighPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5AZ&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusTier">VirtualNetworkGatewaySkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_Status">VirtualNetworkGatewaySku_Status</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_StatusARM">VirtualNetworkGatewaySku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;HighPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5AZ&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuTier">VirtualNetworkGatewaySkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku">VirtualNetworkGatewaySku</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuARM">VirtualNetworkGatewaySkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ErGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;HighPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraPerformance&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw1AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw2AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw3AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw4AZ&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VpnGw5AZ&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_Status">VirtualNetworkGatewaySku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<p>Capacity: The capacity.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusName">
VirtualNetworkGatewaySkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Gateway SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusTier">
VirtualNetworkGatewaySkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Gateway SKU tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_StatusARM">VirtualNetworkGatewaySku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>)
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
<p>Capacity: The capacity.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusName">
VirtualNetworkGatewaySkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Gateway SKU name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuStatusTier">
VirtualNetworkGatewaySkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: Gateway SKU tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway">VirtualNetworkGateway</a>)
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
<code>activeActive</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ActiveActive: ActiveActive flag.</p>
</td>
</tr>
<tr>
<td>
<code>bgpSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BgpSettings_Status">
BgpSettings_Status
</a>
</em>
</td>
<td>
<p>BgpSettings: Virtual network gateway&rsquo;s BGP speaker settings.</p>
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
<code>customRoutes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_Status">
AddressSpace_Status
</a>
</em>
</td>
<td>
<p>CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
the customer for virtual network gateway and VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>enableBgp</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBgp: Whether BGP is enabled for this virtual network gateway or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableDnsForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDnsForwarding: Whether dns forwarding is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateIpAddress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of type local virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayDefaultSite</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
routes. Assign Null value in case of removing existing default site setting.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusGatewayType">
VirtualNetworkGatewayPropertiesFormatStatusGatewayType
</a>
</em>
</td>
<td>
<p>GatewayType: The type of this virtual network gateway.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>inboundDnsForwardingEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>InboundDnsForwardingEndpoint: The IP address allocated by the gateway to which dns requests can be sent.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfiguration_Status">
[]VirtualNetworkGatewayIPConfiguration_Status
</a>
</em>
</td>
<td>
<p>IpConfigurations: IP configurations for virtual network gateway.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network gateway resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resource GUID property of the virtual network gateway resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku_Status">
VirtualNetworkGatewaySku_Status
</a>
</em>
</td>
<td>
<p>Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
gateway.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
<tr>
<td>
<code>vNetExtendedLocationResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>VNetExtendedLocationResourceId: Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated
with the customer vnet.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">
VpnClientConfiguration_Status
</a>
</em>
</td>
<td>
<p>VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
configurations.</p>
</td>
</tr>
<tr>
<td>
<code>vpnGatewayGeneration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration">
VirtualNetworkGatewayPropertiesFormatStatusVpnGatewayGeneration
</a>
</em>
</td>
<td>
<p>VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.</p>
</td>
</tr>
<tr>
<td>
<code>vpnType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormatStatusVpnType">
VirtualNetworkGatewayPropertiesFormatStatusVpnType
</a>
</em>
</td>
<td>
<p>VpnType: The type of this virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateway_StatusARM">VirtualNetworkGateway_StatusARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of type local virtual network gateway.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">
VirtualNetworkGatewayPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network gateway.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecAPIVersion">VirtualNetworkGatewaysSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesGatewayType">VirtualNetworkGatewaysSpecPropertiesGatewayType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
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
<tbody><tr><td><p>&#34;ExpressRoute&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;HyperNet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LocalGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Vpn&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes">VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
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
<tbody><tr><td><p>&#34;AAD&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Certificate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Radius&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols">VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
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
<tbody><tr><td><p>&#34;IkeV2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OpenVPN&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SSTP&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration">VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
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
<tbody><tr><td><p>&#34;Generation1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Generation2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnType">VirtualNetworkGatewaysSpecPropertiesVpnType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
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
<tbody><tr><td><p>&#34;PolicyBased&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RouteBased&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway">VirtualNetworkGateway</a>)
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
<code>activeActive</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ActiveActive: ActiveActive flag.</p>
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
<code>bgpSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BgpSettings">
BgpSettings
</a>
</em>
</td>
<td>
<p>BgpSettings: Virtual network gateway&rsquo;s BGP speaker settings.</p>
</td>
</tr>
<tr>
<td>
<code>customRoutes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
the customer for virtual network gateway and VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>enableBgp</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBgp: Whether BGP is enabled for this virtual network gateway or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableDnsForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDnsForwarding: Whether dns forwarding is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateIpAddress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayDefaultSite</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
routes. Assign Null value in case of removing existing default site setting.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesGatewayType">
VirtualNetworkGatewaysSpecPropertiesGatewayType
</a>
</em>
</td>
<td>
<p>GatewayType: The type of this virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations">
[]VirtualNetworkGateways_Spec_Properties_IpConfigurations
</a>
</em>
</td>
<td>
<p>IpConfigurations: IP configurations for virtual network gateway.</p>
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
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySku">
VirtualNetworkGatewaySku
</a>
</em>
</td>
<td>
<p>Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
gateway.</p>
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
<code>vNetExtendedLocationResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VNetExtendedLocationResourceReference: MAS FIJI customer vnet resource id. VirtualNetworkGateway of type local gateway
is associated with the customer vnet.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkExtendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>VirtualNetworkExtendedLocation: The extended location of type local virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">
VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration
</a>
</em>
</td>
<td>
<p>VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
configurations.</p>
</td>
</tr>
<tr>
<td>
<code>vpnGatewayGeneration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration">
VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration
</a>
</em>
</td>
<td>
<p>VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.</p>
</td>
</tr>
<tr>
<td>
<code>vpnType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnType">
VirtualNetworkGatewaysSpecPropertiesVpnType
</a>
</em>
</td>
<td>
<p>VpnType: The type of this virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_SpecARM">VirtualNetworkGateways_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">
VirtualNetworkGateways_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network gateway.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_SpecARM">VirtualNetworkGateways_SpecARM</a>)
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
<code>activeActive</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ActiveActive: ActiveActive flag.</p>
</td>
</tr>
<tr>
<td>
<code>bgpSettings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.BgpSettingsARM">
BgpSettingsARM
</a>
</em>
</td>
<td>
<p>BgpSettings: Virtual network gateway&rsquo;s BGP speaker settings.</p>
</td>
</tr>
<tr>
<td>
<code>customRoutes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpaceARM">
AddressSpaceARM
</a>
</em>
</td>
<td>
<p>CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
the customer for virtual network gateway and VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>enableBgp</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBgp: Whether BGP is enabled for this virtual network gateway or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableDnsForwarding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDnsForwarding: Whether dns forwarding is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateIpAddress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayDefaultSite</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
routes. Assign Null value in case of removing existing default site setting.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesGatewayType">
VirtualNetworkGatewaysSpecPropertiesGatewayType
</a>
</em>
</td>
<td>
<p>GatewayType: The type of this virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>ipConfigurations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM">
[]VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM
</a>
</em>
</td>
<td>
<p>IpConfigurations: IP configurations for virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaySkuARM">
VirtualNetworkGatewaySkuARM
</a>
</em>
</td>
<td>
<p>Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vNetExtendedLocationResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkExtendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkExtendedLocation: The extended location of type local virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientConfiguration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">
VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM
</a>
</em>
</td>
<td>
<p>VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
configurations.</p>
</td>
</tr>
<tr>
<td>
<code>vpnGatewayGeneration</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration">
VirtualNetworkGatewaysSpecPropertiesVpnGatewayGeneration
</a>
</em>
</td>
<td>
<p>VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.</p>
</td>
</tr>
<tr>
<td>
<code>vpnType</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnType">
VirtualNetworkGatewaysSpecPropertiesVpnType
</a>
</em>
</td>
<td>
<p>VpnType: The type of this virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurations">VirtualNetworkGateways_Spec_Properties_IpConfigurations
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAllocationMethod</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod">
VirtualNetworkGatewayIPConfigurationPropertiesFormatPrivateIPAllocationMethod
</a>
</em>
</td>
<td>
<p>PrivateIPAllocationMethod: The private IP address allocation method.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddress</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>PublicIPAddress: The reference to the public IP resource.</p>
</td>
</tr>
<tr>
<td>
<code>subnet</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>Subnet: The reference to the subnet resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM">VirtualNetworkGateways_Spec_Properties_IpConfigurationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormatARM">
VirtualNetworkGatewayIPConfigurationPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network gateway ip configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec">VirtualNetworkGateways_Spec</a>)
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
<code>aadAudience</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadIssuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadTenant</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServers</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RadiusServer">
[]RadiusServer
</a>
</em>
</td>
<td>
<p>RadiusServers: The radiusServers property for multiple radius server configuration.</p>
</td>
</tr>
<tr>
<td>
<code>vpnAuthenticationTypes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes">
[]VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes
</a>
</em>
</td>
<td>
<p>VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientIpsecPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicy">
[]IpsecPolicy
</a>
</em>
</td>
<td>
<p>VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientProtocols</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols">
[]VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols
</a>
</em>
</td>
<td>
<p>VpnClientProtocols: VpnClientProtocols for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRevokedCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates">
[]VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates
</a>
</em>
</td>
<td>
<p>VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRootCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates">
[]VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates
</a>
</em>
</td>
<td>
<p>VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_PropertiesARM">VirtualNetworkGateways_Spec_PropertiesARM</a>)
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
<code>aadAudience</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadIssuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadTenant</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServers</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RadiusServerARM">
[]RadiusServerARM
</a>
</em>
</td>
<td>
<p>RadiusServers: The radiusServers property for multiple radius server configuration.</p>
</td>
</tr>
<tr>
<td>
<code>vpnAuthenticationTypes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes">
[]VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnAuthenticationTypes
</a>
</em>
</td>
<td>
<p>VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpaceARM">
AddressSpaceARM
</a>
</em>
</td>
<td>
<p>VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientIpsecPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicyARM">
[]IpsecPolicyARM
</a>
</em>
</td>
<td>
<p>VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientProtocols</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols">
[]VirtualNetworkGatewaysSpecPropertiesVpnClientConfigurationVpnClientProtocols
</a>
</em>
</td>
<td>
<p>VpnClientProtocols: VpnClientProtocols for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRevokedCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM">
[]VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM
</a>
</em>
</td>
<td>
<p>VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRootCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM">
[]VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM
</a>
</em>
</td>
<td>
<p>VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificates
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: The revoked VPN client certificate thumbprint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificatePropertiesFormatARM">
VpnClientRevokedCertificatePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vpn client revoked certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificates
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicCertData</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicCertData: The certificate public data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfigurationARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRootCertificatePropertiesFormatARM">
VpnClientRootCertificatePropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vpn client root certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">VirtualNetworkPeeringPropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_SpecARM">VirtualNetworksVirtualNetworkPeerings_SpecARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM">VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkPeeringPropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkPeeringPropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowForwardedTraffic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
allowed/disallowed in remote virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowGatewayTransit</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowVirtualNetworkAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
virtual network space.</p>
</td>
</tr>
<tr>
<td>
<code>peeringState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatPeeringState">
VirtualNetworkPeeringPropertiesFormatPeeringState
</a>
</em>
</td>
<td>
<p>PeeringState: The status of the virtual network peering.</p>
</td>
</tr>
<tr>
<td>
<code>remoteAddressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpaceARM">
AddressSpaceARM
</a>
</em>
</td>
<td>
<p>RemoteAddressSpace: The reference to the remote virtual network address space.</p>
</td>
</tr>
<tr>
<td>
<code>remoteBgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunitiesARM">
VirtualNetworkBgpCommunitiesARM
</a>
</em>
</td>
<td>
<p>RemoteBgpCommunities: The reference to the remote virtual network&rsquo;s Bgp Communities.</p>
</td>
</tr>
<tr>
<td>
<code>remoteVirtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
different region (preview). See here to register for the preview and learn more
(<a href="https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering">https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering</a>).</p>
</td>
</tr>
<tr>
<td>
<code>useRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatPeeringState">VirtualNetworkPeeringPropertiesFormatPeeringState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">VirtualNetworkPeeringPropertiesFormatARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">VirtualNetworksVirtualNetworkPeerings_Spec</a>)
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
<tbody><tr><td><p>&#34;Connected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disconnected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Initiated&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatStatusPeeringState">VirtualNetworkPeeringPropertiesFormatStatusPeeringState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM</a>, <a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status</a>)
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
<tbody><tr><td><p>&#34;Connected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disconnected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Initiated&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">VirtualNetworkPeeringPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_StatusARM">VirtualNetworkPeering_StatusARM</a>)
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
<code>allowForwardedTraffic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
allowed/disallowed in remote virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowGatewayTransit</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowVirtualNetworkAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
virtual network space.</p>
</td>
</tr>
<tr>
<td>
<code>doNotVerifyRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.</p>
</td>
</tr>
<tr>
<td>
<code>peeringState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatStatusPeeringState">
VirtualNetworkPeeringPropertiesFormatStatusPeeringState
</a>
</em>
</td>
<td>
<p>PeeringState: The status of the virtual network peering.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network peering resource.</p>
</td>
</tr>
<tr>
<td>
<code>remoteAddressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_StatusARM">
AddressSpace_StatusARM
</a>
</em>
</td>
<td>
<p>RemoteAddressSpace: The reference to the remote virtual network address space.</p>
</td>
</tr>
<tr>
<td>
<code>remoteBgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_StatusARM">
VirtualNetworkBgpCommunities_StatusARM
</a>
</em>
</td>
<td>
<p>RemoteBgpCommunities: The reference to the remote virtual network&rsquo;s Bgp Communities.</p>
</td>
</tr>
<tr>
<td>
<code>remoteVirtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
different region (preview). See here to register for the preview and learn more
(<a href="https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering">https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering</a>).</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resourceGuid property of the Virtual Network peering resource.</p>
</td>
</tr>
<tr>
<td>
<code>useRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">VirtualNetworkPeering_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeering">VirtualNetworksVirtualNetworkPeering</a>)
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
<code>allowForwardedTraffic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
allowed/disallowed in remote virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowGatewayTransit</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowVirtualNetworkAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
virtual network space.</p>
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
<code>doNotVerifyRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>peeringState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatStatusPeeringState">
VirtualNetworkPeeringPropertiesFormatStatusPeeringState
</a>
</em>
</td>
<td>
<p>PeeringState: The status of the virtual network peering.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network peering resource.</p>
</td>
</tr>
<tr>
<td>
<code>remoteAddressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_Status">
AddressSpace_Status
</a>
</em>
</td>
<td>
<p>RemoteAddressSpace: The reference to the remote virtual network address space.</p>
</td>
</tr>
<tr>
<td>
<code>remoteBgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_Status">
VirtualNetworkBgpCommunities_Status
</a>
</em>
</td>
<td>
<p>RemoteBgpCommunities: The reference to the remote virtual network&rsquo;s Bgp Communities.</p>
</td>
</tr>
<tr>
<td>
<code>remoteVirtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
different region (preview). See here to register for the preview and learn more
(<a href="https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering">https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering</a>).</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resourceGuid property of the Virtual Network peering resource.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
<tr>
<td>
<code>useRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPeering_StatusARM">VirtualNetworkPeering_StatusARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormat_StatusARM">
VirtualNetworkPeeringPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network peering.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">VirtualNetworkPropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetwork_StatusARM">VirtualNetwork_StatusARM</a>)
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
<code>addressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_StatusARM">
AddressSpace_StatusARM
</a>
</em>
</td>
<td>
<p>AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.</p>
</td>
</tr>
<tr>
<td>
<code>bgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_StatusARM">
VirtualNetworkBgpCommunities_StatusARM
</a>
</em>
</td>
<td>
<p>BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>ddosProtectionPlan</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DdosProtectionPlan: The DDoS protection plan associated with the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>dhcpOptions</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhcpOptions_StatusARM">
DhcpOptions_StatusARM
</a>
</em>
</td>
<td>
<p>DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>enableDdosProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
requires a DDoS protection plan associated with the resource.</p>
</td>
</tr>
<tr>
<td>
<code>enableVmProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_StatusARM">
[]SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resourceGuid property of the Virtual Network resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded">VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded">NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM">VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM">NetworkInterfaceIPConfigurationPropertiesFormat_Status_NetworkInterface_SubResourceEmbeddedARM</a>)
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
<p>Id: Resource ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetwork_Status">VirtualNetwork_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetwork">VirtualNetwork</a>)
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
<code>addressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_Status">
AddressSpace_Status
</a>
</em>
</td>
<td>
<p>AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.</p>
</td>
</tr>
<tr>
<td>
<code>bgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities_Status">
VirtualNetworkBgpCommunities_Status
</a>
</em>
</td>
<td>
<p>BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.</p>
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
<code>ddosProtectionPlan</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DdosProtectionPlan: The DDoS protection plan associated with the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>dhcpOptions</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhcpOptions_Status">
DhcpOptions_Status
</a>
</em>
</td>
<td>
<p>DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>enableDdosProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
requires a DDoS protection plan associated with the resource.</p>
</td>
</tr>
<tr>
<td>
<code>enableVmProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the virtual network.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource_Status">
[]SubResource_Status
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this VNET.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the virtual network resource.</p>
</td>
</tr>
<tr>
<td>
<code>resourceGuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceGuid: The resourceGuid property of the Virtual Network resource.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetwork_StatusARM">VirtualNetwork_StatusARM
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the virtual network.</p>
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
<p>Id: Resource ID.</p>
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
<p>Location: Resource location.</p>
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
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPropertiesFormat_StatusARM">
VirtualNetworkPropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network.</p>
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
<p>Tags: Resource tags.</p>
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
<p>Type: Resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSpecAPIVersion">VirtualNetworksSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnet">VirtualNetworksSubnet
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_subnets">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_subnets</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec">
VirtualNetworksSubnets_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
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
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_Delegations">
[]VirtualNetworksSubnets_Spec_Properties_Delegations
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
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
reference to a network.azure.com/VirtualNetwork resource</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat">
[]ServiceEndpointPropertiesFormat
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded">
Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnetsSpecAPIVersion">VirtualNetworksSubnetsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec">VirtualNetworksSubnets_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnet">VirtualNetworksSubnet</a>)
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
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
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
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_Delegations">
[]VirtualNetworksSubnets_Spec_Properties_Delegations
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
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
reference to a network.azure.com/VirtualNetwork resource</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormat">
[]ServiceEndpointPropertiesFormat
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnets_SpecARM">VirtualNetworksSubnets_SpecARM
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_PropertiesARM">
VirtualNetworksSubnets_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_PropertiesARM">VirtualNetworksSubnets_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_SpecARM">VirtualNetworksSubnets_SpecARM</a>)
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
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_DelegationsARM">
[]VirtualNetworksSubnets_Spec_Properties_DelegationsARM
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormatARM">
[]ServiceEndpointPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_Delegations">VirtualNetworksSubnets_Spec_Properties_Delegations
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec">VirtualNetworksSubnets_Spec</a>)
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
<p>Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceName: The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_Properties_DelegationsARM">VirtualNetworksSubnets_Spec_Properties_DelegationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksSubnets_Spec_PropertiesARM">VirtualNetworksSubnets_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormatARM">
ServiceDelegationPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeering">VirtualNetworksVirtualNetworkPeering
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_virtualNetworkPeerings">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_virtualNetworkPeerings</a></p>
</div>
<table>
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
<a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">
VirtualNetworksVirtualNetworkPeerings_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>allowForwardedTraffic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
allowed/disallowed in remote virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowGatewayTransit</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowVirtualNetworkAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
virtual network space.</p>
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
reference to a network.azure.com/VirtualNetwork resource</p>
</td>
</tr>
<tr>
<td>
<code>peeringState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatPeeringState">
VirtualNetworkPeeringPropertiesFormatPeeringState
</a>
</em>
</td>
<td>
<p>PeeringState: The status of the virtual network peering.</p>
</td>
</tr>
<tr>
<td>
<code>remoteAddressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>RemoteAddressSpace: The reference to the remote virtual network address space.</p>
</td>
</tr>
<tr>
<td>
<code>remoteBgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities">
VirtualNetworkBgpCommunities
</a>
</em>
</td>
<td>
<p>RemoteBgpCommunities: The reference to the remote virtual network&rsquo;s Bgp Communities.</p>
</td>
</tr>
<tr>
<td>
<code>remoteVirtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
different region (preview). See here to register for the preview and learn more
(<a href="https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering">https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering</a>).</p>
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
<code>useRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
gateway.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeering_Status">
VirtualNetworkPeering_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeeringsSpecAPIVersion">VirtualNetworksVirtualNetworkPeeringsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec">VirtualNetworksVirtualNetworkPeerings_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeering">VirtualNetworksVirtualNetworkPeering</a>)
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
<code>allowForwardedTraffic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
allowed/disallowed in remote virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowGatewayTransit</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>allowVirtualNetworkAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
virtual network space.</p>
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
reference to a network.azure.com/VirtualNetwork resource</p>
</td>
</tr>
<tr>
<td>
<code>peeringState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatPeeringState">
VirtualNetworkPeeringPropertiesFormatPeeringState
</a>
</em>
</td>
<td>
<p>PeeringState: The status of the virtual network peering.</p>
</td>
</tr>
<tr>
<td>
<code>remoteAddressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>RemoteAddressSpace: The reference to the remote virtual network address space.</p>
</td>
</tr>
<tr>
<td>
<code>remoteBgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities">
VirtualNetworkBgpCommunities
</a>
</em>
</td>
<td>
<p>RemoteBgpCommunities: The reference to the remote virtual network&rsquo;s Bgp Communities.</p>
</td>
</tr>
<tr>
<td>
<code>remoteVirtualNetwork</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
different region (preview). See here to register for the preview and learn more
(<a href="https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering">https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering</a>).</p>
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
<code>useRemoteGateways</code><br/>
<em>
bool
</em>
</td>
<td>
<p>UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworksVirtualNetworkPeerings_SpecARM">VirtualNetworksVirtualNetworkPeerings_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">
VirtualNetworkPeeringPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network peering.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec">VirtualNetworks_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetwork">VirtualNetwork</a>)
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
<code>addressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace">
AddressSpace
</a>
</em>
</td>
<td>
<p>AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.</p>
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
<code>bgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunities">
VirtualNetworkBgpCommunities
</a>
</em>
</td>
<td>
<p>BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>ddosProtectionPlan</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
SubResource
</a>
</em>
</td>
<td>
<p>DdosProtectionPlan: The DDoS protection plan associated with the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>dhcpOptions</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhcpOptions">
DhcpOptions
</a>
</em>
</td>
<td>
<p>DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>enableDdosProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
requires a DDoS protection plan associated with the resource.</p>
</td>
</tr>
<tr>
<td>
<code>enableVmProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResource">
[]SubResource
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this VNET.</p>
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_SpecARM">VirtualNetworks_SpecARM
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
<a href="#network.azure.com/v1beta20201101.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the virtual network.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">
VirtualNetworks_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network.</p>
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
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_SpecARM">VirtualNetworks_SpecARM</a>)
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
<code>addressSpace</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpaceARM">
AddressSpaceARM
</a>
</em>
</td>
<td>
<p>AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.</p>
</td>
</tr>
<tr>
<td>
<code>bgpCommunities</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkBgpCommunitiesARM">
VirtualNetworkBgpCommunitiesARM
</a>
</em>
</td>
<td>
<p>BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>ddosProtectionPlan</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>DdosProtectionPlan: The DDoS protection plan associated with the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>dhcpOptions</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.DhcpOptionsARM">
DhcpOptionsARM
</a>
</em>
</td>
<td>
<p>DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>enableDdosProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
requires a DDoS protection plan associated with the resource.</p>
</td>
</tr>
<tr>
<td>
<code>enableVmProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this VNET.</p>
</td>
</tr>
<tr>
<td>
<code>subnets</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_SubnetsARM">
[]VirtualNetworks_Spec_Properties_SubnetsARM
</a>
</em>
</td>
<td>
<p>Subnets: A list of subnets in a Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkPeerings</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM">
[]VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkPeerings: A list of peerings in a Virtual Network.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_SubnetsARM">VirtualNetworks_Spec_Properties_SubnetsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_PropertiesARM">
VirtualNetworks_Spec_Properties_Subnets_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_PropertiesARM">VirtualNetworks_Spec_Properties_Subnets_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_SubnetsARM">VirtualNetworks_Spec_Properties_SubnetsARM</a>)
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
<code>addressPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>AddressPrefix: The address prefix for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>addressPrefixes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AddressPrefixes: List of address prefixes for the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>delegations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM">
[]VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM
</a>
</em>
</td>
<td>
<p>Delegations: An array of references to the delegations on the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>ipAllocations</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>IpAllocations: Array of IpAllocation which reference this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>natGateway</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>NatGateway: Nat gateway associated with this subnet.</p>
</td>
</tr>
<tr>
<td>
<code>networkSecurityGroup</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceNetworkPolicies</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>routeTable</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
<p>RouteTable: The reference to the RouteTable resource.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpointPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.SubResourceARM">
[]SubResourceARM
</a>
</em>
</td>
<td>
<p>ServiceEndpointPolicies: An array of service endpoint policies.</p>
</td>
</tr>
<tr>
<td>
<code>serviceEndpoints</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceEndpointPropertiesFormatARM">
[]ServiceEndpointPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>ServiceEndpoints: An array of service endpoints.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM">VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_Subnets_PropertiesARM">VirtualNetworks_Spec_Properties_Subnets_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ServiceDelegationPropertiesFormatARM">
ServiceDelegationPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the subnet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM">VirtualNetworks_Spec_Properties_VirtualNetworkPeeringsARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworks_Spec_PropertiesARM">VirtualNetworks_Spec_PropertiesARM</a>)
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VirtualNetworkPeeringPropertiesFormatARM">
VirtualNetworkPeeringPropertiesFormatARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the virtual network peering.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnAuthenticationTypes">VpnClientConfigurationStatusVpnAuthenticationTypes
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<tbody><tr><td><p>&#34;AAD&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Certificate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Radius&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnClientProtocols">VpnClientConfigurationStatusVpnClientProtocols
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>, <a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<tbody><tr><td><p>&#34;IkeV2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OpenVPN&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SSTP&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateway_Status">VirtualNetworkGateway_Status</a>)
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
<code>aadAudience</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadIssuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadTenant</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServers</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RadiusServer_Status">
[]RadiusServer_Status
</a>
</em>
</td>
<td>
<p>RadiusServers: The radiusServers property for multiple radius server configuration.</p>
</td>
</tr>
<tr>
<td>
<code>vpnAuthenticationTypes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnAuthenticationTypes">
[]VpnClientConfigurationStatusVpnAuthenticationTypes
</a>
</em>
</td>
<td>
<p>VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_Status">
AddressSpace_Status
</a>
</em>
</td>
<td>
<p>VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientIpsecPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicy_Status">
[]IpsecPolicy_Status
</a>
</em>
</td>
<td>
<p>VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientProtocols</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnClientProtocols">
[]VpnClientConfigurationStatusVpnClientProtocols
</a>
</em>
</td>
<td>
<p>VpnClientProtocols: VpnClientProtocols for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRevokedCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificate_Status">
[]VpnClientRevokedCertificate_Status
</a>
</em>
</td>
<td>
<p>VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRootCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRootCertificate_Status">
[]VpnClientRootCertificate_Status
</a>
</em>
</td>
<td>
<p>VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGatewayPropertiesFormat_StatusARM">VirtualNetworkGatewayPropertiesFormat_StatusARM</a>)
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
<code>aadAudience</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadIssuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>aadTenant</code><br/>
<em>
string
</em>
</td>
<td>
<p>AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
authentication.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServerSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.</p>
</td>
</tr>
<tr>
<td>
<code>radiusServers</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.RadiusServer_StatusARM">
[]RadiusServer_StatusARM
</a>
</em>
</td>
<td>
<p>RadiusServers: The radiusServers property for multiple radius server configuration.</p>
</td>
</tr>
<tr>
<td>
<code>vpnAuthenticationTypes</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnAuthenticationTypes">
[]VpnClientConfigurationStatusVpnAuthenticationTypes
</a>
</em>
</td>
<td>
<p>VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientAddressPool</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.AddressSpace_StatusARM">
AddressSpace_StatusARM
</a>
</em>
</td>
<td>
<p>VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientIpsecPolicies</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.IpsecPolicy_StatusARM">
[]IpsecPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientProtocols</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientConfigurationStatusVpnClientProtocols">
[]VpnClientConfigurationStatusVpnClientProtocols
</a>
</em>
</td>
<td>
<p>VpnClientProtocols: VpnClientProtocols for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRevokedCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificate_StatusARM">
[]VpnClientRevokedCertificate_StatusARM
</a>
</em>
</td>
<td>
<p>VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.</p>
</td>
</tr>
<tr>
<td>
<code>vpnClientRootCertificates</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRootCertificate_StatusARM">
[]VpnClientRootCertificate_StatusARM
</a>
</em>
</td>
<td>
<p>VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRevokedCertificatePropertiesFormatARM">VpnClientRevokedCertificatePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRevokedCertificatesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VpnClientRevokedCertificatePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VpnClientRevokedCertificatePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: The revoked VPN client certificate thumbprint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRevokedCertificatePropertiesFormat_StatusARM">VpnClientRevokedCertificatePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificate_StatusARM">VpnClientRevokedCertificate_StatusARM</a>)
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
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the VPN client revoked certificate resource.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: The revoked VPN client certificate thumbprint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRevokedCertificate_Status">VpnClientRevokedCertificate_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the VPN client revoked certificate resource.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: The revoked VPN client certificate thumbprint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRevokedCertificate_StatusARM">VpnClientRevokedCertificate_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRevokedCertificatePropertiesFormat_StatusARM">
VpnClientRevokedCertificatePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vpn client revoked certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRootCertificatePropertiesFormatARM">VpnClientRootCertificatePropertiesFormatARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM">VirtualNetworkGateways_Spec_Properties_VpnClientConfiguration_VpnClientRootCertificatesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VpnClientRootCertificatePropertiesFormat">https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VpnClientRootCertificatePropertiesFormat</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicCertData</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicCertData: The certificate public data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRootCertificatePropertiesFormat_StatusARM">VpnClientRootCertificatePropertiesFormat_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientRootCertificate_StatusARM">VpnClientRootCertificate_StatusARM</a>)
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
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the VPN client root certificate resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicCertData</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicCertData: The certificate public data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRootCertificate_Status">VpnClientRootCertificate_Status
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_Status">VpnClientConfiguration_Status</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the VPN client root certificate resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicCertData</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicCertData: The certificate public data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="network.azure.com/v1beta20201101.VpnClientRootCertificate_StatusARM">VpnClientRootCertificate_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#network.azure.com/v1beta20201101.VpnClientConfiguration_StatusARM">VpnClientConfiguration_StatusARM</a>)
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
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A unique read-only string that changes whenever the resource is updated.</p>
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#network.azure.com/v1beta20201101.VpnClientRootCertificatePropertiesFormat_StatusARM">
VpnClientRootCertificatePropertiesFormat_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vpn client root certificate.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
