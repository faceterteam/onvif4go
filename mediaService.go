package onvif4go

import (
	trt "github.com/atagirov/onvif4go/media"
	tt "github.com/atagirov/onvif4go/onvif"
)

type MediaService struct {
	Client onvifCaller
}

func NewMediaService(endpoint string, onvifAuth *onvifAuth) *MediaService {
	return &MediaService{
		Client: NewOnvifClient(endpoint, onvifAuth),
	}
}

func (s *MediaService) WithoutAuth() *MediaService {
	return &MediaService{
		Client: s.Client.WithoutAuth(),
	}
}

// GetProfiles using for ask the existing media profiles of a device
// Pre-configured or dynamically configured profiles can be retrieved using this command.
// This command lists all configured profiles in a device.
// The client does not need to know the media profile in order to use the command.
func (s *MediaService) GetProfiles() (res trt.GetProfilesResponse, err error) {
	err = s.Client.Call(trt.GetProfiles{}, &res)
	return
}

// GetProfile return profile is the profile token is already known
func (s *MediaService) GetProfile(profileToken string) (res trt.GetProfileResponse, err error) {
	err = s.Client.Call(trt.GetProfile{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetServiceCapabilities returns the capabilities of the media service.
func (s *MediaService) GetServiceCapabilities() (res trt.GetServiceCapabilitiesResponse, err error) {
	err = s.Client.Call(trt.GetServiceCapabilities{}, &res)
	return
}

// GetVideoSources lists all available physical video inputs of the device.
func (s *MediaService) GetVideoSources() (res trt.GetVideoSourcesResponse, err error) {
	err = s.Client.Call(trt.GetVideoSources{}, &res)
	return
}

// GetAudioSources lists all available physical audio inputs of the device.
func (s *MediaService) GetAudioSources() (res trt.GetAudioSourcesResponse, err error) {
	err = s.Client.Call(trt.GetAudioSources{}, &res)
	return
}

// GetAudioOutputs lists all available physical audio outputs of the device.
func (s *MediaService) GetAudioOutputs() (res trt.GetAudioOutputsResponse, err error) {
	err = s.Client.Call(trt.GetAudioOutputs{}, &res)
	return
}

// CreateProfile creates a new empty media profile. The media profile shall be created in the
// device and shall be persistent (remain after reboot).
//
// A created profile shall be deletable and a device shall set the “fixed” attribute to false in the
// returned Profile.
func (s *MediaService) CreateProfile(token, name string) (res trt.CreateProfileResponse, err error) {
	err = s.Client.Call(trt.CreateProfile{
		Token: tt.ReferenceToken(token),
		Name:  tt.Name(name),
	}, &res)
	return
}

// AddVideoEncoderConfiguration adds a VideoEncoderConfiguration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced. The change shall be persistent.
// A device shall support adding a compatible VideoEncoderConfiguration to a Profile containing a
// VideoSourceConfiguration and shall support streaming video data of such a profile.
func (s *MediaService) AddVideoEncoderConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddVideoEncoderConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddVideoEncoderConfigurationResponse{})
}

// RemoveVideoEncoderConfiguration removes a VideoEncoderConfiguration from an existing media profile.
// If the media profile does not contain a VideoEncoderConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveVideoEncoderConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveVideoEncoderConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveVideoEncoderConfigurationResponse{})
}

// AddVideoSourceConfiguration adds a VideoSourceConfiguration to an existing media profile.
// If such a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
func (s *MediaService) AddVideoSourceConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddVideoSourceConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddVideoSourceConfigurationResponse{})
}

// RemoveVideoSourceConfiguration removes a VideoSourceConfiguration from an existing media profile.
// If the media profile does not contain a VideoSourceConfiguration, the operation has no effect.
// The removal shall be persistent. Video source configurations should only be removed after removing a
// VideoEncoderConfiguration from the media profile.
func (s *MediaService) RemoveVideoSourceConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveVideoSourceConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveVideoSourceConfigurationResponse{})
}

