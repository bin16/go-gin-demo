package apiv1

import (
	"log"
	"net/http"

	"github.com/bin16/go-gin-demo/db"
	"github.com/bin16/go-gin-demo/models"
	"github.com/bin16/go-gin-demo/utils/queryhelper"
	"github.com/gin-gonic/gin"
)

type userCreateNoteValues struct {
	Content string             `json:"content"`
	Visible models.NoteVisible `json:"visible"`
}

func CreateNote(ctx *gin.Context) {
	values := userCreateNoteValues{}

	ctx.ShouldBindJSON(&values)

	note := models.Note{
		Content:   values.Content,
		Visible:   values.Visible,
		ProfileID: 1,
	}
	if err := db.DB.Create(&note).Error; err != nil {
		// error
	}

	ctx.JSON(http.StatusCreated, note)
}

func Notes(ctx *gin.Context) {
	query := models.Query{}

	ctx.ShouldBindQuery(&query)

	notes := []models.Note{}
	db.DB.Preload("Profile").Offset(query.Offset).Limit(query.Limit).Find(&notes)

	ctx.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func Note(ctx *gin.Context) {
	var u queryhelper.URLWithID

	if err := ctx.ShouldBindUri(&u); err != nil {
		log.Panicf("invalid query uri: %s\n", ctx.Request.URL)
	}

	note := models.Note{}

	if err := db.DB.Preload("Profile").First(&note, "id = ?", u.ID).Error; err != nil {
		log.Panicln(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}
