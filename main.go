package main

func main() {
	// Begin by parsing the command line arguments
  parse_flags()

  // Load the GeoIP.dat
	load_geoip()

  // Create a connection pool for Redis
	new_redis_pool()

  // Print the startup message
	print_startup_message()

  // Finally start the server
	start_server()
}
