package onvif4go

import (
	"errors"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/atagirov/onvif4go/onvif"
	"github.com/atagirov/onvif4go/xsd"
)

type OnvifDevice struct {
	xaddr string
	auth  onvifAuth

	Capabilities onvif.Capabilities

	endpoints map[string]onvifCaller

	Device *DeviceService
	Media  *MediaService
	Events *EventsService
}

func (onvifDevice *OnvifDevice) Call(request, response interface{}) error {
	pkgPath := strings.Split(reflect.TypeOf(request).PkgPath(), "/")
	pkg := pkgPath[len(pkgPath)-1]

	client, ok := onvifDevice.On(strings.ToLower(pkg))
	if !ok {
		return errors.New("no endpoint")
	}
	return client.Call(request, response)
}

func (onvifDevice *OnvifDevice) On(endpoint string) (caller onvifCaller, ok bool) {
	caller, ok = onvifDevice.endpoints[endpoint]
	return
}

func NewOnvifDevice(xaddr string) *OnvifDevice {
	return &OnvifDevice{
		xaddr:     xaddr,
		endpoints: make(map[string]onvifCaller),
	}
}

func (onvifDevice *OnvifDevice) Auth(login, password string) {
	onvifDevice.auth.login = login
	onvifDevice.auth.password = password
}

func fixXAddr(uri xsd.AnyURI, xaddr string) string {
	u, _ := url.Parse(string(uri))
	u.Host = xaddr
	return u.String()
}

func (onvifDevice *OnvifDevice) Initialize() error {
	deviceService := NewDeviceService(onvifDevice)

	currentTime := time.Now().UTC()
	systemDateAndTimeResponse, err := deviceService.GetSystemDateAndTime()
	if err != nil {
		return err
	}
	deviceTime := systemDateAndTimeResponse.SystemDateAndTime.UTCDateTime.GetTime()
	onvifDevice.auth.timeDiff = deviceTime.Sub(currentTime)

	capabilitiesResponse, err := deviceService.GetCapabilities()
	if err != nil {
		return err
	}

	capabilities := capabilitiesResponse.Capabilities
	if capabilities.Device != nil {
		onvifDevice.endpoints["device"] = deviceService.Client
		onvifDevice.Device = deviceService
	}
	if capabilities.Events != nil {
		xaddr := fixXAddr(capabilities.Events.XAddr, onvifDevice.xaddr)
		onvifDevice.Events = NewEventsService(xaddr, &onvifDevice.auth)
		onvifDevice.endpoints["events"] = onvifDevice.Events.Client
	}
	if capabilities.Media != nil {
		mediaURI := fixXAddr(capabilities.Media.XAddr, onvifDevice.xaddr)
		onvifDevice.Media = NewMediaService(mediaURI, &onvifDevice.auth)
		onvifDevice.endpoints["media"] = onvifDevice.Media.Client
	}
	if capabilities.Imaging != nil {
		xaddr := fixXAddr(capabilities.Imaging.XAddr, onvifDevice.xaddr)
		onvifDevice.endpoints["imaging"] = NewOnvifClient(xaddr, &onvifDevice.auth)
	}
	if capabilities.Analytics != nil {
		xaddr := fixXAddr(capabilities.Analytics.XAddr, onvifDevice.xaddr)
		onvifDevice.endpoints["analytics"] = NewOnvifClient(xaddr, &onvifDevice.auth)
	}
	if capabilities.PTZ != nil {
		xaddr := fixXAddr(capabilities.PTZ.XAddr, onvifDevice.xaddr)
		onvifDevice.endpoints["ptz"] = NewOnvifClient(xaddr, &onvifDevice.auth)
	}

	if capabilities.Extension != nil {
		extension := capabilities.Extension
		if extension.AnalyticsDevice != nil {
			xaddr := fixXAddr(extension.AnalyticsDevice.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["analyticsdevice"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.DeviceIO != nil {
			xaddr := fixXAddr(extension.DeviceIO.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["deviceio"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.Display != nil {
			xaddr := fixXAddr(extension.Display.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["display"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.Receiver != nil {
			xaddr := fixXAddr(extension.Receiver.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["receiver"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.Recording != nil {
			xaddr := fixXAddr(extension.Recording.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["recording"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.Replay != nil {
			xaddr := fixXAddr(extension.Replay.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["replay"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
		if extension.Search != nil {
			xaddr := fixXAddr(extension.Search.XAddr, onvifDevice.xaddr)
			onvifDevice.endpoints["search"] = NewOnvifClient(xaddr, &onvifDevice.auth)
		}
	}

	onvifDevice.Capabilities = capabilities

	return nil
}
