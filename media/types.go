package media

import (
	"github.com/atagirov/onvif4go/onvif"
)

type Capabilities struct {
	SnapshotUri           bool `xml:"SnapshotUri,attr"`
	Rotation              bool `xml:"Rotation,attr"`
	VideoSourceMode       bool `xml:"VideoSourceMode,attr"`
	OSD                   bool `xml:"OSD,attr"`
	TemporaryOSDText      bool `xml:"TemporaryOSDText,attr"`
	EXICompression        bool `xml:"EXICompression,attr"`
	ProfileCapabilities   ProfileCapabilities
	StreamingCapabilities StreamingCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type StreamingCapabilities struct {
	RTPMulticast        bool `xml:"RTPMulticast,attr"`
	RTP_TCP             bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP        bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr"`
	NoRTSPStreaming     bool `xml:"NoRTSPStreaming,attr"`
}

//Media main types

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetVideoSources struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSources"`
}

type GetVideoSourcesResponse struct {
	VideoSources []onvif.VideoSource
}

type GetAudioSources struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSources"`
}

type GetAudioSourcesResponse struct {
	AudioSources []onvif.AudioSource
}

type GetAudioOutputs struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputs"`
}

type GetAudioOutputsResponse struct {
	AudioOutputs []onvif.AudioOutput
}

type CreateProfile struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/media/wsdl CreateProfile"`
	Name    onvif.Name           `xml:"http://www.onvif.org/ver10/media/wsdl Name"`
	Token   onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Token"`
}

type CreateProfileResponse struct {
	Profile onvif.Profile
}

type GetProfile struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetProfile"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetProfileResponse struct {
	Profiles onvif.Profile
}

type GetProfiles struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetProfiles"`
}

type GetProfilesResponse struct {
	Profiles []onvif.Profile
}

type AddVideoEncoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoEncoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddVideoEncoderConfigurationResponse struct {
}

type RemoveVideoEncoderConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoEncoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveVideoEncoderConfigurationResponse struct {
}

type AddVideoSourceConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoSourceConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddVideoSourceConfigurationResponse struct {
}

type RemoveVideoSourceConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoSourceConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveVideoSourceConfigurationResponse struct {
}

type AddAudioEncoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioEncoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddAudioEncoderConfigurationResponse struct {
}

type RemoveAudioEncoderConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioEncoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveAudioEncoderConfigurationResponse struct {
}

type AddAudioSourceConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioSourceConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddAudioSourceConfigurationResponse struct {
}

type RemoveAudioSourceConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioSourceConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveAudioSourceConfigurationResponse struct {
}

type AddPTZConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddPTZConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddPTZConfigurationResponse struct {
}

type RemovePTZConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemovePTZConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemovePTZConfigurationResponse struct {
}

type AddVideoAnalyticsConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoAnalyticsConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddVideoAnalyticsConfigurationResponse struct {
}

type RemoveVideoAnalyticsConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoAnalyticsConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveVideoAnalyticsConfigurationResponse struct {
}

type AddMetadataConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddMetadataConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddMetadataConfigurationResponse struct {
}

type RemoveMetadataConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveMetadataConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveMetadataConfigurationResponse struct {
}

type AddAudioOutputConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioOutputConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddAudioOutputConfigurationResponse struct {
}

type RemoveAudioOutputConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioOutputConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveAudioOutputConfigurationResponse struct {
}

type AddAudioDecoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioDecoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type AddAudioDecoderConfigurationResponse struct {
}

type RemoveAudioDecoderConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioDecoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type RemoveAudioDecoderConfigurationResponse struct {
}

type DeleteProfile struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl DeleteProfile"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type DeleteProfileResponse struct {
}

type GetVideoSourceConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfigurations"`
}

type GetVideoSourceConfigurationsResponse struct {
	Configurations []onvif.VideoSourceConfiguration
}

type GetVideoEncoderConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfigurations"`
}

type GetVideoEncoderConfigurationsResponse struct {
	Configurations []onvif.VideoEncoderConfiguration
}

type GetAudioSourceConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfigurations"`
}

type GetAudioSourceConfigurationsResponse struct {
	Configurations []onvif.AudioSourceConfiguration
}

type GetAudioEncoderConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfigurations"`
}

