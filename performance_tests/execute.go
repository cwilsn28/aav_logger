package main

import (
	"aav_logger/performance_tests/stubs"
)

//------------------------------------------------------------------------------
// Company endpoint tests.Tests
// -----------------------------------------------------------------------------

func TestNewFlightPost() {
	stubs.PostNewFlights()
}

func main() {
	TestNewFlightPost()
}
