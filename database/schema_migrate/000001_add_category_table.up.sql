CREATE DATABASE IF NOT EXISTS karma_mvc_foodopiaDB;
USE karma_mvc_foodopiaDB;
CREATE TABLE IF NOT EXISTS category (
    cat_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    cat_name VARCHAR(255) NOT NULL,
    cat_description VARCHAR(500)
);
