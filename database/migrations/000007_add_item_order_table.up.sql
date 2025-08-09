CREATE TABLE item_order (
    order_id BIGINT,
    item_id BIGINT,
    quantity INTEGER NOT NULL,
    instruction VARCHAR(500),
    is_complete ENUM('pending', 'taken', 'complete') NOT NULL DEFAULT 'pending',
    cook_id VARCHAR(36) NULL,
    PRIMARY KEY (order_id, item_id),
    FOREIGN KEY (order_id) REFERENCES `order`(order_id),
    FOREIGN KEY (item_id) REFERENCES item(item_id),
    FOREIGN KEY (cook_id) REFERENCES user(user_id)
);