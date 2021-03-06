package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-programing-tour-book/blog-service/global"
	"github.com/go-programing-tour-book/blog-service/internal/model"
	"github.com/go-programing-tour-book/blog-service/internal/routers"
	"github.com/go-programing-tour-book/blog-service/pkg/logger"
	"github.com/go-programing-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go语言编程之旅：一起用Go做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "lw", "blog-service")
	s.ListenAndServe()
}

// 全局配置文件读取
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

//日志
func setupLogger() error {
	// log.Printf("global.ServerSetting: %#v", *global.ServerSetting)    //%#v，打印结构体的名称，key，value
	// log.Printf("global.AppSetting: %+v", *global.AppSetting)          //%+v, 打印结构体的key， value
	// log.Printf("global.DataBaseSetting: %v", *global.DataBaseSetting) //%v，打印结构体的value
	global.Logger = logger.Newlogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,  //MB
		MaxAge:    10,   //10 days
		LocalTime: true, //本地时间
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

//数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting) //不用:=，用:=的话，等于在该函数内生成了一个新的仅在函数能起作用的同名变量global.DBEngine，函数体外的全局global.DBEngine还是nil
	if err != nil {
		return err
	}
	return nil
}