// AddAudioEncoderConfiguration adds a AudioEncoderConfiguration to an existing media profile.
// If such a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
//
// A device shall support adding a compatible AudioEncoderConfiguration to a profile containing an
// AudioSourceConfiguration and shall support streaming audio data of such a profile.
func (s *MediaService) AddAudioEncoderConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddAudioEncoderConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddAudioEncoderConfigurationResponse{})
}

// RemoveAudioEncoderConfiguration removes an AudioEncoderConfiguration from an existing media profile.
// If the media profile does not contain an AudioEncoderConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveAudioEncoderConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveAudioEncoderConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveAudioEncoderConfigurationResponse{})
}

// AddAudioSourceConfiguration adds an AudioSourceConfiguration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
func (s *MediaService) AddAudioSourceConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddAudioSourceConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddAudioSourceConfigurationResponse{})
}

// RemoveAudioSourceConfiguration removes an AudioSourceConfiguration from an existing media profile.
// If the media profile does not contain an AudioSourceConfiguration, the operation has no effect.
// The removal shall be persistent.
//
// Audio source configurations should only be removed after removing an AudioEncoderConfiguration from the media profile.
func (s *MediaService) RemoveAudioSourceConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveAudioSourceConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveAudioSourceConfigurationResponse{})
}

// AddPTZConfiguration adds a PTZConfiguration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
// Adding a PTZConfiguration to a media profile means that streams using that media profile can
// contain PTZ status (in the metadata), and that the media profile can be used for controlling
// PTZ movement.
func (s *MediaService) AddPTZConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddPTZConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddPTZConfigurationResponse{})
}

// RemovePTZConfiguration removes a PTZConfiguration from an existing media profile.
// If the media profile does not contain a PTZConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemovePTZConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemovePTZConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemovePTZConfigurationResponse{})
}

// AddVideoAnalyticsConfiguration adds a VideoAnalytics configuration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
//
// Adding a VideoAnalyticsConfiguration to a media profile means that streams using that media profile
// can contain video analytics data (in the metadata) as defined by the submitted configuration reference.
// A profile containing only a video analytics configuration but no video source configuration is incomplete.
// Therefore, a client should first add a video source configuration to a profile before adding a video analytics configuration.
// The device can deny adding of a video analytics configuration before a video source configuration.
func (s *MediaService) AddVideoAnalyticsConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddVideoAnalyticsConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddVideoAnalyticsConfigurationResponse{})
}

// RemoveVideoAnalyticsConfiguration removes a VideoAnalyticsConfiguration from an existing media profile.
// If the media profile does not contain a VideoAnalyticsConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveVideoAnalyticsConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveVideoAnalyticsConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveVideoAnalyticsConfigurationResponse{})
}

// AddMetadataConfiguration adds a Metadata configuration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
//
// Adding a MetadataConfiguration to a Profile means that streams using that profile contain metadata.
// Metadata can consist of events, PTZ status, and/or video analytics data.
func (s *MediaService) AddMetadataConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddMetadataConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddMetadataConfigurationResponse{})
}

// RemoveMetadataConfiguration removes a MetadataConfiguration from an existing media profile.
// If the media profile does not contain a MetadataConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveMetadataConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveMetadataConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveMetadataConfigurationResponse{})
}

// AddAudioOutputConfiguration adds an AudioOutputConfiguration to an existing media profile.
// If a configuration exists in the media profile, it will be replaced.
// The change shall be persistent.
func (s *MediaService) AddAudioOutputConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddAudioOutputConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddAudioOutputConfigurationResponse{})
}

// RemoveAudioOutputConfiguration removes an AudioOutputConfiguration from an existing media profile.
// If the media profile does not contain an AudioOutputConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveAudioOutputConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveAudioOutputConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveAudioOutputConfigurationResponse{})
}

