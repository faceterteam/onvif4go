package events

import (
	"time"

	tt "github.com/faceterteam/onvif4go/onvif"
	"github.com/faceterteam/onvif4go/xsd"
)

type FilterType string

// AbsoluteOrRelativeTimeType <xsd:union memberTypes="xsd:dateTime xsd:duration"/>
type AbsoluteOrRelativeTimeType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	duration *xsd.Duration
	dateTime *time.Time
}

// UnmarshalText unmarshal AbsoluteOrRelativeTimeType from text
func (v *AbsoluteOrRelativeTimeType) UnmarshalText(text []byte) error {
	str := string(text)
	if str[0] == 'P' {
		var duration xsd.Duration
		err := duration.UnmarshalText(text)
		if err != nil {
			return err
		}
		v.duration = &duration
	} else {
		t, err := time.Parse(time.RFC3339Nano, str)
		if err != nil {
			return err
		}
		v.dateTime = &t
	}

	return nil
}

// MarshalText marshal AbsoluteOrRelativeTimeType to text
func (v AbsoluteOrRelativeTimeType) MarshalText() ([]byte, error) {
	if v.duration == nil {
		return v.dateTime.MarshalText()
	}
	return v.duration.MarshalText()
}

// NewAbsoluteTimeType make *AbsoluteOrRelativeTimeType from time.Time
func NewAbsoluteTimeType(dateTime time.Time) *AbsoluteOrRelativeTimeType {
	return &AbsoluteOrRelativeTimeType{dateTime: &dateTime}
}

// NewRelativeTimeType make *AbsoluteOrRelativeTimeType from time.Duration
func NewRelativeTimeType(duration time.Duration) *AbsoluteOrRelativeTimeType {
	d := xsd.Duration(duration)
	return &AbsoluteOrRelativeTimeType{duration: &d}
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
	SubscriptionReference *SubscriptionReference //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Topic                 Topic
	ProducerReference     *ProducerReference
	Message               OnvifEventsMessage
}

type SubscriptionReference EndpointReferenceType
type Topic TopicExpressionType
type ProducerReference EndpointReferenceType

type OnvifEventsMessage struct {
	Messages []tt.Message `xml:"http://www.onvif.org/ver10/schema Message"`
}

type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect xsd.AnyURI `xml:"Dialect,attr"`
	Value   string     `xml:",chardata"`
}

//Event main types

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/events/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type CreatePullPointSubscription struct {
	XMLName                string                      `xml:"http://www.onvif.org/ver10/events/wsdl CreatePullPointSubscription"`
	Filter                 FilterType                  `xml:"http://www.onvif.org/ver10/events/wsdl Filter,omitempty"`
	InitialTerminationTime *AbsoluteOrRelativeTimeType `xml:"http://www.onvif.org/ver10/events/wsdl InitialTerminationTime,omitempty"`
	SubscriptionPolicy     *SubscriptionPolicy         `xml:"http://www.onvif.org/ver10/events/wsdl SubscriptionPolicy,omitempty"`
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
	MessageLimit int          `xml:"http://www.onvif.org/ver10/events/wsdl MessageLimit"`
}

type PullMessagesResponse struct {
	CurrentTime          CurrentTime
	TerminationTime      TerminationTime
	NotificationMessages []NotificationMessage `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationMessage"`
}

type PullMessagesFaultResponse struct {
	MaxTimeout      xsd.Duration
	MaxMessageLimit int
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

type UnsubscribeRequest struct {
	XMLName string `xml:"http://www.onvif.org/ver10/events/wsdl UnsubscribeRequest"`
}

type UnsubscribeResponse struct {
}
