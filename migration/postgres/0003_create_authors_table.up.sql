-- create_table_authors

create table if not exists authors (
    id UUID primary key default uuid_generate_v4(),
    name varchar(255) not null,
    photo_url varchar(2048) not null,
    nationality varchar(255),
    date_of_birth date,
    date_of_death date default null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);
