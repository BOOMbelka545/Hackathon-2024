package postgres

import (
	"context"
	"donPass/backend/internal/config"
	p "donPass/backend/internal/storage/postgres/sqlc"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
)

var DB *p.Queries

func New() {
	// *Init config
	cfg := config.MustLoad()

	// *Init db
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.DBAddress, cfg.NameDB)
	dbConn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	DB = p.New(dbConn)
}
