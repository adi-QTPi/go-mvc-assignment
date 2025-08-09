CREATE TABLE paid_orders (
    order_id BIGINT PRIMARY KEY,
    customer_review VARCHAR(1000),
    total_amount BIGINT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES `order`(order_id)
);