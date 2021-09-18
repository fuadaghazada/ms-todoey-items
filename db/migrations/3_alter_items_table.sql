-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE items
    ADD COLUMN deleted_at timestamp with time zone default null;
