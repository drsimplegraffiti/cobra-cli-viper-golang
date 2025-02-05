// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package store

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING id, name
`

// Categories
func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.createCategoryStmt, createCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createNote = `-- name: CreateNote :one
INSERT INTO notes (title, content, category_id) VALUES ($1, $2, $3) RETURNING id, category_id, title, content, created_at, updated_at
`

type CreateNoteParams struct {
	Title      string         `json:"title"`
	Content    sql.NullString `json:"content"`
	CategoryID sql.NullInt32  `json:"category_id"`
}

// Notes
func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.queryRow(ctx, q.createNoteStmt, createNote, arg.Title, arg.Content, arg.CategoryID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllCategories = `-- name: DeleteAllCategories :exec
DELETE FROM categories
`

func (q *Queries) DeleteAllCategories(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteAllCategoriesStmt, deleteAllCategories)
	return err
}

const deleteAllNotes = `-- name: DeleteAllNotes :exec
DELETE FROM notes
`

func (q *Queries) DeleteAllNotes(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteAllNotesStmt, deleteAllNotes)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCategoryStmt, deleteCategory, id)
	return err
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1
`

func (q *Queries) DeleteNote(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteNoteStmt, deleteNote, id)
	return err
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name FROM categories
WHERE id = $1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int32) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByIdStmt, getCategoryById, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getNote = `-- name: GetNote :one
SELECT n.id, n.title, c.name AS category, n.content, n.created_at, n.updated_at
FROM notes n
LEFT JOIN categories c ON n.category_id = c.id
WHERE n.id = $1
`

type GetNoteRow struct {
	ID        int32          `json:"id"`
	Title     string         `json:"title"`
	Category  sql.NullString `json:"category"`
	Content   sql.NullString `json:"content"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

func (q *Queries) GetNote(ctx context.Context, id int32) (GetNoteRow, error) {
	row := q.queryRow(ctx, q.getNoteStmt, getNote, id)
	var i GetNoteRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Category,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name FROM categories
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.listCategoriesStmt, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const listNotes = `-- name: ListNotes :many
SELECT n.id, n.title, c.name AS category, n.created_at
FROM notes n
LEFT JOIN categories c ON n.category_id = c.id
ORDER BY n.created_at DESC
`

type ListNotesRow struct {
	ID        int32          `json:"id"`
	Title     string         `json:"title"`
	Category  sql.NullString `json:"category"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

func (q *Queries) ListNotes(ctx context.Context) ([]ListNotesRow, error) {
	rows, err := q.query(ctx, q.listNotesStmt, listNotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListNotesRow{}
	for rows.Next() {
		var i ListNotesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Category,
			&i.CreatedAt,
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

const paginateNotes = `-- name: PaginateNotes :many
SELECT 
    n.id, 
    n.title, 
    c.name AS category, 
    n.created_at
FROM notes n
LEFT JOIN categories c ON n.category_id = c.id
WHERE 
    ($1 = '' OR c.name = $1)
    AND ($2 = '' OR n.title ILIKE '%' || $2 || '%')
ORDER BY n.created_at DESC
LIMIT $3 OFFSET $4
`

type PaginateNotesParams struct {
	Column1 interface{} `json:"column_1"`
	Column2 interface{} `json:"column_2"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

type PaginateNotesRow struct {
	ID        int32          `json:"id"`
	Title     string         `json:"title"`
	Category  sql.NullString `json:"category"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

// @param category: string
// @param title: string
// @param limit: int
// @param offset: int
// @return id: int, title: string, category: string, created_at: timestamp
func (q *Queries) PaginateNotes(ctx context.Context, arg PaginateNotesParams) ([]PaginateNotesRow, error) {
	rows, err := q.query(ctx, q.paginateNotesStmt, paginateNotes,
		arg.Column1,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PaginateNotesRow{}
	for rows.Next() {
		var i PaginateNotesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Category,
			&i.CreatedAt,
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

const updateNoteContent = `-- name: UpdateNoteContent :exec
UPDATE notes SET content = $1, updated_at = now() WHERE id = $2
`

type UpdateNoteContentParams struct {
	Content sql.NullString `json:"content"`
	ID      int32          `json:"id"`
}

func (q *Queries) UpdateNoteContent(ctx context.Context, arg UpdateNoteContentParams) error {
	_, err := q.exec(ctx, q.updateNoteContentStmt, updateNoteContent, arg.Content, arg.ID)
	return err
}
