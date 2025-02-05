-- Categories  
-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- Notes  
-- name: CreateNote :one
INSERT INTO notes (title, content, category_id) VALUES ($1, $2, $3) RETURNING *;

-- name: ListNotes :many
SELECT n.id, n.title, c.name AS category, n.created_at
FROM notes n
LEFT JOIN categories c ON n.category_id = c.id
ORDER BY n.created_at DESC;

-- name: GetCategoryById :one
SELECT * FROM categories
WHERE id = $1;

-- name: UpdateNoteContent :exec
UPDATE notes SET content = $1, updated_at = now() WHERE id = $2;

-- name: GetNote :one
SELECT n.id, n.title, c.name AS category, n.content, n.created_at, n.updated_at
FROM notes n
LEFT JOIN categories c ON n.category_id = c.id
WHERE n.id = $1;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;

-- name: PaginateNotes :many
-- @param category: string
-- @param title: string
-- @param limit: int
-- @param offset: int
-- @return id: int, title: string, category: string, created_at: timestamp
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
LIMIT $3 OFFSET $4;

-- name: DeleteAllNotes :exec
DELETE FROM notes;

-- name: DeleteAllCategories :exec
DELETE FROM categories;
