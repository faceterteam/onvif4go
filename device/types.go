package device

import (
	"time"

	"github.com/faceterteam/onvif4go/onvif"
	"github.com/faceterteam/onvif4go/xsd"
)

type Service struct {
	Namespace xsd.AnyURI
	XAddr     xsd.AnyURI
	Capabilities
	Version onvif.OnvifVersion
}

type Capabilities struct {
	Any string
}

type DeviceServiceCapabilities struct {
	Network  NetworkCapabilities
	Security SecurityCapabilities
	System   SystemCapabilities
	Misc     *MiscCapabilities
}

type NetworkCapabilities struct {
	IPFilter            bool `xml:"IPFilter,attr"`
	ZeroConfiguration   bool `xml:"ZeroConfiguration,attr"`
	IPVersion6          bool `xml:"IPVersion6,attr"`
	DynDNS              bool `xml:"DynDNS,attr"`
	Dot11Configuration  bool `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int  `xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP    bool `xml:"HostnameFromDHCP,attr"`
	NTP                 int  `xml:"NTP,attr"`
	DHCPv6              bool `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS1_0               bool           `xml:"TLS1_0,attr"`
	TLS1_1               bool           `xml:"TLS1_1,attr"`
	TLS1_2               bool           `xml:"TLS1_2,attr"`
	OnboardKeyGeneration bool           `xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig   bool           `xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy  bool           `xml:"DefaultAccessPolicy,attr"`
	Dot1X                bool           `xml:"Dot1X,attr"`
	RemoteUserHandling   bool           `xml:"RemoteUserHandling,attr"`
	X509Token            bool           `xml:"X_509Token,attr"`
	SAMLToken            bool           `xml:"SAMLToken,attr"`
	KerberosToken        bool           `xml:"KerberosToken,attr"`
	UsernameToken        bool           `xml:"UsernameToken,attr"`
	HTTPDigest           bool           `xml:"HttpDigest,attr"`
	RELToken             bool           `xml:"RELToken,attr"`
	SupportedEAPMethods  EAPMethodTypes `xml:"SupportedEAPMethods,attr"`
	MaxUsers             int            `xml:"MaxUsers,attr"`
	MaxUserNameLength    int            `xml:"MaxUserNameLength,attr"`
	MaxPasswordLength    int            `xml:"MaxPasswordLength,attr"`
}

//TODO: <xs:list itemType="xs:int"/>
type EAPMethodTypes struct {
	Types []int
}

type SystemCapabilities struct {
	DiscoveryResolve         bool   `xml:"DiscoveryResolve,attr"`
	DiscoveryBye             bool   `xml:"DiscoveryBye,attr"`
	RemoteDiscovery          bool   `xml:"RemoteDiscovery,attr"`
	SystemBackup             bool   `xml:"SystemBackup,attr"`
	SystemLogging            bool   `xml:"SystemLogging,attr"`
	FirmwareUpgrade          bool   `xml:"FirmwareUpgrade,attr"`
	HTTPFirmwareUpgrade      bool   `xml:"HttpFirmwareUpgrade,attr"`
	HTTPSystemBackup         bool   `xml:"HttpSystemBackup,attr"`
	HTTPSystemLogging        bool   `xml:"HttpSystemLogging,attr"`
	HTTPSupportInformation   bool   `xml:"HttpSupportInformation,attr"`
	StorageConfiguration     bool   `xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations int    `xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries       int    `xml:"GeoLocationEntries,attr"`
	AutoGeo                  string `xml:"AutoGeo,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands onvif.StringAttrList `xml:"AuxiliaryCommands,attr"`
}

type StorageConfiguration struct {
	onvif.DeviceEntity
	Data StorageConfigurationData `xml:"http://www.onvif.org/ver10/device/wsdl Data"`
}

