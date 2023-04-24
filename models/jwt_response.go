package models

import "time"

// JWT Response from main auth microservice
type JwtResponse struct {
	AuthId     int64      `json:"auth_id"`
	Slug       string     `json:"slug"`
	Role       string     `json:"role"`
	Status     int        `json:"status"`
	VerifiedAt *time.Time `json:"verified_at"`
}
