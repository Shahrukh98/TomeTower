-- create_table_user

-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists users (
    id UUID primary key default uuid_generate_v4(),
    username varchar(255),
    nick varchar(255) unique,
    email varchar(255) unique,
    password varchar(255),
    role varchar(50),
    nick_last_updated timestamp default CURRENT_TIMESTAMP,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);
