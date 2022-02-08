package controllers

import (
	"aav_logger/app/models"
	"aav_logger/app/transactions"
	"aav_logger/app/utils"
	"encoding/json"
	"fmt"
	"net/url"

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
		flightID, err := transactions.InsertFlight(DBCONN, newFlight)
		if err != nil {
			// TODO: Log the error
			fmt.Println(err)
			resp := models.APIResponse{Status: "server_error", Message: "A server error occurred"}
			responseJSON, _ := json.Marshal(resp)
			return ServerError(responseJSON)
		}

		// Query the flight we just inserted and echo it back in the response.
		/* --- Stringify the id so we can use the Flights transaction --- */
		queryParams := map[string]string{"id": fmt.Sprint(flightID)}

		flights, err := transactions.Flights(DBCONN, queryParams)
		if err != nil {
			// TODO: Log the error
			fmt.Println(err)
			resp := models.APIResponse{Status: "server_error", Message: "A server error occurred"}
			responseJSON, _ := json.Marshal(resp)
			return ServerError(responseJSON)
		}

		// Marshal the response, send it back!
		resp := models.APIResponse{Status: "success", Results: flights}
		responseJSON, _ := json.Marshal(resp)
		return Created(string(responseJSON))
	}
	return MethodNotAllowed("")
}

/* ---
 * Add flight records in bulk
 * --- */
func (c APIV1) NewFlightBulk() revel.Result {
	if c.Request.Method == "POST" {
		// Get the file from the request
		multipartFile := c.Params.Files["logfile"]

		// TODO: Add a sanity check on the file

		// Save a copy of the log to disk
		logfile, err := utils.SaveLogFile(multipartFile)
		if err != nil {
			//TODO: Log the error
			fmt.Println(err)
			resp := models.APIResponse{Status: "server_error", Message: "A server error occurred"}
			responseJSON, _ := json.Marshal(resp)
			return ServerError(responseJSON)
		}

		// Insert records from the saved file
		insertCount, err := transactions.InsertFlightBulk(DBCONN, logfile)
		if err != nil {
			//TODO: Log the error
			fmt.Println(err)
			resp := models.APIResponse{Status: "server_error", Message: "A server error occurred"}
			responseJSON, _ := json.Marshal(resp)
			return ServerError(responseJSON)
		}

		responseJSON, _ := json.Marshal(map[string]int64{"records_inserted": insertCount})
		return Created(string(responseJSON))
	}
	return MethodNotAllowed("")
}

/* ---
 * Read flights from store
 * --- */
func (c APIV1) Flights() revel.Result {
	if c.Request.Method == "GET" {
		// Parse the query parameters
		params := parseParams(c.Params.Query)

		// Query flights based on supplied params
		flights, err := transactions.Flights(DBCONN, params)
		if err != nil && err.Error() == "sql: no rows in result set" {
			resp := models.APIResponse{Status: "api_error", Message: "No records matched your query"}
			responseJSON, _ := json.Marshal(resp)
			return OK(responseJSON)

		} else if err != nil {
			resp := models.APIResponse{Status: "server_error", Message: "A server error occurred"}
			responseJSON, _ := json.Marshal(resp)
			return ServerError(responseJSON)
		}

		// Marshal and return the results
		// Marshal the response, send it back!
		resp := models.APIResponse{Status: "success", Results: flights}
		responseJSON, _ := json.Marshal(resp)
		return OK(string(responseJSON))
	}
	return MethodNotAllowed("")
}

func parseParams(v url.Values) map[string]string {
	var params = make(map[string]string)

	// Parse expected params based on obj/schema attributes
	params["id"] = v.Get("id")
	params["robot"] = v.Get("robot")
	params["generation"] = v.Get("generation")
	params["start"] = v.Get("start")
	params["stop"] = v.Get("stop")
	params["duration"] = v.Get("duration")
	return params
}
