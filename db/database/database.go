package database

import (
	"khaerus/mini-wallet/conf"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	// "gorm.io/driver/postgres"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error
	db_host := conf.Viper("APP_DB_HOST")  
	db_port := conf.Viper("APP_DB_PORT")  
	db_user := conf.Viper("APP_DB_USERNAME")  
	db_pass := conf.Viper("APP_DB_PASSWORD")
	db_name := conf.Viper("APP_DB_NAME")  
	db_sslmode := conf.Viper("APP_DB_SSL")

	connectionString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString, db_host, db_port, db_user, db_pass, db_name, db_sslmode)
    DBCon, err = gorm.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }
}