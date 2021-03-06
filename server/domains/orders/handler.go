package orders

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/errors"
	apiutils "github.com/rillo-carrillo/restaurant/server/api/utils"
	"github.com/rillo-carrillo/restaurant/server/consts"
)

// CreateOrder godoc
// @Summary Create an order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {object} model.OrderCreateResponse
// @Failure 400 {object} errors.ServerError
// @Failure 404 {object} errors.ServerError
// @Failure 500 {object} errors.ServerError
// @Router /v1/orders [post]
func CreateOrder(c *gin.Context) {
	var or OrderRequest
	session := sessions.Default(c)
	userID := session.Get(consts.CookieUserIdKey)

	ID, err := apiutils.GetUserID(userID)
	if err != nil {
		var res errors.ServerError
		res.Message = "Somethin went wrong with request: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if sess := session.Get(consts.CookieUserIdKey); sess == nil {
		var res errors.ServerError
		res.Message = "Not able to process Order, please log in."
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.BindJSON(&or); err != nil {
		var res errors.ServerError
		res.Message = "Somethin went wrong with request: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	id, err := CreateOrderLogic(ID, &or)
	if err != nil {
		var res errors.ServerError
		res.Message = "Somethin went wrong with request: " + err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := OrderCreateResponse{
		ID:     int(id),
		Status: true,
	}
	c.JSON(http.StatusOK, res)
}

//Handler Define routes of type.
func Handler(g *gin.RouterGroup) {
	orders := g.Group("/orders")
	{
		orders.POST("", CreateOrder)
	}
}
