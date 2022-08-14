CREATE TABLE IF NOT EXISTS blocks (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    district_id integer REFERENCES districts(id)
)