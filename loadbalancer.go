package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/carlhoerberg/go-geoip"
	"github.com/garyburd/redigo/redis"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	defaultPort      = 1983
	defaultGeoIpPath = "/usr/local/share/GeoIP/GeoIP.dat"
	defaultRedisKey  = "redirect_ip"
)

var redis_key = flag.String("redis_key", defaultRedisKey, "The Redis key to use")
var geoip_path = flag.String("geoip", defaultGeoIpPath, "Full path of the GeoIP.dat")
var port = flag.Int("port", defaultPort, "The port number")

var gi *geoip.GeoIP
var gi_err error

func main() {
	parse_flags()
	load_geoip()
	start_server()
}

func parse_flags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(2)
	}

	flag.Parse()
}

func load_geoip() {
	gi, gi_err = geoip.Load(*geoip_path)

	if gi_err != nil {
		fmt.Println("Could not load GeoIP.dat:\n", gi_err.Error())
		os.Exit(2)
	}
}

func start_server() {
	lb_endpoint := "/loadbalancer.json"

	m := pat.New()
	m.Get(lb_endpoint, http.HandlerFunc(LoadbalancerServer))
	m.Get("/", http.RedirectHandler(lb_endpoint, http.StatusMovedPermanently))

	startup_message()

	http.Handle("/", m)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func startup_message() {
	tpl := "Loadbalancer starting on port: %d, redis_key: %s, geoip_path: %s"
	println(fmt.Sprintf(tpl, *port, *redis_key, *geoip_path))
}

type LoadbalancerResponse struct {
	Redirect    string `json:"redirect"`
	ClientIP    string `json:"client_ip"`
	CountryCode string `json:"country_code"`
	Timestamp   int64  `json:"timestamp"`
}

func LoadbalancerServer(w http.ResponseWriter, req *http.Request) {
	// Set some headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "go-streaming-loadbalancer")

	// Connect to Redis
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Make sure that the Redis connection is closed
	defer c.Close()

	// Retrieve the redirect key
	redirect_ip, err := redis.String(c.Do("GET", *redis_key))

	if err != nil {
		fmt.Println("Error:", err)
	}

	ip := ClientIP(req)
	cc := ""
	loc := gi.GetLocationByIP(ip)

	if loc != nil {
		cc = loc.CountryCode
	}

	// Populate the loadbalancer response struct
	response := LoadbalancerResponse{
		Redirect:    redirect_ip,
		Timestamp:   time.Now().Unix(),
		CountryCode: cc,
		ClientIP:    ip,
	}

	// Marshal the JSON
	json, err := json.Marshal(response)

	if err != nil {
		fmt.Println("JSON Error:", err)
	}

	// Return the JSON to the client
	io.WriteString(w, string(json))
}

func ClientIP(req *http.Request) string {
	return strings.Split(req.RemoteAddr, ":")[0]
}
