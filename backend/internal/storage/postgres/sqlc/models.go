// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64
	Number    string
	FirstName string
	Name      string
	LastName  pgtype.Text
	Balance   int64
	CreatedAt pgtype.Timestamptz
}

type Entry struct {
	ID        int64
	AccountID int64
	// can be positive and negative
	Amount    int64
	CreatedAt pgtype.Timestamptz
}
