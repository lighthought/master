// Package main Master Guide Backend API
//
// Master Guide Backend 是一个基于Go + Gin + GORM + PostgreSQL + Redis的技艺传承平台后端系统，
// 为前端提供高性能的API服务，包括用户管理、课程管理、社群互动、实时通信等功能。
//
//	Schemes: http, https
//	Host: localhost:8080
//	BasePath: /api/v1
//	Version: 1.0.0
//	Title: Master Guide Backend API
//	Description: Master Guide Backend 是一个基于Go + Gin + GORM + PostgreSQL + Redis的技艺传承平台后端系统，为前端提供高性能的API服务。
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//	- image/*
//
//	Security:
//	- bearer
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"master-guide-backend/internal/api/middleware"
	"master-guide-backend/internal/api/routes"
	"master-guide-backend/internal/container"
	"master-guide-backend/pkg/cache"
	"master-guide-backend/pkg/config"
	"master-guide-backend/pkg/database"
	"master-guide-backend/pkg/logger"

	_ "master-guide-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token
func main() {
	// 加载配置
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs/config.yaml"
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	if err := logger.Init(cfg.Log.Level, cfg.Log.Format, cfg.Log.Output, cfg.Log.FilePath); err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		os.Exit(1)
	}

	logger.Info("启动Master Guide后端服务")

	// 连接数据库
	dbConfig := &database.Config{
		Host:            cfg.Database.Host,
		Port:            cfg.Database.Port,
		User:            cfg.Database.User,
		Password:        cfg.Database.Password,
		DBName:          cfg.Database.DBName,
		SSLMode:         cfg.Database.SSLMode,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	}

	var db *gorm.DB
	if err := database.Connect(dbConfig); err != nil {
		logger.Warn("连接数据库失败，将启动无数据库模式", logger.String("error", err.Error()))
		logger.Info("注意：某些功能可能不可用")
		db = nil
	} else {
		defer database.Close()
		logger.Info("数据库连接成功")
		db = database.GetDB()
	}

	// 连接Redis
	redisConfig := &cache.Config{
		Host:         cfg.Redis.Host,
		Port:         cfg.Redis.Port,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	}

	if err := cache.Connect(redisConfig); err != nil {
		logger.Warn("连接Redis失败，将使用内存缓存", logger.String("error", err.Error()))
	} else {
		defer cache.Close()
		logger.Info("Redis连接成功")
	}

	// 初始化依赖注入容器
	var the_container *container.Container
	if db != nil {
		the_container = container.NewContainer(db, cfg)
		logger.Info("依赖注入容器初始化成功")
	} else {
		logger.Warn("数据库未连接，部分功能不可用")
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	engine := gin.New()

	// 添加中间件
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recovery())
	engine.Use(middleware.CORS(cfg.CORS))
	engine.Use(middleware.RequestID())
	engine.Use(gin.Recovery())

	// 设置静态文件服务
	engine.Static("/static", "./static")
	engine.Static("/uploads", "./static/uploads")

	// 设置路由
	if the_container != nil {
		routes.SetupRoutes(engine, cfg, the_container.AuthHandler, the_container.UserHandler, the_container.MentorHandler, the_container.CourseHandler, the_container.AppointmentHandler)
	} else {
		// 如果数据库未连接，使用默认路由
		routes.SetupRoutes(engine, cfg, nil, nil, nil, nil, nil)
	}

	// 添加Swagger文档路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查路由
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
			"service":   "master-guide-backend",
			"version":   "1.0.0",
		})
	})

	// 创建HTTP服务器
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        engine,
		ReadTimeout:    time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.Server.MaxHeaderBytes,
	}

	// 启动服务器
	go func() {
		logger.Info("服务器启动", logger.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("服务器启动失败", logger.String("error", err.Error()))
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("正在关闭服务器...")

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("服务器关闭失败", logger.String("error", err.Error()))
	}

	logger.Info("服务器已关闭")
}
