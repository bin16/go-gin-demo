package apiv1

import (
	"log"
	"net/http"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/models"
	"github.com/bin16/go-gin-demo/utils/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpValues struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

func SignUp(ctx *gin.Context) {
	values := SignUpValues{}
	ctx.ShouldBindJSON(&values)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(values.Password), 12)

	user := models.User{
		Username: values.Username,
		Password: string(hashedPassword),
		Profile: &models.Profile{
			Name: values.Username,
		},
	}

	if err := db.DB.Create(&user).Error; err != nil {
		// ...
	}

	ctx.JSON(http.StatusCreated, user)
}

type SignInValues struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(ctx *gin.Context) {
	values := SignInValues{}
	ctx.ShouldBindJSON(&values)

	user := models.User{}

	if err := db.DB.Where("username = ?", values.Username).First(&user).Error; err != nil {
		// user not found
		log.Printf("user not found: %s", values.Username)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok": false,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(values.Password)); err != nil {
		// wrong password
		log.Printf("wrong password: %s", values.Username)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok": false,
		})
	}

	token, err := auth.GenerateTokenString(user.ID)
	if err != nil {
		log.Printf("signing token: %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok": false,
		})
	}

	ctx.SetCookie("TOKEN", token, 365*24*60*60, "/", "", true, true)
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
