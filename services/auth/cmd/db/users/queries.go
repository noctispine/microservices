package userDB

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type LoginParams struct {
	Email          string `db:"email" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db DBTX
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

const getByEmail = `-- name: GetByEmail :one
SELECT id, email, hashed_password, name, surname, role, created_at, last_login_at FROM users
WHERE email = $1
`

func (q *Queries) GetByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.Name,
		&i.Surname,
		&i.Role,
		&i.CreatedAt,
		&i.LastLoginAt,
	)
	return i, err
}

const updateLoginAt = `-- name: updateLoginAt :one
UPDATE users
SET last_login_at = $2
WHERE id = $1
`

func (q *Queries) UpdateLastLoginAt(ctx context.Context, id uuid.UUID, lastLoginAt time.Time) error {
	return q.db.QueryRowContext(ctx, updateLoginAt, id, lastLoginAt).Err()
}

const create = `-- name: Create :exec
INSERT INTO users (
    email, hashed_password, name, surname
) VALUES (
    $1, $2, $3, $4
)
`

type CreateParams struct {
	Email          string `db:"email" json:"email"`
	HashedPassword string `db:"hashed_password" json:"hashedPassword"`
	Name           string `db:"name" json:"name"`
	Surname        string `db:"surname" json:"surname"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) error {
	_, err := q.db.ExecContext(ctx, create,
		arg.Email,
		arg.HashedPassword,
		arg.Name,
		arg.Surname,
	)
	return err
}
