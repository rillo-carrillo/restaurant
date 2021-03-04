package login

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/consts"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login User
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {object} orm.UserResponse
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/orders [post]
func Login(c *gin.Context) {
	var user entities.Employee
	session := sessions.Default(c)
	err := c.BindJSON(&user)
	password := user.Password
	//Validate Request match Struct
	if err != nil {
		var res errors.ServerError
		res.Message = "Somethin went wrong with request: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := db.Connection.First(&user, "username= ?", user.Username)
	//Validate is user is found.
	if res.Error != nil {
		validationError := errors.ValidateUserErrors(res.Error)
		c.JSON(http.StatusOK, validationError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//Validate password match
	if err != nil {
		var res errors.DBErrorResponse
		res.Field = "password"
		res.Message = "Password did not match."
		c.JSON(http.StatusOK, res)
		return
	}
	//create a cookieSession
	session.Set(consts.CookieUserIdKey, user.ID)
	session.Set("roleID", user.RoleID)
	err = session.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"Username": user.Username,
		"RoleID":   user.RoleID})

}

// Logout godoc
// @Summary Loout User
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {object} orm.StatusResponse
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/orders [delete]
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

//Handler Define routes of type.
func Handler(g *gin.RouterGroup) {
	login := g.Group("/login")
	{
		login.POST("", Login)
		login.DELETE("", Logout)
	}
}
