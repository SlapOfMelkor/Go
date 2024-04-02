-- name: BookBorrow :many
INSERT INTO borrow (borrow_date,return_date,user_id,book_id,status) 
VALUES ($1, $2, $3,$4,$5) RETURNING *;

-- name: GetBorrows :many
SELECT
    bb.id AS borrow_id,
    bb.borrow_date,
    bb.return_date,
    bb.status,
    u.username AS user_username,
    b.title AS book_name
FROM
    borrow bb
INNER JOIN
    users u ON bb.user_id = u.id
INNER JOIN
    books b ON bb.book_id = b.id
WHERE status= $1 OR bb.return_date < CURRENT_TIMESTAMP;

-- name: GetBorrowHistory :many
SELECT
    bb.id AS borrow_id,
    bb.borrow_date,
    bb.return_date,
    bb.status,
    u.username AS user_username,
    b.title AS book_name
FROM
    borrow bb
INNER JOIN
    users u ON bb.user_id = u.id
INNER JOIN
    books b ON bb.book_id = b.id;

-- name: Returnbook :exec 
UPDATE borrow
  set status = $2
WHERE id = $1;