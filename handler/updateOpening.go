package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassf2k/gopportunities-back/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id != "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
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
	sendSucess(ctx, "updating-opening", openingResponse)
}
