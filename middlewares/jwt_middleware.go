package middlewares

import (
	"errors"
	"net/http"
	"time"

	"buluttan/shift-service/config"
	"buluttan/shift-service/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtRequiredMiddleware(c *gin.Context) {
	// get auth header to get jwt
	var tokenstr string
	authHeader := c.GetHeader("Authorization")
	// get jwt from header or return 401
	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt required"})
		return
	}
	// get token from authHeader
	tokenstr = authHeader[7:]

	// decode the jwt with public key
	pub, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.C.Auth.JwtPub))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "jwt could not be verified due to internal error"})
		return
	}

	// decode the jwt
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return pub, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt is invalid"})
		return
	}

	// validate the JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt is invalid"})
		return
	}

	// before decoding the JWT, check if user verified at
	user, ok := claims["user"].(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt user claim is invalid"})
		return
	}

	// user.verified_at is time string or null check and convert it
	var verifiedAt *time.Time
	if user["verified_at"] != nil {
		timeValue, err := time.Parse(time.RFC3339, user["verified_at"].(string))
		if err != nil {
			verifiedAt = nil
		}
		verifiedAt = &timeValue
	}

	u := models.JwtResponse{
		AuthId:     int64(user["auth_id"].(float64)),
		Slug:       user["slug"].(string),
		Status:     int(user["status"].(float64)),
		Role:       user["role"].(string),
		VerifiedAt: verifiedAt,
	}

	// If everything is ok, continue
	c.Set("user", &u)
	c.Next()
}
