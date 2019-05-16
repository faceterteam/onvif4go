package xsd

import (
	"encoding/xml"
)

type AnyType struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Childs  []AnyType  `xml:",any"`
	Text    string     `xml:",chardata"`
}
