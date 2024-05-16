package main

import (
	"log"
	"net/http"
	"time"

	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/model"
	"github.com/camtrik/gin-blog/internal/routers"
	"github.com/camtrik/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	// log.Printf("global.ServerSetting: %+v\n", global.ServerSetting)
	// log.Printf("global.AppSetting: %+v\n", global.AppSetting)
	// log.Printf("global.DatabaseSetting: %+v\n", global.DatabaseSetting)

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

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
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
