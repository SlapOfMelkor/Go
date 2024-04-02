-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    rol TEXT NOT NULL,
    username TEXT NOT NULL,
    pasword TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE users;
-- +goose StatementEnd
