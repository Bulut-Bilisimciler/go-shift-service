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
		db:           db,
		// s3sess:       s3sess,
	}
}

type RespondJson struct {
	Status  bool        `json:"status"`
	Intent  string      `json:"intent"`
	Message interface{} `json:"message"`
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

func (ss *ShiftService) InitRouter(r *gin.Engine) {
	// -- my service routes (group)
	v1 := r.Group(API_PREFIX)

	// USERS

	// Create a new user
	v1.POST("/users", func(ctx *gin.Context) {
		code, data, err := ss.HandleCreateUser(ctx)
		respondJson(ctx, code, RN_PREFIX+"/users", data, err)
	})

	// Get all users
	v1.GET("/users", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetUsers(ctx)
		respondJson(ctx, code, API_PREFIX+"/users", data, err)
	})

	// Get User by id
	v1.GET("/users/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetUserById(ctx)
		respondJson(ctx, code, API_PREFIX+"/users/:id", data, err)
	})

	// Delete a specific user by id
	v1.DELETE("/users/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleDeleteUser(ctx)
		respondJson(ctx, code, RN_PREFIX+"/users", data, err)
	})

	// Update a specific user by id
	v1.PUT("/users/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleUpdateUser(ctx)
		respondJson(ctx, code, RN_PREFIX+"/users", data, err)
	})

	//DEMANDS

	// create a new demands
	v1.POST("/demands", func(ctx *gin.Context) {
		code, data, err := ss.HandleCreateDemand(ctx)
		respondJson(ctx, code, RN_PREFIX+"/demands", data, err)
	})
	// get all demands
	v1.GET("/demands", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetDemand(ctx)
		respondJson(ctx, code, API_PREFIX+"/demands", data, err)
	})
	// get demands by id
	v1.GET("/demands/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetDemandById(ctx)
		respondJson(ctx, code, API_PREFIX+"/demands/:id", data, err)
	})

	// SHIFTS

	// Create a new shift
	v1.POST("/shifts", func(ctx *gin.Context) {
		code, data, err := ss.HandleCreateShift(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shifts", data, err)
	})

	// Get all shifts
	v1.GET("/shifts", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetShift(ctx)
		respondJson(ctx, code, API_PREFIX+"/shifts", data, err)
	})

	// Shift by id
	v1.GET("/shifts/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetShiftById(ctx)
		respondJson(ctx, code, API_PREFIX+"/shifts/:id", data, err)
	})

	// Delete a specific shift by id
	v1.DELETE("/shifts/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleDeleteShift(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shifts", data, err)
	})

	// Update a specific shift by id
	v1.PUT("/shifts/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleUpdateShift(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shifts", data, err)
	})
	// SHIFT_PERIODS

	// Create a new shift period
	v1.POST("/shift-periods", func(ctx *gin.Context) {
		code, data, err := ss.HandleCreateShiftPeriod(ctx)
		respondJson(ctx, code, RN_PREFIX+"/shift-periods", data, err)
	})

	// Get all shift periods
	v1.GET("/shift-periods", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetShiftPeriods(ctx)
		respondJson(ctx, code, API_PREFIX+"/shift-periods", data, err)
	})

	// Get a shift period by ID
	v1.GET("/shift-periods/:id", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetShiftPeriodByID(ctx)
		respondJson(ctx, code, API_PREFIX+"/shift-periods/:id", data, err)

	})

	v1.GET("/shift-periods", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetShiftPeriod(ctx)
		respondJson(ctx, code, API_PREFIX+"/shift-periods", data, err)
	})

	v1.GET("/demands", func(ctx *gin.Context) {
		code, data, err := ss.HandleGetDemand(ctx)
		respondJson(ctx, code, API_PREFIX+"/demands", data, err)
	})

}
