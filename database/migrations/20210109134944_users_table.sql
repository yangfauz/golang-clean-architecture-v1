-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    id          serial NOT NULL PRIMARY KEY,
	name        varchar(100) NOT NULL,
	role        varchar(255) NOT NULL,
	created_at  TIMESTAMP NULL DEFAULT NULL,
	updated_at  TIMESTAMP NULL DEFAULT NULL
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;