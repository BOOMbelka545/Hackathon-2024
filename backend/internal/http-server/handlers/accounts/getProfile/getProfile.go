package profile

import (
	"context"
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	postgres "donPass/backend/internal/storage/postgres"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetProfile(c echo.Context) error {
	claims, err := jwt.GetClaims(c.Request())
	if err != nil {
		log.Errorf("Cannot get claims from request: %v", err)
		return err
	}

	userId := claims["id"].(float64)
	account, err := postgres.DB.GetAccountByID(context.Background(), int64(userId))
	if err != nil {
		log.Errorf("Cannot get account by id: %v", err)
		return err
	}

	account.Password = ""
	return c.JSON(http.StatusOK, account)
}
