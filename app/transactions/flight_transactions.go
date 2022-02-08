package transactions

import (
	"aav_logger/app/models"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

func InsertFlight(db *sql.DB, flight models.Flight) (int64, error) {
	var err error
	var flightID int64

	// Just write the SQL in this case.
	var SQL string = `INSERT INTO flights VALUES(DEFAULT,$1,$2,$3,$4,$5,$6) RETURNING id;`
	stmt, err := db.Prepare(SQL)
	if err != nil {
		return flightID, err
	}

	err = stmt.QueryRow(
		flight.Robot,
		flight.Generation,
		flight.Start,
		flight.Stop,
		flight.Lat,
		flight.Lon,
	).Scan(&flightID)

	stmt.Close()
	return flightID, nil
}

func InsertFlightBulk(db *sql.DB, flightlog string) (int64, error) {
	var err error
	var recordCount int64

	// Create a transaction so we can import all records in batch
	txn, err := db.Begin()
	if err != nil {
		return recordCount, err
	}

	stmt, err := txn.Prepare(
		pq.CopyIn(
			"flights",
			"robot",
			"generation",
			"start",
			"stop",
			"lat",
			"lon",
		),
	)
	if err != nil {
		return recordCount, err
	}

	// Open a handle to the flightlog
	csvFile, err := os.Open(flightlog)
	if err != nil {
		// TODO: Log the error
		return recordCount, err
	}

	// Create new csv reader
	r := csv.NewReader(csvFile)

	// Pack all records into an exec buffer
	rowNum := 0
	for {
		record, err := r.Read()
		if rowNum == 0 {
			rowNum++
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return recordCount, err
		}
		flight := parseRecord(record)
		_, err = stmt.Exec(
			flight.Robot,
			flight.Generation,
			flight.Start,
			flight.Stop,
			flight.Lat,
			flight.Lon,
		)
		if err != nil {
			return recordCount, err
		}
		rowNum++
	}

	recordCount = int64(rowNum) + 1

	// Flush the buffer.
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return recordCount, err
	}

	// Close the statement.
	err = stmt.Close()
	if err != nil {
		fmt.Println(err)
		return recordCount, err
	}

	// Close the transaction.
	err = txn.Commit()
	return recordCount, err
}

/* ---
 * All flights
 * --- */
func Flights(db *sql.DB, params map[string]string) ([]models.Flight, error) {
	var err error
	var flights []models.Flight

	// Use squirrel to dynamically build the query
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("*").From("flights")

	// Build a query based on available parameters
	// Need to consider corner cases associate with this approach
	if params["id"] != "" {
		query = query.Where("id=?", params["id"])
	}
	if params["robot"] != "" {
		query = query.Where("robot=?", params["robot"])
	}
	if params["generation"] != "" {
		query = query.Where("generation=?", params["generation"])
	}
	if params["start"] != "" {
		query = query.Where("start >= ?", params["start"])
	}
	if params["stop"] != "" {
		query = query.Where("stop <= ?", params["stop"])
	}
	if params["duration"] != "" {
		query = query.Where("extract(MINUTE FROM (stop - start)) <= ?", params["duration"])
	}

	SQL, args, err := query.ToSql()
	if err != nil {
		return flights, err
	}

	stmt, err := db.Prepare(SQL)
	if err != nil {
		return flights, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return flights, err
	}
	fmt.Printf("%+v\n", rows)
	for rows.Next() {
		flight := models.Flight{}
		err := rows.Scan(
			&flight.ID,
			&flight.Robot,
			&flight.Generation,
			&flight.Start,
			&flight.Stop,
			&flight.Lat,
			&flight.Lon,
		)
		if err != nil {
			return flights, err
		}
		flights = append(flights, flight)
	}
	stmt.Close()
	return flights, err
}

// Local helpers
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
