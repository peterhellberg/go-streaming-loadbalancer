package main

func main() {
	// Begin by parsing the command line arguments
	parseFlags()

	// Load the GeoIP.dat
	loadGeoip()

	// Create a connection pool for Redis
	createRedisPool()

	// Print the startup message
	printStartupMessage()

	// Finally start the server
	startServer()
}
