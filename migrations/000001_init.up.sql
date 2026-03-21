CREATE SCHEMA agilewebapp;

CREATE TABLE agilewebapp.users (
	id	  SERIAL                PRIMARY KEY,
	version   INTEGER      NOT NULL DEFAULT 1,
	full_name VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
	email     VARCHAR(50)  NOT NULL 
);

CREATE TABLE agilewebapp.tasks (
	id	       SERIAL                PRIMARY KEY,
	version        INTEGER      NOT NULL DEFAULT 1,
	title 	       VARCHAR(100) NOT NULL,
	description    VARCHAR(1000),
	completed      BOOLEAN      NOT NULL,
	created_at     TIMESTAMPTZ  NOT NULL,
	completed_at   TIMESTAMPTZ,

	CHECK (
		(completed=FALSE AND completed_at IS NULL)
		OR
		(completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
	),

	author_user_id INTEGER      NOT NULL REFERENCES agilewebapp.users(id),
);
