package onvif

import (
	"encoding/xml"

	"github.com/faceterteam/onvif4go/xsd"
)

type ItemList struct {
	SimpleItems  map[string]string        `xml:"http://www.onvif.org/ver10/schema SimpleItem"`
	ElementItems map[string][]xsd.AnyType `xml:"http://www.onvif.org/ver10/schema ElementItem"`
	Extension    *ItemListExtension       `xml:"http://www.onvif.org/ver10/schema Extension"`
}

func (l *ItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*l = ItemList{}

	var (
		token xml.Token
		err   error
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Space != "http://www.onvif.org/ver10/schema" {
				return xml.UnmarshalError("expected tt:SimpleItem, tt:ElementItem... but receive")
			}

			if se.Name.Local == "SimpleItem" {
				simpleItem := &ItemListSimpleItem{}
				err = d.DecodeElement(simpleItem, &se)
				if err != nil {
					return err
				}
				if l.SimpleItems == nil {
					l.SimpleItems = make(map[string]string, 0)
				}
				l.SimpleItems[simpleItem.Name] = simpleItem.Value
			} else if se.Name.Local == "ElementItem" {
				elementItem := &ItemListElementItem{}
				err = d.DecodeElement(elementItem, &se)
				if err != nil {
					return err
				}
				if l.ElementItems == nil {
					l.ElementItems = make(map[string][]xsd.AnyType, 0)
				}
				l.ElementItems[elementItem.Name] = elementItem.Elements
			} else if se.Name.Local == "Extension" {
				err = d.DecodeElement(l.Extension, &se)
				if err != nil {
					return err
				}
			} else {
				return xml.UnmarshalError("expected tt:SimpleItem, tt:ElementItem... but receive")
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

type ItemListSimpleItem struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type ItemListElementItem struct {
	Name     string        `xml:"Name,attr"`
	Elements []xsd.AnyType `xml:",any"`
}

type ItemListExtension xsd.AnyType
