package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassf2k/gopportunities-back/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx, "list-openings", openings)
}
