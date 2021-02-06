package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api/login"
	"github.com/rillo-carrillo/restaurant/server/api/register"
	"github.com/rillo-carrillo/restaurant/server/consts"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"github.com/rillo-carrillo/restaurant/server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	name := consts.GetCookieName()
	store := utils.CreateNewStore()
	r.Use(sessions.Sessions(name, store))
	r = register.Register(db, r)
	r = login.Login(db, r)

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

func main() {
	//Creates db Connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  consts.GetDNS(), // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,            // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&orm.User{})

	r := setupRouter(db)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
