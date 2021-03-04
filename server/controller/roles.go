package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
)

//GetRoles godoc
func (ctr *Controller) GetRoles(c *gin.Context) {
	var roles []entities.Role
	if err := db.Connection.Order("code").Find(&roles).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, roles)
}

//CreateRole godoc
func (ctr *Controller) CreateRole(c *gin.Context) {
	var role entities.Role
	if err := c.BindJSON(&role); err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	if err := db.Connection.Save(&role).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, role)
}

//DeleteRole godoc
func (ctr *Controller) DeleteRole(c *gin.Context) {
	pID := c.Param("id")
	id, err := strconv.ParseInt(pID, 10, 0)
	if err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	var role entities.Role
	if err := db.Connection.Delete(&role, "code=?", id).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, role)
}

//UpdateRole godoc
func (ctr *Controller) UpdateRole(c *gin.Context) {
	var uRole entities.Role
	if err := c.BindJSON(&uRole); err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	var role entities.Role
	if err := db.Connection.Find(&role, "code = ?", uRole.Code).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	role.Name = uRole.Name
	if err := db.Connection.Save(&role).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, role)
}
