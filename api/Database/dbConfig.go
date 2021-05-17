package Database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"cryptoCurrencies/Utils"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     uint
	DBName   string
}

var DB *gorm.DB

const (
	dbConfigFile = "Config/DatabaseConfiguration.csv"
	username     = "root"
	password     = "Bala!989"
	host         = "localhost"
	port         = 3306
)

func BuildDBConfig() *DBConfig {
	dbConfig := &DBConfig{
		User:     username,
		Password: password,
		Host:     host,
		Port:     port,
	}

	return dbConfig
}

func DBConnectionString(dbConfig *DBConfig) string {
	connectionStr := Utils.ProcessDBConnectionString(Utils.ReadDBConfiguration(dbConfigFile))

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/"+connectionStr,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
	)
}
