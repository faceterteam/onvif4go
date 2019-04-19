package main

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/atagirov/onvif4go"
	"github.com/atagirov/onvif4go/xsd"
)

func main() {
	camera := onvif4go.NewOnvifDevice("10.110.3.235:8899")
	camera.Auth("admin", "12345678q")

	err := camera.Initialize()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	di, err := camera.Device.GetDeviceInformation()
	fmt.Println(di)
	users, err := camera.Device.GetUsers()
	fmt.Println(users)

	//camera.Media.Client.SetLogger(logClient, logClient)
	profiles, _ := camera.Media.GetProfiles()
	for _, profile := range profiles.Profiles {
		fmt.Println(profile.VideoEncoderConfiguration.SessionTimeout.String())
	}

	//camera.Events.Client.SetLogger(logClient, logClient)
	pullPoint, err := camera.Events.CreatePullPointSubscription("", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pullPointTicker := time.NewTicker(time.Second)
	go func() {
		//pullPoint.Client.SetLogger(logClient, logClient)

		for _ = range pullPointTicker.C {
			//fmt.Println("Tick at", t)

			messages, err := pullPoint.PullMessages(time.Second, 0)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for _, message := range messages.NotificationMessages {

				if message.Topic.Value == "tns1:RuleEngine/CellMotionDetector/Motion" {
					fmt.Println(messages.CurrentTime, message.Topic.Value, message.Message.Childs[0].Childs[1].Childs[0].Attrs[1].Value)
				} else {
					m := xsd.AnyType(message.Message)
					mergeNamespaces(&m)

					mes, _ := xml.MarshalIndent(m, "", "  ")
					fmt.Println(string(mes))
				}
			}
		}

		pullPoint.Unsubscribe()
	}()

	time.Sleep(60 * time.Minute)
	pullPointTicker.Stop()
}

func mergeNamespaces(x *xsd.AnyType) {
	x.Attrs = append(x.Attrs, xml.Attr{
		Name: xml.Name{
			Space: "xmlns",
			Local: "tt",
		},
		Value: "http://www.onvif.org/ver10/schema",
	})

	mergeNamespacesInner(x)
}

func mergeNamespacesInner(x *xsd.AnyType) {
	if x.XMLName.Space == "http://www.onvif.org/ver10/schema" {
		x.XMLName.Space = ""
		x.XMLName.Local = "tt:" + x.XMLName.Local
	}

	for idx := range x.Childs {
		mergeNamespacesInner(&x.Childs[idx])
	}
}

func logClient(message string) {
	fmt.Println(message)
}
