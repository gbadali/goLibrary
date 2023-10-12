-- name: LinkBookAuthor :one
INSERT INTO book_authors (
   "book_id",
    "author_id"
) VALUES (
    $1,
    $2
) RETURNING *;