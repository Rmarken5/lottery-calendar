-- +goose Up
-- +goose StatementBegin
create database calendar_db;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    drop database calendar_db;
-- +goose StatementEnd
