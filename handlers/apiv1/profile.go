package apiv1

import (
	"net/http"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/models"
	"github.com/gin-gonic/gin"
)

// get me
func Profile(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	profile := models.Profile{}
	db.DB.Where("user_id = ?", userId).First(&profile)

	ctx.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}
