# AAV Logger

An interactive logger for Autonomous Aerial Vehicle flight logs.


## To run the service:
    clone the repo
    cd <path_to>/aav_logger
    docker compose up (or docker-compose up)

## Usage

Once the service is running, the API will be exposed at:

http://localhost:9000/api/v1

The following endpoints are available:

    POST /api/v1/flight     Insert a single log record
    POST /api/v1/flights    Insert multiple logs via csv upload
    GET /api/v1/flights     Query flight logs

Inserting a single record

    POST /api/v1/flight
    
    Request fields:
    
    robot: string
    The unique name of the drone.

    generation: integer
    The generation number of the drone.

    start: string
    Flight start time as a UTC timestamp.

    stop: string
    Flight end time as a UTC timestamp.

    lat: float
    Flight latitude coordinate

    lon: float
    Flight longitude coordinate

    EXAMPLE:

    curl -X POST https://sandbox.plaid.com/transactions/get \
    -H 'Content-Type: application/json' \
    -d '{
            "robot":"drone-10",
            "generation":11,
            "start":"2022-02-01T18:59:19Z",
            "stop":"2022-02-01T19:05:19Z",
            "lat":21.3069, 
            "lon":-157.8583
        }'   
## Code Layout

The directory structure of a generic Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        models/       Data models
        transactions/ Database transaction functions
        utils/        Various helper/utility functions used throughout the service
        views/        Templates directory for any HTML rendering

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    sql/
        pg/
            load/     SQL files for initializing the database on launch

    tests/            Test suites

    test_data/        Contains a .csv with generated log data for testing
    uploads/          A local dir for storing log files that get uploaded





