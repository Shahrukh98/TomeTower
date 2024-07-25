-- create_table_genre

create table if not exists genres (
    id UUID primary key default uuid_generate_v4(),
    name varchar(255) unique,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);
