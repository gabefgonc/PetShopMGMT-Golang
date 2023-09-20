-- +goose Up
-- +goose StatementBegin
CREATE TABLE costumers (
  id serial PRIMARY KEY,
  name VARCHAR (250) NOT NULL,
  phone_number VARCHAR (50) UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE costumers;
-- +goose StatementEnd
