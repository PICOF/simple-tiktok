package util

import (
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LoggerConfig struct {
	MsgLogsRefreshCycle int
	ErrLogsRefreshCycle int
}

var (
	ErrLog       *log.Logger
	MsgLog       *log.Logger
	loggerConfig LoggerConfig
)

func setLogger() {
	mf, err := os.OpenFile("log/common/"+time.Now().Format("2006-01-02T3PM")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("日志创建失败！error:", err)
	}
	MsgLog = log.New(io.MultiWriter(mf, os.Stdout), "", log.LstdFlags)
	ef, err := os.OpenFile("log/error/"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("日志创建失败！error:", err)
	}
	ErrLog = log.New(io.MultiWriter(ef, os.Stdout), "[Error]", log.Lshortfile|log.LstdFlags)
}
func cleanLogs() {
	cleanMsgLogs()
	cleanErrLogs()
}

func cleanMsgLogs() {
	date := time.Now().Add(-time.Duration(loggerConfig.MsgLogsRefreshCycle) * time.Hour).Format("2006-01-02T3PM")
	root := "log/common"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if info.Name()[:13] == date {
				err := os.Remove(path)
				if err != nil {
					ErrLog.Println("刪除信息日志時出現問題！Error:", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		ErrLog.Println("清理信息日志文件時出現問題！Error:", err)
	}
}
func cleanErrLogs() {
	date := time.Now().AddDate(0, 0, -loggerConfig.ErrLogsRefreshCycle).Format("2006-01-02")
	root := "log/error"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if info.Name()[:10] == date {
				err := os.Remove(path)
				if err != nil {
					ErrLog.Println("刪除錯誤日志時出現問題！Error:", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		ErrLog.Println("清理錯誤日志文件時出現問題！Error:", err)
	}
}
func RenewLoggers() {
	for {
		//每个整点进行日志清理
		time.Sleep(time.Second * time.Duration(60*(60-time.Now().Minute())-time.Now().Second()))
		setLogger()
		go cleanLogs()
	}
}

func GetLoggerConfig() {
	err := viper.Sub("logger").Unmarshal(&loggerConfig)
	if err != nil {
		log.Fatal("获取日志相关配置失败，error:", err)
	}
}

func InitLoggers() {
	GetLoggerConfig()
	setLogger()
}
