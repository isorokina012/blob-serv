-- +migrate Up

create table blobs (
    id text primary key,
    user_id text not null,
    blob jsonb default '{}'::jsonb not null
);

-- +migrate Down

drop table blobs;