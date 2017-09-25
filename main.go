package main

import (
	"web-push-go/webpush"
	"fmt"
)

func main() {

	key256 := "BClMhqu1+FmXeruyMqNfiifYWtCXG5Df4rBomMTbS+f9gZ+3L3bRbxL20y11/WFhZN6+II8vQkbOltk9GvIedNI="
	auth   := "OGtdsqxVRRLOT3/nEv6V/Q=="

	// The values that make up the Subscription struct come from the browser

	sub := &webpush.Subscription{
		Endpoint:"https://gcm-http.googleapis.com/gcm/send",
		Key:  []byte(key256),
		Auth: []byte(auth)}

	res , err := webpush.Send(nil, sub, "hi", "AIzaSyAYLY9qTgd8M4t7yFXsGKJFtEqwp4YasfI")

	if err != nil {
		fmt.Println(err)
		return
	}else{
	fmt.Println(res)
		return
	}

}