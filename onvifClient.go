package onvif4go

import (
	"time"

	"github.com/atagirov/onvif4go/soap"
)

type onvifCaller interface {
	Call(request, response interface{}, headers ...interface{}) error

	WithoutAuth() onvifCaller

	SetLogger(logRequest, logResponse func(message string))
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

func (c *onvifClient) Call(request, response interface{}, headers ...interface{}) error {
	if c.auth.login != "" {
		headers = append(headers, soap.MakeWSSecurity(c.auth.login, c.auth.password, c.auth.timeDiff))
	}

	return c.soapClient.Do(request, response, headers...)
}

func (c *onvifClient) SetLogger(logRequest, logResponse func(message string)) {
	c.soapClient.LogRequest = logRequest
	c.soapClient.LogResponse = logResponse
}

func (c *onvifClient) WithoutAuth() onvifCaller {
	return &onvifClient{
		soapClient: c.soapClient,
		auth:       &onvifAuth{},
	}
}
