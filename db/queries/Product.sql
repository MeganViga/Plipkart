-- name: AddProduct :one
INSERT INTO product(
    product_name,
    price,
    created_at,
    modified_at,
    deleted_at
 ) VALUES($1,$2,$3,$4,$5) RETURNING *;


 -- name: AddProductCategory :one
 INSERT INTO product_category(
    product_id,
    color,
    quantity
 ) VALUES($1,$2,$3) RETURNING *;

 -- name: GetProducts :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id;

 -- name: GetProductById :one
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE id = $1;

 -- name: GetProductByName :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE product_name = $1;

 -- name: GetProductByColor :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE product_id = $1 AND color = $2;
