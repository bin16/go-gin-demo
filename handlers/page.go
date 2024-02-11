package handlers

import (
	"net/http"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/models"
	"github.com/bin16/go-gin-demo/utils/queryhelper"
	"github.com/gin-gonic/gin"
)

func Notes(ctx *gin.Context) {
	q := queryhelper.QueryBody{}
	ctx.ShouldBind(&q)

	notes := []models.Note{}
	db.DB.Debug().Preload("Profile").Order("created_at DESC").Offset(q.Offset).Limit(q.Limit).Find(&notes)

	ctx.HTML(http.StatusOK, "notes.html", gin.H{
		"notes": notes,
	})
}
