package users

import (
	"context"
	"database/sql"
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