type StorageConfigurationData struct {
	Type       string          `xml:"type,attr"`
	LocalPath  *xsd.AnyURI     `xml:"http://www.onvif.org/ver10/device/wsdl LocalPath"`
	StorageURI *xsd.AnyURI     `xml:"http://www.onvif.org/ver10/device/wsdl StorageUri"`
	User       *UserCredential `xml:"http://www.onvif.org/ver10/device/wsdl User"`
	Extension  *xsd.AnyType    `xml:"http://www.onvif.org/ver10/device/wsdl Extension"`
}

type UserCredential struct {
	UserName  string       `xml:"http://www.onvif.org/ver10/device/wsdl UserName"`
	Password  *string      `xml:"http://www.onvif.org/ver10/device/wsdl Password"`
	Extension *xsd.AnyType `xml:"http://www.onvif.org/ver10/device/wsdl Extension"`
}

type GetServices struct {
	XMLName           string `xml:"http://www.onvif.org/ver10/device/wsdl GetServices"`
	IncludeCapability bool   `xml:"http://www.onvif.org/ver10/device/wsdl IncludeCapability"`
}

type GetServicesResponse struct {
	Service []Service
}

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities DeviceServiceCapabilities
}

type GetDeviceInformation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDeviceInformation"`
}

type GetDeviceInformationResponse struct {
	Manufacturer    string `xml:"http://www.onvif.org/ver10/device/wsdl Manufacturer"`
	Model           string `xml:"http://www.onvif.org/ver10/device/wsdl Model"`
	FirmwareVersion string `xml:"http://www.onvif.org/ver10/device/wsdl FirmwareVersion"`
	SerialNumber    string `xml:"http://www.onvif.org/ver10/device/wsdl SerialNumber"`
	HardwareID      string `xml:"http://www.onvif.org/ver10/device/wsdl HardwareId"`
}

type SetSystemDateAndTime struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemDateAndTime"`
	DateTimeType    onvif.SetDateTimeType `xml:"http://www.onvif.org/ver10/device/wsdl DateTimeType"`
	DaylightSavings bool                  `xml:"http://www.onvif.org/ver10/device/wsdl DaylightSavings"`
	TimeZone        *onvif.TimeZone        `xml:"http://www.onvif.org/ver10/device/wsdl TimeZone"`
	UTCDateTime     *onvif.DateTime        `xml:"http://www.onvif.org/ver10/device/wsdl UTCDateTime"`
}

func NewSetSystemDateAndTimeNTP(timeZone string, useDST bool) (s SetSystemDateAndTime, err error) {
	if timeZone != "" {
		var ns xsd.NormalizedString
		ns, err = xsd.NewNormalizedString(timeZone)
		if err != nil {
			return
		}

		var token xsd.Token
		token, err = xsd.NewToken(ns)
		if err != nil {
			return
		}
		tz := onvif.TimeZone{TZ:token}
		s.TimeZone = &tz
	}

	s.DaylightSavings = useDST

	s.DateTimeType = onvif.SetDateTimeType("NTP")

	return
}

func NewSetSystemDateAndTimeManual(datetime time.Time, timeZone string, useDST bool) (s SetSystemDateAndTime, err error) {
	if timeZone != "" {
		var ns xsd.NormalizedString
		ns, err = xsd.NewNormalizedString(timeZone)
		if err != nil {
			return
		}

		var token xsd.Token
		token, err = xsd.NewToken(ns)
		if err != nil {
			return
		}
		tz := onvif.TimeZone{TZ:token}
		s.TimeZone = &tz
	}

	dt := onvif.DateTime{
		Time: onvif.Time{
			Hour: datetime.Hour(),
			Minute: datetime.Minute(),
			Second: datetime.Second(),
		},
		Date: onvif.Date{
			Day: datetime.Day(),
			Month: int(datetime.Month()),
			Year: datetime.Year(),
		},
	}

	s.UTCDateTime = &dt
	s.DaylightSavings = useDST

	s.DateTimeType = onvif.SetDateTimeType("Manual")

	return
}

type SetSystemDateAndTimeResponse struct {
}

