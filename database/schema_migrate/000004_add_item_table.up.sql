USE karma_mvc_foodopiaDB;
CREATE TABLE IF NOT EXISTS item (
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
-- CREATE INDEX idx_item_cat_id ON item(cat_id);
-- CREATE INDEX idx_item_subcat_id ON item(subcat_id);