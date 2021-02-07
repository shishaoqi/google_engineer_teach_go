package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func main(){
	db, err := geoip2.Open("./geoip2/GeoLite2-City.mmdb")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	ipType := net.ParseIP("59.57.13.156")

	record, err := db.City(ipType)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("city name: %v\n", record.City.Names)
	if len(record.Subdivisions) > 0 {
		fmt.Println("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Println("country name: %v\n", record.Country.Names["en"])
	fmt.Println("Time zone: %v\n", record.Location.TimeZone)

	fmt.Println("lat, lng: ", record.Location.Latitude, record.Location.Longitude)

}