type GetSystemDateAndTime struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemDateAndTime"`
}

type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime onvif.SystemDateTime
}

type SetSystemFactoryDefault struct {
	XMLName        string                   `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemFactoryDefault"`
	FactoryDefault onvif.FactoryDefaultType `xml:"http://www.onvif.org/ver10/device/wsdl FactoryDefault"`
}

type SetSystemFactoryDefaultResponse struct {
}

type UpgradeSystemFirmware struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/device/wsdl UpgradeSystemFirmware"`
	Firmware onvif.AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Firmware"`
}

type UpgradeSystemFirmwareResponse struct {
	Message string
}

type SystemReboot struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl SystemReboot"`
}

type SystemRebootResponse struct {
	Message string
}

type RestoreSystem struct {
	XMLName     string             `xml:"http://www.onvif.org/ver10/device/wsdl RestoreSystem"`
	BackupFiles []onvif.BackupFile `xml:"http://www.onvif.org/ver10/device/wsdl BackupFiles"`
}

type RestoreSystemResponse struct {
}

type GetSystemBackup struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemBackup"`
}

type GetSystemBackupResponse struct {
	BackupFiles []onvif.BackupFile
}

type GetSystemLog struct {
	XMLName string              `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemLog"`
	LogType onvif.SystemLogType `xml:"http://www.onvif.org/ver10/device/wsdl LogType"`
}

type GetSystemLogResponse struct {
	SystemLog onvif.SystemLog
}

type GetSystemSupportInformation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemSupportInformation"`
}

type GetSystemSupportInformationResponse struct {
	SupportInformation onvif.SupportInformation
}

type GetScopes struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetScopes"`
}

type GetScopesResponse struct {
	Scopes []onvif.Scope
}

type SetScopes struct {
	XMLName string       `xml:"http://www.onvif.org/ver10/device/wsdl SetScopes"`
	Scopes  []xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl Scopes"`
}

type SetScopesResponse struct {
}

type AddScopes struct {
	XMLName   string       `xml:"http://www.onvif.org/ver10/device/wsdl AddScopes"`
	ScopeItem []xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl ScopeItem"`
}

type AddScopesResponse struct {
}

type RemoveScopes struct {
	XMLName   string       `xml:"http://www.onvif.org/ver10/device/wsdl RemoveScopes"`
	ScopeItem []xsd.AnyURI `xml:"http://www.onvif.org/ver10/schema onvif:ScopeItem"`
}

type RemoveScopesResponse struct {
	ScopeItem []xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl ScopeItem"`
}

type GetDiscoveryMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDiscoveryMode"`
}

type GetDiscoveryModeResponse struct {
	DiscoveryMode onvif.DiscoveryMode
}

type SetDiscoveryMode struct {
	XMLName       string              `xml:"http://www.onvif.org/ver10/device/wsdl SetDiscoveryMode"`
	DiscoveryMode onvif.DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl DiscoveryMode"`
}

type SetDiscoveryModeResponse struct {
}

type GetRemoteDiscoveryMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteDiscoveryMode"`
}

type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode onvif.DiscoveryMode
}

type SetRemoteDiscoveryMode struct {
	XMLName             string              `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteDiscoveryMode"`
	RemoteDiscoveryMode onvif.DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl RemoteDiscoveryMode"`
}

type SetRemoteDiscoveryModeResponse struct {
}

type GetDPAddresses struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDPAddresses"`
}

type GetDPAddressesResponse struct {
	DPAddress []onvif.NetworkHost
}

type SetDPAddresses struct {
	XMLName   string              `xml:"http://www.onvif.org/ver10/device/wsdl SetDPAddresses"`
	DPAddress []onvif.NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl DPAddress"`
}

type SetDPAddressesResponse struct {
}

type GetEndpointReference struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetEndpointReference"`
}

type GetEndpointReferenceResponse struct {
	GUID string
}

