package controller

import (
	"github.com/gin-gonic/gin"
	"go_api/helper"
	"go_api/model"
	"net/http"
)

// @Summary     Add entry
// @Description Add a new entry
// @Accept      json
// @Produce     json
// @Param       entry body     model.Entry   true "Entry data"
// @Success     201   {object} model.Entry   "Created"
// @Failure     400   {object} ErrorResponse "Error message"
// @Security    ApiKeyAuth
// @Router      /api/entry [post]
func AddEntry(context *gin.Context) {
	var input model.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save() //original: input.save(). Maybe make save() in entry.go

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, savedEntry)
}

// @Summary     Get all entries
// @Description Retrieve all entries for the current user
// @Accept      json
// @Produce     json
// @Success     200 {array}  model.Entry   "List of entries"
// @Failure     400 {object} ErrorResponse "Error message"
// @Security    ApiKeyAuth
// @Router      /api/entry [get]
func GetAllEntries(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, user.Entries)
}
