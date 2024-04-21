package signin

import (
	"context"
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	"donPass/backend/internal/storage/postgres"
	postgr "donPass/backend/internal/storage/postgres/sqlc"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	tokenAccess  = 12 * time.Hour
	tokenRefresh = 24 * time.Hour * 5
)

type signInInput struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type Tokens struct {
	AcceessToken string
	RefreshToken string
}

func SignIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		log.Infof("Cannot bind the request: %v \n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something wrong witg request. Please check it :)")
	}

	AccessToken, err := jwt.GenerateToken(input.Number, input.Password, tokenAccess)
	if err != nil {
		log.Errorf("Some problem: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	RefreshToken, err := jwt.GenerateToken(input.Number, input.Password, tokenRefresh)
	if err != nil {
		log.Errorf("Some problem: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, Tokens{AcceessToken: AccessToken,
		RefreshToken: RefreshToken})
}

func RefreshToken(c echo.Context) error {
	var account postgr.Account

	claims, err := jwt.GetClaims(c.Request())
	if err != nil {
		log.Errorf("Cannot get claims from request: %v", err)
		return err
	}

	userId := claims["id"].(float64)
	account, err = postgres.DB.GetAccountByID(context.Background(), int64(userId))
	if err != nil {
		log.Errorf("Cannot get account by id: %v", err)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse JWT with claims")
	}
	fmt.Println(account)
	accessToken, err := jwt.GenerateToken(account.Number, account.Password, tokenAccess)
	if err != nil {
		log.Errorf("Unable to generate the token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to generate the token")
	}

	refreshToken, err := jwt.GenerateToken(account.Number, account.Password, tokenRefresh)
	if err != nil {
		log.Errorf("Unable to generate the token")
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to generate the token")
	}



	return c.JSON(http.StatusOK, Tokens{AcceessToken: accessToken,
		RefreshToken: refreshToken})
}
