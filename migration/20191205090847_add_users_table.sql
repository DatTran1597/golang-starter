-- +goose Up
-- +goose StatementBegin
CREATE table users(
    id serial not null primary key,
    name text not null default '',
    created_at timestamp without time zone not null default (now() at time zone 'utc'),
    updated_at timestamp without time zone not null default (now() at time zone 'utc'),
    deleted_at timestamp without time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table if exists users cascade;
-- +goose StatementEnd
