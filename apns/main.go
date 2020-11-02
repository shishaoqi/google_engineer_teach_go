package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

func main() {

	wd, _ := os.Getwd()
	certFile := wd + "/apns_dev.p12"
	cert, err := certificate.FromP12File(certFile, "tongbu.com")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "f6dfb3f0f6a1040b025950fdb0dc2a5b2f7df693ad0d1a333d4c5f80fd4a7bca"
	notification.Topic = "com.sideshow.Apns2"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below

	// If you want to test push notifications for builds running directly from XCode (Development), use
	// client := apns2.NewClient(cert).Development()
	// For apps published to the app store or installed as an ad-hoc distribution use Production()

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}