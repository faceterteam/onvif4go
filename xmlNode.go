package onvif4go

import (
	"encoding/xml"
)

type XmlNode struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Childs  []XmlNode  `xml:",any"`
	Text    string     `xml:",chardata"`
}
