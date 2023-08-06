package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleare(c *gin.Context) {
	err := validateToken(c.Request.Header)

	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	c.Next()
}

func validateToken(header http.Header) error {
	token, err := getTokenFromHeader(header)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func getTokenFromHeader(header http.Header)(*jwt.Token, error){
	bearerToken := header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return nil, errors.New("invalid")
	}

	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(os.Getenv("API_SECRET")), nil
    })

	if err != nil {
		return nil, err
	}

	return token, nil
}
