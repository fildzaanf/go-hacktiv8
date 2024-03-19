package middlewares

import (
	"assignment-2/app/configs"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var config *configs.Configuration

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, err := VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GenerateToken(id string, role string) (string, error) {
	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["id"] = id
	tokenClaims["role"] = role
	tokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(config.JWT.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(c *gin.Context) (string, string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", "", errors.New("missing authorization token")
	}

	tokenString = tokenString[len("bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := []byte(config.JWT.JWT_SECRET)
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid authorization token")
	}

	id := claims["id"].(string)
	role := claims["role"].(string)

	return id, role, nil
}
