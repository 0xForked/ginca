package config

import (
	"github.com/aasumitro/gorest/src/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func (config AppConfig) SetupDatabaseConnection() {
	// open database connection
	db, err := gorm.Open("mysql", viper.GetString(`DB_DSN_URL`))
	// error validator database connection
	if err != nil {
		panic("Failed to connect to database!")
	}
	// Migrate the schema
	db.AutoMigrate(&domain.Example{})
	// set database connection for global use
	setDBConnection(db)
}

func setDBConnection(DB *gorm.DB) {
	db = DB
}

func (config AppConfig) GetDatabaseConnection() *gorm.DB {
	return db
}