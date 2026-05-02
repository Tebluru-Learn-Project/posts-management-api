ALTER TABLE profiles
DROP FOREIGN KEY fk_profiles_avatar_file_id,
DROP COLUMN avatar_file_id;