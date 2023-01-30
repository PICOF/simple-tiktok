package db

import "gorm.io/gorm"

type DataSourceConfig struct {
	Host     string
	Port     string
	Auth     bool
	Username string
	Password string
	MaxConn  int
	MaxOpen  int
	Timeout  int
}

func init() {

}

func GetDBConfig() {

}

func getConnection() {
	gorm.Open("mysql", "username:password@/database?charset=utf8&parseTime=True&loc=Local")
}
