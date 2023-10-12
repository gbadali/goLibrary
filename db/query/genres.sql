-- name: CreateGenre :one
INSERT INTO genres (
  "genre_name"
) VALUES (
  $1
) RETURNING *;

-- name: GetGenre :one
SELECT * FROM genres 
WHERE "genre_id" = $1;

-- name: ListGenres :many
SELECT genre_name FROM genres
ORDER BY "genre_id";

-- name: SearchGenres :one
SELECT genre_id FROM genres
WHERE genre_name = $1
LIMIT 1;

-- name: UpdateGenre :one
UPDATE genres
SET "genre_name" = $2
WHERE   "genre_id" = $1
RETURNING *;


-- name: DeleteGenre :exec
DELETE FROM genres
WHERE genre_id = $1;