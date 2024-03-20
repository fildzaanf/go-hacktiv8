package middlewares

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

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
	logrus.Infof("generating token for user with ID: %s, Role: %s", id, role)

	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["id"] = id
	tokenClaims["role"] = role
	tokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		logrus.Errorf("error generating token: %v", err)
		return "", err
	}
	logrus.Infof("token generated successfully: %s", tokenString)

	return tokenString, nil
}

func VerifyToken(c *gin.Context) (string, string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", "", errors.New("missing authorization token")
	}

	tokenString = tokenString[len("bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
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
