package model

import "time"

type OTP struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Code      string    `json:"code"`
	Purpose   string    `json:"purpose"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}