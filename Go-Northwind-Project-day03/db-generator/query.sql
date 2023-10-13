-- 1. business_entity
-- name: GetCategory :one
SELECT * FROM business_entity
WHERE entity_id = $1;

-- name: ListCategories :many
SELECT * FROM business_entity
ORDER BY entity_id;
