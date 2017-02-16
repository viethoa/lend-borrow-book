
-- +goose Up
-- +goose StatementBegin
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS books (
   id BIGSERIAL PRIMARY KEY NOT NULL,
   name VARCHAR(100) NOT NULL,
   short_description VARCHAR(250),
   category VARCHAR(100),
   status VARCHAR(100),
   suggestion VARCHAR(250),
   created_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS requests;

