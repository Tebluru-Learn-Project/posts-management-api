package model

import "time"

type Session struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id"`
	Token        string    `json:"token"`
	IPAddress    *string   `json:"ip_address,omitempty"`
	UserAgent    *string   `json:"user_agent,omitempty"`
	ExpiredAt    time.Time `json:"expired_at"`
	LastActivity time.Time `json:"last_activity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}