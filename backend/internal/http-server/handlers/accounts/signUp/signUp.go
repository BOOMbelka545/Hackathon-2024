package signup

import (
	"context"
	"crypto/sha1"
	postgres "donPass/backend/internal/storage/postgres"
	postgr "donPass/backend/internal/storage/postgres/sqlc"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

var v = validator.New()

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SignUp(c echo.Context) error {
	var account postgr.Account

	if err := c.Bind(&account); err != nil {
		log.Infof("Cannot bind the request: %v \n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something wrong witg request. Please check it :)")
	}

	// Check if account already exists
	account1, _ := postgres.DB.GetAccountByNumber(context.Background(), account.Number)
	if account1.Number == account.Number {
		return echo.NewHTTPError(http.StatusConflict, "Account already exists")
	}

	// Validate request
	c.Echo().Validator = &CustomValidator{validator: v}
	if err := c.Validate(account); err != nil {
		return err
	}

	arg := postgr.CreateAccountParams{
		Number:    account.Number,
		Password:  GeneratePasswordHash(account.Password),
		FirstName: account.FirstName,
		Name:      account.Name,
		LastName:  account.LastName,
		Balance:   5000,
	}

	newUser, err := postgres.DB.CreateAccount(context.Background(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUser.ID)
}

const salt = "P4ever-erfinemo"
func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}