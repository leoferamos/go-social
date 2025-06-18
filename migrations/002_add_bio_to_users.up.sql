-- Migration: add bio field to users table

ALTER TABLE users ADD COLUMN bio VARCHAR(255) DEFAULT '' AFTER email;