type GetRemoteUser struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteUser"`
}

type GetRemoteUserResponse struct {
	RemoteUser *onvif.RemoteUser
}

type SetRemoteUser struct {
	XMLName    string            `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteUser"`
	RemoteUser *onvif.RemoteUser `xml:"http://www.onvif.org/ver10/device/wsdl RemoteUser"`
}

type SetRemoteUserResponse struct {
}

type GetUsers struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetUsers"`
}

type GetUsersResponse struct {
	User []onvif.User
}

type CreateUsers struct {
	XMLName string       `xml:"http://www.onvif.org/ver10/device/wsdl CreateUsers"`
	Users   []onvif.User `xml:"http://www.onvif.org/ver10/device/wsdl User"`
}

type CreateUsersResponse struct {
}

type DeleteUsers struct {
	XMLName   string   `xml:"http://www.onvif.org/ver10/device/wsdl DeleteUsers"`
	Usernames []string `xml:"http://www.onvif.org/ver10/device/wsdl Username"`
}

type DeleteUsersResponse struct {
}

type SetUser struct {
	XMLName string       `xml:"http://www.onvif.org/ver10/device/wsdl SetUser"`
	Users   []onvif.User `xml:"http://www.onvif.org/ver10/device/wsdl User"`
}

type SetUserResponse struct {
}

type GetWsdlUrl struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetWsdlUrl"`
}

type GetWsdlUrlResponse struct {
	WsdlURL xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl WsdlUrl"`
}

type GetCapabilities struct {
	XMLName  string                     `xml:"http://www.onvif.org/ver10/device/wsdl GetCapabilities"`
	Category []onvif.CapabilityCategory `xml:"http://www.onvif.org/ver10/device/wsdl Category"`
}

type GetCapabilitiesResponse struct {
	Capabilities onvif.Capabilities
}

type GetHostname struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetHostname"`
}

type GetHostnameResponse struct {
	HostnameInformation onvif.HostnameInformation
}

type SetHostname struct {
	XMLName string    `xml:"http://www.onvif.org/ver10/device/wsdl SetHostname"`
	Name    xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl Name"`
}

type SetHostnameResponse struct {
}

type SetHostnameFromDHCP struct {
	XMLName  string `xml:"http://www.onvif.org/ver10/device/wsdl SetHostnameFromDHCP"`
	FromDHCP bool   `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
}

type SetHostnameFromDHCPResponse struct {
	RebootNeeded bool
}

type GetDNS struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDNS"`
}

type GetDNSResponse struct {
	DNSInformation onvif.DNSInformation
}

type SetDNS struct {
	XMLName      string            `xml:"http://www.onvif.org/ver10/device/wsdl SetDNS"`
	FromDHCP     bool              `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
	SearchDomain []xsd.Token       `xml:"http://www.onvif.org/ver10/device/wsdl SearchDomain"`
	DNSManual    []onvif.IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl DNSManual"`
}

type SetDNSResponse struct {
}

type GetNTP struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNTP"`
}

type GetNTPResponse struct {
	NTPInformation onvif.NTPInformation
}

type SetNTP struct {
	XMLName   string              `xml:"http://www.onvif.org/ver10/device/wsdl SetNTP"`
	FromDHCP  bool                `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
	NTPManual []onvif.NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl NTPManual"`
}

type SetNTPResponse struct {
}

type GetDynamicDNS struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDynamicDNS"`
}

type GetDynamicDNSResponse struct {
	DynamicDNSInformation onvif.DynamicDNSInformation
}

type SetDynamicDNS struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl SetDynamicDNS"`
	Type    onvif.DynamicDNSType `xml:"http://www.onvif.org/ver10/device/wsdl Type"`
	Name    *onvif.DNSName       `xml:"http://www.onvif.org/ver10/device/wsdl Name"`
	TTL     *xsd.Duration        `xml:"http://www.onvif.org/ver10/device/wsdl TTL"`
}

