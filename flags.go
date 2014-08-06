package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultPort        = 1983
	defaultGeoIPPath   = "/usr/local/share/GeoIP/GeoIP.dat"
	defaultRedisServer = "127.0.0.1:6379"
	defaultRedisKey    = "redirect_ip"
)

var port = flag.Int("port", defaultPort, "The port number")
var geoipPath = flag.String("geoip_path", defaultGeoIPPath, "Full path of the GeoIP.dat")
var redisServer = flag.String("redis_server", defaultRedisServer, "The Redis server to use")
var redisKey = flag.String("redis_key", defaultRedisKey, "The Redis key to use")

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(2)
	}

	flag.Parse()
}
