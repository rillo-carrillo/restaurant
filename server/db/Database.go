package db

import (
	"github.com/rillo-carrillo/restaurant/server/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Connection to the DB
var Connection *gorm.DB

//GetConnection creates a connection with DB.
func GetConnection() *gorm.DB {
	if Connection == nil {
		d, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  consts.GetDNS(), // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true,            // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			panic("failed to connect database")
		}
		Connection = d
	}

	return Connection
}

//SaveToDB create the object in DB.
func SaveToDB(obj interface{}) error {
	if err := Connection.Create(obj); err.Error != nil {
		return err.Error
	}
	return nil
}