type SetDynamicDNSResponse struct {
}

type GetNetworkInterfaces struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkInterfaces"`
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces []onvif.NetworkInterface
}

type SetNetworkInterfaces struct {
	XMLName          string                                 `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkInterfaces"`
	InterfaceToken   onvif.ReferenceToken                   `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
	NetworkInterface onvif.NetworkInterfaceSetConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl NetworkInterface"`
}

type SetNetworkInterfacesResponse struct {
	RebootNeeded bool
}

type GetNetworkProtocols struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkProtocols"`
}

type GetNetworkProtocolsResponse struct {
	NetworkProtocols []onvif.NetworkProtocol
}

type SetNetworkProtocols struct {
	XMLName          string                  `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkProtocols"`
	NetworkProtocols []onvif.NetworkProtocol `xml:"http://www.onvif.org/ver10/device/wsdl NetworkProtocols"`
}

type SetNetworkProtocolsResponse struct {
}

type GetNetworkDefaultGateway struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkDefaultGateway"`
}

type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway onvif.NetworkGateway
}

type SetNetworkDefaultGateway struct {
	XMLName     string              `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkDefaultGateway"`
	IPv4Address []onvif.IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address"`
	IPv6Address []onvif.IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address"`
}

type SetNetworkDefaultGatewayResponse struct {
}

type GetZeroConfiguration struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetZeroConfiguration"`
}

type GetZeroConfigurationResponse struct {
	ZeroConfiguration onvif.NetworkZeroConfiguration
}

type SetZeroConfiguration struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl SetZeroConfiguration"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
	Enabled        bool                 `xml:"http://www.onvif.org/ver10/device/wsdl Enabled"`
}

type SetZeroConfigurationResponse struct {
}

type GetIPAddressFilter struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetIPAddressFilter"`
}

type GetIPAddressFilterResponse struct {
	IPAddressFilter onvif.IPAddressFilter
}

type SetIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl SetIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter"`
}

type SetIPAddressFilterResponse struct {
}

type AddIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl AddIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter"`
}

type AddIPAddressFilterResponse struct {
}

type RemoveIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl RemoveIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/schema onvif:IPAddressFilter"`
}

type RemoveIPAddressFilterResponse struct {
}

type GetAccessPolicy struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetAccessPolicy"`
}

type GetAccessPolicyResponse struct {
	PolicyFile onvif.BinaryData
}

type SetAccessPolicy struct {
	XMLName    string           `xml:"http://www.onvif.org/ver10/device/wsdl SetAccessPolicy"`
	PolicyFile onvif.BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl PolicyFile"`
}

type SetAccessPolicyResponse struct {
}

type CreateCertificate struct {
	XMLName        string        `xml:"http://www.onvif.org/ver10/device/wsdl CreateCertificate"`
	CertificateID  *xsd.Token    `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`
	Subject        *string       `xml:"http://www.onvif.org/ver10/device/wsdl Subject,omitempty"`
	ValidNotBefore *xsd.DateTime `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotBefore,omitempty"`
	ValidNotAfter  *xsd.DateTime `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotAfter,omitempty"`
}

type CreateCertificateResponse struct {
	NvtCertificate onvif.Certificate
}

type GetCertificates struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificates"`
}

type GetCertificatesResponse struct {
	NvtCertificate []onvif.Certificate `xml:"http://www.onvif.org/ver10/device/wsdl NvtCertificate"`
}

type GetCertificatesStatus struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificatesStatus"`
}

type GetCertificatesStatusResponse struct {
	CertificateStatus []onvif.CertificateStatus `xml:"http://www.onvif.org/ver10/device/wsdl CertificateStatus"`
}

type SetCertificatesStatus struct {
	XMLName           string                    `xml:"http://www.onvif.org/ver10/device/wsdl SetCertificatesStatus"`
	CertificateStatus []onvif.CertificateStatus `xml:"http://www.onvif.org/ver10/device/wsdl CertificateStatus"`
}

