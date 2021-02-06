package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"github.com/rillo-carrillo/restaurant/server/utils"
	"gorm.io/gorm"
)

type ServeError struct {
	message string
}

//User Endpoint Operations
func User(db *gorm.DB, r *gin.Engine) *gin.Engine {
	var user orm.User
	//CreateUser
	r.POST("/register", func(c *gin.Context) {
		err := c.BindJSON(&user)
		if err != nil {
			var res ServeError
			res.message = "Somethin went wrong with request: " + err.Error()
			c.JSON(http.StatusBadRequest, res)
			return
		}
		user.Password, err = utils.GeneratePassword(user.Password)
		if err != nil {
			var res ServeError
			res.message = "Something went wrong when hashing password: " + err.Error()
			c.JSON(http.StatusBadRequest, res)
			return
		}
		res := db.Create(&user)
		if res.Error != nil {
			validationError := errors.ValidateUserErrors(res.Error)
			c.JSON(http.StatusOK, validationError)
			return
		}
		var userResponse orm.UserCreationResponse
		userResponse.ID = user.ID
		userResponse.Username = user.Username
		c.JSON(http.StatusOK, userResponse)

	})
	return r
}
