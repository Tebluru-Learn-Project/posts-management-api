package repository

import (
	"database/sql"

	"go-api/internal/model"
)

type FileRepository struct {
	DB *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{DB: db}
}

func (r *FileRepository) Create(file *model.File) error {
	query := `
		INSERT INTO files (
			user_id, disk, bucket, path, filename, original_name,
			mime_type, extension, size, visibility, category, checksum, is_used
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		file.UserID,
		file.Disk,
		file.Bucket,
		file.Path,
		file.Filename,
		file.OriginalName,
		file.MimeType,
		file.Extension,
		file.Size,
		file.Visibility,
		file.Category,
		file.Checksum,
		file.IsUsed,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	file.ID = uint64(id)
	return nil
}

func (r *FileRepository) FindByID(id uint64) (*model.File, error) {
	query := `
		SELECT id, user_id, disk, bucket, path, filename, original_name,
		       mime_type, extension, size, visibility, category, checksum,
		       is_used, created_at, updated_at
		FROM files
		WHERE id = ?
		LIMIT 1
	`

	file := &model.File{}

	err := r.DB.QueryRow(query, id).Scan(
		&file.ID,
		&file.UserID,
		&file.Disk,
		&file.Bucket,
		&file.Path,
		&file.Filename,
		&file.OriginalName,
		&file.MimeType,
		&file.Extension,
		&file.Size,
		&file.Visibility,
		&file.Category,
		&file.Checksum,
		&file.IsUsed,
		&file.CreatedAt,
		&file.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return file, nil
}