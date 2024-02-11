package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func MustAuthedMiddleware(role models.UserRole) func(*gin.Context) {
	return func(ctx *gin.Context) {
		log.Printf("MustAuthedMiddleware() -> %s\n", ctx.Request.URL)

		ts, err := ctx.Cookie("TOKEN")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false})
			return
		}

		data, err := ParseTokenString(ts)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false})
			return
		}

		user := models.User{}
		if err := db.DB.Where("id = ?", data.UserID).First(&user).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false})
			return
		}

		if user.Role < role {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"ok": false})
			return
		}

		ctx.Set("userId", user.ID)
		ctx.Next()
	}
}

type TokenData struct {
	UserID int64 `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateTokenString(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenData{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 30)),
			Issuer:    "demo",
		},
	})

	return token.SignedString(getJWTSecret())
}

func ParseTokenString(ts string) (*TokenData, error) {
	token, err := jwt.ParseWithClaims(ts, &TokenData{}, func(t *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	data, ok := token.Claims.(*TokenData)
	if !ok {
		return nil, fmt.Errorf("invalid token type")
	}

	return data, nil
}
