package updatepassword

import (
	"context"
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	signup "donPass/backend/internal/http-server/handlers/accounts/signUp"
	postgr "donPass/backend/internal/storage/postgres"
	postgres "donPass/backend/internal/storage/postgres/sqlc"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func UpdateAccountPassword(c echo.Context) error {
	var updateArg postgres.UpdateAccountPasswordParams

	if err := c.Bind(&updateArg); err != nil {
		log.Infof("Cannot bind the request: %v \n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something wrong witg request. Please check it :)")
	}

	claims, err := jwt.GetClaims(c.Request())
	if err != nil {
		log.Errorf("Cannot get claims from request: %v", err)
		return err
	}
	updateArg.Password = signup.GeneratePasswordHash(updateArg.Password)
	updateArg.ID = int64(claims["id"].(float64))

	_, err = postgr.DB.UpdateAccountPassword(context.Background(), updateArg)
	if err != nil{
		return err
	}

	return c.JSON(http.StatusOK, "Password was updated")
}
