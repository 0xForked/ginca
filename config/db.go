package config

import (
	"github.com/aasumitro/ginca/logs"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func (config AppConfig) SetupDatabaseConnection() {
	// open database connection
	db, err := gorm.Open(
		viper.GetString(`DB_CONNECTION`),
		viper.GetString(`DB_DSN_URL`))
	// error validator database connection
	if err != nil {
		logs.AppError.Println(err)
	}
	// Migrate the schema
	db.AutoMigrate(&domain.Example{})
	// set database connection for global use
	setDBConnectionAndStatus(db)
}

func setDBConnectionAndStatus(DB *gorm.DB) {
	db = DB
}

func (config AppConfig) GetDatabaseConnection() *gorm.DB {
	return db
}

func (config AppConfig) GetDatabaseStatus() string {
	if err := db.DB().Ping(); err != nil {
		return domain.MySQLUnavailable.Error()
	}

	return domain.MySQLAvailable
}