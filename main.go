package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	docs "github.com/Bulut-Bilisimciler/go-shift-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// import local packages /logger
	"github.com/Bulut-Bilisimciler/go-shift-service/logger"

	"github.com/Bulut-Bilisimciler/go-shift-service/config"
	"github.com/Bulut-Bilisimciler/go-shift-service/handlers"
)

// Path: ShiftService
// @Title ShiftService API
// @Description tt.app.ShiftService : microservice for layout.
// @Version 1.0.0
// @Schemes http https
// @BasePath /api-shifts

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

var isConfigSuccess = false
var equals string = strings.Repeat("=", 50)

var (
	TRACE *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	FATAL *log.Logger
)

const (
	// server name
	APP_NAME = "localhost:9097/api-shifts/"
	// server description
	APP_DESCRIPTION = "localhost:9097/api-shifts/ : microservice for layout."
)

func main() {
	env := config.C.App.Env
	port := config.C.App.Port
	logger.INFO.Printf("INIT: %s env: %s, port: %s", APP_NAME, env, port)

	// create 3th party connections
	/*
		s3Config := &models.S3Config{
			Endpoint:  config.C.Cdn.Endpoint,
			Region:    config.C.Cdn.Region,
			Bucket:    config.C.Cdn.Bucket,
			AccessKey: config.C.Cdn.AccessKey,
			SecretKey: config.C.Cdn.SecretKey,
		}
	*/
	// s3sess := NewS3Session(s3Config)

	router := gin.New()
	router.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	inAppCache := NewInAppCacheStore(time.Minute)
	cacheConn, cacheContext := NewRedisCacheConnection(config.C.Cache.Url)
	dbConn := NewPostgresDB(config.C.DB.Url)

	// create application service
	shiftsvc := handlers.NewShiftService(
		inAppCache,
		cacheConn,
		cacheContext,
		dbConn,
		// s3sess,
	)

	// check env and set gin mode
	gin.SetMode(gin.ReleaseMode)
	if env == "prod" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// check env and set swagger
	if !(env == "prod" || env == "production") {
		docs.SwaggerInfo.BasePath = handlers.API_PREFIX
		router.GET(handlers.API_PREFIX+"/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	shiftsvc.InitRouter(router)

	// start server
	logger.INFO.Println("INIT: Application " + APP_NAME + " started on port " + port)
	logger.INFO.Println(router.Run(":" + port))
}

// Initialize Application
func init() {
	logger.InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	isConfigSuccess = configureApplication()
	if !isConfigSuccess {
		logger.ERROR.Println("INIT: Application configuration failed. Please check your config file.")
		os.Exit(1)
	} else {
		logger.INFO.Println("INIT: Application configuration success.")
	}
}

// Configure Application
func configureApplication() bool {
	fmt.Println(equals)
	dir, err := os.Getwd()
	if err != nil {
		// log.Fatal("INIT: Cannot get current working directory os.Getwd()")
		logger.ERROR.Println("INIT: Cannot get current working directory os.Getwd()")
		return false
	} else {
		config.ReadConfig(dir)
		logger.INFO.Println("INIT: Application configuration file read success.")
		return true
	}
}
