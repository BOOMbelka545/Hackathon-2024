// versions:
//   sqlc v1.26.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64
	Number    string      `json:"number"`
	Password  string      `json:"password,omitempty" validate:"min=8,max=300,required"`
	FirstName string      `json:"first_name,omitempty", validate:"required"`
	Name      string      `json:"name,omitempty", validate:"required"`
	LastName  pgtype.Text `json:"last_name,omitempty", validate:"required"`
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