type SetCertificatesStatusResponse struct {
}

type DeleteCertificates struct {
	XMLName       string      `xml:"http://www.onvif.org/ver10/device/wsdl DeleteCertificates"`
	CertificateID []xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
}

type DeleteCertificatesResponse struct {
}

type GetPkcs10Request struct {
	XMLName       string            `xml:"http://www.onvif.org/ver10/device/wsdl GetPkcs10Request"`
	CertificateID xsd.Token         `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
	Subject       *string           `xml:"http://www.onvif.org/ver10/device/wsdl Subject"`
	Attributes    *onvif.BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl Attributes"`
}

type GetPkcs10RequestResponse struct {
	Pkcs10Request onvif.BinaryData
}

type LoadCertificates struct {
	XMLName        string              `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificates"`
	NVTCertificate []onvif.Certificate `xml:"http://www.onvif.org/ver10/device/wsdl NVTCertificate"`
}

type LoadCertificatesResponse struct {
}

type GetClientCertificateMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetClientCertificateMode"`
}

type GetClientCertificateModeResponse struct {
	Enabled bool
}

type SetClientCertificateMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl SetClientCertificateMode"`
	Enabled bool   `xml:"http://www.onvif.org/ver10/device/wsdl Enabled"`
}

type SetClientCertificateModeResponse struct {
}

type GetRelayOutputs struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRelayOutputs"`
}

type GetRelayOutputsResponse struct {
	RelayOutputs []onvif.RelayOutput `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputs"`
}

type SetRelayOutputSettings struct {
	XMLName          string                    `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputSettings"`
	RelayOutputToken onvif.ReferenceToken      `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken"`
	Properties       onvif.RelayOutputSettings `xml:"http://www.onvif.org/ver10/device/wsdl Properties"`
}

type SetRelayOutputSettingsResponse struct {
}

type SetRelayOutputState struct {
	XMLName          string                  `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputState"`
	RelayOutputToken onvif.ReferenceToken    `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken"`
	LogicalState     onvif.RelayLogicalState `xml:"http://www.onvif.org/ver10/device/wsdl LogicalState"`
}

type SetRelayOutputStateResponse struct {
}

type SendAuxiliaryCommand struct {
	XMLName          string              `xml:"http://www.onvif.org/ver10/device/wsdl SendAuxiliaryCommand"`
	AuxiliaryCommand onvif.AuxiliaryData `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommand"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse *onvif.AuxiliaryData
}

type GetCACertificates struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCACertificates"`
}

type GetCACertificatesResponse struct {
	CACertificate []onvif.Certificate
}

type LoadCertificateWithPrivateKey struct {
	XMLName                   string                            `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificateWithPrivateKey"`
	CertificateWithPrivateKey []onvif.CertificateWithPrivateKey `xml:"http://www.onvif.org/ver10/device/wsdl CertificateWithPrivateKey"`
}

type LoadCertificateWithPrivateKeyResponse struct {
}

type GetCertificateInformation struct {
	XMLName       string    `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificateInformation"`
	CertificateID xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
}

type GetCertificateInformationResponse struct {
	CertificateInformation onvif.CertificateInformation
}

type LoadCACertificates struct {
	XMLName       string              `xml:"http://www.onvif.org/ver10/device/wsdl LoadCACertificates"`
	CACertificate []onvif.Certificate `xml:"http://www.onvif.org/ver10/device/wsdl CACertificate"`
}

type LoadCACertificatesResponse struct {
}

type CreateDot1XConfiguration struct {
	XMLName            string                   `xml:"http://www.onvif.org/ver10/device/wsdl CreateDot1XConfiguration"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration"`
}

type CreateDot1XConfigurationResponse struct {
}

type SetDot1XConfiguration struct {
	XMLName            string                   `xml:"http://www.onvif.org/ver10/device/wsdl SetDot1XConfiguration"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration"`
}

