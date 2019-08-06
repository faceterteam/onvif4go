package onvif

import (
	"net"
	"time"

	"github.com/atagirov/onvif4go/xsd"
)

type ContentType string // minLength value="3"
type DNSName xsd.Token

type DeviceEntity struct {
	Token ReferenceToken `xml:"token,attr"`
}

type ReferenceToken xsd.String

type Name xsd.String

type IntRectangle struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type IntRectangleRange struct {
	XRange      IntRange
	YRange      IntRange
	WidthRange  IntRange
	HeightRange IntRange
}

type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64 `xml:"http://www.onvif.org/ver10/schema Min"`
	Max float64 `xml:"http://www.onvif.org/ver10/schema Max"`
}

type OSDConfiguration struct {
	DeviceEntity                  `xml:"token,attr"`
	VideoSourceConfigurationToken OSDReference               `xml:"http://www.onvif.org/ver10/schema VideoSourceConfigurationToken"`
	Type                          OSDType                    `xml:"http://www.onvif.org/ver10/schema Type"`
	Position                      OSDPosConfiguration        `xml:"http://www.onvif.org/ver10/schema Position"`
	TextString                    *OSDTextConfiguration      `xml:"http://www.onvif.org/ver10/schema TextString"`
	Image                         *OSDImgConfiguration       `xml:"http://www.onvif.org/ver10/schema Image"`
	Extension                     *OSDConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type OSDType xsd.String

type OSDPosConfiguration struct {
	Type      string                        `xml:"http://www.onvif.org/ver10/schema Type"`
	Pos       *Vector                       `xml:"http://www.onvif.org/ver10/schema Pos"`
	Extension *OSDPosConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type Vector struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type OSDPosConfigurationExtension xsd.AnyType

type OSDReference ReferenceToken

type OSDTextConfiguration struct {
	IsPersistentText bool `xml:"IsPersistentText,attr"`

	Type            string                         `xml:"http://www.onvif.org/ver10/schema Type"`
	DateFormat      string                         `xml:"http://www.onvif.org/ver10/schema DateFormat,omitempty"`
	TimeFormat      string                         `xml:"http://www.onvif.org/ver10/schema TimeFormat,omitempty"`
	FontSize        int                            `xml:"http://www.onvif.org/ver10/schema FontSize,omitempty"`
	FontColor       *OSDColor                      `xml:"http://www.onvif.org/ver10/schema FontColor"`
	BackgroundColor *OSDColor                      `xml:"http://www.onvif.org/ver10/schema BackgroundColor"`
	PlainText       string                         `xml:"http://www.onvif.org/ver10/schema PlainText,omitempty"`
	Extension       *OSDTextConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type OSDColor struct {
	Transparent int `xml:"Transparent,attr"`

	Color Color `xml:"http://www.onvif.org/ver10/schema Color"`
}

type Color struct {
	X          float64    `xml:"X,attr"`
	Y          float64    `xml:"Y,attr"`
	Z          float64    `xml:"Z,attr"`
	Colorspace xsd.AnyURI `xml:"Colorspace,attr"`
}

type OSDTextConfigurationExtension xsd.AnyType

type OSDImgConfiguration struct {
	ImgPath   xsd.AnyURI                    `xml:"http://www.onvif.org/ver10/schema ImgPath"`
	Extension *OSDImgConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type OSDImgConfigurationExtension xsd.AnyType

type OSDConfigurationExtension xsd.AnyType

type VideoSource struct {
	DeviceEntity
	Framerate  float64
	Resolution VideoResolution
	Imaging    *ImagingSettings
	Extension  *VideoSourceExtension
}

type VideoResolution struct {
	Width  int `xml:"http://www.onvif.org/ver10/schema Width"`
	Height int `xml:"http://www.onvif.org/ver10/schema Height"`
}

type ImagingSettings struct {
	BacklightCompensation BacklightCompensation
	Brightness            float64
	ColorSaturation       float64
	Contrast              float64
	Exposure              Exposure
	Focus                 FocusConfiguration
	IrCutFilter           IrCutFilterMode
	Sharpness             float64
	WideDynamicRange      WideDynamicRange
	WhiteBalance          WhiteBalance
	Extension             ImagingSettingsExtension
}

type BacklightCompensation struct {
	Mode  BacklightCompensationMode
	Level float64
}

type BacklightCompensationMode xsd.String

type Exposure struct {
	Mode            ExposureMode
	Priority        ExposurePriority
	Window          Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain         float64
	MaxGain         float64
	MinIris         float64
	MaxIris         float64
	ExposureTime    float64
	Gain            float64
	Iris            float64
}

type ExposureMode xsd.String

type ExposurePriority xsd.String

type Rectangle struct {
	Bottom float64 `xml:"bottom,attr"`
	Top    float64 `xml:"top,attr"`
	Right  float64 `xml:"right,attr"`
	Left   float64 `xml:"left,attr"`
}

type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed  float64
	NearLimit     float64
	FarLimit      float64
}

type AutoFocusMode xsd.String

type IrCutFilterMode xsd.String

type WideDynamicRange struct {
	Mode  WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode"`
	Level float64         `xml:"http://www.onvif.org/ver10/schema Level"`
}

type WideDynamicMode xsd.String

type WhiteBalance struct {
	Mode   WhiteBalanceMode
	CrGain float64
	CbGain float64
}

type WhiteBalanceMode xsd.String

type ImagingSettingsExtension xsd.AnyType

type VideoSourceExtension struct {
	Imaging   ImagingSettings20
	Extension VideoSourceExtension2
}

type ImagingSettings20 struct {
	BacklightCompensation BacklightCompensation20    `xml:"http://www.onvif.org/ver10/schema BacklightCompensation"`
	Brightness            float64                    `xml:"http://www.onvif.org/ver10/schema Brightness"`
	ColorSaturation       float64                    `xml:"http://www.onvif.org/ver10/schema ColorSaturation"`
	Contrast              float64                    `xml:"http://www.onvif.org/ver10/schema Contrast"`
	Exposure              Exposure20                 `xml:"http://www.onvif.org/ver10/schema Exposure"`
	Focus                 FocusConfiguration20       `xml:"http://www.onvif.org/ver10/schema Focus"`
	IrCutFilter           IrCutFilterMode            `xml:"http://www.onvif.org/ver10/schema IrCutFilter"`
	Sharpness             float64                    `xml:"http://www.onvif.org/ver10/schema Sharpness"`
	WideDynamicRange      WideDynamicRange20         `xml:"http://www.onvif.org/ver10/schema WideDynamicRange"`
	WhiteBalance          WhiteBalance20             `xml:"http://www.onvif.org/ver10/schema WhiteBalance"`
	Extension             ImagingSettingsExtension20 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type BacklightCompensation20 struct {
	Mode  BacklightCompensationMode `xml:"http://www.onvif.org/ver10/schema Mode"`
	Level float64                   `xml:"http://www.onvif.org/ver10/schema Level"`
}

type Exposure20 struct {
	Mode            ExposureMode     `xml:"http://www.onvif.org/ver10/schema Mode"`
	Priority        ExposurePriority `xml:"http://www.onvif.org/ver10/schema Priority"`
	Window          Rectangle        `xml:"http://www.onvif.org/ver10/schema Window"`
	MinExposureTime float64          `xml:"http://www.onvif.org/ver10/schema MinExposureTime"`
	MaxExposureTime float64          `xml:"http://www.onvif.org/ver10/schema MaxExposureTime"`
	MinGain         float64          `xml:"http://www.onvif.org/ver10/schema MinGain"`
	MaxGain         float64          `xml:"http://www.onvif.org/ver10/schema MaxGain"`
	MinIris         float64          `xml:"http://www.onvif.org/ver10/schema MinIris"`
	MaxIris         float64          `xml:"http://www.onvif.org/ver10/schema MaxIris"`
	ExposureTime    float64          `xml:"http://www.onvif.org/ver10/schema ExposureTime"`
	Gain            float64          `xml:"http://www.onvif.org/ver10/schema Gain"`
	Iris            float64          `xml:"http://www.onvif.org/ver10/schema Iris"`
}

type FocusConfiguration20 struct {
	AutoFocusMode AutoFocusMode                 `xml:"http://www.onvif.org/ver10/schema AutoFocusMode"`
	DefaultSpeed  float64                       `xml:"http://www.onvif.org/ver10/schema DefaultSpeed"`
	NearLimit     float64                       `xml:"http://www.onvif.org/ver10/schema NearLimit"`
	FarLimit      float64                       `xml:"http://www.onvif.org/ver10/schema FarLimit"`
	Extension     FocusConfiguration20Extension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type FocusConfiguration20Extension xsd.AnyType

type WideDynamicRange20 struct {
	Mode  WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode"`
	Level float64         `xml:"http://www.onvif.org/ver10/schema Level"`
}

type WhiteBalance20 struct {
	Mode      WhiteBalanceMode        `xml:"http://www.onvif.org/ver10/schema Mode"`
	CrGain    float64                 `xml:"http://www.onvif.org/ver10/schema CrGain"`
	CbGain    float64                 `xml:"http://www.onvif.org/ver10/schema CbGain"`
	Extension WhiteBalance20Extension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type WhiteBalance20Extension xsd.AnyType

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization          `xml:"http://www.onvif.org/ver10/schema ImageStabilization"`
	Extension          ImagingSettingsExtension202 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type ImageStabilization struct {
	Mode      ImageStabilizationMode      `xml:"http://www.onvif.org/ver10/schema Mode"`
	Level     float64                     `xml:"http://www.onvif.org/ver10/schema Level"`
	Extension ImageStabilizationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type ImageStabilizationMode xsd.String

type ImageStabilizationExtension xsd.AnyType

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment   `xml:"http://www.onvif.org/ver10/schema IrCutFilterAutoAdjustment"`
	Extension                 ImagingSettingsExtension203 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string                             `xml:"http://www.onvif.org/ver10/schema BoundaryType"`
	BoundaryOffset float64                            `xml:"http://www.onvif.org/ver10/schema BoundaryOffset"`
	ResponseTime   xsd.Duration                       `xml:"http://www.onvif.org/ver10/schema ResponseTime"`
	Extension      IrCutFilterAutoAdjustmentExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type IrCutFilterAutoAdjustmentExtension xsd.AnyType

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation            `xml:"http://www.onvif.org/ver10/schema ToneCompensation"`
	Defogging        Defogging                   `xml:"http://www.onvif.org/ver10/schema Defogging"`
	NoiseReduction   NoiseReduction              `xml:"http://www.onvif.org/ver10/schema NoiseReduction"`
	Extension        ImagingSettingsExtension204 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type ToneCompensation struct {
	Mode      string                    `xml:"http://www.onvif.org/ver10/schema Mode"`
	Level     float64                   `xml:"http://www.onvif.org/ver10/schema Level"`
	Extension ToneCompensationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type ToneCompensationExtension xsd.AnyType

type Defogging struct {
	Mode      string
	Level     float64
	Extension DefoggingExtension
}

type DefoggingExtension xsd.AnyType

type NoiseReduction struct {
	Level float64 `xml:"http://www.onvif.org/ver10/schema Level"`
}

type ImagingSettingsExtension204 xsd.AnyType

type VideoSourceExtension2 xsd.AnyType

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token                       ReferenceToken `xml:"token,attr"`
	Fixed                       bool           `xml:"fixed,attr"`
	Name                        Name
	VideoSourceConfiguration    *VideoSourceConfiguration
	AudioSourceConfiguration    *AudioSourceConfiguration
	VideoEncoderConfiguration   *VideoEncoderConfiguration
	AudioEncoderConfiguration   *AudioEncoderConfiguration
	VideoAnalyticsConfiguration *VideoAnalyticsConfiguration
	PTZConfiguration            *PTZConfiguration
	MetadataConfiguration       *MetadataConfiguration
	Extension                   *ProfileExtension
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode    string                             `xml:"ViewMode,attr"`
	SourceToken ReferenceToken                     `xml:"http://www.onvif.org/ver10/schema SourceToken"`
	Bounds      IntRectangle                       `xml:"http://www.onvif.org/ver10/schema Bounds"`
	Extension   *VideoSourceConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type ConfigurationEntity struct {
	Token    ReferenceToken `xml:"token,attr"`
	Name     Name           `xml:"http://www.onvif.org/ver10/schema Name"`
	UseCount int            `xml:"http://www.onvif.org/ver10/schema UseCount"`
}

type VideoSourceConfigurationExtension struct {
	Rotate    Rotate                             `xml:"http://www.onvif.org/ver10/schema Rotate"`
	Extension VideoSourceConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type Rotate struct {
	Mode      RotateMode      `xml:"http://www.onvif.org/ver10/schema Mode"`
	Degree    int             `xml:"http://www.onvif.org/ver10/schema Degree"`
	Extension RotateExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type RotateMode xsd.String

type RotateExtension xsd.AnyType

type VideoSourceConfigurationExtension2 struct {
	LensDescription  LensDescription  `xml:"http://www.onvif.org/ver10/schema LensDescription"`
	SceneOrientation SceneOrientation `xml:"http://www.onvif.org/ver10/schema SceneOrientation"`
}

type LensDescription struct {
	FocalLength float64        `xml:"FocalLength,attr"`
	Offset      LensOffset     `xml:"http://www.onvif.org/ver10/schema Offset"`
	Projection  LensProjection `xml:"http://www.onvif.org/ver10/schema Projection"`
	XFactor     float64        `xml:"http://www.onvif.org/ver10/schema XFactor"`
}

type LensOffset struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type LensProjection struct {
	Angle         float64 `xml:"http://www.onvif.org/ver10/schema Angle"`
	Radius        float64 `xml:"http://www.onvif.org/ver10/schema Radius"`
	Transmittance float64 `xml:"http://www.onvif.org/ver10/schema Transmittance"`
}

type SceneOrientation struct {
	Mode        SceneOrientationMode `xml:"http://www.onvif.org/ver10/schema Mode"`
	Orientation xsd.String           `xml:"http://www.onvif.org/ver10/schema Orientation"`
}

type SceneOrientationMode xsd.String

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema SourceToken"`
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       VideoEncoding          `xml:"http://www.onvif.org/ver10/schema Encoding"`
	Resolution     VideoResolution        `xml:"http://www.onvif.org/ver10/schema Resolution"`
	Quality        float64                `xml:"http://www.onvif.org/ver10/schema Quality"`
	RateControl    *VideoRateControl      `xml:"http://www.onvif.org/ver10/schema RateControl"`
	MPEG4          *Mpeg4Configuration    `xml:"http://www.onvif.org/ver10/schema MPEG4"`
	H264           *H264Configuration     `xml:"http://www.onvif.org/ver10/schema H264"`
	Multicast      MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast"`
	SessionTimeout xsd.Duration           `xml:"http://www.onvif.org/ver10/schema SessionTimeout"`
}

type VideoEncoding xsd.String

type VideoRateControl struct {
	FrameRateLimit   int `xml:"http://www.onvif.org/ver10/schema FrameRateLimit"`
	EncodingInterval int `xml:"http://www.onvif.org/ver10/schema EncodingInterval"`
	BitrateLimit     int `xml:"http://www.onvif.org/ver10/schema BitrateLimit"`
}

type Mpeg4Configuration struct {
	GovLength    int          `xml:"http://www.onvif.org/ver10/schema GovLength"`
	Mpeg4Profile Mpeg4Profile `xml:"http://www.onvif.org/ver10/schema Mpeg4Profile"`
}

type Mpeg4Profile xsd.String

type H264Configuration struct {
	GovLength   int         `xml:"http://www.onvif.org/ver10/schema GovLength"`
	H264Profile H264Profile `xml:"http://www.onvif.org/ver10/schema H264Profile"`
}

type H264Profile xsd.String

type MulticastConfiguration struct {
	Address   IPAddress `xml:"http://www.onvif.org/ver10/schema Address"`
	Port      int       `xml:"http://www.onvif.org/ver10/schema Port"`
	TTL       int       `xml:"http://www.onvif.org/ver10/schema TTL"`
	AutoStart bool      `xml:"http://www.onvif.org/ver10/schema AutoStart"`
}

type IPAddress struct {
	Type        IPType      `xml:"http://www.onvif.org/ver10/schema Type"`
	IPv4Address IPv4Address `xml:"http://www.onvif.org/ver10/schema IPv4Address"`
	IPv6Address IPv6Address `xml:"http://www.onvif.org/ver10/schema IPv6Address"`
}

func NewIPAddress(ip net.IP) (res IPAddress, err error) {
	ns, err := xsd.NewNormalizedString(ip.String())
	if err != nil {
		return
	}

	token, err := xsd.NewToken(ns)
	if err != nil {
		return
	}

	if ip.To4() == nil {
		res.Type = IPType("IPv6")
		res.IPv6Address = IPv6Address(token)
	} else {
		res.Type = IPType("IPv4")
		res.IPv4Address = IPv4Address(token)
	}

	return
}

type IPType xsd.String

//IPv4 address
type IPv4Address xsd.Token

//IPv6 address
type IPv6Address xsd.Token

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       AudioEncoding          `xml:"http://www.onvif.org/ver10/schema Encoding"`
	Bitrate        int                    `xml:"http://www.onvif.org/ver10/schema Bitrate"`
	SampleRate     int                    `xml:"http://www.onvif.org/ver10/schema SampleRate"`
	Multicast      MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast"`
	SessionTimeout xsd.Duration           `xml:"http://www.onvif.org/ver10/schema SessionTimeout"`
}

type AudioEncoding xsd.String

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration"`
	RuleEngineConfiguration      RuleEngineConfiguration      `xml:"http://www.onvif.org/ver10/schema RuleEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule []Config                               `xml:"http://www.onvif.org/ver10/schema AnalyticsModule"`
	Extension       *AnalyticsEngineConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type Config struct {
	Name       string    `xml:"Name,attr"`
	Type       xsd.QName `xml:"Type,attr"`
	Parameters ItemList  `xml:"http://www.onvif.org/ver10/schema Parameters"`
}

type AnalyticsEngineConfigurationExtension xsd.AnyType

type RuleEngineConfiguration struct {
	Rule      Config                           `xml:"http://www.onvif.org/ver10/schema Rule"`
	Extension RuleEngineConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type RuleEngineConfigurationExtension xsd.AnyType

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int                       `xml:"MoveRamp,attr"`
	PresetRamp                             int                       `xml:"PresetRamp,attr"`
	PresetTourRamp                         int                       `xml:"PresetTourRamp,attr"`
	NodeToken                              ReferenceToken            `xml:"http://www.onvif.org/ver10/schema NodeToken"`
	DefaultAbsolutePantTiltPositionSpace   xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultAbsolutePantTiltPositionSpace"`
	DefaultAbsoluteZoomPositionSpace       xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultAbsoluteZoomPositionSpace"`
	DefaultRelativePanTiltTranslationSpace xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultRelativePanTiltTranslationSpace"`
	DefaultRelativeZoomTranslationSpace    xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultRelativeZoomTranslationSpace"`
	DefaultContinuousPanTiltVelocitySpace  xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultContinuousPanTiltVelocitySpace"`
	DefaultContinuousZoomVelocitySpace     xsd.AnyURI                `xml:"http://www.onvif.org/ver10/schema DefaultContinuousZoomVelocitySpace"`
	DefaultPTZSpeed                        PTZSpeed                  `xml:"http://www.onvif.org/ver10/schema DefaultPTZSpeed"`
	DefaultPTZTimeout                      xsd.Duration              `xml:"http://www.onvif.org/ver10/schema DefaultPTZTimeout"`
	PanTiltLimits                          PanTiltLimits             `xml:"http://www.onvif.org/ver10/schema PanTiltLimits"`
	ZoomLimits                             ZoomLimits                `xml:"http://www.onvif.org/ver10/schema ZoomLimits"`
	Extension                              PTZConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZSpeed struct {
	PanTilt Vector2D `xml:"http://www.onvif.org/ver10/schema PanTilt"`
	Zoom    Vector1D `xml:"http://www.onvif.org/ver10/schema Zoom"`
}

type Vector2D struct {
	X     float64    `xml:"x,attr"`
	Y     float64    `xml:"y,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type Vector1D struct {
	X     float64    `xml:"x,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type PanTiltLimits struct {
	Range Space2DDescription `xml:"http://www.onvif.org/ver10/schema Range"`
}

type Space2DDescription struct {
	URI    xsd.AnyURI `xml:"http://www.onvif.org/ver10/schema URI"`
	XRange FloatRange `xml:"http://www.onvif.org/ver10/schema XRange"`
	YRange FloatRange `xml:"http://www.onvif.org/ver10/schema YRange"`
}

type ZoomLimits struct {
	Range Space1DDescription `xml:"http://www.onvif.org/ver10/schema Range"`
}

type Space1DDescription struct {
	URI    xsd.AnyURI `xml:"http://www.onvif.org/ver10/schema URI"`
	XRange FloatRange `xml:"http://www.onvif.org/ver10/schema XRange"`
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection         `xml:"http://www.onvif.org/ver10/schema PTControlDirection"`
	Extension          PTZConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTControlDirection struct {
	EFlip     EFlip                       `xml:"http://www.onvif.org/ver10/schema EFlip"`
	Reverse   Reverse                     `xml:"http://www.onvif.org/ver10/schema Reverse"`
	Extension PTControlDirectionExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type EFlip struct {
	Mode EFlipMode `xml:"http://www.onvif.org/ver10/schema Mode"`
}

type EFlipMode xsd.String

type Reverse struct {
	Mode ReverseMode `xml:"http://www.onvif.org/ver10/schema Mode"`
}

type ReverseMode xsd.String

type PTControlDirectionExtension xsd.AnyType

type PTZConfigurationExtension2 xsd.AnyType

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string                          `xml:"CompressionType,attr"`
	PTZStatus                    *PTZFilter                      `xml:"http://www.onvif.org/ver10/schema PTZStatus"`
	Events                       *EventSubscription              `xml:"http://www.onvif.org/ver10/schema Events"`
	Analytics                    bool                            `xml:"http://www.onvif.org/ver10/schema Analytics"`
	Multicast                    MulticastConfiguration          `xml:"http://www.onvif.org/ver10/schema Multicast"`
	SessionTimeout               xsd.Duration                    `xml:"http://www.onvif.org/ver10/schema SessionTimeout"`
	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration   `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration"`
	Extension                    *MetadataConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZFilter struct {
	Status   bool `xml:"http://www.onvif.org/ver10/schema Status"`
	Position bool `xml:"http://www.onvif.org/ver10/schema Position"`
}

type EventSubscription struct {
	Filter             FilterType `xml:"http://www.onvif.org/ver10/schema Filter"`
	SubscriptionPolicy `xml:"http://www.onvif.org/ver10/schema SubscriptionPolicy"`
}

type FilterType xsd.AnyType

type SubscriptionPolicy xsd.AnyType

type MetadataConfigurationExtension xsd.AnyType

type ProfileExtension struct {
	AudioOutputConfiguration  *AudioOutputConfiguration
	AudioDecoderConfiguration *AudioDecoderConfiguration
	Extension                 *ProfileExtension2
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema OutputToken"`
	SendPrimacy xsd.AnyURI     `xml:"http://www.onvif.org/ver10/schema SendPrimacy"`
	OutputLevel int            `xml:"http://www.onvif.org/ver10/schema OutputLevel"`
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type ProfileExtension2 xsd.AnyType

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles    int `xml:"MaximumNumberOfProfiles,attr"`
	BoundsRange                IntRectangleRange
	VideoSourceTokensAvailable ReferenceToken
	Extension                  VideoSourceConfigurationOptionsExtension
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    RotateOptions
	Extension VideoSourceConfigurationOptionsExtension2
}

type RotateOptions struct {
	Mode       RotateMode
	DegreeList IntList
	Extension  RotateOptionsExtension
}

type IntList struct {
	Items []int
}

type RotateOptionsExtension xsd.AnyType

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode SceneOrientationMode
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	JPEG         JpegOptions
	MPEG4        Mpeg4Options
	H264         H264Options
	Extension    VideoEncoderOptionsExtension
}

type JpegOptions struct {
	ResolutionsAvailable  VideoResolution
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
}

type Mpeg4Options struct {
	ResolutionsAvailable   VideoResolution
	GovLengthRange         IntRange
	FrameRateRange         IntRange
	EncodingIntervalRange  IntRange
	Mpeg4ProfilesSupported Mpeg4Profile
}

type H264Options struct {
	ResolutionsAvailable  VideoResolution
	GovLengthRange        IntRange
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
	H264ProfilesSupported H264Profile
}

type VideoEncoderOptionsExtension struct {
	JPEG      JpegOptions2
	MPEG4     Mpeg4Options2
	H264      H264Options2
	Extension VideoEncoderOptionsExtension2
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange
}

type VideoEncoderOptionsExtension2 xsd.AnyType

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable ReferenceToken
	Extension            AudioSourceOptionsExtension
}

type AudioSourceOptionsExtension xsd.AnyType

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption
}

type AudioEncoderConfigurationOption struct {
	Encoding       AudioEncoding
	BitrateList    IntList
	SampleRateList IntList
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions
	Extension              *MetadataConfigurationOptionsExtension
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   bool
	ZoomStatusSupported      bool
	PanTiltPositionSupported bool
	ZoomPositionSupported    bool
	Extension                *PTZStatusFilterOptionsExtension
}

type PTZStatusFilterOptionsExtension xsd.AnyType

type MetadataConfigurationOptionsExtension struct {
	CompressionType string
	Extension       MetadataConfigurationOptionsExtension2
}

type MetadataConfigurationOptionsExtension2 xsd.AnyType

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable ReferenceToken
	SendPrimacyOptions    xsd.AnyURI
	OutputLevelRange      IntRange
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  AACDecOptions
	G711DecOptions G711DecOptions
	G726DecOptions G726DecOptions
	Extension      AudioDecoderConfigurationOptionsExtension
}

type AACDecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G711DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G726DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type AudioDecoderConfigurationOptionsExtension xsd.AnyType

type StreamSetup struct {
	Stream    StreamType `xml:"http://www.onvif.org/ver10/schema Stream"`
	Transport Transport  `xml:"http://www.onvif.org/ver10/schema Transport"`
}

type StreamType xsd.String

type Transport struct {
	Protocol TransportProtocol `xml:"http://www.onvif.org/ver10/schema Protocol"`
	Tunnel   *Transport        `xml:"http://www.onvif.org/ver10/schema Tunnel"`
}

//enum
type TransportProtocol xsd.String

type MediaUri struct {
	Uri                 xsd.AnyURI
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             xsd.Duration
}

type VideoSourceMode struct {
	Token         ReferenceToken `xml:"token,attr"`
	Enabled       bool           `xml:"Enabled,attr"`
	MaxFramerate  float64
	MaxResolution VideoResolution
	Encodings     EncodingTypes
	Reboot        bool
	Description   *Description
	Extension     *VideoSourceModeExtension
}

type EncodingTypes struct {
	EncodingTypes []string
}

type Description struct {
	Description string
}

type VideoSourceModeExtension xsd.AnyType

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs
	Type                []OSDType
	PositionOption      []string
	TextOption          *OSDTextOptions
	ImageOption         *OSDImgOptions
	Extension           *OSDConfigurationOptionsExtension
}

type MaximumNumberOfOSDs struct {
	Total       int `xml:"Total,attr"`
	Image       int `xml:"Image,attr"`
	PlainText   int `xml:"PlainText,attr"`
	Date        int `xml:"Date,attr"`
	Time        int `xml:"Time,attr"`
	DateAndTime int `xml:"DateAndTime,attr"`
}

type OSDTextOptions struct {
	Type            []string
	FontSizeRange   *IntRange
	DateFormat      []string
	TimeFormat      []string
	FontColor       *OSDColorOptions
	BackgroundColor *OSDColorOptions
	Extension       *OSDTextOptionsExtension
}

type OSDColorOptions struct {
	Color       *ColorOptions
	Transparent *IntRange
	Extension   *OSDColorOptionsExtension
}

// ColorOptions describe the option of the color supported.
// Either list each color or define the range of color value.
// The following values are acceptable for Colourspace attribute.
//
//	* http://www.onvif.org/ver10/colorspace/YCbCr - YCbCr colourspace
//	* http://www.onvif.org/ver10/colorspace/CIELUV - CIE LUV
//	* http://www.onvif.org/ver10/colorspace/CIELAB - CIE 1976 (L*a*b*)
//	* http://www.onvif.org/ver10/colorspace/HSV - HSV colourspace
type ColorOptions struct {
	ColorList       []Color
	ColorspaceRange []ColorspaceRange
}

type ColorspaceRange struct {
	X          FloatRange
	Y          FloatRange
	Z          FloatRange
	Colorspace xsd.AnyURI
}

type OSDColorOptionsExtension xsd.AnyType

type OSDTextOptionsExtension xsd.AnyType

type OSDImgOptions struct {
	FormatsSupported StringAttrList `xml:"FormatsSupported,attr"`
	MaxSize          int            `xml:"MaxSize,attr"`
	MaxWidth         int            `xml:"MaxWidth,attr"`
	MaxHeight        int            `xml:"MaxHeight,attr"`

	ImagePath []xsd.AnyURI
	Extension *OSDImgOptionsExtension
}

//TODO: <xs:list itemType="xs:string"/>
type StringAttrList struct {
	AttrList []string
}

type OSDImgOptionsExtension xsd.AnyType

type OSDConfigurationOptionsExtension xsd.AnyType

//PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      bool `xml:"FixedHomePosition,attr"`
	GeoMove                bool `xml:"GeoMove,attr"`
	Name                   Name
	SupportedPTZSpaces     PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported          bool
	AuxiliaryCommands      AuxiliaryData
	Extension              PTZNodeExtension
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    Space2DDescription
	AbsoluteZoomPositionSpace       Space1DDescription
	RelativePanTiltTranslationSpace Space2DDescription
	RelativeZoomTranslationSpace    Space1DDescription
	ContinuousPanTiltVelocitySpace  Space2DDescription
	ContinuousZoomVelocitySpace     Space1DDescription
	PanTiltSpeedSpace               Space1DDescription
	ZoomSpeedSpace                  Space1DDescription
	Extension                       PTZSpacesExtension
}

type PTZSpacesExtension xsd.AnyType

//TODO: restriction
type AuxiliaryData xsd.String

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported
	Extension           PTZNodeExtension2
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int
	PTZPresetTourOperation     PTZPresetTourOperation
	Extension                  PTZPresetTourSupportedExtension
}

type PTZPresetTourOperation xsd.String
type PTZPresetTourSupportedExtension xsd.AnyType

type PTZNodeExtension2 xsd.AnyType

type PTZConfigurationOptions struct {
	PTZRamps           IntAttrList `xml:"PTZRamps,attr"`
	Spaces             PTZSpaces
	PTZTimeout         DurationRange
	PTControlDirection PTControlDirectionOptions
	Extension          PTZConfigurationOptions2
}

type IntAttrList struct {
	IntAttrList []int
}

type DurationRange struct {
	Min xsd.Duration
	Max xsd.Duration
}

type PTControlDirectionOptions struct {
	EFlip     EFlipOptions
	Reverse   ReverseOptions
	Extension PTControlDirectionOptionsExtension
}

type EFlipOptions struct {
	Mode      EFlipMode
	Extension EFlipOptionsExtension
}

type EFlipOptionsExtension xsd.AnyType

type ReverseOptions struct {
	Mode      ReverseMode
	Extension ReverseOptionsExtension
}

type ReverseOptionsExtension xsd.AnyType

type PTControlDirectionOptionsExtension xsd.AnyType

type PTZConfigurationOptions2 xsd.AnyType

type PTZPreset struct {
	Token       ReferenceToken `xml:"token,attr"`
	Name        Name
	PTZPosition PTZVector
}

type PTZVector struct {
	PanTilt Vector2D `xml:"http://www.onvif.org/ver10/schema PanTilt"`
	Zoom    Vector1D `xml:"http://www.onvif.org/ver10/schema Zoom"`
}

type PTZStatus struct {
	Position   PTZVector
	MoveStatus PTZMoveStatus
	Error      string
	UtcTime    xsd.DateTime
}

type PTZMoveStatus struct {
	PanTilt MoveStatus
	Zoom    MoveStatus
}

type MoveStatus struct {
	Status string
}

type GeoLocation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type PresetTour struct {
	Token             ReferenceToken                 `xml:"token,attr"`
	Name              Name                           `xml:"http://www.onvif.org/ver10/schema Name"`
	Status            PTZPresetTourStatus            `xml:"http://www.onvif.org/ver10/schema Status"`
	AutoStart         bool                           `xml:"http://www.onvif.org/ver10/schema AutoStart"`
	StartingCondition PTZPresetTourStartingCondition `xml:"http://www.onvif.org/ver10/schema StartingCondition"`
	TourSpot          PTZPresetTourSpot              `xml:"http://www.onvif.org/ver10/schema TourSpot"`
	Extension         PTZPresetTourExtension         `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZPresetTourStatus struct {
	State           PTZPresetTourState           `xml:"http://www.onvif.org/ver10/schema State"`
	CurrentTourSpot PTZPresetTourSpot            `xml:"http://www.onvif.org/ver10/schema CurrentTourSpot"`
	Extension       PTZPresetTourStatusExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZPresetTourState xsd.String

type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail  `xml:"http://www.onvif.org/ver10/schema PresetDetail"`
	Speed        PTZSpeed                   `xml:"http://www.onvif.org/ver10/schema Speed"`
	StayTime     xsd.Duration               `xml:"http://www.onvif.org/ver10/schema StayTime"`
	Extension    PTZPresetTourSpotExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZPresetTourPresetDetail struct {
	PresetToken   ReferenceToken             `xml:"http://www.onvif.org/ver10/schema PresetToken"`
	Home          bool                       `xml:"http://www.onvif.org/ver10/schema Home"`
	PTZPosition   PTZVector                  `xml:"http://www.onvif.org/ver10/schema PTZPosition"`
	TypeExtension PTZPresetTourTypeExtension `xml:"http://www.onvif.org/ver10/schema TypeExtension"`
}

type PTZPresetTourTypeExtension xsd.AnyType

type PTZPresetTourSpotExtension xsd.AnyType

type PTZPresetTourStatusExtension xsd.AnyType

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder bool                                    `xml:"RandomPresetOrder,attr"`
	RecurringTime     int                                     `xml:"http://www.onvif.org/ver10/schema RecurringTime"`
	RecurringDuration xsd.Duration                            `xml:"http://www.onvif.org/ver10/schema RecurringDuration"`
	Direction         PTZPresetTourDirection                  `xml:"http://www.onvif.org/ver10/schema Direction"`
	Extension         PTZPresetTourStartingConditionExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type PTZPresetTourDirection xsd.String

type PTZPresetTourStartingConditionExtension xsd.AnyType

type PTZPresetTourExtension xsd.AnyType

type PTZPresetTourOptions struct {
	AutoStart         bool
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot          PTZPresetTourSpotOptions
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     IntRange
	RecurringDuration DurationRange
	Direction         PTZPresetTourDirection
	Extension         PTZPresetTourStartingConditionOptionsExtension
}

type PTZPresetTourStartingConditionOptionsExtension xsd.AnyType

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions
	StayTime     DurationRange
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          ReferenceToken
	Home                 bool
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace    Space1DDescription
	Extension            PTZPresetTourPresetDetailOptionsExtension
}

type PTZPresetTourPresetDetailOptionsExtension xsd.AnyType

//Device

type OnvifVersion struct {
	Major int
	Minor int
}

type SetDateTimeType xsd.String

type TimeZone struct {
	TZ xsd.Token `xml:"http://www.onvif.org/ver10/schema TZ"`
}

type SystemDateTime struct {
	DateTimeType    SetDateTimeType
	DaylightSavings bool
	TimeZone        TimeZone
	UTCDateTime     *DateTime
	LocalDateTime   DateTime
	Extension       SystemDateTimeExtension
}

func (d *SystemDateTime) GetUTCTime() (time.Time, error) {
	if d.UTCDateTime != nil {
		return d.UTCDateTime.getTime(), nil
	}

	tz, err := ParsePosixTimezone(string(d.TimeZone.TZ))
	if err != nil {
		return time.Now().UTC(), err
	}

	return tz.LocalToUTC(d.LocalDateTime.getTime()), nil
}

func (d *DateTime) getTime() time.Time {
	return time.Date(
		d.Date.Year, time.Month(d.Date.Month), d.Date.Day,
		d.Time.Hour, d.Time.Minute, d.Time.Second, 0,
		time.UTC)
}

type SystemDateTimeExtension xsd.AnyType

type FactoryDefaultType xsd.String

type AttachmentData struct {
	ContentType ContentType `xml:"contentType,attr"`
	Include     Include     `xml:"inc:Include"`
}

type Include struct {
	Href xsd.AnyURI `xml:"href,attr"`
}

type BackupFile struct {
	Name string         `xml:"http://www.onvif.org/ver10/schema Name"`
	Data AttachmentData `xml:"http://www.onvif.org/ver10/schema Data"`
}

type SystemLogType xsd.String

type SystemLog struct {
	Binary AttachmentData
	String string
}

type SupportInformation struct {
	Binary AttachmentData
	String string
}

type Scope struct {
	ScopeDef  ScopeDefinition
	ScopeItem xsd.AnyURI
}

type ScopeDefinition xsd.String

type DiscoveryMode xsd.String

type NetworkHost struct {
	Type        NetworkHostType       `xml:"http://www.onvif.org/ver10/schema Type"`
	IPv4Address *IPv4Address          `xml:"http://www.onvif.org/ver10/schema IPv4Address"`
	IPv6Address *IPv6Address          `xml:"http://www.onvif.org/ver10/schema IPv6Address"`
	DNSname     *DNSName              `xml:"http://www.onvif.org/ver10/schema DNSname"`
	Extension   *NetworkHostExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type NetworkHostType xsd.String

type NetworkHostExtension xsd.String

type RemoteUser struct {
	Username           string `xml:"http://www.onvif.org/ver10/schema Username"`
	Password           string `xml:"http://www.onvif.org/ver10/schema Password"`
	UseDerivedPassword bool   `xml:"http://www.onvif.org/ver10/schema UseDerivedPassword"`
}

type User struct {
	Username  string         `xml:"http://www.onvif.org/ver10/schema Username"`
	Password  string         `xml:"http://www.onvif.org/ver10/schema Password"`
	UserLevel UserLevel      `xml:"http://www.onvif.org/ver10/schema UserLevel"`
	Extension *UserExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

func NewUser(username, password, level string) User {
	return User {
		Username: username,
		Password: password,
		UserLevel: UserLevel(level),
	}
}

type UserLevel xsd.String

type UserExtension xsd.String

type CapabilityCategory xsd.String

type Capabilities struct {
	Analytics *AnalyticsCapabilities
	Device    *DeviceCapabilities
	Events    *EventCapabilities
	Imaging   *ImagingCapabilities
	Media     *MediaCapabilities
	PTZ       *PTZCapabilities
	Extension *CapabilitiesExtension
}

type AnalyticsCapabilities struct {
	XAddr                  xsd.AnyURI
	RuleSupport            bool
	AnalyticsModuleSupport bool
}

type DeviceCapabilities struct {
	XAddr     xsd.AnyURI
	Network   NetworkCapabilities
	System    SystemCapabilities
	IO        IOCapabilities
	Security  SecurityCapabilities
	Extension DeviceCapabilitiesExtension
}

type NetworkCapabilities struct {
	IPFilter          bool
	ZeroConfiguration bool
	IPVersion6        bool
	DynDNS            bool
	Extension         NetworkCapabilitiesExtension
}

type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool
	Extension          NetworkCapabilitiesExtension2
}

type NetworkCapabilitiesExtension2 xsd.AnyType

type SystemCapabilities struct {
	DiscoveryResolve  bool
	DiscoveryBye      bool
	RemoteDiscovery   bool
	SystemBackup      bool
	SystemLogging     bool
	FirmwareUpgrade   bool
	SupportedVersions OnvifVersion
	Extension         SystemCapabilitiesExtension
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    bool
	HttpSystemBackup       bool
	HttpSystemLogging      bool
	HttpSupportInformation bool
	Extension              SystemCapabilitiesExtension2
}

type SystemCapabilitiesExtension2 xsd.AnyType

type IOCapabilities struct {
	InputConnectors int
	RelayOutputs    int
	Extension       IOCapabilitiesExtension
}

type IOCapabilitiesExtension struct {
	Auxiliary         bool
	AuxiliaryCommands AuxiliaryData
	Extension         IOCapabilitiesExtension2
}

type IOCapabilitiesExtension2 xsd.AnyType

type SecurityCapabilities struct {
	TLS1_1               bool
	TLS1_2               bool
	OnboardKeyGeneration bool
	AccessPolicyConfig   bool
	X_509Token           bool
	SAMLToken            bool
	KerberosToken        bool
	RELToken             bool
	Extension            SecurityCapabilitiesExtension
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    bool
	Extension SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              bool
	SupportedEAPMethod int
	RemoteUserHandling bool
}

type DeviceCapabilitiesExtension xsd.AnyType

type EventCapabilities struct {
	XAddr                                         xsd.AnyURI
	WSSubscriptionPolicySupport                   bool
	WSPullPointSupport                            bool
	WSPausableSubscriptionManagerInterfaceSupport bool
}

type ImagingCapabilities struct {
	XAddr xsd.AnyURI
}

type MediaCapabilities struct {
	XAddr                 xsd.AnyURI
	StreamingCapabilities RealTimeStreamingCapabilities
	Extension             MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast bool
	RTP_TCP      bool
	RTP_RTSP_TCP bool
	Extension    RealTimeStreamingCapabilitiesExtension
}

type RealTimeStreamingCapabilitiesExtension xsd.AnyType

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int
}

type PTZCapabilities struct {
	XAddr xsd.AnyURI
}

type CapabilitiesExtension struct {
	DeviceIO        *DeviceIOCapabilities
	Display         *DisplayCapabilities
	Recording       *RecordingCapabilities
	Search          *SearchCapabilities
	Replay          *ReplayCapabilities
	Receiver        *ReceiverCapabilities
	AnalyticsDevice *AnalyticsDeviceCapabilities
	Extensions      *CapabilitiesExtension2
}

type DeviceIOCapabilities struct {
	XAddr        xsd.AnyURI
	VideoSources int
	VideoOutputs int
	AudioSources int
	AudioOutputs int
	RelayOutputs int
}

type DisplayCapabilities struct {
	XAddr       xsd.AnyURI
	FixedLayout bool
}

type RecordingCapabilities struct {
	XAddr              xsd.AnyURI
	ReceiverSource     bool
	MediaProfileSource bool
	DynamicRecordings  bool
	DynamicTracks      bool
	MaxStringLength    int
}

type SearchCapabilities struct {
	XAddr          xsd.AnyURI
	MetadataSearch bool
}

type ReplayCapabilities struct {
	XAddr xsd.AnyURI
}

type ReceiverCapabilities struct {
	XAddr                xsd.AnyURI
	RTP_Multicast        bool
	RTP_TCP              bool
	RTP_RTSP_TCP         bool
	SupportedReceivers   int
	MaximumRTSPURILength int
}

type AnalyticsDeviceCapabilities struct {
	XAddr       xsd.AnyURI
	RuleSupport bool
	Extension   AnalyticsDeviceExtension
}

type AnalyticsDeviceExtension xsd.AnyType

type CapabilitiesExtension2 xsd.AnyType

type HostnameInformation struct {
	FromDHCP  bool
	Name      xsd.Token
	Extension *HostnameInformationExtension
}

type HostnameInformationExtension xsd.AnyType

type DNSInformation struct {
	FromDHCP     bool
	SearchDomain xsd.Token
	DNSFromDHCP  IPAddress
	DNSManual    IPAddress
	Extension    *DNSInformationExtension
}

type DNSInformationExtension xsd.AnyType

type NTPInformation struct {
	FromDHCP    bool
	NTPFromDHCP *NetworkHost
	NTPManual   *NetworkHost
	Extension   *NTPInformationExtension
}

type NTPInformationExtension xsd.AnyType

type DynamicDNSInformation struct {
	Type      DynamicDNSType
	Name      DNSName
	TTL       xsd.Duration
	Extension DynamicDNSInformationExtension
}

//TODO: enumeration
type DynamicDNSType xsd.String

type DynamicDNSInformationExtension xsd.AnyType

type NetworkInterface struct {
	DeviceEntity
	Enabled   bool
	Info      *NetworkInterfaceInfo
	Link      *NetworkInterfaceLink
	IPv4      *IPv4NetworkInterface
	IPv6      *IPv6NetworkInterface
	Extension *NetworkInterfaceExtension
}

type NetworkInterfaceInfo struct {
	Name      xsd.String
	HwAddress HwAddress
	MTU       int
}

type HwAddress xsd.Token

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting
	OperSettings  NetworkInterfaceConnectionSetting
	InterfaceType IANA_IfTypes `xml:"IANA-IfTypes"`
}

type IANA_IfTypes int

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation bool   `xml:"http://www.onvif.org/ver10/schema AutoNegotiation"`
	Speed           int    `xml:"http://www.onvif.org/ver10/schema Speed"`
	Duplex          Duplex `xml:"http://www.onvif.org/ver10/schema Duplex"`
}

//TODO: enum
type Duplex xsd.String

type NetworkInterfaceExtension struct {
	InterfaceType IANA_IfTypes
	Dot3          Dot3Configuration
	Dot11         Dot11Configuration
	Extension     NetworkInterfaceExtension2
}

type NetworkInterfaceExtension2 xsd.AnyType

type Dot11Configuration struct {
	SSID     Dot11SSIDType                  `xml:"http://www.onvif.org/ver10/schema SSID"`
	Mode     Dot11StationMode               `xml:"http://www.onvif.org/ver10/schema Mode"`
	Alias    Name                           `xml:"http://www.onvif.org/ver10/schema Alias"`
	Priority NetworkInterfaceConfigPriority `xml:"http://www.onvif.org/ver10/schema Priority"`
	Security Dot11SecurityConfiguration     `xml:"http://www.onvif.org/ver10/schema Security"`
}

type Dot11SecurityConfiguration struct {
	Mode      Dot11SecurityMode                   `xml:"http://www.onvif.org/ver10/schema Mode"`
	Algorithm Dot11Cipher                         `xml:"http://www.onvif.org/ver10/schema Algorithm"`
	PSK       Dot11PSKSet                         `xml:"http://www.onvif.org/ver10/schema PSK"`
	Dot1X     ReferenceToken                      `xml:"http://www.onvif.org/ver10/schema Dot1X"`
	Extension Dot11SecurityConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type Dot11SecurityConfigurationExtension xsd.AnyType

type Dot11PSKSet struct {
	Key        Dot11PSK             `xml:"http://www.onvif.org/ver10/schema Key"`
	Passphrase Dot11PSKPassphrase   `xml:"http://www.onvif.org/ver10/schema Passphrase"`
	Extension  Dot11PSKSetExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type Dot11PSKSetExtension xsd.AnyType

type Dot11PSKPassphrase xsd.String

type Dot11PSK xsd.HexBinary

//TODO: enumeration
type Dot11Cipher xsd.String

//TODO: enumeration
type Dot11SecurityMode xsd.String

//TODO: restrictions
type NetworkInterfaceConfigPriority xsd.Integer

//TODO: enumeration
type Dot11StationMode xsd.String

//TODO: restrictions
type Dot11SSIDType xsd.HexBinary

type Dot3Configuration xsd.String

type IPv6NetworkInterface struct {
	Enabled bool
	Config  IPv6Configuration
}

type IPv6Configuration struct {
	AcceptRouterAdvert bool
	DHCP               IPv6DHCPConfiguration
	Manual             PrefixedIPv6Address
	LinkLocal          PrefixedIPv6Address
	FromDHCP           PrefixedIPv6Address
	FromRA             PrefixedIPv6Address
	Extension          IPv6ConfigurationExtension
}

type IPv6ConfigurationExtension xsd.AnyType

type PrefixedIPv6Address struct {
	Address      IPv6Address `xml:"http://www.onvif.org/ver10/schema Address"`
	PrefixLength int         `xml:"http://www.onvif.org/ver10/schema PrefixLength"`
}

//TODO: enumeration
type IPv6DHCPConfiguration xsd.String

type IPv4NetworkInterface struct {
	Enabled bool
	Config  IPv4Configuration
}

type IPv4Configuration struct {
	Manual    PrefixedIPv4Address
	LinkLocal PrefixedIPv4Address
	FromDHCP  PrefixedIPv4Address
	DHCP      bool
}

//optional, unbounded
type PrefixedIPv4Address struct {
	Address      IPv4Address `xml:"http://www.onvif.org/ver10/schema Address"`
	PrefixLength int         `xml:"http://www.onvif.org/ver10/schema PrefixLength"`
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   bool                                      `xml:"http://www.onvif.org/ver10/schema Enabled"`
	Link      NetworkInterfaceConnectionSetting         `xml:"http://www.onvif.org/ver10/schema Link"`
	MTU       int                                       `xml:"http://www.onvif.org/ver10/schema MTU"`
	IPv4      IPv4NetworkInterfaceSetConfiguration      `xml:"http://www.onvif.org/ver10/schema IPv4"`
	IPv6      IPv6NetworkInterfaceSetConfiguration      `xml:"http://www.onvif.org/ver10/schema IPv6"`
	Extension NetworkInterfaceSetConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      Dot3Configuration                          `xml:"http://www.onvif.org/ver10/schema Dot3"`
	Dot11     Dot11Configuration                         `xml:"http://www.onvif.org/ver10/schema Dot11"`
	Extension NetworkInterfaceSetConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type NetworkInterfaceSetConfigurationExtension2 xsd.AnyType

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            bool                  `xml:"http://www.onvif.org/ver10/schema Enabled"`
	AcceptRouterAdvert bool                  `xml:"http://www.onvif.org/ver10/schema AcceptRouterAdvert"`
	Manual             PrefixedIPv6Address   `xml:"http://www.onvif.org/ver10/schema Manual"`
	DHCP               IPv6DHCPConfiguration `xml:"http://www.onvif.org/ver10/schema DHCP"`
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled bool                `xml:"http://www.onvif.org/ver10/schema Enabled"`
	Manual  PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema Manual"`
	DHCP    bool                `xml:"http://www.onvif.org/ver10/schema DHCP"`
}

type NetworkProtocol struct {
	Name      NetworkProtocolType       `xml:"http://www.onvif.org/ver10/schema Name"`
	Enabled   bool                      `xml:"http://www.onvif.org/ver10/schema Enabled"`
	Port      int                       `xml:"http://www.onvif.org/ver10/schema Port"`
	Extension *NetworkProtocolExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type NetworkProtocolExtension xsd.AnyType

//TODO: enumeration
type NetworkProtocolType xsd.String

type NetworkGateway struct {
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type NetworkZeroConfiguration struct {
	InterfaceToken ReferenceToken
	Enabled        bool
	Addresses      IPv4Address
	Extension      NetworkZeroConfigurationExtension
}

type NetworkZeroConfigurationExtension struct {
	Additional *NetworkZeroConfiguration
	Extension  NetworkZeroConfigurationExtension2
}

type NetworkZeroConfigurationExtension2 xsd.AnyType

type IPAddressFilter struct {
	Type        IPAddressFilterType      `xml:"http://www.onvif.org/ver10/schema Type"`
	IPv4Address PrefixedIPv4Address      `xml:"http://www.onvif.org/ver10/schema IPv4Address,omitempty"`
	IPv6Address PrefixedIPv6Address      `xml:"http://www.onvif.org/ver10/schema IPv6Address,omitempty"`
	Extension   IPAddressFilterExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

type IPAddressFilterExtension xsd.AnyType

//enum { 'Allow', 'Deny' }
//TODO: enumeration
type IPAddressFilterType xsd.String

//TODO: attribite <xs:attribute ref="xmime:contentType" use="optional"/>
type BinaryData struct {
	X    ContentType      `xml:"xmime:contentType,attr"`
	Data xsd.Base64Binary `xml:"http://www.onvif.org/ver10/schema Data"`
}

type Certificate struct {
	CertificateID xsd.Token  `xml:"http://www.onvif.org/ver10/schema CertificateID"`
	Certificate   BinaryData `xml:"http://www.onvif.org/ver10/schema Certificate"`
}

type CertificateStatus struct {
	CertificateID xsd.Token `xml:"http://www.onvif.org/ver10/schema CertificateID"`
	Status        bool      `xml:"http://www.onvif.org/ver10/schema Status"`
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings
}

type RelayOutputSettings struct {
	Mode      RelayMode      `xml:"http://www.onvif.org/ver10/schema Mode"`
	DelayTime xsd.Duration   `xml:"http://www.onvif.org/ver10/schema DelayTime"`
	IdleState RelayIdleState `xml:"http://www.onvif.org/ver10/schema IdleState"`
}

//TODO:enumeration
type RelayIdleState xsd.String

//TODO: enumeration
type RelayMode xsd.String

//TODO: enumeration
type RelayLogicalState xsd.String

type CertificateWithPrivateKey struct {
	CertificateID xsd.Token  `xml:"http://www.onvif.org/ver10/schema CertificateID"`
	Certificate   BinaryData `xml:"http://www.onvif.org/ver10/schema Certificate"`
	PrivateKey    BinaryData `xml:"http://www.onvif.org/ver10/schema PrivateKey"`
}

type CertificateInformation struct {
	CertificateID      xsd.Token
	IssuerDN           xsd.String
	SubjectDN          xsd.String
	KeyUsage           CertificateUsage
	ExtendedKeyUsage   CertificateUsage
	KeyLength          int
	Version            xsd.String
	SerialNum          xsd.String
	SignatureAlgorithm xsd.String
	Validity           DateTimeRange
	Extension          CertificateInformationExtension
}

type CertificateInformationExtension xsd.AnyType

type DateTimeRange struct {
	From  xsd.DateTime
	Until xsd.DateTime
}

type CertificateUsage struct {
	Critical         bool `xml:"Critical,attr"`
	CertificateUsage xsd.String
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken              `xml:"http://www.onvif.org/ver10/schema Dot1XConfigurationToken"`
	Identity                xsd.String                  `xml:"http://www.onvif.org/ver10/schema Identity"`
	AnonymousID             xsd.String                  `xml:"http://www.onvif.org/ver10/schema AnonymousID,omitempty"`
	EAPMethod               int                         `xml:"http://www.onvif.org/ver10/schema EAPMethod"`
	CACertificateID         xsd.Token                   `xml:"http://www.onvif.org/ver10/schema CACertificateID,omitempty"`
	EAPMethodConfiguration  EAPMethodConfiguration      `xml:"http://www.onvif.org/ver10/schema EAPMethodConfiguration,omitempty"`
	Extension               Dot1XConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

type Dot1XConfigurationExtension xsd.AnyType

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration   `xml:"http://www.onvif.org/ver10/schema TLSConfiguration,omitempty"`
	Password         xsd.String         `xml:"http://www.onvif.org/ver10/schema Password,omitempty"`
	Extension        EapMethodExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

type EapMethodExtension xsd.AnyType

type TLSConfiguration struct {
	CertificateID xsd.Token `xml:"http://www.onvif.org/ver10/schema CertificateID,omitempty"`
}

type Dot11Capabilities struct {
	TKIP                  bool
	ScanAvailableNetworks bool
	MultipleConfiguration bool
	AdHocStationMode      bool
	WEP                   bool
}

type Dot11Status struct {
	SSID              Dot11SSIDType
	BSSID             xsd.String
	PairCipher        Dot11Cipher
	GroupCipher       Dot11Cipher
	SignalStrength    Dot11SignalStrength
	ActiveConfigAlias ReferenceToken
}

//TODO: enumeration
type Dot11SignalStrength xsd.String

type Dot11AvailableNetworks struct {
	SSID                  Dot11SSIDType
	BSSID                 xsd.String
	AuthAndMangementSuite Dot11AuthAndMangementSuite
	PairCipher            Dot11Cipher
	GroupCipher           Dot11Cipher
	SignalStrength        Dot11SignalStrength
	Extension             Dot11AvailableNetworksExtension
}

type Dot11AvailableNetworksExtension xsd.AnyType

//TODO: enumeration
type Dot11AuthAndMangementSuite xsd.String

type SystemLogUriList struct {
	SystemLog SystemLogUri
}

type SystemLogUri struct {
	Type SystemLogType
	Uri  xsd.AnyURI
}

type LocationEntity struct {
	Entity    xsd.String     `xml:"Entity,attr"`
	Token     ReferenceToken `xml:"Token,attr"`
	Fixed     bool           `xml:"Fixed,attr"`
	GeoSource xsd.AnyURI     `xml:"GeoSource,attr"`
	AutoGeo   bool           `xml:"AutoGeo,attr"`

	GeoLocation      GeoLocation      `xml:"http://www.onvif.org/ver10/schema GeoLocation"`
	GeoOrientation   GeoOrientation   `xml:"http://www.onvif.org/ver10/schema GeoOrientation"`
	LocalLocation    LocalLocation    `xml:"http://www.onvif.org/ver10/schema LocalLocation"`
	LocalOrientation LocalOrientation `xml:"http://www.onvif.org/ver10/schema LocalOrientation"`
}

type LocalOrientation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type LocalLocation struct {
	X xsd.Float `xml:"x,attr"`
	Y xsd.Float `xml:"y,attr"`
	Z xsd.Float `xml:"z,attr"`
}

type GeoOrientation struct {
	Roll  xsd.Float `xml:"roll,attr"`
	Pitch xsd.Float `xml:"pitch,attr"`
	Yaw   xsd.Float `xml:"yaw,attr"`
}

type FocusMove struct {
	Absolute   AbsoluteFocus   `xml:"http://www.onvif.org/ver10/schema Absolute"`
	Relative   RelativeFocus   `xml:"http://www.onvif.org/ver10/schema Relative"`
	Continuous ContinuousFocus `xml:"http://www.onvif.org/ver10/schema Continuous"`
}

type ContinuousFocus struct {
	Speed xsd.Float `xml:"http://www.onvif.org/ver10/schema Speed"`
}

type RelativeFocus struct {
	Distance xsd.Float `xml:"http://www.onvif.org/ver10/schema Distance"`
	Speed    xsd.Float `xml:"http://www.onvif.org/ver10/schema Speed"`
}

type AbsoluteFocus struct {
	Position xsd.Float `xml:"http://www.onvif.org/ver10/schema Position"`
	Speed    xsd.Float `xml:"http://www.onvif.org/ver10/schema Speed"`
}

type DateTime struct {
	Time Time `xml:"http://www.onvif.org/ver10/schema Time"`
	Date Date `xml:"http://www.onvif.org/ver10/schema Date"`
}

type Time struct {
	Hour   int `xml:"http://www.onvif.org/ver10/schema Hour"`
	Minute int `xml:"http://www.onvif.org/ver10/schema Minute"`
	Second int `xml:"http://www.onvif.org/ver10/schema Second"`
}

type Date struct {
	Year  int `xml:"http://www.onvif.org/ver10/schema Year"`
	Month int `xml:"http://www.onvif.org/ver10/schema Month"`
	Day   int `xml:"http://www.onvif.org/ver10/schema Day"`
}
