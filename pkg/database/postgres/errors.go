package postgres

import (
	"errors"
	"marketplace/pkg/database"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	PgForeignKeyViolation = "23503"
	PgUniqueViolation = "23505"
)

var (
	codeErrors = map[string]error {
		PgForeignKeyViolation: database.ErrForeignKeyViolation,
		PgUniqueViolation: database.ErrUniqueViolation,
	}
)

func DetectError(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return database.ErrUndocumented
	}

	if dbErr, ok := codeErrors[pgErr.Code]; ok {
		return dbErr
	}

	return database.ErrUndocumented
}
