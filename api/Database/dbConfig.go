package Database

import (
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

func BuildDBConfig(fileName string) string {
	return Utils.ProcessDBServerConfiguration(Utils.ReadDBConfiguration(fileName))

}

func BuildDBParams(fileName string) string {
	return Utils.ProcessDBConnectionString(Utils.ReadDBConfiguration(fileName))
}
