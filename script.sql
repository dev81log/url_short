DROP TABLE IF EXISTS url_codes;

DROP TABLE IF EXISTS urls;

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    url VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

CREATE TABLE url_codes (
    id SERIAL PRIMARY KEY,
    url_id INT NOT NULL REFERENCES urls(id),
    code CHAR(6) NOT NULL,
    created_at TIMESTAMP NOT NULL
);


