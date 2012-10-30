package main

import "log"
import "github.com/carlhoerberg/go-geoip"

var gi *geoip.GeoIP

func load_geoip() {
	var err error

	gi, err = geoip.Load(*geoip_path)

	if err != nil {
		log.Fatal("Could not load GeoIP.dat: ", err.Error())
	}
}
