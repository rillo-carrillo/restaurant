package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/consts"
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
	"github.com/rillo-carrillo/restaurant/server/model"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"github.com/rillo-carrillo/restaurant/server/utils"
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
func (ctr *Controller) Login(c *gin.Context) {
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
	//Generate response
	response := orm.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		RoleID:   user.RoleID,
	}
	c.JSON(http.StatusOK, response)

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
func (ctr *Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	res := orm.StatusResponse{Status: true}
	c.JSON(http.StatusOK, res)
}

// CreateUser godoc
// @Summary Create new User
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {object} orm.UserResponse
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/login [post]
func (ctr *Controller) CreateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		var res errors.ServerError
		res.Message = "Somethin went wrong with request: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user.Password, err = utils.GeneratePassword(user.Password)
	if err != nil {
		var res errors.ServerError
		res.Message = "Something went wrong when hashing password: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := db.SaveToDB(&user)
	if res != nil {
		validationError := errors.ValidateUserErrors(res)
		c.JSON(http.StatusConflict, validationError)
		return
	}
	var userResponse orm.UserResponse
	userResponse.ID = user.ID
	userResponse.Username = user.Username
	c.JSON(http.StatusOK, userResponse)
}

//Me  Get user info from session and return it
func (ctr *Controller) Me(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get(consts.CookieUserIdKey)
	roleid := session.Get("roleID")
	var user entities.Employee
	if userid != nil && roleid != nil {
		user.ID = userid.(uint)
		user.RoleID = roleid.(int)
	}
	c.JSON(http.StatusOK, user)

}
