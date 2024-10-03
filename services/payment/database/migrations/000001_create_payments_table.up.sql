CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL check (order_id > 0),
    payment_method VARCHAR(255) NOT NULL,
    paid_amount DECIMAL(10, 2) NOT NULL,
    card_last_four VARCHAR(4) NOT NULL,
    payment_info JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);