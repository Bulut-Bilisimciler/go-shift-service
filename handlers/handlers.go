package handlers

import (
	"context"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	API_PREFIX = "/api-shifts"
	RN_PREFIX  = "bbrn:::shiftservice:::"
)

// MyShiftService is a struct for auth core
type ShiftService struct {
	inAppCache   *persistence.InMemoryStore
	cache        *redis.Client
	cacheContext context.Context
	db           *gorm.DB
	// s3sess       *session.Session
}

func NewShiftService(
	inAppCache *persistence.InMemoryStore,
	cache *redis.Client,
	cacheContext context.Context,
	db *gorm.DB,
	// s3sess *session.Session,
) *ShiftService {
	return &ShiftService{
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

func (mss *ShiftService) InitRouter(r *gin.Engine) {
	// -- my service routes (group)
	v1 := r.Group(API_PREFIX)

	// SHIFT SECTİON
	// create shift
	v1.POST("/shifts", func(ctx *gin.Context) {
		code, data, err := mss.HandleCreateShift(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shifts", data, err)
	})

	v1.GET("/shifts", func(ctx *gin.Context) {
		code, data, err := mss.HandleGetShift(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shifts", data, err)
	})

}
