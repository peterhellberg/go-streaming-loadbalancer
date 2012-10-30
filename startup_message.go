package main

import "fmt"

func printStartupMessage() {
	fmt.Println("Loadbalancer is starting up\n")
	fmt.Printf("port: %d\n", *port)
	fmt.Printf("geoipPath: %s\n", *geoipPath)
	fmt.Printf("redisKey: %s\n", *redisKey)
	fmt.Printf("redisServer: %s\n\n", *redisServer)
}
