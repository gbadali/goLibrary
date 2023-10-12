-- name: CreateAuthor :one
INSERT INTO authors (
  "first_name",
  "last_name"
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors 
WHERE "author_id" = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY "author_id"
LIMIT $1
OFFSET $2;

-- name: SearchAuthors :one
SELECT * FROM authors
WHERE "first_name" = $1 AND "last_name" = $2
LIMIT 1;


-- name: UpdateAuthor :one
UPDATE authors
SET "first_name" = $2, "last_name" = $3
WHERE   "author_id" = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE author_id = $1;