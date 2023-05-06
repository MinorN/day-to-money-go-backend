-- name: ListUsers :many
SELECT * from users
ORDER BY id
OFFSET $1
LIMIT $2;


-- name: FindUser :one
SELECT * from users
WHERE id = $1;

-- name: FindByEmail :one
SELECT * from users
WHERE email = $1;

-- name: FindByPhone :one
SELECT * from users
WHERE phone = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


-- name: CreateUser :one
INSERT INTO users (email) values ($1) RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET 
    email = $2,
    phone = $3
WHERE
    id = $1;