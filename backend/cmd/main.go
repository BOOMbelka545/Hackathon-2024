package main

import (
	"donPass/backend/internal/config"
	deleteAccount "donPass/backend/internal/http-server/handlers/accounts/delete"
	profile "donPass/backend/internal/http-server/handlers/accounts/getProfile"
	"donPass/backend/internal/http-server/handlers/accounts/payment"
	signin "donPass/backend/internal/http-server/handlers/accounts/signIn"
	signup "donPass/backend/internal/http-server/handlers/accounts/signUp"
	updatepassword "donPass/backend/internal/http-server/handlers/accounts/updatePassword"
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
	postgres.NewDB()
	postgres.NewTX()

	// *Handlers
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/user/signUp", signup.SignUp)
	e.POST("/user/signIn", signin.SignIn)
	e.GET("/user/getProfile", profile.GetProfile)
	e.PUT("/payment", payment.Payment)
	e.DELETE("/user/delete", deleteAccount.DeleteAccount)
	e.PUT("/user/update-password", updatepassword.UpdateAccountPassword)
	e.POST("/refresh", signin.RefreshToken)

	// *Run server
	e.Logger.Infof("Listening on %s", cfg.HTTPAddress)
	e.Logger.Fatal(e.Start(cfg.HTTPAddress))
}
