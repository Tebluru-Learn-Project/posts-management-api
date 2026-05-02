package repository

import (
	"database/sql"
	"go-api/internal/model"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) FindUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password, role, is_active, last_login, token_expiry, created_at, updated_at FROM users WHERE email = ?`
	user := &model.User{}
	
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.LastLogin, &user.TokenExpiry, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *AuthRepository) CreateSession(session *model.Session) error {
	query := `INSERT INTO sessions (user_id, token, ip_address, user_agent, expired_at, last_activity) VALUES (?, ?, ?, ?, ?, ?)`
	
	_, err := r.DB.Exec(query, session.UserID, session.Token, session.IPAddress, session.UserAgent, session.ExpiredAt, session.LastActivity)
	
	return err
}

func (r *AuthRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (name, email, password, is_active)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.IsActive)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint64(id)
	return nil
}

func (r *AuthRepository) ActivateUser(userID uint64) error {
	query := `
		UPDATE users
		SET is_active = true,
			verified = true,
		    updated_at = NOW()
		WHERE id = ?
	`

	_, err := r.DB.Exec(query, userID)
	return err
}

func (r *AuthRepository) FindUserByID(id uint64) (*model.User, error) {
	query := `
		SELECT id, name, email, role, is_active, created_at, updated_at
		FROM users
		WHERE id = ? AND is_active = true
		LIMIT 1
	`

	user := &model.User{}

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}