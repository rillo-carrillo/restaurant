package login

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	"github.com/rillo-carrillo/restaurant/server/consts"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Login Endpoint
func Login(db *gorm.DB, r *gin.Engine) *gin.Engine {
	var user orm.User

	//CreateUser
	r.POST("/login", func(c *gin.Context) {
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
		res := db.First(&user, "username= ?", user.Username)
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
		err = session.Save()
		if err != nil {
			fmt.Println(err)
			return
		}
		//Generate response
		response := orm.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		}
		c.JSON(http.StatusOK, response)

	})
	return r
}
