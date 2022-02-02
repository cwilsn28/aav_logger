CREATE TABLE flights (
	id bigserial,
	robot text,
	generation bigint,
	start timestamp without time zone,
	stop timestamp without time zone,
	lat double precision,
	lon double precision
);
GRANT ALL ON TABLE flights TO aav_logger;
