package repository

import (
	"database/sql"

	"go-api/internal/model"
)

type SessionRepository struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

func (r *SessionRepository) FindByToken(token string) (*model.Session, error) {
	query := `
		SELECT id, user_id, token, ip_address, user_agent, expired_at, created_at
		FROM sessions
		WHERE token = ?
		LIMIT 1
	`

	session := &model.Session{}

	err := r.DB.QueryRow(query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.IPAddress,
		&session.UserAgent,
		&session.ExpiredAt,
		&session.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}