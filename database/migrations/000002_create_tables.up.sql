CREATE TABLE category (
    cat_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    cat_name VARCHAR(255) NOT NULL,
    cat_description VARCHAR(500)
);

CREATE TABLE user (
    user_id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_name VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    pwd_hash VARCHAR(255) NOT NULL,
    profile_pic VARCHAR(500),
    role ENUM('customer', 'cook', 'admin') NOT NULL
);

CREATE TABLE `table` (
    table_id BIGINT PRIMARY KEY,
    is_empty BOOLEAN
);

CREATE TABLE item (
    item_id BIGINT AUTO_INCREMENT PRIMARY KEY ,
    item_name VARCHAR(255) NOT NULL,
    cook_time_min BIGINT,
    price BIGINT NOT NULL,
    display_pic VARCHAR(500),
    cat_id BIGINT NOT NULL,
    subcat_id BIGINT,
    is_available BOOLEAN DEFAULT 1,
    FOREIGN KEY (cat_id) REFERENCES category(cat_id),
    FOREIGN KEY (subcat_id) REFERENCES category(cat_id)
);
CREATE INDEX idx_item_cat_id ON item(cat_id);
CREATE INDEX idx_item_subcat_id ON item(subcat_id);

CREATE TABLE `order` (
    order_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_at TIMESTAMP NOT NULL,
    table_no BIGINT,
    customer_id VARCHAR(36) NOT NULL,
    status ENUM('received', 'cooking', 'ready_to_serve', 'paid'),
    total_price BIGINT NOT NULL DEFAULT 0,
    FOREIGN KEY (customer_id) REFERENCES user(user_id),
    FOREIGN KEY (table_no) REFERENCES `table`(table_id)
);
CREATE INDEX idx_order_customer_id ON `order`(customer_id);
CREATE INDEX idx_order_status ON `order`(status);

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

CREATE TABLE paid_orders (
    order_id BIGINT PRIMARY KEY,
    customer_review VARCHAR(1000),
    total_amount BIGINT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES `order`(order_id)
);