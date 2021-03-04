package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/consts"
	"github.com/rillo-carrillo/restaurant/server/db"
	_ "github.com/rillo-carrillo/restaurant/server/docs"
	"github.com/rillo-carrillo/restaurant/server/domains/categories"
	"github.com/rillo-carrillo/restaurant/server/domains/employees"
	"github.com/rillo-carrillo/restaurant/server/domains/login"
	"github.com/rillo-carrillo/restaurant/server/domains/me"
	"github.com/rillo-carrillo/restaurant/server/domains/orders"
	"github.com/rillo-carrillo/restaurant/server/domains/restaurants"
	"github.com/rillo-carrillo/restaurant/server/domains/roles"
	eutils "github.com/rillo-carrillo/restaurant/server/entities/utils"
	"github.com/rillo-carrillo/restaurant/server/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	name := consts.GetCookieName()
	// store := cookie.NewStore([]byte(name))
	// store.Options(sessions.Options{
	// 	Secure: false,
	// 	MaxAge: 1000 * 60 * 60 * 24 * 3,
	// })
	store := utils.CreateNewStore()
	corsConf := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.15.6:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowHeaders:     []string{"content-type"},
	}
	r.Use(sessions.Sessions(name, store),
		cors.New(corsConf))

	v1 := r.Group("/v1")
	{
		orders.Handler(v1)
		restaurants.Handler(v1)
		roles.Handler(v1)
		me.Handler(v1)
		employees.Handler(v1)
		categories.Handler(v1)
		login.Handler(v1)

	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// // Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	//user := c.Params.ByName("name")
	// // Authorized group (uses gin.BasicAuth() middleware)
	// // Same than:
	// // authorized := r.Group("/")
	// // authorized.Use(gin.BasicAuth(gin.Credentials{
	// //	  "foo":  "bar",
	// //	  "manu": "123",
	// //}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	return r
}

// @title Restaurant.io API
// @version 1.0
// @description This is the api for the Restaurant.io application
// @termsOfService http://swagger.io/terms/

// @contact.name restaurant.io Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	//Creates db Connection

	db := db.GetConnection()

	if err := eutils.CreateEntities(db); err != nil {
		fmt.Println(err)
	}
	eutils.InitializeDatabase(db)
	r := setupRouter(db)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
