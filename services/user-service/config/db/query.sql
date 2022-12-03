-- name: GetAllByNames :many
SELECT * FROM users
ORDER BY name;

-- name: GetAll :many
SELECT * FROM users;

-- name: GetByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetById :one
SELECT * FROM users
WHERE id = $1;

-- name: Create :exec
INSERT INTO users (
    name, surname, email
) VALUES (
    $1, $2, $3
);