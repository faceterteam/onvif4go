package onvif4go

import (
	tev "github.com/atagirov/onvif4go/events"
)

type EventsService struct {
	Client    onvifCaller
	onvifAuth *onvifAuth
}

func NewEventsService(endpoint string, onvifAuth *onvifAuth) *EventsService {
	return &EventsService{
		Client:    NewOnvifClient(endpoint, onvifAuth),
		onvifAuth: onvifAuth,
	}
}

// GetServiceCapabilities returns the capabilities of the event service.
func (s *EventsService) GetServiceCapabilities() (res tev.GetServiceCapabilitiesResponse, err error) {
	err = s.Client.Call(tev.GetServiceCapabilities{}, &res)
	return
}

/*
GetEventProperties returns information about the FilterDialects, Schema files and
topics supported by the device.

The WS-BaseNotification specification defines a set of OPTIONAL WS-ResouceProperties.
This specification does not require the implementation of the WS-ResourceProperty interface.
Instead, the subsequent direct interface shall be implemented by an ONVIF compliant device
in order to provide information about the FilterDialects, Schema files and topics supported by
the device.
*/
func (s *EventsService) GetEventProperties() (res tev.GetEventPropertiesResponse, err error) {
	err = s.Client.Call(tev.GetEventProperties{}, &res)
	return
}

/*
CreatePullPointSubscription returns a PullPointSubscription that can be polled using PullMessages.
This message contains the same elements as the SubscriptionRequest of the WS-BaseNotification
without the ConsumerReference.

If no Filter is specified the pullpoint notifies all occurring events to the client.
*/
func (s *EventsService) CreatePullPointSubscription(filter string, changeOnly bool) (service *PullPointSubscription, err error) {
	var res tev.CreatePullPointSubscriptionResponse
	err = s.Client.Call(tev.CreatePullPointSubscription{
		Filter: tev.FilterType(filter),
		SubscriptionPolicy: &tev.SubscriptionPolicy{
			ChangedOnly: changeOnly,
		},
	}, &res)
	if err != nil {
		return
	}
	/*
		<wsdl:fault name="ResourceUnknownFault" message="wsrf-rw:ResourceUnknownFault"/>
		<wsdl:fault name="InvalidFilterFault" message="wsntw:InvalidFilterFault"/>
		<wsdl:fault name="TopicExpressionDialectUnknownFault" message="wsntw:TopicExpressionDialectUnknownFault"/>
		<wsdl:fault name="InvalidTopicExpressionFault" message="wsntw:InvalidTopicExpressionFault"/>
		<wsdl:fault name="TopicNotSupportedFault" message="wsntw:TopicNotSupportedFault"/>
		<wsdl:fault name="InvalidProducerPropertiesExpressionFault" message="wsntw:InvalidProducerPropertiesExpressionFault"/>
		<wsdl:fault name="InvalidMessageContentExpressionFault" message="wsntw:InvalidMessageContentExpressionFault"/>
		<wsdl:fault name="UnacceptableInitialTerminationTimeFault" message="wsntw:UnacceptableInitialTerminationTimeFault"/>
		<wsdl:fault name="UnrecognizedPolicyRequestFault" message="wsntw:UnrecognizedPolicyRequestFault"/>
		<wsdl:fault name="UnsupportedPolicyRequestFault" message="wsntw:UnsupportedPolicyRequestFault"/>
		<wsdl:fault name="NotifyMessageNotSupportedFault" message="wsntw:NotifyMessageNotSupportedFault"/>
		<wsdl:fault name="SubscribeCreationFailedFault" message="wsntw:SubscribeCreationFailedFault"/>
	*/

	service = NewPullPointSubscription(res, s.onvifAuth)
	return
}
