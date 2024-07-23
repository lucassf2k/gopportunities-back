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
	openingsReponse := []schemas.OpeningResponse{}
	for _, opening := range openings {
		openingResponse := schemas.OpeningResponse{
			ID:        opening.ID,
			Role:      opening.Role,
			Company:   opening.Company,
			Location:  opening.Location,
			Remote:    opening.Remote,
			Link:      opening.Link,
			Salary:    opening.Salary,
			CreatedAt: opening.CreatedAt,
			UpdatedAt: opening.UpdatedAt,
			DeletedAt: opening.DeletedAt.Time,
		}
		openingsReponse = append(openingsReponse, openingResponse)
	}
	sendSucess(ctx, "list-openings", openingsReponse)
}
