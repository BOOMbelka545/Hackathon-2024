package deleteAccount

import (
	"context"
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	"donPass/backend/internal/storage/postgres"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func DeleteAccount(c echo.Context) error {
	claims, err := jwt.GetClaims(c.Request())
	if err != nil {
		log.Errorf("Cannot get claims from request: %v", err)
		return err
	}
	userId := claims["id"].(float64)
	fmt.Println(userId)

	err = postgres.DB.DeletePayments(context.Background(), int64(userId))
	if err != nil {
		return err
	}

	err = postgres.DB.DeleteAccount(context.Background(), int64(userId))
	if err != nil {
		log.Errorf("Cannot delete account:%v", err)
		return err
	}

	return c.JSON(http.StatusOK, "Account was deleted")
}
