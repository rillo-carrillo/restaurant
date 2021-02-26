DROP TABLE modifiers;
DROP TABLE order_products;
DROP TABLE order_details;
DROP TABLE orders;
DROP TABLE employees;
DROP TABLE roles;
DROP TABLE products;
DROP TABLE categories;
DROP TABLE clients;
DROP TABLE restaurants;
DROP TABLE users;





TRUNCATE TABLE restaurants;
TRUNCATE TABLE clients;
TRUNCATE TABLE categories;
TRUNCATE TABLE products;
TRUNCATE TABLE roles;
TRUNCATE TABLE employees;
TRUNCATE TABLE orders;
TRUNCATE TABLE order_products;
TRUNCATE TABLE modifiers;


INSERT INTO "modifiers" ("created_at","updated_at","deleted_at","product_id","description","price","id") VALUES 
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,1,'Lata 355ml',0.000000,1),
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,1,'Botella Vidrio 500ml',4.000000,2),
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,3,'Verdes',0.000000,4),
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,3,'Verdes',0.000000,4),
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,3,'Pollo',10.000000,6),
('2021-02-07 22:19:13.132','2021-02-07 22:19:13.132',NULL,3,'Queso',10.000000,7) 
ON CONFLICT ("id") DO UPDATE SET "product_id"="excluded"."product_id" RETURNING "id"


SELECT p.name, m.description, op.comment,op.number from order_products op
INNER JOIN modifiers m ON m.id = op.modifier_id
INNER JOIN products p ON p.id = m.product_id;
