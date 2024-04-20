package main

import (
	"context"
	"donPass/backend/internal/config"
	p "donPass/backend/internal/storage/postgres/sqlc"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var dbConn *pgx.Conn

func main() {
	// *Init config
	cfg := config.MustLoad()
	log.Infof("starting: %v", cfg)

	fmt.Println(cfg.DBAddress)

	// *Init db
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.DBAddress, cfg.NameDB)
	dbConn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	_ = p.New(dbConn)

	// *Handlers
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/user/signUp", )

	// *Run server
	e.Logger.Infof("Listening on %s", cfg.HTTPAddress)
	e.Logger.Fatal(e.Start(cfg.HTTPAddress))
}
