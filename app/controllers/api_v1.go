package controllers

import (
	"aav_logger/app/models"
	"aav_logger/app/transactions"
	"encoding/json"
	"fmt"

	"github.com/revel/revel"
)

type APIV1 struct {
	App
}

/* ---
 * Add a single flight record
 * --- */
func (c APIV1) NewFlight() revel.Result {
	if c.Request.Method == "POST" {
		// Create new flight record and bind the JSON to it
		var newFlight models.Flight
		c.Params.BindJSON(&newFlight)

		// Insert the flight record
		recordID, err := transactions.InsertFlight(DBCONN, newFlight)
		if err != nil {
			// TODO: Log the error
			return ServerError("A server error occurred")
		}

		// Send back a copy of the record from the database
		flight, err := transactions.FlightWithID(DBCONN, recordID)
		if err != nil {
			// TODO: Log the error
			return ServerError("A server error occurred")
		}

		responseJSON, _ := json.Marshal(flight)
		return Created(string(responseJSON))
	}
	return MethodNotAllowed("")
}

/* ---
 * Add flight records in bulk
 * --- */
func (c APIV1) NewFlightBulk() revel.Result {
	if c.Request.Method == "POST" {

		// For now, test the endpoint with a file currently on disk.
		logfile := "uploads/csv/test_flights.csv"
		insertCount, err := transactions.InsertFlightBulk(DBCONN, logfile)
		if err != nil {
			fmt.Println(err)
			return ServerError("A server error occurred")
		}

		responseJSON, _ := json.Marshal(map[string]int64{"records_inserted": insertCount})
		return Created(string(responseJSON))
	}
	return MethodNotAllowed("")
}
