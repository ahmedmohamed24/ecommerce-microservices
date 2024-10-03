CREATE TABLE order_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL CHECK(name <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    order_status (name)
VALUES
    ('Pending'),
    ('Completed'),
    ('Cancelled');