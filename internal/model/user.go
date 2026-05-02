package model

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      *string       `json:"role"`
	Verified  bool      `json:"verified"`
	IsActive  bool      `json:"is_active"`
	LastLogin *time.Time `json:"last_login"`
	TokenExpiry *time.Time `json:"token_expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}