-- create database demo_db;

-- \c demo_db;

create table todo_list(
    id integer,
    title varchar(64) not null,
    completed boolean not null default false,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz,
    primary key(id)
);
