package main

import (
	"web-push-go/webpush"
	"fmt"
	"encoding/base64"

)

func main() {

	//key256 := "BClMhqu1+FmXeruyMqNfiifYWtCXG5Df4rBomMTbS+f9gZ+3L3bRbxL20y11/WFhZN6+II8vQkbOltk9GvIedNI="
	//auth   := "OGtdsqxVRRLOT3/nEv6V/Q=="

	key256 , _ := base64.StdEncoding.DecodeString("BClMhqu1+FmXeruyMqNfiifYWtCXG5Df4rBomMTbS+f9gZ+3L3bRbxL20y11/WFhZN6+II8vQkbOltk9GvIedNI=")
	auth, _ := base64.StdEncoding.DecodeString("OGtdsqxVRRLOT3/nEv6V/Q==")

	//// The values that make up the Subscription struct come from the browser
	//var exampleJSON string
	// exampleJSON = `{"endpoint": "https://gcm-http.googleapis.com/gcm/send", "keys": {"p256dh": "BClMhqu1+FmXeruyMqNfiifYWtCXG5Df4rBomMTbS+f9gZ+3L3bRbxL20y11/WFhZN6+II8vQkbOltk9GvIedNI=", "auth": "OGtdsqxVRRLOT3/nEv6V/Q=="}}`
	//
	//sub, err := webpush.SubscriptionFromJSON([]byte(exampleJSON))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}


	sub := &webpush.Subscription{
		Endpoint:"https://gcm-http.googleapis.com/gcm/send",
		Key:  key256,
		Auth: auth}

	res , err := webpush.Send(nil, sub, "hii", "AIzaSyAYLY9qTgd8M4t7yFXsGKJFtEqwp4YasfI")

	if err != nil {
		fmt.Println(err)
		return
	}else{
	fmt.Println(res)
		return
	}

}