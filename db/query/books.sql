-- name: CreateBook :one
INSERT INTO books (
  "title",
  "isbn",
  "publication_year",
  "author_id",
  "publisher_id",
  "genre_id",
  "price",
  "stock_quantity"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetBook :one
SELECT * FROM books 
WHERE "book_id" = $1;


-- name: UpdateBook :one
UPDATE books
SET 
    "title" = $2,
    "isbn" = $3,
    "publication_year" = $4,
    "author_id" = $5,
    "publisher_id" = $6,
    "genre_id" = $7,
    "price" = $8,
    "stock_quantity" = $9
WHERE   "book_id" = $1
RETURNING *;


-- name: DeleteBook :exec
DELETE FROM books
WHERE book_id = $1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY "title"
LIMIT $1
OFFSET $2;

-- name: ListBooksJoin :many
SELECT
    b.title AS book_title,
    b.isbn,
    b.publication_year,
    a.first_name AS author_first_name,
    a.last_name AS author_last_name,
    p.publisher_name,
    g.genre_name,
    b.price,
    b.stock_quantity
FROM
    books AS b
INNER JOIN
    authors AS a ON b.author_id = a.author_id
INNER JOIN
    publishers AS p ON b.publisher_id = p.publisher_id
INNER JOIN
    genres AS g ON b.genre_id = g.genre_id
LIMIT $1
OFFSET $2;
