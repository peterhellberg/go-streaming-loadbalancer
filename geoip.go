package main

import "log"
import "github.com/carlhoerberg/go-geoip"

var gi *geoip.GeoIP

func loadGeoip() {
	var err error

	gi, err = geoip.Load(*geoipPath)

	if err != nil {
		log.Fatal("Could not load GeoIP.dat: ", err.Error())
	}
}
