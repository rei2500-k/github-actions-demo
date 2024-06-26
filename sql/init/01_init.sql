create table todo_list(
    id integer,
    title varchar(64) not null,
    completed boolean not null default false,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz,
    primary key(id)
);

insert into todo_list values
(1, 'test', false)
;
