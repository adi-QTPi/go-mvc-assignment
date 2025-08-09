USE karma_mvc_foodopiaDB;
DROP INDEX idx_item_cat_id ON item;
DROP INDEX idx_item_subcat_id ON item;

DROP TABLE IF EXISTS item;