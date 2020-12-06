package config

import (
	"fmt"
	"github.com/aasumitro/ginca/logs"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var rdbConn *gorm.DB

func (config AppConfig) SetupRDBConnection() {
	// open up (rdb) relation database connection
	conn, err := gorm.Open(
		viper.GetString(`DB_CONNECTION`),
		viper.GetString(`DB_DSN_URL`),
	)
	// error validator database connection
	if err != nil {
		logs.AppError.Println(fmt.Sprintf(
			"failed to connect to relation database, cause: %s",
			err,
		))
	}
	// Migrate the schema
	conn.AutoMigrate(&domain.Example{})
	// set the rdb connection for global usage
	setRDBConnection(conn)
}

func setRDBConnection(currentRDBConnection *gorm.DB) {
	rdbConn = currentRDBConnection
}

func (config AppConfig) GetRDBConnection() *gorm.DB {
	return rdbConn
}

func (config AppConfig) GetRDBStatus() string {
	if err := rdbConn.DB().Ping(); err != nil {
		return domain.MySQLUnavailable.Error()
	}

	return domain.MySQLAvailable
}