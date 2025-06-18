-- Rollback: remove bio field from users table

ALTER TABLE users DROP COLUMN bio;
