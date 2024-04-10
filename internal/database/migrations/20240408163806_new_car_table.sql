-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS people (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS car (
 id SERIAL PRIMARY KEY,
 regNum VARCHAR(255) NOT NULL,
 mark VARCHAR(255) NOT NULL,
 model VARCHAR(255) NOT NULL
);

ALTER TABLE people ADD COLUMN patronymic VARCHAR(20) DEFAULT 'none';
ALTER TABLE car ADD COLUMN year VARCHAR(20) DEFAULT 'none';
ALTER TABLE car ADD COLUMN owner INTEGER REFERENCES people(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE people;
DROP TABLE car;
-- +goose StatementEnd