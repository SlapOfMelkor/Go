-- name: CreateUser :one
INSERT INTO users (rol,username,pasword) 
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
  set username = $2,
  pasword=$3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: Login :one
SELECT rol, username, pasword FROM users WHERE username = $1;