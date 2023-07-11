package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	// cache
	"buluttan/shift-service/models"

	"github.com/go-redis/redis/v8"

	// postgres
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// in-app-cache
	"github.com/gin-contrib/cache/persistence"

	// s3
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewS3Session(conf *models.S3Config) *session.Session {
	// open s3 session
	creds := credentials.NewStaticCredentials(
		conf.AccessKey,
		conf.SecretKey,
		"",
	)
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(conf.Endpoint),
		Region:           aws.String(conf.Region),
		Credentials:      creds,
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatalln("error_create_s3_session: ", err)
	}
	return sess
}

func NewPostgresDB(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("INIT: error_postgres_initial_conn: ", err)
	}

	return db
}

func NewRedisCacheConnection(redisUrl string) (*redis.Client, context.Context) {
	// redis://username:password@host:port/db
	_, username, password, host, port, db := UrlStringToOptions(redisUrl)
	// convert db to int
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalln("error_parse_redis_db_to_int: db option is not a number ", err)
	}

	rContext := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Username: username,
		Password: password,
		DB:       dbInt,
	})
	// ping redis for check connection
	_, err = rdb.Ping(rContext).Result()
	if err != nil {
		log.Fatalln("error_redis_initial_ping: ", err)
	}
	return rdb, rContext
}

func NewInAppCacheStore(defaultTTL time.Duration) *persistence.InMemoryStore {
	return persistence.NewInMemoryStore(defaultTTL)
}

func UrlStringToOptions(url string) (string, string, string, string, string, string) {
	// parse redis url to username password and host and port and db
	// app_name://username:password@host:port/db

	// split all options (app_name, username, password, host, port, db)
	options := strings.Split(url, "://")
	// split protocol and info
	protocol := options[0]
	info := options[1]
	// split info to username, password, host, port, db
	infoOptions := strings.Split(info, "@")
	// split username and password
	usernamePassword := infoOptions[0]
	hostPortDb := infoOptions[1]
	// split username and password
	usernamePasswordOptions := strings.Split(usernamePassword, ":")
	username := usernamePasswordOptions[0]
	password := usernamePasswordOptions[1]
	// split host, port and db
	hostPortDbOptions := strings.Split(hostPortDb, "/")
	hostPort := hostPortDbOptions[0]
	db := hostPortDbOptions[1]
	// split host and port
	hostPortOptions := strings.Split(hostPort, ":")
	host := hostPortOptions[0]
	port := hostPortOptions[1]

	return protocol, username, password, host, port, db
}
