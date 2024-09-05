CREATE TABLE IF NOT EXISTS tasks (
    id serial primary key,
    title varchar(255) not null,
    description varchar(255) not null,
    due_date timestamp without time zone null default null,
    done boolean default false,
    created_at timestamp without time zone default current_timestamp not null,
    updated_at timestamp without time zone default current_timestamp not null,
    deleted_at timestamp without time zone null default null
);