// AddAudioDecoderConfiguration adds an AudioDecoderConfiguration to an existing media profile.
// If a configuration exists in the media profile, it shall be replaced.
// The change shall be persistent.
func (s *MediaService) AddAudioDecoderConfiguration(profileToken, configurationToken string) error {
	return s.Client.Call(trt.AddAudioDecoderConfiguration{
		ProfileToken:       tt.ReferenceToken(profileToken),
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &trt.AddAudioDecoderConfigurationResponse{})
}

// RemoveAudioDecoderConfiguration removes an AudioDecoderConfiguration from an existing media profile.
// If the media profile does not contain an AudioDecoderConfiguration, the operation has no effect.
// The removal shall be persistent.
func (s *MediaService) RemoveAudioDecoderConfiguration(profileToken string) error {
	return s.Client.Call(trt.RemoveAudioDecoderConfiguration{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.RemoveAudioDecoderConfigurationResponse{})
}

// DeleteProfile deletes a profile. This change shall always be persistent.
// Deletion of a profile is only possible for non-fixed profiles
func (s *MediaService) DeleteProfile(profileToken string) error {
	return s.Client.Call(trt.DeleteProfile{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &trt.DeleteProfileResponse{})
}

// GetVideoSourceConfigurations lists all existing video source configurations for a device.
// The client need not know anything about the video source configurations in order to use the command.
func (s *MediaService) GetVideoSourceConfigurations() (res trt.GetVideoSourceConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetVideoSourceConfigurations{}, &res)
	return
}

//GetVideoEncoderConfigurations lists all existing video encoder configurations of a device.
// The client need not know anything apriori about the video encoder configurations in order to use the command.
func (s *MediaService) GetVideoEncoderConfigurations() (res trt.GetVideoEncoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetVideoEncoderConfigurations{}, &res)
	return
}

// GetAudioSourceConfigurations lists all existing audio source configurations of a device.
// The client need not know anything apriori about the audio source configurations in order to use the command.
func (s *MediaService) GetAudioSourceConfigurations() (res trt.GetAudioSourceConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetAudioSourceConfigurations{}, &res)
	return
}

// GetAudioEncoderConfigurations lists all existing device audio encoder configurations.
// The client need not know anything apriori about the audio encoder configurations in order to use the command.
func (s *MediaService) GetAudioEncoderConfigurations() (res trt.GetAudioEncoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetAudioEncoderConfigurations{}, &res)
	return
}

// GetVideoAnalyticsConfigurations lists all video analytics configurations of a device.
// The client need not know anything apriori about the video analytics in order to use the command.
func (s *MediaService) GetVideoAnalyticsConfigurations() (res trt.GetVideoAnalyticsConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetVideoAnalyticsConfigurations{}, &res)
	return
}

// GetMetadataConfigurations lists all existing metadata configurations.
// The client need not know anything apriori about the metadata in order to use the command.
func (s *MediaService) GetMetadataConfigurations() (res trt.GetMetadataConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetMetadataConfigurations{}, &res)
	return
}

// GetAudioOutputConfigurations lists all existing AudioOutputConfigurations of a device.
// The NVC need not know anything apriori about the audio configurations to use this command.
func (s *MediaService) GetAudioOutputConfigurations() (res trt.GetAudioOutputConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetAudioOutputConfigurations{}, &res)
	return
}

// GetAudioDecoderConfigurations lists all existing AudioDecoderConfigurations of a device.
// The NVC need not know anything apriori about the audio decoder configurations in order to use this command.
func (s *MediaService) GetAudioDecoderConfigurations() (res trt.GetAudioDecoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetAudioDecoderConfigurations{}, &res)
	return
}

