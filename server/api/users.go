package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"gorm.io/gorm"
)

//User Endpoint Operations
func User(db *gorm.DB, r *gin.Engine) *gin.Engine {
	var user orm.User
	//CreateUser
	r.POST("/user", func(c *gin.Context) {
		err := c.BindJSON(&user)
		if err != nil {
			res := make(map[int]string)
			res[http.StatusBadRequest] = "Bad_Request"
			c.JSON(http.StatusBadRequest, res)
		}
		res := db.Create(&user)
		if res.Error != nil {
			validationError := errors.ValidateUserErrors(res.Error)
			c.JSON(http.StatusOK, validationError)
		}
		var userResponse orm.UserCreationResponse
		userResponse.ID = user.ID
		userResponse.Username = user.Username
		c.JSON(http.StatusOK, userResponse)

	})
	return r
}
