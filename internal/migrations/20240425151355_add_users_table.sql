-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    login VARCHAR UNIQUE NOT NULL ,
    hashed_password VARCHAR NOT NULL,
    username VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd