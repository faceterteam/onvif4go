package onvif4go

import (
	tev "github.com/atagirov/onvif4go/events"
)

type PullPointSubscription struct {
	Client onvifCaller
}

func NewPullPointSubscription(endpoint string, onvifAuth *onvifAuth) *PullPointSubscription {
	return &PullPointSubscription{
		Client: NewOnvifClient(endpoint, onvifAuth),
	}
}

// PullMessages pulls one or more messages from a PullPoint.
// The device shall provide the following PullMessages command for all SubscriptionManager
// endpoints returned by the CreatePullPointSubscription command. This method shall not wait until
// the requested number of messages is available but return as soon as at least one message is available.
//
// The command shall at least support a Timeout of one minute. In case a device supports retrieval
// of less messages than requested it shall return these without generating a fault.
func (s *PullPointSubscription) PullMessages() (res tev.PullMessagesResponse, err error) {
	err = s.Client.Call(tev.PullMessages{}, &res)
	return
	// <wsdl:fault name="PullMessagesFaultResponse" message="tev:PullMessagesFaultResponse" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/PullMessages/Fault/PullMessagesFaultResponse"/>
}

/*
<wsdl:portType name="PullPointSubscription">
	<wsdl:operation name="Seek">
		<wsdl:documentation>
			This method readjusts the pull pointer into the past.
			A device supporting persistent notification storage shall provide the
			following Seek command for all SubscriptionManager endpoints returned by
			the CreatePullPointSubscription command. The optional Reverse argument can
			be used to reverse the pull direction of the PullMessages command.<br/>
			The UtcTime argument will be matched against the UtcTime attribute on a
			NotificationMessage.
		</wsdl:documentation>
		<wsdl:input message="tev:SeekRequest" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/SeekRequest"/>
		<wsdl:output message="tev:SeekResponse" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/SeekResponse"/>
	</wsdl:operation>
	<wsdl:operation name="SetSynchronizationPoint">
		<wsdl:documentation>Properties inform a client about property creation, changes and
			deletion in a uniform way. When a client wants to synchronize its properties with the
			properties of the device, it can request a synchronization point which repeats the current
			status of all properties to which a client has subscribed. The PropertyOperation of all
			produced notifications is set to “Initialized”. The Synchronization Point is
			requested directly from the SubscriptionManager which was returned in either the
			SubscriptionResponse or in the CreatePullPointSubscriptionResponse. The property update is
			transmitted via the notification transportation of the notification interface. This method is mandatory.
		</wsdl:documentation>
		<wsdl:input message="tev:SetSynchronizationPointRequest" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/SetSynchronizationPointRequest"/>
		<wsdl:output message="tev:SetSynchronizationPointResponse" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/SetSynchronizationPointResponse"/>
	</wsdl:operation>
	<wsdl:operation name="Unsubscribe">
		<wsdl:documentation>The device shall provide the following Unsubscribe command for all SubscriptionManager endpoints returned by the CreatePullPointSubscription command.<br/>
			This command shall terminate the lifetime of a pull point.
		</wsdl:documentation>
		<wsdl:input  name="UnsubscribeRequest" message="wsntw:UnsubscribeRequest" />
		<wsdl:output name="UnsubscribeResponse" message="wsntw:UnsubscribeResponse" />
		<wsdl:fault  name="ResourceUnknownFault" message="wsrf-rw:ResourceUnknownFault" />
		<wsdl:fault  name="UnableToDestroySubscriptionFault" message="wsntw:UnableToDestroySubscriptionFault" />
	</wsdl:operation>
</wsdl:portType>
*/
