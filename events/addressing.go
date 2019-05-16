package events

import "encoding/xml"

type addressingAction struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing Action"`
	Value   string   `xml:",chardata"`
}

type addressingMessageID struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing MessageID"`
	Value   string   `xml:",chardata"` // urn:uuid:5e6558cf-e5dc-42e5-9417-6372b3ed9a65
}

type addressingReplyTo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ReplyTo"`
	Address string   `xml:"http://www.w3.org/2005/08/addressing Address"` // http://www.w3.org/2005/08/addressing/anonymous
}

type addressingTo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing To"`
	Value   string   `xml:",chardata"`
}

func MakeAnonymousAddressingHeaders(action, to string) []interface{} {
	return []interface{}{ 
		addressingAction { 
			Value: action,
		},
		addressingReplyTo {
			Address: "http://www.w3.org/2005/08/addressing/anonymous",
		},
		addressingTo {
			Value: to,
		},
	}
}

// <a:Action s:mustUnderstand="1">http://www.onvif.org/ver10/events/wsdl/PullPointSubscription/PullMessagesRequest</a:Action>
// <a:MessageID>urn:uuid:5e6558cf-e5dc-42e5-9417-6372b3ed9a65</a:MessageID>
// <a:ReplyTo>
//   <a:Address>http://www.w3.org/2005/08/addressing/anonymous</a:Address>
// </a:ReplyTo>
// <a:To s:mustUnderstand="1">http://10.110.3.252/onvif/event/subsription_4352</a:To>
//   <a:Address>http://www.w3.org/2005/08/addressing/anonymous</a:Address>
// </a:ReplyTo>
// <a:To s:mustUnderstand="1">http://10.110.3.252/onvif/event/subsription_4352</a:To>
