CREATE TABLE ip_addresses(
    id bigint,
    ip_public varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)

CREATE TABLE users(
    id bigint,
    username varchar(255),
    password varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)