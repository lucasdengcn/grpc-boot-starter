BEGIN;

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	user_name VARCHAR(100) NOT NULL DEFAULT '',
	email VARCHAR(100) NOT NULL DEFAULT '',
	hashed_password TEXT NOT NULL DEFAULT '',
	birthday date NOT NULL DEFAULT 0001-01-01,
    gender INT NOT NULL DEFAULT 0,
    photo_url VARCHAR(100) NOT NULL DEFAULT '',
    active boolean NOT NULL DEFAULT TRUE,
    status INT NOT NULL DEFAULT 0,
    roles TEXT NOT NULL DEFAULT '',
    deleted boolean NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

COMMIT;