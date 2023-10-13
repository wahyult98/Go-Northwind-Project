package dbContext

import (
	"context"
	"net/http"

	"codeid.northwind/models"
	"codeid.northwind/models/features"
)

const createCategory = `-- name: CreateCategory :one

INSERT INTO categories(category_id, category_name, description, picture) VALUES ($1, $2, $3, $4)
RETURNING *
`

type CreateCategoryParams struct {
	CategoryID   int16  `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	Description  string `db:"description" json:"description"`
	Picture      []byte `db:"picture" json:"picture"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (*models.Category, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.CategoryID,
		arg.CategoryName,
		arg.Description,
		arg.Picture,
	)
	i := models.Category{}
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.Category{
		CategoryID:   i.CategoryID,
		CategoryName: i.CategoryName,
		Description:  i.Description,
		Picture:      i.Picture,
	}, nil
}

/* func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (models.Category, models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.CategoryID,
		arg.CategoryName,
		arg.Description,
		arg.Picture,
	)
	var category_id int16
	err := row.Scan(&category_id)
	return category_id, err
} */

const listCategories = `-- name: ListCategories :many
SELECT category_id, category_name, description, picture FROM Categories
ORDER BY category_id
limit $1 offset $2
`

func (q *Queries) ListCategories(ctx context.Context, metadata *features.Metadata) ([]models.Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, metadata.PageSize, metadata.PageNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Category
	for rows.Next() {
		var i models.Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.Description,
			&i.Picture,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategory = `-- name: GetCategory :one
SELECT category_id, category_name, description, picture FROM categories
WHERE category_id = $1
`

func (q *Queries) GetCategory(ctx context.Context, categoryID int16) (models.Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, categoryID)
	var i models.Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
  set category_name = $2,
  description = $3
WHERE category_id = $1
`

func (q *Queries) UpdateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.CategoryID, arg.CategoryName, arg.Description)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM Categories
WHERE category_id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, categoryID int16) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, categoryID)
	return err
}
