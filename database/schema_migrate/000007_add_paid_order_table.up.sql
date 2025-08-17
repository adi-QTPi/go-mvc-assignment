
CREATE TABLE IF NOT EXISTS paid_orders (
    order_id BIGINT PRIMARY KEY,
    customer_review VARCHAR(1000),
    total_amount BIGINT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES `order`(order_id)
);

INSERT INTO `table` (table_id, is_empty) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(9, 1),
(10, 1);