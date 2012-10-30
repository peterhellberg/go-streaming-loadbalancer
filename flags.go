package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultPort        = 1983
	defaultGeoIpPath   = "/usr/local/share/GeoIP/GeoIP.dat"
	defaultRedisServer = "127.0.0.1:6379"
	defaultRedisKey    = "redirect_ip"
)

var port = flag.Int("port", defaultPort, "The port number")
var geoip_path = flag.String("geoip_path", defaultGeoIpPath, "Full path of the GeoIP.dat")
var redis_server = flag.String("redis_server", defaultRedisServer, "The Redis server to use")
var redis_key = flag.String("redis_key", defaultRedisKey, "The Redis key to use")

func parse_flags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(2)
	}

	flag.Parse()
}