type GetAudioEncoderConfigurationsResponse struct {
	Configurations []onvif.AudioEncoderConfiguration
}

type GetVideoAnalyticsConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoAnalyticsConfigurations"`
}

type GetVideoAnalyticsConfigurationsResponse struct {
	Configurations []onvif.VideoAnalyticsConfiguration
}

type GetMetadataConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfigurations"`
}

type GetMetadataConfigurationsResponse struct {
	Configurations []onvif.MetadataConfiguration
}

type GetAudioOutputConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfigurations"`
}

type GetAudioOutputConfigurationsResponse struct {
	Configurations []onvif.AudioOutputConfiguration
}

type GetAudioDecoderConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfigurations"`
}

type GetAudioDecoderConfigurationsResponse struct {
	Configurations []onvif.AudioDecoderConfiguration
}

type GetVideoSourceConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetVideoSourceConfigurationResponse struct {
	Configuration onvif.VideoSourceConfiguration
}

type GetVideoEncoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetVideoEncoderConfigurationResponse struct {
	Configuration onvif.VideoEncoderConfiguration
}

type GetAudioSourceConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioSourceConfigurationResponse struct {
	Configuration onvif.AudioSourceConfiguration
}

type GetAudioEncoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioEncoderConfigurationResponse struct {
	Configuration onvif.AudioEncoderConfiguration
}

type GetVideoAnalyticsConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoAnalyticsConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetVideoAnalyticsConfigurationResponse struct {
	Configuration onvif.VideoAnalyticsConfiguration
}

type GetMetadataConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetMetadataConfigurationResponse struct {
	Configuration onvif.MetadataConfiguration
}

type GetAudioOutputConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioOutputConfigurationResponse struct {
	Configuration onvif.AudioOutputConfiguration
}

type GetAudioDecoderConfiguration struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioDecoderConfigurationResponse struct {
	Configuration onvif.AudioDecoderConfiguration
}

type GetCompatibleVideoEncoderConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoEncoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleVideoEncoderConfigurationsResponse struct {
	Configurations []onvif.VideoEncoderConfiguration
}

type GetCompatibleVideoSourceConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoSourceConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleVideoSourceConfigurationsResponse struct {
	Configurations []onvif.VideoSourceConfiguration
}

type GetCompatibleAudioEncoderConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioEncoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleAudioEncoderConfigurationsResponse struct {
	Configurations []onvif.AudioEncoderConfiguration
}

type GetCompatibleAudioSourceConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioSourceConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleAudioSourceConfigurationsResponse struct {
	Configurations []onvif.AudioSourceConfiguration
}

type GetCompatibleVideoAnalyticsConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoAnalyticsConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	Configurations []onvif.VideoAnalyticsConfiguration
}

type GetCompatibleMetadataConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleMetadataConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleMetadataConfigurationsResponse struct {
	Configurations []onvif.MetadataConfiguration
}

type GetCompatibleAudioOutputConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioOutputConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleAudioOutputConfigurationsResponse struct {
	Configurations []onvif.AudioOutputConfiguration
}

type GetCompatibleAudioDecoderConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioDecoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetCompatibleAudioDecoderConfigurationsResponse struct {
	Configurations []onvif.AudioDecoderConfiguration
}

