package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
)

// GetCategories godoc
// @Summary Get all categories
// @Tags Categories
// @Produce  json
// @Success 200 {object} []entities.Category
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/categories [get]
func (ctr *Controller) GetCategories(c *gin.Context) {
	var categorias []entities.Category
	if err := db.Connection.Find(&categorias).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = "Categories missing."
		c.JSON(http.StatusOK, res)
	}
	c.JSON(http.StatusOK, categorias)
}
