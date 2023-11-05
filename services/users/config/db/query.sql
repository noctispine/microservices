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
    email, hashed_password, name, surname
) VALUES (
    $1, $2, $3, $4
);

-- name: DeleteByEmail :execrows
DELETE FROM users WHERE email = $1;

-- name: DeleteById :execrows
DELETE FROM users WHERE id = $1;

-- name: ActivateUser :execrows
UPDATE users SET is_active = TRUE WHERE ID = $1;