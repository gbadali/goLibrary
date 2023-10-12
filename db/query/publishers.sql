-- name: CreatePublisher :one
INSERT INTO publishers (
  "publisher_name",
  "address",
  "phone"
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPublisher :one
SELECT * FROM publishers 
WHERE "publisher_id" = $1;

-- name: SearchPublisher :one
SELECT publisher_id FROM publishers
WHERE publisher_name = $1;

-- name: ListPublishers :many
SELECT publisher_name FROM publishers
ORDER BY "publisher_id";


-- name: UpdatePublisher :one
UPDATE publishers
SET "publisher_name" = $2, "address" = $3, "phone" = $4
WHERE   "publisher_id" = $1
RETURNING *;


-- name: DeletePublishers :exec
DELETE FROM publishers
WHERE publisher_id = $1;