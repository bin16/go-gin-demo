package main

import (
	"net/http"
	"time"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/handlers"
	"github.com/bin16/go-gin-demo/handlers/apiv1"
	"github.com/bin16/go-gin-demo/models"
	"github.com/bin16/go-gin-demo/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env", ".env.dev1")
	db.Connect()
}

func main() {
	r := gin.New()
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"date": time.Now(),
		})
	})
	r.GET("/notes", handlers.Notes)

	r.POST("/auth/signup", apiv1.SignUp)
	r.POST("/auth/signin", apiv1.SignIn)

	api := r.Group("/api/v1", auth.MustAuthedMiddleware(models.NormalUser))

	api.POST("/", apiv1.CreateNote)
	api.GET("/", apiv1.Notes)
	api.GET("/:id", apiv1.Note)

	api.GET("/profile", apiv1.Profile)

	r.Run(":8080")
}
