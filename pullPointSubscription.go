package onvif4go

import (
	"time"

	tev "github.com/atagirov/onvif4go/events"
	"github.com/atagirov/onvif4go/xsd"
)

type PullPointSubscription struct {
	Client   onvifCaller
	Settings tev.CreatePullPointSubscriptionResponse
}

func NewPullPointSubscription(settings tev.CreatePullPointSubscriptionResponse, onvifAuth *onvifAuth) *PullPointSubscription {
	return &PullPointSubscription{
		Client:   NewOnvifClient(string(settings.SubscriptionReference.Address.Content), onvifAuth),
		Settings: settings,
	}
}

/*
PullMessages pulls one or more messages from a PullPoint.
The device shall provide the following PullMessages command for all SubscriptionManager
endpoints returned by the CreatePullPointSubscription command. This method shall not wait until
the requested number of messages is available but return as soon as at least one message is available.

The command shall at least support a Timeout of one minute. In case a device supports retrieval
of less messages than requested it shall return these without generating a fault.
*/
func (s *PullPointSubscription) PullMessages(timeout time.Duration, messageLimit int) (res tev.PullMessagesResponse, err error) {
	err = s.Client.Call(tev.PullMessages{
		Timeout:      xsd.Duration(timeout),
		MessageLimit: messageLimit,
	}, &res)
	return
	// <wsdl:fault name="PullMessagesFaultResponse" message="tev:PullMessagesFaultResponse" wsaw:Action="http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/PullMessages/Fault/PullMessagesFaultResponse"/>
}

/*
Unsubscribe shall terminate the lifetime of a pull point
*/
func (s *PullPointSubscription) Unsubscribe() (err error) {
	var res tev.UnsubscribeResponse
	err = s.Client.Call(tev.UnsubscribeRequest{}, &res)
	return
	// <wsdl:fault  name="ResourceUnknownFault" message="wsrf-rw:ResourceUnknownFault" />
	// <wsdl:fault  name="UnableToDestroySubscriptionFault" message="wsntw:UnableToDestroySubscriptionFault" />
}

/*
Seek readjusts the pull pointer into the past.
A device supporting persistent notification storage shall provide the
following Seek command for all SubscriptionManager endpoints returned by
the CreatePullPointSubscription command. The optional Reverse argument can
be used to reverse the pull direction of the PullMessages command.

The UtcTime argument will be matched against the UtcTime attribute on a
NotificationMessage.
*/
func (s *PullPointSubscription) Seek(time time.Time, reverse bool) (err error) {
	var res tev.SeekResponse
	err = s.Client.Call(tev.Seek{
		UtcTime: xsd.MakeDateTime(time),
		Reverse: reverse,
	}, &res)
	return
}

/*
SetSynchronizationPoint inform a client about property creation, changes and
deletion in a uniform way. When a client wants to synchronize its properties with the
properties of the device, it can request a synchronization point which repeats the current
status of all properties to which a client has subscribed. The PropertyOperation of all
produced notifications is set to “Initialized”. The Synchronization Point is
requested directly from the SubscriptionManager which was returned in either the
SubscriptionResponse or in the CreatePullPointSubscriptionResponse. The property update is
transmitted via the notification transportation of the notification interface. This method is mandatory.
*/
func (s *PullPointSubscription) SetSynchronizationPoint(time time.Time, reverse bool) (err error) {
	var res tev.SetSynchronizationPointResponse
	err = s.Client.Call(tev.SetSynchronizationPoint{}, &res)
	return
}
