package onvif

import (
	"github.com/atagirov/onvif4go/xsd"
)

type Message struct {
	UtcTime           xsd.DateTime      `xml:"UtcTime,attr"`
	PropertyOperation *string           `xml:"PropertyOperation,attr"`
	Source            *ItemList         `xml:"http://www.onvif.org/ver10/schema Source"`
	Key               *ItemList         `xml:"http://www.onvif.org/ver10/schema Key"`
	Data              *ItemList         `xml:"http://www.onvif.org/ver10/schema Data"`
	Extension         *MessageExtension `xml:"http://www.onvif.org/ver10/schema Extension"`
}

type MessageExtension xsd.AnyType
