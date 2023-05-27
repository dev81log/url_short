CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    expiration TIMESTAMP NOT NULL
);

CREATE TABLE codes (
    id SERIAL PRIMARY KEY,
    code CHAR(6) NOT NULL,
    url_id INTEGER REFERENCES urls(id),
    expiration TIMESTAMP NOT NULL
);


