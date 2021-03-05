package db

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var Engine *xorm.Engine

func InitXorm() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:1350017101@tcp(120.78.122.145:3306)/music")
	Engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	//日志
	//logFile, err := os.OpenFile("logs/music.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	//if err != nil {
	//	panic("log file can't init")
	//}
	logger := log.NewSimpleLogger(os.Stdout)
	logger.SetLevel(log.LOG_INFO)
	Engine.SetLogger(log.NewLoggerAdapter(logger))
	Engine.ShowSQL(true)
	if err != nil {
		panic("mysql connecting error")
	}
}
