package main

import "fmt"

func print_startup_message() {
	fmt.Println("Loadbalancer is starting up\n")
	fmt.Printf("port: %d\n", *port)
	fmt.Printf("geoip_path: %s\n", *geoip_path)
	fmt.Printf("redis_key: %s\n", *redis_key)
	fmt.Printf("redis_server: %s\n\n", *redis_server)
}
