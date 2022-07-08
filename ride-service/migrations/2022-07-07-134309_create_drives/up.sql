CREATE TABLE drives (
  id SERIAL PRIMARY KEY,
  departure_location VARCHAR NOT NULL,
  destination VARCHAR NOT NULL,
  departure_date_time BIGINT NOT NULL,
  departure_address VARCHAR NOT NULL,
  free_places INTEGER NOT NULL,
  planned_arrival_time BIGINT NOT NULL,
  note VARCHAR NOT NULL, 
  finished BOOLEAN DEFAULT 'f' NOT NULL,
  distance INTEGER NOT NULL
)