package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassf2k/gopportunities-back/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating openig: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on databse")
		return
	}
	sendSucess(ctx, "Create Opening", opening)
}
