CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address VARCHAR(255) NULL,
    phone VARCHAR(255) NULL,
    password VARCHAR(255) NULL,
    created_at TIMESTAMPTZ NULL,
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);