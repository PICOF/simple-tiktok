package main

import (
	"tiktok/server/config"
	"tiktok/server/util"
)

func init() {
	config.LoadConfig()
	util.InitLoggers()
	go util.RenewLoggers()
	config.BootConfigMonitor()
}

func main() {
	for {
	}
}
