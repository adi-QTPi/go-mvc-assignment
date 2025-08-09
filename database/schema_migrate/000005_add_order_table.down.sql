USE karma_mvc_foodopiaDB;
DROP INDEX idx_order_customer_id ON `order`;
DROP INDEX idx_order_status ON `order`;

DROP TABLE IF EXISTS `order`;
