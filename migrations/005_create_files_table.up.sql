CREATE TABLE IF NOT EXISTS files (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,

    disk VARCHAR(50) NOT NULL DEFAULT 'local',
    bucket VARCHAR(100) NULL,
    path VARCHAR(255) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    extension VARCHAR(20) NOT NULL,
    size BIGINT UNSIGNED NOT NULL,

    visibility ENUM('public', 'private') NOT NULL DEFAULT 'public',
    category VARCHAR(50) NOT NULL DEFAULT 'general',

    checksum VARCHAR(64) NULL,
    is_used BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_files_user_id (user_id),
    INDEX idx_files_category (category),
    INDEX idx_files_disk (disk),
    INDEX idx_files_is_used (is_used),

    CONSTRAINT fk_files_user_id
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);