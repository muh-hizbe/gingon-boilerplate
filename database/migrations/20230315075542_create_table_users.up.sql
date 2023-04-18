CREATE TABLE users (
    -- id SERIAL PRIMARY KEY,
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address VARCHAR(255) NULL,
    phone VARCHAR(255) NULL,
    password VARCHAR(255) NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);