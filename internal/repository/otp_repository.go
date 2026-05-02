package repository

import (
	"database/sql"
	"go-api/internal/model"
)

type OTPRepository struct {
	DB *sql.DB
}

func NewOTPRepository(db *sql.DB) *OTPRepository {
	return &OTPRepository{DB: db}
}

func (r *OTPRepository) CreateOTP(otp *model.OTP) error {
	query := `
		INSERT INTO otps (user_id, code, purpose, expired_at)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.DB.Exec(query, otp.UserID, otp.Code, otp.Purpose, otp.ExpiredAt)
	return err
}

func (r *OTPRepository) FindByUserIDAndCode(userID uint64, code string) (*model.OTP, error) {
	query := `
		SELECT id, user_id, code, purpose, expired_at, created_at
		FROM otps
		WHERE user_id = ? AND code = ? AND purpose = 'register'
		LIMIT 1
	`

	otp := &model.OTP{}

	err := r.DB.QueryRow(query, userID, code).Scan(
		&otp.ID,
		&otp.UserID,
		&otp.Code,
		&otp.Purpose,
		&otp.ExpiredAt,
		&otp.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return otp, nil
}

func (r *OTPRepository) DeleteByID(id uint64) error {
	query := `DELETE FROM otps WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}