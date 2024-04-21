package postgres

import (
	"context"
	"donPass/backend/internal/config"
	p "donPass/backend/internal/storage/postgres/sqlc"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

var DB *p.Queries
var TX *pgxpool.Pool
var QueriesTX *p.Queries

func NewDB() {
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

func NewTX() {
		// *Init config
		cfg := config.MustLoad()
		// *Init db
		dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.DBAddress, cfg.NameDB)
		TX, err := pgxpool.New(context.Background(), dbURL)
		if err != nil {
			log.Fatal("Cannot connect to db:", err)
		}
		QueriesTX = p.New(TX)
}
