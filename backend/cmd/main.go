package main

import (
	"donPass/backend/internal/config"
	profile "donPass/backend/internal/http-server/handlers/accounts/getProfile"
	signin "donPass/backend/internal/http-server/handlers/accounts/signIn"
	signup "donPass/backend/internal/http-server/handlers/accounts/signUp"
	"donPass/backend/internal/storage/postgres"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// *Init config
	cfg := config.MustLoad()
	log.Infof("starting: %v", cfg)

	// *Init db
	postgres.New()

	// *Handlers
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/user/signUp", signup.SignUp)
	e.POST("/user/signIn", signin.SignIn)
	e.GET("/user/getProfile", profile.GetProfile)

	// *Run server
	e.Logger.Infof("Listening on %s", cfg.HTTPAddress)
	e.Logger.Fatal(e.Start(cfg.HTTPAddress))
}
