package payment

import (
	"context"
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	"donPass/backend/internal/storage/postgres"
	postgr "donPass/backend/internal/storage/postgres/sqlc"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Payment(c echo.Context) error {
	var payment postgr.Payment

	if err := c.Bind(&payment); err != nil {
		log.Infof("Cannot bind the request: %v \n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something wrong witg request. Please check it :)")
	}

	claims, err := jwt.GetClaims(c.Request())
	if err != nil {
		log.Errorf("Cannot get claims from request: %v", err)
		return err
	}
	userId := claims["id"].(float64)

	arg := postgr.CreatePaymentParams{
		AccountID: int64(userId),
		Amount:    payment.Amount,
		Place:     payment.Place,
	}
	_, err = postgres.DB.CreatePayment(context.Background(), arg)
	if err != nil {
		log.Errorf("cannot create new payment: %v", err)
		return err
	}

	account, err := postgres.DB.GetAccountByID(context.Background(), arg.AccountID)
	if err != nil {
		return err
	}

	argUpdateBalance := postgr.UpdateAccountBalanceParams{
		ID:      arg.AccountID,
		Balance: account.Balance - arg.Amount,
	}
	_, err = postgres.DB.UpdateAccountBalance(context.Background(), argUpdateBalance)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Payment was successful")
}
