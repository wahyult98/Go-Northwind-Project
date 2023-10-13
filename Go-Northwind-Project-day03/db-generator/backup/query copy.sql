-- name: GetCategory :one
SELECT * FROM categories
WHERE category_id = $1;

-- name: ListCategories :many
SELECT * FROM Categories
ORDER BY category_name;

-- name: CreateCategory :one

INSERT INTO categories(category_id, category_name, description, picture) VALUES ($1, $2, $3, $4)
RETURNING category_id;

-- name: DeleteCategory :exec
DELETE FROM Categories
WHERE category_id = $1;

-- name: UpdateCategory :exec
UPDATE categories
  set category_name = $2,
  description = $3
WHERE category_id = $1;

-- products

-- name: GetProducts :one
SELECT * FROM products
WHERE category_id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY product_name;

-- name: CreateProduct :one

INSERT INTO products 
(product_id, product_name, supplier_id, category_id, 
quantity_per_unit, unit_price, units_in_stock, 
units_on_order, reorder_level, discontinued)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;

-- name: UpdateProduct :exec
UPDATE products
  set product_name = $2,
  unit_price = $3
WHERE product_id = $1;