package employees

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
	"github.com/rillo-carrillo/restaurant/server/utils"
)

//GetEmployees godoc
func GetEmployees(c *gin.Context) {
	var emps []entities.Employee
	if err := db.Connection.Find(&emps).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	for idx := range emps {
		emp := &emps[idx]
		emp.Password = "*****"
	}
	c.JSON(http.StatusOK, emps)
}

//CreateEmployee godoc
func CreateEmployee(c *gin.Context) {
	var emp entities.Employee
	if err := c.BindJSON(&emp); err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	emp.Password, _ = utils.GeneratePassword(emp.Password)
	if err := db.Connection.Save(&emp).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emp)
}

//DeleteEmployee godoc
func DeleteEmployee(c *gin.Context) {
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
func UpdateEmployee(c *gin.Context) {
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
	if len(uEmp.Password) > 3 {
		emp.Password, _ = utils.GeneratePassword(uEmp.Password)
	}

	if err := db.Connection.Save(&emp).Error; err != nil {
		var res errors.DBErrorResponse
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, emp)
}

//Handler Define routes of type.
func Handler(g *gin.RouterGroup) {
	employees := g.Group("/empleados")
	{
		employees.GET("", GetEmployees)
		employees.POST("", CreateEmployee)
		employees.DELETE(":id", DeleteEmployee)
		employees.PUT("", UpdateEmployee)
	}
}
