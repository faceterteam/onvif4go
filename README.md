# onvif4go

device := onvif4go.NewOnvifDevice("192.168.1.1:80")
device.Auth("admin", "1234456")
err := device.Initialize()
...

deviceInformation, err := device.Device.GetDeviceInformation()
...

err := device.Media.AddVideoAnalyticsConfiguration("Profile_0", "VideoAnalyticsConfiguration_000")
...

response := CustomDeviceResponse{}
err := device.Device.Call(CustomDeviceRequest{}, &response)
...

response := trt.GetProfilesResponse{}
media, ok := device.On("media")
err := media.Call(trt.GetProfiles{}, &response)
...

response := onvif4go.XmlNode{}
analytics, _ := device.On("analytics")
err = analytics.Call(onvif4go.XmlNode{
    XMLName: xml.Name{"http://www.onvif.org/ver20/analytics/wsdl", "GetSupportedAnalyticsModules"},
    Childs: []onvif4go.XmlNode{
        onvif4go.XmlNode{
            XMLName: xml.Name{"http://www.onvif.org/ver20/analytics/wsdl", "ConfigurationToken"},
            Text:    "VideoAnalyticsConfiguration_000",
        },
    },
}, &response)