-- name: CreateBooksTable :exec

-- name: GetAllBooks :many
SELECT * FROM books;

-- name: GetBookByID :one   
SELECT * FROM books WHERE id = $1;

-- name: AddBook :one
INSERT INTO books (title, author, published_date) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateBook :exec
UPDATE books SET title = $1, author = $2, published_date = $3 WHERE id = $4;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;
