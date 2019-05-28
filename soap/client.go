package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SoapClient struct {
	Endpoint    string
	LogRequest  func(request string)
	LogResponse func(response string)
}

func NewSoapClient(endpoint string) *SoapClient {
	return &SoapClient{
		Endpoint: endpoint,
	}
}

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (soap *SoapClient) Do(request, response interface{}, headers ...interface{}) error {
	envelope := NewSOAPEnvelope(request)
	for _, header := range headers {
		envelope.AddHeader(header)
	}

	message, err := xml.Marshal(envelope)
	if err != nil {
		return err
	}

	if soap.LogRequest != nil {
		soap.LogRequest(string(message))
	}

	httpClient := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := httpClient.Post(soap.Endpoint, "application/soap+xml; charset=utf-8", bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusInternalServerError &&
		resp.StatusCode != http.StatusBadRequest {

		//respBody := readResponse(resp)

		return errors.New("camera is not available at " + soap.Endpoint + " or it does not support ONVIF services")
	}

	responseEnvelope := NewSOAPEnvelope(response)

	respBody := readResponse(resp)

	if soap.LogResponse != nil {
		soap.LogResponse(string(respBody))
	}

	decoder := xml.NewDecoder(strings.NewReader(respBody))
	err = decoder.Decode(responseEnvelope)
	if err != nil {
		return err
	}
	if responseEnvelope.Body.Fault != nil {
		return responseEnvelope.Body.Fault
	}

	return nil
}
