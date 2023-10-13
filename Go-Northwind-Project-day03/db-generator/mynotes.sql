SELECT c.*,array_to_json(array_agg(row_to_json (p)))
FROM categories c
JOIN products p ON c.category_id = p.category_id
GROUP by c.category_id

select array_to_json(array_agg(row_to_json (t)))  as categories from (
SELECT c.*,array_to_json(array_agg(row_to_json (p)))[1] as products 
FROM categories c
JOIN products p ON c.category_id = p.category_id 
GROUP by c.category_id)t limit 1