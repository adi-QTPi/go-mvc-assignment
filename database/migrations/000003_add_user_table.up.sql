CREATE TABLE user (
    user_id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_name VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    pwd_hash VARCHAR(255) NOT NULL,
    profile_pic VARCHAR(500),
    role ENUM('customer', 'cook', 'admin') NOT NULL
);