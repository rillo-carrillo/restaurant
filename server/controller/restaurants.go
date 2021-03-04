package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
)

// GetRestaurants godoc
// @Summary Get all restaurants
// @Tags restaurant
// @Produce  json
// @Success 200 {object} []entities.Category
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/restaurants [get]
func (ctr *Controller) GetRestaurants(c *gin.Context) {
	var restaurants []entities.Restaurant
	if err := db.Connection.Find(&restaurants).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = "Restaurants missing."
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

// CreateRestaurant godoc
// @Summary Create a new restaurant
// @Tags restaurant
// @Produce  json
// @Success 200 {object} []entities.Category
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/restaurants [post]
func (ctr *Controller) CreateRestaurant(c *gin.Context) {
	var restaurant entities.Restaurant
	if err := c.BindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	if err := db.Connection.Create(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

// DeleteRestaurant godoc
// @Summary Delete a new restaurant
// @Tags restaurant
// @Produce  json
// @Success 200 {object} entities.Restaurant
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/restaurants [delete]
func (ctr *Controller) DeleteRestaurant(c *gin.Context) {
	parID := c.Param("id")
	var restaurant entities.Restaurant
	id, err := strconv.ParseInt(parID, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	restaurant.ID = uint(id)
	// if err := db.Connection.Find(&restaurant).Where("id=?", restaurant.ID).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": string(err.Error()),
	// 	})
	// 	return
	// }
	if err := db.Connection.Delete(&restaurant).Where("id=?", restaurant.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

// UpdateRestaurant godoc
// @Summary Update restaurant
// @Tags restaurant
// @Produce  json
// @Success 200 {object} entities.Restaurant
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/restaurants [put]
func (ctr *Controller) UpdateRestaurant(c *gin.Context) {
	var newRest entities.Restaurant
	if err := c.BindJSON(&newRest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	var rest entities.Restaurant
	if err := db.Connection.Find(&rest, "ID=?", newRest.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	rest.Name = newRest.Name
	rest.Address = newRest.Address
	rest.CP = newRest.CP
	rest.Logo = newRest.Logo
	rest.Number = newRest.Number
	rest.RFC = newRest.RFC
	if err := db.Connection.Save(&rest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": string(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