// GetVideoSourceConfiguration return the VideoSourceConfiguration if the configuration token is known
func (s *MediaService) GetVideoSourceConfiguration(configurationToken string) (res trt.GetVideoSourceConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetVideoSourceConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetVideoEncoderConfiguration return the VideoEncoderConfiguration if the configuration token is known
func (s *MediaService) GetVideoEncoderConfiguration(configurationToken string) (res trt.GetVideoEncoderConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetVideoEncoderConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetAudioSourceConfiguration return the AudioSourceConfiguration if the configuration token is known
func (s *MediaService) GetAudioSourceConfiguration(configurationToken string) (res trt.GetAudioSourceConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetAudioSourceConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetAudioEncoderConfiguration return the AudioEncoderConfiguration if the configuration token is known
func (s *MediaService) GetAudioEncoderConfiguration(configurationToken string) (res trt.GetAudioEncoderConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetAudioEncoderConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetVideoAnalyticsConfiguration return the VideoAnalyticsConfiguration if the configuration token is known
func (s *MediaService) GetVideoAnalyticsConfiguration(configurationToken string) (res trt.GetVideoAnalyticsConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetVideoAnalyticsConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetMetadataConfiguration return the MetadataConfiguration if the configuration token is known
func (s *MediaService) GetMetadataConfiguration(configurationToken string) (res trt.GetMetadataConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetMetadataConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetAudioOutputConfiguration return the AudioOutputConfiguration if the configuration token is known
func (s *MediaService) GetAudioOutputConfiguration(configurationToken string) (res trt.GetAudioOutputConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetAudioOutputConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetAudioDecoderConfiguration return the AudioDecoderConfiguration if the configuration token is known
func (s *MediaService) GetAudioDecoderConfiguration(configurationToken string) (res trt.GetAudioDecoderConfigurationResponse, err error) {
	err = s.Client.Call(trt.GetAudioDecoderConfiguration{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetCompatibleVideoEncoderConfigurations lists all the video encoder configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddVideoEncoderConfiguration command on the media profile.
// The result will vary depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleVideoEncoderConfigurations(profileToken string) (res trt.GetCompatibleVideoEncoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleVideoEncoderConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleVideoSourceConfigurations requests all the video source configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddVideoSourceConfiguration command on the media profile.
// The result will vary depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleVideoSourceConfigurations(profileToken string) (res trt.GetCompatibleVideoSourceConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleVideoSourceConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleAudioEncoderConfigurations requests all audio encoder configurations of a device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddAudioSourceConfiguration command on the media profile.
// The result varies depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleAudioEncoderConfigurations(profileToken string) (res trt.GetCompatibleAudioEncoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleAudioEncoderConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleAudioSourceConfigurations requests all audio source configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddAudioEncoderConfiguration command on the media profile.
// The result varies depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleAudioSourceConfigurations(profileToken string) (res trt.GetCompatibleAudioSourceConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleAudioSourceConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleVideoAnalyticsConfigurations requests all video analytic configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddVideoAnalyticsConfiguration command on the media profile.
// The result varies depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleVideoAnalyticsConfigurations(profileToken string) (res trt.GetCompatibleVideoAnalyticsConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleVideoAnalyticsConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleMetadataConfigurations requests all the metadata configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddMetadataConfiguration command on the media profile.
// The result varies depending on the capabilities, configurations and settings in the device.
func (s *MediaService) GetCompatibleMetadataConfigurations(profileToken string) (res trt.GetCompatibleMetadataConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleMetadataConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleAudioOutputConfigurations lists all audio output configurations of a device that
// are compatible with a certain media profile. Each returned configuration shall be a valid input for the
// AddAudioOutputConfiguration command.
func (s *MediaService) GetCompatibleAudioOutputConfigurations(profileToken string) (res trt.GetCompatibleAudioOutputConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleAudioOutputConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetCompatibleAudioDecoderConfigurations lists all the audio decoder configurations of the device that
// are compatible with a certain media profile. Each of the returned configurations shall be a valid input
// parameter for the AddAudioDecoderConfiguration command on the media profile.
func (s *MediaService) GetCompatibleAudioDecoderConfigurations(profileToken string) (res trt.GetCompatibleAudioDecoderConfigurationsResponse, err error) {
	err = s.Client.Call(trt.GetCompatibleAudioDecoderConfigurations{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

/*
<wsdl:portType name="Media">
<!--===============================-->
<wsdl:operation name="SetVideoSourceConfiguration">
	<wsdl:documentation>This operation modifies a video source configuration. The ForcePersistence flag indicates if the changes shall remain after reboot of the device. Running streams using this configuration may be immediately updated according to the new settings. The changes are not guaranteed to take effect unless the client requests a new stream URI and restarts any affected stream. NVC methods for changing a running stream are out of scope for this specification.</wsdl:documentation>
	<wsdl:input message="trt:SetVideoSourceConfigurationRequest"/>
	<wsdl:output message="trt:SetVideoSourceConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetVideoEncoderConfiguration">
	<wsdl:documentation>This operation modifies a video encoder configuration. The ForcePersistence flag indicates if the changes shall remain after reboot of the device. Changes in the Multicast settings shall always be persistent. Running streams using this configuration may be immediately updated according to the new settings. The changes are not guaranteed to take effect unless the client requests a new stream URI and restarts any affected stream. NVC methods for changing a running stream are out of scope for this specification. <br/>SessionTimeout is provided as a hint for keeping rtsp session by a device. If necessary the device may adapt parameter values for SessionTimeout elements without returning an error. For the time between keep alive calls the client shall adhere to the timeout value signaled via RTSP.</wsdl:documentation>
	<wsdl:input message="trt:SetVideoEncoderConfigurationRequest"/>
	<wsdl:output message="trt:SetVideoEncoderConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetAudioSourceConfiguration">
	<wsdl:documentation>This operation modifies an audio source configuration. The ForcePersistence flag indicates if
the changes shall remain after reboot of the device. Running streams using this configuration
may be immediately updated according to the new settings. The changes are not guaranteed
to take effect unless the client requests a new stream URI and restarts any affected stream
NVC methods for changing a running stream are out of scope for this specification.</wsdl:documentation>
	<wsdl:input message="trt:SetAudioSourceConfigurationRequest"/>
	<wsdl:output message="trt:SetAudioSourceConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetAudioEncoderConfiguration">
	<wsdl:documentation>This operation modifies an audio encoder configuration. The ForcePersistence flag indicates if
the changes shall remain after reboot of the device. Running streams using this configuration may be immediately updated
according to the new settings. The changes are not guaranteed to take effect unless the client
requests a new stream URI and restarts any affected streams. NVC methods for changing a
running stream are out of scope for this specification.</wsdl:documentation>
	<wsdl:input message="trt:SetAudioEncoderConfigurationRequest"/>
	<wsdl:output message="trt:SetAudioEncoderConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetVideoAnalyticsConfiguration">
	<wsdl:documentation>A video analytics configuration is modified using this command. The ForcePersistence flag
indicates if the changes shall remain after reboot of the device or not. Running streams using
this configuration shall be immediately updated according to the new settings. Otherwise
inconsistencies can occur between the scene description processed by the rule engine and
the notifications produced by analytics engine and rule engine which reference the very same
video analytics configuration token.</wsdl:documentation>
	<wsdl:input message="trt:SetVideoAnalyticsConfigurationRequest"/>
	<wsdl:output message="trt:SetVideoAnalyticsConfigurationResponse"/>
</wsdl:operation>
*/

// SetMetadataConfiguration modifies a metadata configuration.
// The ForcePersistence flag indicates if the changes shall remain after reboot of the device.
// Changes in the Multicast settings shall always be persistent.
// Running streams using this configuration may be updated immediately according to the new settings.
// The changes are not guaranteed to take effect unless the client requests a new stream URI and
// restarts any affected streams.
// NVC methods for changing a running stream are out of scope for this specification.
func (s *MediaService) SetMetadataConfiguration(configuration tt.MetadataConfiguration) (res trt.SetMetadataConfigurationResponse, err error) {
	err = s.Client.Call(trt.SetMetadataConfiguration{
		ForcePersistence: true,
		Configuration:    configuration,
	}, &res)
	return
}

/*
<wsdl:operation name="SetAudioOutputConfiguration">
	<wsdl:documentation>This operation modifies an audio output configuration. The ForcePersistence flag indicates if
the changes shall remain after reboot of the device.</wsdl:documentation>
	<wsdl:input message="trt:SetAudioOutputConfigurationRequest"/>
	<wsdl:output message="trt:SetAudioOutputConfigurationResponse"/>
</wsdl:operation>
<wsdl:operation name="SetAudioDecoderConfiguration">
	<wsdl:documentation>This operation modifies an audio decoder configuration. The ForcePersistence flag indicates if
the changes shall remain after reboot of the device.</wsdl:documentation>
	<wsdl:input message="trt:SetAudioDecoderConfigurationRequest"/>
	<wsdl:output message="trt:SetAudioDecoderConfigurationResponse"/>
</wsdl:operation>
<!--===============================-->
<wsdl:operation name="GetVideoSourceConfigurationOptions">
	<wsdl:documentation>This operation returns the available options  (supported values and ranges for video source configuration parameters) when the video source parameters are
reconfigured If a video source configuration is specified, the options shall concern that
particular configuration. If a media profile is specified, the options shall be compatible with
that media profile.</wsdl:documentation>
	<wsdl:input message="trt:GetVideoSourceConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetVideoSourceConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetVideoEncoderConfigurationOptions">
	<wsdl:documentation>This operation returns the available options (supported values and ranges for video encoder
		configuration parameters) when the video encoder parameters are reconfigured. <br/>
		For JPEG, MPEG4 and H264 extension elements have been defined that provide additional information. A device must provide the
		XxxOption information for all encodings supported and should additionally provide the corresponding XxxOption2 information.<br/>
		This response contains the available video encoder configuration options. If a video encoder configuration is specified,
		the options shall concern that particular configuration. If a media profile is specified, the options shall be
		compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	</wsdl:documentation>
	<wsdl:input message="trt:GetVideoEncoderConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetVideoEncoderConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetAudioSourceConfigurationOptions">
	<wsdl:documentation>This operation returns the available options (supported values and ranges for audio source configuration parameters) when the audio source parameters are
reconfigured. If an audio source configuration is specified, the options shall concern that
particular configuration. If a media profile is specified, the options shall be compatible with
that media profile.</wsdl:documentation>
	<wsdl:input message="trt:GetAudioSourceConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetAudioSourceConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetAudioEncoderConfigurationOptions">
	<wsdl:documentation>This operation returns the available options  (supported values and ranges for audio encoder configuration parameters) when the audio encoder parameters are
reconfigured.</wsdl:documentation>
	<wsdl:input message="trt:GetAudioEncoderConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetAudioEncoderConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetMetadataConfigurationOptions">
	<wsdl:documentation>This operation returns the available options (supported values and ranges for metadata configuration parameters) for changing the metadata configuration.</wsdl:documentation>
	<wsdl:input message="trt:GetMetadataConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetMetadataConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetAudioOutputConfigurationOptions">
	<wsdl:documentation>This operation returns the available options (supported values and ranges for audio output configuration parameters) for configuring an audio output.</wsdl:documentation>
	<wsdl:input message="trt:GetAudioOutputConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetAudioOutputConfigurationOptionsResponse"/>
</wsdl:operation>
<wsdl:operation name="GetAudioDecoderConfigurationOptions">
	<wsdl:documentation>This command list the audio decoding capabilities for a given profile and configuration of a
device.</wsdl:documentation>
	<wsdl:input message="trt:GetAudioDecoderConfigurationOptionsRequest"/>
	<wsdl:output message="trt:GetAudioDecoderConfigurationOptionsResponse"/>
</wsdl:operation>
<!--===============================-->
<wsdl:operation name="GetGuaranteedNumberOfVideoEncoderInstances">
	<wsdl:documentation>The GetGuaranteedNumberOfVideoEncoderInstances command can be used to request the
minimum number of guaranteed video encoder instances (applications) per Video Source
Configuration.</wsdl:documentation>
	<wsdl:input message="trt:GetGuaranteedNumberOfVideoEncoderInstancesRequest"/>
	<wsdl:output message="trt:GetGuaranteedNumberOfVideoEncoderInstancesResponse"/>
</wsdl:operation>
<!--===============================-->
*/

// GetStreamURI requests a URI that can be used to initiate a live media stream using RTSP as
// the control protocol. The returned URI shall remain valid indefinitely even if the profile is
// changed. The ValidUntilConnect, ValidUntilReboot and Timeout Parameter shall be set
// accordingly (ValidUntilConnect=false, ValidUntilReboot=false, timeout=PT0S).
//
// The correct syntax for the StreamSetup element for these media stream setups defined in 5.1.1 of the streaming specification are as follows:
//
//		* RTP unicast over UDP: StreamType = "RTP_unicast", TransportProtocol = "UDP"
//		* RTP over RTSP over HTTP over TCP: StreamType = "RTP_unicast", TransportProtocol = "HTTP"
//		* RTP over RTSP over TCP: StreamType = "RTP_unicast", TransportProtocol = "RTSP"
//
// If a multicast stream is requested the VideoEncoderConfiguration, AudioEncoderConfiguration and MetadataConfiguration
// element inside the corresponding media profile must be configured with valid multicast settings.
//
// For full compatibility with other ONVIF services a device should not generate Uris longer than 128 octets.
func (s *MediaService) GetStreamURI(profileToken, streamType, transportProtocol string) (res trt.GetStreamUriResponse, err error) {
	err = s.Client.Call(trt.GetStreamUri{
		ProfileToken: tt.ReferenceToken(profileToken),
		StreamSetup: tt.StreamSetup{
			Stream: tt.StreamType(streamType),
			Transport: tt.Transport{
				Protocol: tt.TransportProtocol(transportProtocol),
			},
		},
	}, &res)
	return
}

/*
<wsdl:operation name="StartMulticastStreaming">
	<wsdl:documentation>This command starts multicast streaming using a specified media profile of a device.
Streaming continues until StopMulticastStreaming is called for the same Profile. The
streaming shall continue after a reboot of the device until a StopMulticastStreaming request is
received. The multicast address, port and TTL are configured in the
VideoEncoderConfiguration, AudioEncoderConfiguration and MetadataConfiguration
respectively.</wsdl:documentation>
	<wsdl:input message="trt:StartMulticastStreamingRequest"/>
	<wsdl:output message="trt:StartMulticastStreamingResponse"/>
</wsdl:operation>
<wsdl:operation name="StopMulticastStreaming">
	<wsdl:documentation>This command stop multicast streaming using a specified media profile of a device</wsdl:documentation>
	<wsdl:input message="trt:StopMulticastStreamingRequest"/>
	<wsdl:output message="trt:StopMulticastStreamingResponse"/>
</wsdl:operation>
<wsdl:operation name="SetSynchronizationPoint">
	<wsdl:documentation>Synchronization points allow clients to decode and correctly use all data after the
synchronization point.
For example, if a video stream is configured with a large I-frame distance and a client loses a
single packet, the client does not display video until the next I-frame is transmitted. In such
cases, the client can request a Synchronization Point which enforces the device to add an I-Frame as soon as possible. Clients can request Synchronization Points for profiles. The device
shall add synchronization points for all streams associated with this profile.
Similarly, a synchronization point is used to get an update on full PTZ or event status through
the metadata stream.
If a video stream is associated with the profile, an I-frame shall be added to this video stream.
If a PTZ metadata stream is associated to the profile,
the PTZ position shall be repeated within the metadata stream.</wsdl:documentation>
	<wsdl:input message="trt:SetSynchronizationPointRequest"/>
	<wsdl:output message="trt:SetSynchronizationPointResponse"/>
</wsdl:operation>
*/

// GetSnapshotURI uses to obtain a JPEG snapshot from the device.
// The returned URI shall remain valid indefinitely even if the profile is changed. The
// ValidUntilConnect, ValidUntilReboot and Timeout Parameter shall be set accordingly
// (ValidUntilConnect=false, ValidUntilReboot=false, timeout=PT0S). The URI can be used for
// acquiring a JPEG image through a HTTP GET operation. The image encoding will always be
// JPEG regardless of the encoding setting in the media profile. The Jpeg settings
// (like resolution or quality) may be taken from the profile if suitable. The provided
// image will be updated automatically and independent from calls to GetSnapshotUri.
func (s *MediaService) GetSnapshotURI(profileToken string) (res trt.GetSnapshotUriResponse, err error) {
	err = s.Client.Call(trt.GetSnapshotUri{
		ProfileToken: tt.ReferenceToken(profileToken),
	}, &res)
	return
}

// GetVideoSourceModes returns the information for current video source mode and settable
// video source modes of specified video source. A device that indicates a capability of
// VideoSourceModes shall support this command.
func (s *MediaService) GetVideoSourceModes(videoSourceToken string) (res trt.GetVideoSourceModesResponse, err error) {
	err = s.Client.Call(trt.GetVideoSourceModes{
		VideoSourceToken: tt.ReferenceToken(videoSourceToken),
	}, &res)
	return
}

// SetVideoSourceMode changes the media profile structure relating to video source for the
// specified video source mode. A device that indicates a capability of VideoSourceModes
// shall support this command.
// The behavior after changing the mode is not defined in this specification.
func (s *MediaService) SetVideoSourceMode(videoSourceToken, videoSourceModeToken string) (res trt.SetVideoSourceModeResponse, err error) {
	err = s.Client.Call(trt.SetVideoSourceMode{
		VideoSourceToken:     tt.ReferenceToken(videoSourceToken),
		VideoSourceModeToken: tt.ReferenceToken(videoSourceModeToken),
	}, &res)
	return
}

// GetOSDs returns the OSDs on specified VideoSourceConfiguration
func (s *MediaService) GetOSDs(configurationToken string) (res trt.GetOSDsResponse, err error) {
	err = s.Client.Call(trt.GetOSDs{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// GetOSD return the OSD
func (s *MediaService) GetOSD(osdToken string) (res trt.GetOSDResponse, err error) {
	err = s.Client.Call(trt.GetOSD{
		OSDToken: tt.ReferenceToken(osdToken),
	}, &res)
	return
}

// GetOSDOptions return the OSD Options.
func (s *MediaService) GetOSDOptions(configurationToken string) (res trt.GetOSDOptionsResponse, err error) {
	err = s.Client.Call(trt.GetOSDOptions{
		ConfigurationToken: tt.ReferenceToken(configurationToken),
	}, &res)
	return
}

// SetOSD set the OSD
func (s *MediaService) SetOSD(configuration tt.OSDConfiguration) (res trt.SetOSDResponse, err error) {
	err = s.Client.Call(trt.SetOSD{
		OSD: configuration,
	}, &res)
	return
}

// CreateOSD create the OSD
func (s *MediaService) CreateOSD(configuration tt.OSDConfiguration) (res trt.CreateOSDResponse, err error) {
	err = s.Client.Call(trt.CreateOSD{
		OSD: configuration,
	}, &res)
	return
}

// DeleteOSD delete the OSD
func (s *MediaService) DeleteOSD(osdToken string) (res trt.DeleteOSDResponse, err error) {
	err = s.Client.Call(trt.DeleteOSD{
		OSDToken: tt.ReferenceToken(osdToken),
	}, &res)
	return
}
