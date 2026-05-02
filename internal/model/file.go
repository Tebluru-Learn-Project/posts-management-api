package model

import "time"

type File struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id"`
	Disk         string    `json:"disk"`
	Bucket       *string   `json:"bucket,omitempty"`
	Path         string    `json:"path"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	MimeType     string    `json:"mime_type"`
	Extension    string    `json:"extension"`
	Size         uint64    `json:"size"`
	Visibility   string    `json:"visibility"`
	Category     string    `json:"category"`
	Checksum     *string   `json:"checksum,omitempty"`
	IsUsed       bool      `json:"is_used"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}