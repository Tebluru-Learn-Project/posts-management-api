ALTER TABLE profiles
ADD COLUMN avatar_file_id BIGINT UNSIGNED NULL AFTER user_id,
ADD CONSTRAINT fk_profiles_avatar_file_id
    FOREIGN KEY (avatar_file_id) REFERENCES files(id)
    ON DELETE SET NULL;