CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL CHECK (name <> ''),
    email VARCHAR(255) NOT NULL CHECK (email <> ''),
    password VARCHAR(255) NOT NULL CHECK (password <> ''),
    mobile VARCHAR(255) NOT NULL CHECK (mobile <> ''),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);