package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/atagirov/onvif4go"
	tev "github.com/atagirov/onvif4go/events"
)

func main() {
	camera := onvif4go.NewOnvifDevice("10.110.3.235:8899") // OK

	camera.Auth("admin", "12345678q")

	err := camera.Initialize()
	if err != nil {
		fmt.Println("camera.Initialize error:", err.Error())
		return
	}

	di, err := camera.Device.GetDeviceInformation()
	fmt.Println(di)
	users, err := camera.Device.GetUsers()
	fmt.Println(users)

	//camera.Media.Client.SetLogger(logClient, logClient)
	// profiles, _ := camera.Media.GetProfiles()
	// for _, profile := range profiles.Profiles {
	// 	fmt.Println(profile.VideoEncoderConfiguration.SessionTimeout.String())
	// }

	eventServiceCapabilities, err := camera.Events.GetServiceCapabilities()
	if err != nil {
		fmt.Println("Events.GetServiceCapabilities error:", err.Error())
		//return
	} else {
		fmt.Println(eventServiceCapabilities.Capabilities)
	}

	//camera.Events.Client.SetLogger(logClient, logClient)

	pullPoint, err := camera.Events.CreatePullPointSubscription("", false, tev.NewRelativeTimeType(time.Second*600))
	if err != nil {
		fmt.Println("Events.CreatePullPointSubscription error:", err.Error())
		return
	}

	fmt.Println(pullPoint.Settings.SubscriptionReference.Address.Content)

	//pullPoint.Client.SetLogger(logClient, logClient)

	pullPointTicker := time.NewTicker(time.Second * 30)

	go func() {

		//currentTime := time.Now()
		currentState := false

		for {
			select {
			case _ = <-pullPointTicker.C:
				//fmt.Println("SetSynchronizationPoint", t)

				err := pullPoint.SetSynchronizationPoint()
				if err != nil {
					fmt.Println("pullPoint.SetSynchronizationPoint error:", err.Error())
				}
			default:
				messages, err := pullPoint.PullMessages(time.Second*5, 1024)
				if err != nil {
					fmt.Println("pullPoint.PullMessages error:", err.Error())
					return
				}

				//fmt.Println(messages.CurrentTime, messages.TerminationTime, len(messages.NotificationMessages))

				for _, message := range messages.NotificationMessages {
					if message.Topic.Value == "tns1:RuleEngine/CellMotionDetector/Motion" {
						for _, onvifMessage := range message.Message.Messages {
							isMotion, _ := strconv.ParseBool(onvifMessage.Data.SimpleItems["IsMotion"])
							if currentState != isMotion {
								currentState = isMotion
								//currentTime = onvifMessage.UtcTime

								fmt.Println(
									message.Topic.Value,
									*onvifMessage.PropertyOperation,
									onvifMessage.UtcTime,
									onvifMessage.Data.SimpleItems["IsMotion"])
							}
						}
					} else if message.Topic.Value == "tns1:VideoSource/MotionAlarm" {
						for _, onvifMessage := range message.Message.Messages {
							state, _ := strconv.ParseBool(onvifMessage.Data.SimpleItems["State"])

							if currentState != state {
								currentState = state
								//currentTime = onvifMessage.UtcTime

								fmt.Println(
									message.Topic.Value,
									*onvifMessage.PropertyOperation,
									onvifMessage.UtcTime,
									onvifMessage.Data.SimpleItems["State"])
							}
						}
					} else {
						fmt.Println(message.Topic.Value)
					}
				}
			}
		}

		pullPoint.Unsubscribe()
	}()

	time.Sleep(60 * time.Minute)
}

func logClient(message string) {
	fmt.Println(message)
}