type SetDot1XConfigurationResponse struct {
}

type GetDot1XConfiguration struct {
	XMLName                 string               `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfiguration"`
	Dot1XConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken"`
}

type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration
}

type GetDot1XConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfigurations"`
}

type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration []onvif.Dot1XConfiguration
}

type DeleteDot1XConfiguration struct {
	XMLName                 string                 `xml:"http://www.onvif.org/ver10/device/wsdl DeleteDot1XConfiguration"`
	Dot1XConfigurationToken []onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken"`
}

type DeleteDot1XConfigurationResponse struct {
}

type GetDot11Capabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Capabilities"`
}

type GetDot11CapabilitiesResponse struct {
	Capabilities onvif.Dot11Capabilities
}

type GetDot11Status struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Status"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
}

type GetDot11StatusResponse struct {
	Status onvif.Dot11Status
}

type ScanAvailableDot11Networks struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl ScanAvailableDot11Networks"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
}

type ScanAvailableDot11NetworksResponse struct {
	Networks []onvif.Dot11AvailableNetworks `xml:"http://www.onvif.org/ver10/device/wsdl Networks"`
}

type GetSystemUris struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemUris"`
}

type GetSystemUrisResponse struct {
	SystemLogUris   *onvif.SystemLogUriList `xml:"http://www.onvif.org/ver10/device/wsdl SystemLogUris"`
	SupportInfoURI  *xsd.AnyURI             `xml:"http://www.onvif.org/ver10/device/wsdl SupportInfoUri"`
	SystemBackupURI *xsd.AnyURI             `xml:"http://www.onvif.org/ver10/device/wsdl SystemBackupUri"`
	Extension       *xsd.AnyType            `xml:"http://www.onvif.org/ver10/device/wsdl Extension"`
}

type StartFirmwareUpgrade struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl StartFirmwareUpgrade"`
}

type StartFirmwareUpgradeResponse struct {
	UploadUri        xsd.AnyURI
	UploadDelay      xsd.Duration
	ExpectedDownTime xsd.Duration
}

type StartSystemRestore struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl StartSystemRestore"`
}

type StartSystemRestoreResponse struct {
	UploadUri        xsd.AnyURI
	ExpectedDownTime xsd.Duration
}

type GetStorageConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfigurations"`
}

type GetStorageConfigurationsResponse struct {
	StorageConfigurations []StorageConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfigurations"`
}

type CreateStorageConfiguration struct {
	XMLName              string `xml:"http://www.onvif.org/ver10/device/wsdl CreateStorageConfiguration"`
	StorageConfiguration StorageConfigurationData
}

type CreateStorageConfigurationResponse struct {
	Token onvif.ReferenceToken
}

type GetStorageConfiguration struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfiguration"`
	Token   onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token"`
}

type GetStorageConfigurationResponse struct {
	StorageConfiguration StorageConfiguration
}

type SetStorageConfiguration struct {
	XMLName              string               `xml:"http://www.onvif.org/ver10/device/wsdl SetStorageConfiguration"`
	StorageConfiguration StorageConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfiguration"`
}

type SetStorageConfigurationResponse struct {
}

type DeleteStorageConfiguration struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl DeleteStorageConfiguration"`
	Token   onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token"`
}

type DeleteStorageConfigurationResponse struct {
}

type GetGeoLocation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetGeoLocation"`
}

type GetGeoLocationResponse struct {
	Location []onvif.LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location"`
}

type SetGeoLocation struct {
	XMLName  string                 `xml:"http://www.onvif.org/ver10/device/wsdl SetGeoLocation"`
	Location []onvif.LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location"`
}

type SetGeoLocationResponse struct {
}

type DeleteGeoLocation struct {
	XMLName  string                 `xml:"http://www.onvif.org/ver10/device/wsdl DeleteGeoLocation"`
	Location []onvif.LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location"`
}

type DeleteGeoLocationResponse struct {
}
