package dal

import (
	"fmt"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var DBConfig *viper.Viper

const ConfigName = "mysql"

func init() {
	DB = getDB()
}

func getDB() (myDB *gorm.DB) {
	DBConfig = config.GetConfig(ConfigName)
	var err error
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		DBConfig.GetString("mysql.dsn.username"),
		DBConfig.GetString("mysql.dsn.password"),
		DBConfig.GetString("mysql.dsn.host"),
		DBConfig.GetInt("mysql.dsn.port"),
		DBConfig.GetString("mysql.dsn.database"),
		DBConfig.GetString("mysql.dsn.charset"),
		DBConfig.GetBool("mysql.dsn.parseTime"),
		DBConfig.GetString("mysql.dsn.location"),
	)
	myDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSN,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}))
	if err != nil {
		fmt.Println("数据库链接error:", err)
	}

	sqlDB, err := myDB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(DBConfig.GetInt("mysql.connPool.maxIdleConn"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(DBConfig.GetInt("mysql.connPool.maxOpenConn"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(DBConfig.GetInt("mysql.connPool.connMaxLifetime")) * time.Hour)
	return
}
