package onvif4go

import (
	"time"

	"github.com/atagirov/onvif4go/soap"
)

type onvifCaller interface {
	Call(request, response interface{}) error

	CallWithoutAuth(request, response interface{}) error
}

type onvifAuth struct {
	login    string
	password string
	timeDiff time.Duration
}

type onvifClient struct {
	soapClient *soap.SoapClient
	auth       *onvifAuth
}

func NewOnvifClient(endpoint string, auth *onvifAuth) *onvifClient {
	return &onvifClient{
		soapClient: soap.NewSoapClient(endpoint),
		auth:       auth,
	}
}

func (c *onvifClient) callOnInternal(request, response interface{}, useAuth bool) error {
	if c.auth.login != "" && useAuth {
		return c.soapClient.Do(request, response,
			soap.MakeWSSecurity(c.auth.login, c.auth.password, c.auth.timeDiff))
	}

	return c.soapClient.Do(request, response)
}

func (c *onvifClient) Call(request, response interface{}) error {
	return c.callOnInternal(request, response, true)
}

func (c *onvifClient) CallWithoutAuth(request, response interface{}) error {
	return c.callOnInternal(request, response, false)
}
