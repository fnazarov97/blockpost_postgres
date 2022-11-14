BEGIN;

CREATE TABLE author (
	id CHAR(36) PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE article (
	id CHAR(36) PRIMARY KEY,
	title VARCHAR(255) UNIQUE NOT NULL,
	body TEXT NOT NULL,
	author_id CHAR(36) REFERENCES author (id),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

COMMIT;