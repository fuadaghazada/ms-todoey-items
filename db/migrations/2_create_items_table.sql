-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table items
(
    id          serial constraint pk_item_id primary key,
    title       varchar(255),
    description text,
    user_id     varchar(255),
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone
);

comment on table items is 'Table for storing todoey item information';