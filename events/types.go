package events

import (
	"github.com/atagirov/onvif4go/xsd"
)

type FilterType string

//<xsd:union memberTypes="xsd:dateTime xsd:duration"/>
type AbsoluteOrRelativeTimeType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	xsd.DateTime
	xsd.Duration
}

type SubscriptionPolicy struct { //tev http://www.onvif.org/ver10/events/wsdl
	ChangedOnly bool `xml:"ChangedOnly,attr"`
}

type Capabilities struct { //tev
	WSSubscriptionPolicySupport                   bool `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            bool `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers                      int  `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints                                 int  `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage                 bool `xml:"PersistentNotificationStorage,attr"`
}

type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address             AttributedURIType `xml:"http://www.w3.org/2005/08/addressing Address"`
	ReferenceParameters *ReferenceParametersType
	Metadata
}

type AttributedURIType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Content xsd.AnyURI `xml:",chardata"`
	//Here can be anyAttribute
}

type ReferenceParametersType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Any string
	//Here can be anyAttribute
}

type Metadata MetadataType //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

type MetadataType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Any string
	//Here can be anyAttribute
}

type CurrentTime xsd.DateTime     //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
type TerminationTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
type FixedTopicSet bool           //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

type TopicSet TopicSetType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

type TopicSetType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	ExtensibleDocumented
	//here can be any element
}

type ExtensibleDocumented struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Documentation Documentation //к xsd-документе documentation с маленькой буквы начинается
	//here can be anyAttribute
}

type Documentation xsd.AnyType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

type TopicExpressionDialect xsd.AnyURI

type NotificationMessage NotificationMessageHolderType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

type NotificationMessageHolderType struct {
	SubscriptionReference SubscriptionReference //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Topic                 Topic
	ProducerReference     ProducerReference
	Message               Message
}

type SubscriptionReference EndpointReferenceType
type Topic TopicExpressionType
type ProducerReference EndpointReferenceType
type Message xsd.AnyType

type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect xsd.AnyURI `xml:"Dialect,attr"`
}

//Event main types

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/events/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

//BUG(r) Bad AbsoluteOrRelativeTimeType type
type CreatePullPointSubscription struct {
	XMLName                string                     `xml:"http://www.onvif.org/ver10/events/wsdl CreatePullPointSubscription"`
	Filter                 FilterType                 `xml:"http://www.onvif.org/ver10/events/wsdl Filter"`
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"http://www.onvif.org/ver10/events/wsdl InitialTerminationTime"`
	SubscriptionPolicy     SubscriptionPolicy         `xml:"http://www.onvif.org/ver10/events/wsdl SubscriptionPolicy"`
}

type CreatePullPointSubscriptionResponse struct {
	SubscriptionReference EndpointReferenceType `xml:"http://www.onvif.org/ver10/events/wsdl SubscriptionReference"`
	CurrentTime           CurrentTime
	TerminationTime       TerminationTime
}

type ResourceUnknownFault struct {
}

type InvalidFilterFault struct {
}

type TopicExpressionDialectUnknownFault struct {
}

type InvalidTopicExpressionFault struct {
}

type TopicNotSupportedFault struct {
}

type InvalidProducerPropertiesExpressionFault struct {
}

type InvalidMessageContentExpressionFault struct {
}

type UnacceptableInitialTerminationTimeFault struct {
}

type UnrecognizedPolicyRequestFault struct {
}

type UnsupportedPolicyRequestFault struct {
}

type NotifyMessageNotSupportedFault struct {
}

type SubscribeCreationFailedFault struct {
}

type GetEventProperties struct {
	XMLName string `xml:"http://www.onvif.org/ver10/events/wsdl GetEventProperties"`
}

type GetEventPropertiesResponse struct {
	TopicNamespaceLocation          xsd.AnyURI
	FixedTopicSet                   FixedTopicSet
	TopicSet                        TopicSet
	TopicExpressionDialect          TopicExpressionDialect
	MessageContentFilterDialect     xsd.AnyURI
	ProducerPropertiesFilterDialect xsd.AnyURI
	MessageContentSchemaLocation    xsd.AnyURI
}

//Port type PullPointSubscription

type PullMessages struct {
	XMLName      string       `xml:"http://www.onvif.org/ver10/events/wsdl PullMessages"`
	Timeout      xsd.Duration `xml:"http://www.onvif.org/ver10/events/wsdl Timeout"`
	MessageLimit xsd.Int      `xml:"http://www.onvif.org/ver10/events/wsdl MessageLimit"`
}

type PullMessagesResponse struct {
	CurrentTime          CurrentTime
	TerminationTime      TerminationTime
	NotificationMessages []NotificationMessage `xml:"http://www.onvif.org/ver10/events/wsdl NotificationMessage"`
}

type PullMessagesFaultResponse struct {
	MaxTimeout      xsd.Duration
	MaxMessageLimit xsd.Int
}

type Seek struct {
	XMLName string       `xml:"http://www.onvif.org/ver10/events/wsdl Seek"`
	UtcTime xsd.DateTime `xml:"http://www.onvif.org/ver10/events/wsdl UtcTime"`
	Reverse bool         `xml:"http://www.onvif.org/ver10/events/wsdl Reverse"`
}

type SeekResponse struct {
}

type SetSynchronizationPoint struct {
	XMLName string `xml:"http://www.onvif.org/ver10/events/wsdl SetSynchronizationPoint"`
}

type SetSynchronizationPointResponse struct {
}
