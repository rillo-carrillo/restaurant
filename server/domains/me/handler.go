package me

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/consts"
	"github.com/rillo-carrillo/restaurant/server/entities"
)

//Me  Get user info from session and return it
func Me(c *gin.Context) {
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

//Handler Define routes of type.
func Handler(g *gin.RouterGroup) {
	me := g.Group("/me")
	{
		me.GET("", Me)
	}
}
