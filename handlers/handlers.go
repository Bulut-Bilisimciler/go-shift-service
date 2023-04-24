package handlers

import (
	"context"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	API_PREFIX = "/api-SERVICE_NAME"
	RN_PREFIX  = "bbrn:::SERVICENAME:::"
)

// MySuperService is a struct for auth core
type MySuperService struct {
	inAppCache   *persistence.InMemoryStore
	cache        *redis.Client
	cacheContext context.Context
	db           *gorm.DB
	// s3sess       *session.Session
}

func NewMySuperService(
	inAppCache *persistence.InMemoryStore,
	cache *redis.Client,
	cacheContext context.Context,
	db *gorm.DB,
	// s3sess *session.Session,
) *MySuperService {
	return &MySuperService{
		inAppCache:   inAppCache,
		cache:        cache,
		cacheContext: cacheContext,
		// s3sess:       s3sess,
	}
}

type RespondJson struct {
	Status  bool        `json:"status" example:"true"`
	Intent  string      `json:"intent" example:"bbrn:::SERVICE_NAME:::/upload"`
	Message interface{} `json:"message" example:nil`
}

func respondJson(ctx *gin.Context, code int, intent string, message interface{}, err error) {
	if err == nil {
		ctx.JSON(code, RespondJson{
			Status:  true,
			Intent:  intent,
			Message: message,
		})
	} else {
		ctx.JSON(code, RespondJson{
			Status:  false,
			Intent:  intent,
			Message: err.Error(),
		})
	}
}

func (mss *MySuperService) InitRouter(r *gin.Engine) {
	// -- my service routes (group)
	v1 := r.Group(API_PREFIX)

	// --- File upload templates

	// get uploaded file
	v1.GET("/uploads/:filename", mss.HandleRetrieveFile)
	// upload file (one per request)
	v1.POST("/upload", func(ctx *gin.Context) {
		code, data, err := mss.HandleUploadFile(ctx)
		respondJson(ctx, code, RN_PREFIX+"/upload", data, err)
	})

	// --- CRUD Templates

}
