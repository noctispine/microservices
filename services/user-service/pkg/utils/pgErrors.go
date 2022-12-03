package utils

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

func CheckPostgreError(err error, code string) bool {
    var pgErr *pq.Error
    if errors.As(err, &pgErr) {
        return pgErr.Code == pgerrcode.UniqueViolation
    }

    return false
}