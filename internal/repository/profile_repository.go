package repository

import (
	"database/sql"

	"go-api/internal/model"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) FindByUserID(userID uint64) (*model.Profile, error) {
	query := `
		SELECT id, user_id, first_name, last_name, phone, avatar_file_id, gender,
		       birth_date, bio, country, province, city, district,
		       postal_code, address, created_at, updated_at
		FROM profiles
		WHERE user_id = ?
		LIMIT 1
	`

	profile := &model.Profile{}

	err := r.DB.QueryRow(query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Phone,
		&profile.AvatarFileID,
		&profile.Gender,
		&profile.BirthDate,
		&profile.Bio,
		&profile.Country,
		&profile.Province,
		&profile.City,
		&profile.District,
		&profile.PostalCode,
		&profile.Address,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *ProfileRepository) Update(profile *model.Profile) error {
	query := `
		UPDATE profiles
		SET first_name = ?, last_name = ?, phone = ?, avatar_file_id = ?, gender = ?,
		    birth_date = ?, bio = ?, country = ?, province = ?, city = ?,
		    district = ?, postal_code = ?, address = ?, updated_at = NOW()
		WHERE user_id = ?
	`

	_, err := r.DB.Exec(
		query,
		profile.FirstName,
		profile.LastName,
		profile.Phone,
		profile.AvatarFileID,
		profile.Gender,
		profile.BirthDate,
		profile.Bio,
		profile.Country,
		profile.Province,
		profile.City,
		profile.District,
		profile.PostalCode,
		profile.Address,
		profile.UserID,
	)

	return err
}

func (r *ProfileRepository) Create(profile *model.Profile) error {
	query := `
		INSERT INTO profiles (
			user_id, first_name, last_name, phone, avatar_file_id, gender,
			birth_date, bio, country, province, city, district,
			postal_code, address
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		profile.UserID,
		profile.FirstName,
		profile.LastName,
		profile.Phone,
		profile.AvatarFileID,
		profile.Gender,
		profile.BirthDate,
		profile.Bio,
		profile.Country,
		profile.Province,
		profile.City,
		profile.District,
		profile.PostalCode,
		profile.Address,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	profile.ID = uint64(id)
	return nil
}