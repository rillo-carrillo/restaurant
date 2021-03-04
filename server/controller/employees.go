package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
)

//GetEmployees godoc
func (ctr *Controller) GetEmployees(c *gin.Context) {
	var emps []entities.Employee
	if err := db.Connection.Find(&emps).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emps)
}

//CreateEmployee godoc
func (ctr *Controller) CreateEmployee(c *gin.Context) {
	var emp entities.Employee
	if err := c.BindJSON(&emp); err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	if err := db.Connection.Save(&emp).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emp)
}

//DeleteEmployee godoc
func (ctr *Controller) DeleteEmployee(c *gin.Context) {
	pID := c.Param("id")
	id, err := strconv.ParseInt(pID, 10, 0)
	if err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	var emp entities.Employee
	if err := db.Connection.Delete(&emp, "ID=?", id).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emp)
}

//UpdateEmployee godoc
func (ctr *Controller) UpdateEmployee(c *gin.Context) {
	change := c.Param("change")
	var uEmp entities.Employee
	if err := c.BindJSON(&uEmp); err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	var emp entities.Employee
	if err := db.Connection.Find(&emp, "ID = ?", uEmp.ID).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	emp.Name = uEmp.Name
	emp.RestaurantID = uEmp.RestaurantID
	emp.RoleID = uEmp.RoleID
	emp.Username = uEmp.Username
	if change == "true" {
		emp.Password = uEmp.Password
	}

	if err := db.Connection.Save(&emp).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emp)
}
