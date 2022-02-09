package stubs

import (
	"aav_logger/app/models"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
)

func PostNewFlights() {
	// Create new request object
	request := gorequest.New()

	// Open a handle to the flightlog
	csvFile, err := os.Open("../uploads/csv/test_flights.csv")
	if err != nil {
		// TODO: Log the error
		panic(err)
	}

	// Create new csv reader
	r := csv.NewReader(csvFile)

	// Pack all records into an exec buffer
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		flight := parseRecord(record)
		fmt.Printf("%+v\n", flight)
		_, body, _ := request.Post("http://127.0.0.1:9000/api/v1/flight").
			Set("Content-Type", "application/json").
			Set("Accept", "application/json").
			Send(fmt.Sprintf(
				`{"robot":"%s","generation":%d,"start":"%s","stop":"%s","lat":%0.6f,"lon":%0.6f}`,
				flight.Robot,
				flight.Generation,
				flight.Start,
				flight.Stop,
				flight.Lat,
				flight.Lon)).End()
		fmt.Println(body)
	}
}

func parseRecord(record []string) models.Flight {
	// Big assumptions incoming!
	// Record follows format: droneName,generation,start_time,stop_time,lat,lon

	// Per spec, timestamps are UTC. Use RFC3339 layout
	layout := "2006-01-02T15:04:05Z07:00"
	drone := record[0]
	generation, _ := strconv.ParseInt(record[1], 10, 64)
	start, _ := time.Parse(layout, record[2])
	stop, _ := time.Parse(layout, record[3])
	lat, _ := strconv.ParseFloat(record[4], 64)
	lon, _ := strconv.ParseFloat(record[5], 64)

	// Return flight object
	return models.Flight{
		Robot:      drone,
		Generation: generation,
		Start:      start,
		Stop:       stop,
		Lat:        lat,
		Lon:        lon,
	}
}
