package users

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID    `db:"id" json:"id" validate:"omitempty"`
	Email          string       `db:"email" json:"email" validate:"required,email"`
	HashedPassword string       `db:"hashed_password" json:"hashedPassword"`
	Name           string       `db:"name" json:"name" validate:"required,max=50"`
	Surname        string       `db:"surname" json:"surname" validate:"required,max=50"`
	Role           int32        `db:"role" json:"role" validate:"required"`
	CreatedAt      time.Time    `db:"created_at" json:"createdAt"`
	LastLoginAt    sql.NullTime `db:"last_login_at" json:"lastLoginAt"`
}
