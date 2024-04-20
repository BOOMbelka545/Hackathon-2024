package jwt

import (
	"context"
	signup "donPass/backend/internal/http-server/handlers/accounts/signUp"
	"donPass/backend/internal/storage/postgres"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	signinKey = "klnljdsflskesdkljlskdf"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken(number, password string) (string, error) {
	// Check if account doesn't exist
	account, err := postgres.DB.GetAccountByNumber(context.Background(), number)
	if err == pgx.ErrNoRows {
		return "", echo.NewHTTPError(http.StatusNotFound, "User doesn't exist")
	}

	if account.Password != signup.GeneratePasswordHash(password) {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["id"] = account.ID
	claims["exp"] = time.Now().Add(tokenTTL).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(signinKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetClaims(req *http.Request) (jwt.MapClaims, error) {
	jwtToken := req.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, "Bearer ")
	// if len(splitToken) != 2 {
	// 	c.JSON(internalHttp.StatusUnauthorized, nil)
	// 	return customError.ErrInvalidCredentials
	// }
	reqToken := splitToken[1]
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signinKey), nil
	})
	// Checking token validity
	if !token.Valid {
		log.Errorf("invalid token: %v", token)
	}

	if err != nil {
		log.Infof("Cannot parse JWT token")
		return claims, echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse JWT with claims")
	}

	return claims, nil
}
