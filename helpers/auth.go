package helpers

import (
	"errors"
	"mygram/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id int, email string) (string, error) {
	// claims := jwt.MapClaims{
	// 	"id":    id,
	// 	"email": email,
	// 	"jwtStandardClaims": jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	// 	},
	// }

	claims := &jwtCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := parseToken.SignedString([]byte(config.Cfg.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(c *gin.Context) (any, error) {
	errResponse := errors.New("token invalid or expired")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(config.Cfg.JWT_SECRET_KEY), nil
	})

	if !token.Valid {
		return nil, errResponse
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errResponse
	}

	return claims, nil
}
