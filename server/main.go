package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rillo-carrillo/restaurant/server/api"
	"github.com/rillo-carrillo/restaurant/server/consts"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r = api.User(db, r)
	// Ping test
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "pong")
	// })

	// // Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	//user := c.Params.ByName("name")
	// 	//value, ok := db[user]
	// 	// if ok {
	// 	// 	c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	// } else {
	// 	// 	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	// }
	// })

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

	// /* example curl for /admin with basicauth header
	//    Zm9vOmJhcg== is base64("foo:bar")

	// 	curl -X POST \
	//   	http://localhost:8080/admin \
	//   	-H 'authorization: Basic Zm9vOmJhcg==' \
	//   	-H 'content-type: application/json' \
	//   	-d '{"value":"bar"}'
	// */
	// authorized.POST("admin", func(c *gin.Context) {
	// 	//user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		//db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  consts.GetDNS(), // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,            // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&orm.User{})
	r := setupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
