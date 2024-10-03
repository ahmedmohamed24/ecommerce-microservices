CREATE TABLE products_orders (
    product_id INT NOT NULL,
    order_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, order_id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);