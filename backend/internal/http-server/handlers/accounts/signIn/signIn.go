package signin

import (
	"donPass/backend/internal/http-server/handlers/accounts/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type signInInput struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

func SignIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		log.Infof("Cannot bind the request: %v \n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something wrong witg request. Please check it :)")
	}

	jwt, err := jwt.GenerateToken(input.Number, input.Password)
	if err != nil {
		log.Errorf("Some problem: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, jwt)
}
