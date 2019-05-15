package onvif4go

import (
	tds "github.com/atagirov/onvif4go/device"
	tt "github.com/atagirov/onvif4go/onvif"
)

type DeviceService struct {
	Client onvifCaller
}

func NewDeviceService(onvifDevice *OnvifDevice) *DeviceService {
	endpoint := "http://" + onvifDevice.xaddr + "/onvif/device_service"

	return &DeviceService{
		Client: NewOnvifClient(endpoint, &onvifDevice.auth),
	}
}

func (s *DeviceService) WithoutAuth() *DeviceService {
	return &DeviceService{
		Client: s.Client.WithoutAuth(),
	}
}

// GetDeviceInformation gets basic device information from the device.
func (s *DeviceService) GetDeviceInformation() (res tds.GetDeviceInformationResponse, err error) {
	err = s.Client.Call(tds.GetDeviceInformation{}, &res)
	return
}

// GetServices returns information about services on the device.
func (s *DeviceService) GetServices(includeCapability bool) (res tds.GetServicesResponse, err error) {
	err = s.Client.Call(tds.GetServices{
		IncludeCapability: includeCapability,
	}, &res)
	return
}

/*
GetSystemDateAndTime gets the device system date and time.
The device shall support the return of the daylight saving setting
and of the manual system date and time (if applicable) or indication
of NTP time (if applicable) through the GetSystemDateAndTime command.

A device shall provide the UTCDateTime information.
*/
func (s *DeviceService) GetSystemDateAndTime() (res tds.GetSystemDateAndTimeResponse, err error) {
	err = s.Client.Call(tds.GetSystemDateAndTime{}, &res)
	return
}

// GetServiceCapabilities Returns the capabilities of the device service.
// The result is returned in a typed answer.
func (s *DeviceService) GetServiceCapabilities() (res tds.GetServiceCapabilitiesResponse, err error) {
	err = s.Client.Call(tds.GetServiceCapabilities{}, &res)
	return
}

// GetCapabilities method has been replaced by the more generic GetServices method.
// For capabilities of individual services refer to the GetServiceCapabilities methods.
func (s *DeviceService) GetCapabilities(categories ...string) (res tds.GetCapabilitiesResponse, err error) {
	binding := make([]tt.CapabilityCategory, 0)
	for _, category := range categories {
		binding = append(binding, tt.CapabilityCategory(category))
	}
	err = s.Client.Call(tds.GetCapabilities{
		Category: binding,
	}, &res)
	return
}

/*
GetScopes requests the scope parameters of a device. The scope parameters are used in
the device discovery to match a probe message, see Section 7. The Scope parameters are of
two different types:
	- Fixed
	- Configurable
Fixed scope parameters are permanent device characteristics and cannot be removed through
the device management interface. The scope type is indicated in the scope list returned
in the get scope parameters response. A device shall support retrieval of discovery scope
parameters through the GetScopes command. As some scope parameters are mandatory,
the device shall return a non-empty scope list in the response.
*/
func (s *DeviceService) GetScopes() (res tds.GetScopesResponse, err error) {
	err = s.Client.Call(tds.GetScopes{}, &res)
	return
}

/*
GetUsers lists the registered users and corresponding credentials on a device.
The device shall support retrieval of registered device users and their credentials
for the user token through the GetUsers command.
*/
func (s *DeviceService) GetUsers() (res tds.GetUsersResponse, err error) {
	err = s.Client.Call(tds.GetUsers{}, &res)
	return
}

/*
GetNetworkProtocols gets defined network protocols from a device.
The device shall support the GetNetworkProtocols command returning configured network protocols.
*/
func (s *DeviceService) GetNetworkProtocols() (res tds.GetNetworkProtocolsResponse, err error) {
	err = s.Client.Call(tds.GetNetworkProtocols{}, &res)
	return
}