type SetVideoSourceConfiguration struct {
	XMLName          string                         `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoSourceConfiguration"`
	Configuration    onvif.VideoSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                           `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetVideoSourceConfigurationResponse struct {
}

type SetVideoEncoderConfiguration struct {
	XMLName          string                          `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoEncoderConfiguration"`
	Configuration    onvif.VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                            `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetVideoEncoderConfigurationResponse struct {
}

type SetAudioSourceConfiguration struct {
	XMLName          string                         `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioSourceConfiguration"`
	Configuration    onvif.AudioSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                           `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetAudioSourceConfigurationResponse struct {
}

type SetAudioEncoderConfiguration struct {
	XMLName          string                          `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioEncoderConfiguration"`
	Configuration    onvif.AudioEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                            `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetAudioEncoderConfigurationResponse struct {
}

type SetVideoAnalyticsConfiguration struct {
	XMLName          string                            `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoAnalyticsConfiguration"`
	Configuration    onvif.VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                              `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetVideoAnalyticsConfigurationResponse struct {
}

type SetMetadataConfiguration struct {
	XMLName          string                      `xml:"http://www.onvif.org/ver10/media/wsdl GetDeviceInformation"`
	Configuration    onvif.MetadataConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                        `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetMetadataConfigurationResponse struct {
}

type SetAudioOutputConfiguration struct {
	XMLName          string                         `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioOutputConfiguration"`
	Configuration    onvif.AudioOutputConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                           `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetAudioOutputConfigurationResponse struct {
}

type SetAudioDecoderConfiguration struct {
	XMLName          string                          `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioDecoderConfiguration"`
	Configuration    onvif.AudioDecoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration"`
	ForcePersistence bool                            `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence"`
}

type SetAudioDecoderConfigurationResponse struct {
}

type GetVideoSourceConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetVideoSourceConfigurationOptionsResponse struct {
	Options onvif.VideoSourceConfigurationOptions
}

type GetVideoEncoderConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetVideoEncoderConfigurationOptionsResponse struct {
	Options onvif.VideoEncoderConfigurationOptions
}

type GetAudioSourceConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioSourceConfigurationOptionsResponse struct {
	Options onvif.AudioSourceConfigurationOptions
}

type GetAudioEncoderConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioEncoderConfigurationOptionsResponse struct {
	Options onvif.AudioEncoderConfigurationOptions
}

type GetMetadataConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetMetadataConfigurationOptionsResponse struct {
	Options onvif.MetadataConfigurationOptions
}

type GetAudioOutputConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioOutputConfigurationOptionsResponse struct {
	Options onvif.AudioOutputConfigurationOptions
}

type GetAudioDecoderConfigurationOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetAudioDecoderConfigurationOptionsResponse struct {
	Options onvif.AudioDecoderConfigurationOptions
}

type GetGuaranteedNumberOfVideoEncoderInstances struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetGuaranteedNumberOfVideoEncoderInstances"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	TotalNumber int
	JPEG        int
	H264        int
	MPEG4       int
}

type GetStreamUri struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetStreamUri"`
	StreamSetup  onvif.StreamSetup    `xml:"http://www.onvif.org/ver10/media/wsdl StreamSetup"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetStreamUriResponse struct {
	MediaUri onvif.MediaUri
}

type StartMulticastStreaming struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl StartMulticastStreaming"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type StartMulticastStreamingResponse struct {
}

type StopMulticastStreaming struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl StopMulticastStreaming"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type StopMulticastStreamingResponse struct {
}

type SetSynchronizationPoint struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl SetSynchronizationPoint"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type SetSynchronizationPointResponse struct {
}

type GetSnapshotUri struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/media/wsdl GetSnapshotUri"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken"`
}

type GetSnapshotUriResponse struct {
	MediaUri onvif.MediaUri
}

type GetVideoSourceModes struct {
	XMLName          string               `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceModes"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceToken"`
}

type GetVideoSourceModesResponse struct {
	VideoSourceModes []onvif.VideoSourceMode
}

type SetVideoSourceMode struct {
	XMLName              string               `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoSourceMode"`
	VideoSourceToken     onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceToken"`
	VideoSourceModeToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceModeToken"`
}

type SetVideoSourceModeResponse struct {
	Reboot bool
}

type GetOSDs struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetOSDs"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetOSDsResponse struct {
	OSDs []onvif.OSDConfiguration
}

type GetOSD struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/media/wsdl GetOSD"`
	OSDToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OSDToken"`
}

type GetOSDResponse struct {
	OSD onvif.OSDConfiguration
}

type GetOSDOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/media/wsdl GetOSDOptions"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken"`
}

type GetOSDOptionsResponse struct {
	OSDOptions onvif.OSDConfigurationOptions
}

type SetOSD struct {
	XMLName string                 `xml:"http://www.onvif.org/ver10/media/wsdl SetOSD"`
	OSD     onvif.OSDConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl OSD"`
}

type SetOSDResponse struct {
}

type CreateOSD struct {
	XMLName string                 `xml:"http://www.onvif.org/ver10/media/wsdl CreateOSD"`
	OSD     onvif.OSDConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl OSD"`
}

type CreateOSDResponse struct {
	OSDToken onvif.ReferenceToken
}

type DeleteOSD struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/media/wsdl DeleteOSD"`
	OSDToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OSDToken"`
}

type DeleteOSDResponse struct {
}
