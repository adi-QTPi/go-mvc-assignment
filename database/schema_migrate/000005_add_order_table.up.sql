
CREATE TABLE IF NOT EXISTS `order` (
    order_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_at TIMESTAMP NOT NULL,
    table_no BIGINT,
    customer_id VARCHAR(36) NOT NULL,
    status ENUM('received', 'cooking', 'ready_to_serve', 'paid'),
    total_price BIGINT NOT NULL DEFAULT 0,
    FOREIGN KEY (customer_id) REFERENCES user(user_id),
    FOREIGN KEY (table_no) REFERENCES `table`(table_id)
);
-- CREATE INDEX idx_order_customer_id ON `order`(customer_id);
-- CREATE INDEX idx_order_status ON `order`(status);