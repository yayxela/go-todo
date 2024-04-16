package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/yayxela/go-todo/api/swagger"
	"github.com/yayxela/go-todo/internal/config"
	"github.com/yayxela/go-todo/internal/db"
	"github.com/yayxela/go-todo/internal/logger"
	"github.com/yayxela/go-todo/internal/middleware"
	"github.com/yayxela/go-todo/internal/todo"
	"github.com/yayxela/go-todo/internal/validate"
)

func main() {
	// создание нового логгера
	log, err := logger.New()
	if err != nil {
		panic(err)
	}
	defer func() { _ = log.Sync() }()

	// создание нового конфига
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// указание временной зоны сервера
	loc, err := time.LoadLocation(cfg.AppConfig.Timezone)
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc

	// обновление параметров для swagger
	swagger.SwaggerInfo.Title = cfg.SwaggerConfig.Title
	swagger.SwaggerInfo.Version = cfg.SwaggerConfig.Version
	swagger.SwaggerInfo.Description = cfg.SwaggerConfig.Description
	swagger.SwaggerInfo.Host = cfg.AppConfig.Host
	swagger.SwaggerInfo.BasePath = cfg.AppConfig.BasePath

	// создание нового кастомного валидатора
	validator := validate.New()

	// новое подключение е бд
	ctx := context.Background()
	idb, err := db.New(ctx, cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	// создание веб-сервера
	server := gin.New()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// создание и регистрация мидлвейра
	mw := middleware.New(log)
	api := server.Group("/api", mw.Error, mw.Panic)
	v1 := api.Group("/v1")

	// создание сервиса для тасков
	todoService := todo.New(idb)
	todo.RegisterHandlers(v1, todoService, validator)

	// запуск севера
	if err = server.Run(cfg.AppConfig.Port); err != nil {
		log.Fatal(err)
	}
	go func() {
		if err = server.Run(); err != nil {
			log.Errorf("error: %s\n", err)
		}
	}()

	// graceful shutdown
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_ = idb.Disconnect(ctx)
	log.Info("Server exiting")
}
