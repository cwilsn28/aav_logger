package transactions

import (
	"aav_logger/app/models"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

func InsertFlight(db *sql.DB, flight models.Flight) (int64, error) {
	var err error
	var recordID int64

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Insert("flights").Values(
		"DEFAULT",
		flight.Robot,
		flight.Generation,
		flight.Start,
		flight.Stop,
		flight.Lat,
		flight.Lon,
	).Suffix("RETURNING id")

	SQL, args, err := query.ToSql()
	if err != nil {
		return recordID, err
	}

	stmt, err := db.Prepare(SQL)
	if err != nil {
		return recordID, err
	}

	err = stmt.QueryRow(args...).Scan(&recordID)
	stmt.Close()
	return recordID, err
}

func FlightWithID(db *sql.DB, flightID int64) (models.Flight, error) {
	var err error
	var flightRecord models.Flight

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("*").From("flights").Where("id=$", flightID)

	SQL, args, err := query.ToSql()
	if err != nil {
		return flightRecord, err
	}

	stmt, err := db.Prepare(SQL)
	if err != nil {
		return flightRecord, err
	}

	err = stmt.QueryRow(args...).Scan(
		&flightRecord.ID,
		&flightRecord.Robot,
		&flightRecord.Generation,
		&flightRecord.Start,
		&flightRecord.Stop,
		&flightRecord.Lat,
		&flightRecord.Lon,
	)
	stmt.Close()
	return flightRecord, err
}
