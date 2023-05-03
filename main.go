package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Bulut-Bilisimciler/go-shift-service/config"
	docs "github.com/Bulut-Bilisimciler/go-shift-service/docs"
	"github.com/Bulut-Bilisimciler/go-shift-service/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

const (
	// server name
	APP_NAME = "bb.app.SERVICE_NAME"
	// server description
	APP_DESCRIPTION = "bb.app.SERVICE_NAME : microservice for layout."
)

func main() {
	// parse application envs
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("INIT: Cannot get current working directory os.Getwd()")
	}
	config.ReadConfig(dir)

	// init env
	env := config.C.App.Env
	port := config.C.App.Port
	// log env and port like "bb.app.SERVICE_NAME env: dev, port: 9097"
	log.Printf("INIT: %s env: %s, port: %s", APP_NAME, env, port)

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
	inAppCache := NewInAppCacheStore(time.Minute)
	cacheConn, cacheContext := NewRedisCacheConnection(config.C.Cache.Url)
	db := NewPostgresDB(config.C.DB.Url)

	// create application service
	mysvc := handlers.NewShiftService(
		inAppCache,
		cacheConn,
		cacheContext,
		db,
		// s3sess,
	)

	// check env and set gin mode
	gin.SetMode(gin.ReleaseMode)
	if env == "prod" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	mysvc.InitRouter(router)

	// check env and set swagger
	if !(env == "prod" || env == "production") {
		docs.SwaggerInfo.BasePath = handlers.API_PREFIX
		router.GET(handlers.API_PREFIX+"/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// start server
	log.Println("INIT: Application " + APP_NAME + " started on port " + port)
	// fatal if error
	log.Fatal(router.Run(":" + port))
}
