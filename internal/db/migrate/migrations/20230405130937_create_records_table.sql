-- +goose Up
-- +goose StatementBegin

-- Crypto extension
create extension if not exists "uuid-ossp";

-- Records table
create table if not exists records
(
    id         uuid                     not null,
    user_id    varchar(50)              not null,
    datatype   varchar(20)              not null default 'TEXT',
    data       bytea                    not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    primary key (id),
    unique (id)
);

-- +goose StatementEnd