package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/internal/model"
	"github.com/camtrik/gin-blog/internal/routers"
	"github.com/camtrik/gin-blog/pkg/logger"
	"github.com/camtrik/gin-blog/pkg/setting"
	"github.com/camtrik/gin-blog/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitId  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title gin-blog
// @version 1.0
// @description A blog system backend framework based on Gin.
// @termsOfService https://github.com/camtrik/gin-blog
func main() {
	// log.Printf("global.ServerSetting: %+v\n", global.ServerSetting)
	// log.Printf("global.AppSetting: %+v\n", global.AppSetting)
	// log.Printf("global.DatabaseSetting: %+v\n", global.DatabaseSetting)
	// global.Logger.Infof(c, "%s: www/%s", "ebbi", "gin-blog")
	if isVersion {
		fmt.Printf("build time: %s\n", buildTime)
		fmt.Printf("build version: %s\n", buildVersion)
		fmt.Printf("git commit id: %s\n", gitCommitId)
		return
	}
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// restart server with a new goroutine
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrAbortHandler {
			log.Fatal("s.ListenAndServe err: %v", err)
		}
	}()

	// wait for interrupt signal
	quit := make(chan os.Signal)
	// receive syscall.SIGINT, syscall.SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// shutdown server with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func setupSetting() error {
	// "," to split multiple config files
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
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
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	if port != "" {
		global.ServerSetting.HttpPort, _ = strconv.Atoi(port)
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	global.JWTSetting.Expire *= time.Second

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

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"gin-blog",
		"localhost:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "server port")
	flag.StringVar(&runMode, "mode", "", "run mode")
	flag.StringVar(&config, "config", "./configs/", "config path")
	flag.BoolVar(&isVersion, "version", false, "show compile info")
	flag.Parse()

	return nil
}
