package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type SoapClient struct {
	Endpoint string
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

	httpClient := new(http.Client)
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
	//fmt.Println(respBody)

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
