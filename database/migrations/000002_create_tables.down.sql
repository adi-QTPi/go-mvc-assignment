DROP INDEX IF EXISTS idx_item_cat_id ON item;
DROP INDEX IF EXISTS idx_item_subcat_id ON item;
DROP INDEX IF EXISTS idx_order_customer_id ON `order`;
DROP INDEX IF EXISTS idx_order_status ON `order`;

DROP TABLE IF EXISTS paid_orders;
DROP TABLE IF EXISTS item_order;
DROP TABLE IF EXISTS `order`;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS `table`;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS category;