/*
<wsdl:operation name="SetSystemDateAndTime">
	<wsdl:documentation>This operation sets the device system date and time. The device shall support the
		configuration of the daylight saving setting and of the manual system date and time (if
		applicable) or indication of NTP time (if applicable) through the SetSystemDateAndTime
		command. <br/>
		If system time and date are set manually, the client shall include UTCDateTime in the request.<br/>
		A TimeZone token which is not formed according to the rules of IEEE 1003.1 section 8.3 is considered as invalid timezone.<br/>
		The DayLightSavings flag should be set to true to activate any DST settings of the TimeZone string.
		Clear the DayLightSavings flag if the DST portion of the TimeZone settings should be ignored.
	</wsdl:documentation>
	<wsdl:input message="tds:SetSystemDateAndTimeRequest"/>
	<wsdl:output message="tds:SetSystemDateAndTimeResponse"/>
</wsdl:operation>
<wsdl:operation name="SetSystemFactoryDefault">
	<wsdl:documentation>This operation reloads the parameters on the device to their factory default values.</wsdl:documentation>
	<wsdl:input message="tds:SetSystemFactoryDefaultRequest"/>
	<wsdl:output message="tds:SetSystemFactoryDefaultResponse"/>
</wsdl:operation>
<wsdl:operation name="UpgradeSystemFirmware">
	<wsdl:documentation>This operation upgrades a device firmware version. After a successful upgrade the response
		message is sent before the device reboots. The device should support firmware upgrade
		through the UpgradeSystemFirmware command. The exact format of the firmware data is
		outside the scope of this standard.</wsdl:documentation>
	<wsdl:input message="tds:UpgradeSystemFirmwareRequest"/>
	<wsdl:output message="tds:UpgradeSystemFirmwareResponse"/>
</wsdl:operation>
<wsdl:operation name="SystemReboot">
	<wsdl:documentation>This operation reboots the device.</wsdl:documentation>
	<wsdl:input message="tds:SystemRebootRequest"/>
	<wsdl:output message="tds:SystemRebootResponse"/>
</wsdl:operation>
<wsdl:operation name="RestoreSystem">
	<wsdl:documentation>This operation restores the system backup configuration files(s) previously retrieved from a
		device. The device should support restore of backup configuration file(s) through the
		RestoreSystem command. The exact format of the backup configuration file(s) is outside the
		scope of this standard. If the command is supported, it shall accept backup files returned by
		the GetSystemBackup command.</wsdl:documentation>
	<wsdl:input message="tds:RestoreSystemRequest"/>
	<wsdl:output message="tds:RestoreSystemResponse"/>
</wsdl:operation>
<wsdl:operation name="GetSystemBackup">
	<wsdl:documentation>This operation is retrieves system backup configuration file(s) from a device. The device
		should support return of back up configuration file(s) through the GetSystemBackup command.
		The backup is returned with reference to a name and mime-type together with binary data.
		The exact format of the backup configuration files is outside the scope of this standard.</wsdl:documentation>
	<wsdl:input message="tds:GetSystemBackupRequest"/>
	<wsdl:output message="tds:GetSystemBackupResponse"/>
</wsdl:operation>
<wsdl:operation name="GetSystemLog">
	<wsdl:documentation>This operation gets a system log from the device. The exact format of the system logs is outside the scope of this standard.</wsdl:documentation>
	<wsdl:input message="tds:GetSystemLogRequest"/>
	<wsdl:output message="tds:GetSystemLogResponse"/>
</wsdl:operation>
<wsdl:operation name="GetSystemSupportInformation">
	<wsdl:documentation>This operation gets arbitary device diagnostics information from the device.</wsdl:documentation>
	<wsdl:input message="tds:GetSystemSupportInformationRequest"/>
	<wsdl:output message="tds:GetSystemSupportInformationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetScopes">
	<wsdl:documentation>This operation sets the scope parameters of a device. The scope parameters are used in the
		device discovery to match a probe message.
		This operation replaces all existing configurable scope parameters (not fixed parameters). If
		this shall be avoided, one should use the scope add command instead. The device shall
		support configuration of discovery scope parameters through the SetScopes command.</wsdl:documentation>
	<wsdl:input message="tds:SetScopesRequest"/>
	<wsdl:output message="tds:SetScopesResponse"/>
</wsdl:operation>
<wsdl:operation name="AddScopes">
	<wsdl:documentation>This operation adds new configurable scope parameters to a device. The scope parameters
		are used in the device discovery to match a probe message. The device shall
		support addition of discovery scope parameters through the AddScopes command.</wsdl:documentation>
	<wsdl:input message="tds:AddScopesRequest"/>
	<wsdl:output message="tds:AddScopesResponse"/>
</wsdl:operation>
<wsdl:operation name="RemoveScopes">
	<wsdl:documentation>This operation deletes scope-configurable scope parameters from a device. The scope
		parameters are used in the device discovery to match a probe message, see Section 7. The
		device shall support deletion of discovery scope parameters through the RemoveScopes
		command.
		Table</wsdl:documentation>
	<wsdl:input message="tds:RemoveScopesRequest"/>
	<wsdl:output message="tds:RemoveScopesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDiscoveryMode">
	<wsdl:documentation>This operation gets the discovery mode of a device. See Section 7.2 for the definition of the
		different device discovery modes. The device shall support retrieval of the discovery mode
		setting through the GetDiscoveryMode command.</wsdl:documentation>
	<wsdl:input message="tds:GetDiscoveryModeRequest"/>
	<wsdl:output message="tds:GetDiscoveryModeResponse"/>
</wsdl:operation>
<wsdl:operation name="SetDiscoveryMode">
	<wsdl:documentation>This operation sets the discovery mode operation of a device. See Section 7.2 for the
		definition of the different device discovery modes. The device shall support configuration of
		the discovery mode setting through the SetDiscoveryMode command.</wsdl:documentation>
	<wsdl:input message="tds:SetDiscoveryModeRequest"/>
	<wsdl:output message="tds:SetDiscoveryModeResponse"/>
</wsdl:operation>
<wsdl:operation name="GetRemoteDiscoveryMode">
	<wsdl:documentation>This operation gets the remote discovery mode of a device. See Section 7.4 for the definition
		of remote discovery extensions. A device that supports remote discovery shall support
		retrieval of the remote discovery mode setting through the GetRemoteDiscoveryMode
		command.</wsdl:documentation>
	<wsdl:input message="tds:GetRemoteDiscoveryModeRequest"/>
	<wsdl:output message="tds:GetRemoteDiscoveryModeResponse"/>
</wsdl:operation>
<wsdl:operation name="SetRemoteDiscoveryMode">
	<wsdl:documentation>This operation sets the remote discovery mode of operation of a device. See Section 7.4 for
		the definition of remote discovery remote extensions. A device that supports remote discovery
		shall support configuration of the discovery mode setting through the
		SetRemoteDiscoveryMode command.</wsdl:documentation>
	<wsdl:input message="tds:SetRemoteDiscoveryModeRequest"/>
	<wsdl:output message="tds:SetRemoteDiscoveryModeResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDPAddresses">
	<wsdl:documentation>This operation gets the remote DP address or addresses from a device. If the device supports
		remote discovery, as specified in Section 7.4, the device shall support retrieval of the remote
		DP address(es) through the GetDPAddresses command.</wsdl:documentation>
	<wsdl:input message="tds:GetDPAddressesRequest"/>
	<wsdl:output message="tds:GetDPAddressesResponse"/>
</wsdl:operation>
<wsdl:operation name="SetDPAddresses">
	<wsdl:documentation>This operation sets the remote DP address or addresses on a device. If the device supports
		remote discovery, as specified in Section 7.4, the device shall support configuration of the
		remote DP address(es) through the SetDPAddresses command.</wsdl:documentation>
	<wsdl:input message="tds:SetDPAddressesRequest"/>
	<wsdl:output message="tds:SetDPAddressesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetEndpointReference">
	<wsdl:documentation>A client can ask for the device service endpoint reference address property that can be used
		to derive the password equivalent for remote user operation. The device shall support the
		GetEndpointReference command returning the address property of the device service
		endpoint reference.</wsdl:documentation>
	<wsdl:input message="tds:GetEndpointReferenceRequest"/>
	<wsdl:output message="tds:GetEndpointReferenceResponse"/>
</wsdl:operation>
<wsdl:operation name="GetRemoteUser">
	<wsdl:documentation>This operation returns the configured remote user (if any). A device supporting remote user
		handling shall support this operation. The user is only valid for the WS-UserToken profile or
		as a HTTP / RTSP user.<br/>
		The algorithm to use for deriving the password is described in section 5.12.2.1 of the core specification.</wsdl:documentation>
	<wsdl:input message="tds:GetRemoteUserRequest"/>
	<wsdl:output message="tds:GetRemoteUserResponse"/>
</wsdl:operation>
<wsdl:operation name="SetRemoteUser">
	<wsdl:documentation>This operation sets the remote user. A device supporting remote user handling shall support this
		operation. The user is only valid for the WS-UserToken profile or as a HTTP / RTSP user.<br/>
		The password that is set shall always be the original (not derived) password.<br/>
		If UseDerivedPassword is set password derivation shall be done by the device when connecting to a
		remote device.The algorithm to use for deriving the password is described in section 5.12.2.1 of the core specification.<br/>
		To remove the remote user SetRemoteUser should be called without the RemoteUser parameter.</wsdl:documentation>
	<wsdl:input message="tds:SetRemoteUserRequest"/>
	<wsdl:output message="tds:SetRemoteUserResponse"/>
</wsdl:operation>
<wsdl:operation name="CreateUsers">
	<wsdl:documentation>This operation creates new device users and corresponding credentials on a device for authentication purposes.
		The device shall support creation of device users and their credentials through the CreateUsers
		command. Either all users are created successfully or a fault message shall be returned
		without creating any user.<br/>
		ONVIF compliant devices are recommended to support password length of at least 28 bytes,
		as clients may follow the password derivation mechanism which results in 'password
		equivalent' of length 28 bytes, as described in section 3.1.2 of the ONVIF security white paper.</wsdl:documentation>
	<wsdl:input message="tds:CreateUsersRequest"/>
	<wsdl:output message="tds:CreateUsersResponse"/>
</wsdl:operation>
<wsdl:operation name="DeleteUsers">
	<wsdl:documentation>This operation deletes users on a device. The device shall support deletion of device users and their credentials
		through the DeleteUsers command. A device may have one or more fixed users
		that cannot be deleted to ensure access to the unit. Either all users are deleted successfully or a
		fault message shall be returned and no users be deleted.</wsdl:documentation>
	<wsdl:input message="tds:DeleteUsersRequest"/>
	<wsdl:output message="tds:DeleteUsersResponse"/>
</wsdl:operation>
<wsdl:operation name="SetUser">
	<wsdl:documentation>This operation updates the settings for one or several users on a device for authentication purposes.
		The device shall support update of device users and their credentials through the SetUser command.
		Either all change requests are processed successfully or a fault message shall be returned and no change requests be processed.</wsdl:documentation>
	<wsdl:input message="tds:SetUserRequest"/>
	<wsdl:output message="tds:SetUserResponse"/>
</wsdl:operation>
<wsdl:operation name="GetWsdlUrl">
	<wsdl:documentation>It is possible for an endpoint to request a URL that can be used to retrieve the complete
		schema and WSDL definitions of a device. The command gives in return a URL entry point
		where all the necessary product specific WSDL and schema definitions can be retrieved. The
		device shall provide a URL for WSDL and schema download through the GetWsdlUrl command.</wsdl:documentation>
	<wsdl:input message="tds:GetWsdlUrlRequest"/>
	<wsdl:output message="tds:GetWsdlUrlResponse"/>
</wsdl:operation>
<wsdl:operation name="GetHostname">
	<wsdl:documentation>This operation is used by an endpoint to get the hostname from a device. The device shall
		return its hostname configurations through the GetHostname command.</wsdl:documentation>
	<wsdl:input message="tds:GetHostnameRequest"/>
	<wsdl:output message="tds:GetHostnameResponse"/>
</wsdl:operation>
<wsdl:operation name="SetHostname">
	<wsdl:documentation>This operation sets the hostname on a device. It shall be possible to set the device hostname
		configurations through the SetHostname command.<br/>
		A device shall accept string formated according to RFC 1123 section 2.1 or alternatively to RFC 952,
		other string shall be considered as invalid strings.
	</wsdl:documentation>
	<wsdl:input message="tds:SetHostnameRequest"/>
	<wsdl:output message="tds:SetHostnameResponse"/>
</wsdl:operation>
<wsdl:operation name="SetHostnameFromDHCP">
	<wsdl:documentation>This operation controls whether the hostname is set manually or retrieved via DHCP.</wsdl:documentation>
	<wsdl:input message="tds:SetHostnameFromDHCPRequest"/>
	<wsdl:output message="tds:SetHostnameFromDHCPResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDNS">
	<wsdl:documentation>This operation gets the DNS settings from a device. The device shall return its DNS
		configurations through the GetDNS command.</wsdl:documentation>
	<wsdl:input message="tds:GetDNSRequest"/>
	<wsdl:output message="tds:GetDNSResponse"/>
</wsdl:operation>
<wsdl:operation name="SetDNS">
	<wsdl:documentation>This operation sets the DNS settings on a device. It shall be possible to set the device DNS
		configurations through the SetDNS command.</wsdl:documentation>
	<wsdl:input message="tds:SetDNSRequest"/>
	<wsdl:output message="tds:SetDNSResponse"/>
</wsdl:operation>
<wsdl:operation name="GetNTP">
	<wsdl:documentation>This operation gets the NTP settings from a device. If the device supports NTP, it shall be
		possible to get the NTP server settings through the GetNTP command.</wsdl:documentation>
	<wsdl:input message="tds:GetNTPRequest"/>
	<wsdl:output message="tds:GetNTPResponse"/>
</wsdl:operation>
<wsdl:operation name="SetNTP">
	<wsdl:documentation>This operation sets the NTP settings on a device. If the device supports NTP, it shall be
		possible to set the NTP server settings through the SetNTP command.<br/>
		A device shall accept string formated according to RFC 1123 section 2.1 or alternatively to RFC 952,
		other string shall be considered as invalid strings. <br/>
		Changes to the NTP server list will not affect the clock mode DateTimeType. Use SetSystemDateAndTime to activate NTP operation.
	</wsdl:documentation>
	<wsdl:input message="tds:SetNTPRequest"/>
	<wsdl:output message="tds:SetNTPResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDynamicDNS">
	<wsdl:documentation>This operation gets the dynamic DNS settings from a device. If the device supports dynamic
		DNS as specified in [RFC 2136] and [RFC 4702], it shall be possible to get the type, name
		and TTL through the GetDynamicDNS command.</wsdl:documentation>
	<wsdl:input message="tds:GetDynamicDNSRequest"/>
	<wsdl:output message="tds:GetDynamicDNSResponse"/>
</wsdl:operation>
<wsdl:operation name="SetDynamicDNS">
	<wsdl:documentation>This operation sets the dynamic DNS settings on a device. If the device supports dynamic
		DNS as specified in [RFC 2136] and [RFC 4702], it shall be possible to set the type, name
		and TTL through the SetDynamicDNS command.</wsdl:documentation>
	<wsdl:input message="tds:SetDynamicDNSRequest"/>
	<wsdl:output message="tds:SetDynamicDNSResponse"/>
</wsdl:operation>
<wsdl:operation name="GetNetworkInterfaces">
	<wsdl:documentation>This operation gets the network interface configuration from a device. The device shall
		support return of network interface configuration settings as defined by the NetworkInterface
		type through the GetNetworkInterfaces command.</wsdl:documentation>
	<wsdl:input message="tds:GetNetworkInterfacesRequest"/>
	<wsdl:output message="tds:GetNetworkInterfacesResponse"/>
</wsdl:operation>
<wsdl:operation name="SetNetworkInterfaces">
	<wsdl:documentation>This operation sets the network interface configuration on a device. The device shall support
		network configuration of supported network interfaces through the SetNetworkInterfaces
		command.<br/>
		For interoperability with a client unaware of the IEEE 802.11 extension a device shall retain
		its IEEE 802.11 configuration if the IEEE 802.11 configuration element isn’t present in the
		request.</wsdl:documentation>
	<wsdl:input message="tds:SetNetworkInterfacesRequest"/>
	<wsdl:output message="tds:SetNetworkInterfacesResponse"/>
</wsdl:operation>
<wsdl:operation name="SetNetworkProtocols">
	<wsdl:documentation>This operation configures defined network protocols on a device. The device shall support
		configuration of defined network protocols through the SetNetworkProtocols command.</wsdl:documentation>
	<wsdl:input message="tds:SetNetworkProtocolsRequest"/>
	<wsdl:output message="tds:SetNetworkProtocolsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetNetworkDefaultGateway">
	<wsdl:documentation>This operation gets the default gateway settings from a device. The device shall support the
		GetNetworkDefaultGateway command returning configured default gateway address(es).</wsdl:documentation>
	<wsdl:input message="tds:GetNetworkDefaultGatewayRequest"/>
	<wsdl:output message="tds:GetNetworkDefaultGatewayResponse"/>
</wsdl:operation>
<wsdl:operation name="SetNetworkDefaultGateway">
	<wsdl:documentation>This operation sets the default gateway settings on a device. The device shall support
		configuration of default gateway through the SetNetworkDefaultGateway command.</wsdl:documentation>
	<wsdl:input message="tds:SetNetworkDefaultGatewayRequest"/>
	<wsdl:output message="tds:SetNetworkDefaultGatewayResponse"/>
</wsdl:operation>
<wsdl:operation name="GetZeroConfiguration">
	<wsdl:documentation>This operation gets the zero-configuration from a device. If the device supports dynamic IP
		configuration according to [RFC3927], it shall support the return of IPv4 zero configuration
		address and status through the GetZeroConfiguration command.<br/>
	Devices supporting zero configuration on more than one interface shall use the extension to list the additional interface settings.</wsdl:documentation>
	<wsdl:input message="tds:GetZeroConfigurationRequest"/>
	<wsdl:output message="tds:GetZeroConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetZeroConfiguration">
	<wsdl:documentation>This operation sets the zero-configuration. Use GetCapalities to get if zero-zero-configuration is supported or not.</wsdl:documentation>
	<wsdl:input message="tds:SetZeroConfigurationRequest"/>
	<wsdl:output message="tds:SetZeroConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="GetIPAddressFilter">
	<wsdl:documentation>This operation gets the IP address filter settings from a device. If the device supports device
		access control based on IP filtering rules (denied or accepted ranges of IP addresses), the
		device shall support the GetIPAddressFilter command.</wsdl:documentation>
	<wsdl:input message="tds:GetIPAddressFilterRequest"/>
	<wsdl:output message="tds:GetIPAddressFilterResponse"/>
</wsdl:operation>
<wsdl:operation name="SetIPAddressFilter">
	<wsdl:documentation>This operation sets the IP address filter settings on a device. If the device supports device
		access control based on IP filtering rules (denied or accepted ranges of IP addresses), the
		device shall support configuration of IP filtering rules through the SetIPAddressFilter
		command.</wsdl:documentation>
	<wsdl:input message="tds:SetIPAddressFilterRequest"/>
	<wsdl:output message="tds:SetIPAddressFilterResponse"/>
</wsdl:operation>
<wsdl:operation name="AddIPAddressFilter">
	<wsdl:documentation>This operation adds an IP filter address to a device. If the device supports device access
		control based on IP filtering rules (denied or accepted ranges of IP addresses), the device
		shall support adding of IP filtering addresses through the AddIPAddressFilter command.</wsdl:documentation>
	<wsdl:input message="tds:AddIPAddressFilterRequest"/>
	<wsdl:output message="tds:AddIPAddressFilterResponse"/>
</wsdl:operation>
<wsdl:operation name="RemoveIPAddressFilter">
	<wsdl:documentation>This operation deletes an IP filter address from a device. If the device supports device access
		control based on IP filtering rules (denied or accepted ranges of IP addresses), the device
		shall support deletion of IP filtering addresses through the RemoveIPAddressFilter command.</wsdl:documentation>
	<wsdl:input message="tds:RemoveIPAddressFilterRequest"/>
	<wsdl:output message="tds:RemoveIPAddressFilterResponse"/>
</wsdl:operation>
<wsdl:operation name="GetAccessPolicy">
	<wsdl:documentation>Access to different services and sub-sets of services should be subject to access control. The
		WS-Security framework gives the prerequisite for end-point authentication. Authorization
		decisions can then be taken using an access security policy. This standard does not mandate
		any particular policy description format or security policy but this is up to the device
		manufacturer or system provider to choose policy and policy description format of choice.
		However, an access policy (in arbitrary format) can be requested using this command. If the
		device supports access policy settings based on WS-Security authentication, then the device
		shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:GetAccessPolicyRequest"/>
	<wsdl:output message="tds:GetAccessPolicyResponse"/>
</wsdl:operation>
<wsdl:operation name="SetAccessPolicy">
	<wsdl:documentation>This command sets the device access security policy (for more details on the access security
		policy see the Get command). If the device supports access policy settings
		based on WS-Security authentication, then the device shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:SetAccessPolicyRequest"/>
	<wsdl:output message="tds:SetAccessPolicyResponse"/>
</wsdl:operation>
<wsdl:operation name="CreateCertificate">
	<wsdl:documentation>This operation generates a private/public key pair and also can create a self-signed device
		certificate as a result of key pair generation. The certificate is created using a suitable
		onboard key pair generation mechanism.<br/>
		If a device supports onboard key pair generation, the device that supports TLS shall support
		this certificate creation command. And also, if a device supports onboard key pair generation,
		the device that support IEEE 802.1X shall support this command for the purpose of key pair
		generation. Certificates and key pairs are identified using certificate IDs. These IDs are either
		chosen by the certificate generation requester or by the device (in case that no ID value is
		given).</wsdl:documentation>
	<wsdl:input message="tds:CreateCertificateRequest"/>
	<wsdl:output message="tds:CreateCertificateResponse"/>
</wsdl:operation>
<wsdl:operation name="GetCertificates">
	<wsdl:documentation>This operation gets all device server certificates (including self-signed) for the purpose of TLS
		authentication and all device client certificates for the purpose of IEEE 802.1X authentication.
		This command lists only the TLS server certificates and IEEE 802.1X client certificates for the
		device (neither trusted CA certificates nor trusted root certificates). The certificates are
		returned as binary data. A device that supports TLS shall support this command and the
		certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
		rules.</wsdl:documentation>
	<wsdl:input message="tds:GetCertificatesRequest"/>
	<wsdl:output message="tds:GetCertificatesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetCertificatesStatus">
	<wsdl:documentation>This operation is specific to TLS functionality. This operation gets the status
		(enabled/disabled) of the device TLS server certificates. A device that supports TLS shall
		support this command.</wsdl:documentation>
	<wsdl:input message="tds:GetCertificatesStatusRequest"/>
	<wsdl:output message="tds:GetCertificatesStatusResponse"/>
</wsdl:operation>
<wsdl:operation name="SetCertificatesStatus">
	<wsdl:documentation>This operation is specific to TLS functionality. This operation sets the status (enable/disable)
		of the device TLS server certificates. A device that supports TLS shall support this command.
		Typically only one device server certificate is allowed to be enabled at a time.</wsdl:documentation>
	<wsdl:input message="tds:SetCertificatesStatusRequest"/>
	<wsdl:output message="tds:SetCertificatesStatusResponse"/>
</wsdl:operation>
<wsdl:operation name="DeleteCertificates">
	<wsdl:documentation>This operation deletes a certificate or multiple certificates. The device MAY also delete a
		private/public key pair which is coupled with the certificate to be deleted. The device that
		support either TLS or IEEE 802.1X shall support the deletion of a certificate or multiple
		certificates through this command. Either all certificates are deleted successfully or a fault
		message shall be returned without deleting any certificate.</wsdl:documentation>
	<wsdl:input message="tds:DeleteCertificatesRequest"/>
	<wsdl:output message="tds:DeleteCertificatesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetPkcs10Request">
	<wsdl:documentation>This operation requests a PKCS #10 certificate signature request from the device. The
		returned information field shall be either formatted exactly as specified in [PKCS#10] or PEM
		encoded [PKCS#10] format. In order for this command to work, the device must already have
		a private/public key pair. This key pair should be referred by CertificateID as specified in the
		input parameter description. This CertificateID refers to the key pair generated using
		CreateCertificate command.<br/>
		A device that support onboard key pair generation that supports either TLS or IEEE 802.1X
		using client certificate shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:GetPkcs10RequestRequest"/>
	<wsdl:output message="tds:GetPkcs10RequestResponse"/>
</wsdl:operation>
<wsdl:operation name="LoadCertificates">
	<wsdl:documentation>TLS server certificate(s) or IEEE 802.1X client certificate(s) created using the PKCS#10
		certificate request command can be loaded into the device using this command (see Section
		8.4.13). The certificate ID in the request shall be present. The device may sort the received
		certificate(s) based on the public key and subject information in the certificate(s).
		The certificate ID in the request will be the ID value the client wish to have. The device is
		supposed to scan the generated key pairs present in the device to identify which is the
		correspondent key pair with the loaded certificate and then make the link between the
		certificate and the key pair.<br/>
		A device that supports onboard key pair generation that support either TLS or IEEE 802.1X
		shall support this command.<br/>
		The certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
		rules.<br/>
		This command is applicable to any device type, although the parameter name is called for
		historical reasons NVTCertificate.</wsdl:documentation>
	<wsdl:input message="tds:LoadCertificatesRequest"/>
	<wsdl:output message="tds:LoadCertificatesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetClientCertificateMode">
	<wsdl:documentation>This operation is specific to TLS functionality. This operation gets the status
		(enabled/disabled) of the device TLS client authentication. A device that supports TLS shall
		support this command.</wsdl:documentation>
	<wsdl:input message="tds:GetClientCertificateModeRequest"/>
	<wsdl:output message="tds:GetClientCertificateModeResponse"/>
</wsdl:operation>
<wsdl:operation name="SetClientCertificateMode">
	<wsdl:documentation>This operation is specific to TLS functionality. This operation sets the status
		(enabled/disabled) of the device TLS client authentication. A device that supports TLS shall
		support this command.</wsdl:documentation>
	<wsdl:input message="tds:SetClientCertificateModeRequest"/>
	<wsdl:output message="tds:SetClientCertificateModeResponse"/>
</wsdl:operation>
<wsdl:operation name="GetRelayOutputs">
	<wsdl:documentation>This operation gets a list of all available relay outputs and their settings.<br/>
		This method has been depricated with version 2.0. Refer to the DeviceIO service.</wsdl:documentation>
	<wsdl:input message="tds:GetRelayOutputsRequest"/>
	<wsdl:output message="tds:GetRelayOutputsResponse"/>
</wsdl:operation>
<wsdl:operation name="SetRelayOutputSettings">
	<wsdl:documentation>This operation sets the settings of a relay output.
		<br/>This method has been depricated with version 2.0. Refer to the DeviceIO service.</wsdl:documentation>
	<wsdl:input message="tds:SetRelayOutputSettingsRequest"/>
	<wsdl:output message="tds:SetRelayOutputSettingsResponse"/>
</wsdl:operation>
<wsdl:operation name="SetRelayOutputState">
	<wsdl:documentation>This operation sets the state of a relay output.
		<br/>This method has been depricated with version 2.0. Refer to the DeviceIO service.</wsdl:documentation>
	<wsdl:input message="tds:SetRelayOutputStateRequest"/>
	<wsdl:output message="tds:SetRelayOutputStateResponse"/>
</wsdl:operation>
<wsdl:operation name="SendAuxiliaryCommand">
	<wsdl:documentation>Manage auxiliary commands supported by a device, such as controlling an Infrared (IR) lamp,
		a heater or a wiper or a thermometer that is connected to the device.<br/>
		The supported commands can be retrieved via the AuxiliaryCommands capability.<br/>
		Although the name of the auxiliary commands can be freely defined, commands starting with the prefix tt: are
		reserved to define frequently used commands and these reserved commands shall all share the "tt:command|parameter" syntax.
		<ul>
			<li>tt:Wiper|On – Request to start the wiper.</li>
			<li>tt:Wiper|Off – Request to stop the wiper.</li>
			<li>tt:Washer|On – Request to start the washer.</li>
			<li>tt:Washer|Off – Request to stop the washer.</li>
			<li>tt:WashingProcedure|On – Request to start the washing procedure.</li>
			<li>tt: WashingProcedure |Off – Request to stop the washing procedure.</li>
			<li>tt:IRLamp|On – Request to turn ON an IR illuminator attached to the unit.</li>
			<li>tt:IRLamp|Off – Request to turn OFF an IR illuminator attached to the unit.</li>
			<li>tt:IRLamp|Auto – Request to configure an IR illuminator attached to the unit so that it automatically turns ON and OFF.</li>
		</ul>
		A device that indicates auxiliary service capability shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:SendAuxiliaryCommandRequest"/>
	<wsdl:output message="tds:SendAuxiliaryCommandResponse"/>
</wsdl:operation>
<wsdl:operation name="GetCACertificates">
	<wsdl:documentation>CA certificates will be loaded into a device and be used for the sake of following two cases.
		The one is for the purpose of TLS client authentication in TLS server function. The other one
		is for the purpose of Authentication Server authentication in IEEE 802.1X function. This
		operation gets all CA certificates loaded into a device. A device that supports either TLS client
		authentication or IEEE 802.1X shall support this command and the returned certificates shall
		be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding rules.</wsdl:documentation>
	<wsdl:input message="tds:GetCACertificatesRequest"/>
	<wsdl:output message="tds:GetCACertificatesResponse"/>
</wsdl:operation>
<wsdl:operation name="LoadCertificateWithPrivateKey">
	<wsdl:documentation>There might be some cases that a Certificate Authority or some other equivalent creates a
		certificate without having PKCS#10 certificate signing request. In such cases, the certificate
		will be bundled in conjunction with its private key. This command will be used for such use
		case scenarios. The certificate ID in the request is optionally set to the ID value the client
		wish to have. If the certificate ID is not specified in the request, device can choose the ID
		accordingly.<br/>
		This operation imports a private/public key pair into the device.
		The certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
		rules.<br/>
		A device that does not support onboard key pair generation and support either TLS or IEEE
		802.1X using client certificate shall support this command. A device that support onboard key
		pair generation MAY support this command. The security policy of a device that supports this
		operation should make sure that the private key is sufficiently protected.</wsdl:documentation>
	<wsdl:input message="tds:LoadCertificateWithPrivateKeyRequest"/>
	<wsdl:output message="tds:LoadCertificateWithPrivateKeyResponse"/>
</wsdl:operation>
<wsdl:operation name="GetCertificateInformation">
	<wsdl:documentation>This operation requests the information of a certificate specified by certificate ID. The device
		should respond with its “Issuer DN”, “Subject DN”, “Key usage”, "Extended key usage”, “Key
		Length”, “Version”, “Serial Number”, “Signature Algorithm” and “Validity” data as the
		information of the certificate, as long as the device can retrieve such information from the
		specified certificate.<br/>
		A device that supports either TLS or IEEE 802.1X should support this command.</wsdl:documentation>
	<wsdl:input message="tds:GetCertificateInformationRequest"/>
	<wsdl:output message="tds:GetCertificateInformationResponse"/>
</wsdl:operation>
<wsdl:operation name="LoadCACertificates">
	<wsdl:documentation>This command is used when it is necessary to load trusted CA certificates or trusted root
		certificates for the purpose of verification for its counterpart i.e. client certificate verification in
		TLS function or server certificate verification in IEEE 802.1X function.<br/>
		A device that support either TLS or IEEE 802.1X shall support this command. As for the
		supported certificate format, either DER format or PEM format is possible to be used. But a
		device that support this command shall support at least DER format as supported format type.
		The device may sort the received certificate(s) based on the public key and subject
		information in the certificate(s). Either all CA certificates are loaded successfully or a fault
		message shall be returned without loading any CA certificate.</wsdl:documentation>
	<wsdl:input message="tds:LoadCACertificatesRequest"/>
	<wsdl:output message="tds:LoadCACertificatesResponse"/>
</wsdl:operation>
<wsdl:operation name="CreateDot1XConfiguration">
	<wsdl:documentation>This operation newly creates IEEE 802.1X configuration parameter set of the device. The
		device shall support this command if it supports IEEE 802.1X. If the device receives this
		request with already existing configuration token (Dot1XConfigurationToken) specification, the
		device should respond with 'ter:ReferenceToken ' error to indicate there is some configuration
		conflict.</wsdl:documentation>
	<wsdl:input message="tds:CreateDot1XConfigurationRequest"/>
	<wsdl:output message="tds:CreateDot1XConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetDot1XConfiguration">
	<wsdl:documentation>While the CreateDot1XConfiguration command is trying to create a new configuration
		parameter set, this operation modifies existing IEEE 802.1X configuration parameter set of
		the device. A device that support IEEE 802.1X shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:SetDot1XConfigurationRequest"/>
	<wsdl:output message="tds:SetDot1XConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDot1XConfiguration">
	<wsdl:documentation>This operation gets one IEEE 802.1X configuration parameter set from the device by
		specifying the configuration token (Dot1XConfigurationToken).<br/>
		A device that supports IEEE 802.1X shall support this command.
		Regardless of whether the 802.1X method in the retrieved configuration has a password or
		not, the device shall not include the Password element in the response.</wsdl:documentation>
	<wsdl:input message="tds:GetDot1XConfigurationRequest"/>
	<wsdl:output message="tds:GetDot1XConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDot1XConfigurations">
	<wsdl:documentation>This operation gets all the existing IEEE 802.1X configuration parameter sets from the device.
		The device shall respond with all the IEEE 802.1X configurations so that the client can get to
		know how many IEEE 802.1X configurations are existing and how they are configured.<br/>
		A device that support IEEE 802.1X shall support this command.<br/>
		Regardless of whether the 802.1X method in the retrieved configuration has a password or
		not, the device shall not include the Password element in the response.</wsdl:documentation>
	<wsdl:input message="tds:GetDot1XConfigurationsRequest"/>
	<wsdl:output message="tds:GetDot1XConfigurationsResponse"/>
</wsdl:operation>
<wsdl:operation name="DeleteDot1XConfiguration">
	<wsdl:documentation>This operation deletes an IEEE 802.1X configuration parameter set from the device. Which
		configuration should be deleted is specified by the 'Dot1XConfigurationToken' in the request.
		A device that support IEEE 802.1X shall support this command.</wsdl:documentation>
	<wsdl:input message="tds:DeleteDot1XConfigurationRequest"/>
	<wsdl:output message="tds:DeleteDot1XConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDot11Capabilities">
	<wsdl:documentation>This operation returns the IEEE802.11 capabilities. The device shall support
		this operation.</wsdl:documentation>
	<wsdl:input message="tds:GetDot11CapabilitiesRequest"/>
	<wsdl:output message="tds:GetDot11CapabilitiesResponse"/>
</wsdl:operation>
<wsdl:operation name="GetDot11Status">
	<wsdl:documentation>This operation returns the status of a wireless network interface. The device shall support this
		command.</wsdl:documentation>
	<wsdl:input message="tds:GetDot11StatusRequest"/>
	<wsdl:output message="tds:GetDot11StatusResponse"/>
</wsdl:operation>
<wsdl:operation name="ScanAvailableDot11Networks">
	<wsdl:documentation>This operation returns a lists of the wireless networks in range of the device. A device should
		support this operation.</wsdl:documentation>
	<wsdl:input message="tds:ScanAvailableDot11NetworksRequest"/>
	<wsdl:output message="tds:ScanAvailableDot11NetworksResponse"/>
</wsdl:operation>
<wsdl:operation name="GetSystemUris">
	<wsdl:documentation>This operation is used to retrieve URIs from which system information may be downloaded
		using HTTP. URIs may be returned for the following system information:<br/>
		System Logs. Multiple system logs may be returned, of different types. The exact format of
		the system logs is outside the scope of this specification.<br/>
		Support Information. This consists of arbitrary device diagnostics information from a device.
		The exact format of the diagnostic information is outside the scope of this specification.<br/>
		System Backup. The received file is a backup file that can be used to restore the current
		device configuration at a later date. The exact format of the backup configuration file is
		outside the scope of this specification.<br/>
		If the device allows retrieval of system logs, support information or system backup data, it
		should make them available via HTTP GET. If it does, it shall support the GetSystemUris
		command.</wsdl:documentation>
	<wsdl:input message="tds:GetSystemUrisRequest"/>
	<wsdl:output message="tds:GetSystemUrisResponse"/>
</wsdl:operation>
<wsdl:operation name="StartFirmwareUpgrade">
	<wsdl:documentation>This operation initiates a firmware upgrade using the HTTP POST mechanism. The response
		to the command includes an HTTP URL to which the upgrade file may be uploaded. The
		actual upgrade takes place as soon as the HTTP POST operation has completed. The device
		should support firmware upgrade through the StartFirmwareUpgrade command. The exact
		format of the firmware data is outside the scope of this specification.
		Firmware upgrade over HTTP may be achieved using the following steps:<ol>
			<li>Client calls StartFirmwareUpgrade.</li>
			<li>Server responds with upload URI and optional delay value.</li>
			<li>Client waits for delay duration if specified by server.</li>
			<li>Client transmits the firmware image to the upload URI using HTTP POST.</li>
			<li>Server reprograms itself using the uploaded image, then reboots.</li>
		</ol>
		If the firmware upgrade fails because the upgrade file was invalid, the HTTP POST response
		shall be “415 Unsupported Media Type”. If the firmware upgrade fails due to an error at the
		device, the HTTP POST response shall be “500 Internal Server Error”.<br/>
		The value of the Content-Type header in the HTTP POST request shall be “application/octetstream”.</wsdl:documentation>
	<wsdl:input message="tds:StartFirmwareUpgradeRequest"/>
	<wsdl:output message="tds:StartFirmwareUpgradeResponse"/>
</wsdl:operation>
<wsdl:operation name="StartSystemRestore">
	<wsdl:documentation>This operation initiates a system restore from backed up configuration data using the HTTP
		POST mechanism. The response to the command includes an HTTP URL to which the backup
		file may be uploaded. The actual restore takes place as soon as the HTTP POST operation
		has completed. Devices should support system restore through the StartSystemRestore
		command. The exact format of the backup configuration data is outside the scope of this
		specification.<br/>
		System restore over HTTP may be achieved using the following steps:<ol>
			<li>Client calls StartSystemRestore.</li>
			<li>Server responds with upload URI.</li>
			<li>Client transmits the configuration data to the upload URI using HTTP POST.</li>
			<li>Server applies the uploaded configuration, then reboots if necessary.</li>
		</ol>
		If the system restore fails because the uploaded file was invalid, the HTTP POST response
		shall be “415 Unsupported Media Type”. If the system restore fails due to an error at the
		device, the HTTP POST response shall be “500 Internal Server Error”.<br/>
		The value of the Content-Type header in the HTTP POST request shall be “application/octetstream”.</wsdl:documentation>
	<wsdl:input message="tds:StartSystemRestoreRequest"/>
	<wsdl:output message="tds:StartSystemRestoreResponse"/>
</wsdl:operation>

<wsdl:operation name="GetStorageConfigurations">
	<wsdl:documentation>
	This operation lists all existing storage configurations for the device.
	</wsdl:documentation>
	<wsdl:input message="tds:GetStorageConfigurationsRequest"/>
	<wsdl:output message="tds:GetStorageConfigurationsResponse"/>
</wsdl:operation>
<wsdl:operation name="CreateStorageConfiguration">
	<wsdl:documentation>
	This operation creates a new storage configuration.
	The configuration data shall be created in the device and shall be persistent (remain after reboot).
	</wsdl:documentation>
	<wsdl:input message="tds:CreateStorageConfigurationRequest"/>
	<wsdl:output message="tds:CreateStorageConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="GetStorageConfiguration">
	<wsdl:documentation>
	This operation retrieves the Storage configuration associated with the given storage configuration token.
	</wsdl:documentation>
	<wsdl:input message="tds:GetStorageConfigurationRequest"/>
	<wsdl:output message="tds:GetStorageConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetStorageConfiguration">
	<wsdl:documentation>
	This operation modifies an existing Storage configuration.
	</wsdl:documentation>
	<wsdl:input message="tds:SetStorageConfigurationRequest"/>
	<wsdl:output message="tds:SetStorageConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="DeleteStorageConfiguration">
	<wsdl:documentation>
	This operation deletes the given storage configuration and configuration change shall always be persistent.
	</wsdl:documentation>
	<wsdl:input message="tds:DeleteStorageConfigurationRequest"/>
	<wsdl:output message="tds:DeleteStorageConfigurationResponse"/>
</wsdl:operation>

<wsdl:operation name="GetGeoLocation">
	<wsdl:documentation>
		This operation lists all existing geo location configurations for the device.
	</wsdl:documentation>
	<wsdl:input message="tds:GetGeoLocationRequest"/>
	<wsdl:output message="tds:GetGeoLocationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetGeoLocation">
	<wsdl:documentation>
		This operation allows to modify one or more geo configuration entries.
	</wsdl:documentation>
	<wsdl:input message="tds:SetGeoLocationRequest"/>
	<wsdl:output message="tds:SetGeoLocationResponse"/>
</wsdl:operation>
<wsdl:operation name="DeleteGeoLocation">
	<wsdl:documentation>
		This operation deletes the given geo location entries.
	</wsdl:documentation>
	<wsdl:input message="tds:DeleteGeoLocationRequest"/>
	<wsdl:output message="tds:DeleteGeoLocationResponse"/>
</wsdl:operation>
*/
